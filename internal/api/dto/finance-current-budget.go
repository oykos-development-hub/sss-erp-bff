package dto

import (
	"bff/structs"
)

type GetCurrentBudgetResponseMS struct {
	Data *structs.CurrentBudget `json:"data"`
}

type GetCurrentBudgetListResponseMS struct {
	Data []structs.CurrentBudget `json:"data"`
}

type GetCurrentBudgetUnitListResponseMS struct {
	Data []int `json:"data"`
}
