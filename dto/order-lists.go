package dto

import "bff/structs"

type GetOrderListsResponseMS struct {
	Data  []structs.OrderListItem `json:"data"`
	Total int                     `json:"total"`
}

type GetOrderListResponseMS struct {
	Data structs.OrderListItem `json:"data"`
}

type GetOrderListInput struct {
	Page                *int    `json:"page"`
	Size                *int    `json:"size"`
	SupplierID          *int    `json:"supplier_id"`
	Search              *string `json:"search"`
	Status              *string `json:"status"`
	PublicProcurementID *int    `json:"public_procurement_id"`
	OrganizationUnitId  *int    `json:"organization_unit_id"`
}

type OrderListOverviewResponse struct {
	Id                  int                                    `json:"id"`
	DateOrder           string                                 `json:"date_order" validate:"required"`
	TotalBruto          float32                                `json:"total_bruto"`
	TotalNeto           float32                                `json:"total_neto"`
	PublicProcurementID int                                    `json:"public_procurement_id"`
	SupplierID          int                                    `json:"supplier_id"`
	Status              string                                 `json:"status"`
	DateSystem          *string                                `json:"date_system"`
	InvoiceDate         *string                                `json:"invoice_date"`
	InvoiceNumber       string                                 `json:"invoice_number"`
	OrganizationUnitID  int                                    `json:"organization_unit_id"`
	OfficeID            int                                    `json:"office_id"`
	RecipientUserID     *int                                   `json:"recipient_user_id"`
	Description         *string                                `json:"description"`
	CreatedAt           string                                 `json:"created_at"`
	UpdatedAt           string                                 `json:"updated_at"`
	PublicProcurement   *DropdownSimple                        `json:"public_procurement"`
	Supplier            *DropdownSimple                        `json:"supplier"`
	RecipientUser       *DropdownSimple                        `json:"recipient_user"`
	Office              *DropdownSimple                        `json:"office"`
	Articles            *[]DropdownProcurementAvailableArticle `json:"articles"`
}
