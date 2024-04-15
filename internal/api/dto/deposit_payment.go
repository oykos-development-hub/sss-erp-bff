package dto

import (
	"bff/structs"
	"time"
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
	Amount                    float64            `json:"amount"`
	MainBankAccount           bool               `json:"main_bank_account"`
	CurrentBankAccount        string             `json:"current_bank_account"`
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
	OrganizationUnitID *int    `json:"organization_unit_id"`
}

type GetDepositPaymentResponseMS struct {
	Data structs.DepositPayment `json:"data"`
}

type GetDepositPaymentListResponseMS struct {
	Data  []structs.DepositPayment `json:"data"`
	Total int                      `json:"total"`
}
