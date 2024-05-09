package dto

import (
	"bff/structs"
	"time"
)

type ObligationForAccounting struct {
	ID         int            `json:"id"`
	InvoiceID  *int           `json:"invoice_id"`
	SalaryID   *int           `json:"salary_id"`
	Type       string         `json:"type"`
	SupplierID *int           `json:"supplier_id"`
	Supplier   DropdownSimple `json:"supplier"`
	Date       time.Time      `json:"date"`
	Title      string         `json:"title"`
	Price      float64        `json:"price"`
	Status     string         `json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
}

type PaymentOrdersForAccounting struct {
	ID             int            `json:"id"`
	PaymentOrderID int            `json:"payment_order_id"`
	SupplierID     *int           `json:"supplier_id"`
	Supplier       DropdownSimple `json:"supplier"`
	Date           time.Time      `json:"date"`
	Title          string         `json:"title"`
	Price          float64        `json:"price"`
	CreatedAt      time.Time      `json:"created_at"`
}

type GetObligationsForAccountingResponseMS struct {
	Data  []ObligationForAccounting `json:"data"`
	Total int                       `json:"total"`
}

type GetPaymentOrdersForAccountingResponseMS struct {
	Data  []PaymentOrdersForAccounting `json:"data"`
	Total int                          `json:"total"`
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
	AccountID             int            `json:"account_id"`
	Title                 string         `json:"title"`
	CreditAmount          float32        `json:"credit_amount"`
	DebitAmount           float32        `json:"debit_amount"`
	Type                  string         `json:"type"`
	SupplierID            int            `json:"supplier_id"`
	Invoice               DropdownSimple `json:"invoice"`
	Salary                DropdownSimple `json:"salary"`
	PaymentOrder          DropdownSimple `json:"payment_order"`
	EnforcedPayment       DropdownSimple `json:"enforced_payment"`
	ReturnEnforcedPayment DropdownSimple `json:"return_enforced_payment"`
}

type AccountingOrderForObligationsResponse struct {
	OrganizationUnit DropdownSimple                               `json:"organization_unit"`
	DateOfBooking    time.Time                                    `json:"date_of_booking"`
	CreditAmount     float32                                      `json:"credit_amount"`
	DebitAmount      float32                                      `json:"debit_amount"`
	Items            []AccountingOrderItemsForObligationsResponse `json:"items"`
}

type AccountingOrderItemsForObligationsResponse struct {
	ID                    int            `json:"id"`
	Account               DropdownSimple `json:"account"`
	Title                 string         `json:"title"`
	CreditAmount          float32        `json:"credit_amount"`
	DebitAmount           float32        `json:"debit_amount"`
	Type                  string         `json:"type"`
	Invoice               DropdownSimple `json:"invoice"`
	Salary                DropdownSimple `json:"salary"`
	PaymentOrder          DropdownSimple `json:"payment_order"`
	EnforcedPayment       DropdownSimple `json:"enforced_payment"`
	ReturnEnforcedPayment DropdownSimple `json:"return_enforced_payment"`
}

type AccountingEntryResponse struct {
	ID               int                           `json:"id"`
	Title            string                        `json:"title"`
	IDOfEntry        int                           `json:"id_of_entry"`
	OrganizationUnit DropdownSimple                `json:"organization_unit"`
	DateOfBooking    time.Time                     `json:"date_of_booking"`
	CreditAmount     float64                       `json:"credit_amount"`
	DebitAmount      float64                       `json:"debit_amount"`
	Items            []AccountingEntryItemResponse `json:"items"`
	CreatedAt        time.Time                     `json:"created_at"`
	UpdatedAt        time.Time                     `json:"updated_at"`
}

type AccountingEntryItemResponse struct {
	ID                    int            `json:"id"`
	Title                 string         `json:"title"`
	EntryID               int            `json:"entry_id"`
	Account               DropdownSimple `json:"account"`
	CreditAmount          float64        `json:"credit_amount"`
	DebitAmount           float64        `json:"debit_amount"`
	Invoice               DropdownSimple `json:"invoice"`
	Salary                DropdownSimple `json:"salary"`
	PaymentOrder          DropdownSimple `json:"payment_order"`
	EnforcedPayment       DropdownSimple `json:"enforced_payment"`
	ReturnEnforcedPayment DropdownSimple `json:"return_enforced_payment"`
	Type                  string         `json:"type"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
}

type AccountingEntryFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Type               *string `json:"type"`
}

type GetAccountingEntryResponseMS struct {
	Data structs.AccountingEntry `json:"data"`
}

type GetAccountingEntryListResponseMS struct {
	Data  []structs.AccountingEntry `json:"data"`
	Total int                       `json:"total"`
}
