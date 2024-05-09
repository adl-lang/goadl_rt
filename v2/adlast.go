package goadl

import (
	"encoding/json"
	"fmt"
)

type ADL struct {
	Modules []*Module
}

func (a ADL) String() string {
	return fmt.Sprintf("%v", a.Modules)
}

type Annotations []Annotation

type Annotation struct {
	Key ScopedName  `json:"v1"`
	Val interface{} `json:"v2"`
}

type ScopedName struct {
	ModuleName string `json:"moduleName"`
	Name       string `json:"name"`
}

// union TypeRef
//
//	{
//	    Ident primitive;
//	    Ident typeParam;
//	    ScopedName reference;
//	};
type TypeRef struct {
	Branch TypeRefBranch
}

// func (obj *TypeRef) Branch() TypeRefBranch {
// 	return obj.branch
// }

type TypeRefBranch interface {
	isTypeRefBranch()
}
type TypeRefBranch_Primitive string
type TypeRefBranch_TypeParam string
type TypeRefBranch_Reference ScopedName

func (TypeRefBranch_Primitive) isTypeRefBranch() {}
func (TypeRefBranch_TypeParam) isTypeRefBranch() {}
func (TypeRefBranch_Reference) isTypeRefBranch() {}

func (*TypeRef) Primitive(primitive string) TypeRef {
	return TypeRef{
		TypeRefBranch_Primitive(primitive),
	}
}
func (*TypeRef) TypeParam(typeParam string) TypeRef {
	return TypeRef{
		TypeRefBranch_TypeParam(typeParam),
	}
}
func (*TypeRef) Reference(scopedName ScopedName) TypeRef {
	return TypeRef{
		TypeRefBranch_Reference(scopedName),
	}
}

func HandleTypeRef[T any](
	in TypeRefBranch,
	primitive func(TypeRefBranch_Primitive) (T, error),
	typeParam func(TypeRefBranch_TypeParam) (T, error),
	reference func(TypeRefBranch_Reference) (T, error),
) (T, error) {
	switch b := in.(type) {
	case TypeRefBranch_Primitive:
		return primitive(b)
	case TypeRefBranch_TypeParam:
		return typeParam(b)
	case TypeRefBranch_Reference:
		return reference(b)
	}
	panic(fmt.Sprintf("code gen error unhandled branch '%#v", in))
}

func (obj *TypeRef) UnmarshalJSON(b []byte) error {
	key, val, err := SplitKV(b)
	if err != nil {
		return err
	}
	switch key {
	case "primitive":
		obj.Branch = TypeRefBranch_Primitive(string(val[1 : len(val)-1]))
	case "typeParam":
		obj.Branch = TypeRefBranch_TypeParam(string(val[1 : len(val)-1]))
	case "reference":
		sn := ScopedName{}
		err = json.Unmarshal(val, &sn)
		if err != nil {
			return err
		}
		obj.Branch = TypeRefBranch_Reference(sn)
	}
	return nil
	// fmt.Printf("TypeRef) UnmarshalJSON '%s'\n", string(b))
	// return fmt.Errorf("ads")
}

func (obj TypeRef) MarshalJSON() ([]byte, error) {
	ba, err := HandleTypeRef(
		obj.Branch,
		func(primitive TypeRefBranch_Primitive) ([]byte, error) {
			return []byte(`{"primitive" : "` + primitive + `"}`), nil
		},
		func(typeParam TypeRefBranch_TypeParam) ([]byte, error) {
			return []byte(`{"typeParam" : "` + typeParam + `"}`), nil
		},
		func(reference TypeRefBranch_Reference) ([]byte, error) {
			ba, err := json.Marshal(&reference)
			if err != nil {
				return nil, err
			}
			ba2 := make([]byte, len(ba)+len(`{"reference" : `+`}`))
			copy(ba2, []byte(`{"reference" : `))
			copy(ba2[len(`{"reference" : `):], ba)
			ba2[len(ba2)-1] = '}'
			return ba2, nil
		},
	)
	if err != nil {
		return nil, err
	}
	return ba, nil
}

type TypeExpr struct {
	TypeRef    TypeRef    `json:"typeRef"`
	Parameters []TypeExpr `json:"parameters"`
}

type Field struct {
	Name           string             `json:"name"`
	SerializedName string             `json:"serializedName"`
	TypeExpr       TypeExpr           `json:"typeExpr"`
	Default        Maybe[interface{}] `json:"default,omitempty"`
	Annotations    `json:"annotations"`
}

// Struct & Union
type Name struct {
	TypeParams []string `json:"typeParams"`
	Field      []Field  `json:"fields"`
}

type Struct struct {
	TypeParams []string `json:"typeParams"`
	Fields     []Field  `json:"fields"`
}

type Union struct {
	TypeParams []string `json:"typeParams"`
	Fields     []Field  `json:"fields"`
}

type TypeDef struct {
	TypeParams []string `json:"typeParams"`
	TypeExpr   TypeExpr `json:"typeExpr"`
}

type NewType struct {
	TypeParams []string    `json:"typeParams"`
	TypeExpr   TypeExpr    `json:"typeExpr"`
	Default    interface{} `json:"default,omitempty"`
}

// union DeclType
//
//	{
//	    Struct struct_;
//	    Union union_;
//	    TypeDef type_;
//	    NewType newtype_;
//	};
type DeclType struct {
	Branch DeclTypeBranch
}

// func (obj *DeclType) Branch() DeclTypeBranch {
// 	return obj.branch
// }

