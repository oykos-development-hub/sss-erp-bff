package structs

type JobPositions struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Abbreviation     string `json:"abbreviation"`
	Description      string `json:"description"`
	Requirements     string `json:"requirements"`
	SerialNumber     string `json:"serial_number"`
	IsJudge          bool   `json:"is_judge"`
	IsActive         bool   `json:"is_active"`
	IsJudgePresident bool   `json:"is_judge_president"`
	Color            string `json:"color"`
	Icon             string `json:"icon"`
}

type JobPositionsInOrganizationUnitsSettings struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type EmployeesInOrganizationUnits struct {
	ID                           int    `json:"id"`
	UserAccountID                int    `json:"user_account_id"`
	UserProfileID                int    `json:"user_profile_id"`
	PositionInOrganizationUnitID int    `json:"position_in_organization_unit_id"`
	Active                       bool   `json:"active"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
}

type ActiveEmployees struct {
	ID           int              `json:"id"`
	FullName     string           `json:"full_name"`
	JobPositions SettingsDropdown `json:"job_position"`
	Sector       string           `json:"sector"`
}

type JobPositionsInOrganizationUnits struct {
	ID                       int     `json:"id"`
	SystematizationID        int     `json:"systematization_id"`
	ParentOrganizationUnitID int     `json:"parent_organization_unit_id"`
	JobPositionID            int     `json:"job_position_id"`
	AvailableSlots           int     `json:"available_slots"`
	Requirements             *string `json:"requirements"`
	Description              *string `json:"description"`
	Employees                []int   `json:"employees"`
}
