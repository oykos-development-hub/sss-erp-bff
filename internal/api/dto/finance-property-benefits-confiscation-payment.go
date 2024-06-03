package dto

import (
	"bff/structs"
	"time"
)

type PropBenConfPaymentMethod string

const (
	FinancialPropBenConfPaymentMethodPayment    PropBenConfPaymentMethod = "Uplata"
	FinancialPropBenConfPaymentMethodForced     PropBenConfPaymentMethod = "Prinudna naplata"
	FinancialPropBenConfPaymentMethodCourtCosts PropBenConfPaymentMethod = "Sudski troškovi"
)

type PropBenConfPaymentStatus string

const (
	FinancialPropBenConfPaymentStatusPaid     PropBenConfPaymentStatus = "Uplaćeno"
	FinancialPropBenConfPaymentStatusCanceled PropBenConfPaymentStatus = "Stornirano"
	FinancialPropBenConfPaymentStatusReturned PropBenConfPaymentStatus = "Povraćaj"
)

type PropBenConfPaymentResponseItem struct {
	ID                     int            `json:"id,omitempty"`
	PropBenConfID          int            `json:"property_benefits_confiscation_id"`
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

type GetPropBenConfPaymentResponseMS struct {
	Data structs.PropBenConfPayment `json:"data"`
}

type GetPropBenConfPaymentListResponseMS struct {
	Data  []structs.PropBenConfPayment `json:"data"`
	Total int                          `json:"total"`
}

type GetPropBenConfPaymentListInputMS struct {
	Subject       *string `json:"subject"`
	Page          *int    `json:"page"`
	Size          *int    `json:"size"`
	PropBenConfID *int    `json:"property_benefits_confiscation_id"`
	Search        *string `json:"search"`
}
