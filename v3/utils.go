package goadl

import (
	"fmt"

	adlast "github.com/adl-lang/goadl_rt/v3/sys/adlast"
)

func SubstituteTypeBindings(m map[string]adlast.TypeExpr, te adlast.TypeExpr) adlast.TypeExpr {
	p0 := make([]adlast.TypeExpr, len(te.Parameters))
	for i := range te.Parameters {
		p0[i] = SubstituteTypeBindings(m, te.Parameters[i])
	}

	if tp, ok := te.TypeRef.Branch.(adlast.TypeRef_TypeParam); ok {
		if te0, ok := m[tp.V]; !ok {
			panic(fmt.Errorf("type param not found %v", tp.V))
			// return adlast.TypeExpr{
			// 	TypeRef:    te.TypeRef,
			// 	Parameters: p0,
			// }
		} else {
			if len(te.Parameters) != 0 {
				panic(fmt.Errorf("type param cannot have type params, not a concrete type"))
			}
			return te0
		}
	}

	return adlast.TypeExpr{
		TypeRef:    te.TypeRef,
		Parameters: p0,
	}
}

func TypeParamsFromDecl(decl adlast.Decl) []string {
	return adlast.Handle_DeclType[[]string](
		decl.Type_.Branch,
		func(struct_ adlast.Struct) []string {
			return struct_.TypeParams
		},
		func(union_ adlast.Union) []string {
			return union_.TypeParams
		},
		func(type_ adlast.TypeDef) []string {
			return type_.TypeParams
		},
		func(newtype_ adlast.NewType) []string {
			return newtype_.TypeParams
		},
		nil,
	)
}
