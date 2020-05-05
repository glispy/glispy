package types

import (
	"github.com/Hatch1fy/errors"
	"github.com/glispyy/glispyy/tokens"
)

// ErrInvalidString is returned when a string is invalid
const ErrInvalidString = errors.Error("invalid string")

// NewString will return a new string
func NewString(t tokens.Token) (s String, err error) {
	if t[0] != '"' {
		err = ErrInvalidString
		return
	}

	if t[len(t)-1] != '"' {
		err = ErrInvalidString
		return
	}

	s = String(t[1 : len(t)-1])
	return
}

// String represents a string type
type String string
