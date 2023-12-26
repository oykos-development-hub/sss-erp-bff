package dto

import (
	"bff/structs"
)

type GetBasicInventoryInsertItem struct {
	Data *structs.BasicInventoryInsertItem `json:"data"`
}

type GetAllBasicInventoryInsertItem struct {
	Data []interface{} `json:"data"`
}

type GetAllBasicInventoryItem struct {
	Data  []*structs.BasicInventoryInsertItem `json:"data"`
	Total int                                 `json:"total"`
}

type InventoryItemFilter struct {
	ID                 *int    `json:"id"`
	Type               *string `json:"type"`
	ClassTypeID        *int    `json:"class_type_id"`
	OfficeID           *int    `json:"office_id"`
	Search             *string `json:"search"`
	SourceType         *string `json:"source_type"`
	DeprecationTypeID  *int    `json:"depreciation_type_id"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	ContractID         *int    `json:"contract_id "`
	ArticleID          *int    `json:"article_id"`
	SerialNumber       *string `json:"serial_number"`
	InventoryNumber    *string `json:"inventory_number"`
	Location           *string `json:"location"`
	IsExternalDonation *bool   `json:"is_external_donation"`
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
}

type GetAllItemsInOrgUnits struct {
	ItemID     int   `json:"item_id"`
	ReversID   int   `json:"revers_id"`
	ReturnID   int   `json:"return_id"`
	MovementID []int `json:"movement_id"`
}

type GetAllItemsInOrgUnitsMS struct {
	Data []GetAllItemsInOrgUnits `json:"data"`
}

type DispatchInventoryItemFilter struct {
	DispatchID *int `json:"dispatch_id"`
}

type BasicInventoryResponseListItem struct {
	ID                     int                                                            `json:"id"`
	Active                 bool                                                           `json:"active"`
	Type                   string                                                         `json:"type"`
	Title                  string                                                         `json:"title"`
	Location               string                                                         `json:"location"`
	InventoryNumber        string                                                         `json:"inventory_number"`
	PurchaseGrossPrice     float32                                                        `json:"purchase_gross_price"`
	GrossPrice             float32                                                        `json:"gross_price"`
	ResidualPrice          *float32                                                       `json:"residual_price"`
	DateOfPurchase         string                                                         `json:"date_of_purchase"`
	DateOfAssessments      string                                                         `json:"date_of_assessments"`
	EstimatedDuration      int                                                            `json:"estimated_duration"`
	Status                 string                                                         `json:"status"`
	SourceType             string                                                         `json:"source_type"`
	HasAssessments         bool                                                           `json:"has_assessments"`
	IsExternalDonation     bool                                                           `json:"is_external_donation"`
	RealEstate             *structs.BasicInventoryRealEstatesItemResponseForInventoryItem `json:"real_estate"`
	DepreciationType       DropdownSimple                                                 `json:"depreciation_type"`
	OrganizationUnit       DropdownSimple                                                 `json:"organization_unit"`
	TargetOrganizationUnit DropdownSimple                                                 `json:"target_organization_unit"`
	ClassType              DropdownSimple                                                 `json:"class_type"`
	Office                 DropdownSimple                                                 `json:"office"`
	Invoice                DropdownSimple                                                 `json:"invoice"`
}

type BasicInventoryResponseItem struct {
	ID                           int                                                            `json:"id"`
	ArticleID                    int                                                            `json:"article_id"`
	Type                         string                                                         `json:"type"`
	ClassType                    DropdownSimple                                                 `json:"class_type"`
	DepreciationType             DropdownSimple                                                 `json:"depreciation_type"`
	Supplier                     DropdownSimple                                                 `json:"supplier"`
	Donor                        DropdownSimple                                                 `json:"donor"`
	Invoice                      DropdownSimple                                                 `json:"invoice"`
	RealEstate                   *structs.BasicInventoryRealEstatesItemResponseForInventoryItem `json:"real_estate"`
	Assessments                  []*BasicInventoryResponseAssessment                            `json:"assessments"`
	Movements                    []*InventoryDispatchResponse                                   `json:"movements"`
	SerialNumber                 string                                                         `json:"serial_number"`
	Status                       string                                                         `json:"status"`
	InventoryNumber              string                                                         `json:"inventory_number"`
	Title                        string                                                         `json:"title"`
	Abbreviation                 string                                                         `json:"abbreviation"`
	InternalOwnership            bool                                                           `json:"internal_ownership"`
	SourceType                   string                                                         `json:"source_type"`
	Office                       DropdownSimple                                                 `json:"office"`
	Location                     string                                                         `json:"location"`
	TargetUserProfile            DropdownSimple                                                 `json:"target_user_profile"`
	OrganizationUnit             DropdownSimple                                                 `json:"organization_unit"`
	TargetOrganizationUnit       DropdownSimple                                                 `json:"target_organization_unit"`
	City                         string                                                         `json:"city"`
	Address                      string                                                         `json:"address"`
	Unit                         string                                                         `json:"unit"`
	Amount                       int                                                            `json:"amount"`
	NetPrice                     float32                                                        `json:"net_price"`
	PurchaseGrossPrice           float32                                                        `json:"purchase_gross_price"`
	GrossPrice                   float32                                                        `json:"gross_price"`
	ResidualPrice                *float32                                                       `json:"residual_price"`
	Description                  string                                                         `json:"description"`
	DateOfPurchase               string                                                         `json:"date_of_purchase"`
	Source                       string                                                         `json:"source"`
	DonorTitle                   string                                                         `json:"donor_title"`
	InvoiceNumber                string                                                         `json:"invoice_number"`
	PriceOfAssessment            int                                                            `json:"price_of_assessment"`
	DateOfAssessment             *string                                                        `json:"date_of_assessment"`
	LifetimeOfAssessmentInMonths int                                                            `json:"lifetime_of_assessment_in_months"`
	DepreciationRate             string                                                         `json:"depreciation_rate"`
	AmortizationValue            float32                                                        `json:"amortization_value"`
	Active                       bool                                                           `json:"active"`
	Inactive                     *string                                                        `json:"inactive"`
	DeactivationDescription      string                                                         `json:"deactivation_description"`
	CreatedAt                    string                                                         `json:"created_at"`
	UpdatedAt                    string                                                         `json:"updated_at"`
	InvoiceFileID                int                                                            `json:"invoice_file_id"`
	FileID                       int                                                            `json:"file_id"`
	DonationDescription          string                                                         `json:"donation_description"`
	DonationFiles                []FileDropdownSimple                                           `json:"donation_files"`
	Owner                        string                                                         `json:"owner"`
	IsExternalDonation           bool                                                           `json:"is_external_donation"`
}

type ReportValueClassInventoryItem struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	Class              string  `json:"class"`
	PurchaseGrossPrice float32 `json:"purchase_gross_price"`
	GrossPrice         float32 `json:"gross_price"`
	PriceOfAssessment  float32 `json:"price_of_assessment"`
}

type ReportValueClassInventory struct {
	Values             []ReportValueClassInventoryItem `json:"items"`
	PurchaseGrossPrice float32                         `json:"purchase_gross_price"`
	GrossPrice         float32                         `json:"gross_price"`
	PriceOfAssessment  float32                         `json:"price_of_assessment"`
}
