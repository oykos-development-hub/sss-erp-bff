package structs

type SalaryParams struct {
	Id               int    `json:"id"`
	UserProfileId    int    `json:"user_profile_id"`
	BenefitedTrack   bool   `json:"benefited_track"`
	WithoutRaise     bool   `json:"without_raise"`
	InsuranceBasis   string `json:"insurance_basis"`
	SalaryRank       string `json:"salary_rank"`
	DailyWorkHours   string `json:"daily_work_hours"`
	WeeklyWorkHours  string `json:"weekly_work_hours"`
	EducationRank    string `json:"education_rank"`
	EducationNaming  string `json:"education_naming"`
	UserResolutionId int    `json:"user_resolution_id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
