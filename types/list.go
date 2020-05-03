package types

import (
	"fmt"

	"github.com/Hatch1fy/errors"
	"github.com/glispy/glispy/common"
	"github.com/glispy/glispy/tokens"
)

const (
	// ErrOutOfBounds is returned when a requested index is out of bounds
	ErrOutOfBounds = errors.Error("cannot access argument, out of bounds")
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
func (l List) GetSymbol(index int) (out Symbol, err error) {
	if len(l) <= index {
		err = ErrOutOfBounds
		return
	}

	val := l[index]

	var ok bool
	if out, ok = val.(Symbol); !ok {
		err = fmt.Errorf("invalid type, expected symbol and received %T", val)
		return
	}

	return
}

// GetString will get a list item (by index) as a string
func (l List) GetString(index int) (out String, err error) {
	if len(l) <= index {
		err = ErrOutOfBounds
		return
	}

	val := l[index]

	var ok bool
	if out, ok = val.(String); !ok {
		err = fmt.Errorf("invalid type, expected string and received %T", val)
		return
	}

	return
}

// GetNumber will get a list item (by index) as a number
func (l List) GetNumber(index int) (out Number, err error) {
	if len(l) <= index {
		err = ErrOutOfBounds
		return
	}

	val := l[index]

	var ok bool
	if out, ok = val.(Number); !ok {
		err = fmt.Errorf("invalid type, expected number and received %T", val)
		return
	}

	return
}

// GetFunction will get a list item (by index) as a function
func (l List) GetFunction(index int) (out Function, err error) {
	if len(l) <= index {
		err = ErrOutOfBounds
		return
	}

	val := l[index]

	var ok bool
	if out, ok = val.(Function); !ok {
		err = fmt.Errorf("invalid type, expected function and received %T", val)
		return
	}

	return
}

// GetAtom will get a list item (by index) as an atom
func (l List) GetAtom(index int) (out Atom, err error) {
	if len(l) <= index {
		err = ErrOutOfBounds
		return
	}

	val := l[index]

	var ok bool
	if out, ok = val.(Atom); !ok {
		err = fmt.Errorf("invalid type, expected atom and received %T", val)
		return
	}

	return
}
