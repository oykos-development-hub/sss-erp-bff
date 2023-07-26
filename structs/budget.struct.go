package structs

type BudgetIndent struct {
	Id           int    `json:"id"`
	ParentId     int    `json:"parent_id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
}
