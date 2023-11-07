package dto

import (
	"bff/structs"
)

type GetProcurementContractResponseMS struct {
	Data structs.PublicProcurementContract `json:"data"`
}

type GetProcurementContractsInput struct {
	Page          *int `json:"page"`
	Size          *int `json:"size"`
	ProcurementID *int `json:"procurement_id"`
	SupplierID    *int `json:"supplier_id"`
}

type GetProcurementContractListResponseMS struct {
	Data  []*structs.PublicProcurementContract `json:"data"`
	Total int                                  `json:"total"`
}

type ProcurementContractResponseItem struct {
	Id                  int              `json:"id"`
	PublicProcurementId int              `json:"public_procurement_id"`
	SupplierId          int              `json:"supplier_id"`
	SerialNumber        string           `json:"serial_number"`
	DateOfSigning       string           `json:"date_of_signing"`
	DateOfExpiry        *string          `json:"date_of_expiry"`
	NetValue            *float32         `json:"net_value"`
	GrossValue          *float32         `json:"gross_value"`
	VatValue            *float32         `json:"vat_value"`
	File                []DropdownSimple `json:"file"`
	CreatedAt           string           `json:"created_at"`
	UpdatedAt           string           `json:"updated_at"`
	PublicProcurement   DropdownSimple   `json:"public_procurement"`
	Supplier            DropdownSimple   `json:"supplier"`
}
