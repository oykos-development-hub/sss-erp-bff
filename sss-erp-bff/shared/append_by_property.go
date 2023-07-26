package shared

import (
	"reflect"
)

func AppendByProperty(collection []interface{}, property string, value interface{}, appendProperty string, appendValue interface{}) []interface{} {
	for index, item := range collection {
		itemValue := reflect.ValueOf(item)

		if itemValue.Kind() == reflect.Map {
			if itemValue.MapIndex(reflect.ValueOf(property)).Interface() == value {
				m := itemValue.Interface().(map[string]interface{})
				var newValue = m[appendProperty].([]interface{})
				newValue = append(newValue, appendValue)
				m[appendProperty] = newValue
				collection[index] = m
				return collection
			}
		} else {
			if itemValue.Kind() == reflect.Ptr {
				itemValue = itemValue.Elem()
			}

			fieldValue := itemValue.FieldByName(property)
			if !fieldValue.IsValid() {
				continue
			}

			if reflect.DeepEqual(fieldValue.Interface(), value) {
				appendFieldValue := itemValue.FieldByName(appendProperty)
				if !appendFieldValue.IsValid() {
					continue
				}

				if appendFieldValue.Kind() == reflect.Slice {
					newSlice := reflect.Append(appendFieldValue, reflect.ValueOf(appendValue))
					appendFieldValue.Set(newSlice)
					collection[index] = itemValue.Interface()
					return collection
				}
			}
		}
	}

	return collection
}
