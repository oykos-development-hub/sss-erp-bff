package files

import (
	"bff/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
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

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var response FileResponseData

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		response.Message = "File is not valid"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		response.Message = "Error during fetching file"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		response.Message = "Error during creating form file"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		response.Message = "Error during copying file to form field"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Close()

	backendRequest, err := http.NewRequest(http.MethodPost, config.FILES_ENDPOINT, &requestBody)
	if err != nil {
		response.Message = "Error during making request"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	backendRequest.Header.Set("Content-Type", writer.FormDataContentType())

	httpClient := &http.Client{}
	backendResponse, err := httpClient.Do(backendRequest)
	if err != nil {
		response.Message = "Error during sending file"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer backendResponse.Body.Close()

	if backendResponse.StatusCode != http.StatusOK {
		response.Message = "Error during processing file on backend"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(backendResponse.Body)
	if err := decoder.Decode(&response); err != nil {
		response.Message = "Error during reading response"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		response.Message = "Error during JSON marshaling"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, _ = w.Write(jsonResponse)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var response FileGetByIDResponse
	id := cutAfterLastOccurrence(r.URL.Path, "/")

	backendFileURL := config.FILES_ENDPOINT + "/" + id

	backendRequest, err := http.NewRequest(http.MethodDelete, backendFileURL, nil)
	if err != nil {
		response.Message = "Error during making request"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpClient := &http.Client{}
	backendResponse, err := httpClient.Do(backendRequest)
	if err != nil {
		response.Message = "Error during executing request"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer backendResponse.Body.Close()

	if backendResponse.StatusCode != http.StatusOK {
		response.Message = "Error during deleting file"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Message = "File was successfully deleted"
	_ = MarshalAndWriteJSON(w, response)
}

func MultipleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var response FileGetByIDResponse

	var input MultipleDeleteFiles
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		response.Message = "Error during decoding input"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		response.Message = "Error during encoding input"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	backendFileURL := config.FILES_MULTIPLE_DELETE_ENDPOINT

	backendRequest, err := http.NewRequest(http.MethodPost, backendFileURL, bytes.NewBuffer(jsonData))

	if err != nil {
		response.Message = "Error during making request"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	httpClient := &http.Client{}
	backendResponse, err := httpClient.Do(backendRequest)
	if err != nil {
		response.Message = "Error during executing request"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer backendResponse.Body.Close()

	if backendResponse.StatusCode != http.StatusOK {
		response.Message = "Error during deleting files"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Message = "Files were successfully deleted"
	_ = MarshalAndWriteJSON(w, response)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	var response FileGetByIDResponse
	id := cutAfterLastOccurrence(r.URL.Path, "/")

	backendFileURL := config.FILES_DOWNLOAD_ENDPOINT + "/" + id

	backendResponse, err := http.Get(backendFileURL)
	if err != nil {
		response.Message = "Error during making request"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer backendResponse.Body.Close()

	if backendResponse.StatusCode != http.StatusOK {
		response.Message = "Error during fetching file"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fileData, err := GetFileData(id)

	if err != nil {
		response.Message = "Error during reading file data"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	filename := fileData.Data.Data.Name
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	_, err = io.Copy(w, backendResponse.Body)
	if err != nil {
		response.Message = "Error during sending file to frontend"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func OverviewHandler(w http.ResponseWriter, r *http.Request) {
	var response FileGetByIDResponse
	id := cutAfterLastOccurrence(r.URL.Path, "/")

	backendFileURL := config.FILES_OVERVIEW_ENDPOINT + "/" + id

	backendResponse, err := http.Get(backendFileURL)
	if err != nil {
		response.Message = "Error during creating request"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	defer backendResponse.Body.Close()

	if backendResponse.StatusCode != http.StatusOK {
		response.Message = "Error during fetching file from backend"
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fileData, err := GetFileData(id)

	if err != nil {
		response.Message = "Error during reading file data"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	filename := fileData.Data.Data.Name
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	_, err = io.Copy(w, backendResponse.Body)
	if err != nil {
		response.Message = "Error during sending file to frontend"
		response.Error = err.Error()
		_ = MarshalAndWriteJSON(w, response)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func cutAfterLastOccurrence(input, substring string) string {
	lastIndex := strings.LastIndex(input, substring)
	if lastIndex == -1 {
		return input
	}
	result := input[lastIndex+len(substring):]
	return result
}

func GetFileData(fileID string) (*FileGetByIDResponse, error) {
	backendURL := fmt.Sprintf(config.FILES_ENDPOINT + "/" + fileID)

	response, err := http.Get(backendURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("backend returned non-ok status code: %d", response.StatusCode)
	}

	var fileData FileGetByIDResponse
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&fileData); err != nil {
		return nil, err
	}

	return &fileData, nil
}
