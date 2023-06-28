package structs

type Resolution struct {
	Id                int    `json:"id"`
	UserProfileId     int    `json:"user_profile_id"`
	ResolutionTypeId  int    `json:"resolution_type_id"`
	ResolutionPurpose string `json:"resolution_purpose"`
	DateOfStart       string `json:"date_of_start"`
	DateOfEnd         string `json:"date_of_end"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	FileId            int    `json:"file_id"`
}

type ResolutionType struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
	Color        string `json:"color"`
}
