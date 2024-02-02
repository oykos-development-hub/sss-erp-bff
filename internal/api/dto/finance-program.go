package dto

import "bff/structs"

type GetFinanceProgramResponseMS struct {
	Data structs.ProgramItem `json:"data"`
}

type GetFinanceProgramListResponseMS struct {
	Data []structs.ProgramItem `json:"data"`
}

type GetFinanceProgramListInputMS struct {
	IsProgram *bool   `json:"is_program"`
	Search    *string `json:"search"`
}

type ProgramResItem struct {
	ID          int             `json:"id"`
	Parent      *DropdownSimple `json:"parent"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Code        string          `json:"code"`
}
