package glisp

import (
	"regexp"

	"github.com/itsmontoya/glisp/tokens"
	"github.com/missionMeteora/toolkit/errors"
)

// ErrInvalidSymbol is returned when a symbol is invalid
const ErrInvalidSymbol = errors.Error("invalid symbol")

var symbolRegExp = regexp.MustCompile(`[^a-zA-Z_<>+*-]`)

// NewSymbol will return a new Symbol
func NewSymbol(t tokens.Token) (s Symbol, err error) {
	if symbolRegExp.Match([]byte(t)) {
		err = ErrInvalidSymbol
		return
	}

	s = Symbol(t)
	return
}

// Symbol represents a symbol
type Symbol string
