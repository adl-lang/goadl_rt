// Code generated by goadlc v3 - DO NOT EDIT.
package goadl

import (
	adlast "github.com/adl-lang/goadl_rt/v3/sys/adlast"
	. "github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_Either[T1 any, T2 any](t1 ATypeExpr[T1], t2 ATypeExpr[T2]) ATypeExpr[Either[T1, T2]] {
	return ATypeExpr[Either[T1, T2]]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.types",
						Name:       "Either",
					},
				},
			},
			Parameters: []adlast.TypeExpr{t1.Value, t2.Value},
		},
	}
}

func AST_Either() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "Either",
		Version: Maybe[uint32]{
			Branch: Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Union_{
				V: adlast.Union{
					TypeParams: []adlast.Ident{
						"T1",
						"T2",
					},
					Fields: []adlast.Field{
						adlast.Field{
							Name:           "left",
							SerializedName: "left",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_TypeParam{
										V: "T1"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: Maybe[any]{
								Branch: Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
						},
						adlast.Field{
							Name:           "right",
							SerializedName: "right",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_TypeParam{
										V: "T2"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: Maybe[any]{
								Branch: Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.types",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.types", Name: "Either"},
		AST_Either(),
	)
}

func Texpr_Map[K any, V any](k ATypeExpr[K], v ATypeExpr[V]) ATypeExpr[Map[K, V]] {
	return ATypeExpr[Map[K, V]]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.types",
						Name:       "Map",
					},
				},
			},
			Parameters: []adlast.TypeExpr{k.Value, v.Value},
		},
	}
}

func AST_Map() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "Map",
		Version: Maybe[uint32]{
			Branch: Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Newtype_{
				V: adlast.NewType{
					TypeParams: []adlast.Ident{
						"K",
						"V",
					},
					TypeExpr: adlast.TypeExpr{
						TypeRef: adlast.TypeRef{
							Branch: adlast.TypeRef_Primitive{
								V: "Vector"},
						},
						Parameters: []adlast.TypeExpr{
							adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Reference{
										V: adlast.ScopedName{
											ModuleName: "sys.types",
											Name:       "MapEntry",
										}},
								},
								Parameters: []adlast.TypeExpr{
									adlast.TypeExpr{
										TypeRef: adlast.TypeRef{
											Branch: adlast.TypeRef_TypeParam{
												V: "K"},
										},
										Parameters: []adlast.TypeExpr{},
									},
									adlast.TypeExpr{
										TypeRef: adlast.TypeRef{
											Branch: adlast.TypeRef_TypeParam{
												V: "V"},
										},
										Parameters: []adlast.TypeExpr{},
									},
								},
							},
						},
					},
					Default: Maybe[any]{
						Branch: Maybe_Nothing{
							V: struct{}{}},
					},
				}},
		},
		Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.types",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.types", Name: "Map"},
		AST_Map(),
	)
}

func Texpr_MapEntry[K any, V any](k ATypeExpr[K], v ATypeExpr[V]) ATypeExpr[MapEntry[K, V]] {
	return ATypeExpr[MapEntry[K, V]]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.types",
						Name:       "MapEntry",
					},
				},
			},
			Parameters: []adlast.TypeExpr{k.Value, v.Value},
		},
	}
}

