package eval

/*
// Expand will perform a macro expansion pass
func Expand(sc types.Scope, e types.Expression) (out types.Expression, err error) {
	switch val := e.(type) {
	case types.Number:
		out = val
	case types.String:
		out = val
	case types.Symbol:
		return expandSymbol(sc, val)
	case types.List:
		return expandList(sc, val)
	// Account for any non-stdlib type
	case types.Atom:
		out = val
	}

	return
}

func expandSymbol(sc types.Scope, s types.Symbol, l types.List) (out types.Expression, err error) {
	var ok bool
	if out, ok = sc.Get(s); !ok {
		out = s
		return
	}

	var fn types.Function
	if fn, ok = out.(types.Function); !ok {
		err = fmt.Errorf("invalid type, expected types.Function and received %T", out)
		return
	}

	return fn(sc, l)
}

func expandList(sc types.Scope, l types.List) (out types.Expression, err error) {
	for i := 0; i < len(l); i++ {
		var exp types.Expression
		atom := l[i]
		switch n := atom.(type) {
		case types.Symbol:
			if exp, err = expandSymbol(sc, n, l[i+1:]); err != nil {
				err = fmt.Errorf("error expanding symbol \"%v\": %v", n, err)
				return
			}

			l = insertToSlice(l, exp, i)
		case types.List:
			if exp, err = expandList(sc, n); err != nil {
				err = fmt.Errorf("error expanding list \"%v\": %v", n, err)
				return
			}

		default:
			exp = atom
		}
	}

	out = l
	return
}

func insertToSlice(in types.List, exp types.Expression, index int) (out types.List) {
	// Append list head
	out = append(out, in[:index]...)

	// Determine if expression is a list
	if list, ok := exp.(types.List); ok {
		// Expression is a list, append list items individually
		out = append(out, list...)
	} else {
		// Expression is something else, append it as an single expression
		out = append(out, exp)
	}

	// Append list tail
	out = append(out, in[index:]...)
	return
}

/*
Delete this if everything works


func replaceSymbols(sc types.Scope, l types.List, startAt int) (out types.List, err error) {
	var ok bool
	out = make(types.List, 0, len(l))
	for i, atom := range l {
		if i < startAt {
			out = append(out, atom)
			continue
		}

		var exp types.Expression
		switch n:= atom.(type) {
		case types.Symbol:
			if exp, err = handleSymbol(sc, n); err != nil {
				return
			}
		case types.Function:
			n(sc, )
		default:
			exp = atom
		}


		out = append(out, exp)
	}

	return
}

*/
