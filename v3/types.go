package goadl

import (
	"reflect"
	"strconv"

	adlast "github.com/adl-lang/goadl_rt/v3/sys/adlast"
)

type MapSet[A comparable] map[A]struct{}

func Texpr_Set[T comparable](t ATypeExpr[T]) ATypeExpr[MapSet[T]] {
	return ATypeExpr[MapSet[T]]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.types",
						Name:       "Set",
					},
				},
			},
			Parameters: []adlast.TypeExpr{t.Value},
		},
	}
}

func SetEncoderFunc(e *EncodeState, v reflect.Value) error {
	return nil
}

func SetDecodeFunc(ds *DecodeState, v any) error {
	return nil
}

type SetHelper struct{}

func (*SetHelper) BuildEncodeFunc(typeparamEnc ...EncoderFunc) EncoderFunc {
	var f EncoderFunc = func(e *EncodeState, v reflect.Value) error {
		e.WriteByte('[')
		iter := v.MapRange()
		rest := false
		for iter.Next() {
			if rest {
				e.WriteByte(',')
			} else {
				rest = true
			}
			k := iter.Key()
			err := typeparamEnc[0](e, k)
			if err != nil {
				return err
			}
		}
		e.WriteByte(']')
		return nil
	}
	return f
}

func (*SetHelper) BuildDecodeFunc(typeparamDec ...DecodeFunc) DecodeFunc {
	var f DecodeFunc = func(ds *DecodeState, v any) error {
		newM := reflect.MakeMap(ds.V.Type())
		rv := reflect.ValueOf(v)
		l := rv.Len()
		for i := 0; i < l; i++ {
			v0 := rv.Index(i).Interface()
			ds0 := DecodeState{
				V:    reflect.New(ds.V.Type().Key()).Elem(),
				Path: append(ds.Path, "["+strconv.Itoa(i)+"]"),
			}
			err := typeparamDec[0](&ds0, v0)
			if err != nil {
				return err
			}
			newM.SetMapIndex(ds0.V, reflect.ValueOf(struct{}{}))
		}
		ds.V.Set(newM)
		return nil
	}
	return f
}
