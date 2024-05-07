package structs

import "time"

type AccountingOrderForObligationsData struct {
	InvoiceID          []int     `json:"invoice_id"`
	SalaryID           []int     `json:"salary_id"`
	PaymentOrderID     []int     `json:"payment_order_id"`
	DateOfBooking      time.Time `json:"date_of_booking"`
	OrganizationUnitID int       `json:"organization_unit_id"`
}

type AccountingEntry struct {
	ID                 int                    `json:"id"`
	Title              string                 `json:"title"`
	IDOfEntry          int                    `json:"id_of_entry"`
	OrganizationUnitID int                    `json:"organization_unit_id"`
	DateOfBooking      time.Time              `json:"date_of_booking"`
	CreditAmount       float64                `json:"credit_amount"`
	DebitAmount        float64                `json:"debit_amount"`
	Items              []AccountingEntryItems `json:"items"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
}

type AccountingEntryItems struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	EntryID        int       `json:"entry_id"`
	AccountID      int       `json:"account_id"`
	CreditAmount   float64   `json:"credit_amount"`
	DebitAmount    float64   `json:"debit_amount"`
	InvoiceID      *int      `json:"invoice_id"`
	SalaryID       *int      `json:"salary_id"`
	PaymentOrderID *int      `json:"payment_order_id"`
	Type           string    `json:"type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
