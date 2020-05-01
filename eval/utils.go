package eval

import (
	"github.com/itsmontoya/glisp/common"
	"github.com/itsmontoya/glisp/types"
)

// GetNumber will get a number from an expression
func GetNumber(sc types.Scope, exp types.Expression) (n types.Number, err error) {
	switch val := exp.(type) {
	case types.Number:
		n = val
	case types.Symbol:
		if exp, err = Eval(sc, val); err != nil {
			return
		}

		return GetNumber(sc, exp)

	case types.List:
		if exp, err = Eval(sc, val); err != nil {
			return
		}

		return GetNumber(sc, exp)

	default:
		err = common.ErrExpectedNumber
	}

	return
}

// GetString will get a string from an expression
func GetString(sc types.Scope, exp types.Expression) (s types.String, err error) {
	switch val := exp.(type) {
	case types.String:
		s = val

	case types.Symbol:
		if exp, err = Eval(sc, val); err != nil {
			return
		}

		return GetString(sc, exp)

	case types.List:
		if exp, err = Eval(sc, val); err != nil {
			return
		}

		return GetString(sc, exp)

	default:
		err = common.ErrExpectedString
	}

	return
}

// Func is the function type
//type Func func(sc types.Scope, args types.List) (types.Expression, error)
