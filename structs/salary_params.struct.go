package structs

type SalaryParams struct {
	ID                  int              `json:"id,omitempty"`
	UserProfileID       int              `json:"user_profile_id"`
	OrganizationUnitID  int              `json:"organization_unit_id"`
	OrganizationUnit    SettingsDropdown `json:"organization_unit"`
	BenefitedTrack      bool             `json:"benefited_track"`
	WithoutRaise        bool             `json:"without_raise"`
	InsuranceBasis      string           `json:"insurance_basis"`
	SalaryRank          string           `json:"salary_rank"`
	DailyWorkHours      string           `json:"daily_work_hours"`
	WeeklyWorkHours     string           `json:"weekly_work_hours"`
	EducationRank       string           `json:"education_rank"`
	EducationNaming     string           `json:"education_naming"`
	UserResolutionID    *int             `json:"user_resolution_id,omitempty"`
	ObligationReduction string           `json:"obligation_reduction"`
	CreatedAt           string           `json:"created_at"`
	UpdatedAt           string           `json:"updated_at"`
}
