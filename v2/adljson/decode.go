package adljson

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"sync"

	goadl "github.com/adl-lang/goadl_rt/v2"
)

type Decoder[T any] struct {
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
	texpr goadl.ATypeExpr[T],
	dres goadl.Resolver,
) *Decoder[T] {
	binding := buildDecodeBinding(dres, texpr.Value, make(boundDecodeTypeParams))
	return &Decoder[T]{
		r:       r,
		binding: binding,
	}
}

func (dec *Decoder[T]) Decode(v *T) error {
	// for now encode into a Go any and pull pieces out into ADL decls
	var v0 any
	jd := json.NewDecoder(dec.r)
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

type boundDecodeTypeParams map[string]decodeFunc

func texprKey(te goadl.TypeExpr) string {
	ref := goadl.Handle_TypeRef[string](
		te.TypeRef.Branch,
		func(primitive string) string {
			if len(te.Parameters) == 0 {
				return primitive + ":"
			}
			return primitive + ":" + texprKey(te.Parameters[0])
		},
		func(typeParam string) string {
			if len(te.Parameters) != 0 {
				panic("type params cannot have params")
			}
			return "[" + typeParam + "]"
		},
		func(reference goadl.ScopedName) string {
			sn := reference.ModuleName + "." + reference.Name + "::"
			for i := range te.Parameters {
				if i != 0 {
					sn = sn + ","
				}
				sn = sn + texprKey(te.Parameters[i])
			}
			return sn
		},
	)
	return ref
}

var decoderCache sync.Map // map[reflect.Type]decodeFunc

func buildDecodeBinding(
	dres goadl.Resolver,
	texpr goadl.TypeExpr,
	boundTypeParams boundDecodeTypeParams,
) decodeFunc {
	// taken from golang stdlib src/encoding/json/encode.go
	key := texprKey(texpr)
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
	dres goadl.Resolver,
	texpr goadl.TypeExpr,
	boundTypeParams boundDecodeTypeParams,
) decodeFunc {
	return goadl.Handle_TypeRef[decodeFunc](
		texpr.TypeRef.Branch,
		func(primitive string) decodeFunc {
			return primitiveDecodeBinding(dres, primitive, texpr.Parameters, boundTypeParams)
		},
		func(typeParam string) decodeFunc {
			return boundTypeParams[typeParam]
		},
		func(reference goadl.ScopedName) decodeFunc {
			ast := dres.Resolve(reference)
			return goadl.Handle_DeclType[decodeFunc](
				ast.SD.Decl.Type.Branch,
				func(struct_ goadl.Struct) decodeFunc {
					return structDecodeBinding(dres, struct_, texpr.Parameters, boundTypeParams)
				},
				func(union_ goadl.Union) decodeFunc {
					if isEnum(union_) {
						return enumDecodeBinding(dres, union_, texpr.Parameters, boundTypeParams)
					}
					return unionDecodeBinding(dres, ast.TypeMap, union_, texpr.Parameters, boundTypeParams)
				},
				func(type_ goadl.TypeDef) decodeFunc {
					return typedefDecodeBinding(dres, type_, texpr.Parameters, boundTypeParams)
				},
				func(newtype_ goadl.NewType) decodeFunc {
					return newtypeDecodeBinding(dres, newtype_, texpr.Parameters, boundTypeParams)
				},
				nil,
			)
		},
	)
}

func primitiveDecodeBinding(
	dres goadl.Resolver,
	primitive string,
	typeExpr []goadl.TypeExpr,
	boundTypeParams boundDecodeTypeParams,
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
				return fmt.Errorf("path: %v, received value cannot be convert to expected type. expected %s:%v received %v %+#v",
					ctx.path, primitive, ds.v.Type(), ro.Kind(), ro.Interface())
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
	dres goadl.Resolver,
	struct_ goadl.Struct,
	typeExpr []goadl.TypeExpr,
	boundTypeParams boundDecodeTypeParams,
) decodeFunc {
	newBoundTypeParams := make(boundDecodeTypeParams)
	for i, paramName := range struct_.TypeParams {
		newBoundTypeParams[paramName] = buildDecodeBinding(dres, typeExpr[i], boundTypeParams)
	}
	fieldJB := make([]decodeFunc, 0, len(struct_.Fields))
	for _, field := range struct_.Fields {
		jb := buildDecodeBinding(dres, field.TypeExpr, newBoundTypeParams)
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
					if f.Default.Nothing != nil {
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
	dres goadl.Resolver,
	union_ goadl.Union,
	typeExpr []goadl.TypeExpr,
	boundTypeParams boundDecodeTypeParams,
) decodeFunc {
	panic("unimplemented")
}

type boundDecField struct {
	decodeFunc decodeFunc
	field      goadl.Field
}

func unionDecodeBinding(
	dres goadl.Resolver,
	typeMap map[string]reflect.Type,
	union_ goadl.Union,
	typeExpr []goadl.TypeExpr,
	boundTypeParams boundDecodeTypeParams,
) decodeFunc {
	decMap := make(map[string]boundDecField)
	for _, f := range union_.Fields {
		decMap[f.SerializedName] = boundDecField{
			buildDecodeBinding(dres, f.TypeExpr, boundTypeParams),
			f,
		}
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
	dres goadl.Resolver,
	type_ goadl.TypeDef,
	typeExpr []goadl.TypeExpr,
	boundTypeParams boundDecodeTypeParams,
) decodeFunc {
	return buildDecodeBinding(dres, type_.TypeExpr, boundTypeParams)
}

func newtypeDecodeBinding(
	dres goadl.Resolver,
	newtype_ goadl.NewType,
	typeExpr []goadl.TypeExpr,
	boundTypeParams boundDecodeTypeParams,
) decodeFunc {
	// TODO different default values
	return buildDecodeBinding(dres, newtype_.TypeExpr, boundTypeParams)
}
