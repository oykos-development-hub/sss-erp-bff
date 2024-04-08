package structs

type SettingsDropdown struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Entity       string `json:"entity"`
	Value        string `json:"value"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
	ParentID     int    `json:"parent_id"`
}
