package structs

type BasicInventoryDepreciationTypesItem struct {
	Id               int    `json:"id"`
	Title            string `json:"title"`
	Abbreviation     string `json:"abbreviation"`
	LifetimeInMonths int    `json:"lifetime_in_months"`
	Description      string `json:"description"`
	Color            string `json:"color"`
	Icon             string `json:"icon"`
}
