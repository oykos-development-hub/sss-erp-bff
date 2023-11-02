package structs

type BasicInventoryDetailsItem struct {
	Id                           int    `json:"id"`
	ArticleId                    int    `json:"article_id"`
	Type                         string `json:"type"`
	ClassTypeId                  int    `json:"class_type_id"`
	DepreciationTypeId           int    `json:"depreciation_type_id"`
	SupplierId                   int    `json:"supplier_id"`
	RealEstateId                 int    `json:"real_estate_id"`
	SerialNumber                 string `json:"serial_number"`
	InventoryNumber              string `json:"inventory_number"`
	Title                        string `json:"title"`
	Abbreviation                 string `json:"abbreviation"`
	InternalOwnership            bool   `json:"internal_ownership"`
	OfficeId                     int    `json:"office_id"`
	Location                     string `json:"location"`
	OrganizationUnitId           int    `json:"organization_unit_id"`
	TargetOrganizationUnitId     int    `json:"target_organization_unit_id"`
	TargetUserProfileId          int    `json:"target_user_profile_id"`
	Unit                         string `json:"unit"`
	Amount                       int    `json:"amount"`
	NetPrice                     int    `json:"net_price"`
	GrossPrice                   int    `json:"gross_price"`
	PurchaseGrossPrice           int    `json:"purchase_gross_price"`
	Description                  string `json:"description"`
	DateOfPurchase               string `json:"date_of_purchase"`
	Source                       string `json:"source"`
	Status                       string `json:"status"`
	DonorTitle                   string `json:"donor_title"`
	InvoiceNumber                int    `json:"invoice_number"`
	PriceOfAssessment            int    `json:"price_of_assessment"`
	DateOfAssessment             string `json:"date_of_assessment"`
	LifetimeOfAssessmentInMonths int    `json:"lifetime_of_assessment_in_months"`
	Active                       bool   `json:"active"`
	DeactivationDescription      string `json:"deactivation_description"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
	InvoiceFileId                string `json:"invoice_file_id"`
	FileId                       string `json:"file_id"`
}

type BasicInventoryInsertItem struct {
	Id                           int                            `json:"id"`
	ArticleId                    int                            `json:"article_id"`
	Type                         string                         `json:"type"`
	ClassTypeId                  int                            `json:"class_type_id"`
	DepreciationTypeId           int                            `json:"depreciation_type_id"`
	SupplierId                   int                            `json:"supplier_id"`
	RealEstate                   *BasicInventoryRealEstatesItem `json:"real_estate"`
	RealEstateId                 int                            `json:"real_estate_id"`
	SerialNumber                 string                         `json:"serial_number"`
	InventoryNumber              string                         `json:"inventory_number"`
	Title                        string                         `json:"title"`
	Abbreviation                 string                         `json:"abbreviation"`
	InternalOwnership            bool                           `json:"internal_ownership"`
	OfficeId                     int                            `json:"office_id"`
	Location                     string                         `json:"location"`
	OrganizationUnitId           int                            `json:"organization_unit_id"`
	TargetOrganizationUnitId     int                            `json:"target_organization_unit_id"`
	TargetUserProfileId          int                            `json:"target_user_profile_id"`
	Unit                         string                         `json:"unit"`
	Amount                       int                            `json:"amount"`
	NetPrice                     int                            `json:"net_price"`
	GrossPrice                   int                            `json:"gross_price"`
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
	InvoiceFileId                int                            `json:"invoice_file_id"`
	FileId                       int                            `json:"file_id"`
}

type BasicInventoryItem struct {
	Id                       int    `json:"id"`
	Type                     string `json:"type"`
	ClassTypeId              int    `json:"class_type_id"`
	DepreciationTypeId       int    `json:"depreciation_type_id"`
	RealEstateId             int    `json:"real_estate_id"`
	InventoryNumber          string `json:"inventory_number"`
	Title                    string `json:"title"`
	OfficeId                 int    `json:"office_id"`
	TargetUserProfileId      int    `json:"target_user_profile_id"`
	OrganizationUnitId       int    `json:"organization_unit_id"`
	TargetOrganizationUnitId int    `json:"target_organization_unit_id"`
	GrossPrice               int    `json:"gross_price"`
	DateOfPurchase           string `json:"date_of_purchase"`
	Source                   string `json:"source"`
	Active                   bool   `json:"active"`
}

type BasicInventoryItemDispatch struct {
	Id                        int    `json:"id"`
	ParentId                  int    `json:"parent_id"`
	BasicInventoryItemId      int    `json:"basic_inventory_item_id"`
	DispatchedByUserProfileId int    `json:"dispatched_by_user_profile_id"`
	DispatchedToUserProfileId int    `json:"dispatched_to_user_profile_id"`
	OrganizationUnitId        int    `json:"organization_unit_id"`
	OfficeId                  int    `json:"office_id"`
	SerialNumber              string `json:"serial_number"`
}
