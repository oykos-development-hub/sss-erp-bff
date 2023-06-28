package structs

type VacationType struct {
	Id                int    `json:"id"`
	ParentId          int    `json:"parent_id"`
	Title             string `json:"title"`
	Abbreviation      string `json:"abbreviation"`
	AccountingDaysOff bool   `json:"accounting_days_off"`
	Description       string `json:"description"`
	Color             string `json:"color"`
	Icon              string `json:"icon"`
}

type Vacation struct {
	Id             int    `json:"id"`
	VacationTypeId int    `json:"vacation_type_id"`
	UserProfileId  int    `json:"user_profile_id"`
	DateOfStart    string `json:"date_of_start"`
	DateOfEnd      string `json:"date_of_end"`
	Location       string `json:"location"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	FileId         int    `json:"file_id"`
}

type Relocation struct {
	Id                       int    `json:"id"`
	TargetOrganizationUnitId int    `json:"target_organization_unit_id"`
	UserProfileId            int    `json:"user_profile_id"`
	Description              string `json:"description"`
	DateOfStart              string `json:"date_of_start"`
	DateOfEnd                string `json:"date_of_end"`
	CreatedAt                string `json:"created_at"`
	UpdatedAt                string `json:"updated_at"`
	FileId                   int    `json:"file_id"`
}

type Absent struct {
	Id                       int    `json:"id"`
	VacationTypeId           int    `json:"vacation_type_id"`
	Location                 string `json:"location"`
	TargetOrganizationUnitId int    `json:"target_organization_unit_id"`
	UserProfileId            int    `json:"user_profile_id"`
	Description              string `json:"description"`
	DateOfStart              string `json:"date_of_start"`
	DateOfEnd                string `json:"date_of_end"`
	CreatedAt                string `json:"created_at"`
	UpdatedAt                string `json:"updated_at"`
	FileId                   int    `json:"file_id"`
}
