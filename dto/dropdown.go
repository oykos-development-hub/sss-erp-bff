package dto

type DropdownSimple struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type FileDropdownSimple struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type DropdownBudgetIndent struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	SerialNumber string `json:"serial_number"`
}

type DropdownProcurementArticle struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	VatPercentage string `json:"vat_percentage"`
	Description   string `json:"description"`
}

type DropdownProcurementAvailableArticle struct {
	Id             int     `json:"id"`
	OrderArticleID int     `json:"order_article_id"`
	Title          string  `json:"title"`
	Manufacturer   string  `json:"manufacturer"`
	Description    string  `json:"description"`
	Unit           string  `json:"unit"`
	Available      int     `json:"available"`
	Amount         int     `json:"amount"`
	Price          float32 `json:"price"`
	TotalPrice     float32 `json:"total_price"`
	VatPercentage  string  `json:"vat_percentage"`
}
