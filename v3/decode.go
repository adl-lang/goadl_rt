package goadl

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"

	adlast "github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

type Decoder[T any] struct {
	r       io.Reader
	binding decodeFunc
}

type UncheckedDecoder struct {
	r       io.Reader
	binding decodeFunc
}

type decodeState struct {
	v reflect.Value
}

type ctxPath []string

func (cp ctxPath) String() string {
	return "[" + strings.Join(cp, ",") + "]"
}

type decContext struct {
	path ctxPath
}

type decodeFunc func(ctx decContext, ds *decodeState, v any) error

func NewDecoder[T any](
	r io.Reader,
	texpr ATypeExpr[T],
	dres Resolver,
) *Decoder[T] {
	binding := buildDecodeBinding(dres, texpr.Value, make(map[string]decodeFunc))
	return &Decoder[T]{
		r:       r,
		binding: binding,
	}
}

func NewDecoderUnchecked(
	r io.Reader,
	texpr adlast.TypeExpr,
	dres Resolver,
) *UncheckedDecoder {
	binding := buildDecodeBinding(dres, texpr, make(map[string]decodeFunc))
	return &UncheckedDecoder{
		r:       r,
		binding: binding,
	}
}

func (dec *Decoder[T]) Decode(v *T) error {
	// for now encode into a Go any and pull pieces out into ADL decls
	var v0 any
	jd := json.NewDecoder(dec.r)
	// jd.UseNumber()
	err := jd.Decode(&v0)
	if err != nil {
		return err
	}
	ds := decodeState{
		v: reflect.ValueOf(v).Elem(),
	}
	dc := decContext{
		path: []string{"$"},
	}
	return dec.binding(dc, &ds, v0)
}

func (dec *UncheckedDecoder) Decode(v any) error {
	// for now encode into a Go any and pull pieces out into ADL decls
	var v0 any
	jd := json.NewDecoder(dec.r)
	err := jd.Decode(&v0)
	if err != nil {
		return err
	}
	rv := reflect.ValueOf(v)
	rv = unwrap(rv)
	ds := decodeState{
		v: rv,
	}
	dc := decContext{
		path: []string{"$"},
	}
	return dec.binding(dc, &ds, v0)
}

// type map[string]decodeFunc map[string]decodeFunc

func texprDecKey(te adlast.TypeExpr) string {
	ref := adlast.Handle_TypeRef[string](
		te.TypeRef.Branch,
		func(primitive string) string {
			if len(te.Parameters) == 0 {
				return primitive + ":"
			}
			return primitive + ":" + texprDecKey(te.Parameters[0])
		},
		func(typeParam string) string {
			panic("type params cannot have params")
			// if len(te.Parameters) != 0 {
			// 	panic("type params cannot have params")
			// }
			// return "[" + typeParam + "]"
		},
		func(reference adlast.ScopedName) string {
			sn := reference.ModuleName + "." + reference.Name + "::"
			for i := range te.Parameters {
				if i != 0 {
					sn = sn + ","
				}
				sn = sn + texprDecKey(te.Parameters[i])
			}
			return sn
		},
		nil,
	)
	return ref
}

var decoderCache sync.Map // map[reflect.Type]decodeFunc

func buildDecodeBinding(
	dres Resolver,
	texpr adlast.TypeExpr,
	boundTypeParams map[string]decodeFunc,
) decodeFunc {
	// taken from golang stdlib src/encoding/json/encode.go
	key := texprDecKey(texpr)
	if fi, ok := decoderCache.Load(key); ok {
		return fi.(decodeFunc)
	}
	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  decodeFunc
	)
	wg.Add(1)
	fi, loaded := decoderCache.LoadOrStore(key, decodeFunc(func(ctx decContext, e *decodeState, v any) error {
		wg.Wait()
		return f(ctx, e, v)
	}))
	if loaded {
		return fi.(decodeFunc)
	}

	// Compute the real encoder and replace the indirect func with it.
	f = buildNewDecodeBinding(dres, texpr, boundTypeParams)
	wg.Done()
	decoderCache.Store(key, f)
	return f
}

