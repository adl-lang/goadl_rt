package goadl

import (
	"reflect"

	. "github.com/adl-lang/goadl_rt/v3/sys/adlast"
)

type BranchFactory interface {
	MakeNewBranch(key string) (any, error)
}

type ScopedInfo struct {
	SD      ScopedDecl
	TypeMap map[string]reflect.Type
}

type Resolver interface {
	Resolve(ScopedName) *ScopedInfo
}

type Registry interface {
	Register(ScopedName, ScopedInfo)
}

type ResolverType struct {
	store map[ScopedName]*ScopedInfo
}

var RESOLVER *ResolverType = &ResolverType{
	store: make(map[ScopedName]*ScopedInfo),
}

func (rt *ResolverType) Resolve(name ScopedName) *ScopedInfo {
	sd, has := rt.store[name]
	if !has {
		return nil
	}
	return sd
}

func (rt *ResolverType) Register(name ScopedName, sd ScopedInfo) {
	rt.store[name] = &sd
}