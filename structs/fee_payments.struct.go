package structs

import "time"

type FeePaymentMethod int

const (
	PaymentFeePeymentMethod FeePaymentMethod = 1
	ForcedFeePeymentMethod  FeePaymentMethod = 2
)

type FeePaymentStatus int

const (
	PaidFeePeymentStatus      FeePaymentStatus = 1
	CancelledFeePeymentStatus FeePaymentStatus = 2
	RetunedFeePeymentStatus   FeePaymentStatus = 3
)

type FeePayment struct {
	ID                     int              `json:"id,omitempty"`
	FeeID                  int              `json:"fee_id"`
	PaymentMethod          FeePaymentMethod `json:"payment_method"`
	Amount                 float64          `json:"amount"`
	PaymentDate            time.Time        `json:"payment_date"`
	PaymentDueDate         time.Time        `json:"payment_due_date"`
	ReceiptNumber          string           `json:"receipt_number"`
	PaymentReferenceNumber string           `json:"payment_reference_number"`
	DebitReferenceNumber   string           `json:"debit_reference_number"`
	Status                 FeePaymentStatus `json:"status"`
	CreatedAt              time.Time        `json:"created_at,omitempty"`
	UpdatedAt              time.Time        `json:"updated_at"`
}
