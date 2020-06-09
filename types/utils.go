package types

import "fmt"

// ToAtom will convert a value to the atom representation
func ToAtom(val interface{}) (a Atom) {
	switch n := val.(type) {
	case string:
		return String(n)
	case float32:
		return Number(n)
	}

	return Atom(val)
}

// ToNumber converts a value to a number
func ToNumber(a Atom) (n Number, err error) {
	switch val := a.(type) {
	case Number:
		n = val
	case int:
		n = Number(val)
	case int32:
		n = Number(val)
	case int64:
		n = Number(val)
	case uint:
		n = Number(val)
	case uint32:
		n = Number(val)
	case uint64:
		n = Number(val)
	case float32:
		n = Number(val)
	case float64:
		n = Number(val)

	default:
		err = fmt.Errorf("value of %v cannot be converted to a number", a)
	}

	return
}
