package reflector

import "reflect"

func makeCached(t reflect.Type, tag string) (c Cached) {
	c = make(Cached)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		alias, ok := field.Tag.Lookup(tag)
		if !ok {
			continue
		}

		c[alias] = newCachedEntry(i, field.Type)
	}

	return
}

// Cached manages the reflection fields by key for a given type
type Cached map[string]*CachedEntry

func newCachedEntry(index int, rType reflect.Type) *CachedEntry {
	var e CachedEntry
	e.Index = index
	e.Type = rType
	return &e
}

// CachedEntry holds the index and type of a field
type CachedEntry struct {
	Index int
	Type  reflect.Type
}
