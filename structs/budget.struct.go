package structs

type BudgetIndent struct {
	Id           int    `json:"id"`
	ParentId     int    `json:"parent_id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
}

type BudgeItem struct {
	Id         int    `json:"id"`
	Year       string `json:"year"`
	ActivityId int    `json:"activity_id"`
	Source     string `json:"source"`
	Type       string `json:"type"`
	Status     string `json:"status"`
}
