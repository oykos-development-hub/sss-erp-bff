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
	CreditAmount       float64                              `json:"credit_amount"`
	DebitAmount        float64                              `json:"debit_amount"`
	Items              []AccountingOrderItemsForObligations `json:"items"`
}

type AccountingOrderItemsForObligations struct {
	AccountID             int            `json:"account_id"`
	Title                 string         `json:"title"`
	CreditAmount          float64        `json:"credit_amount"`
	DebitAmount           float64        `json:"debit_amount"`
	Type                  string         `json:"type"`
	SupplierID            int            `json:"supplier_id"`
	Date                  time.Time      `json:"date"`
	Invoice               DropdownSimple `json:"invoice"`
	Salary                DropdownSimple `json:"salary"`
	PaymentOrder          DropdownSimple `json:"payment_order"`
	EnforcedPayment       DropdownSimple `json:"enforced_payment"`
	ReturnEnforcedPayment DropdownSimple `json:"return_enforced_payment"`
}

type AccountingOrderForObligationsResponse struct {
	OrganizationUnit DropdownSimple                               `json:"organization_unit"`
	DateOfBooking    time.Time                                    `json:"date_of_booking"`
	CreditAmount     float64                                      `json:"credit_amount"`
	DebitAmount      float64                                      `json:"debit_amount"`
	Items            []AccountingOrderItemsForObligationsResponse `json:"items"`
}

type AccountingOrderItemsForObligationsResponse struct {
	ID                    int            `json:"id"`
	Account               DropdownSimple `json:"account"`
	Title                 string         `json:"title"`
	CreditAmount          float64        `json:"credit_amount"`
	DebitAmount           float64        `json:"debit_amount"`
	Type                  string         `json:"type"`
	SupplierID            int            `json:"supplier_id"`
	Date                  time.Time      `json:"date"`
	Invoice               DropdownSimple `json:"invoice"`
	Salary                DropdownSimple `json:"salary"`
	PaymentOrder          DropdownSimple `json:"payment_order"`
	EnforcedPayment       DropdownSimple `json:"enforced_payment"`
	ReturnEnforcedPayment DropdownSimple `json:"return_enforced_payment"`
}

type AccountingEntryResponse struct {
	ID                int                               `json:"id"`
	Title             string                            `json:"title"`
	Type              string                            `json:"type"`
	IDOfEntry         int                               `json:"id_of_entry"`
	FormatedIDOfEntry string                            `json:"formated_id_of_entry"`
	OrganizationUnit  OrganizationUnitsOverviewResponse `json:"organization_unit"`
	DateOfBooking     time.Time                         `json:"date_of_booking"`
	CreditAmount      float64                           `json:"credit_amount"`
	DebitAmount       float64                           `json:"debit_amount"`
	Items             []AccountingEntryItemResponse     `json:"items"`
	CreatedAt         time.Time                         `json:"created_at"`
	UpdatedAt         time.Time                         `json:"updated_at"`
}

type AccountingEntryItemResponse struct {
	ID                    int            `json:"id"`
	Title                 string         `json:"title"`
	EntryID               int            `json:"entry_id"`
	EntryNumber           string         `json:"entry_number"`
	EntryDate             time.Time      `json:"entry_date"`
	Account               DropdownSimple `json:"account"`
	CreditAmount          float64        `json:"credit_amount"`
	DebitAmount           float64        `json:"debit_amount"`
	Invoice               DropdownSimple `json:"invoice"`
	Salary                DropdownSimple `json:"salary"`
	PaymentOrder          DropdownSimple `json:"payment_order"`
	EnforcedPayment       DropdownSimple `json:"enforced_payment"`
	ReturnEnforcedPayment DropdownSimple `json:"return_enforced_payment"`
	Supplier              DropdownSimple `json:"supplier"`
	Date                  string         `json:"date"`
	Type                  string         `json:"type"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
}

type AccountingEntryFilter struct {
	Page               *int       `json:"page"`
	Size               *int       `json:"size"`
	OrganizationUnitID *int       `json:"organization_unit_id"`
	Type               *string    `json:"type"`
	DateOfStart        *time.Time `json:"date_of_start"`
	DateOfEnd          *time.Time `json:"date_of_end"`
	SortForReport      *bool      `json:"sort_for_report"`
}

type GetAccountingEntryResponseMS struct {
	Data structs.AccountingEntry `json:"data"`
}

type GetAccountingEntryListResponseMS struct {
	Data  []structs.AccountingEntry `json:"data"`
	Total int                       `json:"total"`
}

type AnalyticalCardFilter struct {
	AccountID          []int      `json:"account_id"`
	SupplierID         *int       `json:"supplier_id"`
	OrganizationUnitID int        `json:"organization_unit_id"`
	DateOfStart        *time.Time `json:"date_of_start"`
	DateOfEnd          *time.Time `json:"date_of_end"`
	DateOfStartBooking *time.Time `json:"date_of_start_booking"`
	DateOfEndBooking   *time.Time `json:"date_of_end_booking"`
}

type GetAnalyticalCardListResponseMS struct {
	Data []structs.AnalyticalCard `json:"data"`
}

type AnalyticalCardDTO struct {
	InitialState            float64                  `json:"initial_state"`
	SumCreditAmount         float64                  `json:"sum_credit_amount"`
	SumDebitAmount          float64                  `json:"sum_debit_amount"`
	SumCreditAmountInPeriod float64                  `json:"sum_credit_amount_in_period"`
	SumDebitAmountInPeriod  float64                  `json:"sum_debit_amount_in_period"`
	Supplier                DropdownSimple           `json:"supplier"`
	OrganizationUnit        DropdownSimple           `json:"organization_unit"`
	DateOfStart             *time.Time               `json:"date_of_start"`
	DateOfEnd               *time.Time               `json:"date_of_end"`
	Items                   []AnalyticalCardItemsDTO `json:"items"`
}

type AnalyticalCardItemsDTO struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	Type              string    `json:"type"`
	FormatedIDOfEntry string    `json:"formated_id_of_entry"`
	CreditAmount      float64   `json:"credit_amount"`
	DebitAmount       float64   `json:"debit_amount"`
	Balance           float64   `json:"balance"`
	DateOfBooking     string    `json:"date_of_booking"`
	Date              string    `json:"date"`
	DocumentNumber    string    `json:"document_number"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
