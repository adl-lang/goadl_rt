// Code generated by goadlc v3 - DO NOT EDIT.
package adlast

func Texpr_Vector[T any](te ATypeExpr[T]) ATypeExpr[[]T] {
	te0 := Make_TypeExpr(
		Make_TypeRef_primitive("Vector"),
		[]TypeExpr{
			te.Value,
		},
	)
	return Make_ATypeExpr[[]T](te0)
}

func Texpr_StringMap[T any](te ATypeExpr[T]) ATypeExpr[map[string]T] {
	te0 := Make_TypeExpr(
		Make_TypeRef_primitive("StringMap"),
		[]TypeExpr{
			te.Value,
		},
	)
	return Make_ATypeExpr[map[string]T](te0)
}

func Texpr_Nullable[T any](te ATypeExpr[T]) ATypeExpr[*T] {
	te0 := Make_TypeExpr(
		Make_TypeRef_primitive("Nullable"),
		[]TypeExpr{
			te.Value,
		},
	)
	return Make_ATypeExpr[*T](te0)
}

func Texpr_Int8() ATypeExpr[int8] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Int8"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[int8](te)
}

func Texpr_Int16() ATypeExpr[int16] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Int16"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[int16](te)
}

func Texpr_Int32() ATypeExpr[int32] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Int32"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[int32](te)
}

func Texpr_Int64() ATypeExpr[int64] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Int64"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[int64](te)
}

func Texpr_Word8() ATypeExpr[uint8] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Word8"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[uint8](te)
}

func Texpr_Word16() ATypeExpr[uint16] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Word16"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[uint16](te)
}

func Texpr_Word32() ATypeExpr[uint32] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Word32"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[uint32](te)
}

func Texpr_Word64() ATypeExpr[uint64] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Word64"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[uint64](te)
}

func Texpr_Bool() ATypeExpr[bool] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Bool"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[bool](te)
}

func Texpr_Float() ATypeExpr[float64] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Float"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[float64](te)
}

func Texpr_Double() ATypeExpr[float64] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Double"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[float64](te)
}

func Texpr_String() ATypeExpr[string] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("String"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[string](te)
}

func Texpr_Void() ATypeExpr[struct{}] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Void"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[struct{}](te)
}

func Texpr_Json() ATypeExpr[any] {
	te := Make_TypeExpr(
		Make_TypeRef_primitive("Json"),
		[]TypeExpr{},
	)
	return Make_ATypeExpr[any](te)
}
