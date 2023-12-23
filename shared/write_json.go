package shared

import (
	"encoding/json"
	"errors"
	"os"
)

func WriteJSON(path string, data []interface{}) error {
	if len(path) == 0 {
		return errors.New("argument 'path' cannot be empty")
	}
	if data == nil {
		return errors.New("argument 'data' cannot be empty")
	}
	// Marshal the slice to JSON
	var jsonData, err = json.Marshal(data)

	if err != nil {
		return errors.New("argument 'data' cannot be converted to JSON object")
	}
	// Write the JSON data to a file
	err = os.WriteFile(path, jsonData, 0644)

	if err != nil {
		return errors.New("data cannot be writen to a file on provided path")
	}

	return nil
}
