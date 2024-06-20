package dto

import (
	"bff/structs"
	"time"

	"github.com/shopspring/decimal"
)

type ExternalReallocationResponse struct {
	ID                          int                                `json:"id"`
	Title                       string                             `json:"title"`
	Status                      string                             `json:"status"`
	SourceOrganizationUnit      DropdownSimple                     `json:"source_organization_unit"`
	DestinationOrganizationUnit DropdownSimple                     `json:"destination_organization_unit"`
	DateOfRequest               time.Time                          `json:"date_of_request"`
	DateOfActionDestOrgUnit     time.Time                          `json:"date_of_action_dest_org_unit"`
	DateOfActionSSS             time.Time                          `json:"date_of_action_sss"`
	RequestedBy                 DropdownSimple                     `json:"requested_by"`
	AcceptedBy                  DropdownSimple                     `json:"accepted_by"`
	File                        FileDropdownSimple                 `json:"file"`
	DestinationOrgUnitFile      FileDropdownSimple                 `json:"destination_org_unit_file"`
	SSSFile                     FileDropdownSimple                 `json:"sss_file"`
	Budget                      DropdownSimple                     `json:"budget"`
	Items                       []ExternalReallocationItemResponse `json:"items"`
	CreatedAt                   time.Time                          `json:"created_at"`
	UpdatedAt                   time.Time                          `json:"updated_at"`
}

type ExternalReallocationItemResponse struct {
	ID                 int             `json:"id"`
	ReallocationID     int             `json:"reallocation_id"`
	SourceAccount      DropdownSimple  `json:"source_account"`
	DestinationAccount DropdownSimple  `json:"destination_account"`
	Amount             decimal.Decimal `json:"amount"`
	CreatedAt          time.Time       `json:"created_at"`
	UpdatedAt          time.Time       `json:"updated_at"`
}

type GetExternalReallocationResponseMS struct {
	Data  []structs.ExternalReallocation `json:"data"`
	Total int                            `json:"total"`
}

type GetExternalReallocationSingleResponseMS struct {
	Data *structs.ExternalReallocation `json:"data"`
}

type ExternalReallocationFilter struct {
	Page                          *int    `json:"page"`
	Size                          *int    `json:"size"`
	SortByTitle                   *string `json:"sort_by_title"`
	SourceOrganizationUnitID      *int    `json:"source_organization_unit_id"`
	DestinationOrganizationUnitID *int    `json:"destination_organization_unit_id"`
	OrganizationUnitID            *int    `json:"organization_unit_id"`
	Status                        *string `json:"status"`
	RequestedBy                   *int    `json:"requested_by"`
	BudgetID                      *int    `json:"budget_id"`
}
