package glisp

import (
	"math"

	"github.com/itsmontoya/glisp/stdlib/core"
	gmath "github.com/itsmontoya/glisp/stdlib/math"
	"github.com/itsmontoya/glisp/utils"

	"github.com/itsmontoya/glisp/scope"
	"github.com/itsmontoya/glisp/types"
)

// NewGlisp will return a new instance of Glisp
func NewGlisp() (g Glisp) {
	g.sc = scope.NewRoot()
	g.setEnvFn("println", core.Println)
	g.setEnvFn("+", core.Add)
	g.setEnvFn("*", gmath.Multiply)
	g.setEnvFn("define", core.Define)
	g.setEnvFn("defun", core.Defun)
	g.setEnvFn("begin", core.Begin)
	g.setEnvFn(">", core.GreaterThan)
	g.setEnvFn("<", core.LessThan)

	g.sc.Put("greeting", "Hello world")
	g.sc.Put("pi", types.Number(math.Pi))
	return
}

// Glisp is a lisp worker
type Glisp struct {
	sc scope.Scope
}

func (g *Glisp) setEnvFn(key string, fn utils.Func) {
	g.sc.Put(types.Symbol(key), fn)
}

// Eval will evaluate an Expression
func (g *Glisp) Eval(e types.Expression) (out types.Expression, err error) {
	return utils.Eval(g.sc, e)
}
