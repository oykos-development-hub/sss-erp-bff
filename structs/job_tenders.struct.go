package structs

type JobTenders struct {
	Id                           int      `json:"id"`
	PositionInOrganizationUnitId *int     `json:"position_in_organization_unit_id"`
	OrganizationUnitID           int      `json:"organization_unit_id"`
	TypeID                       int      `json:"type"`
	DateOfStart                  JSONDate `json:"date_of_start"`
	DateOfEnd                    JSONDate `json:"date_of_end"`
	Description                  string   `json:"description"`
	SerialNumber                 string   `json:"serial_number"`
	FileId                       int      `json:"file_id"`
	CreatedAt                    string   `json:"created_at"`
	UpdatedAt                    string   `json:"updated_at"`
}

type JobTenderApplications struct {
	Id                 int    `json:"id"`
	JobTenderId        int    `json:"job_tender_id"`
	UserProfileId      *int   `json:"user_profile_id"`
	Active             bool   `json:"active"`
	Type               string `json:"type"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Nationality        string `json:"nationality"`
	DateOfBirth        string `json:"date_of_birth"`
	DateOfApplication  string `json:"date_of_application"`
	OfficialPersonalID string `json:"official_personal_id"`
	Evaluation         string `json:"evaluation"`
	Status             string `json:"status"`
	FileId             int    `json:"file_id"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
}
