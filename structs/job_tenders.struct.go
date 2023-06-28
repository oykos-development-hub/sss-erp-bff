package structs

type JobTenders struct {
	Id                           int    `json:"id"`
	PositionInOrganizationUnitId int    `json:"position_in_organization_unit_id"`
	Active                       bool   `json:"active"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
	DateOfStart                  string `json:"date_of_start"`
	DateOfEnd                    string `json:"date_of_end"`
	Description                  string `json:"description"`
	Type                         string `json:"type"`
	SerialNumber                 string `json:"serial_number"`
	AvailableSlots               int    `json:"available_slots"`
	FileId                       int    `json:"file_id"`
}

type JobTenderApplications struct {
	Id                 int    `json:"id"`
	JobTenderId        int    `json:"job_tender_id"`
	UserProfileId      int    `json:"user_profile_id"`
	Type               string `json:"type"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	OfficialPersonalId string `json:"official_personal_id"`
	Nationality        string `json:"nationality"`
	Evaluation         string `json:"evaluation"`
	DateOfBirth        string `json:"date_of_birth"`
	DateOfApplication  string `json:"date_of_application"`
	FileId             int    `json:"file_id"`
	Status             string `json:"status"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}