func buildNewDecodeBinding(
	dres Resolver,
	texpr adlast.TypeExpr,
	boundTypeParams map[string]decodeFunc,
) decodeFunc {
	return adlast.Handle_TypeRef[decodeFunc](
		texpr.TypeRef.Branch,
		func(primitive string) decodeFunc {
			return primitiveDecodeBinding(dres, primitive, texpr.Parameters, boundTypeParams)
		},
		func(typeParam string) decodeFunc {
			bf := boundTypeParams[typeParam]
			if bf == nil {
				panic(fmt.Errorf("bg == nil - typeParam:%v boundTypeParams:%v", typeParam, boundTypeParams))
			}
			return bf
		},
		func(reference adlast.ScopedName) decodeFunc {
			ast := dres.Resolve(reference)
			if ast == nil {
				panic(fmt.Errorf("cannot find %v", reference))
			}
			if ast.SD.Decl.Type_.Branch == nil {
				panic(fmt.Errorf("nil branch %v\n%+v", reference, ast))

			}

			fbind, tbind := createDecBoundTypeParams(dres, TypeParamsFromDecl(ast.SD.Decl), texpr.Parameters, boundTypeParams)

			return adlast.Handle_DeclType[decodeFunc](
				ast.SD.Decl.Type_.Branch,
				func(struct_ adlast.Struct) decodeFunc {
					return structDecodeBinding(dres, struct_, fbind, tbind)
				},
				func(union_ adlast.Union) decodeFunc {
					if isEnum(union_) {
						return enumDecodeBinding(dres, union_, ast.TypeMap)
					}
					return unionDecodeBinding(dres, ast.TypeMap, union_, fbind, tbind)
				},
				func(type_ adlast.TypeDef) decodeFunc {
					return typedefDecodeBinding(dres, type_, fbind, tbind)
				},
				func(newtype_ adlast.NewType) decodeFunc {
					return newtypeDecodeBinding(dres, newtype_, fbind, tbind)
				},
				nil,
			)
		},
		nil,
	)
}

func primitiveDecodeBinding(
	dres Resolver,
	primitive string,
	typeExpr []adlast.TypeExpr,
	boundTypeParams map[string]decodeFunc,
) decodeFunc {
	switch primitive {
	case "Int8", "Int16", "Int32", "Int64",
		"Word8", "Word16", "Word32", "Word64",
		"Bool",
		"Float",
		"Double",
		"String":
		return func(ctx decContext, ds *decodeState, v any) error {
			// fmt.Println(v, ds.v)
			ro := reflect.ValueOf(v)
			if !ro.CanConvert(ds.v.Type()) {
				return fmt.Errorf("path: %v, received value cannot be convert to expected type. expected %s:%v received type:'%v' val:'%+#v'",
					ctx.path, ds.v.Type(), primitive, ro.Kind(), ro.Interface())
			}
			ro = ro.Convert(ds.v.Type())
			ds.v.Set(ro)
			return nil
		}
	// case "ByteVector":
	case "Void":
		return func(ctx decContext, ds *decodeState, v any) error {
			return nil
		}
	case "Json":
		return func(ctx decContext, ds *decodeState, v any) error {
			if v == nil {
				return nil
			}
			ds.v.Set(reflect.ValueOf(v))
			return nil
		}
	case "Vector":
		elementBinding := buildDecodeBinding(dres, typeExpr[0], boundTypeParams)
		return func(ctx decContext, ds *decodeState, v any) error {
			rv := reflect.ValueOf(v)
			l := rv.Len()
			newSlice := reflect.MakeSlice(ds.v.Type(), l, l)
			for i := 0; i < l; i++ {
				ds0 := decodeState{v: newSlice.Index(i)}
				ctx0 := decContext{
					path: append(ctx.path, "["+strconv.Itoa(i)+"]"),
				}
				err := elementBinding(ctx0, &ds0, rv.Index(i).Interface())
				if err != nil {
					return err
				}
			}
			ds.v.Set(newSlice)

			return nil
		}
	case "StringMap":
		elementBinding := buildDecodeBinding(dres, typeExpr[0], boundTypeParams)
		return func(ctx decContext, ds *decodeState, v any) error {
			newM := reflect.MakeMap(ds.v.Type())
			vT := ds.v.Type().Elem()
			m := reflect.ValueOf(v)
			iter := m.MapRange()
			for iter.Next() {
				k := iter.Key()
				v0 := iter.Value()
				ds0 := decodeState{v: reflect.New(vT).Elem()}
				ctx0 := decContext{
					path: append(ctx.path, k.String()+":"),
				}
				err := elementBinding(ctx0, &ds0, v0.Interface())
				if err != nil {
					return err
				}
				newM.SetMapIndex(k, ds0.v)
			}
			ds.v.Set(newM)
			return nil
		}
	case "Nullable":
		elementBinding := buildDecodeBinding(dres, typeExpr[0], boundTypeParams)
		return func(ctx decContext, ds *decodeState, v any) error {
			if v == nil {
				return nil
			}
			// ds0 := decodeState{v: ds.v}
			ds0 := decodeState{v: reflect.New(ds.v.Type().Elem()).Elem()}
			err := elementBinding(ctx, &ds0, v)
			if err != nil {
				return err
			}
			ds.v.Set(ds0.v.Addr())
			return nil
		}
	}
	panic("unimplemented")
}

