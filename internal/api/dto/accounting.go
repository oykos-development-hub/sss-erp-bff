package dto

import (
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
