package structs

type BasicInventoryRealEstatesItem struct {
	ID                       int     `json:"id"`
	Title                    string  `json:"title"`
	ItemID                   int     `json:"item_id"`
	TypeID                   string  `json:"type_id"`
	SquareArea               float64 `json:"square_area"`
	LandSerialNumber         string  `json:"land_serial_number"`
	EstateSerialNumber       string  `json:"estate_serial_number"`
	OwnershipType            string  `json:"ownership_type"`
	OwnershipScope           string  `json:"ownership_scope"`
	OwnershipInvestmentScope string  `json:"ownership_investment_scope"`
	LimitationsDescription   string  `json:"limitations_description"`
	PropertyDocument         string  `json:"property_document"`
	LimitationID             bool    `json:"limitation_id"`
	Document                 string  `json:"document"`
	FileID                   int     `json:"file_id"`
}

type BasicInventoryRealEstatesItemResponseForInventoryItem struct {
	ID                       int     `json:"id"`
	TypeID                   string  `json:"type_id"`
	SquareArea               float64 `json:"square_area"`
	LandSerialNumber         string  `json:"land_serial_number"`
	EstateSerialNumber       string  `json:"estate_serial_number"`
	OwnershipType            string  `json:"ownership_type"`
	OwnershipScope           string  `json:"ownership_scope"`
	OwnershipInvestmentScope string  `json:"ownership_investment_scope"`
	LimitationsDescription   string  `json:"limitations_description"`
	LimitationsID            bool    `json:"limitation_id"`
	Document                 string  `json:"document"`
	PropertyDocument         string  `json:"property_document"`
	FileID                   int     `json:"file_id"`
}
