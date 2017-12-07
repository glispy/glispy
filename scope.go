package glisp

// Scope represents a scope layer
type Scope interface {
	Get(key string) (out Expression, err error)
	Put(key string, exp Expression)
}
