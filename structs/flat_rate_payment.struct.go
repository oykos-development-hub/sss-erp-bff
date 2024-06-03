package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type FlatRatePaymentMethod int

const (
	PaymentFlatRatePeymentMethod    FlatRatePaymentMethod = 1
	ForcedFlatRatePeymentMethod     FlatRatePaymentMethod = 2
	CourtCostsFlatRatePeymentMethod FlatRatePaymentMethod = 3
)

type FlatRatePaymentStatus int

const (
	PaidFlatRatePeymentStatus      FlatRatePaymentStatus = 1
	CancelledFlatRatePeymentStatus FlatRatePaymentStatus = 2
	RetunedFlatRatePeymentStatus   FlatRatePaymentStatus = 3
)

type FlatRatePayment struct {
	ID                     int                   `json:"id,omitempty"`
	FlatRateID             int                   `json:"flat_rate_id"`
	PaymentMethod          FlatRatePaymentMethod `json:"payment_method"`
	Amount                 decimal.Decimal       `json:"amount"`
	PaymentDate            time.Time             `json:"payment_date"`
	PaymentDueDate         time.Time             `json:"payment_due_date"`
	ReceiptNumber          string                `json:"receipt_number"`
	PaymentReferenceNumber string                `json:"payment_reference_number"`
	DebitReferenceNumber   string                `json:"debit_reference_number"`
	Status                 FlatRatePaymentStatus `json:"status"`
	CreatedAt              time.Time             `json:"created_at,omitempty"`
	UpdatedAt              time.Time             `json:"updated_at"`
}
