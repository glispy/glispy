package tokens

import (
	"regexp"
	"strings"
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
