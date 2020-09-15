package common

import "github.com/hatchify/errors"

const (
	// ErrUnexpectedEOF is returned when an end is encountered before it was expected
	ErrUnexpectedEOF = errors.Error("unexpected end of file")
	// ErrUnexpectedCloseParens is returned when an closing paren is encountered before it was expected
	ErrUnexpectedCloseParens = errors.Error("unexpected close parens")
	// ErrKeyNotFound is returned when a key has not been found
	ErrKeyNotFound = errors.Error("key not found")
	// ErrExpectedSymbol is returned when a symbol is expected
	ErrExpectedSymbol = errors.Error("symbol expected")
	// ErrExpectedList is returned when a list is expected
	ErrExpectedList = errors.Error("list expected")
	// ErrExpectedFn is returned when a function is expected
	ErrExpectedFn = errors.Error("function expected")
	// ErrExpectedNumber is returned when a number is expected
	ErrExpectedNumber = errors.Error("expected number")
	// ErrExpectedString is returned when a string is expected
	ErrExpectedString = errors.Error("expected string")
	// ErrCannotAdd is returned when the provided type cannot be added
	ErrCannotAdd = errors.Error("cannot add the provided type")
	// ErrInvalidArgs is returned when there are the invalid number of arguments
	ErrInvalidArgs = errors.Error("invalid arguments")
	// ErrExpectedExpression is returned when an expression is expected
	ErrExpectedExpression = errors.Error("expected expression")
	// ErrExpectedAtom is returned when an atom is expected
	ErrExpectedAtom = errors.Error("expected atom")
)
