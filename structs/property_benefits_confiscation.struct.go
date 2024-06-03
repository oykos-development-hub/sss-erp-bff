package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type PropBenConfType int

const (
	DissisionPropBenConfType PropBenConfType = 1
	VerdictPropBenConfType   PropBenConfType = 2
)

type PropBenConfStatus int

const (
	UnpaidPropBenConfStatus PropBenConfStatus = 1
	PaidPropBenConfStatus   PropBenConfStatus = 2
	PartPropBenConfStatus   PropBenConfStatus = 3
)

type PropBenConf struct {
	ID                     int                 `json:"id"`
	PropBenConfType        PropBenConfType     `json:"property_benefits_confiscation_type"`
	DecisionNumber         string              `json:"decision_number"`
	DecisionDate           time.Time           `json:"decision_date"`
	Subject                string              `json:"subject"`
	JMBG                   string              `json:"jmbg"`
	Residence              string              `json:"residence"`
	Amount                 decimal.Decimal     `json:"amount"`
	PaymentReferenceNumber string              `json:"payment_reference_number"`
	DebitReferenceNumber   string              `json:"debit_reference_number"`
	AccountID              int                 `json:"account_id"`
	ExecutionDate          time.Time           `json:"execution_date"`
	PaymentDeadlineDate    time.Time           `json:"payment_deadline_date"`
	Description            string              `json:"description"`
	Status                 PropBenConfStatus   `json:"status"`
	CourtCosts             *decimal.Decimal    `json:"court_costs"`
	CourtAccountID         *int                `json:"court_account_id"`
	PropBenConfDetailsDTO  *PropBenConfDetails `json:"property_benefits_confiscation_details"`
	File                   []int               `json:"file"`
	CreatedAt              time.Time           `json:"created_at"`
	UpdatedAt              time.Time           `json:"updated_at"`
}

type PropBenConfDetails struct {
	AllPaymentAmount           decimal.Decimal `json:"all_payments_amount"`
	AmountGracePeriod          decimal.Decimal `json:"amount_grace_period"`
	AmountGracePeriodDueDate   time.Time       `json:"amount_grace_period_due_date"`
	AmountGracePeriodAvailable bool            `json:"amount_grace_period_available"`
	LeftToPayAmount            decimal.Decimal `json:"left_to_pay_amount"`
	CourtCostsPaid             decimal.Decimal `json:"court_costs_paid"`
	CourtCostsLeftToPayAmount  decimal.Decimal `json:"court_costs_left_to_pay_amount"`
}
