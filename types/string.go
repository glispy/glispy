package types

import (
	"github.com/Hatch1fy/errors"
)

// ErrInvalidString is returned when a string is invalid
const ErrInvalidString = errors.Error("invalid string")

// String represents a string type
type String string