func structDecodeBinding(
	dres Resolver,
	struct_ adlast.Struct,
	fbind map[string]decodeFunc,
	tbind map[string]adlast.TypeExpr,
) decodeFunc {
	fieldJB := make([]decodeFunc, 0, len(struct_.Fields))
	for _, field := range struct_.Fields {
		// field.TypeExpr.Parameters
		monoTe := SubstituteTypeBindings(tbind, field.TypeExpr)
		jb := buildDecodeBinding(dres, monoTe, fbind)
		fieldJB = append(fieldJB, jb)
	}
	return func(ctx decContext, ds *decodeState, v any) error {
		switch t := v.(type) {
		case map[string]any:
			for i, f := range struct_.Fields {
				if v0, ok := t[f.SerializedName]; ok {
					ds0 := decodeState{}
					ds0.v = ds.v.Field(i)
					ctx0 := decContext{
						path: append(ctx.path, f.Name),
					}
					err := fieldJB[i](ctx0, &ds0, v0)
					if err != nil {
						return err
					}
				} else {
					if _, ok := any(f.Default).(types.Maybe_Nothing); ok {
						return fmt.Errorf("path %v, required field missing '%v'", ctx.path, f.SerializedName)
					}
					// TODO set from default
				}
			}
			return nil
		}
		return fmt.Errorf("path %v, struct: expect an object received %v '%v'", ctx.path, reflect.TypeOf(v), v)
	}
}

func enumDecodeBinding(
	dres Resolver,
	union_ adlast.Union,
	typeMap map[string]reflect.Type,
) decodeFunc {
	decMap := make(map[string]boundDecField)
	for _, f := range union_.Fields {
		bf := boundDecField{
			buildDecodeBinding(dres, Texpr_Void().Value, map[string]decodeFunc{}),
			f,
		}
		if bf.decodeFunc == nil {
			panic(fmt.Errorf("decodeFunc == nil - %#+v", f.TypeExpr))
		}
		decMap[f.SerializedName] = bf
	}
	return func(ctx decContext, ds *decodeState, v any) error {
		var (
			key string
			val any
		)
		switch t := v.(type) {
		case string:
			key = t
			val = nil
		case map[string]any:
			if len(t) != 1 {
				return fmt.Errorf("path: %v, expect an object with one and only element received %v", ctx.path, len(t))
			}
			for k0, v0 := range t {
				key = k0
				val = v0
			}
		default:
			return fmt.Errorf("path: %v, union: expect an object received %v '%v'", ctx.path, reflect.TypeOf(v), v)
		}

		if bf, ok := decMap[key]; ok {
			if typ, ok := typeMap[key]; ok {
				vn := reflect.New(typ)
				ds0 := decodeState{
					v: vn.Elem().Field(0),
				}
				ctx0 := decContext{
					path: append(ctx.path, key),
				}
				if bf.decodeFunc == nil {
					panic(fmt.Errorf("path: %v, decodeFunc == nil '%v'\n%+v", ctx.path, key, decMap))
				}
				err := bf.decodeFunc(ctx0, &ds0, val)
				if err != nil {
					return err
				}
				r0 := ds.v // for top level Elem() is already called
				r0 = r0.Field(0)
				// r0 = r0.Field(0)
				r0.Set(vn.Elem())
				return nil
			} else {
				return fmt.Errorf("path: %v, unexpected branch - no type registered '%v'", ctx.path, key)
			}
		} else {
			return fmt.Errorf("path: %v, unexpected branch '%v'", ctx.path, key)
		}
	}
}

