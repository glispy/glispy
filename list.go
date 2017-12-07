package glisp

import "github.com/itsmontoya/glisp/tokens"

// List represents a list of Atom's
type List []Atom

// NewList will return a new list
func NewList(ts *tokens.Tokens) (l List, err error) {
	var (
		token tokens.Token
		ok    bool
	)

	for {
		if token, ok = ts.Shift(); !ok {
			err = ErrUnexpectedEOF
			return
		}

		if token == ")" {
			return
		}

		var e Expression
		if e, err = toExpression(ts, token); err != nil {
			return
		}

		l = append(l, e)
	}
}
