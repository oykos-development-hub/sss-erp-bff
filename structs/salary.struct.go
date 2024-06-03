package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type Salary struct {
	ID                       int                       `json:"id"`
	ActivityID               int                       `json:"activity_id"`
	IsDeletable              bool                      `json:"is_deletable"`
	DebtorID                 int                       `json:"debtor_id"`
	Month                    string                    `json:"month"`
	Registred                bool                      `json:"registred"`
	DateOfCalculation        time.Time                 `json:"date_of_calculation"`
	Description              string                    `json:"description"`
	OrganizationUnitID       int                       `json:"organization_unit_id"`
	Status                   string                    `json:"status"`
	NumberOfEmployees        int                       `json:"number_of_employees"`
	GrossPrice               decimal.Decimal           `json:"gross_price"`
	VatPrice                 decimal.Decimal           `json:"vat_price"`
	NetPrice                 decimal.Decimal           `json:"net_price"`
	ObligationsPrice         decimal.Decimal           `json:"obligations_price"`
	SalaryAdditionalExpenses []SalaryAdditionalExpense `json:"salary_additional_expenses"`
	CreatedAt                time.Time                 `json:"created_at"`
	UpdatedAt                time.Time                 `json:"updated_at"`
}

type SalaryAdditionalExpense struct {
	ID                  int             `json:"id"`
	Title               string          `json:"title"`
	SalaryID            int             `json:"salary_id"`
	AccountID           int             `json:"account_id"`
	DebtorID            int             `json:"debtor_id"`
	IdentificatorNumber string          `json:"identificator_number"`
	Amount              decimal.Decimal `json:"amount"`
	SubjectID           int             `json:"subject_id"`
	BankAccount         string          `json:"bank_account"`
	Status              string          `json:"status"`
	OrganizationUnitID  int             `json:"organization_unit_id"`
	Type                string          `json:"type"`
	CreatedAt           time.Time       `json:"created_at"`
	UpdatedAt           time.Time       `json:"updated_at"`
}
