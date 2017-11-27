package glisp

// NewExpression will return a new expression from Tokens
func NewExpression(ts Tokens) (e Expression, err error) {
	token, ok := ts.Shift()
	if !ok {
		err = ErrUnexpectedEOF
		return
	}

	switch token {
	case "(":
		return NewList(ts)
	case ")":
		err = ErrUnexpectedCloseParens
		return

	default:
		return NewAtom(token)
	}
}

// Expression represents an expression (Either an atom or a list)
type Expression interface{}
