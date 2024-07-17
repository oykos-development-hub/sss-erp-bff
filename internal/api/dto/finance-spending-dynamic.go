package dto

import (
	"bff/structs"
	"time"

	"github.com/shopspring/decimal"
)

type GetSpendingDynamicListResponseMS struct {
	Data []SpendingDynamicDTO `json:"data"`
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
	Version   int       `json:"version"`
}

type GetSpendingDynamicHistoryInput struct {
	Version *int `json:"version"`
}

type SpendingDynamicDTO struct {
	ID                     int                   `json:"id"`
	AccountID              int                   `json:"account_id"`
	BudgetID               int                   `json:"budget_id"`
	UnitID                 int                   `json:"unit_id"`
	CurrentBudgetID        int                   `json:"current_budget_id"`
	Actual                 decimal.Decimal       `json:"actual"`
	CurrentAmount          decimal.Decimal       `json:"current_amount"`
	Username               string                `json:"username"`
	January                MonthEntry            `json:"january"`
	February               MonthEntry            `json:"february"`
	March                  MonthEntry            `json:"march"`
	April                  MonthEntry            `json:"april"`
	May                    MonthEntry            `json:"may"`
	June                   MonthEntry            `json:"june"`
	July                   MonthEntry            `json:"july"`
	August                 MonthEntry            `json:"august"`
	September              MonthEntry            `json:"september"`
	October                MonthEntry            `json:"october"`
	November               MonthEntry            `json:"november"`
	December               MonthEntry            `json:"december"`
	TotalSavings           decimal.Decimal       `json:"total_savings"`
	IsCurrentMonthEditable bool                  `json:"is_current_month_editable"`
	CreatedAt              time.Time             `json:"created_at"`
	AccountSerialNumber    string                `json:"account_serial_number"`
	AccountTitle           string                `json:"account_title"`
	Children               []*SpendingDynamicDTO `json:"children"`
}

type MonthEntry struct {
	Value   decimal.Decimal `json:"value"`
	Savings decimal.Decimal `json:"savings"`
}

type SpendingDynamicTreeDTO struct {
	ID              int             `json:"id"`
	AccountID       int             `json:"account_id"`
	UnitID          int             `json:"unit_id"`
	CurrentBudgetID int             `json:"current_budget_id"`
	Actual          decimal.Decimal `json:"actual"`
	Username        string          `json:"username"`
	January         MonthEntry      `json:"january"`
	February        MonthEntry      `json:"february"`
	March           MonthEntry      `json:"march"`
	April           MonthEntry      `json:"april"`
	May             MonthEntry      `json:"may"`
	June            MonthEntry      `json:"june"`
	July            MonthEntry      `json:"july"`
	August          MonthEntry      `json:"august"`
	September       MonthEntry      `json:"september"`
	October         MonthEntry      `json:"october"`
	November        MonthEntry      `json:"november"`
	December        MonthEntry      `json:"december"`
	TotalSavings    decimal.Decimal `json:"total_savings"`
	CreatedAt       time.Time       `json:"created_at"`
	Children        []SpendingDynamicDTO
}
