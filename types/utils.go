package types

import (
	"github.com/glispyy/glispyy/common"
	"github.com/glispyy/glispyy/tokens"
)

func toExpression(ts *tokens.Tokens, token tokens.Token) (e Expression, err error) {
	switch token {
	case "(":
		return NewList(ts)
	case ")":
		err = common.ErrUnexpectedCloseParens
		return

	default:
		return NewAtom(token)
	}
}

// ToAtom will convert a value to the atom representation
func ToAtom(val interface{}) (a Atom) {
	switch n := val.(type) {
	case string:
		return String(n)
	case float32:
		return Number(n)
	}

	return Atom(val)
}
