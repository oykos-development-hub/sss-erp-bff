package dto

import "bff/structs"

type GetResolutionResponseMS struct {
	Data structs.Resolution `json:"data"`
}

type GetResolutionListResponseMS struct {
	Data []*structs.Resolution `json:"data"`
}

type Resolution struct {
	Id                int              `json:"id"`
	UserProfile       DropdownSimple   `json:"user_profile"`
	ResolutionType    DropdownSimple   `json:"resolution_type"`
	ResolutionPurpose string           `json:"resolution_purpose"`
	DateOfStart       structs.JSONDate `json:"date_of_start"`
	DateOfEnd         structs.JSONDate `json:"date_of_end"`
	CreatedAt         string           `json:"created_at"`
	UpdatedAt         string           `json:"updated_at"`
	FileId            int              `json:"file_id"`
}
