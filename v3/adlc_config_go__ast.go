// Code generated by goadlc v3 - DO NOT EDIT.
package goadl

import (
	. "github.com/adl-lang/goadl_rt/v3/adlc/config/go_"
	adlast "github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_GoCustomType() ATypeExpr[GoCustomType] {
	return ATypeExpr[GoCustomType]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "adlc.config.go_",
						Name:       "GoCustomType",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_GoCustomType() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "GoCustomType",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Struct_{
				V: adlast.Struct{
					TypeParams: []adlast.Ident{},
					Fields: []adlast.Field{
						adlast.Field{
							Name:           "gotype",
							SerializedName: "gotype",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Reference{
										V: adlast.ScopedName{
											ModuleName: "adlc.config.go_",
											Name:       "GoType",
										}},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{
								types.MapEntry[adlast.ScopedName, any]{
									Key: adlast.ScopedName{
										ModuleName: "sys.annotations",
										Name:       "Doc",
									},
									Value: "The Go struct to use\n\nNote currently interfaces are not supported\n",
								},
							}),
						},
						adlast.Field{
							Name:           "helpers",
							SerializedName: "helpers",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Reference{
										V: adlast.ScopedName{
											ModuleName: "adlc.config.go_",
											Name:       "GoHelperType",
										}},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{
								types.MapEntry[adlast.ScopedName, any]{
									Key: adlast.ScopedName{
										ModuleName: "sys.annotations",
										Name:       "Doc",
									},
									Value: "The Go structed which implements json Marshal and Unmarshal for the gotype.\n\nNote Marshal and Unmarshal must be implemented for the specified struct's nil ptr\n",
								},
							}),
						},
					},
				}},
		},
		Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{
			types.MapEntry[adlast.ScopedName, any]{
				Key: adlast.ScopedName{
					ModuleName: "sys.annotations",
					Name:       "Doc",
				},
				Value: "ADL declaration annotation to specify the custom type to use\n",
			},
		}),
	}
	return adlast.ScopedDecl{
		ModuleName: "adlc.config.go_",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "adlc.config.go_", Name: "GoCustomType"},
		AST_GoCustomType(),
	)
}

func Texpr_GoHelperType() ATypeExpr[GoHelperType] {
	return ATypeExpr[GoHelperType]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "adlc.config.go_",
						Name:       "GoHelperType",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_GoHelperType() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "GoHelperType",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Struct_{
				V: adlast.Struct{
					TypeParams: []adlast.Ident{},
					Fields: []adlast.Field{
						adlast.Field{
							Name:           "name",
							SerializedName: "name",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "String"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{}),
						},
						adlast.Field{
							Name:           "pkg",
							SerializedName: "pkg",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "String"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{}),
						},
						adlast.Field{
							Name:           "import_path",
							SerializedName: "import_path",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "String"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "adlc.config.go_",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "adlc.config.go_", Name: "GoHelperType"},
		AST_GoHelperType(),
	)
}

func Texpr_GoType() ATypeExpr[GoType] {
	return ATypeExpr[GoType]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "adlc.config.go_",
						Name:       "GoType",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_GoType() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "GoType",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Struct_{
				V: adlast.Struct{
					TypeParams: []adlast.Ident{},
					Fields: []adlast.Field{
						adlast.Field{
							Name:           "name",
							SerializedName: "name",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "String"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{}),
						},
						adlast.Field{
							Name:           "pkg",
							SerializedName: "pkg",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "String"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{}),
						},
						adlast.Field{
							Name:           "import_path",
							SerializedName: "import_path",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "String"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[adlast.ScopedName, any]([]types.MapEntry[adlast.ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "adlc.config.go_",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "adlc.config.go_", Name: "GoType"},
		AST_GoType(),
	)
}