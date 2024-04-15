package dto

import (
	"bff/structs"
	"time"
)

type DepositPaymentOrderResponse struct {
	ID                          int                                        `json:"id"`
	OrganizationUnit            DropdownSimple                             `json:"organization_unit"`
	CaseNumber                  string                                     `json:"case_number"`
	Supplier                    DropdownSimple                             `json:"supplier"`
	NetAmount                   float64                                    `json:"net_amount"`
	BankAccount                 string                                     `json:"bank_account"`
	DateOfPayment               time.Time                                  `json:"date_of_payment"`
	DateOfStatement             *time.Time                                 `json:"date_of_statement"`
	IDOfStatement               *string                                    `json:"id_of_statement"`
	Status                      string                                     `json:"status"`
	AdditionalExpenses          []DepositPaymentAdditionalExpensesResponse `json:"additional_expenses"`
	AdditionalExpensesForPaying []DepositPaymentAdditionalExpensesResponse `json:"additional_expenses_for_paying"`
	File                        FileDropdownSimple                         `json:"file"`
	CreatedAt                   time.Time                                  `json:"created_at"`
	UpdatedAt                   time.Time                                  `json:"updated_at"`
}

type DepositPaymentOrderFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	SortByTitle        *string `json:"sort_by_title"`
	Status             *string `json:"status"`
	Search             *string `json:"search"`
	CaseNumber         *string `json:"case_number"`
	SupplierID         *int    `json:"supplier_id"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
}

type GetDepositPaymentOrderResponseMS struct {
	Data structs.DepositPaymentOrder `json:"data"`
}

type GetDepositPaymentOrderListResponseMS struct {
	Data  []structs.DepositPaymentOrder `json:"data"`
	Total int                           `json:"total"`
}

type DepositPaymentAdditionalExpensesListInputMS struct {
	Page                 *int    `json:"page"`
	Size                 *int    `json:"size"`
	PaymentOrderID       *int    `json:"payment_order_id"`
	PayingPaymentOrderID *int    `json:"paying_payment_order_id"`
	SubjectID            *int    `json:"subject_id"`
	OrganizationUnitID   *int    `json:"organization_unit_id"`
	Year                 *int    `json:"year"`
	Status               *string `json:"status"`
	Search               *string `json:"search"`
}

type GetDepositPaymentAdditionalExpensesListResponseMS struct {
	Data  []structs.DepositPaymentAdditionalExpenses `json:"data"`
	Total int                                        `json:"total"`
}

type DepositPaymentAdditionalExpensesResponse struct {
	ID               int            `json:"id"`
	Title            string         `json:"title"`
	Account          DropdownSimple `json:"account"`
	Price            float32        `json:"price"`
	Subject          DropdownSimple `json:"subject"`
	BankAccount      string         `json:"bank_account"`
	PaymentOrder     DropdownSimple `json:"payment_order"`
	OrganizationUnit DropdownSimple `json:"organization_unit"`
	Status           string         `json:"status"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}
