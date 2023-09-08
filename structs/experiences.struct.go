package structs

import "time"

type Experience struct {
	Id                        int       `json:"id"`
	UserProfileId             int       `json:"user_profile_id"`
	OrganizationUnitId        int       `json:"organization_unit_id,omitempty"`
	Relevant                  bool      `json:"relevant"`
	OrganizationUnit          string    `json:"organization_unit"`
	AmountOfExperience        int       `json:"amount_of_experience"`
	AmountOfInsuredExperience int       `json:"amount_of_insured_experience"`
	DateOfSignature           time.Time `json:"date_of_signature"`
	DateOfStart               time.Time `json:"date_of_start"`
	DateOfEnd                 time.Time `json:"date_of_end"`
	CreatedAt                 string    `json:"created_at"`
	UpdatedAt                 string    `json:"updated_at"`
	ReferenceFileId           int       `json:"reference_file_id"`
}