type boundDecField struct {
	decodeFunc decodeFunc
	field      adlast.Field
}

func unionDecodeBinding(
	dres Resolver,
	typeMap map[string]reflect.Type,
	union_ adlast.Union,
	fbind map[string]decodeFunc,
	tbind map[string]adlast.TypeExpr,
) decodeFunc {
	decMap := make(map[string]boundDecField)
	for _, f := range union_.Fields {
		monoTe := SubstituteTypeBindings(tbind, f.TypeExpr)
		bf := boundDecField{
			buildDecodeBinding(dres, monoTe, fbind),
			f,
		}
		if bf.decodeFunc == nil {
			panic(fmt.Errorf("decodeFunc == nil - %#+v", f.TypeExpr))
		}
		decMap[f.SerializedName] = bf
	}
	return func(ctx decContext, ds *decodeState, v any) error {
		var (
			key string
			val any
		)

		switch t := v.(type) {
		case string:
			key = t
			val = nil
		case map[string]any:
			if len(t) != 1 {
				return fmt.Errorf("path: %v, expect an object with one and only element received %v", ctx.path, len(t))
			}
			for k0, v0 := range t {
				key = k0
				val = v0
			}
		default:
			return fmt.Errorf("path: %v, union: expect an object received %v '%v'", ctx.path, reflect.TypeOf(v), v)
		}

		if bf, ok := decMap[key]; ok {
			if typ, ok := typeMap[key]; ok {
				vn := reflect.New(typ)
				ds0 := decodeState{
					v: vn.Elem().Field(0),
				}
				ctx0 := decContext{
					path: append(ctx.path, key),
				}
				if bf.decodeFunc == nil {
					panic(fmt.Errorf("path: %v, decodeFunc == nil '%v'\n%+v", ctx.path, key, decMap))

				}
				err := bf.decodeFunc(ctx0, &ds0, val)
				if err != nil {
					return err
				}
				r0 := ds.v // for top level Elem() is already called
				r0 = r0.Field(0)
				// r0 = r0.Field(0)
				r0.Set(vn.Elem())
				return nil
			} else {
				return fmt.Errorf("path: %v, unexpected branch - no type registered '%v'", ctx.path, key)
			}
		} else {
			return fmt.Errorf("path: %v, unexpected branch '%v'", ctx.path, key)
		}

	}
}

func typedefDecodeBinding(
	dres Resolver,
	type_ adlast.TypeDef,
	fbind map[string]decodeFunc,
	tbind map[string]adlast.TypeExpr,
) decodeFunc {
	monoTe := SubstituteTypeBindings(tbind, type_.TypeExpr)
	return buildDecodeBinding(dres, monoTe, fbind)
}

func newtypeDecodeBinding(
	dres Resolver,
	newtype_ adlast.NewType,
	fbind map[string]decodeFunc,
	tbind map[string]adlast.TypeExpr,
) decodeFunc {
	monoTe := SubstituteTypeBindings(tbind, newtype_.TypeExpr)
	// TODO different default values
	return buildDecodeBinding(dres, monoTe, fbind)
}

func createDecBoundTypeParams(
	dres Resolver,
	paramNames []string,
	paramTypes []adlast.TypeExpr,
	boundTypeParams map[string]decodeFunc,
) (map[string]decodeFunc, map[string]adlast.TypeExpr) {
	fbind := map[string]decodeFunc{}
	tbind := map[string]adlast.TypeExpr{}
	for i, paramName := range paramNames {
		fbind[paramName] = buildDecodeBinding(dres, paramTypes[i], boundTypeParams)
		tbind[paramName] = paramTypes[i]
	}
	return fbind, tbind
}
