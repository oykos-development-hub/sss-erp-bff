package structs

type OfficesOfOrganizationUnitItem struct {
	Id                 int    `json:"id"`
	OrganizationUnitId int    `json:"organization_unit_id"`
	Title              string `json:"title"`
	Abbreviation       string `json:"abbreviation"`
	Description        string `json:"description"`
	Color              string `json:"color"`
	Icon               string `json:"icon"`
}
