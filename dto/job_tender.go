package dto

import "bff/structs"

type GetJobTenderResponseMS struct {
	Data structs.JobTenders `json:"data"`
}

type GetJobTenderListResponseMS struct {
	Data []*structs.JobTenders `json:"data"`
}

type JobTenderResponseItem struct {
	Id               int                       `json:"id"`
	OrganizationUnit structs.OrganizationUnits `json:"organization_unit"`
	JobPosition      *structs.JobPositions     `json:"job_position"`
	Type             structs.JobTenderTypes    `json:"type"`
	Description      string                    `json:"description"`
	SerialNumber     string                    `json:"serial_number"`
	AvailableSlots   int                       `json:"available_slots"`
	Active           bool                      `json:"active"`
	DateOfStart      structs.JSONDate          `json:"date_of_start"`
	DateOfEnd        structs.JSONDate          `json:"date_of_end"`
	FileId           int                       `json:"file_id"`
	CreatedAt        string                    `json:"created_at"`
	UpdatedAt        string                    `json:"updated_at"`
}

type GetJobTendersInput struct {
	Active *bool `json:"is_active"`
}

type GetJobTenderApplicationsInput struct {
	Page        *int `json:"page"`
	Size        *int `json:"size"`
	JobTenderID *int `json:"job_tender_id"`
}

type GetJobTenderTypeResponseMS struct {
	Data structs.JobTenderTypes `json:"data"`
}

type GetJobTenderTypeInputMS struct {
	Search *string `json:"search"`
}

type GetJobTenderTypeListResponseMS struct {
	Data []*structs.JobTenderTypes `json:"data"`
}

type GetJobTenderApplicationResponseMS struct {
	Data structs.JobTenderApplications `json:"data"`
}

type GetJobTenderApplicationListResponseMS struct {
	Data  []*structs.JobTenderApplications `json:"data"`
	Total int                              `json:"total"`
}

type JobTenderApplicationResponseItem struct {
	Id                 int              `json:"id"`
	UserProfile        *DropdownSimple  `json:"user_profile"`
	JobTender          DropdownSimple   `json:"job_tender"`
	Active             bool             `json:"active"`
	Type               string           `json:"type"`
	FirstName          *string          `json:"first_name"`
	LastName           *string          `json:"last_name"`
	Nationality        *string          `json:"nationality"`
	DateOfBirth        structs.JSONDate `json:"date_of_birth"`
	DateOfApplication  structs.JSONDate `json:"date_of_application"`
	OfficialPersonalID *string          `json:"official_personal_id"`
	Evaluation         *string          `json:"evaluation"`
	Status             string           `json:"status"`
	FileId             int              `json:"file_id"`
	CreatedAt          string           `json:"created_at"`
	UpdatedAt          string           `json:"updated_at"`
}
