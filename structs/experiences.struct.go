package structs

type Experience struct {
	ID                        int    `json:"id"`
	UserProfileID             int    `json:"user_profile_id"`
	OrganizationUnitID        int    `json:"organization_unit_id,omitempty"`
	Relevant                  bool   `json:"relevant"`
	OrganizationUnit          string `json:"organization_unit"`
	YearsOfExperience         int    `json:"years_of_experience"`
	YearsOfInsuredExperience  int    `json:"years_of_insured_experience"`
	MonthsOfExperience        int    `json:"months_of_experience"`
	MonthsOfInsuredExperience int    `json:"months_of_insured_experience"`
	DaysOfExperience          int    `json:"days_of_experience"`
	DaysOfInsuredExperience   int    `json:"days_of_insured_experience"`
	DateOfStart               string `json:"date_of_start"`
	DateOfEnd                 string `json:"date_of_end"`
	CreatedAt                 string `json:"created_at"`
	UpdatedAt                 string `json:"updated_at"`
	FileIDs                   []int  `json:"file_ids"`
}
