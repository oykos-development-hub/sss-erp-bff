package structs

type Evaluation struct {
	ID                  int              `json:"id"`
	UserProfileID       int              `json:"user_profile_id"`
	EvaluationTypeID    int              `json:"evaluation_type_id"`
	EvaluationType      SettingsDropdown `json:"evaluation_type"`
	Score               string           `json:"score"`
	DateOfEvaluation    *string          `json:"date_of_evaluation"`
	Evaluator           string           `json:"evaluator"`
	IsRelevant          bool             `json:"is_relevant"`
	CreatedAt           string           `json:"created_at"`
	UpdatedAt           string           `json:"updated_at"`
	FileIDs             []int            `json:"file_ids"`
	ReasonForEvaluation *string          `json:"reason_for_evaluation"`
	EvaluationPeriod    *string          `json:"evaluation_period"`
	DecisionNumber      *string          `json:"decision_number"`
}

type EvaluationType struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
	Value        string `json:"value"`
	Color        string `json:"color"`
	Icon         string `json:"icon"`
}
