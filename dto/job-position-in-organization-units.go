package dto

import "bff/structs"

type GetJobPositionInOrganizationUnitsResponseMS struct {
	Data structs.JobPositionsInOrganizationUnits `json:"data"`
}

type GetJobPositionInOrganizationUnitsInput struct {
	Page               *int `json:"page"`
	Size               *int `json:"page_size"`
	OrganizationUnitID *int `json:"organization_unit_id"`
	JobPositionID      *int `json:"job_position_id"`
	SystematizationID  *int `json:"systematization_id"`
}

type GetJobPositionsInOrganizationUnitsResponseMS struct {
	Data  []structs.JobPositionsInOrganizationUnits `json:"data"`
	Total int                                       `json:"total"`
}
