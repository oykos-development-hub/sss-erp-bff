package dto

import (
	"bff/structs"
)

type GetSystematizationsResponseMS struct {
	Data  []structs.Systematization `json:"data"`
	Total int                       `json:"total"`
}

type GetSystematizationResponseMS struct {
	Data structs.Systematization `json:"data"`
}

type GetSystematizationsInput struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"page_size"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Active             *int    `json:"active"`
	Year               *string `json:"year"`
	Search             *string `json:"search"`
}

type SystematizationOverviewResponse struct {
	ID                 int                                `json:"id"`
	UserProfileID      int                                `json:"user_profile_id"`
	OrganizationUnitID int                                `json:"organization_unit_id"`
	Description        string                             `json:"description"`
	SerialNumber       string                             `json:"serial_number"`
	Active             int                                `json:"active"`
	DateOfActivation   *string                            `json:"date_of_activation"`
	CreatedAt          string                             `json:"created_at"`
	UpdatedAt          string                             `json:"updated_at"`
	FileIDs            []int                              `json:"file_ids"`
	OrganizationUnit   *structs.OrganizationUnits         `json:"organization_unit"`
	Sectors            *[]OrganizationUnitsSectorResponse `json:"sectors"`
	ActiveEmployees    []structs.ActiveEmployees          `json:"active_employees"`
	Files              []FileDropdownSimple               `json:"files"`
}
