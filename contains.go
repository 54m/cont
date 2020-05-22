/*
Package cont - Checks if the target is contained in a string, slice or map
*/
package cont

import (
	"reflect"
	"strings"
)

// Contains returns true if there is something in the target
// supports type: string, slice, array, map
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
		iter := targetValue.MapRange()
		for iter.Next() {
			k := iter.Key().Interface()
			v := iter.Value().Interface()
			if k == obj {
				return true
			}
			if v == obj {
				return true
			}
		}
	}
	return false
}
