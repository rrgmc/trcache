package trcache

import "fmt"

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