package types

import (
	"github.com/itsmontoya/glisp/common"
	"github.com/itsmontoya/glisp/tokens"
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
