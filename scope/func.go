package scope

import "github.com/glispy/glispy/types"

// NewFunc will return a new function scope
func NewFunc(parent types.Scope) *Func {
	var f Func
	f.p = parent
	f.d = make(types.Dict)
	return &f
}

// Func represents a function scope
type Func struct {
	// Parent scope
	p types.Scope
	// Local scope
	d types.Dict
}

// Get will get a value
func (f *Func) Get(key types.Symbol) (out types.Expression, ok bool) {
	if out, ok = f.d[key]; ok {
		return
	}

	if f.p == nil {
		return
	}

	return f.p.Get(key)
}

// Put will set a value
func (f *Func) Put(key types.Symbol, in types.Expression) {
	f.d[key] = in
}

// PutRoot will set a value to root
func (f *Func) PutRoot(key types.Symbol, in types.Expression) {
	if f.p == nil {
		f.Put(key, in)
		return
	}

	f.p.PutRoot(key, in)
	return
}

// Root will return the root Scope
func (f *Func) Root() (root types.Scope) {
	if f.p == nil {
		return f
	}

	return f.p.Root()
}

// Len will return the length of the local scope
func (f *Func) Len() int {
	return len(f.d)
}
