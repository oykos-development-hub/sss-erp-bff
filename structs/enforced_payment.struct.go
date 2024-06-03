package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

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
	ReturnAmount       *decimal.Decimal       `json:"return_amount"`
	Items              []EnforcedPaymentItems `json:"items"`
	Amount             decimal.Decimal        `json:"amount"`
	AmountForLawyer    decimal.Decimal        `json:"amount_for_lawyer"`
	AmountForAgent     decimal.Decimal        `json:"amount_for_agent"`
	AmountForBank      decimal.Decimal        `json:"amount_for_bank"`
	AgentID            int                    `json:"agent_id"`
	ExecutionNumber    string                 `json:"execution_number"`
	Description        string                 `json:"description"`
	Status             string                 `json:"status"`
	CreatedAt          time.Time              `json:"created_at"`
	UpdatedAt          time.Time              `json:"updated_at"`
}

type EnforcedPaymentItems struct {
	ID             int             `json:"id"`
	PaymentOrderID int             `json:"payment_order_id"`
	InvoiceID      *int            `json:"invoice_id"`
	AccountID      int             `json:"account_id"`
	Amount         decimal.Decimal `json:"amount"`
	Title          string          `json:"title"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
