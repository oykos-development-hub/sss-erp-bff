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
	ID                       *int    `json:"id"`
	Type                     *string `json:"type"`
	ClassTypeID              *int    `json:"class_type_id"`
	OfficeID                 *int    `json:"office_id"`
	Search                   *string `json:"search"`
	SourceType               *string `json:"source_type"`
	DeprecationTypeID        *int    `json:"depreciation_type_id"`
	OrganizationUnitID       *int    `json:"organization_unit_id"`
	SourceOrganizationUnitID *int    `json:"source_organization_unit_id"`
	ContractID               *int    `json:"contract_id"`
	InvoiceArticleID         *int    `json:"invoice_article_id"`
	ArticleID                *int    `json:"article_id"`
	SerialNumber             *string `json:"serial_number"`
	InventoryNumber          *string `json:"inventory_number"`
	Location                 *string `json:"location"`
	IsExternalDonation       *bool   `json:"is_external_donation"`
	Expire                   *bool   `json:"expire"`
	Status                   *string `json:"status"`
	TypeOfImmovableProperty  *string `json:"type_of_immovable_property"`
	CurrentOrganizationUnit  int     `json:"current_organization_unit_id"`
	Page                     *int    `json:"page"`
	Size                     *int    `json:"size"`
}

type ItemReportFilterDTO struct {
	Type               *string `json:"type"`
	SourceType         *string `json:"source_type"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
	OfficeID           *int    `json:"office_id"`
	Date               *string `json:"date"`
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

type ItemReportResponse struct {
	ID               int     `json:"id"`
	Title            string  `json:"title"`
	SourceType       string  `json:"source_type"`
	Type             string  `json:"type"`
	InventoryNumber  string  `json:"inventory_number"`
	OfficeID         int     `json:"office_id"`
	Office           string  `json:"office"`
	ProcurementPrice float32 `json:"procurement_price"`
	LostValue        float32 `json:"lost_value"`
	Price            float32 `json:"price"`
	Date             string  `json:"date"`
	DateOfPurchase   string  `json:"date_of_purchase"`
}

type GetAllItemsReportMS struct {
	Data []ItemReportResponse `json:"data"`
}

type DispatchInventoryItemFilter struct {
	DispatchID *int `json:"dispatch_id"`
}

type BasicInventoryResponseListItem struct {
	ID                           int                                                            `json:"id"`
	Active                       bool                                                           `json:"active"`
	Type                         string                                                         `json:"type"`
	Title                        string                                                         `json:"title"`
	Location                     string                                                         `json:"location"`
	InventoryNumber              string                                                         `json:"inventory_number"`
	PurchaseGrossPrice           float32                                                        `json:"purchase_gross_price"`
	GrossPrice                   float32                                                        `json:"gross_price"`
	ResidualPrice                *float32                                                       `json:"residual_price"`
	DateOfPurchase               string                                                         `json:"date_of_purchase"`
	Description                  string                                                         `json:"description"`
	DateOfAssessments            string                                                         `json:"date_of_assessments"`
	DateOfEndOfAssessment        string                                                         `json:"date_of_end_of_assessment"`
	Inactive                     *string                                                        `json:"inactive"`
	EstimatedDuration            int                                                            `json:"estimated_duration"`
	Status                       string                                                         `json:"status"`
	SourceType                   string                                                         `json:"source_type"`
	Source                       string                                                         `json:"source"`
	City                         *string                                                        `json:"city"`
	Address                      *string                                                        `json:"address"`
	LifetimeOfAssessmentInMonths int                                                            `json:"lifetime_of_assessment_in_months"`
	AmortizationValue            float32                                                        `json:"amortization_value"`
	HasAssessments               bool                                                           `json:"has_assessments"`
	IsExternalDonation           bool                                                           `json:"is_external_donation"`
	RealEstate                   *structs.BasicInventoryRealEstatesItemResponseForInventoryItem `json:"real_estate"`
	DepreciationType             DropdownSimple                                                 `json:"depreciation_type"`
	OrganizationUnit             DropdownOUSimple                                               `json:"organization_unit"`
	TargetOrganizationUnit       DropdownOUSimple                                               `json:"target_organization_unit"`
	ClassType                    DropdownSimple                                                 `json:"class_type"`
	Office                       DropdownSimple                                                 `json:"office"`
	Invoice                      DropdownSimple                                                 `json:"invoice"`
}

type BasicInventoryResponseItem struct {
	ID                           int                                                            `json:"id"`
	ArticleID                    int                                                            `json:"article_id"`
	InvoiceArticleID             int                                                            `json:"invoice_article_id"`
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
	LostValue          float32 `json:"lost_value"`
	Price              float32 `json:"price"`
}

type ReportValueClassInventory struct {
	Values             []ReportValueClassInventoryItem `json:"items"`
	PurchaseGrossPrice float32                         `json:"purchase_gross_price"`
	LostValue          float32                         `json:"lost_value"`
	Price              float32                         `json:"price"`
}
