package dto

import (
	"bff/structs"
	"time"

	"github.com/shopspring/decimal"
)

type GetSpendingDynamicListResponseMS struct {
	Data []structs.SpendingDynamic `json:"data"`
}

type GetSpendingDynamicResponseMS struct {
	Data structs.SpendingDynamic `json:"data"`
}

type GetSpendingDynamicActualResponseMS struct {
	Data decimal.NullDecimal `json:"data"`
}

type GetSpendingDynamicHistoryResponseMS struct {
	Data []SpendingDynamicHistoryDTO `json:"data"`
}

type SpendingDynamicHistoryDTO struct {
	BudgetID  int       `json:"budget_id"`
	UnitID    int       `json:"unit_id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}
