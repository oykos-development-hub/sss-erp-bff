package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type SpendingDynamicInsert struct {
	ID        int             `json:"id"`
	BudgetID  int             `json:"budget_id"`
	UnitID    int             `json:"unit_id"`
	Username  string          `json:"username"`
	AccountID int             `json:"account_id"`
	January   decimal.Decimal `json:"january"`
	February  decimal.Decimal `json:"february"`
	March     decimal.Decimal `json:"march"`
	April     decimal.Decimal `json:"april"`
	May       decimal.Decimal `json:"may"`
	June      decimal.Decimal `json:"june"`
	July      decimal.Decimal `json:"july"`
	August    decimal.Decimal `json:"august"`
	September decimal.Decimal `json:"september"`
	October   decimal.Decimal `json:"october"`
	November  decimal.Decimal `json:"november"`
	December  decimal.Decimal `json:"december"`
	CreatedAt time.Time       `json:"created_at"`
}

type SpendingDynamic struct {
	ID                int             `json:"id"`
	BudgetID          int             `json:"budget_id"`
	UnitID            int             `json:"unit_id"`
	AccountID         int             `json:"account_id"`
	Actual            decimal.Decimal `json:"actual"`
	SpendingDynamicID int             `json:"spending_dynamic_id"`
	Username          string          `json:"username"`
	January           decimal.Decimal `json:"january"`
	February          decimal.Decimal `json:"february"`
	March             decimal.Decimal `json:"march"`
	April             decimal.Decimal `json:"april"`
	May               decimal.Decimal `json:"may"`
	June              decimal.Decimal `json:"june"`
	July              decimal.Decimal `json:"july"`
	August            decimal.Decimal `json:"august"`
	September         decimal.Decimal `json:"september"`
	October           decimal.Decimal `json:"october"`
	November          decimal.Decimal `json:"november"`
	December          decimal.Decimal `json:"december"`
	CreatedAt         time.Time       `json:"created_at"`
}
