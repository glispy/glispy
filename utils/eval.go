package utils

import (
	"github.com/itsmontoya/glisp/common"
	"github.com/itsmontoya/glisp/scope"
	"github.com/itsmontoya/glisp/types"
)

const ifSymbol = types.Symbol("if")

// Eval will evaluate an Expression
func Eval(sc scope.Scope, e types.Expression) (out types.Expression, err error) {
	switch val := e.(type) {
	case types.Number:
		out = val
	case types.String:
		out = val
	case types.Symbol:
		return handleSymbol(sc, val)
	case types.List:
		return handleList(sc, val)
	}

	return
}

func handleSymbol(sc scope.Scope, s types.Symbol) (out types.Expression, err error) {
	var ok bool
	if out, ok = sc.Get(s); !ok {
		err = common.ErrKeyNotFound
	}

	return
}

func handleList(sc scope.Scope, l types.List) (out types.Expression, err error) {
	tkn := l[0]
	switch tkn {
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

func handleFn(sc scope.Scope, l types.List) (out types.Expression, err error) {
	var (
		sym  types.Symbol
		ref  types.Expression
		fn   Func
		args types.List
		ok   bool
	)

	if sym, ok = l[0].(types.Symbol); !ok {
		err = common.ErrExpectedSymbol
		return
	}

	if ref, ok = sc.Get(sym); !ok {
		err = common.ErrKeyNotFound
		return
	}

	if fn, ok = ref.(Func); !ok {
		err = common.ErrExpectedFn
		return
	}

	args = l[1:]
	return fn(sc, args)
}
