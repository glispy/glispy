package glispy

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/glispy/glispy/eval"
	"github.com/glispy/glispy/reader"
	"github.com/glispy/glispy/scope"
	"github.com/glispy/glispy/stdlib/core"
	"github.com/glispy/glispy/stdlib/math"
	"github.com/glispy/glispy/types"
)

// New will return a new instance of Glispy
func New() (g Glispy) {
	s := scope.NewRoot()
	setFunc(s, "quote", core.Quote)
	setFunc(s, "define", core.Define)
	setFunc(s, "defun", core.Defun)
	setFunc(s, "begin", core.Begin)
	setFunc(s, "println", core.Println)
	setFunc(s, "+", core.Add)
	setFunc(s, "*", math.Multiply)
	setFunc(s, ">", core.GreaterThan)
	setFunc(s, "<", core.LessThan)
	setFunc(s, "square", math.Square)
	setFunc(s, "make-hash-map", core.MakeHashMap)
	setFunc(s, "get-value", core.GetValue)
	setFunc(s, "set-value", core.SetValue)
	setFunc(s, "remove-value", core.RemoveValue)
	// TODO: Bring this back when net library has been implemented
	// setFunc(s, "http-get", net.HTTPGetRequest)
	return NewWithScope(s)
}

// NewWithScope will return a new instance of Glispy with a provided scope
func NewWithScope(s types.Scope) (g Glispy) {
	g.readmacros = scope.NewRoot()
	g.macros = scope.NewRoot()
	setFunc(g.macros, "'", core.ToQuoteMacro)
	g.sc = s
	return
}

// Glispy is a lisp worker
type Glispy struct {
	// Underlying scope
	sc types.Scope

	// Read-Macros scope, used during read
	readmacros types.Scope
	// Macros scope, used during compile
	macros types.Scope
}

// Eval will evaluate an Expression
func (g *Glispy) Eval(e types.Expression) (out types.Expression, err error) {
	// Run macro expansion pass
	//if e, err = eval.Eval(g.macros, e); err != nil {
	//	err = fmt.Errorf("error encountered during macro expansion phase: %v", err)
	//	return
	//}

	return eval.Eval(g.sc, e)
}

// EvalReader will evaluate an io.Reader as a series of input characters
func (g *Glispy) EvalReader(input io.Reader) (out types.Expression, err error) {
	r := reader.New(input, g.readmacros)
	var exp types.Expression
	if exp, err = r.Read(); err != nil {
		return
	}

	return g.Eval(exp)
}

// EvalString will evaluate a string as an Expression
func (g *Glispy) EvalString(str string) (out types.Expression, err error) {
	r := strings.NewReader(str)
	return g.EvalReader(r)
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
func (g *Glispy) CallFunc(key string, args ...types.Atom) (out types.Expression, err error) {
	var (
		exp types.Expression
		ok  bool
	)

	if exp, ok = g.sc.Get(types.Symbol(key)); !ok {
		err = fmt.Errorf("key of \"%s\" not found", key)
		return
	}

	var fn types.Function
	if fn, ok = exp.(types.Function); !ok {
		err = fmt.Errorf("invalid type, cannot assert %T as function", exp)
		return
	}

	return fn(g.sc, types.List(args))
}
