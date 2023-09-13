package dto

type OfficesOfOrganizationResponse struct {
	Id               int            `json:"id"`
	OrganizationUnit DropdownSimple `json:"organization_unit"`
	Title            string         `json:"title"`
	Abbreviation     string         `json:"abbreviation"`
	Description      string         `json:"description"`
	Color            string         `json:"color"`
	Icon             string         `json:"icon"`
}
