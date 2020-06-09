package core

import (
	"fmt"
	"reflect"

	"github.com/glispy/glispy/types"
)

// GetIndexValue will get an index value from within a slice
func GetIndexValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var (
		target types.Atom
		index  types.Number
	)

	if err = args.GetValues(&target, &index); err != nil {
		return
	}

	return getIndexValue(target, int(index))
}

// SetIndexValue will set an index value within a slice
func SetIndexValue(sc types.Scope, args types.List) (exp types.Expression, err error) {
	var (
		target types.Atom
		index  types.Number
		value  types.Atom
	)

	if err = args.GetValues(&target, &index, &value); err != nil {
		return
	}

	return setIndexValue(target, int(index), value)
}

func getIndexValue(target types.Atom, index int) (exp types.Expression, err error) {
	switch v := target.(type) {
	case []string:
		exp = v[index]
	case []float32:
		exp = v[index]
	case []float64:
		exp = v[index]
	case []int:
		exp = v[index]
	case []uint:
		exp = v[index]
	case []int64:
		exp = v[index]
	case []uint64:
		exp = v[index]
	case types.List:
		exp = v[index]

	default:
		err = fmt.Errorf("unsupported list type: %T is not supported", target)
		return
	}

	return
}

func setIndexValue(target types.Atom, index int, value types.Atom) (exp types.Expression, err error) {
	var ok bool
	switch v := target.(type) {
	case []string:
		var val string
		if val, ok = value.(string); !ok {
			err = fmt.Errorf("invalid value, expected %T and received %T", val, value)
			return
		}

	case []float32:
		var val float32
		if val, ok = value.(float32); !ok {
			err = fmt.Errorf("invalid value, expected %T and received %T", val, value)
			return
		}

	case []float64:
		var val float64
		if val, ok = value.(float64); !ok {
			err = fmt.Errorf("invalid value, expected %T and received %T", val, value)
			return
		}

	case []int:
		var val int
		if val, ok = value.(int); !ok {
			err = fmt.Errorf("invalid value, expected %T and received %T", val, value)
			return
		}

	case []uint:
		var val uint
		if val, ok = value.(uint); !ok {
			err = fmt.Errorf("invalid value, expected %T and received %T", val, value)
			return
		}

	case []int64:
		var val int64
		if val, ok = value.(int64); !ok {
			err = fmt.Errorf("invalid value, expected %T and received %T", val, value)
			return
		}

	case []uint64:
		var val uint64
		if val, ok = value.(uint64); !ok {
			err = fmt.Errorf("invalid value, expected %T and received %T", val, value)
			return
		}

	case types.List:
		v[index] = value

	default:
		rval := reflect.ValueOf(value)
		if rval.Kind() != reflect.Slice {
			err = fmt.Errorf("unsupported list type: %T is not supported", target)
			return
		}

		reference := rval.Index(index)

		if !reference.CanSet() {
			err = fmt.Errorf("reference type of %T cannot be set", target)
			return
		}

		reference.Set(reflect.ValueOf(value))
		return
	}

	exp = target
	return
}
