package dto

import (
	"bff/structs"
)

type BudgetResponseItem struct {
	ID         int    `json:"id"`
	Year       int    `json:"year"`
	Source     int    `json:"source"`
	BudgetType int    `json:"budget_type"`
	Status     string `json:"status"`
}

type GetBudgetResponseMS struct {
	Data structs.Budget `json:"data"`
}

type GetBudgetListResponseMS struct {
	Data []structs.Budget `json:"data"`
}

type GetBudgetListInputMS struct {
	Year       *int    `json:"year"`
	BudgetType *int    `json:"budget_type"`
	Status     *string `json:"status"`
	ID         *int    `json:"id"`
}

type GetFinancialBudgetResponseMS struct {
	Data structs.FinancialBudget `json:"data"`
}

type FinancialBudgetOverviewResponse struct {
	AccountVersion int                        `json:"account_version"`
	Accounts       []*AccountItemResponseItem `json:"accounts"`
}
