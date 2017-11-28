package glisp

import (
	"fmt"
	"testing"

	"github.com/missionMeteora/journaler"
)

const (
	program = "(begin (define r 10) (* pi (* r r)))"
)

func TestGlisp(t *testing.T) {
	g := NewGlisp()
	tkns := NewTokens(`(+ 1 3 (+ 2 5))`)
	journaler.Debug("Tokens: %v", tkns)

	exp, err := NewExpression(&tkns)
	if err != nil {
		t.Fatal(err)
	}

	journaler.Debug("Expression: %v", exp)
	out, err := g.Eval(exp)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)
}
