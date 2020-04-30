package glisp

import "github.com/itsmontoya/glisp/types"

func setFunc(s types.Scope, key string, fn types.Function) {
	s.Put(types.Symbol(key), fn)
}
