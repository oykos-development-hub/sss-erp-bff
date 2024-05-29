package structs

import "time"

type EnforcedPayment struct {
	ID                 int                    `json:"id"`
	OrganizationUnitID int                    `json:"organization_unit_id"`
	SupplierID         int                    `json:"supplier_id"`
	BankAccount        string                 `json:"bank_account"`
	DateOfPayment      time.Time              `json:"date_of_payment"`
	IDOfStatement      *int                   `json:"id_of_statement"`
	SAPID              *string                `json:"sap_id"`
	DateOfSAP          *time.Time             `json:"date_of_sap"`
	DateOfOrder        *time.Time             `json:"date_of_order"`
	FileID             *int                   `json:"file_id"`
	ReturnFileID       *int                   `json:"return_file_id"`
	ReturnDate         *time.Time             `json:"return_date"`
	ReturnAmount       *float64               `json:"return_amount"`
	Items              []EnforcedPaymentItems `json:"items"`
	Amount             float64                `json:"amount"`
	AmountForLawyer    float64                `json:"amount_for_lawyer"`
	AmountForAgent     float64                `json:"amount_for_agent"`
	Description        string                 `json:"description"`
	Status             string                 `json:"status"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
}

type EnforcedPaymentItems struct {
	ID             int       `json:"id"`
	PaymentOrderID int       `json:"payment_order_id"`
	InvoiceID      *int      `json:"invoice_id"`
	AccountID      int       `json:"account_id"`
	Amount         float64   `json:"amount"`
	Title          string    `json:"title"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
