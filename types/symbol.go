package types

import (
	"github.com/hatchify/errors"
)

// ErrInvalidSymbol is returned when a symbol is invalid
const ErrInvalidSymbol = errors.Error("invalid symbol")

// Symbol represents a symbol
type Symbol string
