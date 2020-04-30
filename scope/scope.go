package scope

import "github.com/itsmontoya/glisp/types"

// Scope represents a scope layer
type Scope interface {
	Get(key types.Symbol) (out types.Expression, ok bool)
	Put(key types.Symbol, exp types.Expression)
	PutRoot(key types.Symbol, exp types.Expression)
}
