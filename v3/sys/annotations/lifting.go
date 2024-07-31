package annotations

import (
	"fmt"
	"strings"
	"sync"

	goadl "github.com/adl-lang/goadl_rt/v3"
	"github.com/adl-lang/goadl_rt/v3/sys/adlast"
	"github.com/samber/lo"
)

type Json = any
type JsonArray = []Json
type JsonObject = map[string]Json

type Lifter func(Json) (Json, error)

var idLifter = func(json Json) (Json, error) { return json, nil }

func CreateLifter(
	resolver *goadl.ResolverType,
	texpr adlast.TypeExpr,
) Lifter {
	return buildLifter(resolver, texpr, map[string]lftr_texpr{})
}

var lifterCache sync.Map

func texpr2Key(
	te adlast.TypeExpr,
) string {
	sb := strings.Builder{}
	sb.Grow(100)
	var recurse func(te adlast.TypeExpr)
	recurse = func(te adlast.TypeExpr) {
		adlast.Handle_TypeRef[*struct{}](
			te.TypeRef,
			func(primitive string) *struct{} {
				sb.WriteString(primitive + ":")
				if len(te.Parameters) == 1 {
					recurse(te.Parameters[0])
				}
				return nil
			},
			func(typeParam string) *struct{} {
				panic(fmt.Errorf("%s", typeParam))
			},
			func(reference adlast.ScopedName) *struct{} {
				sb.WriteString(reference.ModuleName + "." + reference.Name + "::")
				for i := range te.Parameters {
					if i != 0 {
						sb.WriteString(",")
					}
					recurse(te.Parameters[i])
				}
				return nil
			},
			nil,
		)
	}
	recurse(te)
	return sb.String()
}

type lftr_texpr struct {
	lftr  Lifter
	tepxr adlast.TypeExpr
}

func buildLifter(
	resolver *goadl.ResolverType,
	texpr adlast.TypeExpr,
	boundTypeParams map[string]lftr_texpr,
) Lifter {
	if tp, ok := texpr.TypeRef.Cast_typeParam(); ok {
		return boundTypeParams[tp].lftr
	}
	key := texpr2Key(texpr)
	// taken from golang stdlib src/encoding/json/encode.go
	if fi, ok := lifterCache.Load(key); ok {
		return fi.(Lifter)
	}
	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  Lifter
	)
	wg.Add(1)
	fi, loaded := lifterCache.LoadOrStore(key, Lifter(func(json Json) (Json, error) {
		wg.Wait()
		return f(json)
	}))
	if loaded {
		return fi.(Lifter)
	}

	// Compute the real encoder and replace the indirect func with it.
	f = buildLifter0(resolver, texpr, boundTypeParams)
	wg.Done()
	lifterCache.Store(key, f)
	return f
}

func buildLifter0(
	resolver *goadl.ResolverType,
	texpr adlast.TypeExpr,
	boundTypeParams map[string]lftr_texpr,
) Lifter {
	if !hasTypeDiscrimination(resolver, texpr) {
		return idLifter
	}
	return adlast.Handle_TypeRef[Lifter](
		texpr.TypeRef,
		func(primitive string) Lifter {
			if len(texpr.Parameters) == 0 {
				return idLifter
			}
			elem_lifter := buildLifter(resolver, texpr.Parameters[0], boundTypeParams)
			switch primitive {
			case "Nullable":
				return func(j Json) (Json, error) {
					if j == nil {
						return nil, nil
					}
					return elem_lifter(j)
				}
			case "Vector":
				return func(j Json) (Json, error) {
					if ja, ok := j.(JsonArray); ok {
						res := make([]Json, len(ja))
						var err error
						for i := range ja {
							res[i], err = elem_lifter(ja[i])
							if err != nil {
								return nil, err
							}
						}
						return res, nil
					}
					return nil, fmt.Errorf("expected arrays got %T", j)
				}
			case "StringMap":
				return func(j Json) (Json, error) {
					if jo, ok := j.(JsonObject); ok {
						res := make(JsonObject)
						var err error
						for k, v := range jo {
							res[k], err = elem_lifter(v)
							if err != nil {
								return nil, err
							}
						}
						return res, nil
					}
					return nil, fmt.Errorf("expected object got %T", j)
				}
			}
			panic(fmt.Errorf("unexpected primitive %s", primitive))
		},
		func(typeParam string) Lifter {
			return boundTypeParams[typeParam].lftr
		},
		func(reference adlast.ScopedName) Lifter {
			ast := resolver.Resolve(reference)
			return adlast.Handle_DeclType[Lifter](
				ast.Decl.Type_,
				func(struct_ adlast.Struct) Lifter {
					return buildStructLifter(resolver, struct_, texpr, boundTypeParams)
				},
				func(union_ adlast.Union) Lifter {
					if goadl.IsEnum(union_) {
						return idLifter
					}
					return buildUnionLifter(resolver, union_, texpr, boundTypeParams)
				},
				func(type_ adlast.TypeDef) Lifter {
					newBoundTypeParams := BindTypeParams(
						type_.TypeParams,
						texpr.Parameters,
						func(te adlast.TypeExpr) lftr_texpr {
							return lftr_texpr{buildLifter(resolver, te, boundTypeParams), te}
						},
					)
					return buildLifter(resolver, type_.TypeExpr, newBoundTypeParams)
				},
				func(newtype_ adlast.NewType) Lifter {
					newBoundTypeParams := BindTypeParams(
						newtype_.TypeParams,
						texpr.Parameters,
						func(te adlast.TypeExpr) lftr_texpr {
							return lftr_texpr{buildLifter(resolver, te, boundTypeParams), te}
						},
					)
					return buildLifter(resolver, newtype_.TypeExpr, newBoundTypeParams)
				},
				nil,
			)
		},
		nil,
	)
}

