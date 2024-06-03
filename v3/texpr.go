// Code generated by goadlc v3 - DO NOT EDIT.
package goadl

import (
	. "github.com/adl-lang/goadl_rt/v3/sys/adlast"
)

// func Texpr_Vector[T any](te ATypeExpr[T]) ATypeExpr[[]T] {
// 	te0 := Make_TypeExpr(
// 		Make_TypeRef_primitive("Vector"),
// 		[]TypeExpr{
// 			te.Value,
// 		},
// 	)
// 	return ATypeExpr[[]T]{
// 		Value: te0,
// 	}
// }

func Texpr_StringMap[T any](te ATypeExpr[T]) ATypeExpr[map[string]T] {
	te0 := Make_TypeExpr(
		Make_TypeRef_primitive("StringMap"),
		[]TypeExpr{
			te.Value,
		},
	)
	return Make_ATypeExpr[map[string]T](te0)
}

// func Texpr_Nullable[T any](te ATypeExpr[T]) ATypeExpr[*T] {
// 	te0 := Make_TypeExpr(
// 		Make_TypeRef_primitive("Nullable"),
// 		[]TypeExpr{
// 			te.Value,
// 		},
// 	)
// 	return ATypeExpr[*T]{
// 		Value: te0,
// 	}
// }

// func Texpr_Int8() ATypeExpr[int8] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Int8"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[int8]{
// 		Value: te,
// 	}
// }

// func Texpr_Int16() ATypeExpr[int16] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Int16"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[int16]{
// 		Value: te,
// 	}
// }

// func Texpr_Int32() ATypeExpr[int32] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Int32"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[int32]{
// 		Value: te,
// 	}
// }

// func Texpr_Int64() ATypeExpr[int64] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Int64"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[int64]{
// 		Value: te,
// 	}
// }

// func Texpr_Word8() ATypeExpr[uint8] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Word8"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[uint8]{
// 		Value: te,
// 	}
// }

// func Texpr_Word16() ATypeExpr[uint16] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Word16"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[uint16]{
// 		Value: te,
// 	}
// }

// func Texpr_Word32() ATypeExpr[uint32] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Word32"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[uint32]{
// 		Value: te,
// 	}
// }

// func Texpr_Word64() ATypeExpr[uint64] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Word64"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[uint64]{
// 		Value: te,
// 	}
// }

// func Texpr_Bool() ATypeExpr[bool] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Bool"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[bool]{
// 		Value: te,
// 	}
// }

// func Texpr_Float() ATypeExpr[float64] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Float"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[float64]{
// 		Value: te,
// 	}
// }

// func Texpr_Double() ATypeExpr[float64] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Double"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[float64]{
// 		Value: te,
// 	}
// }

// func Texpr_String() ATypeExpr[string] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("String"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[string]{
// 		Value: te,
// 	}
// }

func Texpr_Void() ATypeExpr[struct{}] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Void"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[struct{}](te)
}

// func Texpr_Json() ATypeExpr[any] {
// 	te := Make_TypeExpr(
// 		Make_TypeRef_primitive("Json"),
// 		[]TypeExpr{},
// 	)
// 	return ATypeExpr[any]{
// 		Value: te,
// 	}
// }
