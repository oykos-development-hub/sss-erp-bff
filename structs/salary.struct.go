package structs

import "time"

type Salary struct {
	ID                   int        `json:"id"`
	OrganizationUnitID   int        `json:"organization_unit_id"`
	Subject              string     `json:"subject"`
	JudgeID              int        `json:"judge_id"`
	CaseNumber           string     `json:"case_number"`
	DateOfRecipiet       *time.Time `json:"date_of_recipiet"`
	DateOfCase           *time.Time `json:"date_of_case"`
	DateOfFinality       *time.Time `json:"date_of_finality"`
	DateOfEnforceability *time.Time `json:"date_of_enforceability"`
	DateOfEnd            *time.Time `json:"date_of_end"`
	AccountID            int        `json:"account_id"`
	FileID               int        `json:"file_id"`
	Status               string     `json:"status"`
	Type                 string     `json:"type"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

type SalaryWill struct {
	ID                 int        `json:"id"`
	OrganizationUnitID int        `json:"organization_unit_id"`
	Subject            string     `json:"subject"`
	FatherName         string     `json:"father_name"`
	DateOfBirth        time.Time  `json:"date_of_birth"`
	JMBG               string     `json:"jmbg"`
	CaseNumberSI       string     `json:"case_number_si"`
	CaseNumberRS       string     `json:"case_number_rs"`
	DateOfReceiptSI    *time.Time `json:"date_of_receipt_si"`
	DateOfReceiptRS    *time.Time `json:"date_of_receipt_rs"`
	DateOfEnd          *time.Time `json:"date_of_end"`
	Status             string     `json:"status"`
	FileID             int        `json:"file_id"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}
