package reader

import (
	"fmt"
	"strings"
	"testing"

	"github.com/glispy/glispy/scope"
	"github.com/glispy/glispy/types"
)

func TestReader_Read(t *testing.T) {
	var err error
	src := `(
	define 'foo "bar"
)`

	sc := scope.NewRoot()
	r := New(strings.NewReader(src), sc)

	var exp types.Expression
	if exp, err = r.Read(); err != nil {
		t.Fatal(err)
	}

	fmt.Println("EXP!", exp)
}
