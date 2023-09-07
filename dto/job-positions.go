package dto

import "bff/structs"

type GetJobPositionsResponseMS struct {
	Data  []structs.JobPositions `json:"data"`
	Total int                    `json:"total"`
}

type GetJobPositionResponseMS struct {
	Data structs.JobPositions `json:"data"`
}

type GetJobPositionsInput struct {
	Page    *int    `json:"page"`
	Size    *int    `json:"page_size"`
	Search  *string `json:"search"`
	IsJudge *bool   `json:"is_judge"`
}

type GetEmployeesInOrganizationUnitInput struct {
	PositionInOrganizationUnit *int  `json:"position_in_organization_unit_id"`
	UserProfileId              *int  `json:"user_profile_id"`
	Active                     *bool `json:"active"`
}

type JobPositionsInSectorResponse struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	Abbreviation     string `json:"abbreviation"`
	Description      string `json:"description"`
	Requirements     string `json:"requirements"`
	SerialNumber     string `json:"serial_number"`
	IsJudge          bool   `json:"is_judge"`
	IsJudgePresident bool   `json:"is_judge_president"`
	Color            string `json:"color"`
	Icon             string `json:"icon"`
	AvailableSlots   int    `json:"available_slots"`
}
