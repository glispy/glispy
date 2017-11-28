package glisp

import (
	"github.com/missionMeteora/toolkit/errors"
)

// ErrInvalidAtom is returned when an atom is invalid
const ErrInvalidAtom = errors.Error("atom must be a number or a string")

// NewAtom will return a new atom (Number, string, or symbol)
func NewAtom(t Token) (a Atom, err error) {
	if a, err = NewSymbol(t); err == nil {
		return
	}

	if a, err = NewNumber(t); err == nil {
		return
	}

	if a, err = NewString(t); err == nil {
		return
	}

	err = ErrInvalidAtom
	return
}

// Atom is an atom type
type Atom interface{}
