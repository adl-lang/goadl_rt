// Code generated by goadlc v3 - DO NOT EDIT.
package haskell

import ()

type HaskellCustomType struct {
	_HaskellCustomType
}

type _HaskellCustomType struct {
	Haskellname         string             `json:"haskellname"`
	Haskellimports      []string           `json:"haskellimports"`
	Haskellextraexports []string           `json:"haskellextraexports"`
	InsertCode          []string           `json:"insertCode"`
	GenerateOrigADLType string             `json:"generateOrigADLType"`
	StructConstructor   string             `json:"structConstructor"`
	UnionConstructors   []UnionConstructor `json:"unionConstructors"`
}

func MakeAll_HaskellCustomType(
	haskellname string,
	haskellimports []string,
	haskellextraexports []string,
	insertcode []string,
	generateorigadltype string,
	structconstructor string,
	unionconstructors []UnionConstructor,
) HaskellCustomType {
	return HaskellCustomType{
		_HaskellCustomType{
			Haskellname:         haskellname,
			Haskellimports:      haskellimports,
			Haskellextraexports: haskellextraexports,
			InsertCode:          insertcode,
			GenerateOrigADLType: generateorigadltype,
			StructConstructor:   structconstructor,
			UnionConstructors:   unionconstructors,
		},
	}
}

func Make_HaskellCustomType(
	haskellname string,
	haskellimports []string,
	insertcode []string,
) HaskellCustomType {
	ret := HaskellCustomType{
		_HaskellCustomType{
			Haskellname:         haskellname,
			Haskellimports:      haskellimports,
			Haskellextraexports: ((*HaskellCustomType)(nil)).Default_haskellextraexports(),
			InsertCode:          insertcode,
			GenerateOrigADLType: ((*HaskellCustomType)(nil)).Default_generateOrigADLType(),
			StructConstructor:   ((*HaskellCustomType)(nil)).Default_structConstructor(),
			UnionConstructors:   ((*HaskellCustomType)(nil)).Default_unionConstructors(),
		},
	}
	return ret
}

func (*HaskellCustomType) Default_haskellextraexports() []string {
	return []string{}
}
func (*HaskellCustomType) Default_generateOrigADLType() string {
	return ""
}
func (*HaskellCustomType) Default_structConstructor() string {
	return ""
}
func (*HaskellCustomType) Default_unionConstructors() []UnionConstructor {
	return []UnionConstructor{}
}

type HaskellFieldPrefix = string

type UnionConstructor struct {
	_UnionConstructor
}

type _UnionConstructor struct {
	FieldName   string `json:"fieldName"`
	Constructor string `json:"constructor"`
}

func MakeAll_UnionConstructor(
	fieldname string,
	constructor string,
) UnionConstructor {
	return UnionConstructor{
		_UnionConstructor{
			FieldName:   fieldname,
			Constructor: constructor,
		},
	}
}

func Make_UnionConstructor(
	fieldname string,
	constructor string,
) UnionConstructor {
	ret := UnionConstructor{
		_UnionConstructor{
			FieldName:   fieldname,
			Constructor: constructor,
		},
	}
	return ret
}
