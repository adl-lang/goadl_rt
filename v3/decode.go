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

type JsonDecodeBinder[T any] struct {
	binding DecodeFunc
}

type JsonDecodeBinderUnchecked struct {
	binding DecodeFunc
}

func CreateJsonDecodeBinding[T any](
	texpr ATypeExpr[T],
	dres Resolver,
) JsonDecodeBinder[T] {
	return JsonDecodeBinder[T]{
		binding: buildDecodeBinding(dres, texpr.Value),
	}
}

func CreateUncheckedJsonDecodeBinding(
	texpr adlast.TypeExpr,
	dres Resolver,
) JsonDecodeBinderUnchecked {
	return JsonDecodeBinderUnchecked{
		binding: buildDecodeBinding(dres, texpr),
	}
}

func (jdb *JsonDecodeBinder[T]) Decode(
	r io.Reader,
	dst *T,
) error {
	// for now encode into a Go any and pull pieces out into ADL
	var src any
	jd := json.NewDecoder(r)
	// jd.UseNumber()
	err := jd.Decode(&src)
	if err != nil {
		return err
	}
	return jdb.DecodeFromAny(src, dst)
}

func (jdb *JsonDecodeBinderUnchecked) Decode(
	r io.Reader,
	dst any,
) error {
	// for now encode into a Go any and pull pieces out into ADL
	var src any
	jd := json.NewDecoder(r)
	err := jd.Decode(&src)
	if err != nil {
		return err
	}
	return jdb.DecodeFromAny(src, dst)
}

func (jdb *JsonDecodeBinder[T]) DecodeFromAny(
	src any,
	dst *T,
) error {
	ds := DecodeState{
		V:    unwrap(reflect.ValueOf(dst)),
		Path: []string{"$"},
	}
	return jdb.binding(&ds, src)
}

func (jdb *JsonDecodeBinderUnchecked) DecodeFromAny(
	src any,
	dst any,
) error {
	ds := DecodeState{
		V:    unwrap(reflect.ValueOf(dst)),
		Path: []string{"$"},
	}
	return jdb.binding(&ds, src)
}

func texprDecKey(
	te adlast.TypeExpr,
) string {
	sb := strings.Builder{}
	sb.Grow(100)
	var recurse func(te adlast.TypeExpr)
	recurse = func(te adlast.TypeExpr) {
		adlast.Handle_TypeRef[*struct{}](
			te.TypeRef.Branch,
			func(primitive string) *struct{} {
				sb.WriteString(primitive + ":")
				if len(te.Parameters) == 1 {
					recurse(te.Parameters[0])
				}
				return nil
			},
			func(typeParam string) *struct{} {
				panic(fmt.Errorf("%s", typeParam))
			},
			func(reference adlast.ScopedName) *struct{} {
				sb.WriteString(reference.ModuleName + "." + reference.Name + "::")
				for i := range te.Parameters {
					if i != 0 {
						sb.WriteString(",")
					}
					recurse(te.Parameters[i])
				}
				return nil
			},
			nil,
		)
	}
	recurse(te)
	return sb.String()
}

var decoderCache sync.Map // map[reflect.Type]DecodeFunc

func buildDecodeBinding(
	dres Resolver,
	texpr adlast.TypeExpr,
) DecodeFunc {
	key := texprDecKey(texpr)
	// taken from golang stdlib src/encoding/json/encode.go
	if fi, ok := decoderCache.Load(key); ok {
		return fi.(DecodeFunc)
	}
	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  DecodeFunc
	)
	wg.Add(1)
	fi, loaded := decoderCache.LoadOrStore(key, DecodeFunc(func(e *DecodeState, v any) error {
		wg.Wait()
		return f(e, v)
	}))
	if loaded {
		return fi.(DecodeFunc)
	}

	// Compute the real encoder and replace the indirect func with it.
	f = buildNewDecodeBinding(dres, texpr)
	wg.Done()
	decoderCache.Store(key, f)
	return f
}

