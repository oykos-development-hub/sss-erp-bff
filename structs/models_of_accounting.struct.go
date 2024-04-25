package structs

import "time"

type ModelsOfAccounting struct {
	ID        int                 `json:"id"`
	Title     string              `json:"title"`
	Type      string              `json:"type"`
	Items     []ModelOfAccounting `json:"items"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}

type ModelOfAccounting struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	ModelID         int       `json:"model_id"`
	DebitAccountID  int       `json:"debit_account_id"`
	CreditAccountID int       `json:"credit_account_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
