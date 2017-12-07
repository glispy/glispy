package glisp

import (
	"math"

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
	// ErrExpectedSymbol is returned when a symbol is expected
	ErrExpectedSymbol = errors.Error("symbol expected")
	// ErrExpectedList is returned when a list is expected
	ErrExpectedList = errors.Error("list expected")
	// ErrExpectedFn is returned when a function is expected
	ErrExpectedFn = errors.Error("function expected")
	// ErrExpectedNumber is returned when a number is expected
	ErrExpectedNumber = errors.Error("expected number")
	// ErrExpectedString is returned when a string is expected
	ErrExpectedString = errors.Error("expected string")
	// ErrCannotAdd is returned when the provided type cannot be added
	ErrCannotAdd = errors.Error("cannot add the provided type")
	// ErrInvalidArgs is returned when there are the invalid number of arguments
	ErrInvalidArgs = errors.Error("invalid arguments")
	// ErrExpectedExpression is returned when an expression is expected
	ErrExpectedExpression = errors.Error("expected expression")
	// ErrExpectedAtom is returned when an atom is expected
	ErrExpectedAtom = errors.Error("expected atom")
)

const (
	ifSymbol = Symbol("if")
)

// NewGlisp will return a new instance of Glisp
func NewGlisp() (g Glisp) {
	g.env = make(Dict)
	g.setEnvFn("println", g.println)
	g.setEnvFn("+", g.add)
	g.setEnvFn("*", g.multiply)
	g.setEnvFn("define", g.define)
	g.setEnvFn("defun", g.defun)
	g.setEnvFn("begin", g.begin)
	g.setEnvFn(">", g.greaterThan)
	g.setEnvFn("<", g.lessThan)
	g.env["greeting"] = "Hello world"
	g.env["pi"] = Number(math.Pi)
	return
}

// Glisp is a lisp worker
type Glisp struct {
	env Dict
}

func (g *Glisp) setEnvFn(key string, fn Fn) {
	g.env[Symbol(key)] = fn
}

func (g *Glisp) println(args List) (exp Expression, err error) {
	vals := make([]interface{}, len(args))
	for i, v := range args {
		if vals[i], err = g.Eval(v); err != nil {
			return
		}
	}

	journaler.Notification("Glisp: %v", vals)
	return
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
	)

	for _, exp := range args {
		if str, err = g.getString(exp); err != nil {
			return
		}

		val += str
	}

	out = val
	return
}

func (g *Glisp) getNumber(exp Expression) (n Number, err error) {
	switch val := exp.(type) {
	case Number:
		n = val
	case Symbol:
		if exp, err = g.Eval(val); err != nil {
			return
		}

		return g.getNumber(exp)

	case List:
		if exp, err = g.Eval(val); err != nil {
			return
		}

		return g.getNumber(exp)

	default:
		journaler.Debug("Uhh: %v", exp)
		err = ErrExpectedNumber
	}

	return
}

func (g *Glisp) getString(exp Expression) (s String, err error) {
	switch val := exp.(type) {
	case String:
		s = val

	case Symbol:
		if exp, err = g.Eval(val); err != nil {
			return
		}

		return g.getString(exp)

	case List:
		if exp, err = g.Eval(val); err != nil {
			return
		}

		return g.getString(exp)

	default:
		journaler.Debug("Uhh: %v", exp)
		err = ErrExpectedString
	}

	return
}

func (g *Glisp) multiply(args List) (out Expression, err error) {
	var (
		n   Number
		num Number
	)

	for i, exp := range args {
		if num, err = g.getNumber(exp); err != nil {
			return
		}

		if i == 0 {
			n = num
		} else {
			n *= num
		}
	}

	out = n
	return
}

// define never returns a value
func (g *Glisp) define(args List) (_ Expression, err error) {
	var (
		sym Symbol
		exp Expression
		ok  bool
	)

	if len(args) < 2 {
		err = ErrInvalidArgs
		return
	}

	if sym, ok = args[0].(Symbol); !ok {
		err = ErrExpectedSymbol
		return
	}

	if exp, ok = args[1].(Expression); !ok {
		err = ErrExpectedExpression
		return
	}

	if exp, err = g.Eval(exp); err != nil {
		return
	}

	g.env[sym] = exp
	return
}

