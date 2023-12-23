package structs

type BudgetIndent struct {
	ID           int    `json:"id"`
	ParentID     int    `json:"parent_id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
}

type BudgetItem struct {
	ID     int    `json:"id"`
	Year   string `json:"year"`
	Source string `json:"source"`
	Type   string `json:"type"`
	Status string `json:"status"`
}
