package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type PropBenConfPaymentMethod int

const (
	PaymentPropBenConfPeymentMethod    PropBenConfPaymentMethod = 1
	ForcedPropBenConfPeymentMethod     PropBenConfPaymentMethod = 2
	CourtCostsPropBenConfPeymentMethod PropBenConfPaymentMethod = 3
)

type PropBenConfPaymentStatus int

const (
	PaidPropBenConfPeymentStatus      PropBenConfPaymentStatus = 1
	CancelledPropBenConfPeymentStatus PropBenConfPaymentStatus = 2
	RetunedPropBenConfPeymentStatus   PropBenConfPaymentStatus = 3
)

type PropBenConfPayment struct {
	ID                     int                      `json:"id,omitempty"`
	PropBenConfID          int                      `json:"property_benefits_confiscation_id"`
	PaymentMethod          PropBenConfPaymentMethod `json:"payment_method"`
	Amount                 decimal.Decimal          `json:"amount"`
	PaymentDate            time.Time                `json:"payment_date"`
	PaymentDueDate         time.Time                `json:"payment_due_date"`
	ReceiptNumber          string                   `json:"receipt_number"`
	PaymentReferenceNumber string                   `json:"payment_reference_number"`
	DebitReferenceNumber   string                   `json:"debit_reference_number"`
	Status                 PropBenConfPaymentStatus `json:"status"`
	CreatedAt              time.Time                `json:"created_at,omitempty"`
	UpdatedAt              time.Time                `json:"updated_at"`
}
