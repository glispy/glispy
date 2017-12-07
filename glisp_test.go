package glisp

import (
	"fmt"
	"testing"

	"github.com/janne/go-lisp/lisp"
)

const (
	//	program = "(begin (define r 10) (* pi (* r r)))"

	program = "(begin (define r 10) (println (* pi (* r r)))))"
)

var (
	glispSink  Expression
	goLispSink lisp.Value
)

func TestGlisp(t *testing.T) {
	g := NewGlisp()
	//tkns := NewTokens(`(begin (define foo "bar") (println foo pi))`)
	tkns := NewTokens(program)
	//tkns := NewTokens(`(begin (println ("foo")) (println ("bar")))`)

	exp, err := NewExpression(&tkns)
	if err != nil {
		t.Fatal(err)
	}

	out, err := g.Eval(exp)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)
}

func BenchmarkGlispAdd(b *testing.B) {
	var (
		exp Expression
		val Expression
		err error
	)

	g := NewGlisp()

	for i := 0; i < b.N; i++ {
		tkns := NewTokens(`(+ 1 3 (+ 2 5))`)
		if exp, err = NewExpression(&tkns); err != nil {
			b.Fatal(err)
		}

		if val, err = g.Eval(exp); err != nil {
			b.Fatal(err)
		}

		if val.(Number) != 11 {
			b.Fatalf("invalid value, expected %v and received %v", 11, val)
		}

		glispSink = val
	}

	b.ReportAllocs()
}

func BenchmarkGlispAdd_PreProcessed(b *testing.B) {
	var (
		exp Expression
		val Expression
		err error
	)

	g := NewGlisp()
	tkns := NewTokens(`(+ 1 3 (+ 2 5))`)

	if exp, err = NewExpression(&tkns); err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if val, err = g.Eval(exp); err != nil {
			b.Fatal(err)
		}

		if val.(Number) != 11 {
			b.Fatalf("invalid value, expected %v and received %v", 11, val)
		}

		glispSink = val
	}

	b.ReportAllocs()
}

func BenchmarkGoLispAdd(b *testing.B) {
	var (
		val lisp.Value
		err error
	)
	for i := 0; i < b.N; i++ {
		if val, err = lisp.EvalString(`(+ 1 3 (+ 2 5))`); err != nil {
			b.Fatal(err)
		}

		if val.Number() != 11 {
			b.Fatalf("invalid value, expected %v and received %v", 11, val)
		}

		goLispSink = val
	}

	b.ReportAllocs()
}
