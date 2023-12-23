package dto

import (
	"bff/structs"
)

type GetProcurementContractResponseMS struct {
	Data structs.PublicProcurementContract `json:"data"`
}

type GetProcurementContractsInput struct {
	Page                *int    `json:"page"`
	Size                *int    `json:"size"`
	ProcurementID       *int    `json:"procurement_id"`
	SupplierID          *int    `json:"supplier_id"`
	Year                *string `json:"year"`
	SortByDateOfExpiry  *string `json:"sort_by_date_of_expiry"`
	SortByDateOfSigning *string `json:"sort_by_date_of_signing"`
	SortByGrossValue    *string `json:"sort_by_gross_value"`
	SortBySupplierID    *string `json:"sort_by_supplier_id"`
	SortBySerialNumber  *string `json:"sort_by_serial_number"`
}

type GetProcurementContractListResponseMS struct {
	Data  []*structs.PublicProcurementContract `json:"data"`
	Total int                                  `json:"total"`
}

type ProcurementContractResponseItem struct {
	ID                  int                  `json:"id"`
	PublicProcurementID int                  `json:"public_procurement_id"`
	SupplierID          int                  `json:"supplier_id"`
	SerialNumber        string               `json:"serial_number"`
	DateOfSigning       string               `json:"date_of_signing"`
	DateOfExpiry        *string              `json:"date_of_expiry"`
	NetValue            *float32             `json:"net_value"`
	GrossValue          *float32             `json:"gross_value"`
	VatValue            *float32             `json:"vat_value"`
	File                []FileDropdownSimple `json:"file"`
	CreatedAt           string               `json:"created_at"`
	UpdatedAt           string               `json:"updated_at"`
	PublicProcurement   DropdownSimple       `json:"public_procurement"`
	Supplier            DropdownSimple       `json:"supplier"`
	DaysUntilExpiry     int                  `json:"days_until_expiry"`
}
