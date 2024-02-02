package structs

type ActivitiesItem struct {
	ID                 int    `json:"id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Code               string `json:"code"`
	SubProgramID       int    `json:"sub_program_id"`
	OrganizationUnitID int    `json:"organization_unit_id"`
}
