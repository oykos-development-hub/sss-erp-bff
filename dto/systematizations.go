package dto

import "bff/structs"

type GetSystematizationsResponseMS struct {
	Data  []structs.Systematization `json:"data"`
	Total int                       `json:"total"`
}

type GetSystematizationResponseMS struct {
	Data structs.Systematization `json:"data"`
}

type GetSystematizationsInput struct {
	Page               *int  `json:"page"`
	Size               *int  `json:"page_size"`
	OrganizationUnitID *int  `json:"organization_unit_id"`
	Active             *bool `json:"active"`
}

type SystematizationOverviewResponse struct {
	Id                 int                                `json:"id"`
	UserProfileId      int                                `json:"user_profile_id"`
	OrganizationUnitId int                                `json:"organization_unit_id"`
	Description        string                             `json:"description"`
	SerialNumber       string                             `json:"serial_number"`
	Active             bool                               `json:"active"`
	DateOfActivation   structs.JSONDate                   `json:"date_of_activation"`
	CreatedAt          string                             `json:"created_at"`
	UpdatedAt          string                             `json:"updated_at"`
	FileId             int                                `json:"file_id"`
	OrganizationUnit   *structs.OrganizationUnits         `json:"organization_unit"`
	Sectors            *[]OrganizationUnitsSectorResponse `json:"sectors"`
}

type SystematizationOverviewResponse struct {
	Id                 int                                `json:"id"`
	UserProfileId      int                                `json:"user_profile_id"`
	OrganizationUnitId int                                `json:"organization_unit_id"`
	Description        string                             `json:"description"`
	SerialNumber       string                             `json:"serial_number"`
	Active             bool                               `json:"active"`
	DateOfActivation   structs.JSONDate                   `json:"date_of_activation"`
	CreatedAt          string                             `json:"created_at"`
	UpdatedAt          string                             `json:"updated_at"`
	FileId             int                                `json:"file_id"`
	OrganizationUnit   *structs.OrganizationUnits         `json:"organization_unit"`
	Sectors            *[]OrganizationUnitsSectorResponse `json:"sectors"`
}
