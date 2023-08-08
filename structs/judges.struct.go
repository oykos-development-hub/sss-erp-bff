package structs

type JudgeNorms struct {
	Id                       int      `json:"id"`
	UserProfileId            int      `json:"user_profile_id"`
	Topic                    string   `json:"topic"`
	Title                    string   `json:"title"`
	PercentageOfNormDecrease float32  `json:"percentage_of_norm_decrease"`
	NumberOfNormDecrease     int      `json:"number_of_norm_decrease"`
	NumberOfItems            int      `json:"number_of_items"`
	NumberOfItemsSolved      int      `json:"number_of_items_solved"`
	EvaluationID             int      `json:"evaluation_id"`
	DateOfEvaluation         JSONDate `json:"date_of_evaluation"`
	DateOfEvaluationValidity JSONDate `json:"date_of_evaluation_validity"`
	FileID                   int      `json:"file_id"`
	RelocationID             *int     `json:"relocation_id"`
	CreatedAt                string   `json:"created_at"`
	UpdatedAt                string   `json:"updated_at"`
}

type JudgeResolutions struct {
	Id           int                     `json:"id"`
	SerialNumber string                  `json:"serial_number"`
	Year         string                  `json:"year"`
	CreatedAt    string                  `json:"created_at"`
	UpdatedAt    string                  `json:"updated_at"`
	Active       bool                    `json:"active"`
	Items        []*JudgeResolutionItems `json:"items"`
}

type JudgeResolutionItems struct {
	Id                 int `json:"id"`
	ResolutionId       int `json:"resolution_id"`
	OrganizationUnitId int `json:"organization_unit_id"`
	NumberOfJudges     int `json:"number_of_judges"`
	NumberOfPresidents int `json:"number_of_presidents"`
}
