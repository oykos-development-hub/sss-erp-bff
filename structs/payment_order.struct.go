package structs

import "time"

type PaymentOrder struct {
	ID                 int                 `json:"id"`
	OrganizationUnitID int                 `json:"organization_unit_id"`
	SupplierID         int                 `json:"supplier_id"`
	BankAccount        string              `json:"bank_account"`
	DateOfPayment      time.Time           `json:"date_of_payment"`
	IDOfStatement      *int                `json:"id_of_statement"`
	SAPID              *string             `json:"sap_id"`
	SourceOfFunding    string              `json:"source_of_funding"`
	DateOfSAP          *time.Time          `json:"date_of_sap"`
	Registred          bool                `json:"registred"`
	DateOfOrder        *time.Time          `json:"date_of_order"`
	FileID             *int                `json:"file_id"`
	Items              []PaymentOrderItems `json:"items"`
	Amount             float64             `json:"amount"`
	Description        string              `json:"description"`
	Status             string              `json:"status"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
}

type PaymentOrderItems struct {
	ID                        int       `json:"id"`
	PaymentOrderID            int       `json:"payment_order_id"`
	InvoiceID                 *int      `json:"invoice_id"`
	AdditionalExpenseID       *int      `json:"additional_expense_id"`
	SalaryAdditionalExpenseID *int      `json:"salary_additional_expense_id"`
	Type                      string    `json:"type"`
	AccountID                 int       `json:"account_id"`
	SourceAccountID           int       `json:"source_account_id"`
	Amount                    float64   `json:"amount"`
	Title                     string    `json:"title"`
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
}
