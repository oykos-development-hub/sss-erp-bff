package dto

import (
	"bff/structs"
	"sort"
	"time"

	"github.com/shopspring/decimal"
)

type BudgetRequestStatus string

const (
	BudgetRequestTakeActionStatus BudgetRequestStatus = "Obradi"
	BudgetRequestFinishedStatus   BudgetRequestStatus = "Obrađen"
	BudgetRequestAcceptedStatus   BudgetRequestStatus = "Odobreno"
	BudgetRequestOnHoldStatus     BudgetRequestStatus = "Na čekanju"
	BudgetRequestFilledStatus     BudgetRequestStatus = "Popunjeno"
)

func RequestStatusForOfficial(s structs.BudgetRequestStatus) BudgetRequestStatus {
	switch s {
	case structs.BudgetRequestSentOnReviewStatus:
		return BudgetRequestTakeActionStatus
	case structs.BudgetRequestAcceptedStatus:
		return BudgetRequestAcceptedStatus
	case structs.BudgetRequestWaitingForActual:
		return BudgetRequestTakeActionStatus
	case structs.BudgetRequestCompletedActualStatus:
		return BudgetRequestFinishedStatus
	default:
		return BudgetRequestOnHoldStatus
	}
}

func RequestStatusForManager(s structs.BudgetRequestStatus) BudgetRequestStatus {
	switch s {
	case structs.BudgetRequestSentStatus, structs.BudgetRequestRejectedStatus:
		return BudgetRequestTakeActionStatus
	case structs.BudgetRequestFilledStatus:
		return BudgetRequestFilledStatus
	case structs.BudgetRequestAcceptedStatus, structs.BudgetRequestWaitingForActual, structs.BudgetRequestCompletedActualStatus:
		return BudgetRequestAcceptedStatus
	default:
		return BudgetRequestOnHoldStatus
	}
}

type BudgetStatus string

const (
	BudgetCreatedStatus         BudgetStatus = "Kreiran"
	BudgetSentStatus            BudgetStatus = "Poslat"
	BudgetWaitForActualStatus   BudgetStatus = "Čekanje odobrenog budžeta"
	BudgetCompletedActualStatus BudgetStatus = "Odobreno"
)

func GetBudgetStatus(s structs.BudgetStatus) BudgetStatus {
	switch s {
	case structs.BudgetSentStatus:
		return BudgetSentStatus
	case structs.BudgetAcceptedStatus:
		return BudgetWaitForActualStatus
	case structs.BudgetCompletedActualStatus:
		return BudgetCompletedActualStatus
	default:
		return BudgetCreatedStatus
	}
}

type BudgetResponseItem struct {
	ID                          int                            `json:"id"`
	Year                        int                            `json:"year"`
	Source                      int                            `json:"source"`
	BudgetType                  int                            `json:"budget_type"`
	Status                      DropdownSimple                 `json:"status"`
	Limits                      []structs.FinancialBudgetLimit `json:"limits"`
	NumberOfRequestsForOfficial int                            `json:"number_of_requests"`
}

type GetBudgetResponseMS struct {
	Data structs.Budget `json:"data"`
}

type GetBudgetListResponseMS struct {
	Data []structs.Budget `json:"data"`
}

type GetBudgetListInputMS struct {
	Year       *int `json:"year"`
	BudgetType *int `json:"budget_type"`
	Status     *int `json:"budget_status"`
	ID         *int `json:"id"`
}

type GetBudgetRequestResponseMS struct {
	Data structs.BudgetRequest `json:"data"`
}

type GetBudgetRequestListResponseMS struct {
	Data []structs.BudgetRequest `json:"data"`
}

type GetBudgetRequestListInputMS struct {
	OrganizationUnitID *int                          `json:"organization_unit_id"`
	BudgetID           *int                          `json:"budget_id"`
	RequestType        *structs.RequestType          `json:"request_type"`
	RequestTypes       []structs.RequestType         `json:"request_types"`
	Statuses           []structs.BudgetRequestStatus `json:"statuses"`
	ParentID           *int                          `json:"parent_id"`
}

type RequestType string

const (
	CurrentFinancialRequestType  RequestType = "Tekući"
	DonationFinancialRequestType RequestType = "Donacija"
	NonFinancialRequestType      RequestType = "Nefinansijski"
)

