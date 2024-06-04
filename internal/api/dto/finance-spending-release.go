package dto

import (
	"bff/structs"

	"github.com/shopspring/decimal"
)

type GetSpendingReleaseResponseMS struct {
	Data structs.SpendingRelease `json:"data"`
}

type GetSpendingReleaseListInput struct {
	Year      *int `json:"year"`
	Month     *int `json:"month"`
	UnitID    int  `json:"unit_id"`
	AccountID int  `json:"account_id"`
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
	Month int             `json:"month"`
	Year  int             `json:"year"`
	Value decimal.Decimal `json:"value"`
}
