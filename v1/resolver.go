package goadl

type Resolver interface {
	Resolve(ScopedName) func() interface{}
	Register(ScopedName, func() interface{})
}

type ResolverType map[ScopedName]func() interface{}

var RESOLVER ResolverType = make(map[ScopedName]func() interface{})

func (rt *ResolverType) Resolve(name ScopedName) func() interface{} {
	fact, has := (*rt)[name]
	if !has {
		return nil
	}
	return fact
}

func (rt *ResolverType) Register(name ScopedName, fact func() interface{}) {
	(*rt)[name] = fact
}
