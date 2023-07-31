package structs

type OrganizationUnits struct {
	Id             int    `json:"id"`
	ParentId       int    `json:"parent_id"`
	NumberOfJudges int    `json:"number_of_judges"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Abbreviation   string `json:"abbreviation"`
	Address        string `json:"address"`
	Color          string `json:"color"`
	Icon           string `json:"icon"`
	FolderId       int    `json:"folder_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
