package dto

import "bff/structs"

type GetFileResponse struct {
	Data *structs.File `json:"data"`
}

type GetFileResponsePom struct {
	Data GetFileResponse `json:"data"`
}
