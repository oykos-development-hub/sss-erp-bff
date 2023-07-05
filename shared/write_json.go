package shared

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func WriteJson(path string, data []interface{}) error {
	if len(path) == 0 {
		return errors.New("Argument 'path' cannot be empty!")
	}
	if data == nil {
		return errors.New("Argument 'data' cannot be empty!")
	}
	// Marshal the slice to JSON
	var jsonData, err = json.Marshal(data)

	if err != nil {
		return errors.New("Argument 'data' cannot be converted to JSON object!")
	}
	// Write the JSON data to a file
	err = ioutil.WriteFile(path, jsonData, 0644)

	if err != nil {
		return errors.New("Data cannot be writen to a file on provided path!")
	}

	return nil
}
