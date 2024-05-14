package goadl

type ATypeExpr[T any] struct {
	Value TypeExpr
}

func Texpr_Vector[T any](te ATypeExpr[T]) ATypeExpr[[]T] {
	return ATypeExpr[[]T]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Vector"),
			},
			Parameters: []TypeExpr{
				te.Value,
			},
		},
	}
}

func Texpr_StringMap[T any](te ATypeExpr[T]) ATypeExpr[map[string]T] {
	return ATypeExpr[map[string]T]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("StringMap"),
			},
			Parameters: []TypeExpr{
				te.Value,
			},
		},
	}
}

func Texpr_Nullable[T any](te ATypeExpr[T]) ATypeExpr[*T] {
	return ATypeExpr[*T]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Nullable"),
			},
			Parameters: []TypeExpr{
				te.Value,
			},
		},
	}
}

func Texpr_Int8() ATypeExpr[int8] {
	return ATypeExpr[int8]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Int8"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Int16() ATypeExpr[int16] {
	return ATypeExpr[int16]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Int16"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Int32() ATypeExpr[int32] {
	return ATypeExpr[int32]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Int32"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Int64() ATypeExpr[int64] {
	return ATypeExpr[int64]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Int64"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Word8() ATypeExpr[uint8] {
	return ATypeExpr[uint8]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Word8"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Word16() ATypeExpr[uint16] {
	return ATypeExpr[uint16]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Word16"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Word32() ATypeExpr[uint32] {
	return ATypeExpr[uint32]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Word32"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Word64() ATypeExpr[uint64] {
	return ATypeExpr[uint64]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Word64"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Bool() ATypeExpr[bool] {
	return ATypeExpr[bool]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Bool"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Float() ATypeExpr[float64] {
	return ATypeExpr[float64]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Float"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Double() ATypeExpr[float64] {
	return ATypeExpr[float64]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Double"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_String() ATypeExpr[string] {
	return ATypeExpr[string]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("String"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Void() ATypeExpr[struct{}] {
	return ATypeExpr[struct{}]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Void"),
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Json() ATypeExpr[any] {
	return ATypeExpr[any]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRefBranch_Primitive("Json"),
			},
			Parameters: []TypeExpr{},
		},
	}
}
