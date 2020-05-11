package scope

import "github.com/glispy/glispy/types"

// NewRoot will return a new root scope
func NewRoot() *Root {
	var r Root
	r.d = make(types.Dict)
	return &r
}

// Root scope
type Root struct {
	d types.Dict
}

// Get will get a value
func (r *Root) Get(key types.Symbol) (out types.Expression, ok bool) {
	out, ok = r.d[key]
	return
}

// Put will set a value
func (r *Root) Put(key types.Symbol, exp types.Expression) {
	r.d[key] = exp
}

// PutRoot will set a value to root
func (r *Root) PutRoot(key types.Symbol, exp types.Expression) {
	r.d[key] = exp
}

// Root will return the root Scope
func (r *Root) Root() (root types.Scope) {
	return r
}

// Len will return the length of the local scope
func (r *Root) Len() int {
	return len(r.d)
}
