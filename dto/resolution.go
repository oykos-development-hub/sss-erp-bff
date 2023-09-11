package dto

import (
	"bff/structs"
	"time"
)

type GetResolutionResponseMS struct {
	Data structs.Resolution `json:"data"`
}

type GetResolutionListResponseMS struct {
	Data []*structs.Resolution `json:"data"`
}

type EmployeeResolutionListInput struct {
	From *time.Time `json:"from"`
	To   *time.Time `json:"to"`
}

type Resolution struct {
	Id                int            `json:"id"`
	UserProfile       DropdownSimple `json:"user_profile"`
	ResolutionType    DropdownSimple `json:"resolution_type"`
	ResolutionPurpose string         `json:"resolution_purpose"`
	DateOfStart       string         `json:"date_of_start"`
	DateOfEnd         string         `json:"date_of_end"`
	Value             string         `json:"value"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         string         `json:"updated_at"`
	FileId            int            `json:"file_id"`
}
