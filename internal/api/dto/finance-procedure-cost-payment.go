package dto

import (
	"bff/structs"
	"time"
)

type ProcedureCostPaymentMethod string

const (
	FinancialProcedureCostPaymentMethodPayment    ProcedureCostPaymentMethod = "Uplata"
	FinancialProcedureCostPaymentMethodForced     ProcedureCostPaymentMethod = "Prinudna naplata"
	FinancialProcedureCostPaymentMethodCourtCosts ProcedureCostPaymentMethod = "Sudski troškovi"
)

type ProcedureCostPaymentStatus string

const (
	FinancialProcedureCostPaymentStatusPaid     ProcedureCostPaymentStatus = "Uplaćeno"
	FinancialProcedureCostPaymentStatusCanceled ProcedureCostPaymentStatus = "Stornirano"
	FinancialProcedureCostPaymentStatusReturned ProcedureCostPaymentStatus = "Povraćaj"
)

type ProcedureCostPaymentResponseItem struct {
	ID                     int                        `json:"id,omitempty"`
	ProcedureCostID        int                        `json:"procedure_cost_id"`
	PaymentMethod          ProcedureCostPaymentMethod `json:"payment_method"`
	Amount                 float64                    `json:"amount"`
	PaymentDate            time.Time                  `json:"payment_date"`
	PaymentDueDate         time.Time                  `json:"payment_due_date"`
	ReceiptNumber          string                     `json:"receipt_number"`
	PaymentReferenceNumber string                     `json:"payment_reference_number"`
	DebitReferenceNumber   string                     `json:"debit_reference_number"`
	Status                 ProcedureCostPaymentStatus `json:"status"`
	CreatedAt              time.Time                  `json:"created_at,omitempty"`
	UpdatedAt              time.Time                  `json:"updated_at"`
}

type GetProcedureCostPaymentResponseMS struct {
	Data structs.ProcedureCostPayment `json:"data"`
}

type GetProcedureCostPaymentListResponseMS struct {
	Data  []structs.ProcedureCostPayment `json:"data"`
	Total int                            `json:"total"`
}

type GetProcedureCostPaymentListInputMS struct {
	Subject         *string `json:"subject"`
	Page            *int    `json:"page"`
	Size            *int    `json:"size"`
	ProcedureCostID *int    `json:"procedure_cost_id"`
	Search          *string `json:"search"`
}
