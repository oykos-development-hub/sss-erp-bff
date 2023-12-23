package structs

type Experience struct {
	ID                        int    `json:"id"`
	UserProfileID             int    `json:"user_profile_id"`
	OrganizationUnitID        int    `json:"organization_unit_id,omitempty"`
	Relevant                  bool   `json:"relevant"`
	OrganizationUnit          string `json:"organization_unit"`
	AmountOfExperience        int    `json:"amount_of_experience"`
	AmountOfInsuredExperience int    `json:"amount_of_insured_experience"`
	DateOfStart               string `json:"date_of_start"`
	DateOfEnd                 string `json:"date_of_end"`
	CreatedAt                 string `json:"created_at"`
	UpdatedAt                 string `json:"updated_at"`
	ReferenceFileID           int    `json:"reference_file_id"`
}
