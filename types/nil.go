package types

import "github.com/hatchify/errors"

// ErrInvalidNil is returned when a nil is invalid
const ErrInvalidNil = errors.Error("invalid nil")

// Nil represents a nil atom value
type Nil struct{}
