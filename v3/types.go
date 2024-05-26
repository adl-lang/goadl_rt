package goadl

import (
	"fmt"
	"reflect"
)

type MapSet[A comparable] map[A]struct{}

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
		fmt.Printf("SetHelper DecodeFunc %+#v\n", v)
		rv := reflect.ValueOf(v)
		l := rv.Len()
		for i := 0; i < l; i++ {
			v0 := rv.Index(i).Interface()
			typeparamDec[0](ds, v0)
		}
		return nil
	}
	return f
}
