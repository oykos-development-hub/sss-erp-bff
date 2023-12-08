package files

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func MarshalAndWriteJSON(w http.ResponseWriter, obj interface{}) error {
	jsonResponse, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, "Error during JSON marshaling", http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)

	if err != nil {
		return err
	}

	return nil
}

func openExcelFile(r *http.Request) (*excelize.File, error) {
	maxFileSize := int64(100 * 1024 * 1024) // file maximum 100 MB

	err := r.ParseMultipartForm(maxFileSize)
	if err != nil {
		return nil, err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	xlsFile, err := excelize.OpenReader(file)

	if err != nil {
		return nil, err
	}

	return xlsFile, nil
}

func handleError(w http.ResponseWriter, err error, statusCode int) {
	log.Printf("Error: %s - %v", err.Error(), err)
	w.WriteHeader(statusCode)
	_ = MarshalAndWriteJSON(w, errorResponse{Message: err.Error()})
}

func makeBackendRequest(method, url string, body io.Reader) (*http.Response, int, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, 0, err
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	if resp.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		var errorStruct errorResponse
		if err := decoder.Decode(&errorStruct); err != nil {
			return nil, resp.StatusCode, err
		}

		if errorStruct.Message != "" {
			return nil, resp.StatusCode, errors.New(errorStruct.Message)
		}

		resp.Body.Close()
		return nil, resp.StatusCode, fmt.Errorf("backend returned non-OK status: %d", resp.StatusCode)
	}

	return resp, 0, nil
}
