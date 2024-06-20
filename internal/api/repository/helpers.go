package repository

import (
	"bff/internal/api/errors"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func makeAPIRequest(method, url string, data interface{}, response interface{}, headers ...map[string]string) ([]*http.Cookie, error) {
	reqBytes, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "json marshal")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, errors.Wrap(err, "http new request")
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
		return nil, errors.Wrap(err, "client do")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "io read all")
	}

	if resp.StatusCode != http.StatusOK {
		apiError := &errors.APIError{
			StatusCode: resp.StatusCode,
		}

		err = json.Unmarshal(body, &apiError)
		if err != nil {
			apiError.Message = "Unexpected error"
		}

		return nil, errors.Wrap(apiError, "backend error")
	}

	if len(body) > 0 && response != nil {
		err = json.Unmarshal(body, response)
		if err != nil {
			return nil, errors.Wrap(err, "json unmarshal")
		}
	}

	return resp.Cookies(), nil
}

func makeAPIRequestWithCookie(method, url string, data interface{}, response interface{}, cookie *http.Cookie, headers ...map[string]string) ([]*http.Cookie, error) {
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

	req.AddCookie(cookie)

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
		apiError := &errors.APIError{
			StatusCode: resp.StatusCode,
		}

		err = json.Unmarshal(body, &apiError)
		if err != nil {
			apiError.Message = "Unexpected error"
		}

		return nil, apiError
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
