package shared

import (
	"reflect"
)

func FilterByProperty(collection []interface{}, property string, value interface{}, contain ...bool) []interface{} {
	if len(property) == 0 {
		return nil
	}
	containValue := false
	if len(contain) > 0 {
		containValue = contain[0]
	}

	var isValueString = reflect.TypeOf(value).Kind() == reflect.String
	var filteredData []interface{}

	for _, item := range collection {
		switch item := item.(type) {
		case map[string]interface{}:
			if item[property] != value {
				filteredData = append(filteredData, item)
			}
		default:
			s := reflect.ValueOf(item)
			if s.Kind() == reflect.Ptr {
				s = s.Elem()
			}
			v := s.FieldByName(property)

			if v.IsValid() {
				var sourceValue = v.Interface()

				if isValueString {
					if (!containValue && sourceValue != value) || (containValue && !StringContains(sourceValue.(string), value.(string))) {
						filteredData = append(filteredData, item)
					}
				} else if sourceValue != value {
					filteredData = append(filteredData, item)
				}
			}
		}
	}

	return filteredData
}