func BindTypeParams[T any](
	paramNames []string,
	paramTypes []adlast.TypeExpr,
	fn func(adlast.TypeExpr) T,
) map[string]T {
	result := map[string]T{}
	for i := range paramNames {
		result[paramNames[i]] = fn(paramTypes[i])
	}
	return result
}

func buildStructLifter(
	resolver *goadl.ResolverType,
	struct_ adlast.Struct,
	texpr adlast.TypeExpr,
	boundTypeParams map[string]lftr_texpr,
) Lifter {
	newBoundTypeParams := BindTypeParams(
		struct_.TypeParams,
		texpr.Parameters,
		func(te adlast.TypeExpr) lftr_texpr {
			return lftr_texpr{buildLifter(resolver, te, boundTypeParams), te}
		},
	)
	fieldDetails := map[string]func() Lifter{}
	for _, fld := range struct_.Fields {
		fieldDetails[fld.SerializedName] = once(func() Lifter { return buildLifter(resolver, fld.TypeExpr, newBoundTypeParams) })
	}
	return func(j Json) (Json, error) {
		if jo, ok := j.(JsonObject); ok {
			jv2 := map[string]Json{}
			var err error
			for k, v := range jo {
				if elem_lifter, exists := fieldDetails[k]; exists {
					jv2[k], err = elem_lifter()(v)
					if err != nil {
						return nil, err
					}
				} else {
					// keep field not defined in adl
					jv2[k] = v
				}
			}
			return jv2, nil
		}
		return nil, fmt.Errorf("expected object got %T", j)
	}
}

type Ancestor struct {
	Name       string
	MaxVersion int
}

type UnionFieldDetails struct {
	Ancestors  []Ancestor
	MaxVersion int
	Lifter     func() Lifter
}

type TypeDisc struct {
	Field adlast.Field
	Disc  TypeDiscrimination
}

func buildUnionLifter(
	resolver *goadl.ResolverType,
	union_ adlast.Union,
	texpr adlast.TypeExpr,
	boundTypeParams map[string]lftr_texpr,
) Lifter {
	newBoundTypeParams := BindTypeParams(
		union_.TypeParams,
		texpr.Parameters,
		func(te adlast.TypeExpr) lftr_texpr {
			return lftr_texpr{buildLifter(resolver, te, boundTypeParams), te}
		},
	)
	fields := make(map[string]UnionFieldDetails)
	for _, fld := range union_.Fields {
		lifter := once(func() Lifter { return buildLifter(resolver, fld.TypeExpr, newBoundTypeParams) })
		fields[fld.SerializedName] = UnionFieldDetails{
			Ancestors:  []Ancestor{},
			MaxVersion: -1,
			Lifter:     lifter,
		}
	}
	typeDiscs := []TypeDisc{}
	jb := goadl.CreateJsonDecodeBinding(Texpr_TypeDiscrimination(), goadl.RESOLVER)

	for _, fld := range union_.Fields {
		disc, err := goadl.GetAnnotation(fld.Annotations, tdSN, jb)
		if err != nil {
			panic(err)
		}
		if disc != nil {
			typeDiscs = append(typeDiscs, TypeDisc{Field: fld, Disc: *disc})
		}
	}
	for _, fld := range union_.Fields {
		anc := Ancestor{Name: fld.SerializedName, MaxVersion: -1}
		transDiscs := transitiveTypeDisc(
			resolver,
			fld.TypeExpr,
			[]Ancestor{anc},
			newBoundTypeParams,
			jb,
			map[string]struct{}{},
		)
		for _, td := range transDiscs {
			fields[td.fld.Field.SerializedName] = td.ufd
			typeDiscs = append(typeDiscs, td.fld)
		}
	}
	return buildLiftUnion(resolver, typeDiscs, fields)
}

