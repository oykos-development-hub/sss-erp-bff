package dto

type DropdownSimple struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type FileDropdownSimple struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type DropdownBudgetIndent struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	SerialNumber string `json:"serial_number"`
}

type DropdownProcurementArticle struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	VatPercentage string `json:"vat_percentage"`
	Description   string `json:"description"`
}

type DropdownProcurementAvailableArticle struct {
	ID             int     `json:"id"`
	OrderArticleID int     `json:"order_article_id"`
	Title          string  `json:"title"`
	Manufacturer   string  `json:"manufacturer"`
	Description    string  `json:"description"`
	Unit           string  `json:"unit"`
	Available      int     `json:"available"`
	Amount         int     `json:"amount"`
	Price          float32 `json:"price"`
	TotalPrice     float32 `json:"total_price"`
	NetPrice       float32 `json:"net_price"`
	VatPercentage  string  `json:"vat_percentage"`
}

type DropdownOUSimple struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	City    string `json:"city"`
	Address string `json:"address"`
}
