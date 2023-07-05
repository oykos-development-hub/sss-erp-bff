package structs

type BasicInventoryRealEstatesItem struct {
	Id                       int    `json:"id"`
	Title                    string `json:"title"`
	SquareArea               int    `json:"square_area"`
	LandSerialNumber         string `json:"land_serial_number"`
	EstateSerialNumber       string `json:"estate_serial_number"`
	OwnershipType            string `json:"ownership_type"`
	OwnershipScope           string `json:"ownership_scope"`
	OwnershipInvestmentScope string `json:"ownership_investment_scope"`
	LimitationsDescription   string `json:"limitations_description"`
	FileId                   string `json:"file_id"`
}
