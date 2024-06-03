package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type DepositPayment struct {
	ID                        int             `json:"id"`
	Payer                     string          `json:"payer"`
	OrganizationUnitID        int             `json:"organization_unit_id"`
	CaseNumber                string          `json:"case_number"`
	PartyName                 string          `json:"party_name"`
	NumberOfBankStatement     string          `json:"number_of_bank_statement"`
	DateOfBankStatement       string          `json:"date_of_bank_statement"`
	AccountID                 int             `json:"account_id"`
	Amount                    decimal.Decimal `json:"amount"`
	MainBankAccount           bool            `json:"main_bank_account"`
	CurrentBankAccount        string          `json:"current_bank_account"`
	DateOfTransferMainAccount *time.Time      `json:"date_of_transfer_main_account"`
	FileID                    int             `json:"file_id"`
	Status                    string          `json:"status"`
	CreatedAt                 time.Time       `json:"created_at"`
	UpdatedAt                 time.Time       `json:"updated_at"`
}
