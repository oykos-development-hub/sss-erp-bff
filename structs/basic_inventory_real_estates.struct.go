package structs

type BasicInventoryRealEstatesItem struct {
	Id                       int     `json:"id"`
	Title                    string  `json:"title"`
	ItemId                   int     `json:"item_id"`
	TypeId                   string  `json:"type_id"`
	SquareArea               float32 `json:"square_area"`
	LandSerialNumber         string  `json:"land_serial_number"`
	EstateSerialNumber       string  `json:"estate_serial_number"`
	OwnershipType            string  `json:"ownership_type"`
	OwnershipScope           string  `json:"ownership_scope"`
	OwnershipInvestmentScope string  `json:"ownership_investment_scope"`
	LimitationsDescription   string  `json:"limitations_description"`
	PropertyDocument         string  `json:"property_document"`
	LimitationId             bool    `json:"limitation_id"`
	Document                 string  `json:"document"`
	FileId                   int     `json:"file_id"`
}

type BasicInventoryRealEstatesItemResponseForInventoryItem struct {
	Id                       int     `json:"id"`
	TypeId                   string  `json:"type_id"`
	SquareArea               float32 `json:"square_area"`
	LandSerialNumber         string  `json:"land_serial_number"`
	EstateSerialNumber       string  `json:"estate_serial_number"`
	OwnershipType            string  `json:"ownership_type"`
	OwnershipScope           string  `json:"ownership_scope"`
	OwnershipInvestmentScope string  `json:"ownership_investment_scope"`
	LimitationsDescription   string  `json:"limitations_description"`
	LimitationsId            bool    `json:"limitation_id"`
	Document                 string  `json:"document"`
	PropertyDocument         string  `json:"property_document"`
	FileId                   int     `json:"file_id"`
}
