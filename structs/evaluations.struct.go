package structs

type Evaluation struct {
	Id               int              `json:"id"`
	UserProfileId    int              `json:"user_profile_id"`
	EvaluationTypeId int              `json:"evaluation_type_id"`
	EvaluationType   SettingsDropdown `json:"evaluation_type"`
	Score            string           `json:"score"`
	DateOfEvaluation *string          `json:"date_of_evaluation"`
	Evaluator        string           `json:"evaluator"`
	IsRelevant       bool             `json:"is_relevant"`
	CreatedAt        string           `json:"created_at"`
	UpdatedAt        string           `json:"updated_at"`
	FileId           int              `json:"file_id"`
}

type EvaluationType struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
	Value        string `json:"value"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
}
