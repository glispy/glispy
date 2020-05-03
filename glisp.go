package glisp

import (
	"io/ioutil"

	"github.com/glispy/glispy/eval"
	"github.com/glispy/glispy/scope"
	"github.com/glispy/glispy/stdlib/core"
	gmath "github.com/glispy/glispy/stdlib/math"
	"github.com/glispy/glispy/tokens"
	"github.com/glispy/glispy/types"
)

// New will return a new instance of Glisp
func New() (g Glisp) {
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
	setFunc(s, "set-hash-value", core.SetHashValue)
	setFunc(s, "get-hash-value", core.GetHashValue)
	return NewWithScope(s)
}

// NewWithScope will return a new instance of Glisp with a provided scope
func NewWithScope(s types.Scope) (g Glisp) {
	g.sc = s
	return
}

// Glisp is a lisp worker
type Glisp struct {
	sc types.Scope
}

// Eval will evaluate an Expression
func (g *Glisp) Eval(e types.Expression) (out types.Expression, err error) {
	return eval.Eval(g.sc, e)
}

// EvalTokens will evaluate tokens as an Expression
func (g *Glisp) EvalTokens(ts *tokens.Tokens) (out types.Expression, err error) {
	var e types.Expression
	if e, err = types.NewExpression(ts); err != nil {
		return
	}

	return g.Eval(e)
}

// EvalString will evaluate a string as an Expression
func (g *Glisp) EvalString(str string) (out types.Expression, err error) {
	ts := tokens.NewTokens(str)
	return g.EvalTokens(&ts)
}

// EvalFile will evaluate a file as an Expression
func (g *Glisp) EvalFile(filename string) (out types.Expression, err error) {
	var bs []byte
	if bs, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	return g.EvalString(string(bs))
}

// SetFunc allows Go funcs to be set as lisp funcs
func (g *Glisp) SetFunc(key string, fn types.Function) {
	setFunc(g.sc, key, fn)
}
