package goadl

import (
	. "github.com/adl-lang/goadl_rt/v3/sys/adlast"
)

type BranchFactory interface {
	MakeNewBranch(key string) (any, error)
}

type Resolver interface {
	Resolve(ScopedName) *ScopedDecl
}

type Registry interface {
	Register(ScopedName, ScopedDecl)
}

type ResolverType struct {
	store map[ScopedName]*ScopedDecl
}

var RESOLVER *ResolverType = &ResolverType{
	store: make(map[ScopedName]*ScopedDecl),
}

func (rt *ResolverType) Resolve(name ScopedName) *ScopedDecl {
	sd, has := rt.store[name]
	if !has {
		return nil
	}
	return sd
}

func (rt *ResolverType) Register(name ScopedName, sd ScopedDecl) {
	rt.store[name] = &sd
}
