package dto

import (
	"bff/structs"
	"time"
)

type ModelsOfAccountingResponse struct {
	ID        int                             `json:"id"`
	Title     string                          `json:"title"`
	Type      string                          `json:"type"`
	Items     []ModelOfAccountingItemResponse `json:"items"`
	CreatedAt time.Time                       `json:"created_at"`
	UpdatedAt time.Time                       `json:"updated_at"`
}

type ModelOfAccountingItemResponse struct {
	ID            int                     `json:"id"`
	Title         string                  `json:"title"`
	ModelID       int                     `json:"model_id"`
	DebitAccount  AccountItemResponseItem `json:"debit_account"`
	CreditAccount AccountItemResponseItem `json:"credit_account"`
	CreatedAt     time.Time               `json:"created_at"`
	UpdatedAt     time.Time               `json:"updated_at"`
}

type ModelsOfAccountingFilter struct {
	Search *string `json:"search"`
}

type GetModelsOfAccountingResponseMS struct {
	Data structs.ModelsOfAccounting `json:"data"`
}

type GetModelsOfAccountingListResponseMS struct {
	Data  []structs.ModelsOfAccounting `json:"data"`
	Total int                          `json:"total"`
}
