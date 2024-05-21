package goadl

import (
	. "github.com/adl-lang/goadl_rt/v3/sys/adlast"
)

type ATypeExpr[T any] struct {
	Value TypeExpr
}

func Texpr_Vector[T any](te ATypeExpr[T]) ATypeExpr[[]T] {
	return ATypeExpr[[]T]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Vector"},
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
				Branch: TypeRef_Primitive{V: "StringMap"},
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
				Branch: TypeRef_Primitive{V: "Nullable"},
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
				Branch: TypeRef_Primitive{V: "Int8"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Int16() ATypeExpr[int16] {
	return ATypeExpr[int16]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Int16"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Int32() ATypeExpr[int32] {
	return ATypeExpr[int32]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Int32"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Int64() ATypeExpr[int64] {
	return ATypeExpr[int64]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Int64"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Word8() ATypeExpr[uint8] {
	return ATypeExpr[uint8]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Word8"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Word16() ATypeExpr[uint16] {
	return ATypeExpr[uint16]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Word16"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Word32() ATypeExpr[uint32] {
	return ATypeExpr[uint32]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Word32"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Word64() ATypeExpr[uint64] {
	return ATypeExpr[uint64]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Word64"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Bool() ATypeExpr[bool] {
	return ATypeExpr[bool]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Bool"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Float() ATypeExpr[float64] {
	return ATypeExpr[float64]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Float"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Double() ATypeExpr[float64] {
	return ATypeExpr[float64]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Double"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_String() ATypeExpr[string] {
	return ATypeExpr[string]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "String"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Void() ATypeExpr[struct{}] {
	return ATypeExpr[struct{}]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Void"},
			},
			Parameters: []TypeExpr{},
		},
	}
}

func Texpr_Json() ATypeExpr[any] {
	return ATypeExpr[any]{
		Value: TypeExpr{
			TypeRef: TypeRef{
				Branch: TypeRef_Primitive{V: "Json"},
			},
			Parameters: []TypeExpr{},
		},
	}
}
