package core

import (
	"fmt"
	"reflect"

	"github.com/glispy/glispy/types"
)

// Method will call a method of a given data structure
func Method(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var (
		target types.Atom
		key    types.String
	)

	if err = args.GetValues(&target, &key); err != nil {
		return
	}

	// TODO: Add caching here to speed this up. Under the hood, stdlib iterates through exported
	// fields to find the match. If we can get the index, we can save ourselves lookup time here.
	fn := reflect.ValueOf(target).MethodByName(string(key))

	// Check to see if the method exists
	if fn.IsZero() {
		// Method does not exist, return
		err = fmt.Errorf("method of \"%s\" not found within type of %T", key, target)
		return
	}

	vals := getReflectValues(args[2:])

	output := fn.Call(vals)
	if len(output) == 0 {
		exp = types.Nil{}
		return
	}

	return processReflectOutput(output)
}

func processReflectOutput(output []reflect.Value) (exp types.Expression, err error) {
	if len(output) == 0 {
		exp = types.Nil{}
		return
	}

	// Get last return value
	last := output[len(output)-1]

	var ok bool
	if ok, err = getError(last); ok {
		// Last value is an error, remove it from the returning expression
		output = output[:len(output)-1]
	}

	var l types.List
	l = append(l, types.Symbol("quote"))

	for _, val := range output {
		l = append(l, val.Interface())
	}

	exp = l
	return
}

func getReflectValues(l types.List) (rs []reflect.Value) {
	rs = make([]reflect.Value, 0, 2)
	for _, atom := range l {
		switch n := atom.(type) {
		case types.String:
			atom = string(n)
		}

		rs = append(rs, reflect.ValueOf(atom))
	}

	return
}

var errorInterface = reflect.TypeOf((*error)(nil)).Elem()

func getError(v reflect.Value) (match bool, err error) {
	t := v.Type()
	if match = t.Implements(errorInterface); !match {
		return
	}

	val := v.Interface()
	if val == nil {
		return
	}

	err = val.(error)
	return
}
