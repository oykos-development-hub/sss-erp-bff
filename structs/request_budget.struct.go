package structs

type RequestBudgetType struct {
	ID                   int    `json:"id"`
	OrganizationUnitID   int    `json:"organization_unit_id"`
	DateCreate           string `json:"date_create"`
	BudgetID             int    `json:"budget_id"`
	ActivityID           int    `json:"activity_id"`
	StatusNotFinancially string `json:"status_not_financially"`
	StatusFinancially    string `json:"status_financially"`
}
