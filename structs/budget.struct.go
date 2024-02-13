package structs

type BudgetIndent struct {
	ID           int    `json:"id"`
	ParentID     int    `json:"parent_id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
}

type BudgetStatus int

const (
BudgetCreatedStatus BudgetStatus = 1
	BudgetSentStatus    BudgetStatus = 2
	BudgetClosedStatus  BudgetStatus = 3
)

type Budget struct {
	ID         int          `json:"id"`
	Year       int          `json:"year"`
	BudgetType int          `json:"budget_type"`
	Status     BudgetStatus `json:"budget_status"`
}

type BudgetRequestStatus int

const (
	BudgetRequestCreatedStatus BudgetRequestStatus = 1
	BudgetRequestSentStatus    BudgetRequestStatus = 2
	BudgetRequestClosedStatus  BudgetRequestStatus = 3
)

type RequestType int

const (
	DonationFinancialRequestType RequestType = 1
	CurrentFinancialRequestType  RequestType = 2
	NonFinancialRequestType      RequestType = 3
)

type BudgetRequest struct {
	ID                 int                 `json:"id"`
	OrganizationUnitID int                 `json:"organization_unit_id"`
	BudgetID           int                 `json:"budget_id"`
	RequestType        RequestType         `json:"request_type"`
	Status             BudgetRequestStatus `json:"status"`
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

type FilledFinanceBudget struct {
	ID                 int    `json:"id"`
	OrganizationUnitID int    `json:"organization_unit_id"`
	FinanceBudgetID    int    `json:"finance_budget_id"`
	AccountID          int    `json:"account_id"`
	CurrentYear        int    `json:"current_year"`
	NextYear           int    `json:"next_year"`
	YearAfterNext      int    `json:"year_after_next"`
	Description        string `json:"description"`
}
