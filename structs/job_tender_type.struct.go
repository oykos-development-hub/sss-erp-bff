package structs

type JobTenderTypes struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Abbreviation     string `json:"abbreviation"`
	Description      string `json:"description"`
	Value            string `json:"value"`
	IsJudge          bool   `json:"is_judge"`
	IsJudgePresident bool   `json:"is_judge_president"`
	Color            string `json:"color"`
	Icon             string `json:"icon"`
}
