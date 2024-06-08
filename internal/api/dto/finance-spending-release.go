package dto

import (
	"bff/structs"
	"time"

	"github.com/shopspring/decimal"
)

type GetSpendingReleaseResponseMS struct {
	Data structs.SpendingRelease `json:"data"`
}

type GetSpendingReleaseListResponseMS struct {
	Data []structs.SpendingRelease `json:"data"`
}

type GetSpendingReleaseListInput struct {
	Year     int `json:"year"`
	Month    int `json:"month"`
	UnitID   int `json:"unit_id"`
	BudgetID int `json:"budget_id"`
}

type SpendingReleaseOverviewFilterDTO struct {
	Month    int `json:"month"`
	Year     int `json:"year" validate:"required"`
	BudgetID int `json:"budget_id" validate:"required"`
	UnitID   int `json:"unit_id" validate:"required"`
}

type GetSpendingReleaseOverviewResponseMS struct {
	Data []SpendingReleaseOverviewItem `json:"data"`
}

type SpendingReleaseOverviewItem struct {
	Month     int             `json:"month"`
	Year      int             `json:"year"`
	CreatedAt time.Time       `json:"created_at"`
	Value     decimal.Decimal `json:"value"`
}

type SpendingReleaseDTO struct {
	ID                  int                   `json:"id"`
	AccountID           int                   `json:"account_id"`
	BudgetID            int                   `json:"budget_id"`
	UnitID              int                   `json:"unit_id"`
	CurrentBudgetID     int                   `json:"current_budget_id"`
	Value               decimal.Decimal       `json:"value"`
	AccountSerialNumber string                `json:"account_serial_number"`
	AccountTitle        string                `json:"account_title"`
	CreatedAt           time.Time             `json:"created_at"`
	Children            []*SpendingReleaseDTO `json:"children"`
}
