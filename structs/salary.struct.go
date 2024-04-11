package structs

import "time"

type Salary struct {
	ID                       int                       `json:"id"`
	ActivityID               int                       `json:"activity_id"`
	Month                    string                    `json:"month"`
	DateOfCalculation        time.Time                 `json:"date_of_calculation"`
	Description              string                    `json:"description"`
	OrganizationUnitID       int                       `json:"organization_unit_id"`
	Status                   string                    `json:"status"`
	GrossPrice               float64                   `json:"gross_price"`
	VatPrice                 float64                   `json:"vat_price"`
	NetPrice                 float64                   `json:"net_price"`
	SalaryAdditionalExpenses []SalaryAdditionalExpense `json:"salary_additional_expenses"`
	CreatedAt                time.Time                 `json:"created_at"`
	UpdatedAt                time.Time                 `json:"updated_at"`
}

type SalaryAdditionalExpense struct {
	ID                 int       `json:"id"`
	Title              string    `json:"title"`
	SalaryID           int       `json:"salary_id"`
	AccountID          int       `json:"account_id"`
	Amount             float64   `json:"amount"`
	SubjectID          int       `json:"subject_id"`
	BankAccount        string    `json:"bank_account"`
	Status             string    `json:"status"`
	OrganizationUnitID int       `json:"organization_unit_id"`
	Type               string    `json:"type"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
