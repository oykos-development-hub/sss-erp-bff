package structs

type RequestBudgetType struct {
	Id                   int    `json:"id"`
	OrganizationUnitId   int    `json:"organization_unit_id"`
	DateCreate           string `json:"date_create"`
	BudgetId             int    `json:"budget_id"`
	ActivityId           int    `json:"activity_id"`
	StatusNotFinancially string `json:"status_not_financially"`
	StatusFinancially    string `json:"status_financially"`
}
