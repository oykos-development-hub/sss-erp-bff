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
	SourceOrganizationUnitID *int    `json:"source_organiation_unit_id"`
	Accepted                 *bool   `json:"accepted"`
	Page                     *int    `json:"page"`
	Size                     *int    `json:"size"`
}

type InventoryDispatchResponse struct {
	ID                     int                          `json:"id"`
	Type                   string                       `json:"type"`
	SourceUserProfile      DropdownSimple               `json:"source_user_profile"`
	TargetUserProfile      DropdownSimple               `json:"target_user_profile"`
	SourceOrganizationUnit DropdownSimple               `json:"source_organization_unit"`
	TargetOrganizationUnit DropdownSimple               `json:"target_organization_unit"`
	Office                 DropdownSimple               `json:"office"`
	IsAccepted             bool                         `json:"is_accepted"`
	SerialNumber           string                       `json:"serial_number"`
	DispatchDescription    string                       `json:"dispatch_description"`
	Inventory              []BasicInventoryResponseItem `json:"inventory"`
	CreatedAt              string                       `json:"created_at"`
	UpdatedAt              string                       `json:"updated_at"`
	FileId                 int                          `json:"file_id"`
}
