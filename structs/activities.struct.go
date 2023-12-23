package structs

type ActivitiesItem struct {
	ID                 int    `json:"id"`
	Title              string `json:"title"`
	SubroutineID       int    `json:"subroutine_id"`
	OrganizationUnitID int    `json:"organization_unit_id"`
}
