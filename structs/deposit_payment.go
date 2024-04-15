package structs

import "time"

type DepositPayment struct {
	ID                        int        `json:"id"`
	Payer                     string     `json:"payer"`
	OrganizationUnitID        int        `json:"organization_unit_id"`
	CaseNumber                string     `json:"case_number"`
	PartyName                 string     `json:"party_name"`
	NumberOfBankStatement     string     `json:"number_of_bank_statement"`
	DateOfBankStatement       string     `json:"date_of_bank_statement"`
	AccountID                 int        `json:"account_id"`
	Amount                    float64    `json:"amount"`
	MainBankAccount           bool       `json:"main_bank_account"`
	DateOfTransferMainAccount *time.Time `json:"date_of_transfer_main_account"`
	FileID                    int        `json:"file_id"`
	CreatedAt                 time.Time  `json:"created_at"`
	UpdatedAt                 time.Time  `json:"updated_at"`
}
