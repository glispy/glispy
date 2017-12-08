package core

import (
	"fmt"

	gmath "github.com/itsmontoya/glisp/stdlib/math"
	gstrings "github.com/itsmontoya/glisp/stdlib/strings"

	"github.com/itsmontoya/glisp/common"
	"github.com/itsmontoya/glisp/scope"
	"github.com/itsmontoya/glisp/types"
	"github.com/itsmontoya/glisp/utils"
	"github.com/missionMeteora/journaler"
)

// Println will print a line to stdout
func Println(sc scope.Scope, args types.List) (exp types.Expression, err error) {
	journaler.Debug("Println called")
	vals := make([]interface{}, len(args))
	for i, v := range args {
		vals[i] = v
	}

	journaler.Notification("Glisp: %v", vals...)
	return
}

// Define will define a value
func Define(sc scope.Scope, args types.List) (_ types.Expression, err error) {
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

	if exp, err = utils.Eval(sc, exp); err != nil {
		return
	}

	sc.Put(sym, exp)
	return
}

// LessThan will return if the first argument is less than the second
func LessThan(sc scope.Scope, args types.List) (out types.Expression, err error) {
	if len(args) != 2 {
		err = common.ErrInvalidArgs
		return
	}

	if out, err = utils.Eval(sc, args[0]); err != nil {
		return
	}

	switch val := out.(type) {
	case types.Number:
		return gmath.LessThan(sc, val, args[1])

	case types.String:
		return gstrings.LessThan(sc, val, args[1])

	default:
		err = common.ErrExpectedAtom
		return
	}
}

// GreaterThan will return if the first argument is greater than the second
func GreaterThan(sc scope.Scope, args types.List) (out types.Expression, err error) {
	if len(args) != 2 {
		err = common.ErrInvalidArgs
		return
	}

	if out, err = utils.Eval(sc, args[0]); err != nil {
		return
	}

	switch val := out.(type) {
	case types.Number:
		return gmath.GreaterThan(sc, val, args[1])

	case types.String:
		return gstrings.GreaterThan(sc, val, args[1])

	default:
		err = common.ErrExpectedAtom
		return
	}
}

// Defun defines a function
func Defun(sc scope.Scope, args types.List) (_ types.Expression, err error) {
	var (
		sym   types.Symbol
		fargs types.List
		exp   types.Expression
		ok    bool
	)

	journaler.Debug("Defun! %v", args)

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

	fn := func(fargs types.List, exp types.Expression) utils.Func {
		return func(sc scope.Scope, args types.List) (out types.Expression, err error) {
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

			fmt.Print("\n\n\n\n\n\n\n\n\n")
			journaler.Debug("Being called! %v", fsc)
			if exp, err = utils.Eval(fsc, exp); err != nil {
				return
			}

			return
		}
	}(fargs, exp)

	sc.Put(sym, fn)
	return
}

// Add will add a list
func Add(sc scope.Scope, args types.List) (exp types.Expression, err error) {
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
func Begin(sc scope.Scope, args types.List) (_ types.Expression, err error) {
	for _, arg := range args {
		if _, err = utils.Eval(sc, arg); err != nil {
			return
		}
	}

	return
}
