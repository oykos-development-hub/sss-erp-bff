package structs

type Template struct {
	ID                 int    `json:"id,omitempty"`
	Title              string `json:"title"`
	OrganizationUnitID int    `json:"organization_unit_id"`
	FileID             int    `json:"file_id"`
	TemplateID         int    `json:"template_id"`
}

type CustomerSupport struct {
	ID                      int `json:"id"`
	UserDocumentationFileID int `json:"user_documentation_file_id"`
}

type ListOfParameters struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
