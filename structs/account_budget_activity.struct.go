package structs

type AccountBudgetActivityItem struct {
	ID                 int `json:"id"`
	AccountID          int `json:"account_id"`
	BudgetID           int `json:"budget_id"`
	ActivityID         int `json:"activity_id"`
	ValueCurrentYear   int `json:"value_current_year"`
	ValueNextYear      int `json:"value_next_year"`
	ValueAfterNextYear int `json:"value_after_next_year"`
}
