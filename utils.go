package glisp

import (
	"github.com/itsmontoya/glisp/tokens"
	"github.com/missionMeteora/journaler"
)

func toExpression(ts *tokens.Tokens, token tokens.Token) (e Expression, err error) {
	switch token {
	case "(":
		return NewList(ts)
	case ")":
		err = ErrUnexpectedCloseParens
		return

	default:
		return NewAtom(token)
	}
}

func println(args List) (exp Expression, err error) {
	vals := make([]interface{}, len(args))
	for i, v := range args {
		vals[i] = v
	}

	journaler.Notification("Glisp: %v", vals...)
	return
}

// Fn is the function type
type Fn func(args List) (Expression, error)