// defun never returns a value
func (g *Glisp) defun(args List) (_ Expression, err error) {
	var (
		sym   Symbol
		fargs List
		exp   Expression
		ok    bool
	)

	journaler.Debug("Defun! %v", args)

	if len(args) < 3 {
		err = ErrInvalidArgs
		return
	}

	if sym, ok = args[0].(Symbol); !ok {
		err = ErrExpectedSymbol
		return
	}

	if fargs, ok = args[1].(List); !ok {
		err = ErrExpectedList
		return
	}

	if exp, ok = args[2].(Expression); !ok {
		err = ErrExpectedExpression
		return
	}

	g.env[sym] = func(fargs List, exp Expression) Fn {
		return func(args List) (out Expression, err error) {
			journaler.Debug("Being called!")
			if exp, err = g.Eval(exp); err != nil {
				return
			}

			return
		}
	}(fargs, exp)

	journaler.Debug("YAS? %v", g.env[sym])
	return
}

func (g *Glisp) begin(args List) (_ Expression, err error) {
	for _, arg := range args {
		if _, err = g.Eval(arg); err != nil {
			return
		}
	}

	return
}

func (g *Glisp) lessThan(args List) (out Expression, err error) {
	if len(args) != 2 {
		err = ErrInvalidArgs
		return
	}

	if out, err = g.Eval(args[0]); err != nil {
		return
	}

	switch val := out.(type) {
	case Number:
		return g.lessThanNumber(val, args[1])

	case String:
		return g.lessThanString(val, args[1])

	default:
		err = ErrExpectedAtom
		return
	}
}

func (g *Glisp) greaterThan(args List) (out Expression, err error) {
	if len(args) != 2 {
		err = ErrInvalidArgs
		return
	}

	if out, err = g.Eval(args[0]); err != nil {
		return
	}

	switch val := out.(type) {
	case Number:
		return g.greaterThanNumber(val, args[1])

	case String:
		return g.greaterThanString(val, args[1])

	default:
		err = ErrExpectedAtom
		return
	}
}

func (g *Glisp) lessThanNumber(an Number, b Atom) (out Expression, err error) {
	var bn Number
	if bn, err = g.getNumber(b); err != nil {
		return
	}

	if an < bn {
		out = "true"
	}

	return
}

func (g *Glisp) greaterThanNumber(an Number, b Atom) (out Expression, err error) {
	var bn Number
	if bn, err = g.getNumber(b); err != nil {
		return
	}

	if an > bn {
		out = "true"
	}

	return
}

func (g *Glisp) lessThanString(as String, b Atom) (out Expression, err error) {
	var bs String
	if bs, err = g.getString(b); err != nil {
		return
	}

	if as < bs {
		out = "true"
	}

	return
}

func (g *Glisp) greaterThanString(as String, b Atom) (out Expression, err error) {
	var bs String
	if bs, err = g.getString(b); err != nil {
		return
	}

	if as > bs {
		out = "true"
	}

	return
}

/*

   elif x[0] == 'if':               # conditional
       (_, test, conseq, alt) = x
       exp = (conseq if eval(test, env) else alt)
       return eval(exp, env)
   else:                            # procedure call
*/

// Eval will evaluate an Expression
func (g *Glisp) Eval(e Expression) (out Expression, err error) {
	switch val := e.(type) {
	case Number:
		out = val
	case String:
		out = val
	case Symbol:
		return g.handleSymbol(val)
	case List:
		return g.handleList(val)
	}

	return
}

func (g *Glisp) handleSymbol(s Symbol) (out Expression, err error) {
	var ok bool
	if out, ok = g.env[s]; !ok {
		journaler.Debug("Key not found: %v", s)
		err = ErrKeyNotFound
	}

	return
}

func (g *Glisp) handleList(l List) (out Expression, err error) {
	tkn := l[0]
	switch tkn {
	case ifSymbol:
		test := l[1]
		conseq := l[2]
		alt := l[3]
		if out, err = g.Eval(test); err != nil {
			return
		}

		if out != nil {
			return g.Eval(conseq)
		}

		return g.Eval(alt)

		// Define should be able to be set in the env..
	//case "define":
	/*
		(_, symbol, exp) = x
		env[symbol] = eval(exp, env)
	*/
	default:
		return g.handleFn(l)
	}
}

func (g *Glisp) handleFn(l List) (out Expression, err error) {
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

	if ref, ok = g.env[sym]; !ok {
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
type Dict map[Symbol]Expression
