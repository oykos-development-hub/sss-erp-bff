package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type DepositPaymentOrder struct {
	ID                          int                                `json:"id"`
	OrganizationUnitID          int                                `json:"organization_unit_id"`
	CaseNumber                  string                             `json:"case_number"`
	SupplierID                  int                                `json:"supplier_id"`
	SubjectTypeID               int                                `json:"subject_type_id"`
	NetAmount                   decimal.Decimal                    `json:"net_amount"`
	BankAccount                 string                             `json:"bank_account"`
	SourceBankAccount           string                             `json:"source_bank_account"`
	DateOfPayment               time.Time                          `json:"date_of_payment"`
	DateOfStatement             *time.Time                         `json:"date_of_statement"`
	IDOfStatement               *string                            `json:"id_of_statement"`
	Status                      string                             `json:"status"`
	MunicipalityID              *int                               `json:"municipality_id"`
	TaxAuthorityCodebookID      *int                               `json:"tax_authority_codebook_id"`
	AdditionalExpenses          []DepositPaymentAdditionalExpenses `json:"additional_expenses"`
	AdditionalExpensesForPaying []DepositPaymentAdditionalExpenses `json:"additional_expenses_for_paying"`
	FileID                      int                                `json:"file_id"`
	CreatedAt                   time.Time                          `json:"created_at"`
	UpdatedAt                   time.Time                          `json:"updated_at"`
}

type DepositPaymentAdditionalExpenses struct {
	ID                 int             `json:"id"`
	Title              string          `json:"title"`
	CaseNumber         string          `json:"case_number"`
	AccountID          int             `json:"account_id"`
	Price              decimal.Decimal `json:"price"`
	SubjectID          int             `json:"subject_id"`
	BankAccount        string          `json:"bank_account"`
	SourceBankAccount  string          `json:"source_bank_account"`
	PaymentOrderID     int             `json:"payment_order_id"`
	OrganizationUnitID int             `json:"organization_unit_id"`
	Status             string          `json:"status"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}
