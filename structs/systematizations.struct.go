package structs

type Systematization struct {
	ID                 int     `json:"id"`
	UserProfileID      int     `json:"user_profile_id"`
	OrganizationUnitID int     `json:"organization_unit_id"`
	Description        string  `json:"description"`
	SerialNumber       string  `json:"serial_number"`
	Active             int     `json:"active"`
	DateOfActivation   *string `json:"date_of_activation"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
	FileID             int     `json:"file_id"`
}
