package glisp

// NewTokens will return a new set of tokens
func NewTokens(program string) Tokens {
	tks := toTokens(splitSpaces(expand(program)))
	return tks
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
