package dto

import "bff/structs"

type GetProcurementOULimitListInputMS struct {
	ItemID *int `json:"procurement_id" validate:"omitempty"`
}

type GetProcurementOULimitListResponseMS struct {
	Data []*structs.PublicProcurementLimit `json:"data"`
}

type GetProcurementOULimitResponseMS struct {
	Data structs.PublicProcurementLimit `json:"data"`
}

type ProcurementOULimitResponseItem struct {
	ID                int            `json:"id"`
	PublicProcurement DropdownSimple `json:"public_procurement"`
	OrganizationUnit  DropdownSimple `json:"organization_unit"`
	Limit             int            `json:"limit"`
}
