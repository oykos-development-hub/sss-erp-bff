package structs

import (
	"time"

	"github.com/shopspring/decimal"
)

type ExternalReallocation struct {
	ID                            int                        `json:"id"`
	Title                         string                     `json:"title"`
	Status                        string                     `json:"status"`
	SourceOrganizationUnitID      int                        `json:"source_organization_unit_id"`
	DestinationOrganizationUnitID int                        `json:"destination_organization_unit_id"`
	DateOfRequest                 time.Time                  `json:"date_of_request"`
	DateOfActionDestOrgUnit       time.Time                  `json:"date_of_action_dest_org_unit"`
	DateOfActionSSS               time.Time                  `json:"date_of_action_sss"`
	RequestedBy                   int                        `json:"requested_by"`
	AcceptedBy                    int                        `json:"accepted_by"`
	FileID                        int                        `json:"file_id"`
	DestinationOrgUnitFileID      int                        `json:"destination_org_unit_file_id"`
	SSSFileID                     int                        `json:"sss_file_id"`
	BudgetID                      int                        `json:"budget_id"`
	Items                         []ExternalReallocationItem `json:"items"`
	CreatedAt                     time.Time                  `json:"created_at"`
	UpdatedAt                     time.Time                  `json:"updated_at"`
}

type ExternalReallocationItem struct {
	ID                   int             `json:"id"`
	ReallocationID       int             `json:"reallocation_id"`
	SourceAccountID      int             `json:"source_account_id"`
	DestinationAccountID int             `json:"destination_account_id"`
	Amount               decimal.Decimal `json:"amount"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
}
