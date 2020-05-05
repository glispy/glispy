package glispy

import "github.com/glispy/glispy/types"

func setFunc(s types.Scope, key string, fn types.Function) {
	s.Put(types.Symbol(key), fn)
}
