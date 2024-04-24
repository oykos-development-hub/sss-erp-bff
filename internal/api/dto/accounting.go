package dto

import "time"

type ObligationForAccounting struct {
	InvoiceID                 *int      `json:"invoice_id"`
	AdditionalExpenseID       *int      `json:"additional_expense_id"`
	SalaryAdditionalExpenseID *int      `json:"salary_additional_expense_id"`
	Type                      string    `json:"type"`
	Title                     string    `json:"title"`
	TotalPrice                float64   `json:"total_price"`
	RemainPrice               float64   `json:"remain_price"`
	Status                    string    `json:"status"`
	CreatedAt                 time.Time `json:"created_at"`
}

type GetObligationsForAccountingResponseMS struct {
	Data  []ObligationForAccounting `json:"data"`
	Total int                       `json:"total"`
}
