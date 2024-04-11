package dto

import (
	"bff/structs"
	"time"
)

type SalaryParams struct {
	ID                  int            `json:"id,omitempty"`
	UserProfile         DropdownSimple `json:"user_profile"`
	OrganizationUnit    DropdownSimple `json:"organization_unit"`
	Resolution          Resolution     `json:"resolution"`
	BenefitedTrack      bool           `json:"benefited_track"`
	WithoutRaise        bool           `json:"without_raise"`
	InsuranceBasis      string         `json:"insurance_basis"`
	SalaryRank          string         `json:"salary_rank"`
	DailyWorkHours      string         `json:"daily_work_hours"`
	WeeklyWorkHours     string         `json:"weekly_work_hours"`
	EducationRank       string         `json:"education_rank"`
	EducationNaming     string         `json:"education_naming"`
	ObligationReduction string         `json:"obligation_reduction"`
	CreatedAt           string         `json:"created_at"`
	UpdatedAt           string         `json:"updated_at"`
}

type GetSalaryResponseMS struct {
	Data structs.Salary `json:"data"`
}

type GetSalaryListResponseMS struct {
	Data  []structs.Salary `json:"data"`
	Total int              `json:"total"`
}

type SalaryFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	Year               *int    `json:"year"`
	ActivityID         *int    `json:"activity_id"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Search             *string `json:"search"`
}

type SalaryResponse struct {
	ID                       int                                `json:"id"`
	Activity                 DropdownSimple                     `json:"activity"`
	Month                    string                             `json:"month"`
	DateOfCalculation        time.Time                          `json:"date_of_calculation"`
	Description              string                             `json:"description"`
	Status                   string                             `json:"status"`
	OrganizationUnit         DropdownSimple                     `json:"organization_unit"`
	SalaryAdditionalExpenses []SalaryAdditionalExpensesResponse `json:"salary_additional_expenses"`
	GrossPrice               float64                            `json:"gross_price"`
	VatPrice                 float64                            `json:"vat_price"`
	NetPrice                 float64                            `json:"net_price"`
	CreatedAt                time.Time                          `json:"created_at"`
	UpdatedAt                time.Time                          `json:"updated_at"`
}

type SalaryAdditionalExpensesResponse struct {
	ID               int            `json:"id"`
	SalaryID         int            `json:"salary_id"`
	Account          DropdownSimple `json:"account"`
	Amount           float64        `json:"amount"`
	Subject          DropdownSimple `json:"subject"`
	BankAccount      string         `json:"bank_account"`
	Status           string         `json:"status"`
	Title            string         `json:"title"`
	OrganizationUnit DropdownSimple `json:"organization_unit"`
	Type             string         `json:"type"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}
