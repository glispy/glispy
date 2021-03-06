package glispy

import (
	"fmt"
	"strings"
	"testing"

	"github.com/glispy/glispy/reader"
	"github.com/glispy/glispy/types"
	"github.com/janne/go-lisp/lisp"
)

const (
	square = `(defun square (x) (* x x))`
)

var (
	glispySink types.Expression
	goLispSink lisp.Value
)

func TestGlispy(t *testing.T) {
	g := New()
	src := `(
	(defun square (x)
		(* x x)
	)
	(println (
			square (
				+ 3 3
			)
		)
	)
)`

	out, err := g.EvalString(src)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestGlispyDefine(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	g := New()
	src := `(
	(define x 1337)
	(println x)
)`

	//(define (quote x) 1337)

	if val, err = g.EvalString(src); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Returned", val)
}

func TestGlispyAdd(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	g := New()
	src := `(+ 1 3 (+ 2 5))`

	if val, err = g.EvalString(src); err != nil {
		t.Fatal(err)
	}

	if val.(types.Number) != 11 {
		t.Fatalf("invalid value, expected %v and received %v", 11, val)
	}
}

func TestGetSetValue_map(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	g := New()

	if val, err = g.EvalString(`(
		(make-hash-map 'foo)
		(set-value foo "bar" 1337)
		(get-value foo "bar")
	)`); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Value", val)
}

func TestGetValue_nil_map(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	g := New()
	var m map[string]string
	g.sc.Put("foo", m)
	if val, err = g.EvalString(`(
		(get-value foo "bar")
	)`); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Value", val)
}

func TestGetSetValue_struct(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	type S struct {
		A string  `glispy:"a"`
		B float32 `glispy:"b"`
		C string
	}

	g := New()
	s := S{}

	g.sc.Put(types.Symbol("foo"), &s)
	if val, err = g.EvalString(`(
		(set-value foo "a" "hello world")
		(set-value foo "b" 1337)
		(get-value foo "a")
	)`); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Value", val)
}

func TestGetIndexValue(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	g := New()

	if val, err = g.EvalString(`(
		(define foo '(1 2 3))
		(println (get-index-value foo 1))
	)`); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Value", val)
}

func TestSetIndexValue(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	g := New()

	if val, err = g.EvalString(`(
		(define foo '(1 2 3))
		(set-index-value foo 1 1337)
		(println (get-index-value foo 1))
	)`); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Value", val)
}

func TestGlispyMacro_println(t *testing.T) {
	var err error
	g := New()

	if _, err = g.EvalString(`(
		defmacro speak (x) (
			println x
		)
	)`); err != nil {
		t.Fatal(err)
	}

	var compiled types.Expression
	if compiled, err = g.CompileString(`(
			speak 26
		)`); err != nil {
		t.Fatal(err)
	}

	var val types.Expression
	if val, err = g.Eval(compiled); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Value", val)
}

func TestGlispyMacro_shrug(t *testing.T) {
	var err error
	g := New()

	if _, err = g.EvalString(`(
	defmacro shrug (x) (
		println x "¯\_(ツ)_/¯"
	)
)`); err != nil {
		t.Fatal(err)
	}

	var compiled types.Expression
	if compiled, err = g.CompileString(`(shrug "Oh well")`); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Compiled", compiled)

	var val types.Expression
	if val, err = g.Eval(compiled); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Value", val)
}

func TestHTTPGet(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	g := New()
	if val, err = g.EvalString(`(
	(define resp (http-get nil "https://cat-fact.herokuapp.com/facts/random"))
	(get-value resp "text")
)`); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Value", val)
}

func TestEval_multi_quote(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	g := New()
	if val, err = g.EvalString(`''''''(1 2 3)`); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Value", val)
}

func TestEndWithValue(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	g := New()

	if val, err = g.EvalString(`(
		(make-hash-map 'foo)
		(set-value foo "bar" 1337)
		foo
	)`); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Value", val)
}

func TestIfStatement(t *testing.T) {
	var (
		val types.Expression
		err error
	)

	g := New()

	if val, err = g.EvalString(`(
		(define foo '(1 2 3))
		(define length (get-length foo))
		(if (> length 5) ("yes") ("no"))
	)`); err != nil {
		t.Fatal(err)
	}

	var (
		str types.String
		ok  bool
	)

	if str, ok = val.(types.String); !ok {
		t.Fatalf("invalid return type, expected %T and received %T", str, val)
	}

	if str != "no" {
		t.Fatalf("invalid return value, expected %v and received %v", "no", str)
	}
}
func BenchmarkGlispyAdd(b *testing.B) {
	var (
		val types.Expression
		err error
	)

	g := New()

	for i := 0; i < b.N; i++ {
		src := `(+ 1 3 (+ 2 5))`

		if val, err = g.EvalString(src); err != nil {
			b.Fatal(err)
		}

		if val.(types.Number) != 11 {
			b.Fatalf("invalid value, expected %v and received %v", 11, val)
		}

		glispySink = val
	}

	b.ReportAllocs()
}

func BenchmarkGlispySquare(b *testing.B) {
	var (
		val types.Expression
		err error
	)

	g := New()

	if _, err = g.EvalString(square); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if val, err = g.EvalString(`(square 3)`); err != nil {
			b.Fatal(err)
		}

		if val.(types.Number) != 9 {
			b.Fatalf("invalid value, expected %v and received %v", 11, val)
		}

		glispySink = val
	}

	b.ReportAllocs()
}

func BenchmarkGlispySquare_PreProcessed(b *testing.B) {
	var (
		exp types.Expression
		val types.Expression
		err error
	)

	g := New()
	if _, err = g.EvalString(square); err != nil {
		b.Fatal(err)
	}

	r := reader.New(strings.NewReader(`(square 3)`), g.readermacros)

	if exp, err = r.Read(); err != nil {
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

		glispySink = val
	}

	b.ReportAllocs()
}

func BenchmarkGlispyAdd_PreProcessed(b *testing.B) {
	var (
		exp types.Expression
		val types.Expression
		err error
	)

	g := New()
	r := reader.New(strings.NewReader(`(+ 1 3 (+ 2 5))`), g.readermacros)

	if exp, err = r.Read(); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if val, err = g.Eval(exp); err != nil {
			b.Fatal(err)
		}

		if val.(types.Number) != 11 {
			b.Fatalf("invalid value, expected %v and received %v", 11, val)
		}

		glispySink = val
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
