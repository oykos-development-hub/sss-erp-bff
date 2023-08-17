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

type InventoryItemFilter struct {
	ID                 int    `json:"id"`
	Type               string `json:"type"`
	ClassTypeID        int    `json:"class_type_id"`
	OfficeID           int    `json:"office_id"`
	Search             string `json:"search"`
	SourceType         string `json:"source_type"`
	DeprecationTypeID  int    `json:"deprecation_type_id"`
	OrganizationUnitID int    `json:"organiation_unit_id"`
	Page               int    `json:"page"`
	Size               int    `json:"size"`
}

type BasicInventoryInsertItemResponseMS struct {
	Id                           int                                                            `json:"id"`
	ArticleId                    int                                                            `json:"article_id"`
	Type                         string                                                         `json:"type"`
	ClassType                    structs.SettingsDropdown                                       `json:"class_type"`
	DepreciationType             structs.SettingsDropdown                                       `json:"depreciation_type"`
	Supplier                     structs.SettingsDropdown                                       `json:"supplier"`
	RealEstate                   *structs.BasicInventoryRealEstatesItemResponseForInventoryItem `json:"real_estate"`
	SerialNumber                 string                                                         `json:"serial_number"`
	InventoryNumber              string                                                         `json:"inventory_number"`
	Title                        string                                                         `json:"title"`
	Abbreviation                 string                                                         `json:"abbreviation"`
	InternalOwnership            bool                                                           `json:"internal_ownership"`
	Office                       structs.SettingsDropdown                                       `json:"office"`
	Location                     string                                                         `json:"location"`
	TargetUserProfile            structs.SettingsDropdown                                       `json:"target_user_profile"`
	Unit                         string                                                         `json:"unit"`
	Amount                       int                                                            `json:"amount"`
	NetPrice                     int                                                            `json:"net_price"`
	GrossPrice                   int                                                            `json:"gross_price"`
	Description                  string                                                         `json:"description"`
	DateOfPurchase               string                                                         `json:"date_of_purchase"`
	Source                       string                                                         `json:"source"`
	DonorTitle                   string                                                         `json:"donor_title"`
	InvoiceNumber                string                                                         `json:"invoice_number"`
	PriceOfAssessment            int                                                            `json:"price_of_assessment"`
	DateOfAssessment             string                                                         `json:"date_of_assessment"`
	LifetimeOfAssessmentInMonths int                                                            `json:"lifetime_of_assessment_in_months"`
	Active                       bool                                                           `json:"active"`
	DeactivationDescription      string                                                         `json:"deactivation_description"`
	CreatedAt                    string                                                         `json:"created_at"`
	UpdatedAt                    string                                                         `json:"updated_at"`
	InvoiceFileId                int                                                            `json:"invoice_file_id"`
	FileId                       int                                                            `json:"file_id"`
}
