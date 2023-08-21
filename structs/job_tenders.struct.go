package structs

type JobTenders struct {
	Id                           int      `json:"id"`
	PositionInOrganizationUnitId *int     `json:"position_in_organization_unit_id"`
	OrganizationUnitID           int      `json:"organization_unit_id"`
	TypeID                       int      `json:"type"`
	Active                       bool     `json:"active"`
	DateOfStart                  JSONDate `json:"date_of_start"`
	DateOfEnd                    JSONDate `json:"date_of_end"`
	Description                  string   `json:"description"`
	SerialNumber                 string   `json:"serial_number"`
	AvailableSlots               int      `json:"available_slots"`
	FileId                       int      `json:"file_id"`
	CreatedAt                    string   `json:"created_at"`
	UpdatedAt                    string   `json:"updated_at"`
}

type JobTenderApplications struct {
	Id                 int      `json:"id"`
	JobTenderId        int      `json:"job_tender_id"`
	UserProfileId      *int     `json:"user_profile_id"`
	FileId             int      `json:"file_id"`
	Active             bool     `json:"active"`
	Type               string   `json:"type"`
	FirstName          *string  `json:"first_name"`
	LastName           *string  `json:"last_name"`
	Nationality        *string  `json:"nationality"`
	DateOfBirth        JSONDate `json:"date_of_birth"`
	DateOfApplication  JSONDate `json:"date_of_application"`
	OfficialPersonalID *string  `json:"official_personal_id"`
	Evaluation         *string  `json:"evaluation"`
	Status             string   `json:"status"`
	CreatedAt          string   `json:"created_at"`
	UpdatedAt          string   `json:"updated_at"`
}
