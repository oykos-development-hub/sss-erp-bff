package structs

type Resolution struct {
	ID                int               `json:"id"`
	UserProfileID     int               `json:"user_profile_id"`
	ResolutionTypeID  int               `json:"resolution_type_id"`
	ResolutionType    *SettingsDropdown `json:"resolution_type"`
	IsAffect          bool              `json:"is_affect"`
	ResolutionPurpose string            `json:"resolution_purpose"`
	DateOfStart       string            `json:"date_of_start"`
	DateOfEnd         string            `json:"date_of_end"`
	Year              int               `json:"year"`
	Value             string            `json:"value"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
	FileID            int               `json:"file_id"`
}

type ResolutionType struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
	Color        string `json:"color"`
}
