package glisp

import (
	"fmt"
	"strings"
)

func expand(program string) string {
	program = strings.Replace(program, "(", " ( ", -1)
	return strings.Replace(program, ")", " ) ", -1)
}

func splitSpaces(program string) []string {
	return strings.Split(program, " ")
}

func toTokens(split []string) (ts Tokens) {
	for _, str := range split {
		ts.Push(str)
	}

	return
}

func println(args []interface{}) {
	fmt.Println(args...)
}
