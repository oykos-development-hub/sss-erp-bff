package dto

import (
	"bff/structs"
	"sort"
)

type BudgetRequestStatus string

const (
	FinancialBudgetTakeActionStatus BudgetRequestStatus = "Obradi"
	FinancialBudgetFinishedStatus   BudgetRequestStatus = "Obrađen"
)

type BudgetStatus string

const (
	BudgetCreatedStatus BudgetStatus = "Kreiran"
	BudgetClosedStatus  BudgetStatus = "Završen"

	ManagerBudgetProcessStatus  BudgetStatus = "Obradi"
	ManagerBudgetOnReviewStatus BudgetStatus = "Na čekanju"
	ManagerBudgetClosedStatus   BudgetStatus = "Završen"

	OfficialBudgetSentStatus     BudgetStatus = "Poslat"
	OfficialBudgetOnReviewStatus BudgetStatus = "Obradi"
)

type BudgetResponseItem struct {
	ID         int                            `json:"id"`
	Year       int                            `json:"year"`
	Source     int                            `json:"source"`
	BudgetType int                            `json:"budget_type"`
	Status     BudgetStatus                   `json:"status"`
	Limits     []structs.FinancialBudgetLimit `json:"limits"`
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

type BudgetRequestResponseItem struct {
	ID               int                 `json:"id"`
	OrganizationUnit DropdownSimple      `json:"organization_unit_id"`
	BudgetID         int                 `json:"budget_id"`
	RequestType      structs.RequestType `json:"request_type"`
	Status           BudgetRequestStatus `json:"status"`
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

type GetFinancialBudgetLimitListResponseMS struct {
	Data []structs.FinancialBudgetLimit `json:"data"`
}

type GetFinancialBudgetListInputMS struct {
	BudgetID int `json:"budget_id"`
}

type FinancialBudgetOverviewResponse struct {
	AccountVersion                 int                               `json:"account_version"`
	RequestID                      int                               `json:"request_id"`
	Status                         BudgetRequestStatus               `json:"status"`
	CurrentAccountsWithFilledData  []*AccountWithFilledFinanceBudget `json:"current_accounts"`
	CurrentRequestID               int                               `json:"current_request_id"`
	DonationAccountsWithFilledData []*AccountWithFilledFinanceBudget `json:"donation_accounts"`
	DonationRequestID              int                               `json:"donation_request_id"`
}

type CreateBudget struct {
	ID         int                            `json:"id"`
	Year       int                            `json:"year"`
	BudgetType int                            `json:"budget_type"`
	Limits     []structs.FinancialBudgetLimit `json:"limits"`
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
	BudgetRequestID int `json:"budget_request_id"`
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