type DeclTypeBranch interface {
	isDeclType()
}
type DeclTypeBranch_Struct_ Struct
type DeclTypeBranch_Union_ Union
type DeclTypeBranch_Type_ TypeDef
type DeclTypeBranch_Newtype_ NewType

func (DeclTypeBranch_Struct_) isDeclType()  {}
func (DeclTypeBranch_Union_) isDeclType()   {}
func (DeclTypeBranch_Type_) isDeclType()    {}
func (DeclTypeBranch_Newtype_) isDeclType() {}

func (*DeclType) Struct_(struct_ Struct) DeclType {
	return DeclType{DeclTypeBranch_Struct_(struct_)}
}
func (*DeclType) Union_(union_ Union) DeclType {
	return DeclType{DeclTypeBranch_Union_(union_)}
}
func (*DeclType) Type_(type_ TypeDef) DeclType {
	return DeclType{DeclTypeBranch_Type_(type_)}
}
func (*DeclType) Newtype_(newtype_ NewType) DeclType {
	return DeclType{DeclTypeBranch_Newtype_(newtype_)}
}

func HandleDeclType[T any](
	in DeclTypeBranch,
	struct_ func(DeclTypeBranch_Struct_) (T, error),
	union_ func(DeclTypeBranch_Union_) (T, error),
	type_ func(DeclTypeBranch_Type_) (T, error),
	newtype_ func(DeclTypeBranch_Newtype_) (T, error),
) (T, error) {
	switch b := in.(type) {
	case DeclTypeBranch_Struct_:
		return struct_(b)
	case DeclTypeBranch_Union_:
		return union_(b)
	case DeclTypeBranch_Type_:
		return type_(b)
	case DeclTypeBranch_Newtype_:
		return newtype_(b)
	}
	panic(fmt.Sprintf("code gen error unhandled branch '%#v", in))
}

func (obj *DeclType) UnmarshalJSON(b []byte) error {
	key, val, err := SplitKV(b)
	if err != nil {
		return err
	}
	switch key {
	case "struct_":
		struct_ := Struct{}
		err = json.Unmarshal(val, &struct_)
		if err != nil {
			return err
		}
		obj.Branch = DeclTypeBranch_Struct_(struct_)
	case "union_":
		union_ := Union{}
		err = json.Unmarshal(val, &union_)
		if err != nil {
			return err
		}
		obj.Branch = DeclTypeBranch_Union_(union_)
	case "type_":
		type_ := TypeDef{}
		err = json.Unmarshal(val, &type_)
		if err != nil {
			return err
		}
		obj.Branch = DeclTypeBranch_Type_(type_)
	case "newtype_":
		newtype_ := NewType{}
		err = json.Unmarshal(val, &newtype_)
		if err != nil {
			return err
		}
		obj.Branch = DeclTypeBranch_Newtype_(newtype_)
	}
	return nil
}

func (obj DeclType) MarshalJSON() ([]byte, error) {
	discriminator, err := HandleDeclType(
		obj.Branch,
		func(struct_ DeclTypeBranch_Struct_) (string, error) {
			return "struct_", nil
		},
		func(union_ DeclTypeBranch_Union_) (string, error) {
			return "union_", nil
		},
		func(type_ DeclTypeBranch_Type_) (string, error) {
			return "type_", nil
		},
		func(newtype_ DeclTypeBranch_Newtype_) (string, error) {
			return "newtype_", nil
		},
	)
	if err != nil {
		return nil, err
	}
	if ba, err := json.Marshal(&obj.Branch); err != nil {
		return nil, err
	} else {
		discLen := len(discriminator)
		ba2 := make([]byte, len(ba)+len(`{"": }`)+discLen)
		copy(ba2, []byte(`{"`))
		copy(ba2[len(`{"`):], []byte(discriminator))
		copy(ba2[len(`{"`)+discLen:], []byte(`": `))
		copy(ba2[len(`{"`)+discLen+len(`": `):], ba)
		ba2[len(ba2)-1] = '}'
		return ba2, nil
	}
}

type Decl struct {
	Name        string   `json:"name"`
	Version     string   `json:"version,omitempty"`
	Type        DeclType `json:"type_"`
	Annotations `json:"annotations"`
}

type ScopedDecl struct {
	ScopedName ScopedName `json:"scopedName"`
	Decl       Decl       `json:"decl"`
}

// type DeclVersions = Vector<Decl>;

type Import struct {
	ModuleName *string     `json:"moduleName,omitempty"`
	ScopedName *ScopedName `json:"scopedName,omitempty"`
}

type Module struct {
	Name        string          `json:"name"`
	Imports     []Import        `json:"imports"`
	Decls       map[string]Decl `json:"decls"`
	Annotations `json:"annotations"`
}

type AnnotateAble interface {
	AddAnnotation(an Annotation)
}
type ImportableAble interface {
	AddImport(Import)
}
type Setable interface {
	Set(val interface{})
}

func (ans *Annotations) AddAnnotation(an Annotation) {
	*ans = append(*ans, an)
}
func (mo *Module) AddImport(im Import) {
	mo.Imports = append(mo.Imports, im)
}
func (an *Annotation) Set(val interface{}) {
	an.Val = val
}

func (m Module) String() string { return m.Name }
func (m Decl) String() string   { return m.Name }
func (i Import) String() string {
	if i.ModuleName != nil {
		return *i.ModuleName
	}
	return i.ScopedName.ModuleName + ":" + i.ScopedName.Name
}
