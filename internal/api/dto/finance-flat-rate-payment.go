package dto

import (
	"bff/structs"
	"time"
)

type FlatRatePaymentMethod string

const (
	FinancialFlatRatePaymentMethodPayment    FlatRatePaymentMethod = "Uplata"
	FinancialFlatRatePaymentMethodForced     FlatRatePaymentMethod = "Prinudna naplata"
	FinancialFlatRatePaymentMethodCourtCosts FlatRatePaymentMethod = "Sudski troškovi"
)

type FlatRatePaymentStatus string

const (
	FinancialFlatRatePaymentStatusPaid     FlatRatePaymentStatus = "Uplaćeno"
	FinancialFlatRatePaymentStatusCanceled FlatRatePaymentStatus = "Stornirano"
	FinancialFlatRatePaymentStatusReturned FlatRatePaymentStatus = "Povraćaj"
)

type FlatRatePaymentResponseItem struct {
	ID                     int                   `json:"id,omitempty"`
	FlatRateID             int                   `json:"flat_rate_id"`
	PaymentMethod          FlatRatePaymentMethod `json:"payment_method"`
	Amount                 float64               `json:"amount"`
	PaymentDate            time.Time             `json:"payment_date"`
	PaymentDueDate         time.Time             `json:"payment_due_date"`
	ReceiptNumber          string                `json:"receipt_number"`
	PaymentReferenceNumber string                `json:"payment_reference_number"`
	DebitReferenceNumber   string                `json:"debit_reference_number"`
	Status                 FlatRatePaymentStatus `json:"status"`
	CreatedAt              time.Time             `json:"created_at,omitempty"`
	UpdatedAt              time.Time             `json:"updated_at"`
}

type GetFlatRatePaymentResponseMS struct {
	Data structs.FlatRatePayment `json:"data"`
}

type GetFlatRatePaymentListResponseMS struct {
	Data  []structs.FlatRatePayment `json:"data"`
	Total int                       `json:"total"`
}

type GetFlatRatePaymentListInputMS struct {
	Subject    *string `json:"subject"`
	Page       *int    `json:"page"`
	Size       *int    `json:"size"`
	FlatRateID *int    `json:"flat_rate_id"`
	Search     *string `json:"search"`
}
