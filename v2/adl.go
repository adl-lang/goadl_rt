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
