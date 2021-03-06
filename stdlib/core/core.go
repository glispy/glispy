package core

import (
	"fmt"

	gmath "github.com/glispy/glispy/stdlib/math"
	gstrings "github.com/glispy/glispy/stdlib/strings"

	"github.com/glispy/glispy/common"
	"github.com/glispy/glispy/eval"
	"github.com/glispy/glispy/scope"
	"github.com/glispy/glispy/types"
	"github.com/hatchify/scribe"
)

var out = scribe.New("Glispy")

// Quote will quote a value
func Quote(sc types.Scope, args types.List) (out types.Expression, err error) {
	if len(args) < 1 {
		err = common.ErrInvalidArgs
		return
	}

	out = args[0]
	return
}

// Println will print a line to stdout
func Println(sc types.Scope, args types.List) (_ types.Expression, err error) {
	out.Notificationf("%+v", args)
	return
}

// Define will define a value
func Define(sc types.Scope, args types.List) (_ types.Expression, err error) {
	var (
		sym types.Symbol
		exp types.Expression
		ok  bool
	)

	if len(args) < 2 {
		err = common.ErrInvalidArgs
		return
	}

	if sym, ok = args[0].(types.Symbol); !ok {
		err = common.ErrExpectedSymbol
		return
	}

	if exp, ok = args[1].(types.Expression); !ok {
		err = common.ErrExpectedExpression
		return
	}

	if exp, err = eval.Eval(sc, exp); err != nil {
		return
	}

	sc.Put(sym, exp)
	return
}

// LessThan will return if the first argument is less than the second
func LessThan(sc types.Scope, args types.List) (out types.Expression, err error) {
	if len(args) != 2 {
		err = common.ErrInvalidArgs
		return
	}

	var firstValue types.Atom
	if firstValue, err = args.GetAtom(0); err != nil {
		return
	}

	switch val := firstValue.(type) {
	case types.Number:
		return gmath.LessThan(sc, val, args[1])

	case types.String:
		return gstrings.LessThan(sc, val, args[1])

	default:
		err = fmt.Errorf("invalid value, type of %T is not supported", firstValue)
		return
	}
}

// GreaterThan will return if the first argument is greater than the second
func GreaterThan(sc types.Scope, args types.List) (out types.Expression, err error) {
	if len(args) != 2 {
		err = common.ErrInvalidArgs
		return
	}

	var firstValue types.Atom
	if firstValue, err = args.GetAtom(0); err != nil {
		return
	}

	switch val := firstValue.(type) {
	case types.Number:
		return gmath.GreaterThan(sc, val, args[1])

	case types.String:
		return gstrings.GreaterThan(sc, val, args[1])

	default:
		err = fmt.Errorf("invalid value, type of %T is not supported", firstValue)
		return
	}
}

// Defun defines a function
func Defun(sc types.Scope, args types.List) (_ types.Expression, err error) {
	var (
		sym   types.Symbol
		fargs types.List
		exp   types.Expression
		ok    bool
	)

	if len(args) < 3 {
		err = common.ErrInvalidArgs
		return
	}

	if sym, ok = args[0].(types.Symbol); !ok {
		err = common.ErrExpectedSymbol
		return
	}

	if fargs, ok = args[1].(types.List); !ok {
		err = common.ErrExpectedList
		return
	}

	if exp, ok = args[2].(types.Expression); !ok {
		err = common.ErrExpectedExpression
		return
	}

	fn := NewFunc(fargs, exp)
	sc.Put(sym, fn)
	return
}

// Add will add a list
func Add(sc types.Scope, args types.List) (exp types.Expression, err error) {
	switch args[0].(type) {
	case types.Number:
		return gmath.Add(sc, args)
	case types.String:
		return gstrings.Add(sc, args)

	default:
		err = common.ErrCannotAdd
		return
	}
}

// Begin will begin an expression
func Begin(sc types.Scope, args types.List) (_ types.Expression, err error) {
	for _, arg := range args {
		if _, err = eval.Eval(sc, arg); err != nil {
			return
		}
	}

	return
}

// NewFunc will return a new func
func NewFunc(fargs types.List, exp types.Expression) (fn types.Function) {
	return func(sc types.Scope, args types.List) (out types.Expression, err error) {
		if len(args) != len(fargs) {
			err = common.ErrInvalidArgs
			return
		}

		fsc := scope.NewFunc(sc)
		for i, arg := range args {
			var (
				sym types.Symbol
				ok  bool
			)

			if sym, ok = fargs[i].(types.Symbol); !ok {
				err = common.ErrExpectedSymbol
				return
			}

			fsc.Put(sym, arg)
		}

		if out, err = eval.Eval(fsc, exp); err != nil {
			return
		}

		return
	}
}
