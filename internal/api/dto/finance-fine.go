package dto

import (
	"bff/structs"
	"time"
)

type FineActType string

const (
	FinancialFineDissisionActType FineActType = "Rješenje"
	FinancialFineVerdictActType   FineActType = "Presuda"
)

type FineStatus string

const (
	FinancialFineUnpaidFineStatus FineStatus = "Neplaćeno"
	FinancialFinePaidFineStatus   FineStatus = "Plaćeno"
	FinancialFinePartFineStatus   FineStatus = "Djelimično plaćeno"
)

type FineResponseItem struct {
	ID                     int                     `json:"id"`
	ActType                DropdownSimple          `json:"act_type"`
	DecisionNumber         string                  `json:"decision_number"`
	DecisionDate           time.Time               `json:"decision_date"`
	Subject                string                  `json:"subject"`
	JMBG                   string                  `json:"jmbg"`
	Residence              string                  `json:"residence"`
	Amount                 float64                 `json:"amount"`
	PaymentReferenceNumber string                  `json:"payment_reference_number"`
	DebitReferenceNumber   string                  `json:"debit_reference_number"`
	Account                DropdownSimple          `json:"account"`
	ExecutionDate          time.Time               `json:"execution_date"`
	PaymentDeadlineDate    time.Time               `json:"payment_deadline_date"`
	Description            string                  `json:"description"`
	Status                 DropdownSimple          `json:"status"`
	CourtCosts             *float64                `json:"court_costs"`
	CourtAccount           *DropdownSimple         `json:"court_account"`
	FineFeeDetailsDTO      *structs.FineFeeDetails `json:"fine_fee_details"`
	File                   []FileDropdownSimple    `json:"file"`
	CreatedAt              time.Time               `json:"created_at"`
	UpdatedAt              time.Time               `json:"updated_at"`
}

type GetFineResponseMS struct {
	Data structs.Fine `json:"data"`
}

type GetFineListResponseMS struct {
	Data  []structs.Fine `json:"data"`
	Total int            `json:"total"`
}

type GetFineListInputMS struct {
	Subject           *string `json:"subject"`
	Page              *int    `json:"page"`
	Size              *int    `json:"size"`
	FilterByActTypeID *int    `json:"act_type_id"`
	Search            *string `json:"search"`
}
