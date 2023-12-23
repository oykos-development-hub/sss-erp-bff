package structs

type BasicInventoryAssessmentsTypesItem struct {
	ID                   int      `json:"id"`
	Type                 string   `json:"type"`
	InventoryID          int      `json:"inventory_id"`
	EstimatedDuration    int      `json:"estimated_duration"`
	Active               bool     `json:"active"`
	DepreciationTypeID   int      `json:"depreciation_type_id"`
	UserProfileID        int      `json:"user_profile_id"`
	GrossPriceNew        float32  `json:"gross_price_new"`
	GrossPriceDifference float32  `json:"gross_price_difference"`
	ResidualPrice        *float32 `json:"residual_price"`
	DateOfAssessment     *string  `json:"date_of_assessment"`
	CreatedAt            string   `json:"created_at"`
	UpdatedAt            string   `json:"updated_at"`
	FileID               *int     `json:"file_id"`
}