func buildLiftUnion(
	resolver *goadl.ResolverType,
	type_discs []TypeDisc,
	fields map[string]UnionFieldDetails,
) Lifter {
	return func(json0 Json) (Json, error) {
		mtd := []TypeDisc{}
		for _, td := range type_discs {
			expanded_texpr := expandTypes(resolver, td.Field.TypeExpr, map[string]adlast.TypeExpr{})
			mt, err := matchTypeDiscrimination(resolver, json0, expanded_texpr)
			if err != nil {
				return nil, err
			}
			if mt {
				mtd = append(mtd, td)
			}
		}
		if len(mtd) > 1 {
			return nil, fmt.Errorf(`ambiguous matching type discriminators %v`, lo.Map[TypeDisc, string](mtd, func(item TypeDisc, index int) string {
				return item.Field.Name
			}))
		}
		if len(mtd) == 1 {
			json1 := map[string]any{}
			json1[mtd[0].Field.SerializedName] = json0
			// json1["@v"] = mtd[0].disc.version
			json0 = json1
		}
		if json1, ok := json0.(map[string]any); ok {
			keys := lo.Filter[string](lo.Keys(json1), func(item string, index int) bool { return item != "@v" })
			if len(keys) != 1 {
				return nil, fmt.Errorf("not the shape of a union")
			}
			if ufd, exist := fields[keys[0]]; !exist {
				return nil, fmt.Errorf("branch not defined '${keys[0]}'\n%v\n%v", json0, json1)
			} else {
				if ufd.MaxVersion > -1 {
					json1["@v"] = ufd.MaxVersion
				}
				var err error
				json1[keys[0]], err = ufd.Lifter()(json1[keys[0]])
				if err != nil {
					return nil, err
				}
				if len(ufd.Ancestors) > 0 {
					for _, an := range ufd.Ancestors {
						j3 := map[string]any{}
						j3[an.Name] = json1
						j3["@v"] = an.MaxVersion
						json1 = j3
					}
				}
				return json1, nil
			}
		} else {
			return nil, fmt.Errorf("expecting union, value isn't even an object\n%v\n%v", json0, json1)
		}
	}
}

func matchTypeDiscrimination(
	resolver *goadl.ResolverType,
	json Json,
	texpr adlast.TypeExpr,
) (bool, error) {
	typeRef := texpr.TypeRef
	if _, ok := texpr.TypeRef.Cast_typeParam(); ok {
		return false, nil
	}
	if primitive, ok := typeRef.Cast_primitive(); ok && (primitive == "Json" || primitive == "Void") {
		return false, fmt.Errorf(`cannot use Json or Void as a type discriminator`)
	}
	if json == nil {
		if primitive, ok := typeRef.Cast_primitive(); ok && primitive == "Nullable" {
			return true, nil
		}
		return false, fmt.Errorf(`primitive type mismatch. expected "Nullable" received %v`, typeRef)
	}
	if aj, ok := json.(JsonArray); ok {
		if primitive, ok := typeRef.Cast_primitive(); !ok || primitive != "Vector" {
			return false, nil
		}
		for i := range aj {
			if ok, err := matchTypeDiscrimination(resolver, aj[i], texpr.Parameters[0]); !ok && err == nil {
				return false, nil
			} else if err != nil {
				return false, err
			}
		}
		return true, nil
	}
	if prim, ok := typeRef.Cast_primitive(); ok && prim == "Nullable" {
		typeRef = texpr.Parameters[0].TypeRef
		if prim, ok := typeRef.Cast_primitive(); ok {
			if prim == "Vector" {
				return false, fmt.Errorf("lifting of Nullable<Vector<>> not implemented")
			}
			if prim == "Nullable" {
				return false, fmt.Errorf("lifting of Nullable<Nullable<>> not implemented")
			}
		}
	}
	switch v := json.(type) {
	case string:
		if primitive, ok := typeRef.Cast_primitive(); ok && primitive == "String" {
			return true, nil
		}
		return false, nil
	case float64:
		if primitive, ok := typeRef.Cast_primitive(); ok && lo.Contains(adlNumbers, primitive) {
			return true, nil
		}
		return false, nil
	case bool:
		if primitive, ok := typeRef.Cast_primitive(); ok && primitive == "Bool" {
			return true, nil
		}
		return false, nil
	case map[string]any:
		return matchObject(resolver, v, texpr), nil
	}
	return false, nil
}

