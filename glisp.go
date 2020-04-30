package glisp

import (
	"math"

	"github.com/itsmontoya/glisp/stdlib/core"
	gmath "github.com/itsmontoya/glisp/stdlib/math"
	"github.com/itsmontoya/glisp/tokens"
	"github.com/itsmontoya/glisp/utils"

	"github.com/itsmontoya/glisp/scope"
	"github.com/itsmontoya/glisp/types"
)

// New will return a new instance of Glisp
func New() (g Glisp) {
	g.sc = scope.NewRoot()
	g.setEnvFn("println", core.Println)
	g.setEnvFn("+", core.Add)
	g.setEnvFn("*", gmath.Multiply)
	g.setEnvFn("define", core.Define)
	g.setEnvFn("defun", core.Defun)
	g.setEnvFn("begin", core.Begin)
	g.setEnvFn(">", core.GreaterThan)
	g.setEnvFn("<", core.LessThan)
	g.setEnvFn("make-hash-map", core.MakeHashMap)
	g.setEnvFn("set-hash-value", core.SetHashValue)
	g.setEnvFn("get-hash-value", core.GetHashValue)

	g.sc.Put("greeting", "Hello world")
	g.sc.Put("pi", types.Number(math.Pi))
	return
}

// Glisp is a lisp worker
type Glisp struct {
	sc scope.Scope
}

func (g *Glisp) setEnvFn(key string, fn types.Function) {
	g.sc.Put(types.Symbol(key), fn)
}

// Eval will evaluate an Expression
func (g *Glisp) Eval(e types.Expression) (out types.Expression, err error) {
	return utils.Eval(g.sc, e)
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

// SetFunc allows Go funcs to be set as lisp funcs
func (g *Glisp) SetFunc(key string, fn types.Function) {
	g.setEnvFn(key, fn)
}
