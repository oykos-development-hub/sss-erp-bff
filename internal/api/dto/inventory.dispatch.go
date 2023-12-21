package dto

import "bff/structs"

type GetAllBasicInventoryDispatches struct {
	Data  []*structs.BasicInventoryDispatchItem `json:"data"`
	Total int                                   `json:"total"`
}

type GetBasicInventoryDispatch struct {
	Data *structs.BasicInventoryDispatchItem `json:"data"`
}

type GetAllBasicInventoryDispatchItems struct {
	Data []*structs.BasicInventoryDispatchItemsItem `json:"data"`
}

type InventoryDispatchFilter struct {
	ID                       *int    `json:"id"`
	Type                     *string `json:"type"`
	InventoryType            *string `json:"inventory_type"`
	SourceOrganizationUnitID *int    `json:"source_organization_unit_id"`
	OrganizationUnitID       *int    `json:"organization_unit_id"`
	Accepted                 *bool   `json:"accepted"`
	Page                     *int    `json:"page"`
	Size                     *int    `json:"size"`
}

type InventoryDispatchResponse struct {
	ID                      int                          `json:"id"`
	DispatchID              int                          `json:"dispatch_id"`
	Type                    string                       `json:"type"`
	SourceUserProfile       DropdownSimple               `json:"source_user_profile"`
	TargetUserProfile       DropdownSimple               `json:"target_user_profile"`
	SourceOrganizationUnit  DropdownSimple               `json:"source_organization_unit"`
	TargetOrganizationUnit  DropdownSimple               `json:"target_organization_unit"`
	Office                  DropdownSimple               `json:"office"`
	IsAccepted              bool                         `json:"is_accepted"`
	SerialNumber            string                       `json:"serial_number"`
	DispatchDescription     string                       `json:"dispatch_description"`
	InventoryType           string                       `json:"inventory_type"`
	Inventory               []BasicInventoryResponseItem `json:"inventory"`
	Date                    string                       `json:"date"`
	City                    string                       `json:"city"`
	CreatedAt               string                       `json:"created_at"`
	UpdatedAt               string                       `json:"updated_at"`
	File                    FileDropdownSimple           `json:"file"`
	DeactivationDescription string                       `json:"deactivation_description"`
	DeactivationFile        FileDropdownSimple           `json:"deactivation_file_id"`
	DateOfDeactivation      string                       `json:"date_of_deactivation"`
}
