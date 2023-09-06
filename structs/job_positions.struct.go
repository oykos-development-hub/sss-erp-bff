package structs

type JobPositions struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	Abbreviation     string `json:"abbreviation"`
	Description      string `json:"description"`
	Requirements     string `json:"requirements"`
	SerialNumber     string `json:"serial_number"`
	IsJudge          bool   `json:"is_judge"`
	IsJudgePresident bool   `json:"is_judge_president"`
	Color            string `json:"color"`
	Icon             string `json:"icon"`
}

type JobPositionsInOrganizationUnitsSettings struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type EmployeesInOrganizationUnits struct {
	Id                           int    `json:"id"`
	UserAccountId                int    `json:"user_account_id"`
	UserProfileId                int    `json:"user_profile_id"`
	PositionInOrganizationUnitId int    `json:"position_in_organization_unit_id"`
	Active                       bool   `json:"active"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
}

type JobPositionsInOrganizationUnits struct {
	Id                       int   `json:"id"`
	SystematizationId        int   `json:"systematization_id"`
	ParentOrganizationUnitId int   `json:"parent_organization_unit_id"`
	JobPositionId            int   `json:"job_position_id"`
	AvailableSlots           int   `json:"available_slots"`
	Employees                []int `json:"employees"`
}
