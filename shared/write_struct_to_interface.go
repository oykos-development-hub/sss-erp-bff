package shared

import (
	"fmt"
	"reflect"
)

var WriteStructToInterface = func(source interface{}) map[string]interface{} {
	item := make(map[string]interface{})

	if source == nil {
		fmt.Printf("WriteStructToInterface | Argument 'source' must not be empty!")
		return item
	}

	sourceValue := reflect.ValueOf(source)

	if sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}

	for i := 0; i < sourceValue.NumField(); i++ {
		field := sourceValue.Type().Field(i)
		fieldName := field.Name
		key := ToSnakeCase(fieldName)
		value := sourceValue.FieldByName(fieldName).Interface()
		item[key] = value
	}

	return item
}
