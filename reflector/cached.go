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

		c[alias] = i
	}

	return
}

// Cached manages the reflection fields by key for a given type
type Cached map[string]int
