package dto

import "time"

type ObligationForAccounting struct {
	InvoiceID *int      `json:"invoice_id"`
	SalaryID  *int      `json:"salary_id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type GetObligationsForAccountingResponseMS struct {
	Data  []ObligationForAccounting `json:"data"`
	Total int                       `json:"total"`
}