func GetRequestType(r structs.RequestType) RequestType {
	switch r {
	case structs.RequestTypeCurrentFinancial:
		return CurrentFinancialRequestType
	case structs.RequestTypeDonationFinancial:
		return DonationFinancialRequestType
	case structs.RequestTypeNonFinancial:
		return NonFinancialRequestType
	}

	return ""
}

type BudgetRequestResponseItem struct {
	ID               int            `json:"id"`
	OrganizationUnit DropdownSimple `json:"organization_unit"`
	BudgetID         int            `json:"budget_id"`
	RequestType      RequestType    `json:"request_type"`
	Status           DropdownSimple `json:"status"`
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
	BudgetID int  `json:"budget_id"`
	UnitID   *int `json:"unit_id"`
}

type FinancialBudgetOverviewResponse struct {
	AccountVersion                 int                               `json:"account_version"`
	RequestID                      int                               `json:"request_id"`
	Status                         DropdownSimple                    `json:"status"`
	DonationBudgetStatus           DropdownSimple                    `json:"donation_status"`
	CurrentBudgetStatus            DropdownSimple                    `json:"current_status"`
	CurrentAccountsWithFilledData  []*AccountWithFilledFinanceBudget `json:"current_accounts"`
	CurrentRequestID               int                               `json:"current_request_id"`
	DonationAccountsWithFilledData []*AccountWithFilledFinanceBudget `json:"donation_accounts"`
	DonationRequestID              int                               `json:"donation_request_id"`
	DonationBudgetComment          string                            `json:"donation_budget_comment"`
	CurrentBudgetComment           string                            `json:"current_budget_comment"`
	OfficialComment                string                            `json:"official_comment"`
}

type CreateBudget struct {
	ID         int                            `json:"id"`
	Year       int                            `json:"year"`
	BudgetType int                            `json:"budget_type"`
	Limits     []structs.FinancialBudgetLimit `json:"limits"`
}

type FillActualFinanceBudgetInput struct {
	ID     int             `json:"id"`
	Actual decimal.Decimal `json:"actual"`
	Type   int             `json:"type"`
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

type CurrentBudgetAccounts struct {
	ID                  int                      `json:"id"`
	Title               string                   `json:"title"`
	ParentID            *int                     `json:"parent_id"`
	SerialNumber        string                   `json:"serial_number"`
	Children            []*CurrentBudgetAccounts `json:"children"`
	FilledFinanceBudget CurrentBudgetResponse    `json:"filled_data"`
}

type CurrentBudgetAccountsResponse struct {
	CurrentAccounts  []*CurrentBudgetAccounts `json:"current_accounts"`
	DonationAccounts []*CurrentBudgetAccounts `json:"donation_accounts"`
	BudgetID         int                      `json:"budget_id"`
	Version          int                      `json:"version"`
	Units            []DropdownOUSimple       `json:"units"`
}

type FilledFinancialBudgetInputMS struct {
	BudgetRequestID int   `json:"budget_request_id"`
	Accounts        []int `json:"accounts"`
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

type FinancialBudgetDetails struct {
	Version       int                        `json:"version"`
	LatestVersion int                        `json:"latest_version"`
	Accounts      []*AccountItemResponseItem `json:"accounts"`
}

type BudgetRequestOfficialOverview struct {
	Unit        DropdownOUSimple `json:"unit"`
	Status      string           `json:"status"`
	ReceiveDate *time.Time       `json:"receive_date"`
	Total       decimal.Decimal  `json:"total"`
	Limit       decimal.Decimal  `json:"limit"`
	Year        int              `json:"year"`
}

type BudgetRequestsDetails struct {
	Budget                    DropdownSimple                   `json:"budget"`
	Status                    DropdownSimple                   `json:"status"`
	RequestID                 int                              `json:"request_id"`
	FinancialBudgetDetails    *FinancialBudgetOverviewResponse `json:"financial"`
	NonFinancialBudgetDetails *NonFinancialBudgetResItem       `json:"non_financial"`
	Limit                     int                              `json:"limits"`
}

type CurrentBudgetResponse struct {
	ID             int             `json:"id"`
	BudgetID       int             `json:"budget_id"`
	UnitID         int             `json:"unit_id"`
	Account        DropdownSimple  `json:"account"`
	InititalActual decimal.Decimal `json:"initial_actual"`
	Actual         decimal.Decimal `json:"actual"`
	Balance        decimal.Decimal `json:"balance"`
	CurrentAmount  decimal.Decimal `json:"current_amount"`
	Type           int             `json:"type"`
	CreatedAt      time.Time       `json:"created_at"`
}
