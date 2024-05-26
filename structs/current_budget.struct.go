package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type CurrentBudget struct {
	ID             int             `json:"id"`
	BudgetID       int             `json:"budget_id"`
	UnitID         int             `json:"unit_id"`
	AccountID      int             `json:"account_id"`
	InititalActual decimal.Decimal `json:"actual"`
	Actual         decimal.Decimal `json:"spending_dynamic_id"`
	Balance        decimal.Decimal `json:"balance"`
	CreatedAt      time.Time       `json:"created_at"`
}
