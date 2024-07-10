package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type CurrentBudget struct {
	ID            int             `json:"id"`
	BudgetID      int             `json:"budget_id"`
	UnitID        int             `json:"unit_id"`
	AccountID     int             `json:"account_id"`
	InitialActual decimal.Decimal `json:"initial_actual"`
	Actual        decimal.Decimal `json:"actual"`
	Balance       decimal.Decimal `json:"balance"`
	CurrentAmount decimal.Decimal `json:"current_amount"`
	Type          int             `json:"type"`
	CreatedAt     time.Time       `json:"created_at"`
}
