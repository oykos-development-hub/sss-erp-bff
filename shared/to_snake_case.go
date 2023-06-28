package shared

import (
	"fmt"
	"unicode"
)

func ToSnakeCase(s string, lowercase bool) string {
	var result string

	if !IsString(s) {
		fmt.Printf("ToSnakeCase | Argument 's' can not be empty!")
		return ""
	}

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result += "_"
			}
			result += string(unicode.ToLower(r))
		} else {
			result += string(r)
		}
	}

	return result
}
