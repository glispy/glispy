package glisp

// List represents a list of Atom's
type List []Atom

// NewList will return a new list
func NewList(ts Tokens) (l List, err error) {
	var (
		tkn Token
		ok  bool
	)

	for {
		if tkn, ok = ts.Shift(); !ok {
			err = ErrUnexpectedEOF
			return
		}

		if tkn == ")" {
			return
		}

		var e Expression
		if e, err = NewExpression(ts); err != nil {
			return
		}

		l = append(l, e)
	}
}
