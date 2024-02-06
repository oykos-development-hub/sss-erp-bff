package dto

import "bff/structs"

type GetFinanceActivityResponseMS struct {
	Data structs.ActivitiesItem `json:"data"`
}

type GetFinanceActivityListResponseMS struct {
	Data []structs.ActivitiesItem `json:"data"`
}

type GetFinanceActivityListInputMS struct {
	OrganizationUnitID *int    `json:"organization_unit_id"`
	SubProgramID       *int    `json:"sub_program_id"`
	Search             *string `json:"search"`
}

type ActivityResItem struct {
	ID               int            `json:"id"`
	SubProgram       DropdownSimple `json:"sub_program"`
	Program          DropdownSimple `json:"program"`
	OrganizationUnit DropdownSimple `json:"organization_unit"`
	Title            string         `json:"title"`
	Description      string         `json:"description"`
	Code             string         `json:"code"`
}
