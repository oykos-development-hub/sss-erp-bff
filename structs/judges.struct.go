package structs

type JudgeNorms struct {
	Id                       int    `json:"id"`
	UserProfileId            int    `json:"user_profile_id"`
	NumberOfItems            int    `json:"number_of_items"`
	NumberOfSolvedItems      int    `json:"number_of_solved_items"`
	Norm                     int    `json:"norm"`
	Area                     string `json:"area"`
	PercentageOfNormDecrease string `json:"percentage_of_norm_decrease"`
	StartDate                string `json:"start_date"`
	EndDate                  string `json:"end_date"`
	Evaluation               string `json:"evaluation"`
	EvaluationValidTo        string `json:"evaluation_valid_to"`
	Relocation               string `json:"relocation"`
}

type JudgeResolutionData struct {
	Id            int                    `json:"id"`
	SerialNumber  string                 `json:"serial_number"`
	Year          string                 `json:"year"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
	UserProfileId int                    `json:"user_profile_id"`
	Items         []JudgeResolutionItems `json:"items"`
}

type JudgeResolutions struct {
	Id            int    `json:"id"`
	SerialNumber  string `json:"serial_number"`
	Year          string `json:"year"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	UserProfileId int    `json:"user_profile_id"`
}

type JudgeResolutionItems struct {
	Id                       int `json:"id"`
	ResolutionId             int `json:"resolution_id"`
	OrganizationUnitId       int `json:"organization_unit_id"`
	AvailableSlotsPresidents int `json:"available_slots_presidents"`
	AvailableSlotsJudges     int `json:"available_slots_judges"`
}
