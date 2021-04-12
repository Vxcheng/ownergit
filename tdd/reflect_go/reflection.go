package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	xv := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	numberOfValues := 0
	var getField func(int) reflect.Value
	switch xv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < xv.Len(); i++ {
			walkValue(xv.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < xv.NumField(); i++ {
			walkValue(xv.Field(i))
		}
	case reflect.String:
		fn(xv.String())
	case reflect.Map:
		for _, k := range xv.MapKeys() {
			walkValue(xv.MapIndex(k))
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

}

func getValue(x interface{}) reflect.Value {
	xv := reflect.ValueOf(x)
	if xv.Kind() == reflect.Ptr {
		xv = xv.Elem()
	}

	return xv
}
