package dto

import (
	"bff/structs"
	"time"

	"github.com/shopspring/decimal"
)

type InternalReallocationResponse struct {
	ID               int                                `json:"id"`
	Title            string                             `json:"title"`
	OrganizationUnit DropdownSimple                     `json:"organization_unit"`
	DateOfRequest    time.Time                          `json:"date_of_request"`
	RequestedBy      DropdownSimple                     `json:"requested_by"`
	File             FileDropdownSimple                 `json:"file"`
	Budget           DropdownSimple                     `json:"budget"`
	Sum              decimal.Decimal                    `json:"sum"`
	Items            []InternalReallocationItemResponse `json:"items"`
	CreatedAt        time.Time                          `json:"created_at"`
	UpdatedAt        time.Time                          `json:"updated_at"`
}

type InternalReallocationItemResponse struct {
	ID                 int             `json:"id"`
	ReallocationID     int             `json:"reallocation_id"`
	SourceAccount      DropdownSimple  `json:"source_account"`
	DestinationAccount DropdownSimple  `json:"destination_account"`
	Amount             decimal.Decimal `json:"amount"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

type GetInternalReallocationResponseMS struct {
	Data  []structs.InternalReallocation `json:"data"`
	Total int                            `json:"total"`
}

type GetInternalReallocationSingleResponseMS struct {
	Data *structs.InternalReallocation `json:"data"`
}

type InternalReallocationFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	SortByTitle        *string `json:"sort_by_title"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	Year               *int    `json:"year"`
	RequestedBy        *int    `json:"requested_by"`
	BudgetID           *int    `json:"budget_id"`
}
