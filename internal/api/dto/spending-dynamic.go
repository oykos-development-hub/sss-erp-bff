package dto

import (
	"bff/structs"

	"github.com/shopspring/decimal"
)

type GetSpendingDynamicResponseMS struct {
	Data structs.SpendingDynamic `json:"data"`
}

type GetSpendingDynamicActualResponseMS struct {
	Data decimal.NullDecimal `json:"data"`
}
