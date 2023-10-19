package shared

import (
	"errors"
	"reflect"
)

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

func Paginate(slice interface{}, page, size int) (interface{}, error) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("slice must be a slice type")
	}

	// If the slice is empty, return an empty slice of the same type without error.
	if sliceValue.Len() == 0 {
		return reflect.MakeSlice(sliceValue.Type(), 0, 0).Interface(), nil
	}

	startIndex := (page - 1) * size
	if startIndex < 0 || startIndex > sliceValue.Len() {
		return nil, errors.New("invalid page number")
	}

	endIndex := startIndex + size
	if endIndex > sliceValue.Len() {
		endIndex = sliceValue.Len()
	}

	// Get a slice of the original with the appropriate range and return it.
	paginatedSlice := sliceValue.Slice(startIndex, endIndex)
	return paginatedSlice.Interface(), nil
}
