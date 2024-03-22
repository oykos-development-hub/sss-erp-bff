package dto

import (
	"bff/structs"
	"time"
)

type FlatRateType string

const (
	FinancialFlatRateDissisionType FlatRateType = "Rješenje"
	FinancialFlatRateVerdictType   FlatRateType = "Presuda"
)

type FlatRateStatus string

const (
	FinancialFlatRateUnpaidStatus FlatRateStatus = "Neplaćeno"
	FinancialFlatRatePaidStatus   FlatRateStatus = "Plaćeno"
	FinancialFlatRatePartStatus   FlatRateStatus = "Djelimično plaćeno"
)

type FlatRateResponseItem struct {
	ID                     int                      `json:"id"`
	FlatRateType           FlatRateType             `json:"flat_rate_type"`
	DecisionNumber         string                   `json:"decision_number"`
	DecisionDate           time.Time                `json:"decision_date"`
	Subject                string                   `json:"subject"`
	JMBG                   string                   `json:"jmbg"`
	Residence              string                   `json:"residence"`
	Amount                 float64                  `json:"amount"`
	PaymentReferenceNumber string                   `json:"payment_reference_number"`
	DebitReferenceNumber   string                   `json:"debit_reference_number"`
	Account                DropdownSimple           `json:"account"`
	ExecutionDate          time.Time                `json:"execution_date"`
	PaymentDeadlineDate    time.Time                `json:"payment_deadline_date"`
	Description            string                   `json:"description"`
	Status                 DropdownSimple           `json:"status"`
	CourtCosts             *float64                 `json:"court_costs"`
	CourtAccount           *DropdownSimple          `json:"court_account"`
	FlatRateDetailsDTO     *structs.FlatRateDetails `json:"flat_rate_details"`
	File                   []FileDropdownSimple     `json:"file"`
	CreatedAt              time.Time                `json:"created_at"`
	UpdatedAt              time.Time                `json:"updated_at"`
}

type GetFlatRateResponseMS struct {
	Data structs.FlatRate `json:"data"`
}

type GetFlatRateListResponseMS struct {
	Data  []structs.FlatRate `json:"data"`
	Total int                `json:"total"`
}

type GetFlatRateListInputMS struct {
	Subject        *string `json:"subject"`
	Page           *int    `json:"page"`
	Size           *int    `json:"size"`
	FilterByTypeID *int    `json:"flat_rate_type_id"`
	Search         *string `json:"search"`
}
