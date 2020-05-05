package reflector

import (
	"reflect"
	"sync"
)

// New will return a new instance of reflector
func New(tag string) *Reflector {
	var r Reflector
	r.tag = tag
	r.m = make(map[reflect.Type]Cached)
	return &r
}

// Reflector manages parsed reflections for struct types
type Reflector struct {
	mux sync.RWMutex

	// Tag being managed for an instance of reflector
	tag string
	// Lookup map for cached reflections
	m map[reflect.Type]Cached
}

// Get will get a cached reflection for a given type
func (r *Reflector) Get(t reflect.Type) (c Cached, ok bool) {
	r.mux.RLock()
	defer r.mux.RUnlock()
	c, ok = r.m[t]
	return
}

// Create will create a cached reflection for a given type
func (r *Reflector) Create(t reflect.Type) (c Cached) {
	var ok bool
	if c, ok = r.Get(t); ok {
		return
	}

	r.mux.Lock()
	defer r.mux.Unlock()
	c = makeCached(t, r.tag)
	r.m[t] = c
	return
}

// GetOrCreate will attempt to get a cached reflection for a given type. If it doesn't exist, one will be created.
func (r *Reflector) GetOrCreate(t reflect.Type) (c Cached) {
	var ok bool
	if c, ok = r.Get(t); ok {
		return
	}

	return r.Create(t)
}
