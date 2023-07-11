package shared

import (
	"bff/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ErrorResponse(message string) dto.Response {
	return dto.Response{
		Status:  "error",
		Message: message,
	}
}

type APIError struct {
	StatusCode int
	Message    string      `json:"error"`
	Data       interface{} `json:"data"`
}

type Response struct {
	Message string `json:"message"`
}

func ErrorResponse(message string) dto.Response {
	return dto.Response{
		Status:  "error",
		Message: message,
	}
}

func (e *APIError) Error() string {
	if e.Data != nil {
		dataBytes, err := json.Marshal(e.Data)
		if err != nil {
			return fmt.Sprintf("API error: %d - %s", e.StatusCode, e.Message)
		}
		return fmt.Sprintf("API error: %d - %s - %s", e.StatusCode, e.Message, string(dataBytes))
	}
	return fmt.Sprintf("API error: %d - %s", e.StatusCode, e.Message)
}

func MakeAPIRequest(method, url string, data interface{}, response interface{}, headers ...map[string]string) ([]*http.Cookie, error) {
	reqBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Append headers to the request
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		apiError := &APIError{
			StatusCode: resp.StatusCode,
		}

		err = json.Unmarshal(body, apiError)
		if err != nil {
			apiError.Message = "Unexpected error"
		}

		return nil, apiError
	}
	if response == nil {
		response = &Response{}
	}

	if len(body) > 0 && response != nil {
		err = json.Unmarshal(body, response)
		if err != nil {
			fmt.Println(string(body))
			return nil, fmt.Errorf("failed to parse HTTP response: %w", err)
		}
	}

	return resp.Cookies(), nil
}
