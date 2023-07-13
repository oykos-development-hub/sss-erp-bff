package dto

import "bff/structs"

type GetResolutionResponseMS struct {
	Data structs.Resolution `json:"data"`
}

type GetResolutionListResponseMS struct {
	Data []*structs.Resolution `json:"data"`
}
