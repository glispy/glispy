package glisp

import (
	"fmt"
	"testing"

	"github.com/itsmontoya/glisp/tokens"
	"github.com/itsmontoya/glisp/types"
	"github.com/janne/go-lisp/lisp"
)

const (
	//	program = "(begin (define r 10) (* pi (* r r)))"

	program = "(begin (define r 10) (println (* pi (* r r)))))"
	square  = `(define square (x) (* x x))`
)

var (
	glispSink  types.Expression
	goLispSink lisp.Value
)

func TestGlisp(t *testing.T) {
	g := New()
	//tkns := NewTokens(`(begin (define foo "bar") (println foo pi))`)
	//tkns := tokens.NewTokens(`(if (> 3 2) 11 22)`)
	tkns := tokens.NewTokens(`(
	begin 
		(defun square (x)
			(* x x)
		)
		(println (
				square (
					+ 3 3
				)
			)
		)
)`)
	//tkns := NewTokens(`(begin (println ("foo")) (println ("bar")))`)

	exp, err := types.NewExpression(&tkns)
	if err != nil {
		t.Fatal(err)
	}

	out, err := g.Eval(exp)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestGlispAdd(t *testing.T) {
	var (
		exp types.Expression
		val types.Expression
		err error
	)

	g := New()
	tkns := tokens.NewTokens(`(+ 1 3 (+ 2 5))`)
	if exp, err = types.NewExpression(&tkns); err != nil {
		t.Fatal(err)
	}

	if val, err = g.Eval(exp); err != nil {
		t.Fatal(err)
	}

	if val.(types.Number) != 11 {
		t.Fatalf("invalid value, expected %v and received %v", 11, val)
	}
}

func BenchmarkGlispAdd(b *testing.B) {
	var (
		exp types.Expression
		val types.Expression
		err error
	)

	g := New()

	for i := 0; i < b.N; i++ {
		tkns := tokens.NewTokens(`(+ 1 3 (+ 2 5))`)
		if exp, err = types.NewExpression(&tkns); err != nil {
			b.Fatal(err)
		}

		if val, err = g.Eval(exp); err != nil {
			b.Fatal(err)
		}

		if val.(types.Number) != 11 {
			b.Fatalf("invalid value, expected %v and received %v", 11, val)
		}

		glispSink = val
	}

	b.ReportAllocs()
}

func BenchmarkGlispSquare(b *testing.B) {
	var (
		exp  types.Expression
		val  types.Expression
		tkns tokens.Tokens
		err  error
	)

	g := New()
	tkns = tokens.NewTokens(square)
	if exp, err = types.NewExpression(&tkns); err != nil {
		b.Fatal(err)
	}

	if _, err = g.Eval(exp); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tkns = tokens.NewTokens(`(square 3)`)
		if exp, err = types.NewExpression(&tkns); err != nil {
			b.Fatal(err)
		}

		if val, err = g.Eval(exp); err != nil {
			b.Fatal(err)
		}

		if val.(types.Number) != 9 {
			b.Fatalf("invalid value, expected %v and received %v", 11, val)
		}

		glispSink = val
	}

	b.ReportAllocs()
}

func BenchmarkGlispSquare_PreProcessed(b *testing.B) {
	var (
		exp  types.Expression
		val  types.Expression
		tkns tokens.Tokens
		err  error
	)

	g := New()
	tkns = tokens.NewTokens(square)
	if exp, err = types.NewExpression(&tkns); err != nil {
		b.Fatal(err)
	}

	if _, err = g.Eval(exp); err != nil {
		b.Fatal(err)
	}

	tkns = tokens.NewTokens(`(square 3)`)
	if exp, err = types.NewExpression(&tkns); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if val, err = g.Eval(exp); err != nil {
			b.Fatal(err)
		}

		if val.(types.Number) != 9 {
			b.Fatalf("invalid value, expected %v and received %v", 11, val)
		}

		glispSink = val
	}

	b.ReportAllocs()
}

func BenchmarkGlispAdd_PreProcessed(b *testing.B) {
	var (
		exp types.Expression
		val types.Expression
		err error
	)

	g := New()
	tkns := tokens.NewTokens(`(+ 1 3 (+ 2 5))`)

	if exp, err = types.NewExpression(&tkns); err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if val, err = g.Eval(exp); err != nil {
			b.Fatal(err)
		}

		if val.(types.Number) != 11 {
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

func BenchmarkGoLisp(b *testing.B) {
	var (
		val lisp.Value
		err error
	)

	if val, err = lisp.EvalString(square); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if val, err = lisp.EvalString(`(square 3)`); err != nil {
			b.Fatal(err)
		}

		if val.Number() != 9 {
			b.Fatalf("invalid value, expected %v and received %v", 11, val)
		}

		goLispSink = val
	}

	b.ReportAllocs()
}
