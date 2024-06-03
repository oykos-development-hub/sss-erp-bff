package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type FlatRateType int

const (
	DissisionType FlatRateType = 1
	VerdictType   FlatRateType = 2
)

type FlatRateStatus int

const (
	UnpaidFlatRateStatus FlatRateStatus = 1
	PaidFlatRateStatus   FlatRateStatus = 2
	PartFlatRateStatus   FlatRateStatus = 3
)

type FlatRate struct {
	ID                     int              `json:"id"`
	FlatRateType           FlatRateType     `json:"flat_rate_type"`
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
	Status                 FlatRateStatus   `json:"status"`
	CourtCosts             *decimal.Decimal `json:"court_costs"`
	CourtAccountID         *int             `json:"court_account_id"`
	FlatRateDetailsDTO     *FlatRateDetails `json:"flat_rate_details"`
	File                   []int            `json:"file"`
	CreatedAt              time.Time        `json:"created_at"`
	UpdatedAt              time.Time        `json:"updated_at"`
}

type FlatRateDetails struct {
	AllPaymentAmount           decimal.Decimal `json:"all_payments_amount"`
	AmountGracePeriod          decimal.Decimal `json:"amount_grace_period"`
	AmountGracePeriodDueDate   time.Time       `json:"amount_grace_period_due_date"`
	AmountGracePeriodAvailable bool            `json:"amount_grace_period_available"`
	LeftToPayAmount            decimal.Decimal `json:"left_to_pay_amount"`
	CourtCostsPaid             decimal.Decimal `json:"court_costs_paid"`
	CourtCostsLeftToPayAmount  decimal.Decimal `json:"court_costs_left_to_pay_amount"`
}
