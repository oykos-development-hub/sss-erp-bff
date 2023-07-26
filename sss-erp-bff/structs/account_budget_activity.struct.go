package structs

type AccountBudgetActivityItem struct {
	Id                 int `json:"id"`
	AccountId          int `json:"account_id"`
	BudgetId           int `json:"budget_id"`
	ActivityId         int `json:"activity_id"`
	ValueCurrentYear   int `json:"value_current_year"`
	ValueNextYear      int `json:"value_next_year"`
	ValueAfterNextYear int `json:"value_after_next_year"`
}
