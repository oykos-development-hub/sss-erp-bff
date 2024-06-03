package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type FinePaymentMethod int

const (
	PaymentFinePeymentMethod    FinePaymentMethod = 1
	ForcedFinePeymentMethod     FinePaymentMethod = 2
	CourtCostsFinePeymentMethod FinePaymentMethod = 3
)

type FinePaymentStatus int

const (
	PaidFinePeymentStatus      FinePaymentStatus = 1
	CancelledFinePeymentStatus FinePaymentStatus = 2
	RetunedFinePeymentStatus   FinePaymentStatus = 3
)

type FinePayment struct {
	ID                     int               `json:"id,omitempty"`
	FineID                 int               `json:"fine_id"`
	PaymentMethod          FinePaymentMethod `json:"payment_method"`
	Amount                 decimal.Decimal   `json:"amount"`
	PaymentDate            time.Time         `json:"payment_date"`
	PaymentDueDate         time.Time         `json:"payment_due_date"`
	ReceiptNumber          string            `json:"receipt_number"`
	PaymentReferenceNumber string            `json:"payment_reference_number"`
	DebitReferenceNumber   string            `json:"debit_reference_number"`
	Status                 FinePaymentStatus `json:"status"`
	CreatedAt              time.Time         `json:"created_at,omitempty"`
	UpdatedAt              time.Time         `json:"updated_at"`
}
