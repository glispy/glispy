package strings

import (
	"github.com/itsmontoya/glisp/scope"
	"github.com/itsmontoya/glisp/types"
	"github.com/itsmontoya/glisp/utils"
)

// Add will add a series of strings
func Add(sc scope.Scope, args types.List) (out types.Expression, err error) {
	var (
		val types.String
		str types.String
	)

	for _, exp := range args {
		if str, err = utils.GetString(sc, exp); err != nil {
			return
		}

		val += str
	}

	out = val
	return
}

// LessThan will return if a is less than b
func LessThan(sc scope.Scope, as types.String, b types.Atom) (out types.Expression, err error) {
	var bs types.String
	if bs, err = utils.GetString(sc, b); err != nil {
		return
	}

	if as < bs {
		out = "true"
	}

	return
}

// GreaterThan will return if a is greater than b
func GreaterThan(sc scope.Scope, as types.String, b types.Atom) (out types.Expression, err error) {
	var bs types.String
	if bs, err = utils.GetString(sc, b); err != nil {
		return
	}

	if as > bs {
		out = "true"
	}

	return
}