func buildNewDecodeBinding(
	dres Resolver,
	texpr adlast.TypeExpr,
) DecodeFunc {
	return adlast.Handle_TypeRef[DecodeFunc](
		texpr.TypeRef.Branch,
		func(primitive string) DecodeFunc {
			return primitiveDecodeBinding(dres, primitive, texpr.Parameters)
		},
		func(typeParam string) DecodeFunc {
			panic(fmt.Errorf("typeParam:%v", typeParam))
		},
		func(reference adlast.ScopedName) DecodeFunc {
			ast := dres.Resolve(reference)
			if ast == nil {
				panic(fmt.Errorf("cannot find %v", reference))
			}
			if ast.Decl.Type_.Branch == nil {
				panic(fmt.Errorf("nil branch %v\n%+v", reference, ast))

			}
			tbind := CreateDecBoundTypeParams(TypeParamsFromDecl(ast.Decl), texpr.Parameters)
			// custom types
			if helper, has := dres.ResolveHelper(reference); has {
				typeparamDec := make([]DecodeFunc, len(texpr.Parameters))
				for i := range texpr.Parameters {
					monoTe := SubstituteTypeBindings(tbind, texpr.Parameters[i])
					typeparamDec[i] = buildDecodeBinding(dres, monoTe)
				}
				return helper.BuildDecodeFunc(typeparamDec...)
			}
			return adlast.Handle_DeclType[DecodeFunc](
				ast.Decl.Type_.Branch,
				func(struct_ adlast.Struct) DecodeFunc {
					return structDecodeBinding(dres, struct_, tbind)
				},
				func(union_ adlast.Union) DecodeFunc {
					if isEnum(union_) {
						return enumDecodeBinding(dres, union_)
					}
					return unionDecodeBinding(dres, union_, tbind)
				},
				func(type_ adlast.TypeDef) DecodeFunc {
					monoTe := SubstituteTypeBindings(tbind, type_.TypeExpr)
					return buildDecodeBinding(dres, monoTe)
				},
				func(newtype_ adlast.NewType) DecodeFunc {
					monoTe := SubstituteTypeBindings(tbind, newtype_.TypeExpr)
					// TODO different default values
					return buildDecodeBinding(dres, monoTe)
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
) DecodeFunc {
	switch primitive {
	case "Int8", "Int16", "Int32", "Int64",
		"Word8", "Word16", "Word32", "Word64",
		"Bool",
		"Float",
		"Double",
		"String":
		return func(ds *DecodeState, v any) error {
			// fmt.Println(v, ds.V)
			ro := reflect.ValueOf(v)
			if !ro.CanConvert(ds.V.Type()) {
				return fmt.Errorf("path: %v, received value cannot be convert to expected type. expected %s:%v received type:'%v' val:'%+#v'",
					ds.Path, ds.V.Type(), primitive, ro.Kind(), ro.Interface())
			}
			ro = ro.Convert(ds.V.Type())
			ds.V.Set(ro)
			return nil
		}
	// case "ByteVector":
	case "Void":
		return func(ds *DecodeState, v any) error {
			return nil
		}
	case "Json":
		return func(ds *DecodeState, v any) error {
			if v == nil {
				return nil
			}
			ds.V.Set(reflect.ValueOf(v))
			return nil
		}
	case "Vector":
		elementBinding := buildDecodeBinding(dres, typeExpr[0])
		return func(ds *DecodeState, v any) error {
			rv := reflect.ValueOf(v)
			l := rv.Len()
			newSlice := reflect.MakeSlice(ds.V.Type(), l, l)
			for i := 0; i < l; i++ {
				ds0 := DecodeState{
					V:    newSlice.Index(i),
					Path: append(ds.Path, "["+strconv.Itoa(i)+"]"),
				}
				err := elementBinding(&ds0, rv.Index(i).Interface())
				if err != nil {
					return err
				}
			}
			ds.V.Set(newSlice)

			return nil
		}
	case "StringMap":
		elementBinding := buildDecodeBinding(dres, typeExpr[0])
		return func(ds *DecodeState, v any) error {
			newM := reflect.MakeMap(ds.V.Type())
			vT := ds.V.Type().Elem()
			m := reflect.ValueOf(v)
			iter := m.MapRange()
			for iter.Next() {
				k := iter.Key()
				v0 := iter.Value()
				ds0 := DecodeState{
					V:    reflect.New(vT).Elem(),
					Path: append(ds.Path, k.String()+":"),
				}
				err := elementBinding(&ds0, v0.Interface())
				if err != nil {
					return err
				}
				newM.SetMapIndex(k, ds0.V)
			}
			ds.V.Set(newM)
			return nil
		}
	case "Nullable":
		elementBinding := buildDecodeBinding(dres, typeExpr[0])
		return func(ds *DecodeState, v any) error {
			if v == nil {
				return nil
			}
			// ds0 := DecodeState{v: ds.V}
			ds0 := DecodeState{
				V:    reflect.New(ds.V.Type().Elem()).Elem(),
				Path: ds.Path,
			}
			err := elementBinding(&ds0, v)
			if err != nil {
				return err
			}
			ds.V.Set(ds0.V.Addr())
			return nil
		}
	}
	panic("unimplemented")
}

func structDecodeBinding(
	dres Resolver,
	struct_ adlast.Struct,
	tbind []TypeBinding,
) DecodeFunc {
	fieldJB := make([]DecodeFunc, 0, len(struct_.Fields))
	for _, field := range struct_.Fields {
		monoTe := SubstituteTypeBindings(tbind, field.TypeExpr)
		jb := buildDecodeBinding(dres, monoTe)
		fieldJB = append(fieldJB, jb)
	}
	return func(ds *DecodeState, v any) error {
		switch t := v.(type) {
		case map[string]any:
			for i, f := range struct_.Fields {
				if v0, ok := t[f.SerializedName]; ok {
					ds0 := DecodeState{
						V:    ds.V.Field(i),
						Path: append(ds.Path, f.Name),
					}
					err := fieldJB[i](&ds0, v0)
					if err != nil {
						return err
					}
					continue
				}
				if _, ok := any(f.Default.Branch).(types.Maybe_Nothing); ok {
					return fmt.Errorf("path %v, required field missing '%v'", ds.Path, f.SerializedName)
				}
				// set from default field value
				rv := reflect.ValueOf(f.Default.Branch).Field(0)
				ds0 := DecodeState{
					V:    ds.V.Field(i),
					Path: append(ds.Path, f.Name),
				}
				err := fieldJB[i](&ds0, rv.Interface())
				if err != nil {
					return err
				}
			}
			return nil
		}
		return fmt.Errorf("path %v, struct: expect an object received %v '%v'", ds.Path, reflect.TypeOf(v), v)
	}
}

func enumDecodeBinding(
	dres Resolver,
	union_ adlast.Union,
) DecodeFunc {
	decMap := make(map[string]DecodeFunc)
	for _, f := range union_.Fields {
		bf := buildDecodeBinding(dres, Texpr_Void().Value)
		if bf == nil {
			panic(fmt.Errorf("DecodeFunc == nil - %#+v", f.TypeExpr))
		}
		decMap[f.SerializedName] = bf
	}
	return func(ds *DecodeState, v any) error {
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
				return fmt.Errorf("path: %v, expect an object with one and only element received %v", ds.Path, len(t))
			}
			for k0, v0 := range t {
				key = k0
				val = v0
			}
		default:
			return fmt.Errorf("path: %v, union: expect an object received %v '%v'", ds.Path, reflect.TypeOf(v), v)
		}

		if bf, ok := decMap[key]; ok {
			var vn reflect.Value
			if ds.V.CanAddr() && ds.V.Addr().Type().Implements(reflect.TypeFor[BranchFactory]()) {
				meth := ds.V.Addr().MethodByName("MakeNewBranch")
				resps := meth.Call([]reflect.Value{reflect.ValueOf(key)})
				if resps[1].Interface() != nil {
					return fmt.Errorf("path: %v, unexpected branch - no type in branch factory '%v'", ds.Path, key)
				}
				vn = resps[0].Elem()
			} else {
				return fmt.Errorf("path: %v, MakeNewBranch not implemented '%v'", ds.Path, ds.V.Type())
			}
			ds0 := DecodeState{
				V:    vn.Elem().Field(0),
				Path: append(ds.Path, key),
			}
			err := bf(&ds0, val)
			if err != nil {
				return err
			}
			r0 := ds.V // for top level Elem() is already called
			r0 = r0.Field(0)
			// r0 = r0.Field(0)
			r0.Set(vn.Elem())
			return nil
		} else {
			return fmt.Errorf("path: %v, unexpected branch '%v'", ds.Path, key)
		}
	}
}

func unionDecodeBinding(
	dres Resolver,
	union_ adlast.Union,
	tbind []TypeBinding,
) DecodeFunc {
	decMap := make(map[string]DecodeFunc)
	for _, f := range union_.Fields {
		monoTe := SubstituteTypeBindings(tbind, f.TypeExpr)
		bf := buildDecodeBinding(dres, monoTe)
		if bf == nil {
			panic(fmt.Errorf("DecodeFunc == nil - %#+v", f.TypeExpr))
		}
		decMap[f.SerializedName] = bf
	}
	return func(ds *DecodeState, v any) error {
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
				return fmt.Errorf("path: %v, expect an object with one and only element received %v", ds.Path, len(t))
			}
			for k0, v0 := range t {
				key = k0
				val = v0
			}
		default:
			return fmt.Errorf("path: %v, union: expect an object received %v '%v'", ds.Path, reflect.TypeOf(v), v)
		}

		if bf, ok := decMap[key]; ok {
			var vn reflect.Value
			if ds.V.CanAddr() && ds.V.Addr().Type().Implements(reflect.TypeFor[BranchFactory]()) {
				meth := ds.V.Addr().MethodByName("MakeNewBranch")
				resps := meth.Call([]reflect.Value{reflect.ValueOf(key)})
				if resps[1].Interface() != nil {
					return fmt.Errorf("path: %v, unexpected branch - no type in branch factory '%v'", ds.Path, key)
				}
				vn = resps[0].Elem()
			} else {
				return fmt.Errorf("path: %v, MakeNewBranch not implemented '%v'", ds.Path, ds.V.Type())
			}
			ds0 := DecodeState{
				V:    vn.Elem().Field(0),
				Path: append(ds.Path, key),
			}
			err := bf(&ds0, val)
			if err != nil {
				return err
			}
			r0 := ds.V // for top level Elem() is already called
			r0 = r0.Field(0)
			// r0 = r0.Field(0)
			r0.Set(vn.Elem())
			return nil
		} else {
			return fmt.Errorf("path: %v, unexpected branch '%v'", ds.Path, key)
		}

	}
}
