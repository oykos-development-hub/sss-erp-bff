package structs

type NonFinancialBudgetItem struct {
	ID                 int `json:"id"`
	RequestID          int `json:"request_id"`
	OrganizationUnitID int `json:"organization_unit_id"`
	BudgetID           int `json:"budget_id"`

	ImplContactFullName     string `json:"impl_contact_fullname"`
	ImplContactWorkingPlace string `json:"impl_contact_working_place"`
	ImplContactPhone        string `json:"impl_contact_phone"`
	ImplContactEmail        string `json:"impl_contact_email"`

	ContactFullName     string `json:"contact_fullname"`
	ContactWorkingPlace string `json:"contact_working_place"`
	ContactPhone        string `json:"contact_phone"`
	ContactEmail        string `json:"contact_email"`
}

type BudgetActivityNotFinanciallyProgramItem struct {
	ID                     int    `json:"id"`
	BudgetNotFinanciallyID int    `json:"budget_not_financially_id"` // foreign key BudgetActivityNotFinanciallyItem
	ProgramID              int    `json:"program_id"`                // foreign key ProgramItem
	Description            string `json:"description"`
}

type NonFinancialGoalItem struct {
	ID                   int    `json:"id"`
	NonFinancialBudgetID int    `json:"non_financial_budget_id"`
	Title                string `json:"title"`
	Description          string `json:"description"`
}

type NonFinancialGoalIndicatorItem struct {
	ID                       int    `json:"id"`
	GoalID                   int    `json:"goal_id"`
	PerformanceIndicatorCode string `json:"performance_indicator_code"`
	IndicatorSource          string `json:"indicator_source"`
	BaseYear                 string `json:"base_year"`
	GenderEquality           string `json:"gender_equality"`
	BaseValue                string `json:"base_value"`
	SourceOfInformation      string `json:"source_of_information"`
	UnitOfMeasure            string `json:"unit_of_measure"`
	IndicatorDescription     string `json:"indicator_description"`
	PlannedValue1            string `json:"planned_value_1"`
	RevisedValue1            string `json:"revised_value_1"`
	AchievedValue1           string `json:"achieved_value_1"`
	PlannedValue2            string `json:"planned_value_2"`
	RevisedValue2            string `json:"revised_value_2"`
	AchievedValue2           string `json:"achieved_value_2"`
	PlannedValue3            string `json:"planned_value_3"`
	RevisedValue3            string `json:"revised_value_3"`
	AchievedValue3           string `json:"achieved_value_3"`
}
