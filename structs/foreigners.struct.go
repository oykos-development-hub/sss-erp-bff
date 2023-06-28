package structs

type Foreigners struct {
	Id                              int    `json:"id"`
	UserProfileId                   int    `json:"user_profile_id"`
	WorkPermitNumber                string `json:"work_permit_number"`
	WorkPermitIssuer                string `json:"work_permit_issuer"`
	WorkPermitDateOfStart           string `json:"work_permit_date_of_start"`
	WorkPermitDateOfEnd             string `json:"work_permit_date_of_end"`
	WorkPermitIndefiniteLength      bool   `json:"work_permit_indefinite_length"`
	ResidencePermitDateOfStart      string `json:"residence_permit_date_of_start"`
	ResidencePermitDateOfEnd        string `json:"residence_permit_date_of_end"`
	ResidencePermitIndefiniteLength bool   `json:"residence_permit_indefinite_length"`
	ResidencePermitNumber           string `json:"residence_permit_number"`
	ResidencePermitIssuer           string `json:"residence_permit_issuer"`
	CountryOfOrigin                 string `json:"country_of_origin"`
	CreatedAt                       string `json:"created_at"`
	UpdatedAt                       string `json:"updated_at"`
	WorkPermitFileId                int    `json:"work_permit_file_id"`
	ResidencePermitFileId           int    `json:"residence_permit_file_id"`
}
