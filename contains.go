package cont

import (
	"reflect"
	"strings"
)

// Contains returns true if there is something in the target
//   └── Supports type: string, slice, array, map
func Contains(target, obj interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.String:
		return strings.Contains(target.(string), obj.(string))
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		return targetValue.MapIndex(reflect.ValueOf(obj)).IsValid()
	}
	return false
}
