package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	value := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			walkValue(value.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkValue(value.Field(i))
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}
	case reflect.String:
		fn(value.String())
	}

}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value
}
