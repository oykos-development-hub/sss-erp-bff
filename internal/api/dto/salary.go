package dto

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
