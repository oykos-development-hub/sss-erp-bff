package shared

func Pagination(collection []interface{}, page int, size int) []interface{} {
	startIndex := (page - 1) * size
	endIndex := startIndex + size

	if startIndex > len(collection) {
		return Pagination(collection, startIndex-1, size)
	}

	if startIndex < 0 {
		startIndex = 0
		endIndex = startIndex + size
	}

	if endIndex > len(collection) {
		endIndex = len(collection)
	}

	return collection[startIndex:endIndex]
}
