package glisp

import (
	"github.com/missionMeteora/journaler"
	"github.com/missionMeteora/toolkit/errors"
)

const (
	// ErrUnexpectedEOF is returned when an end is encountered before it was expected
	ErrUnexpectedEOF = errors.Error("unexpected end of file")
	// ErrUnexpectedCloseParens is returned when an closing paren is encountered before it was expected
	ErrUnexpectedCloseParens = errors.Error("unexpected close parens")
	// ErrKeyNotFound is returned when a key has not been found
	ErrKeyNotFound = errors.Error("key not found")
)

// NewGlisp will return a new instance of Glisp
func NewGlisp() (g Glisp) {
	g.env = make(Dict)
	g.env["greeting"] = "Hello world"
	g.env["println"] = println
	return
}

// Glisp is a lisp worker
type Glisp struct {
	env Dict
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
	case String:
		var ok bool
		if out, ok = g.env[string(val)]; !ok {
			err = ErrKeyNotFound
		}

		return
	case Number:
		journaler.Debug("Number: %v", val)

	}

	return
}

// Dict represents a dictionary type
type Dict map[string]Expression
