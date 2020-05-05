package scope

import (
	"testing"

	"github.com/glispy/glispy/types"
)

const (
	testKey  = types.Symbol("foo")
	testVal1 = types.String("bar")
	testVal2 = types.String("baz")
)

func TestScope(t *testing.T) {
	var s Scope
	root := NewRoot()
	s = NewFunc(root)

	root.Put(testKey, testVal1)
	if v, _ := s.Get(testKey); v != testVal1 {
		t.Fatalf("invalid value, expected %v and received %v", testVal1, v)
	}

	s.Put(testKey, testVal2)
	if v, _ := s.Get(testKey); v != testVal2 {
		t.Fatalf("invalid value, expected %v and received %v", testVal2, v)
	}
}
