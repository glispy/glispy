package core

import (
	"fmt"

	"github.com/glispy/glispy/types"
)

// MakeHashMap will make a new hash map
func MakeHashMap(sc types.Scope, args types.List) (exp types.Expression, err error) {
	exp = make(map[string]interface{})

	// Check to see if we have any args provided with this call
	if len(args) == 0 {
		// No args provided, calling func is expecting to use exp return value. Bail out!
		return
	}

	key, ok := args[0].(types.Symbol)
	if !ok {
		err = fmt.Errorf("invalid key type: expected symbol and received %T", args[0])
		return
	}

	sc.Put(key, exp)
	return
}
