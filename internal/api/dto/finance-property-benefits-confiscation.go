package dto

import (
	"bff/structs"
	"time"
)

type PropBenConfType string

const (
	FinancialPropBenConfDissisionType PropBenConfType = "Rješenje"
	FinancialPropBenConfVerdictType   PropBenConfType = "Presuda"
)

type PropBenConfStatus string

const (
	FinancialPropBenConfUnpaidStatus PropBenConfStatus = "Neplaćeno"
	FinancialPropBenConfPaidStatus   PropBenConfStatus = "Plaćeno"
	FinancialPropBenConfPartStatus   PropBenConfStatus = "Djelimično plaćeno"
)

type PropBenConfResponseItem struct {
	ID                     int                         `json:"id"`
	PropBenConfType        DropdownSimple              `json:"property_benefits_confiscation_type"`
	OrganizationUnit       DropdownSimple              `json:"organization_unit"`
	DecisionNumber         string                      `json:"decision_number"`
	DecisionDate           time.Time                   `json:"decision_date"`
	Subject                string                      `json:"subject"`
	JMBG                   string                      `json:"jmbg"`
	Residence              string                      `json:"residence"`
	Amount                 float64                     `json:"amount"`
	PaymentReferenceNumber string                      `json:"payment_reference_number"`
	DebitReferenceNumber   string                      `json:"debit_reference_number"`
	Account                DropdownSimple              `json:"account"`
	ExecutionDate          time.Time                   `json:"execution_date"`
	PaymentDeadlineDate    time.Time                   `json:"payment_deadline_date"`
	Description            string                      `json:"description"`
	Status                 DropdownSimple              `json:"status"`
	CourtCosts             *float64                    `json:"court_costs"`
	CourtAccount           *DropdownSimple             `json:"court_account"`
	PropBenConfDetailsDTO  *structs.PropBenConfDetails `json:"property_benefits_confiscation_details"`
	File                   []FileDropdownSimple        `json:"file"`
	CreatedAt              time.Time                   `json:"created_at"`
	UpdatedAt              time.Time                   `json:"updated_at"`
}

type GetPropBenConfResponseMS struct {
	Data structs.PropBenConf `json:"data"`
}

type GetPropBenConfListResponseMS struct {
	Data  []structs.PropBenConf `json:"data"`
	Total int                   `json:"total"`
}

type GetPropBenConfListInputMS struct {
	Subject                   *string `json:"subject"`
	Page                      *int    `json:"page"`
	Size                      *int    `json:"size"`
	FilterByPropBenConfTypeID *int    `json:"property_benefits_confiscation_type_id"`
	Search                    *string `json:"search"`
	OrganizationUnitID        *int    `json:"organization_unit_id"`
}
