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
	SortByTitle        *string `json:"sort_by_title"`
	Type               *string `json:"type"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Search             *string `json:"search"`
	JudgeID            *int    `json:"judge_id"`
	Status             *string `json:"status"`
}

type SalaryWillFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	SortByTitle        *string `json:"sort_by_title"`
	Search             *string `json:"search"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Status             *string `json:"status"`
}

type SalaryResponse struct {
	ID                   int                `json:"id"`
	OrganizationUnit     DropdownSimple     `json:"organization_unit"`
	Subject              string             `json:"subject"`
	Judge                DropdownSimple     `json:"judge"`
	CaseNumber           string             `json:"case_number"`
	DateOfRecipiet       *time.Time         `json:"date_of_recipiet"`
	DateOfCase           *time.Time         `json:"date_of_case"`
	DateOfFinality       *time.Time         `json:"date_of_finality"`
	DateOfEnforceability *time.Time         `json:"date_of_enforceability"`
	DateOfEnd            *time.Time         `json:"date_of_end"`
	Account              DropdownSimple     `json:"account"`
	File                 FileDropdownSimple `json:"file"`
	Status               string             `json:"status"`
	Type                 string             `json:"type"`
	CreatedAt            time.Time          `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
}

type SalaryItemResponse struct {
	ID                 int                `json:"id"`
	DepositID          int                `json:"deposit_id"`
	Category           DropdownSimple     `json:"category"`
	Type               DropdownSimple     `json:"type"`
	Unit               string             `json:"unit"`
	Currency           string             `json:"currency"`
	Amount             float32            `json:"amount"`
	SerialNumber       string             `json:"serial_number"`
	DateOfConfiscation *time.Time         `json:"date_of_confiscation"`
	CaseNumber         string             `json:"case_number"`
	Judge              DropdownSimple     `json:"judge"`
	File               FileDropdownSimple `json:"file"`
	CreatedAt          time.Time          `json:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at"`
}
