package structs

type BasicInventoryDetailsItem struct {
	ID                           int     `json:"id"`
	ArticleID                    int     `json:"article_id"`
	Type                         string  `json:"type"`
	ClassTypeID                  int     `json:"class_type_id"`
	DepreciationTypeID           int     `json:"depreciation_type_id"`
	SupplierID                   int     `json:"supplier_id"`
	RealEstateID                 int     `json:"real_estate_id"`
	SerialNumber                 string  `json:"serial_number"`
	InventoryNumber              string  `json:"inventory_number"`
	Title                        string  `json:"title"`
	Abbreviation                 string  `json:"abbreviation"`
	InternalOwnership            bool    `json:"internal_ownership"`
	OfficeID                     int     `json:"office_id"`
	InvoiceID                    int     `json:"invoice_id"`
	Location                     string  `json:"location"`
	OrganizationUnitID           int     `json:"organization_unit_id"`
	TargetOrganizationUnitID     int     `json:"target_organization_unit_id"`
	TargetUserProfileID          int     `json:"target_user_profile_id"`
	Unit                         string  `json:"unit"`
	Amount                       int     `json:"amount"`
	NetPrice                     float32 `json:"net_price"`
	GrossPrice                   float32 `json:"gross_price"`
	PurchaseGrossPrice           float32 `json:"purchase_gross_price"`
	Description                  string  `json:"description"`
	DateOfPurchase               string  `json:"date_of_purchase"`
	Source                       string  `json:"source"`
	Status                       string  `json:"status"`
	DonorTitle                   string  `json:"donor_title"`
	InvoiceNumber                int     `json:"invoice_number"`
	PriceOfAssessment            int     `json:"price_of_assessment"`
	DateOfAssessment             string  `json:"date_of_assessment"`
	LifetimeOfAssessmentInMonths int     `json:"lifetime_of_assessment_in_months"`
	Active                       bool    `json:"active"`
	DeactivationDescription      string  `json:"deactivation_description"`
	CreatedAt                    string  `json:"created_at"`
	UpdatedAt                    string  `json:"updated_at"`
	InvoiceFileID                string  `json:"invoice_file_id"`
	FileID                       string  `json:"file_id"`
}

type BasicInventoryInsertItem struct {
	ID                           int                            `json:"id"`
	ArticleID                    int                            `json:"article_id"`
	Type                         string                         `json:"type"`
	ClassTypeID                  int                            `json:"class_type_id"`
	DepreciationTypeID           int                            `json:"depreciation_type_id"`
	SupplierID                   int                            `json:"supplier_id"`
	DonorID                      int                            `json:"donor_id"`
	RealEstate                   *BasicInventoryRealEstatesItem `json:"real_estate"`
	RealEstateID                 int                            `json:"real_estate_id"`
	SerialNumber                 string                         `json:"serial_number"`
	InventoryNumber              string                         `json:"inventory_number"`
	Title                        string                         `json:"title"`
	Owner                        string                         `json:"owner"`
	Abbreviation                 string                         `json:"abbreviation"`
	InternalOwnership            bool                           `json:"internal_ownership"`
	InvoiceID                    int                            `json:"invoice_id"`
	OfficeID                     int                            `json:"office_id"`
	Location                     string                         `json:"location"`
	OrganizationUnitID           int                            `json:"organization_unit_id"`
	TargetOrganizationUnitID     int                            `json:"target_organization_unit_id"`
	TargetUserProfileID          int                            `json:"target_user_profile_id"`
	Unit                         string                         `json:"unit"`
	Amount                       int                            `json:"amount"`
	NetPrice                     float32                        `json:"net_price"`
	GrossPrice                   float32                        `json:"gross_price"`
	Description                  string                         `json:"description"`
	DateOfPurchase               string                         `json:"date_of_purchase"`
	Source                       string                         `json:"source"`
	SourceType                   string                         `json:"source_type"`
	DonorTitle                   string                         `json:"donor_title"`
	InvoiceNumber                string                         `json:"invoice_number"`
	PriceOfAssessment            int                            `json:"price_of_assessment"`
	DateOfAssessment             *string                        `json:"date_of_assessment"`
	LifetimeOfAssessmentInMonths int                            `json:"lifetime_of_assessment_in_months"`
	Active                       bool                           `json:"active"`
	DeactivationDescription      string                         `json:"deactivation_description"`
	Inactive                     *string                        `json:"inactive"`
	CreatedAt                    string                         `json:"created_at"`
	UpdatedAt                    string                         `json:"updated_at"`
	InvoiceFileID                int                            `json:"invoice_file_id"`
	DeactivationFileID           int                            `json:"deactivation_file_id"`
	FileID                       int                            `json:"file_id"`
	ContractID                   int                            `json:"contract_id"`
	ContractArticleID            int                            `json:"contract_article_id"`
	DonationDescription          string                         `json:"donation_description"`
	DonationFiles                []int                          `json:"donation_files"`
	IsExternalDonation           bool                           `json:"is_external_donation"`
}

type BasicInventoryItem struct {
	ID                       int     `json:"id"`
	Type                     string  `json:"type"`
	ClassTypeID              int     `json:"class_type_id"`
	DepreciationTypeID       int     `json:"depreciation_type_id"`
	RealEstateID             int     `json:"real_estate_id"`
	InventoryNumber          string  `json:"inventory_number"`
	Title                    string  `json:"title"`
	InvoiceID                int     `json:"invoice_id"`
	OfficeID                 int     `json:"office_id"`
	TargetUserProfileID      int     `json:"target_user_profile_id"`
	OrganizationUnitID       int     `json:"organization_unit_id"`
	TargetOrganizationUnitID int     `json:"target_organization_unit_id"`
	GrossPrice               float32 `json:"gross_price"`
	DateOfPurchase           string  `json:"date_of_purchase"`
	Source                   string  `json:"source"`
	Active                   bool    `json:"active"`
	Location                 string  `json:"location"`
}

type BasicInventoryItemDispatch struct {
	ID                        int    `json:"id"`
	ParentID                  int    `json:"parent_id"`
	BasicInventoryItemID      int    `json:"basic_inventory_item_id"`
	DispatchedByUserProfileID int    `json:"dispatched_by_user_profile_id"`
	DispatchedToUserProfileID int    `json:"dispatched_to_user_profile_id"`
	OrganizationUnitID        int    `json:"organization_unit_id"`
	OfficeID                  int    `json:"office_id"`
	SerialNumber              string `json:"serial_number"`
}

type BasicInventoryInsertValidator struct {
	Value  string `json:"value"`
	Entity string `json:"entity"`
}
