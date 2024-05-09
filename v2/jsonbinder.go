package goadl

import (
	"bytes"
	"reflect"
)

type Json any

// type JsonBinding[T any] interface {
// 	// Convert a value of type T to json
// 	ToJson(t T) Json
// 	// Parse json into a value of type T.
// 	//
// 	// Throws JsonParseExceptions on failure.
// 	fromJson(json Json) T
// }

type JsonBinding[T any] struct {
	// Convert a value of type T to json
	ToJson func(t T) Json
	// Parse json into a value of type T.
	FromJson func(json Json) (T, error)
}

type BoundTypeParams map[string]*JsonBinding[any]

type encodeState struct {
	bytes.Buffer // accumulated output
}

type encoderFunc func(e *encodeState, v reflect.Value)

func CreateJsonEncodeBinding[T any](
	dres Resolver,
	texpr ATypeExpr[T],
) *JsonBinding[T] {
	return buildJsonBinding[T](dres, texpr.Value, map[string]*JsonBinding[any]{})
}

func buildJsonBinding[T any](
	dres Resolver,
	texpr TypeExpr,
	boundTypeParams BoundTypeParams,
) *JsonBinding[T] {
	jb0 := buildJsonBinding0(dres, texpr, boundTypeParams)
	jb := &JsonBinding[T]{
		ToJson: func(t T) Json { return jb0.ToJson(t) },
		FromJson: func(json Json) (T, error) {
			r, e := jb0.FromJson(json)
			return r.(T), e
		},
	}
	return jb
}

func buildJsonBinding0(
	dres Resolver,
	texpr TypeExpr,
	boundTypeParams BoundTypeParams,
) *JsonBinding[any] {
	return Ok(HandleTypeRef[*JsonBinding[any]](
		texpr.TypeRef.Branch,
		func(trb TypeRefBranch_Primitive) (*JsonBinding[any], error) {
			return primitiveJsonBinding(dres, string(trb), texpr.Parameters, boundTypeParams), nil
		},
		func(trb TypeRefBranch_TypeParam) (*JsonBinding[any], error) {
			return boundTypeParams[string(trb)], nil
		},
		func(trb TypeRefBranch_Reference) (*JsonBinding[any], error) {
			ast := dres.Resolve(ScopedName(trb))
			return HandleDeclType[*JsonBinding[any]](
				ast.Decl.Type.Branch,
				func(dtb DeclTypeBranch_Struct_) (*JsonBinding[any], error) {
					return structJsonBinding(dres, Struct(dtb), texpr.Parameters, boundTypeParams), nil
				},
				func(dtb DeclTypeBranch_Union_) (*JsonBinding[any], error) {
					// union := Union(dtb)
					return nil, nil
				},
				func(dtb DeclTypeBranch_Type_) (*JsonBinding[any], error) {
					return typedefJsonBinding(dres, TypeDef(dtb), texpr.Parameters, boundTypeParams), nil
				},
				func(dtb DeclTypeBranch_Newtype_) (*JsonBinding[any], error) {
					return newtypeJsonBinding(dres, NewType(dtb), texpr.Parameters, boundTypeParams), nil
				},
			)
		},
	))
}

func structJsonBinding(
	dres Resolver,
	struct_ Struct,
	params []TypeExpr,
	boundTypeParams BoundTypeParams,
) *JsonBinding[any] {
	newBoundTypeParams := createBoundTypeParams(dres, struct_.TypeParams, params, boundTypeParams)
	fieldJB := make([]*JsonBinding[any], 0)
	for _, field := range struct_.Fields {
		jb := buildJsonBinding0(dres, field.TypeExpr, newBoundTypeParams)
		fieldJB = append(fieldJB, jb)
	}

	ret := JsonBinding[any]{
		ToJson: func(t any) Json {
			v := make(map[string]any)

			return v
		},
		FromJson: func(json Json) (any, error) {
			return nil, nil
		},
	}
	return &ret
}

func enumJsonBinding(
	dres Resolver,
	union_ Union,
	params []TypeExpr,
	boundTypeParams BoundTypeParams,
) *JsonBinding[any] {
	return nil
}

func unionJsonBinding(
	dres Resolver,
	union_ Union,
	params []TypeExpr,
	boundTypeParams BoundTypeParams,
) *JsonBinding[any] {
	return nil
}

func newtypeJsonBinding(
	dres Resolver,
	newtype NewType,
	params []TypeExpr,
	boundTypeParams BoundTypeParams,
) *JsonBinding[any] {
	return nil
}

func typedefJsonBinding(
	dres Resolver,
	typedef TypeDef,
	params []TypeExpr,
	boundTypeParams BoundTypeParams,
) *JsonBinding[any] {
	return nil
}

func primitiveJsonBinding(
	dres Resolver,
	ptype string,
	params []TypeExpr,
	boundTypeParams BoundTypeParams,
) *JsonBinding[any] {
	return nil
}

func createBoundTypeParams(
	dresolver Resolver,
	paramNames []string,
	paramTypes []TypeExpr,
	boundTypeParams BoundTypeParams,
) BoundTypeParams {
	result := make(BoundTypeParams)
	for i, paramName := range paramNames {
		result[paramName] = buildJsonBinding0(dresolver, paramTypes[i], boundTypeParams)
	}
	return result
}

func Ok[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
