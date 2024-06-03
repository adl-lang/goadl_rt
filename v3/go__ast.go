// Code generated by goadlc v3 - DO NOT EDIT.
package goadl

import (
	. "github.com/adl-lang/goadl_rt/v3/adlc/config/go_"
	"github.com/adl-lang/goadl_rt/v3/customtypes"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_GoCustomType() adlast.ATypeExpr[GoCustomType] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("adlc.config.go_", "GoCustomType"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[GoCustomType](te)
}

func AST_GoCustomType() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"GoCustomType",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"gotype",
						"gotype",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"adlc.config.go_",
									"GoType",
								),
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "The Go struct to use\n\nNote currently interfaces are not supported\n"},
					),
					adlast.MakeAll_Field(
						"helpers",
						"helpers",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_reference(
								adlast.MakeAll_ScopedName(
									"adlc.config.go_",
									"GoHelperType",
								),
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "The Go structed which implements json Marshal and Unmarshal for the gotype.\n\nNote Marshal and Unmarshal must be implemented for the specified struct's nil ptr\n"},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "ADL declaration annotation to specify the custom type to use\n"},
	)
	return adlast.Make_ScopedDecl("adlc.config.go_", decl)
}

func init() {
	RESOLVER.Register(
		adlast.Make_ScopedName("adlc.config.go_", "GoCustomType"),
		AST_GoCustomType(),
	)
}

func Texpr_GoHelperType() adlast.ATypeExpr[GoHelperType] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("adlc.config.go_", "GoHelperType"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[GoHelperType](te)
}

func AST_GoHelperType() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"GoHelperType",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"name",
						"name",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"pkg",
						"pkg",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"import_path",
						"import_path",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("adlc.config.go_", decl)
}

func init() {
	RESOLVER.Register(
		adlast.Make_ScopedName("adlc.config.go_", "GoHelperType"),
		AST_GoHelperType(),
	)
}

func Texpr_GoType() adlast.ATypeExpr[GoType] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("adlc.config.go_", "GoType"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[GoType](te)
}

func AST_GoType() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"GoType",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"name",
						"name",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"pkg",
						"pkg",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"type_constraints",
						"type_constraints",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Vector",
							),
							[]adlast.TypeExpr{
								adlast.MakeAll_TypeExpr(
									adlast.Make_TypeRef_primitive(
										"String",
									),
									[]adlast.TypeExpr{},
								),
							},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "list of type constraints e.g. for Map this would be [comparable, any]\n"},
					),
					adlast.MakeAll_Field(
						"import_path",
						"import_path",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_nothing[any](),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{},
	)
	return adlast.Make_ScopedDecl("adlc.config.go_", decl)
}

func init() {
	RESOLVER.Register(
		adlast.Make_ScopedName("adlc.config.go_", "GoType"),
		AST_GoType(),
	)
}