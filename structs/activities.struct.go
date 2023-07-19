package structs

type ActivitiesItem struct {
	Id                 int    `json:"id"`
	Title              string `json:"title"`
	SubroutineId       int    `json:"subroutine_id"`
	OrganizationUnitId int    `json:"organization_unit_id"`
}
