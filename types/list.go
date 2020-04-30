package types

import (
	"github.com/itsmontoya/glisp/common"
	"github.com/itsmontoya/glisp/tokens"
)

// NewList will return a new list
func NewList(ts *tokens.Tokens) (l List, err error) {
	var (
		token tokens.Token
		ok    bool
	)

	for {
		if token, ok = ts.Shift(); !ok {
			err = common.ErrUnexpectedEOF
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

// List represents a list of Atom's
type List []Atom

// GetSymbol will get a list item (by index) as a symbol
func (l List) GetSymbol(index int) (out Symbol, ok bool) {
	if len(l) <= index {
		return
	}

	val := l[index]
	out, ok = val.(Symbol)
	return
}

// GetString will get a list item (by index) as a string
func (l List) GetString(index int) (out String, ok bool) {
	if len(l) <= index {
		return
	}

	val := l[index]
	out, ok = val.(String)
	return
}

// GetNumber will get a list item (by index) as a number
func (l List) GetNumber(index int) (out Number, ok bool) {
	if len(l) <= index {
		return
	}

	val := l[index]
	out, ok = val.(Number)
	return
}

// GetFunction will get a list item (by index) as a function
func (l List) GetFunction(index int) (out Function, ok bool) {
	if len(l) <= index {
		return
	}

	val := l[index]
	out, ok = val.(Function)
	return
}

// GetAtom will get a list item (by index) as an atom
func (l List) GetAtom(index int) (out Atom, ok bool) {
	if len(l) <= index {
		return
	}

	val := l[index]
	out, ok = val.(Atom)
	return
}
