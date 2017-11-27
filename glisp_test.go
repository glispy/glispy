package glisp

import (
	"fmt"
	"testing"
)

const (
	program = "(begin (define r 10) (* pi (* r r)))"
)

func TestGlisp(t *testing.T) {
	g := NewGlisp()

	exp, err := NewExpression(NewTokens("greeting"))
	if err != nil {
		t.Fatal(err)
	}

	out, err := g.Eval(exp)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)
	//Parse(program)
	//	>>> parse(program)
	//	['begin', ['define', 'r', 10], ['*', 'pi', ['*', 'r', 'r']]]

}
