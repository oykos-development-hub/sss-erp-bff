package dto

import (
	"bff/structs"
	"time"
)

type FeePaymentMethod string

const (
	FinancialFeePaymentMethodPayment FeePaymentMethod = "Uplata"
	FinancialFeePaymentMethodForced  FeePaymentMethod = "Prinudna naplata"
)

type FeePaymentStatus string

const (
	FinancialFeePaymentStatusPaid     FeePaymentStatus = "Uplaćeno"
	FinancialFeePaymentStatusCanceled FeePaymentStatus = "Stornirano"
	FinancialFeePaymentStatusReturned FeePaymentStatus = "Povraćaj"
)

type FeePaymentResponseItem struct {
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

type GetFeePaymentResponseMS struct {
	Data structs.FeePayment `json:"data"`
}

type GetFeePaymentListResponseMS struct {
	Data  []structs.FeePayment `json:"data"`
	Total int                  `json:"total"`
}

type GetFeePaymentListInputMS struct {
	Subject *string `json:"subject"`
	Page    *int    `json:"page"`
	Size    *int    `json:"size"`
	FeeID   *int    `json:"fee_id"`
	Search  *string `json:"search"`
}
