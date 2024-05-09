package dto

import (
	"bff/structs"
	"time"
)

type EnforcedPaymentResponse struct {
	ID               int                           `json:"id"`
	OrganizationUnit DropdownSimple                `json:"organization_unit"`
	Supplier         DropdownSimple                `json:"supplier"`
	BankAccount      string                        `json:"bank_account"`
	DateOfPayment    time.Time                     `json:"date_of_payment"`
	DateOfOrder      *time.Time                    `json:"date_of_order"`
	ReturnFile       FileDropdownSimple            `json:"return_file"`
	ReturnDate       *time.Time                    `json:"return_date"`
	ReturnAmount     *float64                      `json:"return_amount"`
	IDOfStatement    *string                       `json:"id_of_statement"`
	SAPID            *string                       `json:"sap_id"`
	Description      string                        `json:"description"`
	DateOfSAP        *time.Time                    `json:"date_of_sap"`
	File             FileDropdownSimple            `json:"file"`
	Items            []EnforcedPaymentItemResponse `json:"items"`
	Amount           float64                       `json:"amount"`
	AmountForLawyer  float64                       `json:"amount_for_lawyer"`
	AmountForAgent   float64                       `json:"amount_for_agent"`
	Status           string                        `json:"status"`
	CreatedAt        time.Time                     `json:"created_at"`
	UpdatedAt        time.Time                     `json:"updated_at"`
}

type EnforcedPaymentFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	SortByTitle        *string `json:"sort_by_title"`
	Status             *string `json:"status"`
	Search             *string `json:"search"`
	Year               *int    `json:"year"`
	SupplierID         *int    `json:"supplier_id"`
	Registred          *bool   `json:"registred"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
}

type GetEnforcedPaymentResponseMS struct {
	Data structs.EnforcedPayment `json:"data"`
}

type GetEnforcedPaymentListResponseMS struct {
	Data  []structs.EnforcedPayment `json:"data"`
	Total int                       `json:"total"`
}

type EnforcedPaymentItemResponse struct {
	ID             int            `json:"id"`
	PaymentOrderID int            `json:"payment_order_id"`
	InvoiceID      *int           `json:"invoice_id"`
	Account        DropdownSimple `json:"account"`
	Amount         float64        `json:"amount"`
	Title          string         `json:"title"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}
