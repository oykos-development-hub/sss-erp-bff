package shared

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	// "fmt"
)

func ReadJson(path string, responseType interface{}) ([]interface{}, error) {
	// Create a new instance of the response type
	response := reflect.New(reflect.SliceOf(reflect.TypeOf(responseType))).Interface()

	if len(path) == 0 {
		return nil, errors.New("Argument 'path' cannot be empty!")
	}
	// Read the contents of the JSON file
	resp, err := http.Get(path)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	file, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &response)

	if err != nil {
		return nil, err
	}

	// Convert the response slice to a []interface{} slice
	result := make([]interface{}, reflect.ValueOf(response).Elem().Len())
	for i := 0; i < reflect.ValueOf(response).Elem().Len(); i++ {
		result[i] = reflect.ValueOf(response).Elem().Index(i).Interface()
	}

	// Return the JSON data as the response
	return result, nil
}
