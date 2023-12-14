package files

import (
	"bff/config"
	"bytes"
	"encoding/json"
	"errors"
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

	files := r.MultipartForm.File["file"]

	if len(files) == 0 {
		handleError(w, errors.New("you must provide files"), http.StatusBadRequest)
		return
	}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			handleError(w, err, http.StatusBadRequest)
			return
		}

		defer file.Close()

		var requestBody bytes.Buffer
		writer := multipart.NewWriter(&requestBody)

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

		writer.Close() //mora ovako jer se iz nekog neznanog razloga ne kopira fajl kako treba
		backendResponse, status, err := makeBackendRequest(http.MethodPost, config.FILES_ENDPOINT, &requestBody, writer.FormDataContentType())
		if err != nil {
			handleError(w, err, status)
			return
		}
		defer backendResponse.Body.Close()

		var resp FileResponseData

		decoder := json.NewDecoder(backendResponse.Body)
		if err := decoder.Decode(&resp); err != nil {
			handleError(w, err, http.StatusInternalServerError)
			return
		}

		response.Data = append(response.Data, resp.Data...)

	}

	response.Status = "success"
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, _ = w.Write(jsonResponse)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var response SingleFileResponse
	id := chi.URLParam(r, "id")

	backendFileURL := config.FILES_ENDPOINT + "/" + id

	backendResponse, status, err := makeBackendRequest(http.MethodDelete, backendFileURL, nil, "")
	if err != nil {
		handleError(w, err, status)
		return
	}
	defer backendResponse.Body.Close()

	response.Message = "File was successfully deleted"
	response.Status = "success"
	_ = MarshalAndWriteJSON(w, response)
}

func MultipleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var response SingleFileResponse

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

	backendResponse, status, err := makeBackendRequest(http.MethodPost, backendFileURL, bytes.NewBuffer(jsonData), "")
	if err != nil {
		handleError(w, err, status)
		return
	}
	defer backendResponse.Body.Close()

	response.Status = "success"
	response.Message = "Files were successfully deleted"
	_ = MarshalAndWriteJSON(w, response)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	backendFileURL := config.FILES_DOWNLOAD_ENDPOINT + "/" + id

	backendResponse, status, err := makeBackendRequest(http.MethodGet, backendFileURL, nil, "")
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

	filename := fileData.Data.Name
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

	backendResponse, status, err := makeBackendRequest(http.MethodGet, backendFileURL, nil, "")
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

	filename := fileData.Data.Name
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	_, err = io.Copy(w, backendResponse.Body)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}

func GetFileData(fileID string) (*SingleFileResponse, int, error) {
	backendURL := fmt.Sprintf(config.FILES_ENDPOINT + "/" + fileID)

	response, status, err := makeBackendRequest(http.MethodGet, backendURL, nil, "")

	if err != nil {
		return nil, status, err
	}

	var fileData SingleFileResponse
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&fileData); err != nil {
		return nil, status, err
	}

	return &fileData, status, nil
}
