package structs

type Suppliers struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	OfficialId   string `json:"official_id"`
	Address      string `json:"address"`
	Description  string `json:"description"`
	FolderId     int    `json:"folder_id"`
}
