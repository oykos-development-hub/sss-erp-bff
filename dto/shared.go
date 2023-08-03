package dto

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
	Items   interface{} `json:"items"`
	Summary interface{} `json:"summary"`
}

func ErrorResponse(err error) Response {
	return Response{
		Status:  "error",
		Message: err.Error(),
	}
}

type ResponseSingle struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Item    interface{} `json:"item,omitempty"`
}
