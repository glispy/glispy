package expand

import (
	"github.com/glispy/glispy/common"
	"github.com/glispy/glispy/scope"
	"github.com/glispy/glispy/types"
)

func expandFunc(sc types.Scope, l types.List) (out types.Expression, err error) {
	var (
		sym  types.Symbol
		ref  types.Expression
		list types.List

		fargs types.List
		args  types.List
		body  types.List

		ok bool
	)

	// Set returning expression as the provided list, just in case no match is found
	out = l

	if len(l) == 0 {
		return
	}

	// Attempt to assert the first list item as a symbol
	if sym, err = l.GetSymbol(0); err != nil {
		err = nil
		return
	}

	args = l[1:]

	if ref, ok = sc.Get(sym); !ok {
		return
	}

	if list, ok = ref.(types.List); !ok {
		return
	}

	if err = list.GetValues(&fargs, &body); err != nil {
		return
	}

	// Initialize a new function scope with the input root as the root
	fsc := scope.NewFunc(sc.Root())

	// Populate the scope with the argument values
	populateScope(fsc, fargs, args)

	// Replace values if possible
	return replaceValues(fsc, body)
}

func replaceValues(s types.Scope, body types.List) (out types.Expression, err error) {
	var (
		list types.List
		ok   bool
	)

	for _, exp := range body {
		switch n := exp.(type) {
		case types.Symbol:
			if exp, ok = s.Get(n); !ok {
				exp = n
			}

		case types.List:
			if exp, err = expandList(s, n); err != nil {
				return
			}

		default:
			continue
		}

		list = append(list, exp)
	}

	out = list
	return
}

func populateScope(s types.Scope, fargs, args types.List) (err error) {
	for i, arg := range args {
		var (
			sym types.Symbol
			ok  bool
		)

		if sym, ok = fargs[i].(types.Symbol); !ok {
			err = common.ErrExpectedSymbol
			return
		}

		s.Put(sym, arg)
	}

	return
}
