package structs

type Template struct {
	ID                 int    `json:"id,omitempty"`
	Title              string `json:"title"`
	OrganizationUnitID int    `json:"organization_unit_id"`
	FileID             int    `json:"file_id"`
	TemplateID         int    `json:"template_id"`
}
