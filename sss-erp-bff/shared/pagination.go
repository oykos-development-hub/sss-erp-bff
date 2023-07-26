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
		return reflect.MakeSlice(sliceValue.Type(), 0, 0).Interface(), errors.New("slice must be a slice type")
	}

	startIndex := (page - 1) * size
	endIndex := startIndex + size

	if startIndex < 0 || startIndex >= sliceValue.Len() {
		return reflect.MakeSlice(sliceValue.Type(), 0, 0).Interface(), errors.New("invalid page number")
	}

	if endIndex > sliceValue.Len() {
		endIndex = sliceValue.Len()
	}

	paginatedSlice := sliceValue.Slice(startIndex, endIndex)
	return paginatedSlice.Interface(), nil
}
