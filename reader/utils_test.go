package reader

import "testing"

func TestIsWhitespace(t *testing.T) {
	positiveCases := []rune{' ', '\n', '\r', '\t'}
	negativeCases := []rune{'j', '0', 'E'}

	for _, c := range positiveCases {
		if !isWhitespace(c) {
			t.Errorf("%c should be whitespace, but isn't", c)
		}
	}

	for _, c := range negativeCases {
		if isWhitespace(c) {
			t.Errorf("%d should not be whitespace, but is", c)
		}
	}
}
