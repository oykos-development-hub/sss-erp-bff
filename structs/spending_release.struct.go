package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type SpendingReleaseInsert struct {
	AccountID int             `json:"account_id"`
	Month     int             `json:"month"`
	Value     decimal.Decimal `json:"value"`
	Username  string          `json:"username"`
}

type SpendingReleaseRequest struct {
	ID                     int       `json:"id"`
	Year                   int       `json:"year"`
	Month                  int       `json:"month"`
	OrganizationUnitID     int       `json:"organization_unit_id"`
	OrganizationUnitFileID int       `json:"organization_unit_file_id"`
	SSSFileID              int       `json:"sss_file_id"`
	Status                 string    `json:"status"`
	CreatedAt              time.Time `json:"created_at"`
}

type SpendingRelease struct {
	ID              int             `json:"id"`
	CurrentBudgetID int             `json:"current_budget_id"`
	BudgetID        int             `json:"budget_id"`
	UnitID          int             `json:"unit_id"`
	AccountID       int             `json:"account_id"`
	Value           decimal.Decimal `json:"value"`
	Month           int             `json:"month"`
	Year            int             `json:"year"`
	CreatedAt       time.Time       `json:"created_at"`
	Username        string          `json:"username"`
}