func AST_MapEntry() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "MapEntry",
		Version: Maybe[uint32]{
			Branch: Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Struct_{
				V: adlast.Struct{
					TypeParams: []adlast.Ident{
						"K",
						"V",
					},
					Fields: []adlast.Field{
						adlast.Field{
							Name:           "key",
							SerializedName: "k",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_TypeParam{
										V: "K"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: Maybe[any]{
								Branch: Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
						},
						adlast.Field{
							Name:           "value",
							SerializedName: "v",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_TypeParam{
										V: "V"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: Maybe[any]{
								Branch: Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.types",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.types", Name: "MapEntry"},
		AST_MapEntry(),
	)
}

func Texpr_Maybe[T any](t ATypeExpr[T]) ATypeExpr[Maybe[T]] {
	return ATypeExpr[Maybe[T]]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.types",
						Name:       "Maybe",
					},
				},
			},
			Parameters: []adlast.TypeExpr{t.Value},
		},
	}
}

func AST_Maybe() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "Maybe",
		Version: Maybe[uint32]{
			Branch: Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Union_{
				V: adlast.Union{
					TypeParams: []adlast.Ident{
						"T",
					},
					Fields: []adlast.Field{
						adlast.Field{
							Name:           "nothing",
							SerializedName: "nothing",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_Primitive{
										V: "Void"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: Maybe[any]{
								Branch: Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
						},
						adlast.Field{
							Name:           "just",
							SerializedName: "just",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_TypeParam{
										V: "T"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: Maybe[any]{
								Branch: Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.types",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.types", Name: "Maybe"},
		AST_Maybe(),
	)
}

func Texpr_Pair[T1 any, T2 any](t1 ATypeExpr[T1], t2 ATypeExpr[T2]) ATypeExpr[Pair[T1, T2]] {
	return ATypeExpr[Pair[T1, T2]]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.types",
						Name:       "Pair",
					},
				},
			},
			Parameters: []adlast.TypeExpr{t1.Value, t2.Value},
		},
	}
}

func AST_Pair() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "Pair",
		Version: Maybe[uint32]{
			Branch: Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Struct_{
				V: adlast.Struct{
					TypeParams: []adlast.Ident{
						"T1",
						"T2",
					},
					Fields: []adlast.Field{
						adlast.Field{
							Name:           "v1",
							SerializedName: "v1",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_TypeParam{
										V: "T1"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: Maybe[any]{
								Branch: Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
						},
						adlast.Field{
							Name:           "v2",
							SerializedName: "v2",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_TypeParam{
										V: "T2"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: Maybe[any]{
								Branch: Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.types",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.types", Name: "Pair"},
		AST_Pair(),
	)
}

func Texpr_Result[T any, E any](t ATypeExpr[T], e ATypeExpr[E]) ATypeExpr[Result[T, E]] {
	return ATypeExpr[Result[T, E]]{
		Value: adlast.TypeExpr{
			TypeRef: adlast.TypeRef{
				Branch: adlast.TypeRef_Reference{
					V: adlast.ScopedName{
						ModuleName: "sys.types",
						Name:       "Result",
					},
				},
			},
			Parameters: []adlast.TypeExpr{t.Value, e.Value},
		},
	}
}

func AST_Result() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "Result",
		Version: Maybe[uint32]{
			Branch: Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Union_{
				V: adlast.Union{
					TypeParams: []adlast.Ident{
						"T",
						"E",
					},
					Fields: []adlast.Field{
						adlast.Field{
							Name:           "ok",
							SerializedName: "ok",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_TypeParam{
										V: "T"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: Maybe[any]{
								Branch: Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
						},
						adlast.Field{
							Name:           "error",
							SerializedName: "error",
							TypeExpr: adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_TypeParam{
										V: "E"},
								},
								Parameters: []adlast.TypeExpr{},
							},
							Default: Maybe[any]{
								Branch: Maybe_Nothing{
									V: struct{}{}},
							},
							Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
						},
					},
				}},
		},
		Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.types",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.types", Name: "Result"},
		AST_Result(),
	)
}

func AST_Set() adlast.ScopedDecl {
	decl := adlast.Decl{
		Name: "Set",
		Version: Maybe[uint32]{
			Branch: Maybe_Nothing{
				V: struct{}{}},
		},
		Type_: adlast.DeclType{
			Branch: adlast.DeclType_Newtype_{
				V: adlast.NewType{
					TypeParams: []adlast.Ident{
						"T",
					},
					TypeExpr: adlast.TypeExpr{
						TypeRef: adlast.TypeRef{
							Branch: adlast.TypeRef_Primitive{
								V: "Vector"},
						},
						Parameters: []adlast.TypeExpr{
							adlast.TypeExpr{
								TypeRef: adlast.TypeRef{
									Branch: adlast.TypeRef_TypeParam{
										V: "T"},
								},
								Parameters: []adlast.TypeExpr{},
							},
						},
					},
					Default: Maybe[any]{
						Branch: Maybe_Nothing{
							V: struct{}{}},
					},
				}},
		},
		Annotations: Map[adlast.ScopedName, any]([]MapEntry[adlast.ScopedName, any]{
			MapEntry[adlast.ScopedName, any]{
				Key: adlast.ScopedName{
					ModuleName: "adlc.config.go_",
					Name:       "GoCustomType",
				},
				Value: map[string]interface{}{"gotype": map[string]interface{}{"import_path": "github.com/adl-lang/goadl_rt/v3", "name": "MapSet", "pkg": "goadl"}, "helpers": map[string]interface{}{"import_path": "github.com/adl-lang/goadl_rt/v3", "name": "SetHelper", "pkg": "goadl"}},
			},
		}),
	}
	return adlast.ScopedDecl{
		ModuleName: "sys.types",
		Decl:       decl,
	}
}

func init() {
	RESOLVER.Register(
		adlast.ScopedName{ModuleName: "sys.types", Name: "Set"},
		AST_Set(),
	)
	RESOLVER.RegisterHelper(
		adlast.ScopedName{ModuleName: "sys.types", Name: "Set"},
		(*SetHelper)(nil),
	)
}
