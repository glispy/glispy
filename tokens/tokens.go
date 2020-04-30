package tokens

import (
	"io/ioutil"
	"os"
)

// NewTokens will return a new set of tokens
func NewTokens(program string) Tokens {
	tks := toTokens(splitSpaces(expand(program)))
	return tks
}

// NewTokensFromFile will return a new set of tokens from a provided filename
func NewTokensFromFile(filename string) (ts Tokens, err error) {
	var f *os.File
	if f, err = os.Open(filename); err != nil {
		return
	}
	defer f.Close()

	var bs []byte
	if bs, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	ts = NewTokens(string(bs))
	return
}

// Tokens represents a list of tokens
type Tokens []Token

// Shift will remove the first token and return it
func (ts *Tokens) Shift() (t Token, ok bool) {
	if len(*ts) == 0 {
		return
	}

	t = (*ts)[0]
	*ts = (*ts)[1:]
	ok = true
	return
}

// Push will add an item to the tokens list
func (ts *Tokens) Push(str string) {
	if len(str) == 0 {
		return
	}

	tsr := *ts
	tsr = append(tsr, Token(str))
	*ts = tsr
	return
}

// Token is a basic token type
type Token string
