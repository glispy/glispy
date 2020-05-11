package types

// Scope represents a scope type
type Scope interface {
	Get(key Symbol) (out Expression, ok bool)
	Put(key Symbol, in Expression)
	PutRoot(key Symbol, in Expression)
	Root() Scope
	Len() int
}
