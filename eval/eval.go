package eval

import (
	"fmt"

	"github.com/glispy/glispy/common"
	"github.com/glispy/glispy/types"
)

const ifSymbol = types.Symbol("if")

// Eval will evaluate an Expression
func Eval(sc types.Scope, e types.Expression) (out types.Expression, err error) {
	switch val := e.(type) {
	case types.Number:
		out = val
	case types.String:
		out = val
	case types.Symbol:
		return handleSymbol(sc, val)
	case types.List:
		return handleList(sc, val)
	// Account for any non-stdlib type
	case types.Atom:
		out = val
	}

	return
}

func handleSymbol(sc types.Scope, s types.Symbol) (out types.Expression, err error) {
	var ok bool
	if out, ok = sc.Get(s); !ok {
		err = fmt.Errorf("symbol of \"%s\" was not found", s)
		return
	}

	return
}

func tryHandleSymbol(sc types.Scope, a types.Atom) (out types.Expression, err error) {
	var (
		sym types.Symbol
		ok  bool
	)

	if sym, ok = a.(types.Symbol); !ok {
		err = common.ErrExpectedSymbol
		return
	}

	return handleSymbol(sc, sym)
}

func handleList(sc types.Scope, l types.List) (out types.Expression, err error) {
	var (
		list types.List
		ok   bool
	)

	if len(l) == 0 {
		return
	}

	if list, ok = l[0].(types.List); !ok {
		return processList(sc, l)
	}

	if out, err = handleList(sc, list); err != nil {
		return
	}

	if l = l[1:]; len(l) == 0 {
		return
	}

	return handleList(sc, l)
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
		if out, err = Eval(sc, test); err != nil {
			return
		}

		if out != nil {
			return Eval(sc, conseq)
		}

		return Eval(sc, alt)

	default:
		return handleFn(sc, l)
	}
}

func handleFn(sc types.Scope, l types.List) (out types.Expression, err error) {
	var (
		ref  types.Expression
		fn   types.Function
		args types.List
		ok   bool
	)

	switch {
	// We check to see if the symbol is define or defun. If either, we do not want to replace the values
	case isSpecialOperator(l[0]):
		// // TODO: Ensure this isn't needed
		// if l, err = replaceSymbols(sc, l, 200); err != nil {
		//	return
		// }

		if ref, err = tryHandleSymbol(sc, l[0]); err != nil {
			return
		}

	default:
		if l, err = replaceSymbols(sc, l, 0); err != nil {
			return
		}

		ref = l[0]
	}

	// TODO: Figure out how to stop lists from being processed after being quoted
	fmt.Printf("REF? %+v / %T / %v\n", ref, ref, l[1:])
	if fn, ok = ref.(types.Function); !ok {
		err = common.ErrExpectedFn
		return
	}

	args = l[1:]
	return fn(sc, args)
}

func isSpecialOperator(symbol types.Atom) (ok bool) {
	switch symbol {
	case types.Symbol("quote"):
	case types.Symbol("defun"):
	case types.Symbol("defmacro"):

	default:
		// The provided symbol is not a symbol referencing func, return false
		return false
	}

	return true
}

func replaceSymbols(sc types.Scope, l types.List, startAt int) (out types.List, err error) {
	out = make(types.List, 0, len(l))
	for i, atom := range l {
		if startAt == -1 || i < startAt {
			out = append(out, atom)
			continue
		}

		var exp types.Expression
		switch n := atom.(type) {
		case types.Symbol:
			if exp, err = handleSymbol(sc, n); err != nil {
				return
			}

		case types.List:
			if exp, err = handleList(sc, n); err != nil {
				return
			}

		default:
			exp = atom
		}

		out = append(out, exp)
	}

	return
}

/*
Delete this if everything works


func replaceSymbols(sc types.Scope, l types.List, startAt int) (out types.List, err error) {
	var ok bool
	out = make(types.List, 0, len(l))
	for i, atom := range l {
		if i < startAt {
			out = append(out, atom)
			continue
		}

		var exp types.Expression
		switch n:= atom.(type) {
		case types.Symbol:
			if exp, err = handleSymbol(sc, n); err != nil {
				return
			}
		case types.Function:
			n(sc, )
		default:
			exp = atom
		}


		out = append(out, exp)
	}

	return
}

*/
