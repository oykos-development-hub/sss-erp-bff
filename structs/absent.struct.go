package structs

type AbsentType struct {
	ID                int    `json:"id"`
	ParentID          int    `json:"parent_id"`
	Title             string `json:"title"`
	Abbreviation      string `json:"abbreviation"`
	AccountingDaysOff bool   `json:"accounting_days_off"`
	Relocation        bool   `json:"relocation"`
	Description       string `json:"description"`
	Color             string `json:"color"`
	Icon              string `json:"icon"`
}

type Absent struct {
	ID                       int                `json:"id"`
	AbsentTypeID             int                `json:"absent_type_id"`
	AbsentType               AbsentType         `json:"absent_type"`
	UserProfileID            int                `json:"user_profile_id"`
	Location                 string             `json:"location"`
	TargetOrganizationUnitID *int               `json:"target_organization_unit_id"`
	TargetOrganizationUnit   *OrganizationUnits `json:"target_organization_unit"`
	Description              string             `json:"description"`
	DateOfStart              string             `json:"date_of_start"`
	DateOfEnd                string             `json:"date_of_end"`
	CreatedAt                string             `json:"created_at"`
	UpdatedAt                string             `json:"updated_at"`
	FileID                   int                `json:"file_id"`
	File                     FileDropdownSimple `json:"file"`
}

type Vacation struct {
	ID                int    `json:"id"`
	UserProfileID     int    `json:"user_profile_id"`
	ResolutionTypeID  int    `json:"resolution_type_id"`
	ResolutionPurpose string `json:"resolution_purpose"`
	Year              int    `json:"year"`
	Value             string `json:"value"`
	NumberOfDays      int    `json:"number_of_days"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	FileID            int    `json:"file_id"`
}

type VacationArray struct {
	Year int                 `json:"year"`
	Data []VacationArrayItem `json:"data"`
}

type VacationArrayItem struct {
	UserProfileID int `json:"user_profile_id"`
	NumberOfDays  int `json:"number_of_days"`
}

type FileDropdownSimple struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
