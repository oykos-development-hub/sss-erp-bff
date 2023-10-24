package structs

type BasicInventoryDispatchItem struct {
	Id                       int    `json:"id"`
	SourceUserProfileId      int    `json:"source_user_profile_id"`
	TargetUserProfileId      int    `json:"target_user_profile_id"`
	SourceOrganizationUnitId int    `json:"source_organization_unit_id"`
	TargetOrganizationUnitId int    `json:"target_organization_unit_id"`
	OfficeId                 int    `json:"office_id"`
	Type                     string `json:"type"`
	IsAccepted               bool   `json:"is_accepted"`
	SerialNumber             string `json:"serial_number"`
	DispatchDescription      string `json:"dispatch_description"`
	InventoryType            string `json:"inventory_type"`
	InventoryId              []int  `json:"inventory_id"`
	Date                     string `json:"date"`
	CreatedAt                string `json:"created_at"`
	UpdatedAt                string `json:"updated_at"`
	FileId                   int    `json:"file_id"`
}

type BasicInventoryDispatchItemsItem struct {
	Id          int `json:"id"`
	InventoryId int `json:"inventory_id"`
	DispatchId  int `json:"dispatch_id"`
}
