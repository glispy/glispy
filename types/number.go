package types

import (
	"github.com/hatchify/errors"
)

// ErrInvalidNumber is returned when a number is invalid
const ErrInvalidNumber = errors.Error("invalid number")

// Number represents a number type
type Number float32
