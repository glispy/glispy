package types

import (
	"fmt"

	"github.com/Hatch1fy/errors"
)

const (
	// ErrOutOfBounds is returned when a requested index is out of bounds
	ErrOutOfBounds = errors.Error("cannot access argument, out of bounds")
)

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

// GetList will get a list item (by index) as a list
func (l List) GetList(index int) (out List, err error) {
	if len(l) <= index {
		err = ErrOutOfBounds
		return
	}

	val := l[index]

	var ok bool
	if out, ok = val.(List); !ok {
		err = fmt.Errorf("invalid type, expected list and received %T", val)
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

// GetValues will get a set of values from a list
func (l List) GetValues(values ...interface{}) (err error) {
	for i, value := range values {
		if err = l.GetValue(i, value); err != nil {
			return
		}
	}

	return
}

// GetValue will get a value from a list
func (l List) GetValue(i int, value interface{}) (err error) {
	switch n := value.(type) {
	case *Symbol:
		var sym Symbol
		if sym, err = l.GetSymbol(i); err != nil {
			err = fmt.Errorf("error getting value #%d: %v", i, err)
			return
		}

		*n = sym
	case *String:
		var str String
		if str, err = l.GetString(i); err != nil {
			err = fmt.Errorf("error getting value #%d: %v", i, err)
			return
		}

		*n = str
	case *Number:
		var num Number
		if num, err = l.GetNumber(i); err != nil {
			err = fmt.Errorf("error getting value #%d: %v", i, err)
			return
		}

		*n = num
	case *List:
		var list List
		if list, err = l.GetList(i); err != nil {
			err = fmt.Errorf("error getting value #%d: %v", i, err)
			return
		}

		*n = list
	case *Function:
		var fn Function
		if fn, err = l.GetFunction(i); err != nil {
			err = fmt.Errorf("error getting value #%d: %v", i, err)
			return
		}

		*n = fn

	case *Expression:
		var exp Atom
		if exp, err = l.GetAtom(i); err != nil {
			err = fmt.Errorf("error getting value #%d: %v", i, err)
			return
		}

		*n = exp

	case *Atom:
		var exp Atom
		if exp, err = l.GetAtom(i); err != nil {
			err = fmt.Errorf("error getting value #%d: %v", i, err)
			return
		}

		*n = exp

	default:
		err = fmt.Errorf("error getting value #%d: type of %T not supported", i, value)
		return
	}

	return
}
