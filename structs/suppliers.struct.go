package structs

type Suppliers struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Entity       string `json:"entity"`
	Abbreviation string `json:"abbreviation"`
	OfficialID   string `json:"official_id"`
	Address      string `json:"address"`
	Description  string `json:"description"`
	FolderID     int    `json:"folder_id"`
}
