package reader

import "unicode"

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
	return unicode.IsSpace(char)
}
