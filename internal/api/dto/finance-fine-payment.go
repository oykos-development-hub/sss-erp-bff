package dto

import (
	"bff/structs"
	"time"
)

type FinePaymentMethod string

const (
	FinancialFinePaymentMethodPayment    FinePaymentMethod = "Uplata"
	FinancialFinePaymentMethodForced     FinePaymentMethod = "Prinudna naplata"
	FinancialFinePaymentMethodCourtCosts FinePaymentMethod = "Sudski troškovi"
)

type FinePaymentStatus string

const (
	FinancialFinePaymentStatusPaid     FinePaymentStatus = "Uplaćeno"
	FinancialFinePaymentStatusCanceled FinePaymentStatus = "Stornirano"
	FinancialFinePaymentStatusReturned FinePaymentStatus = "Povraćaj"
)

type FinePaymentResponseItem struct {
	ID                     int            `json:"id,omitempty"`
	FineID                 int            `json:"fine_id"`
	PaymentMethod          DropdownSimple `json:"payment_method"`
	Amount                 float64        `json:"amount"`
	PaymentDate            time.Time      `json:"payment_date"`
	PaymentDueDate         time.Time      `json:"payment_due_date"`
	ReceiptNumber          string         `json:"receipt_number"`
	PaymentReferenceNumber string         `json:"payment_reference_number"`
	DebitReferenceNumber   string         `json:"debit_reference_number"`
	Status                 DropdownSimple `json:"status"`
	CreatedAt              time.Time      `json:"created_at,omitempty"`
	UpdatedAt              time.Time      `json:"updated_at"`
}

type GetFinePaymentResponseMS struct {
	Data structs.FinePayment `json:"data"`
}

type GetFinePaymentListResponseMS struct {
	Data  []structs.FinePayment `json:"data"`
	Total int                   `json:"total"`
}

type GetFinePaymentListInputMS struct {
	Subject *string `json:"subject"`
	Page    *int    `json:"page"`
	Size    *int    `json:"size"`
	FineID  *int    `json:"fine_id"`
	Search  *string `json:"search"`
}
