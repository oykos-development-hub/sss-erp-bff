package files

import (
	"bff/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var response FileResponseData

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}
	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	defer writer.Close()

	part, err := writer.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	backendResponse, status, err := makeBackendRequest(http.MethodPost, config.FILES_ENDPOINT, &requestBody)
	if err != nil {
		handleError(w, err, status)
		return
	}
	defer backendResponse.Body.Close()

	decoder := json.NewDecoder(backendResponse.Body)
	if err := decoder.Decode(&response); err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, _ = w.Write(jsonResponse)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var response FileGetByIDResponse
	id := chi.URLParam(r, "id")

	backendFileURL := config.FILES_ENDPOINT + "/" + id

	backendResponse, status, err := makeBackendRequest(http.MethodDelete, backendFileURL, nil)
	if err != nil {
		handleError(w, err, status)
		return
	}
	defer backendResponse.Body.Close()

	response.Message = "File was successfully deleted"
	_ = MarshalAndWriteJSON(w, response)
}

func MultipleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var response FileGetByIDResponse

	var input MultipleDeleteFiles
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(input)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	backendFileURL := config.FILES_MULTIPLE_DELETE_ENDPOINT

	backendResponse, status, err := makeBackendRequest(http.MethodPost, backendFileURL, bytes.NewBuffer(jsonData))
	if err != nil {
		handleError(w, err, status)
		return
	}
	defer backendResponse.Body.Close()

	response.Message = "Files were successfully deleted"
	_ = MarshalAndWriteJSON(w, response)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	backendFileURL := config.FILES_DOWNLOAD_ENDPOINT + "/" + id

	backendResponse, status, err := makeBackendRequest(http.MethodGet, backendFileURL, nil)
	if err != nil {
		handleError(w, err, status)
		return
	}
	defer backendResponse.Body.Close()

	fileData, status, err := GetFileData(id)

	if err != nil {
		handleError(w, err, status)
		return
	}

	filename := fileData.Data.Data.Name
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	_, err = io.Copy(w, backendResponse.Body)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func OverviewHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	backendFileURL := config.FILES_OVERVIEW_ENDPOINT + "/" + id

	backendResponse, status, err := makeBackendRequest(http.MethodGet, backendFileURL, nil)
	if err != nil {
		handleError(w, err, status)
		return

	}
	defer backendResponse.Body.Close()

	fileData, status, err := GetFileData(id)

	if err != nil {
		handleError(w, err, status)
		return
	}

	filename := fileData.Data.Data.Name
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	_, err = io.Copy(w, backendResponse.Body)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func GetFileData(fileID string) (*FileGetByIDResponse, int, error) {
	backendURL := fmt.Sprintf(config.FILES_ENDPOINT + "/" + fileID)

	response, status, err := makeBackendRequest(http.MethodGet, backendURL, nil)

	if err != nil {
		return nil, status, err
	}

	var fileData FileGetByIDResponse
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&fileData); err != nil {
		return nil, status, err
	}

	return &fileData, status, nil
}
