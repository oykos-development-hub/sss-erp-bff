package structs

import "time"

type DepositPaymentOrder struct {
	ID                          int                                `json:"id"`
	OrganizationUnitID          int                                `json:"organization_unit_id"`
	CaseNumber                  string                             `json:"case_number"`
	SupplierID                  int                                `json:"supplier_id"`
	NetAmount                   float64                            `json:"net_amount"`
	BankAccount                 string                             `json:"bank_account"`
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
	ID                 int       `json:"id"`
	Title              string    `json:"title"`
	AccountID          int       `json:"account_id"`
	Price              float32   `json:"price"`
	SubjectID          int       `json:"subject_id"`
	BankAccount        string    `json:"bank_account"`
	PaymentOrderID     int       `json:"payment_order_id"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	Status             string    `json:"status"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