func matchObject(
	resolver *goadl.ResolverType,
	json JsonObject,
	expanded_texpr adlast.TypeExpr,
) bool {
	typeRef := expanded_texpr.TypeRef
	if pr, ok := typeRef.Cast_primitive(); ok && pr == "StringMap" {
		// TODO check stringmap objects
		return true
	}
	if ref, ok := typeRef.Cast_reference(); ok {
		sd := resolver.Resolve(ref)
		if st, ok := sd.Decl.Type_.Cast_struct_(); ok {
			for _, fld := range st.Fields {
				if _, ok0 := fld.Default.Cast_nothing(); ok0 {
					if _, ex := json[fld.SerializedName]; !ex {
						return false
					}
				}
			}
			return true
		}
		if un, ok2 := sd.Decl.Type_.Cast_union_(); ok2 {
			keys := lo.Filter[string](lo.Keys(json), func(item string, index int) bool { return item != "@v" })
			if len(keys) == 1 {
				for _, fld := range un.Fields {
					if fld.SerializedName == keys[0] {
						return true
					}
				}
			}
			return false
		}
	}
	return false
}

var adlNumbers = []string{
	"Int8",
	"Int16",
	"Int32",
	"Int64",
	"Word8",
	"Word16",
	"Word32",
	"Word64",
	"Float",
	"Double",
}

func expandTypes(
	resolver *goadl.ResolverType,
	texpr adlast.TypeExpr,
	boundTypeParams map[string]adlast.TypeExpr,
) adlast.TypeExpr {
	return adlast.Handle_TypeRef[adlast.TypeExpr](
		texpr.TypeRef,
		func(primitive string) adlast.TypeExpr {
			if len(texpr.Parameters) == 0 {
				return texpr
			}
			te0 := expandTypes(resolver, texpr.Parameters[0], boundTypeParams)
			return adlast.Make_TypeExpr(texpr.TypeRef, []adlast.TypeExpr{te0})
		},
		func(typeParam string) adlast.TypeExpr {
			return boundTypeParams[typeParam]
		},
		func(reference adlast.ScopedName) adlast.TypeExpr {
			ast := resolver.Resolve(reference)
			return adlast.Handle_DeclType[adlast.TypeExpr](
				ast.Decl.Type_,
				func(struct_ adlast.Struct) adlast.TypeExpr {
					if len(texpr.Parameters) == 0 {
						return texpr
					}
					return texpr
					// TODO need to sub fields?
				},
				func(union_ adlast.Union) adlast.TypeExpr {
					if len(texpr.Parameters) == 0 {
						return texpr
					}
					return texpr
					// TODO need to sub fields?
				},
				func(type_ adlast.TypeDef) adlast.TypeExpr {
					nbp := BindTypeParams(type_.TypeParams, texpr.Parameters, func(te adlast.TypeExpr) adlast.TypeExpr { return te })
					return expandTypes(resolver, type_.TypeExpr, nbp)
				},
				func(newtype_ adlast.NewType) adlast.TypeExpr {
					nbp := BindTypeParams(newtype_.TypeParams, texpr.Parameters, func(te adlast.TypeExpr) adlast.TypeExpr { return te })
					return expandTypes(resolver, newtype_.TypeExpr, nbp)
				},
				nil,
			)
		},
		nil,
	)
}

var tdSN = adlast.Make_ScopedName("sys.annotations", "TypeDiscrimination")

type ttdR struct {
	fld TypeDisc
	ufd UnionFieldDetails
}

