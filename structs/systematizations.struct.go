package structs

type Systematization struct {
	Id                 int     `json:"id"`
	UserProfileId      int     `json:"user_profile_id"`
	OrganizationUnitId int     `json:"organization_unit_id"`
	Description        string  `json:"description"`
	SerialNumber       string  `json:"serial_number"`
	Active             bool    `json:"active"`
	DateOfActivation   *string `json:"date_of_activation"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
	FileId             int     `json:"file_id"`
}
