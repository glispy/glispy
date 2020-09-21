package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/glispy/glispy/common"
	"github.com/glispy/glispy/types"
	"github.com/hatchify/errors"
)

// ErrInvalidAtom is returned when an atom is invalid
const ErrInvalidAtom = errors.Error("atom must be a number or a string")

// New will return a new instance of reader
func New(input io.Reader, sc types.Scope) *Reader {
	var r Reader
	r.r = bufio.NewReader(input)
	r.sc = sc
	return &r
}

// NewFromFile will return a new instance of reader from a file
func NewFromFile(filename string) (r *Reader, err error) {
	var f *os.File
	if f, err = os.Open(filename); err != nil {
		return
	}
	defer f.Close()
	return
	//return New(f)
}

// Reader represents a list of reader
type Reader struct {
	r  io.RuneReader
	sc types.Scope

	// State used during Read
	// Rune buffer
	buf []rune
	// Whether or not the read process is currently within quotes
	inQuotes bool
	// List of processed tokens, waiting to be pulled
	tokens []Token

	lastRune    rune
	unreadState bool
}

func (r *Reader) flush() (token Token, ok bool) {
	if len(r.buf) == 0 {
		return
	}

	token = Token(r.buf)
	ok = true
	r.buf = r.buf[:0]
	return
}

func (r *Reader) pushProtected(char rune) (token Token, ok bool) {
	if r.inQuotes {
		r.buf = append(r.buf, char)
		return
	}

	if token, ok = r.flush(); ok {
		// We've retrieved a token from our flush, set our unread state to true and bail out
		r.unreadState = true
		return
	}

	token = Token(char)
	ok = true
	return
}

func (r *Reader) pushWhitespace(char rune) (token Token, ok bool) {
	if r.inQuotes {
		r.buf = append(r.buf, char)
		return
	}

	return r.flush()
}

func (r *Reader) pushQuotes(char rune) (token Token, ok bool) {
	if !r.inQuotes && len(r.buf) > 0 {
		// We're transitioning from a symbol into a new string, set unreadState to true
		r.unreadState = true
		return r.flush()
	}

	r.buf = append(r.buf, char)
	// Invert inQuotes state
	r.inQuotes = !r.inQuotes

	if !r.inQuotes {
		return r.flush()
	}

	return
}

func (r *Reader) unreadChar() {
	// We're transitioning from a symbol into a new string, set unreadState to true
	r.unreadState = true
}

// ReadChar will read a single character
func (r *Reader) ReadChar() (char rune, err error) {
	if r.unreadState {
		char = r.lastRune
		r.unreadState = false
		return
	}

	if char, _, err = r.r.ReadRune(); err != nil {
		return
	}

	r.lastRune = char
	return
}

// ReadNextNonWhitespaceChar will read the next non-whitespace single character
func (r *Reader) ReadNextNonWhitespaceChar() (char rune, err error) {
	for {
		if char, err = r.ReadChar(); err != nil {
			return
		}

		if !unicode.IsSpace(char) {
			return
		}
	}
}

// ReadToken will read a single token
func (r *Reader) ReadToken() (token Token, err error) {
	var (
		char rune
		ok   bool
	)

	for char, err = r.ReadChar(); err == nil; char, err = r.ReadChar() {
		switch {
		case char == '"':
			token, ok = r.pushQuotes(char)
		case isProtected(char):
			token, ok = r.pushProtected(char)
		case unicode.IsSpace(char):
			token, ok = r.pushWhitespace(char)

		default:
			r.buf = append(r.buf, char)
		}

		if ok {
			return
		}
	}

	if token, ok = r.flush(); ok {
		return
	}

	err = io.EOF
	return
}

func (r *Reader) performReadMacro() (exp types.Expression, ok bool, err error) {
	var char rune
	if char, err = r.ReadChar(); err != nil {
		return
	}

	switch char {
	case '\'':
		if exp, err = quote(r, char); err != nil {
			return
		}

	case ':':
		if exp, err = method(r, char); err != nil {
			return
		}

	default:
		r.unreadState = true
		return
	}

	ok = true
	return
}

func (r *Reader) Read() (exp types.Expression, err error) {
	var ok bool
	if exp, ok, err = r.performReadMacro(); ok || err != nil {
		return
	}

	var token Token
	if token, err = r.ReadToken(); err != nil {
		return
	}

	switch token {
	case "(":
		exp, err = r.newList()
		return
	case ")":
		err = common.ErrUnexpectedCloseParens
		return

	default:
		return r.newAtom(token)
	}
}

func (r *Reader) newList() (l types.List, err error) {
	return r.appendList(nil)
}

func (r *Reader) appendList(in types.List) (l types.List, err error) {
	l = in
	for {
		var char rune
		if char, err = r.ReadNextNonWhitespaceChar(); err != nil {
			return
		}

		if char == ')' {
			return
		}

		r.unreadState = true

		var e types.Expression
		if e, err = r.Read(); err != nil {
			return
		}

		l = append(l, e)
	}
}

func (r *Reader) newAtom(token Token) (a types.Atom, err error) {
	if a, err = r.newNil(token); err == nil {
		return
	}

	if a, err = r.newSymbol(token); err == nil {
		return
	}

	if a, err = r.newNumber(token); err == nil {
		return
	}

	if a, err = r.newString(token); err == nil {
		return
	}

	err = ErrInvalidAtom
	return
}

func (r *Reader) newNil(t Token) (n types.Nil, err error) {
	if t != "nil" {
		err = types.ErrInvalidNil
		return
	}

	return
}

func (r *Reader) newSymbol(t Token) (s types.Symbol, err error) {
	if symbolRegExp.Match([]byte(t)) {
		err = types.ErrInvalidSymbol
		return
	}

	s = types.Symbol(t)
	return
}

func (r *Reader) newNumber(t Token) (n types.Number, err error) {
	var float float64
	if float, err = strconv.ParseFloat(string(t), 32); err != nil {
		return
	}

	n = types.Number(float)
	return
}

func (r *Reader) newString(t Token) (s types.String, err error) {
	if t[0] != '"' {
		err = types.ErrInvalidString
		return
	}

	if t[len(t)-1] != '"' {
		err = types.ErrInvalidString
		return
	}

	s = types.String(t[1 : len(t)-1])
	return
}

var symbolRegExp = regexp.MustCompile(`[^a-zA-Z_<>+*-]`)

// Token is a basic token type
type Token string

func quote(r *Reader, char rune) (exp types.Expression, err error) {
	var parsed types.Expression
	if parsed, err = r.Read(); err != nil {
		return
	}

	var l types.List
	l = append(l, types.Symbol("quote"), parsed)
	exp = l
	return
}

func method(r *Reader, char rune) (exp types.Expression, err error) {
	var token Token
	if token, err = r.ReadToken(); err != nil {
		return
	}

	spl := strings.SplitN(string(token), ".", 2)
	if len(spl) == 1 {
		err = fmt.Errorf("expected a struct reference and method call, received \"%s\"", token)
		return
	}

	var l types.List
	l = append(l, types.Symbol("method"))
	l = append(l, types.Symbol(spl[0]))
	l = append(l, types.String(spl[1]))

	if l, err = r.appendList(l); err != nil {
		err = fmt.Errorf("error appending list:%v", err)
		return
	}

	exp = l
	r.unreadState = true
	return
}
