package structs

type EngagementType struct {
    ID           int    `json:"id"`
    Title        string `json:"title"`
    Abbreviation string `json:"abbreviation"`
    Color        string `json:"color"`
    Icon         string `json:"icon"`
}