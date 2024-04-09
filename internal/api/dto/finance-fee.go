package dto

import (
	"bff/structs"
	"time"
)

type FeeType string

const (
	LawsuitFeeType  FeeType = "Tužba"
	JudgmentFeeType FeeType = "Presuda"
)

type FeeStatus string

const (
	FinancialFeeUnpaidFeeStatus FeeStatus = "Neplaćeno"
	FinancialFeePaidFeeStatus   FeeStatus = "Plaćeno"
	FinancialFeePartFeeStatus   FeeStatus = "Djelimično plaćeno"
)

type FeeSubcategory string

const (
	CopyingFeeSubcategory FeeSubcategory = "Kopiranje"
)

type FeeResponseItem struct {
	ID                     int                  `json:"id"`
	FeeType                DropdownSimple       `json:"fee_type"`
	FeeSubcategory         DropdownSimple       `json:"fee_subcategory"`
	DecisionNumber         string               `json:"decision_number"`
	DecisionDate           time.Time            `json:"decision_date"`
	Subject                string               `json:"subject"`
	JMBG                   string               `json:"jmbg"`
	Residence              string               `json:"residence"`
	Amount                 float64              `json:"amount"`
	PaymentReferenceNumber string               `json:"payment_reference_number"`
	DebitReferenceNumber   string               `json:"debit_reference_number"`
	ExecutionDate          time.Time            `json:"execution_date"`
	PaymentDeadlineDate    time.Time            `json:"payment_deadline_date"`
	Description            string               `json:"description"`
	Status                 DropdownSimple       `json:"status"`
	CourtAccount           *DropdownSimple      `json:"court_account"`
	FeeDetails             *structs.FeeDetails  `json:"fee_details"`
	File                   []FileDropdownSimple `json:"file"`
	CreatedAt              time.Time            `json:"created_at"`
	UpdatedAt              time.Time            `json:"updated_at"`
}

type GetFeeResponseMS struct {
	Data structs.Fee `json:"data"`
}

type GetFeeListResponseMS struct {
	Data  []structs.Fee `json:"data"`
	Total int           `json:"total"`
}

type GetFeeListInputMS struct {
	Subject               *string `json:"subject"`
	Page                  *int    `json:"page"`
	Size                  *int    `json:"size"`
	FilterByFeeTypeID     *int    `json:"fee_type_id"`
	FilterBySubcategoryID *int    `json:"fee_subcategory_id"`
	Search                *string `json:"search"`
}
