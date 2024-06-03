package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProcedureCostPaymentMethod int

const (
	PaymentProcedureCostPeymentMethod    ProcedureCostPaymentMethod = 1
	ForcedProcedureCostPeymentMethod     ProcedureCostPaymentMethod = 2
	CourtCostsProcedureCostPeymentMethod ProcedureCostPaymentMethod = 3
)

type ProcedureCostPaymentStatus int

const (
	PaidProcedureCostPeymentStatus      ProcedureCostPaymentStatus = 1
	CancelledProcedureCostPeymentStatus ProcedureCostPaymentStatus = 2
	RetunedProcedureCostPeymentStatus   ProcedureCostPaymentStatus = 3
)

type ProcedureCostPayment struct {
	ID                     int                        `json:"id,omitempty"`
	ProcedureCostID        int                        `json:"procedure_cost_id"`
	PaymentMethod          ProcedureCostPaymentMethod `json:"payment_method"`
	Amount                 decimal.Decimal            `json:"amount"`
	PaymentDate            time.Time                  `json:"payment_date"`
	PaymentDueDate         time.Time                  `json:"payment_due_date"`
	ReceiptNumber          string                     `json:"receipt_number"`
	PaymentReferenceNumber string                     `json:"payment_reference_number"`
	DebitReferenceNumber   string                     `json:"debit_reference_number"`
	Status                 ProcedureCostPaymentStatus `json:"status"`
	CreatedAt              time.Time                  `json:"created_at,omitempty"`
	UpdatedAt              time.Time                  `json:"updated_at"`
}
