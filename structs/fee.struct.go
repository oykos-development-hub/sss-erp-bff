package structs

import "time"

type FeeType int

const (
	LawsuitFeeType  FeeType = 1
	JudgmentFeeType FeeType = 2
)

type FeeStatus int

const (
	UnpaidFeeStatus FeeStatus = 1
	PaidFeeStatus   FeeStatus = 2
	PartFeeStatus   FeeStatus = 3
)

type FeeSubcategory int

const (
	CopyingFeeSubcategory FeeSubcategory = 1
)

type Fee struct {
	ID                     int            `json:"id,omitempty"`
	FeeType                FeeType        `json:"fee_type"`
	FeeSubcategory         FeeSubcategory `json:"fee_subcategory"`
	OrganizationUnitID     int            `json:"organization_unit_id"`
	DecisionNumber         string         `json:"decision_number"`
	DecisionDate           time.Time      `json:"decision_date"`
	Subject                string         `json:"subject"`
	JMBG                   string         `json:"jmbg"`
	Amount                 float64        `json:"amount"`
	PaymentReferenceNumber string         `json:"payment_reference_number"`
	DebitReferenceNumber   string         `json:"debit_reference_number"`
	ExecutionDate          time.Time      `json:"execution_date"`
	PaymentDeadlineDate    time.Time      `json:"payment_deadline_date"`
	Description            string         `json:"description"`
	Status                 FeeStatus      `json:"status"`
	CourtAccountID         *int           `json:"court_account"`
	FeeDetails             *FeeDetails    `json:"fee_details"`
	File                   []int          `json:"file"`
	Residence              string         `json:"residence"`
	CreatedAt              time.Time      `json:"created_at,omitempty"`
	UpdatedAt              time.Time      `json:"updated_at"`
}

type FeeDetails struct {
	FeeAllPaymentAmount float64 `json:"fee_all_payments_amount"`
	FeeLeftToPayAmount  float64 `json:"fee_left_to_pay_amount"`
}
