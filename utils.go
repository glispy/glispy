package glisp

import (
	"regexp"
	"strings"

	"github.com/missionMeteora/journaler"
)

var splitRegExp = regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)'`)

func expand(program string) string {
	program = strings.Replace(program, "(", " ( ", -1)
	return strings.Replace(program, ")", " ) ", -1)
}

func splitSpaces(program string) []string {
	return splitRegExp.FindAllString(program, -1)
}

func toTokens(split []string) (ts Tokens) {
	for _, str := range split {
		ts.Push(str)
	}

	return
}

func toExpression(ts *Tokens, token Token) (e Expression, err error) {
	switch token {
	case "(":
		return NewList(ts)
	case ")":
		err = ErrUnexpectedCloseParens
		return

	default:
		return NewAtom(token)
	}
}

func println(args List) (exp Expression, err error) {
	vals := make([]interface{}, len(args))
	for i, v := range args {
		vals[i] = v
	}

	journaler.Notification("Glisp: %v", vals...)
	return
}

// Fn is the function type
type Fn func(args List) (Expression, error)
