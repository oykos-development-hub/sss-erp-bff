package dto

import (
	"time"
)

type Contract struct {
	Id                            int             `json:"id"`
	Title                         string          `json:"title"`
	ContractType                  DropdownSimple  `json:"contract_type"`
	OrganizationUnit              DropdownSimple  `json:"organization_unit"`
	Department                    *DropdownSimple `json:"department"`
	JobPositionInOrganizationUnit DropdownSimple  `json:"job_position_in_organization_unit"`
	UserProfile                   DropdownSimple  `json:"user_profile"`
	Abbreviation                  string          `json:"abbreviation"`
	Description                   string          `json:"description"`
	Active                        bool            `json:"active"`
	SerialNumber                  string          `json:"serial_number"`
	NetSalary                     string          `json:"net_salary"`
	GrossSalary                   string          `json:"gross_salary"`
	BankAccount                   string          `json:"bank_account"`
	BankName                      string          `json:"bank_name"`
	DateOfSignature               time.Time       `json:"date_of_signature"`
	DateOfEligibility             time.Time       `json:"date_of_eligibility"`
	DateOfStart                   time.Time       `json:"date_of_start"`
	DateOfEnd                     time.Time       `json:"date_of_end"`
	CreatedAt                     string          `json:"created_at"`
	UpdatedAt                     string          `json:"updated_at"`
	FileId                        int             `json:"file_id"`
}
