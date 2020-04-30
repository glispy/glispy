package core

import (
	"fmt"

	"github.com/itsmontoya/glisp/types"
)

// MakeHashMap will make a new hash map
func MakeHashMap(sc types.Scope, args types.List) (exp types.Expression, err error) {
	key, ok := args[0].(types.Symbol)
	if !ok {
		err = fmt.Errorf("invalid key type: expected symbol and received %T", args[0])
		return
	}

	list := types.List{
		key,
		make(HashMap),
	}

	if _, err = Define(sc, list); err != nil {
		return
	}

	return
}

// GetHashValue will get a value from a HashMap
func GetHashValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	sym, ok := args[0].(types.Symbol)
	if !ok {
		err = fmt.Errorf("invalid key type: expected symbol and received %T", args[0])
		return
	}

	val, ok := sc.Get(sym)
	if !ok {
		err = fmt.Errorf("symbol \"%s\" does not exist within scope", sym)
		return
	}

	hm, ok := val.(HashMap)
	if !ok {
		err = fmt.Errorf("expected hashmap as the first argument, received %T", val)
		return
	}

	key, ok := args[1].(types.String)
	if !ok {
		err = fmt.Errorf("expected key with type of string as the second argument, received %T", args[1])
		return
	}

	exp = hm[key]
	return
}

// SetHashValue will set a value within a HashMap
func SetHashValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	sym, ok := args[0].(types.Symbol)
	if !ok {
		err = fmt.Errorf("invalid key type: expected symbol and received %T", args[0])
		return
	}

	val, ok := sc.Get(sym)
	if !ok {
		err = fmt.Errorf("symbol \"%s\" does not exist within scope", sym)
		return
	}

	hm, ok := val.(HashMap)
	if !ok {
		err = fmt.Errorf("expected hashmap as the first argument, received %T", val)
		return
	}

	key, ok := args[1].(types.String)
	if !ok {
		err = fmt.Errorf("expected key with type of string as the second argument, received %T", args[1])
		return
	}

	hm[key] = args[2]
	return
}

// RemoveHashValue will remove a key within a HashMap
func RemoveHashValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	hm, ok := args[0].(HashMap)
	if !ok {
		err = fmt.Errorf("expected hashmap as the first argument, received %T", args[0])
		return
	}

	key, ok := args[1].(types.String)
	if !ok {
		err = fmt.Errorf("expected key with type of string as the second argument, received %T", args[1])
		return
	}

	delete(hm, key)
	return
}

// HashMap represents a hash map
type HashMap map[types.String]interface{}
