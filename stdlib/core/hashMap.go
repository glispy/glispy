package core

import (
	"fmt"

	"github.com/glispy/glispy/types"
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
		make(map[string]interface{}),
	}

	if _, err = Define(sc, list); err != nil {
		return
	}

	return
}

// GetHashValue will get a value from a HashMap
func GetHashValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var atom types.Atom
	if atom, err = args.GetAtom(0); err != nil {
		return
	}

	hm, ok := atom.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("expected hashmap as the first argument, received %T", atom)
		return
	}

	var key types.String
	if key, err = args.GetString(1); err != nil {
		return
	}

	exp = hm[string(key)]
	return
}

// SetHashValue will set a value within a HashMap
func SetHashValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var atom types.Atom
	if atom, err = args.GetAtom(0); err != nil {
		return
	}

	hm, ok := atom.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("expected hashmap as the first argument, received %T", atom)
		return
	}

	var key types.String
	if key, err = args.GetString(1); err != nil {
		err = fmt.Errorf("error getting key: %v", err)
		return
	}

	var val types.Atom
	if val, err = args.GetAtom(2); err != nil {
		err = fmt.Errorf("error getting value: %v", err)
		return
	}

	hm[string(key)] = val
	return
}

// RemoveHashValue will remove a key within a HashMap
func RemoveHashValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var atom types.Atom
	if atom, err = args.GetAtom(0); err != nil {
		return
	}

	hm, ok := atom.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("expected hashmap as the first argument, received %T", atom)
		return
	}

	var key types.String
	if key, err = args.GetString(1); err != nil {
		err = fmt.Errorf("error getting key: %v", err)
		return
	}

	delete(hm, string(key))
	return
}
