package core

import (
	"fmt"
	"reflect"

	"github.com/glispy/glispy/reflector"
	"github.com/glispy/glispy/types"
)

var rl = reflector.New("glispy")

// GetValue will get a value from a data structure
func GetValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var target types.Atom
	if target, err = args.GetAtom(0); err != nil {
		err = fmt.Errorf("error accessing first argument: %v", err)
		return
	}

	var lKey types.String
	if lKey, err = args.GetString(1); err != nil {
		return
	}

	return getValueFromAtom(target, string(lKey))
}

// SetValue will set a value within a data structure
func SetValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var target types.Atom
	if target, err = args.GetAtom(0); err != nil {
		err = fmt.Errorf("error accessing first argument: %v", err)
		return
	}

	var lKey types.String
	if lKey, err = args.GetString(1); err != nil {
		err = fmt.Errorf("error accessing second argument: %v", err)
		return
	}

	var value types.Atom
	if value, err = args.GetAtom(2); err != nil {
		err = fmt.Errorf("error accessing third argument: %v", err)
		return
	}

	return setValueToAtom(target, string(lKey), value)
}

// RemoveValue will remove a value within a data structure
func RemoveValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var target types.Atom
	if target, err = args.GetAtom(0); err != nil {
		err = fmt.Errorf("error accessing first argument: %v", err)
		return
	}

	var lKey types.String
	if lKey, err = args.GetString(1); err != nil {
		return
	}

	return removeValueFromAtom(target, string(lKey))
}

func getValueFromAtom(target types.Atom, key string) (exp types.Expression, err error) {
	var ok bool
	switch v := target.(type) {
	case Getter:
		exp, ok = v.GetGlispValue(key)
	case map[string]interface{}:
		exp, ok = v[key]

	default:
		if exp, ok, err = getReflectValueFromAtom(target, key); err != nil {
			return
		}
	}

	if !ok {
		err = fmt.Errorf("field of \"%s\" not found", key)
		return
	}

	return
}

func setValueToAtom(target types.Atom, key string, value types.Atom) (exp types.Expression, err error) {
	var ok bool
	switch v := target.(type) {
	case Setter:
		if ok, err = v.SetGlispValue(key, value); err != nil {
			return
		}
	case map[string]interface{}:
		v[key] = value
		ok = true

	default:
		if exp, ok, err = setReflectValueToAtom(target, key, value); err != nil {
			return
		}
	}

	if !ok {
		err = fmt.Errorf("field of \"%s\" not found", key)
		return
	}

	exp = value
	return
}

func removeValueFromAtom(target types.Atom, key string) (exp types.Expression, err error) {
	var ok bool
	switch v := target.(type) {
	case Remover:
		ok = v.RemoveGlispValue(key)
	case map[string]interface{}:
		delete(v, key)

	default:
		if ok, err = removeReflectValueFromAtom(target, key); err != nil {
			return
		}
	}

	if !ok {
		err = fmt.Errorf("field of \"%s\" not found", key)
		return
	}

	return
}

func getReflectValueFromAtom(target types.Atom, key string) (exp types.Expression, ok bool, err error) {
	// Create target from first argument
	rTarget := reflect.ValueOf(target)
	// Get target's kind
	kind := rTarget.Kind()

	if kind == reflect.Ptr {
		rTarget = reflect.Indirect(rTarget)
		kind = rTarget.Kind()
	}

	if !rTarget.CanSet() {
		err = fmt.Errorf("type of %v cannot be set", kind)
		return
	}

	var rval reflect.Value
	switch kind {
	case reflect.Struct:
		c := rl.GetOrCreate(rTarget.Type())
		rval = rTarget.Field(c[key])
	case reflect.Map:
		rval = rTarget.MapIndex(reflect.ValueOf(key))

	default:
		err = fmt.Errorf("type of %v not supported", kind)
		return
	}

	if rval.IsZero() {
		return
	}

	exp = rval.Interface()
	ok = true
	return
}

func setReflectValueToAtom(target types.Atom, key string, value types.Atom) (exp types.Expression, ok bool, err error) {
	// Create target from first argument
	rTarget := reflect.ValueOf(target)
	// Get target's kind
	kind := rTarget.Kind()

	if kind == reflect.Ptr {
		rTarget = reflect.Indirect(rTarget)
		kind = rTarget.Kind()
	}

	if !rTarget.CanSet() {
		err = fmt.Errorf("type of %v cannot be set", kind)
		return
	}

	switch kind {
	case reflect.Struct:
		c := rl.GetOrCreate(rTarget.Type())
		field, exists := c[key]
		if !exists {
			err = fmt.Errorf("field of \"%s\" does not exist", key)
			return
		}

		reflectedValue := reflect.ValueOf(value)
		targetField := rTarget.Field(field)

		// Check to see if the value is convertable to the field's type
		if !reflectedValue.Type().ConvertibleTo(targetField.Type()) {
			// Value is not convertable to field's type, return
			err = fmt.Errorf("type of %v is not convertable to %v", reflectedValue.Type(), targetField.Type())
			return
		}

		// Convert value to field's type
		convertedVal := reflectedValue.Convert(targetField.Type())

		// Set target field as the converted value
		targetField.Set(convertedVal)

	case reflect.Map:
		rTarget.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))

	default:
		err = fmt.Errorf("type of %v not supported", kind)
		return
	}

	exp = value
	ok = true
	return
}

func removeReflectValueFromAtom(target types.Atom, key string) (ok bool, err error) {
	// Create target from first argument
	rTarget := reflect.ValueOf(target)
	// Get target's kind
	kind := rTarget.Kind()

	if kind == reflect.Ptr {
		rTarget = reflect.Indirect(rTarget)
		kind = rTarget.Kind()
	}

	if !rTarget.CanSet() {
		err = fmt.Errorf("type of %v cannot be set", kind)
		return
	}

	switch kind {
	case reflect.Struct:
		if rTarget.CanSet() {
			err = fmt.Errorf("type of %v cannot be set", kind)
			return
		}

		c := rl.GetOrCreate(rTarget.Type())

		var field int
		field, ok = c[key]
		if !ok {
			return
		}

		// Set field value
		fieldValue := rTarget.Field(field)

		// Get zero value of type
		zeroValue := reflect.Zero(fieldValue.Type())

		// Set target field to zero value
		rTarget.Field(field).Set(zeroValue)
	case reflect.Map:
		// Set field value
		fieldValue := rTarget.MapIndex(reflect.ValueOf(key))

		// Get zero value of type
		zeroValue := reflect.Zero(fieldValue.Type())

		// Set map index as zero value (this will remove the key from the map)
		rTarget.SetMapIndex(reflect.ValueOf(key), zeroValue)

	default:
		err = fmt.Errorf("type of %v not supported", kind)
		return
	}

	ok = true
	return
}

// Getter represents a data structure that can be accessed from within Glisp
type Getter interface {
	GetGlispValue(key string) (value types.Atom, ok bool)
}

// Setter represents a data structure that can be set from within Glisp
type Setter interface {
	SetGlispValue(key string, value types.Atom) (ok bool, err error)
}

// Remover represents a data structure that can be removed from within Glisp
type Remover interface {
	RemoveGlispValue(key string) (ok bool)
}
