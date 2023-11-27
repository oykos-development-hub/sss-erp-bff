package structs

type BasicInventoryAssessmentsTypesItem struct {
	Id                   int     `json:"id"`
	Type                 string  `json:"type"`
	InventoryId          int     `json:"inventory_id"`
	EstimatedDuration    int     `json:"estimated_duration"`
	Active               bool    `json:"active"`
	DepreciationTypeId   int     `json:"depreciation_type_id"`
	UserProfileId        int     `json:"user_profile_id"`
	GrossPriceNew        float32 `json:"gross_price_new"`
	GrossPriceDifference float32 `json:"gross_price_difference"`
	DateOfAssessment     *string `json:"date_of_assessment"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
	FileId               *int    `json:"file_id"`
}
