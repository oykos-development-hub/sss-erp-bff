package errors

import (
	"bff/internal/api/dto"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
)

type APIError struct {
	StatusCode int
	Message    string      `json:"error"`
	Data       interface{} `json:"data"`
}

type Response struct {
	Message string `json:"message"`
}

func (e APIError) Error() string {
	if e.Data != nil {
		dataBytes, err := json.Marshal(e.Data)
		if err != nil {
			return fmt.Sprintf("API error: %d - %s", e.StatusCode, e.Message)
		}
		return fmt.Sprintf("API error: %d - %s - %s", e.StatusCode, e.Message, string(dataBytes))
	}
	return fmt.Sprintf("API error: %d - %s", e.StatusCode, e.Message)
}

func HandleAPIError(err error) (dto.Response, error) {
	_, file, line, _ := runtime.Caller(1) // 1 is the number of stack frames to ascend
	log.Printf("Error occurred in file %s at line %d: %v", file, line, err)

	if apiError, ok := err.(APIError); ok {
		return ErrorResponse(apiError.Message, apiError.Data), nil
	}
	return ErrorResponse(err.Error()), nil
}

func ErrorResponse(message string, data ...interface{}) dto.Response {
	response := dto.Response{
		Status:  "error",
		Message: message,
	}

	if len(data) > 0 && data[0] != nil {
		response.Data = data[0]
	}

	return response
}
