package structs

type AbsentType struct {
	Id                int    `json:"id"`
	ParentId          int    `json:"parent_id"`
	Title             string `json:"title"`
	Abbreviation      string `json:"abbreviation"`
	AccountingDaysOff bool   `json:"accounting_days_off"`
	Relocation        bool   `json:"relocation"`
	Description       string `json:"description"`
	Color             string `json:"color"`
	Icon              string `json:"icon"`
}

type Absent struct {
	Id                       int                `json:"id"`
	AbsentTypeId             int                `json:"absent_type_id"`
	AbsentType               AbsentType         `json:"absent_type"`
	UserProfileId            int                `json:"user_profile_id"`
	Location                 string             `json:"location"`
	TargetOrganizationUnitID *int               `json:"target_organization_unit_id"`
	TargetOrganizationUnit   *OrganizationUnits `json:"target_organization_unit"`
	Description              string             `json:"description"`
	DateOfStart              string             `json:"date_of_start"`
	DateOfEnd                string             `json:"date_of_end"`
	CreatedAt                string             `json:"created_at"`
	UpdatedAt                string             `json:"updated_at"`
	FileId                   int                `json:"file_id"`
}

type Vacation struct {
	Id                int    `json:"id"`
	UserProfileId     int    `json:"user_profile_id"`
	ResolutionTypeId  int    `json:"resolution_type_id"`
	ResolutionPurpose string `json:"resolution_purpose"`
	Year              int    `json:"year"`
	Value             string `json:"value"`
	NumberOfDays      int    `json:"number_of_days"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	FileId            int    `json:"file_id"`
}
