package glisp

import "github.com/itsmontoya/glisp/tokens"

// NewExpression will return a new expression from Tokens
func NewExpression(ts *tokens.Tokens) (e Expression, err error) {
	token, ok := ts.Shift()
	if !ok {
		err = ErrUnexpectedEOF
		return
	}

	return toExpression(ts, token)
}

// Expression represents an expression (Either an atom or a list)
type Expression interface{}
