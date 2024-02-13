package dto

import (
	"bff/structs"
	"sort"
)

type FinancialBudgetStatus string

const (
	FinancialBudgetTakeActionStatus FinancialBudgetStatus = "Obradi"
	FinancialBudgetFinishedStatus   FinancialBudgetStatus = "Obrađen"
)

type BudgetStatus string

const (
	BudgetCreatedStatus BudgetStatus = "Kreiran"
	BudgetSentStatus    BudgetStatus = "Poslat"
	BudgetClosedStatus  BudgetStatus = "Zatvoren"
)

type BudgetResponseItem struct {
	ID         int          `json:"id"`
	Year       int          `json:"year"`
	Source     int          `json:"source"`
	BudgetType int          `json:"budget_type"`
	Status     BudgetStatus `json:"status"`
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

type GetBudgetRequestResponseMS struct {
	Data structs.BudgetRequest `json:"data"`
}

type GetBudgetRequestListResponseMS struct {
	Data []structs.BudgetRequest `json:"data"`
}

type GetBudgetRequestListInputMS struct {
	OrganizationUnitID *int                 `json:"organization_unit_id"`
	BudgetID           int                  `json:"budget_id"`
	RequestType        *structs.RequestType `json:"request_type"`
}

type FinancialBudgetResponseItem struct {
	ID             int    `json:"id"`
	AccountVersion int    `json:"account_version"`
	BudgetID       int    `json:"budget_id"`
	Status         string `json:"status"`
}

type GetFinancialBudgetResponseMS struct {
	Data structs.FinancialBudget `json:"data"`
}

type GetFinancialBudgetLimitResponseMS struct {
	Data structs.FinancialBudgetLimit `json:"data"`
}

type FinancialBudgetOverviewResponse struct {
	AccountVersion         int                               `json:"account_version"`
	Status                 FinancialBudgetStatus             `json:"status"`
	AccountsWithFilledData []*AccountWithFilledFinanceBudget `json:"accounts"`
}

type CreateBudget struct {
	ID         int                    `json:"id"`
	Year       int                    `json:"year"`
	BudgetType int                    `json:"budget_type"`
	Limits     []FinancialBudgetLimit `json:"limits"`
}

type FinancialBudgetLimit struct {
	ID                 int `json:"id"`
	Limit              int `json:"limit"`
	OrganizationUnitID int `json:"organization_unit_id"`
	FinancialBudgetID  int `json:"financial_budget_id"`
}

type GetFilledFinancialBudgetResponseItemMS struct {
	Data structs.FilledFinanceBudget `json:"data"`
}

type GetFilledFinancialBudgetResponseMS struct {
	Data []structs.FilledFinanceBudget `json:"data"`
}

type AccountWithFilledFinanceBudgetResponseList struct {
	Data []*AccountWithFilledFinanceBudget
}

type AccountWithFilledFinanceBudget struct {
	ID                  int                               `json:"id"`
	Title               string                            `json:"title"`
	ParentID            *int                              `json:"parent_id"`
	SerialNumber        string                            `json:"serial_number"`
	Children            []*AccountWithFilledFinanceBudget `json:"children"`
	FilledFinanceBudget *structs.FilledFinanceBudget      `json:"filled_data"`
}

type FilledFinancialBudgetInputMS struct {
	OrganizationUnitID int `json:"organization_unit_id"`
	FinancialBudgetID  int `json:"financial_budget_id"`
}

func (a *AccountWithFilledFinanceBudgetResponseList) CreateTree() []*AccountWithFilledFinanceBudget {
	mappedNodes := make(map[int]*AccountWithFilledFinanceBudget, len(a.Data))
	var rootNodes []*AccountWithFilledFinanceBudget

	for _, node := range a.Data {
		mappedNodes[node.ID] = node
		if node.ParentID == nil {
			rootNodes = append(rootNodes, node)
		}
	}

	// Populate children for each node
	for _, node := range a.Data {
		if node.ParentID != nil {
			if parentNode, exists := mappedNodes[*node.ParentID]; exists {
				parentNode.Children = append(parentNode.Children, node)
			}
		}
	}

	sort.Slice(rootNodes, func(i, j int) bool {
		return rootNodes[i].SerialNumber < rootNodes[j].SerialNumber
	})

	return rootNodes
}
