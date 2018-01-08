package flatten

import "reflect"

func Flatten(input interface{}) []interface{} {
	res := []interface{}{}

	if isSlice(input) {
		for _, item := range input.([]interface{}) {
			if item == nil {
				continue
			}

			if isSlice(item) {
				res = append(res, Flatten(item)...)
			} else {
				res = append(res, item)
			}
		}
	}

	return res
}

func isSlice(item interface{}) bool {
	return reflect.ValueOf(item).Kind() == reflect.Slice
}
