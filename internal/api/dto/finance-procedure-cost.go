package dto

import (
	"bff/structs"
	"time"
)

type ProcedureCostType string

const (
	FinancialProcedureCostDissisionType ProcedureCostType = "Rješenje"
	FinancialProcedureCostVerdictType   ProcedureCostType = "Presuda"
)

type ProcedureCostStatus string

const (
	FinancialProcedureCostUnpaidStatus ProcedureCostStatus = "Neplaćeno"
	FinancialProcedureCostPaidStatus   ProcedureCostStatus = "Plaćeno"
	FinancialProcedureCostPartStatus   ProcedureCostStatus = "Djelimično plaćeno"
)

type ProcedureCostResponseItem struct {
	ID                      int                           `json:"id"`
	ActType                 DropdownSimple                `json:"procedure_cost_type"`
	DecisionNumber          string                        `json:"decision_number"`
	DecisionDate            time.Time                     `json:"decision_date"`
	OrganizationUnit        DropdownSimple                `json:"organization_unit"`
	Subject                 string                        `json:"subject"`
	JMBG                    string                        `json:"jmbg"`
	Residence               string                        `json:"residence"`
	Amount                  float64                       `json:"amount"`
	PaymentReferenceNumber  string                        `json:"payment_reference_number"`
	DebitReferenceNumber    string                        `json:"debit_reference_number"`
	Account                 DropdownSimple                `json:"account"`
	ExecutionDate           time.Time                     `json:"execution_date"`
	PaymentDeadlineDate     time.Time                     `json:"payment_deadline_date"`
	Description             string                        `json:"description"`
	Status                  DropdownSimple                `json:"status"`
	CourtCosts              *float64                      `json:"court_costs"`
	CourtAccount            *DropdownSimple               `json:"court_account"`
	ProcedureCostDetailsDTO *structs.ProcedureCostDetails `json:"procedure_cost_details"`
	File                    []FileDropdownSimple          `json:"file"`
	CreatedAt               time.Time                     `json:"created_at"`
	UpdatedAt               time.Time                     `json:"updated_at"`
}

type GetProcedureCostResponseMS struct {
	Data structs.ProcedureCost `json:"data"`
}

type GetProcedureCostListResponseMS struct {
	Data  []structs.ProcedureCost `json:"data"`
	Total int                     `json:"total"`
}

type GetProcedureCostListInputMS struct {
	Subject                     *string `json:"subject"`
	Page                        *int    `json:"page"`
	Size                        *int    `json:"size"`
	FilterByProcedureCostTypeID *int    `json:"procedure_cost_type_id"`
	Search                      *string `json:"search"`
	OrganizationUnitID          *int    `json:"organization_unit_id"`
}
