package shared

import (
	"reflect"
)

func FindByProperty(collection []interface{}, property string, value interface{}, contain ...bool) []interface{} {
	if len(property) == 0 {
		return nil
	}
	containValue := false
	if len(contain) > 0 {
		containValue = contain[0]
	}

	var isValueString = reflect.TypeOf(value).Kind() == reflect.String
	var matches []interface{}

	for _, item := range collection {
		switch item := item.(type) {
		case map[string]interface{}:
			if v, ok := item[property]; ok && v == value {
				matches = append(matches, item)
			}
		default:
			s := reflect.ValueOf(item)
			if s.Kind() == reflect.Ptr {
				s = s.Elem()
			}
			v := s.FieldByName(property)

			if v.IsValid() {
				sourceValue := v.Interface()

				if isValueString {
					if (!containValue && sourceValue == value) || (containValue && StringContains(sourceValue.(string), value.(string))) {
						matches = append(matches, item)
					}
				} else if sourceValue == value {
					matches = append(matches, item)
				}
			}
		}
	}

	return matches
}
