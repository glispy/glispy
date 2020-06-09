package expand

import (
	"github.com/glispy/glispy/types"
)

const (
	ifSymbol = types.Symbol("if")
)

// Expand will perform a macro expansion pass
func Expand(sc types.Scope, e types.Expression) (out types.Expression, err error) {
	if sc.Len() == 0 {
		out = e
		return
	}

	switch val := e.(type) {
	case types.List:
		return expandList(sc, val)
	// Account for any non-stdlib type
	case types.Atom:
		out = val
	}

	return
}
