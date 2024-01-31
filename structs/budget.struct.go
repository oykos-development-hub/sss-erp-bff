package structs

type BudgetIndent struct {
	ID           int    `json:"id"`
	ParentID     int    `json:"parent_id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
}

type Budget struct {
	ID         int `json:"id"`
	Year       int `json:"year"`
	BudgetType int `json:"budget_type"`
}

type FinancialBudget struct {
	ID             int `json:"id"`
	AccountVersion int `json:"account_version"`
	BudgetID       int `json:"budget_id"`
}

type FinancialBudgetLimit struct {
	ID                 int `json:"id"`
	Limit              int `json:"limit"`
	OrganizationUnitID int `json:"organization_unit_id"`
	FinancialBudgetID  int `json:"financial_budget_id"`
}
