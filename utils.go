package glispy

import "github.com/glispyy/glispyy/types"

func setFunc(s types.Scope, key string, fn types.Function) {
	s.Put(types.Symbol(key), fn)
}
