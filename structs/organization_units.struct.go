package structs

type OrganizationUnits struct {
	Id             int    `json:"id"`
	ParentId       *int   `json:"parent_id,omitempty"`
	NumberOfJudges int    `json:"number_of_judges"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Pib            string `json:"pib"`
	Abbreviation   string `json:"abbreviation"`
	Address        string `json:"address"`
	Color          string `json:"color"`
	Icon           string `json:"icon"`
	FolderId       int    `json:"folder_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
