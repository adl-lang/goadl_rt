// Code generated by goadlc v3 - DO NOT EDIT.
package goadl

import (
	. "github.com/adl-lang/goadl_rt/v3/sys/adlast"
	adlast "github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_Annotations() ATypeExpr[Annotations] {
	return ATypeExpr[Annotations]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "Annotations",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_Annotations() adlast.ScopedDecl {
	decl := Decl{
		Name: "Annotations",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Type_{
				V: TypeDef{
					TypeParams: []Ident{},
					TypeExpr: TypeExpr{
						TypeRef: TypeRef{
							Branch: TypeRef_Reference{
								V: ScopedName{
									ModuleName: "sys.types",
									Name:       "Map",
								}},
						},
						Parameters: []TypeExpr{
							TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "ScopedName",
										}},
								},
								Parameters: []TypeExpr{},
							},
							TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Primitive{
										V: "Json"},
								},
								Parameters: []TypeExpr{},
							},
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "Annotations"},
		AST_Annotations(),
	)
}

func Texpr_Decl() ATypeExpr[Decl] {
	return ATypeExpr[Decl]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "Decl",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_Decl() adlast.ScopedDecl {
	decl := Decl{
		Name: "Decl",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Struct_{
				V: Struct{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "name",
							SerializedName: "name",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Ident",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "version",
							SerializedName: "version",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.types",
											Name:       "Maybe",
										}},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Primitive{
												V: "Word32"},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "type_",
							SerializedName: "type_",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "DeclType",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "annotations",
							SerializedName: "annotations",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Annotations",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "Decl"},
		AST_Decl(),
	)
}

func Texpr_DeclType() ATypeExpr[DeclType] {
	return ATypeExpr[DeclType]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "DeclType",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_DeclType() adlast.ScopedDecl {
	decl := Decl{
		Name: "DeclType",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Union_{
				V: Union{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "struct_",
							SerializedName: "struct_",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Struct",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "union_",
							SerializedName: "union_",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Union",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "type_",
							SerializedName: "type_",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "TypeDef",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "newtype_",
							SerializedName: "newtype_",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "NewType",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "DeclType"},
		AST_DeclType(),
	)
}

func Texpr_DeclVersions() ATypeExpr[DeclVersions] {
	return ATypeExpr[DeclVersions]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "DeclVersions",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_DeclVersions() adlast.ScopedDecl {
	decl := Decl{
		Name: "DeclVersions",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Type_{
				V: TypeDef{
					TypeParams: []Ident{},
					TypeExpr: TypeExpr{
						TypeRef: TypeRef{
							Branch: TypeRef_Primitive{
								V: "Vector"},
						},
						Parameters: []TypeExpr{
							TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Decl",
										}},
								},
								Parameters: []TypeExpr{},
							},
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "DeclVersions"},
		AST_DeclVersions(),
	)
}

func Texpr_Field() ATypeExpr[Field] {
	return ATypeExpr[Field]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "Field",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_Field() adlast.ScopedDecl {
	decl := Decl{
		Name: "Field",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Struct_{
				V: Struct{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "name",
							SerializedName: "name",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Ident",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "serializedName",
							SerializedName: "serializedName",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Ident",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "typeExpr",
							SerializedName: "typeExpr",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "TypeExpr",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "default",
							SerializedName: "default",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.types",
											Name:       "Maybe",
										}},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Primitive{
												V: "Json"},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "annotations",
							SerializedName: "annotations",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Annotations",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "Field"},
		AST_Field(),
	)
}

func Texpr_Ident() ATypeExpr[Ident] {
	return ATypeExpr[Ident]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "Ident",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_Ident() adlast.ScopedDecl {
	decl := Decl{
		Name: "Ident",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Type_{
				V: TypeDef{
					TypeParams: []Ident{},
					TypeExpr: TypeExpr{
						TypeRef: TypeRef{
							Branch: TypeRef_Primitive{
								V: "String"},
						},
						Parameters: []TypeExpr{},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "Ident"},
		AST_Ident(),
	)
}

func Texpr_Import() ATypeExpr[Import] {
	return ATypeExpr[Import]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "Import",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_Import() adlast.ScopedDecl {
	decl := Decl{
		Name: "Import",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Union_{
				V: Union{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "moduleName",
							SerializedName: "moduleName",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "ModuleName",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "scopedName",
							SerializedName: "scopedName",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "ScopedName",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "Import"},
		AST_Import(),
	)
}

func Texpr_Module() ATypeExpr[Module] {
	return ATypeExpr[Module]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "Module",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_Module() adlast.ScopedDecl {
	decl := Decl{
		Name: "Module",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Struct_{
				V: Struct{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "name",
							SerializedName: "name",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "ModuleName",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "imports",
							SerializedName: "imports",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Primitive{
										V: "Vector"},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Reference{
												V: ScopedName{
													ModuleName: "sys.adlast",
													Name:       "Import",
												}},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "decls",
							SerializedName: "decls",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Primitive{
										V: "StringMap"},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Reference{
												V: ScopedName{
													ModuleName: "sys.adlast",
													Name:       "Decl",
												}},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "annotations",
							SerializedName: "annotations",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Annotations",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "Module"},
		AST_Module(),
	)
}

func Texpr_ModuleName() ATypeExpr[ModuleName] {
	return ATypeExpr[ModuleName]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "ModuleName",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_ModuleName() adlast.ScopedDecl {
	decl := Decl{
		Name: "ModuleName",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Type_{
				V: TypeDef{
					TypeParams: []Ident{},
					TypeExpr: TypeExpr{
						TypeRef: TypeRef{
							Branch: TypeRef_Primitive{
								V: "String"},
						},
						Parameters: []TypeExpr{},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "ModuleName"},
		AST_ModuleName(),
	)
}

func Texpr_NewType() ATypeExpr[NewType] {
	return ATypeExpr[NewType]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "NewType",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_NewType() adlast.ScopedDecl {
	decl := Decl{
		Name: "NewType",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Struct_{
				V: Struct{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "typeParams",
							SerializedName: "typeParams",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Primitive{
										V: "Vector"},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Reference{
												V: ScopedName{
													ModuleName: "sys.adlast",
													Name:       "Ident",
												}},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "typeExpr",
							SerializedName: "typeExpr",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "TypeExpr",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "default",
							SerializedName: "default",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.types",
											Name:       "Maybe",
										}},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Primitive{
												V: "Json"},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "NewType"},
		AST_NewType(),
	)
}

func Texpr_ScopedDecl() ATypeExpr[ScopedDecl] {
	return ATypeExpr[ScopedDecl]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "ScopedDecl",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_ScopedDecl() adlast.ScopedDecl {
	decl := Decl{
		Name: "ScopedDecl",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Struct_{
				V: Struct{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "moduleName",
							SerializedName: "moduleName",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "ModuleName",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "decl",
							SerializedName: "decl",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Decl",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "ScopedDecl"},
		AST_ScopedDecl(),
	)
}

func Texpr_ScopedName() ATypeExpr[ScopedName] {
	return ATypeExpr[ScopedName]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "ScopedName",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_ScopedName() adlast.ScopedDecl {
	decl := Decl{
		Name: "ScopedName",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Struct_{
				V: Struct{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "moduleName",
							SerializedName: "moduleName",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "ModuleName",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "name",
							SerializedName: "name",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Ident",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "ScopedName"},
		AST_ScopedName(),
	)
}

func Texpr_Struct() ATypeExpr[Struct] {
	return ATypeExpr[Struct]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "Struct",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_Struct() adlast.ScopedDecl {
	decl := Decl{
		Name: "Struct",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Struct_{
				V: Struct{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "typeParams",
							SerializedName: "typeParams",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Primitive{
										V: "Vector"},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Reference{
												V: ScopedName{
													ModuleName: "sys.adlast",
													Name:       "Ident",
												}},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "fields",
							SerializedName: "fields",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Primitive{
										V: "Vector"},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Reference{
												V: ScopedName{
													ModuleName: "sys.adlast",
													Name:       "Field",
												}},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "Struct"},
		AST_Struct(),
	)
}

func Texpr_TypeDef() ATypeExpr[TypeDef] {
	return ATypeExpr[TypeDef]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "TypeDef",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_TypeDef() adlast.ScopedDecl {
	decl := Decl{
		Name: "TypeDef",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Struct_{
				V: Struct{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "typeParams",
							SerializedName: "typeParams",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Primitive{
										V: "Vector"},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Reference{
												V: ScopedName{
													ModuleName: "sys.adlast",
													Name:       "Ident",
												}},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "typeExpr",
							SerializedName: "typeExpr",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "TypeExpr",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "TypeDef"},
		AST_TypeDef(),
	)
}

func Texpr_TypeExpr() ATypeExpr[TypeExpr] {
	return ATypeExpr[TypeExpr]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "TypeExpr",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_TypeExpr() adlast.ScopedDecl {
	decl := Decl{
		Name: "TypeExpr",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Struct_{
				V: Struct{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "typeRef",
							SerializedName: "typeRef",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "TypeRef",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "parameters",
							SerializedName: "parameters",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Primitive{
										V: "Vector"},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Reference{
												V: ScopedName{
													ModuleName: "sys.adlast",
													Name:       "TypeExpr",
												}},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "TypeExpr"},
		AST_TypeExpr(),
	)
}

func Texpr_TypeRef() ATypeExpr[TypeRef] {
	return ATypeExpr[TypeRef]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "TypeRef",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_TypeRef() adlast.ScopedDecl {
	decl := Decl{
		Name: "TypeRef",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Union_{
				V: Union{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "primitive",
							SerializedName: "primitive",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Ident",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "typeParam",
							SerializedName: "typeParam",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "Ident",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "reference",
							SerializedName: "reference",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Reference{
										V: ScopedName{
											ModuleName: "sys.adlast",
											Name:       "ScopedName",
										}},
								},
								Parameters: []TypeExpr{},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "TypeRef"},
		AST_TypeRef(),
	)
}

func Texpr_Union() ATypeExpr[Union] {
	return ATypeExpr[Union]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.adlast",
						Name:       "Union",
					},
				},
			},
			Parameters: []adlast.TypeExpr{},
		},
	}
}

func AST_Union() adlast.ScopedDecl {
	decl := Decl{
		Name: "Union",
		Version: types.Maybe[uint32]{
			Branch: types.Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: DeclType{
			Branch: DeclType_Struct_{
				V: Struct{
					TypeParams: []Ident{},
					Fields: []Field{
						Field{
							Name:           "typeParams",
							SerializedName: "typeParams",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Primitive{
										V: "Vector"},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Reference{
												V: ScopedName{
													ModuleName: "sys.adlast",
													Name:       "Ident",
												}},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
						Field{
							Name:           "fields",
							SerializedName: "fields",
							TypeExpr: TypeExpr{
								TypeRef: TypeRef{
									Branch: TypeRef_Primitive{
										V: "Vector"},
								},
								Parameters: []TypeExpr{
									TypeExpr{
										TypeRef: TypeRef{
											Branch: TypeRef_Reference{
												V: ScopedName{
													ModuleName: "sys.adlast",
													Name:       "Field",
												}},
										},
										Parameters: []TypeExpr{},
									},
								},
							},
							Default: types.Maybe[any]{
								Branch: types.Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: types.Map[ScopedName, any]([]types.MapEntry[ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.adlast",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.adlast", Name: "Union"},
		AST_Union(),
	)
}
