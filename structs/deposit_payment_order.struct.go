package structs

import "time"

type DepositPaymentOrder struct {
	ID                 int        `json:"id"`
	OrganizationUnitID int        `json:"organization_unit_id"`
	CaseNumber         string     `json:"case_number"`
	SupplierID         int        `json:"supplier_id"`
	NetAmount          float64    `json:"net_amount"`
	BankAccount        string     `json:"bank_account"`
	DateOfPayment      time.Time  `json:"date_of_payment"`
	DateOfStatement    *time.Time `json:"date_of_statement"`
	IDOfStatement      *string    `json:"id_of_statement"`
	Status             string     `json:"status"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}
