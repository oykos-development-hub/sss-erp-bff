package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type InternalReallocation struct {
	ID                 int                        `json:"id"`
	Title              string                     `json:"title"`
	OrganizationUnitID int                        `json:"organization_unit_id"`
	DateOfRequest      time.Time                  `json:"date_of_request"`
	RequestedBy        int                        `json:"requested_by"`
	FileID             int                        `json:"file_id"`
	BudgetID           int                        `json:"budget_id"`
	Items              []InternalReallocationItem `json:"items"`
	CreatedAt          time.Time                  `json:"created_at"`
	UpdatedAt          time.Time                  `json:"updated_at"`
}

type InternalReallocationItem struct {
	ID                   int             `json:"id"`
	ReallocationID       int             `json:"reallocation_id"`
	SourceAccountID      int             `json:"source_account_id"`
	DestinationAccountID int             `json:"destination_account_id"`
	Amount               decimal.Decimal `json:"amount"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
}
