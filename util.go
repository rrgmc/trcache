package trcache

import (
	"fmt"
	"reflect"
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

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
