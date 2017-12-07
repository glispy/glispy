package glisp

import (
	"strconv"

	"github.com/itsmontoya/glisp/tokens"
	"github.com/missionMeteora/toolkit/errors"
)

// ErrInvalidNumber is returned when a number is invalid
const ErrInvalidNumber = errors.Error("invalid number")

// NewNumber will return a new number
func NewNumber(t tokens.Token) (n Number, err error) {
	var float float64
	if float, err = strconv.ParseFloat(string(t), 32); err != nil {
		return
	}

	n = Number(float)
	return
}

// Number represents a number type
type Number float32
