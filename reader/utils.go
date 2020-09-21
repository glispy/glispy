package reader

func isProtected(char rune) (ok bool) {
	switch char {
	case '(':
	case ')':

	default:
		return false
	}

	return true
}
