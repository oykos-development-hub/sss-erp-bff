package dto

import "bff/structs"

type Contract struct {
	ID                            int                  `json:"id"`
	Title                         string               `json:"title"`
	ContractType                  DropdownSimple       `json:"contract_type"`
	OrganizationUnit              DropdownSimple       `json:"organization_unit"`
	Department                    *DropdownSimple      `json:"department"`
	JobPositionInOrganizationUnit DropdownSimple       `json:"job_position_in_organization_unit"`
	UserProfile                   DropdownSimple       `json:"user_profile"`
	Abbreviation                  *string              `json:"abbreviation"`
	NumberOfConference            *string              `json:"number_of_conference"`
	Description                   *string              `json:"description"`
	Active                        bool                 `json:"active"`
	SerialNumber                  *string              `json:"serial_number"`
	NetSalary                     *string              `json:"net_salary"`
	GrossSalary                   *string              `json:"gross_salary"`
	BankAccount                   *string              `json:"bank_account"`
	BankName                      *string              `json:"bank_name"`
	DateOfSignature               *string              `json:"date_of_signature"`
	DateOfEligibility             *string              `json:"date_of_eligibility"`
	DateOfStart                   *string              `json:"date_of_start"`
	DateOfEnd                     *string              `json:"date_of_end"`
	CreatedAt                     string               `json:"created_at"`
	UpdatedAt                     string               `json:"updated_at"`
	Files                         []FileDropdownSimple `json:"files"`

	IsJudge                        bool    `json:"is_judge"`
	IsPresident                    bool    `json:"is_president"`
	JudgeApplicationSubmissionDate *string `json:"judge_application_submission_date"`
}

type ContractInsert struct {
	ID                              int     `json:"id"`
	Title                           string  `json:"title"`
	ContractTypeID                  int     `json:"contract_type_id"`
	OrganizationUnitID              int     `json:"organization_unit_id"`
	OrganizationUnitDepartmentID    *int    `json:"organization_unit_department_id"`
	NumberOfConference              *string `json:"number_of_conference"`
	JobPositionInOrganizationUnitID int     `json:"job_position_in_organization_unit_id"`
	UserProfileID                   int     `json:"user_profile_id"`
	Active                          bool    `json:"active"`
	DateOfSignature                 *string `json:"date_of_signature"`
	DateOfEligibility               *string `json:"date_of_eligibility"`
	DateOfStart                     *string `json:"date_of_start"`
	CreatedAt                       string  `json:"created_at"`
	UpdatedAt                       string  `json:"updated_at"`
	FileIDs                         []int   `json:"file_ids"`

	IsJudge                        bool    `json:"is_judge"`
	IsPresident                    bool    `json:"is_president"`
	JudgeApplicationSubmissionDate *string `json:"judge_application_submission_date"`
}

func (c *ContractInsert) ToEntity() *structs.Contracts {
	return &structs.Contracts{
		ID:                              c.ID,
		Title:                           c.Title,
		ContractTypeID:                  c.ContractTypeID,
		OrganizationUnitID:              c.OrganizationUnitID,
		OrganizationUnitDepartmentID:    c.OrganizationUnitDepartmentID,
		NumberOfConference:              c.NumberOfConference,
		JobPositionInOrganizationUnitID: c.JobPositionInOrganizationUnitID,
		UserProfileID:                   c.UserProfileID,
		Active:                          c.Active,
		DateOfSignature:                 c.DateOfSignature,
		DateOfEligibility:               c.DateOfEligibility,
		DateOfStart:                     c.DateOfStart,
		CreatedAt:                       c.CreatedAt,
		UpdatedAt:                       c.UpdatedAt,
		FileIDs:                         c.FileIDs,
	}
}
