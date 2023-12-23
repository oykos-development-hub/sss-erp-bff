package structs

type OfficesOfOrganizationUnitItem struct {
	ID                 int    `json:"id"`
	OrganizationUnitID int    `json:"organization_unit_id"`
	Title              string `json:"title"`
	Abbreviation       string `json:"abbreviation"`
	Description        string `json:"description"`
	Color              string `json:"color"`
	Icon               string `json:"icon"`
}
