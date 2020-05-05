package glispy

import (
	"fmt"
	"io/ioutil"

	"github.com/glispy/glispy/eval"
	"github.com/glispy/glispy/scope"
	"github.com/glispy/glispy/stdlib/core"
	gmath "github.com/glispy/glispy/stdlib/math"
	"github.com/glispy/glispy/tokens"
	"github.com/glispy/glispy/types"
)

// New will return a new instance of Glispy
func New() (g Glispy) {
	s := scope.NewRoot()
	setFunc(s, "println", core.Println)
	setFunc(s, "+", core.Add)
	setFunc(s, "*", gmath.Multiply)
	setFunc(s, "define", core.Define)
	setFunc(s, "defun", core.Defun)
	setFunc(s, "begin", core.Begin)
	setFunc(s, ">", core.GreaterThan)
	setFunc(s, "<", core.LessThan)
	setFunc(s, "make-hash-map", core.MakeHashMap)
	setFunc(s, "get-value", core.GetValue)
	setFunc(s, "set-value", core.SetValue)
	setFunc(s, "remove-value", core.RemoveValue)
	return NewWithScope(s)
}

// NewWithScope will return a new instance of Glispy with a provided scope
func NewWithScope(s types.Scope) (g Glispy) {
	g.sc = s
	return
}

// Glispy is a lisp worker
type Glispy struct {
	sc types.Scope
}

// Eval will evaluate an Expression
func (g *Glispy) Eval(e types.Expression) (out types.Expression, err error) {
	return eval.Eval(g.sc, e)
}

// EvalTokens will evaluate tokens as an Expression
func (g *Glispy) EvalTokens(ts *tokens.Tokens) (out types.Expression, err error) {
	var e types.Expression
	if e, err = types.NewExpression(ts); err != nil {
		return
	}

	return g.Eval(e)
}

// EvalString will evaluate a string as an Expression
func (g *Glispy) EvalString(str string) (out types.Expression, err error) {
	ts := tokens.NewTokens(str)
	return g.EvalTokens(&ts)
}

// EvalFile will evaluate a file as an Expression
func (g *Glispy) EvalFile(filename string) (out types.Expression, err error) {
	var bs []byte
	if bs, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	return g.EvalString(string(bs))
}

// SetFunc allows Go funcs to be set as lisp funcs
func (g *Glispy) SetFunc(key string, fn types.Function) {
	setFunc(g.sc, key, fn)
}

// CallFunc will call a func within the global scope
func (g *Glispy) CallFunc(key string, args ...types.Expression) (out types.Expression, err error) {
	var exp types.Expression
	if exp, err = g.sc.Get(types.Symbol(key)); err != nil {
		return
	}

	var (
		fn types.Function
		ok bool
	)

	if fn, ok = exp.(types.Function); !ok {
		err = fmt.Errorf("invalid type, cannot assert %T as function", exp)
		return
	}

	return fn(g.sc, types.List(args))
}
