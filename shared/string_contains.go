package shared

import (
	"strings"
)

func StringContains(source string, target string) bool {
	var sourceValue = strings.ReplaceAll(strings.ToLower(source), " ", "")
	var targetValue = strings.ReplaceAll(strings.ToLower(target), " ", "")

	return strings.Contains(sourceValue, targetValue)
}
