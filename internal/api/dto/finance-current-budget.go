package dto

import (
	"bff/structs"
)

type GetCurrentBudgetResponseMS struct {
	Data *structs.CurrentBudget `json:"data"`
}
