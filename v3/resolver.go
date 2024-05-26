package goadl

import (
	. "github.com/adl-lang/goadl_rt/v3/sys/adlast"
)

type BranchFactory interface {
	MakeNewBranch(key string) (any, error)
}

type CustomTypeHelper interface {
	BuildEncodeFunc(typeparamEnc ...EncoderFunc) EncoderFunc
	BuildDecodeFunc(typeparamDec ...DecodeFunc) DecodeFunc
}

type Resolver interface {
	Resolve(ScopedName) *ScopedDecl
	ResolveHelper(name ScopedName) (CustomTypeHelper, bool)
}

type Registry interface {
	Register(ScopedName, ScopedDecl)
}

type ResolverType struct {
	store   map[ScopedName]*ScopedDecl
	helpers map[ScopedName]CustomTypeHelper
}

var RESOLVER *ResolverType = &ResolverType{
	store:   make(map[ScopedName]*ScopedDecl),
	helpers: make(map[ScopedName]CustomTypeHelper),
}

func (rt *ResolverType) Resolve(name ScopedName) *ScopedDecl {
	sd, has := rt.store[name]
	if !has {
		return nil
	}
	return sd
}

func (rt *ResolverType) ResolveHelper(name ScopedName) (CustomTypeHelper, bool) {
	cth, has := rt.helpers[name]
	if !has {
		return cth, false
	}
	return cth, true
}

func (rt *ResolverType) Register(name ScopedName, sd ScopedDecl) {
	rt.store[name] = &sd
}

func (rt *ResolverType) RegisterHelper(name ScopedName, cth CustomTypeHelper) {
	rt.helpers[name] = cth
}
