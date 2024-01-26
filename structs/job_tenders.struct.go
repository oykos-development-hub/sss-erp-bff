package structs

type JobTenders struct {
	ID                           int     `json:"id"`
	PositionInOrganizationUnitID *int    `json:"position_in_organization_unit_id"`
	OrganizationUnitID           int     `json:"organization_unit_id"`
	TypeID                       int     `json:"type"`
	DateOfStart                  string  `json:"date_of_start"`
	DateOfEnd                    *string `json:"date_of_end"`
	Description                  string  `json:"description"`
	SerialNumber                 string  `json:"serial_number"`
	FileID                       int     `json:"file_id"`
	NumberOfVacantSeats          int     `json:"number_of_vacant_seats"`
	CreatedAt                    string  `json:"created_at"`
	UpdatedAt                    string  `json:"updated_at"`
}

type JobTenderApplications struct {
	ID                             int     `json:"id"`
	JobTenderID                    int     `json:"job_tender_id"`
	UserProfileID                  *int    `json:"user_profile_id"`
	Active                         bool    `json:"active"`
	Type                           string  `json:"type"`
	FirstName                      string  `json:"first_name"`
	LastName                       string  `json:"last_name"`
	Nationality                    string  `json:"citizenship"`
	DateOfBirth                    *string `json:"date_of_birth"`
	DateOfApplication              *string `json:"date_of_application"`
	OfficialPersonalDocumentNumber string  `json:"official_personal_document_number"`
	Evaluation                     int     `json:"evaluation"`
	Status                         string  `json:"status"`
	FileID                         int     `json:"file_id"`
	CreatedAt                      string  `json:"created_at"`
	UpdatedAt                      string  `json:"updated_at"`
}
