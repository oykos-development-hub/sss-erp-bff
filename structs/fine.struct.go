package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type FineActType int

const (
	DissisionActType FineActType = 1
	VerdictActType   FineActType = 2
)

type FineStatus int

const (
	UnpaidFineStatus FineStatus = 1
	PaidFineStatus   FineStatus = 2
	PartFineStatus   FineStatus = 3
)

type Fine struct {
	ID                     int              `json:"id"`
	ActType                FineActType      `json:"act_type"`
	DecisionNumber         string           `json:"decision_number"`
	DecisionDate           time.Time        `json:"decision_date"`
	Subject                string           `json:"subject"`
	JMBG                   string           `json:"jmbg"`
	Residence              string           `json:"residence"`
	Amount                 decimal.Decimal  `json:"amount"`
	PaymentReferenceNumber string           `json:"payment_reference_number"`
	DebitReferenceNumber   string           `json:"debit_reference_number"`
	AccountID              int              `json:"account_id"`
	ExecutionDate          time.Time        `json:"execution_date"`
	PaymentDeadlineDate    time.Time        `json:"payment_deadline_date"`
	Description            string           `json:"description"`
	Status                 FineStatus       `json:"status"`
	CourtCosts             *decimal.Decimal `json:"court_costs"`
	CourtAccountID         *int             `json:"court_account_id"`
	FineFeeDetailsDTO      *FineFeeDetails  `json:"fine_fee_details"`
	File                   []int            `json:"file"`
	CreatedAt              time.Time        `json:"created_at"`
	UpdatedAt              time.Time        `json:"updated_at"`
}

type FineFeeDetails struct {
	FeeAllPaymentAmount           decimal.Decimal `json:"fee_all_payments_amount"`
	FeeAmountGracePeriod          decimal.Decimal `json:"fee_amount_grace_period"`
	FeeAmountGracePeriodDueDate   time.Time       `json:"fee_amount_grace_period_due_date"`
	FeeAmountGracePeriodAvailable bool            `json:"fee_amount_grace_period_available"`
	FeeLeftToPayAmount            decimal.Decimal `json:"fee_left_to_pay_amount"`
	FeeCourtCostsPaid             decimal.Decimal `json:"fee_court_costs_paid"`
	FeeCourtCostsLeftToPayAmount  decimal.Decimal `json:"fee_court_costs_left_to_pay_amount"`
}
