package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type AccountingOrderForObligationsData struct {
	InvoiceID               []int     `json:"invoice_id"`
	SalaryID                []int     `json:"salary_id"`
	PaymentOrderID          []int     `json:"payment_order_id"`
	EnforcedPaymentID       []int     `json:"enforced_payment_id"`
	ReturnEnforcedPaymentID []int     `json:"return_enforced_payment_id"`
	DateOfBooking           time.Time `json:"date_of_booking"`
	OrganizationUnitID      int       `json:"organization_unit_id"`
}

type AccountingEntry struct {
	ID                 int                    `json:"id"`
	Title              string                 `json:"title"`
	Type               string                 `json:"type"`
	IDOfEntry          int                    `json:"id_of_entry"`
	OrganizationUnitID int                    `json:"organization_unit_id"`
	DateOfBooking      time.Time              `json:"date_of_booking"`
	CreditAmount       decimal.Decimal        `json:"credit_amount"`
	DebitAmount        decimal.Decimal        `json:"debit_amount"`
	Items              []AccountingEntryItems `json:"items"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
}

type AccountingEntryItems struct {
	ID                      int             `json:"id"`
	Title                   string          `json:"title"`
	EntryID                 int             `json:"entry_id"`
	AccountID               int             `json:"account_id"`
	CreditAmount            decimal.Decimal `json:"credit_amount"`
	DebitAmount             decimal.Decimal `json:"debit_amount"`
	InvoiceID               *int            `json:"invoice_id"`
	SalaryID                *int            `json:"salary_id"`
	SupplierID              int             `json:"supplier_id"`
	EnforcedPaymentID       *int            `json:"enforced_payment_id"`
	ReturnEnforcedPaymentID *int            `json:"return_enforced_payment_id"`
	PaymentOrderID          *int            `json:"payment_order_id"`
	Type                    string          `json:"type"`
	Date                    string          `json:"date"`
	EntryNumber             string          `json:"entry_number"`
	EntryDate               time.Time       `json:"entry_date"`
	CreatedAt               time.Time       `json:"created_at"`
	UpdatedAt               time.Time       `json:"updated_at"`
}

type AnalyticalCard struct {
	InitialState            decimal.Decimal       `json:"initial_state"`
	SumCreditAmount         decimal.Decimal       `json:"sum_credit_amount"`
	SumDebitAmount          decimal.Decimal       `json:"sum_debit_amount"`
	SumCreditAmountInPeriod decimal.Decimal       `json:"sum_credit_amount_in_period"`
	SumDebitAmountInPeriod  decimal.Decimal       `json:"sum_debit_amount_in_period"`
	SupplierID              int                   `json:"supplier_id"`
	OrganizationUnitID      int                   `json:"organization_unit_id"`
	DateOfStart             *time.Time            `json:"date_of_start"`
	DateOfEnd               *time.Time            `json:"date_of_end"`
	Items                   []AnalyticalCardItems `json:"items"`
}

type AnalyticalCardItems struct {
	ID             int             `json:"id"`
	Title          string          `json:"title"`
	Type           string          `json:"type"`
	IDOfEntry      int             `json:"id_of_entry"`
	CreditAmount   decimal.Decimal `json:"credit_amount"`
	DebitAmount    decimal.Decimal `json:"debit_amount"`
	Balance        decimal.Decimal `json:"balance"`
	DateOfBooking  string          `json:"date_of_booking"`
	Date           string          `json:"date"`
	DocumentNumber string          `json:"document_number"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
