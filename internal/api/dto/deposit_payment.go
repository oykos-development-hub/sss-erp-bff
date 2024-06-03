package dto

import (
	"bff/structs"
	"time"

	"github.com/shopspring/decimal"
)

type DepositPaymentResponse struct {
	ID                        int                `json:"id"`
	Payer                     string             `json:"payer"`
	OrganizationUnit          DropdownSimple     `json:"organization_unit"`
	CaseNumber                string             `json:"case_number"`
	PartyName                 string             `json:"party_name"`
	NumberOfBankStatement     string             `json:"number_of_bank_statement"`
	DateOfBankStatement       string             `json:"date_of_bank_statement"`
	Account                   DropdownSimple     `json:"account"`
	Amount                    decimal.Decimal    `json:"amount"`
	MainBankAccount           bool               `json:"main_bank_account"`
	CurrentBankAccount        string             `json:"current_bank_account"`
	Status                    string             `json:"status"`
	DateOfTransferMainAccount *time.Time         `json:"date_of_transfer_main_account"`
	File                      FileDropdownSimple `json:"file"`
	CreatedAt                 time.Time          `json:"created_at"`
	UpdatedAt                 time.Time          `json:"updated_at"`
}

type DepositPaymentFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	SortByTitle        *string `json:"sort_by_title"`
	Status             *string `json:"status"`
	Search             *string `json:"search"`
	CaseuNumber        *string `json:"case_number"`
	BankAccount        *string `json:"bank_account"`
	SourceBankAccount  *string `json:"source_bank_account"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
}

type DepositInitialStateFilter struct {
	BankAccount             *string   `json:"bank_account"`
	OrganizationUnitID      *int      `json:"organization_unit_id"`
	Date                    time.Time `json:"date"`
	TransitionalBankAccount *bool     `json:"transitional_bank_account"`
}

type GetDepositPaymentResponseMS struct {
	Data structs.DepositPayment `json:"data"`
}

type GetDepositPaymentListResponseMS struct {
	Data  []structs.DepositPayment `json:"data"`
	Total int                      `json:"total"`
}
