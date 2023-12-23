package structs

type ProgramItem struct {
	ID                 int    `json:"id"`
	ParentID           int    `json:"parent_id"`
	Title              string `json:"title"`
	OrganizationUnitID int    `json:"organization_unit_id"`
}
