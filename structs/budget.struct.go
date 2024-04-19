package structs

import "time"

type BudgetIndent struct {
	ID           int    `json:"id"`
	ParentID     int    `json:"parent_id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
}

type BudgetStatus int

const (
	BudgetCreatedStatus         BudgetStatus = 1
	BudgetSentStatus            BudgetStatus = 2
	BudgetClosedStatus          BudgetStatus = 3
	ManagerBudgetProcessStatus  BudgetStatus = 4
	ManagerBudgetOnReviewStatus BudgetStatus = 5
	ManagerBudgetClosedStatus   BudgetStatus = 6
	OfficialBudgetSentStatus    BudgetStatus = 7
)

type Budget struct {
	ID         int          `json:"id"`
	Year       int          `json:"year"`
	BudgetType int          `json:"budget_type"`
	Status     BudgetStatus `json:"budget_status"`
}

type BudgetRequestStatus int

const (
	BudgetRequestSentStatus         BudgetRequestStatus = 1
	BudgetRequestFilledStatus       BudgetRequestStatus = 2
	BudgetRequestSentOnReviewStatus BudgetRequestStatus = 3
	BudgetRequestAcceptedStatus     BudgetRequestStatus = 4
	BudgetRequestRejectedStatus     BudgetRequestStatus = 5
)

func (s BudgetRequestStatus) StatusForOfficial() string {
	switch s {
	case BudgetRequestSentOnReviewStatus:
		return "Obradi"
	case BudgetRequestAcceptedStatus:
		return "Odobreno"
	default:
		return "Na čekanju"
	}
}

func (s BudgetRequestStatus) StatusForManager() string {
	switch s {
	case BudgetRequestSentStatus:
		return "Obradi"
	case BudgetRequestFilledStatus:
		return "Popunjeno"
	case BudgetRequestAcceptedStatus:
		return "Odobreno"
	default:
		return "Na čekanju"
	}
}

type RequestType int

const (
	CurrentFinancialRequestType  RequestType = 1
	DonationFinancialRequestType RequestType = 2
	NonFinancialRequestType      RequestType = 3
)

type BudgetRequest struct {
	ID                 int                 `json:"id"`
	OrganizationUnitID int                 `json:"organization_unit_id"`
	BudgetID           int                 `json:"budget_id"`
	RequestType        RequestType         `json:"request_type"`
	Status             BudgetRequestStatus `json:"status"`
	Comment            string              `json:"comment"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
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
	BudgetID           int `json:"budget_id"`
}

type FilledFinanceBudget struct {
	ID              int    `json:"id"`
	BudgetRequestID int    `json:"budget_request_id"`
	AccountID       int    `json:"account_id"`
	CurrentYear     int    `json:"current_year"`
	NextYear        int    `json:"next_year"`
	YearAfterNext   int    `json:"year_after_next"`
	Description     string `json:"description"`
}
