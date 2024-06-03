// Code generated by goadlc v3 - DO NOT EDIT.
package java

import (
	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/customtypes"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/adl-lang/goadl_rt/v3/sys/types"
)

func Texpr_JavaCustomType() adlast.ATypeExpr[JavaCustomType] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("adlc.config.java", "JavaCustomType"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[JavaCustomType](te)
}

func AST_JavaCustomType() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"JavaCustomType",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_struct_(
			adlast.MakeAll_Struct(
				[]adlast.Ident{},
				[]adlast.Field{
					adlast.MakeAll_Field(
						"javaname",
						"javaname",
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
						"ctorjavaname",
						"ctorjavaname",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"String",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_just[any](
							"",
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
					adlast.MakeAll_Field(
						"helpers",
						"helpers",
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
						"generateType",
						"generateType",
						adlast.MakeAll_TypeExpr(
							adlast.Make_TypeRef_primitive(
								"Bool",
							),
							[]adlast.TypeExpr{},
						),
						types.Make_Maybe_just[any](
							false,
						),
						customtypes.MapMap[adlast.ScopedName, any]{},
					),
				},
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "ADL declaration annotation to specify that a custom type\nshould be used\n"},
	)
	return adlast.Make_ScopedDecl("adlc.config.java", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("adlc.config.java", "JavaCustomType"),
		AST_JavaCustomType(),
	)
}

func Texpr_JavaGenerate() adlast.ATypeExpr[JavaGenerate] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("adlc.config.java", "JavaGenerate"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[JavaGenerate](te)
}

func AST_JavaGenerate() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"JavaGenerate",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_type_(
			adlast.MakeAll_TypeDef(
				[]adlast.Ident{},
				adlast.MakeAll_TypeExpr(
					adlast.Make_TypeRef_primitive(
						"Bool",
					),
					[]adlast.TypeExpr{},
				),
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "ADL module or declaration annotation to control\nwhether code is actually generated.\n"},
	)
	return adlast.Make_ScopedDecl("adlc.config.java", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("adlc.config.java", "JavaGenerate"),
		AST_JavaGenerate(),
	)
}

func Texpr_JavaPackage() adlast.ATypeExpr[JavaPackage] {
	te := adlast.Make_TypeExpr(
		adlast.Make_TypeRef_reference(
			adlast.Make_ScopedName("adlc.config.java", "JavaPackage"),
		),
		[]adlast.TypeExpr{},
	)
	return adlast.Make_ATypeExpr[JavaPackage](te)
}

func AST_JavaPackage() adlast.ScopedDecl {
	decl := adlast.MakeAll_Decl(
		"JavaPackage",
		types.Make_Maybe_nothing[uint32](),
		adlast.Make_DeclType_type_(
			adlast.MakeAll_TypeDef(
				[]adlast.Ident{},
				adlast.MakeAll_TypeExpr(
					adlast.Make_TypeRef_primitive(
						"String",
					),
					[]adlast.TypeExpr{},
				),
			),
		),
		customtypes.MapMap[adlast.ScopedName, any]{adlast.Make_ScopedName("sys.annotations", "Doc"): "ADL module annotation to specify a target java package.\n"},
	)
	return adlast.Make_ScopedDecl("adlc.config.java", decl)
}

func init() {
	goadl.RESOLVER.Register(
		adlast.Make_ScopedName("adlc.config.java", "JavaPackage"),
		AST_JavaPackage(),
	)
}