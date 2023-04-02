package trcache

import (
	"fmt"
	"reflect"
	"runtime"
)

func StringValue(key any) string {
	switch tkey := key.(type) {
	case string:
		return tkey
	case []byte:
		return string(tkey)
	case fmt.Stringer:
		return tkey.String()
	default:
		return fmt.Sprint(tkey)
	}
}

func getValueName(t reflect.Value) string {
	switch t.Kind() {
	case reflect.Ptr:
		return "*" + getValueName(t.Elem())
	case reflect.Func:
		f := runtime.FuncForPC(t.Pointer())
		if f != nil {
			return f.Name()
		}
		fallthrough
	default:
		return t.Type().Name()
	}
}

func getName(myvar interface{}) string {
	return getValueName(reflect.ValueOf(myvar))
}
