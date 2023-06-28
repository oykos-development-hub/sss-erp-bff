package shared

func MergeMaps(map1, map2 map[string]interface{}, skipEmpty ...bool) map[string]interface{} {
	mergedMap := make(map[string]interface{})

	skipEmptyValue := false
	if len(skipEmpty) > 0 {
		skipEmptyValue = skipEmpty[0]
	}

	for key, value := range map1 {
		mergedMap[key] = value
	}
	for key, value := range map2 {
		if skipEmptyValue {
			switch v := value.(type) {
			case string:
				if v == "" {
					continue
				}
			case int:
				if v == 0 {
					continue
				}
			case float64:
				if v == 0.0 {
					continue
				}
			case nil:
				continue
			default:
				// other types are not checked for empty values
			}
		}
		mergedMap[key] = value
	}
	return mergedMap
}
