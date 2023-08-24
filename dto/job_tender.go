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
	Active           bool                      `json:"active"`
	DateOfStart      structs.JSONDate          `json:"date_of_start"`
	DateOfEnd        structs.JSONDate          `json:"date_of_end"`
	FileId           int                       `json:"file_id"`
	CreatedAt        string                    `json:"created_at"`
	UpdatedAt        string                    `json:"updated_at"`
}

type GetJobTenderApplicationsInput struct {
	Page          *int `json:"page"`
	Size          *int `json:"size"`
	JobTenderID   *int `json:"job_tender_id"`
	UserProfileId *int `json:"user_profile_id"`
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
	Id                 int             `json:"id"`
	UserProfile        *DropdownSimple `json:"user_profile"`
	JobTender          *DropdownSimple `json:"job_tender"`
	Active             bool            `json:"active"`
	Type               string          `json:"type"`
	FirstName          string          `json:"first_name"`
	LastName           string          `json:"last_name"`
	OfficialPersonalID string          `json:"official_personal_id"`
	DateOfBirth        string          `json:"date_of_birth"`
	Nationality        string          `json:"nationality"`
	Evaluation         string          `json:"evaluation"`
	DateOfAplication   string          `json:"date_of_application"`
	Status             string          `json:"status"`
	FileId             int             `json:"file_id"`
	CreatedAt          string          `json:"created_at"`
	UpdatedAt          string          `json:"updated_at"`
}
