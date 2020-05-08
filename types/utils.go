package types

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
