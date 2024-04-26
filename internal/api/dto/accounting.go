package dto

import (
	"bff/structs"
	"time"
)

type ObligationForAccounting struct {
	InvoiceID *int      `json:"invoice_id"`
	SalaryID  *int      `json:"salary_id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type GetObligationsForAccountingResponseMS struct {
	Data  []ObligationForAccounting `json:"data"`
	Total int                       `json:"total"`
}

type GetAccountingOrderForObligations struct {
	Data *AccountingOrderForObligations `json:"data"`
}

type AccountingOrderForObligations struct {
	OrganizationUnitID int                                  `json:"organization_unit_id"`
	DateOfBooking      time.Time                            `json:"date_of_booking"`
	CreditAmount       float32                              `json:"credit_amount"`
	DebitAmount        float32                              `json:"debit_amount"`
	Items              []AccountingOrderItemsForObligations `json:"items"`
}

type AccountingOrderItemsForObligations struct {
	AccountID    int            `json:"account_id"`
	Title        string         `json:"title"`
	CreditAmount float32        `json:"credit_amount"`
	DebitAmount  float32        `json:"debit_amount"`
	Type         string         `json:"type"`
	SupplierID   int            `json:"supplier_id"`
	Invoice      DropdownSimple `json:"invoice"`
	Salary       DropdownSimple `json:"salary"`
}

type AccountingOrderForObligationsResponse struct {
	OrganizationUnit DropdownSimple                               `json:"organization_unit"`
	DateOfBooking    time.Time                                    `json:"date_of_booking"`
	CreditAmount     float32                                      `json:"credit_amount"`
	DebitAmount      float32                                      `json:"debit_amount"`
	Items            []AccountingOrderItemsForObligationsResponse `json:"items"`
}

type AccountingOrderItemsForObligationsResponse struct {
	Account      DropdownSimple `json:"account"`
	Title        string         `json:"title"`
	CreditAmount float32        `json:"credit_amount"`
	DebitAmount  float32        `json:"debit_amount"`
	Type         string         `json:"type"`
	Invoice      DropdownSimple `json:"invoice"`
	Salary       DropdownSimple `json:"salary"`
}

type AccountingEntryResponse struct {
	ID               int                           `json:"id"`
	OrganizationUnit DropdownSimple                `json:"organization_unit"`
	Supplier         DropdownSimple                `json:"supplier"`
	BankAccount      string                        `json:"bank_account"`
	DateOfPayment    time.Time                     `json:"date_of_payment"`
	DateOfOrder      *time.Time                    `json:"date_of_order"`
	IDOfStatement    *string                       `json:"id_of_statement"`
	SAPID            *string                       `json:"sap_id"`
	SourceOfFunding  string                        `json:"source_of_funding"`
	Description      string                        `json:"description"`
	DateOfSAP        *time.Time                    `json:"date_of_sap"`
	File             FileDropdownSimple            `json:"file"`
	Items            []AccountingEntryItemResponse `json:"items"`
	Amount           float64                       `json:"amount"`
	Status           string                        `json:"status"`
	CreatedAt        time.Time                     `json:"created_at"`
	UpdatedAt        time.Time                     `json:"updated_at"`
}

type AccountingEntryItemResponse struct {
	ID               int                `json:"id"`
	OrganizationUnit DropdownSimple     `json:"organization_unit"`
	Supplier         DropdownSimple     `json:"supplier"`
	BankAccount      string             `json:"bank_account"`
	DateOfPayment    time.Time          `json:"date_of_payment"`
	DateOfOrder      *time.Time         `json:"date_of_order"`
	IDOfStatement    *string            `json:"id_of_statement"`
	SAPID            *string            `json:"sap_id"`
	SourceOfFunding  string             `json:"source_of_funding"`
	Description      string             `json:"description"`
	DateOfSAP        *time.Time         `json:"date_of_sap"`
	File             FileDropdownSimple `json:"file"`
	Amount           float64            `json:"amount"`
	Status           string             `json:"status"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

type AccountingEntryFilter struct {
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

type GetAccountingEntryResponseMS struct {
	Data structs.AccountingEntry `json:"data"`
}

type GetAccountingEntryListResponseMS struct {
	Data  []structs.AccountingEntry `json:"data"`
	Total int                       `json:"total"`
}
