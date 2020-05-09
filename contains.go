package cont

import (
	"reflect"
	"strings"
)

// Contains returns true if there is something in the target
//   └── Supports type: string, slice, array, map
func Contains(target, obj interface{}) (_ bool) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.String:
		if strings.Contains(target.(string), obj.(string)) {
			return true
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return
}