func transitiveTypeDisc(
	resolver *goadl.ResolverType,
	ftexpr adlast.TypeExpr,
	ancestors []Ancestor,
	boundTypeParams map[string]lftr_texpr,
	jb goadl.JsonDecodeBinder[TypeDiscrimination],
	seen map[string]struct{},
) []ttdR {
	return adlast.Handle_TypeRef[[]ttdR](
		ftexpr.TypeRef,
		func(primitive string) []ttdR {
			return []ttdR{}
		},
		func(typeParam string) []ttdR {
			return []ttdR{}
		},
		func(reference adlast.ScopedName) []ttdR {
			ast := resolver.Resolve(reference)
			return adlast.Handle_DeclType[[]ttdR](
				ast.Decl.Type_,
				func(struct_ adlast.Struct) []ttdR {
					return []ttdR{}
				},
				func(union_ adlast.Union) []ttdR {
					ret := []ttdR{}
					newBoundTypeParams := BindTypeParams(
						union_.TypeParams,
						ftexpr.Parameters,
						func(te adlast.TypeExpr) lftr_texpr {
							return lftr_texpr{buildLifter(resolver, te, boundTypeParams), te}
						},
					)
					max_version := lo.Reduce[adlast.Field, int](union_.Fields,
						func(agg int, item adlast.Field, index int) int {
							disc, err := goadl.GetAnnotation(item.Annotations, tdSN, jb)
							if err != nil {
								panic(err)
							}
							if disc != nil {
								if int(disc.Version) > agg {
									return agg
								}
							}
							return agg
						},
						-1,
					)
					for _, fld := range union_.Fields {
						te := fld.TypeExpr
						if tp, ok := fld.TypeExpr.TypeRef.Cast_typeParam(); ok {
							te = newBoundTypeParams[tp].tepxr
						}
						if _, ex := seen[texpr2Key(te)]; ex {
							continue
						}
						disc, err := goadl.GetAnnotation(fld.Annotations, tdSN, jb)
						if err != nil {
							panic(err)
						}
						if disc != nil {
							bldr := once(func() Lifter { return buildLifter(resolver, fld.TypeExpr, newBoundTypeParams) })
							ret = append(ret, ttdR{
								fld: TypeDisc{
									Field: fld,
									Disc:  *disc,
								},
								ufd: UnionFieldDetails{
									MaxVersion: max_version,
									Ancestors:  ancestors,
									Lifter:     bldr,
								},
							})
						}
						parent := Ancestor{
							Name:       fld.SerializedName,
							MaxVersion: max_version,
						}
						ancestors0 := make([]Ancestor, 1, len(ancestors)+1)
						ancestors0[0] = parent
						ancestors0 = append(ancestors0, ancestors...)
						seen[texpr2Key(te)] = struct{}{}
						decendants := transitiveTypeDisc(
							resolver,
							fld.TypeExpr,
							ancestors0,
							newBoundTypeParams,
							jb,
							seen,
						)
						ret = append(ret, decendants...)
					}
					return ret
				},
				func(type_ adlast.TypeDef) []ttdR {
					return transitiveTypeDisc(
						resolver,
						type_.TypeExpr,
						ancestors,
						boundTypeParams,
						jb,
						seen,
					)
				},
				func(newtype_ adlast.NewType) []ttdR {
					return transitiveTypeDisc(
						resolver,
						newtype_.TypeExpr,
						ancestors,
						boundTypeParams,
						jb,
						seen,
					)
				},
				nil,
			)
		},
		nil,
	)
}

func hasTypeDiscrimination(
	resolver *goadl.ResolverType,
	texpr adlast.TypeExpr,
) bool {
	return adlast.Handle_TypeRef[bool](
		texpr.TypeRef,
		func(primitive string) bool {
			if len(texpr.Parameters) == 0 {
				return false
			}
			return hasTypeDiscrimination(resolver, texpr.Parameters[0])
		},
		func(typeParam string) bool {
			// todo use a binder to see if this is really needed
			return true
		},
		func(reference adlast.ScopedName) bool {
			ast := resolver.Resolve(reference)
			return adlast.Handle_DeclType[bool](
				ast.Decl.Type_,
				func(struct_ adlast.Struct) bool {
					for _, fld := range struct_.Fields {
						if hasTypeDiscrimination(resolver, fld.TypeExpr) {
							return true
						}
					}
					return false
				},
				func(union_ adlast.Union) bool {
					for _, fld := range union_.Fields {
						if goadl.HasAnnotation(fld.Annotations, tdSN) {
							return true
						}
					}
					for _, fld := range union_.Fields {
						if hasTypeDiscrimination(resolver, fld.TypeExpr) {
							return true
						}
					}
					return false
				},
				func(type_ adlast.TypeDef) bool {
					return hasTypeDiscrimination(resolver, type_.TypeExpr)
				},
				func(newtype_ adlast.NewType) bool {
					return hasTypeDiscrimination(resolver, newtype_.TypeExpr)
				},
				nil,
			)
		},
		nil,
	)
}

type Once[T any] func(fn func() T) func() T

func once[T any](fn func() T) func() T {
	var result *T = nil
	return func() T {
		if result == nil {
			result = goadl.Addr(fn())
		}
		return *result
	}
}
