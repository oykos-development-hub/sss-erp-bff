package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type BudgetIndent struct {
	ID           int    `json:"id"`
	ParentID     int    `json:"parent_id"`
	Title        string `json:"title"`
	Abbreviation string `json:"abbreviation"`
}

type BudgetStatus int

const (
	BudgetCreatedStatus  BudgetStatus = 1
	BudgetSentStatus     BudgetStatus = 2
	BudgetAcceptedStatus BudgetStatus = 3
)

type Budget struct {
	ID         int          `json:"id"`
	Year       int          `json:"year"`
	BudgetType int          `json:"budget_type"`
	Status     BudgetStatus `json:"budget_status"`
}

type BudgetRequestStatus int

const (
	BudgetRequestSentStatus            BudgetRequestStatus = 1
	BudgetRequestFilledStatus          BudgetRequestStatus = 2
	BudgetRequestSentOnReviewStatus    BudgetRequestStatus = 3
	BudgetRequestAcceptedStatus        BudgetRequestStatus = 4
	BudgetRequestRejectedStatus        BudgetRequestStatus = 5
	BudgetRequestWaitingForActual      BudgetRequestStatus = 6
	BudgetRequestCompletedActualStatus BudgetRequestStatus = 7
)

type RequestType int

const (
	RequestTypeGeneral           RequestType = 1
	RequestTypeNonFinancial      RequestType = 2
	RequestTypeFinancial         RequestType = 3
	RequestTypeCurrentFinancial  RequestType = 4
	RequestTypeDonationFinancial RequestType = 5
)

type BudgetRequest struct {
	ID                 int                 `json:"id"`
	ParentID           *int                `json:"parent_id"`
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
	ID              int                 `json:"id"`
	BudgetRequestID int                 `json:"budget_request_id"`
	AccountID       int                 `json:"account_id"`
	CurrentYear     decimal.Decimal     `json:"current_year"`
	NextYear        decimal.Decimal     `json:"next_year"`
	YearAfterNext   decimal.Decimal     `json:"year_after_next"`
	Actual          decimal.NullDecimal `json:"actual"`
	Description     string              `json:"description"`
}
