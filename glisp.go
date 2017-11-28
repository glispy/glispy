package glisp

import (
	"github.com/missionMeteora/toolkit/errors"
)

const (
	// ErrUnexpectedEOF is returned when an end is encountered before it was expected
	ErrUnexpectedEOF = errors.Error("unexpected end of file")
	// ErrUnexpectedCloseParens is returned when an closing paren is encountered before it was expected
	ErrUnexpectedCloseParens = errors.Error("unexpected close parens")
	// ErrKeyNotFound is returned when a key has not been found
	ErrKeyNotFound = errors.Error("key not found")
	// ErrExpectedSymbol is returned when a symbol is expected
	ErrExpectedSymbol = errors.Error("symbol expected")
	// ErrExpectedFn is returned when a function is expected
	ErrExpectedFn = errors.Error("function expected")
	// ErrExpectedNumber is returned when a number is expected
	ErrExpectedNumber = errors.Error("expected number")
	// ErrExpectedString is returned when a string is expected
	ErrExpectedString = errors.Error("expected string")
	// ErrCannotAdd is returned when the provided type cannot be added
	ErrCannotAdd = errors.Error("cannot add the provided type")
)

// NewGlisp will return a new instance of Glisp
func NewGlisp() (g Glisp) {
	g.env = make(Dict)
	g.setEnvFn("println", println)
	g.setEnvFn("+", g.add)

	g.env["greeting"] = "Hello world"
	return
}

// Glisp is a lisp worker
type Glisp struct {
	env Dict
}

func (g *Glisp) setEnvFn(key string, fn Fn) {
	g.env[key] = fn
}

func (g *Glisp) add(args List) (exp Expression, err error) {
	switch args[0].(type) {
	case Number:
		return g.addNumbers(args)
	case String:
		return g.addStrings(args)

	default:
		err = ErrCannotAdd
		return
	}
}

func (g *Glisp) addNumbers(args List) (out Expression, err error) {
	var (
		n   Number
		num Number
		ok  bool
	)

	out = 0

	for _, exp := range args {
		switch val := exp.(type) {
		case Number:
			n += val
		case List:
			var exp Expression
			if exp, err = g.handleList(val); err != nil {
				return
			}

			if num, ok = exp.(Number); !ok {
				err = ErrExpectedNumber
				return
			}

			n += num

		default:
			err = ErrExpectedNumber
			return
		}
	}

	out = n
	return
}

func (g *Glisp) addStrings(args List) (out Expression, err error) {
	var (
		val String
		str String
		ok  bool
	)

	for _, exp := range args {
		if str, ok = exp.(String); !ok {
			out = ""
			err = ErrExpectedString
			return
		}

		val += str
	}

	out = val
	return
}

/*
def eval(x: Exp, env=global_env) -> Exp:
    "Evaluate an expression in an environment."
    if isinstance(x, Symbol):        # variable reference
        return env[x]
    elif not isinstance(x, Number):  # constant number
        return x
    elif x[0] == 'if':               # conditional
        (_, test, conseq, alt) = x
        exp = (conseq if eval(test, env) else alt)
        return eval(exp, env)
    elif x[0] == 'define':           # definition
        (_, symbol, exp) = x
        env[symbol] = eval(exp, env)
    else:                            # procedure call
        proc = eval(x[0], env)
        args = [eval(arg, env) for arg in x[1:]]
        return proc(*args)
*/

// Eval will evaluate an Expression
func (g *Glisp) Eval(e Expression) (out Expression, err error) {
	switch val := e.(type) {
	case Symbol:
		return g.handleSymbol(string(val))

	case Number:
		//		journaler.Debug("Number: %v", val)

	case List:
		return g.handleList(val)
	}

	return
}

func (g *Glisp) handleSymbol(key string) (out Expression, err error) {
	var ok bool
	if out, ok = g.env[key]; !ok {
		err = ErrKeyNotFound
	}

	return
}

func (g *Glisp) handleList(l List) (out Expression, err error) {
	var (
		sym  Symbol
		ref  Expression
		fn   Fn
		args List
		ok   bool
	)

	if sym, ok = l[0].(Symbol); !ok {
		err = ErrExpectedSymbol
		return
	}

	if ref, ok = g.env[string(sym)]; !ok {
		err = ErrKeyNotFound
		return
	}

	if fn, ok = ref.(Fn); !ok {
		err = ErrExpectedFn
		return
	}

	args = l[1:]
	return fn(args)
}

// Dict represents a dictionary type
type Dict map[string]Expression
