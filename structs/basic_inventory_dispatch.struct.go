package structs

type BasicInventoryDispatchItem struct {
	ID                       int    `json:"id"`
	DispatchID               int    `json:"dispatch_id"`
	SourceUserProfileID      int    `json:"source_user_profile_id"`
	TargetUserProfileID      int    `json:"target_user_profile_id"`
	SourceOrganizationUnitID int    `json:"source_organization_unit_id"`
	TargetOrganizationUnitID int    `json:"target_organization_unit_id"`
	OfficeID                 int    `json:"office_id"`
	Type                     string `json:"type"`
	IsAccepted               bool   `json:"is_accepted"`
	SerialNumber             string `json:"serial_number"`
	DispatchDescription      string `json:"dispatch_description"`
	InventoryType            string `json:"inventory_type"`
	InventoryID              []int  `json:"inventory_id"`
	Date                     string `json:"date"`
	CreatedAt                string `json:"created_at"`
	UpdatedAt                string `json:"updated_at"`
	FileID                   int    `json:"file_id"`
	IsExternalDonation       bool   `json:"is_external_donation"`
	DonationFiles            []int  `json:"donation_files"`
}

type BasicInventoryDispatchItemsItem struct {
	ID          int `json:"id"`
	InventoryID int `json:"inventory_id"`
	DispatchID  int `json:"dispatch_id"`
}
