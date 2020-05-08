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

func isWhitespace(char rune) (ok bool) {
	switch char {
	case ' ':
	case '\t':
	case '\n':

	default:
		return false
	}

	return true
}
