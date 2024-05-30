package dto

import (
	"bff/structs"
	"time"
)

type PaymentOrderResponse struct {
	ID               int                        `json:"id"`
	OrganizationUnit DropdownSimple             `json:"organization_unit"`
	Supplier         DropdownSimple             `json:"supplier"`
	BankAccount      string                     `json:"bank_account"`
	DateOfPayment    time.Time                  `json:"date_of_payment"`
	DateOfOrder      *time.Time                 `json:"date_of_order"`
	IDOfStatement    *int                       `json:"id_of_statement"`
	SAPID            *string                    `json:"sap_id"`
	SourceOfFunding  string                     `json:"source_of_funding"`
	Description      string                     `json:"description"`
	DateOfSAP        *time.Time                 `json:"date_of_sap"`
	File             FileDropdownSimple         `json:"file"`
	Items            []PaymentOrderItemResponse `json:"items"`
	Amount           float64                    `json:"amount"`
	Status           string                     `json:"status"`
	CreatedAt        time.Time                  `json:"created_at"`
	UpdatedAt        time.Time                  `json:"updated_at"`
}

type PaymentOrderFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	SortByTitle        *string `json:"sort_by_title"`
	Status             *string `json:"status"`
	Search             *string `json:"search"`
	Year               *int    `json:"year"`
	SupplierID         *int    `json:"supplier_id"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Registred          *bool   `json:"registred"`
}

type ObligationsFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	OrganizationUnitID int     `json:"organization_unit_id"`
	SupplierID         int     `json:"supplier_id"`
	Search             *string `json:"search"`
	Type               *string `json:"type"`
}

type Obligation struct {
	InvoiceID                 *int         `json:"invoice_id"`
	AdditionalExpenseID       *int         `json:"additional_expense_id"`
	SalaryAdditionalExpenseID *int         `json:"salary_additional_expense_id"`
	SourceAccount             string       `json:"source_account"`
	InvoiceItems              []Obligation `json:"invoice_items"`
	Type                      string       `json:"type"`
	Title                     string       `json:"title"`
	TotalPrice                float64      `json:"total_price"`
	RemainPrice               float64      `json:"remain_price"`
	Status                    string       `json:"status"`
	CreatedAt                 time.Time    `json:"created_at"`
}

type GetObligationsResponseMS struct {
	Data  []Obligation `json:"data"`
	Total int          `json:"total"`
}

type GetPaymentOrderResponseMS struct {
	Data structs.PaymentOrder `json:"data"`
}

type GetPaymentOrderListResponseMS struct {
	Data  []structs.PaymentOrder `json:"data"`
	Total int                    `json:"total"`
}

type PaymentAdditionalExpensesListInputMS struct {
	Page                 *int    `json:"page"`
	Size                 *int    `json:"size"`
	PaymentOrderID       *int    `json:"payment_order_id"`
	PayingPaymentOrderID *int    `json:"paying_payment_order_id"`
	SubjectID            *int    `json:"subject_id"`
	OrganizationUnitID   *int    `json:"organization_unit_id"`
	Year                 *int    `json:"year"`
	Status               *string `json:"status"`
	Search               *string `json:"search"`
}

type GetPaymentAdditionalExpensesListResponseMS struct {
	Data  []structs.PaymentOrderItems `json:"data"`
	Total int                         `json:"total"`
}

type PaymentOrderItemResponse struct {
	ID                        int            `json:"id"`
	PaymentOrderID            int            `json:"payment_order_id"`
	InvoiceID                 *int           `json:"invoice_id"`
	AdditionalExpenseID       *int           `json:"additional_expense"`
	SalaryAdditionalExpenseID *int           `json:"salary_additional_expense"`
	Type                      string         `json:"type"`
	Account                   DropdownSimple `json:"account"`
	Amount                    float64        `json:"amount"`
	Title                     string         `json:"title"`
	CreatedAt                 time.Time      `json:"created_at"`
	UpdatedAt                 time.Time      `json:"updated_at"`
}
