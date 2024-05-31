package goadl

import (
	"fmt"

	adlast "github.com/adl-lang/goadl_rt/v3/sys/adlast"
)

type TypeBinding struct {
	Name  string
	Value adlast.TypeExpr
}

func CreateDecBoundTypeParams(
	paramNames []string,
	paramTypes []adlast.TypeExpr,
) []TypeBinding {
	binding := make([]TypeBinding, len(paramNames))
	for i, paramName := range paramNames {
		binding[i] = TypeBinding{
			Name:  paramName,
			Value: paramTypes[i],
		}
	}
	return binding
}

func SubstituteTypeBindings(binding []TypeBinding, te adlast.TypeExpr) adlast.TypeExpr {
	usedTp := make([]*adlast.TypeExpr, len(binding))
	var recurse func(binding []TypeBinding, te adlast.TypeExpr) adlast.TypeExpr
	recurse = func(binding []TypeBinding, te adlast.TypeExpr) adlast.TypeExpr {
		parameters := make([]adlast.TypeExpr, len(te.Parameters))
		for i := range te.Parameters {
			parameters[i] = recurse(binding, te.Parameters[i])
		}

		if tp, ok := te.TypeRef.Cast_typeParam(); ok {
			// if tp, ok := te.TypeRef.Branch.(adlast.TypeRef_TypeParam); ok {
			// tp := tp.V
			for i, b := range binding {
				if b.Name == tp {
					if len(te.Parameters) != 0 {
						panic(fmt.Errorf("type param cannot have type params, not a concrete type"))
					}
					usedTp[i] = &b.Value
					return b.Value
				}
			}
			panic(fmt.Errorf("type param not found %v", tp))
		}

		return adlast.TypeExpr{
			TypeRef:    te.TypeRef,
			Parameters: parameters,
		}
	}
	return recurse(binding, te)
}

func TypeParamsFromDecl(decl adlast.Decl) []string {
	return adlast.Handle_DeclType[[]string](
		decl.Type_,
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

// func HasAnnotation(anns adlast.Annotations, sn adlast.ScopedName) bool {
// 	for i := range anns {
// 		ann := anns[i]
// 		if ScopedNamesEqual(ann.Key, sn) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func GetAnnotation[T any](anns adlast.Annotations, sn adlast.ScopedName, jb JsonDecodeBinder[T]) (*T, error) {
// 	for i := range anns {
// 		ann := anns[i]
// 		if ScopedNamesEqual(ann.Key, sn) {
// 			var dst T
// 			err := jb.DecodeFromAny(ann.Value, &dst)
// 			if err != nil {
// 				return nil, err
// 			}
// 			return &dst, nil
// 		}
// 	}
// 	return nil, nil
// }

func GetAnnotation[T any](anns adlast.Annotations, sn adlast.ScopedName, jb JsonDecodeBinder[T]) (*T, error) {
	if val, ok := anns[sn]; ok {
		var dst T
		err := jb.DecodeFromAny(val, &dst)
		if err != nil {
			return nil, err
		}
		return &dst, nil
	}
	return nil, nil
}

func HasAnnotation(anns adlast.Annotations, sn adlast.ScopedName) bool {
	_, ok := anns[sn]
	return ok
}

func ScopedNamesEqual(sn1 adlast.ScopedName, sn2 adlast.ScopedName) bool {
	return sn1.ModuleName == sn2.ModuleName && sn1.Name == sn2.Name
}
