package expand

import (
	"github.com/glispy/glispy/types"
)

func expandList(sc types.Scope, l types.List) (out types.Expression, err error) {
	if len(l) == 0 {
		out = l
		return
	}

	if _, ok := l[0].(types.List); !ok {
		return processList(sc, l)
	}

	var list types.List
	for _, listItem := range l {
		exp := listItem
		if value, ok := listItem.(types.List); ok {
			if exp, err = processList(sc, value); err != nil {
				return
			}
		}

		list = append(list, exp)
	}

	out = list
	return
}

func processList(sc types.Scope, l types.List) (out types.Expression, err error) {
	if len(l) == 0 {
		return
	}

	switch l[0] {
	case ifSymbol:
		test := l[1]
		conseq := l[2]
		alt := l[3]
		if out, err = Expand(sc, test); err != nil {
			return
		}

		if out != nil {
			return Expand(sc, conseq)
		}

		return Expand(sc, alt)

	default:
		return expandFunc(sc, l)
	}
}
