package structs

type OrderListItem struct {
	Id                  int    `json:"id"`
	DataOrder           string `json:"data_order"`
	TotalPrice          int    `json:"total_price"`
	PublicProcurementId int    `json:"public_procurement_id"`
	SupplierId          int    `json:"supplier_id"`
	Status              string `json:"status"`
	DateSystem          string `json:"date_system"`
	InvoiceDate         string `json:"invoice_date"`
	InvoiceNumber       string `json:"invoice_number"`
	OrganizationUnitId  int    `json:"organization_unit_id"`
	UserProfileId       int    `json:"user_profile_id"`
}

type OrderProcurementArticleItem struct {
	Id        int `json:"id"`
	OrderId   int `json:"order_id"`
	ArticleId int `json:"article_id"`
	Amount    int `json:"amount"`
}

type OrderArticleItem struct {
	Id                  int    `json:"id"`
	BudgetIndentId      int    `json:"budget_indent_id"`
	PublicProcurementId int    `json:"public_procurement_id"`
	Title               string `json:"title"`
	Description         string `json:"description"`
	NetPrice            string `json:"net_price"`
	VatPercentage       string `json:"vat_percentage"`
	Amount              int    `json:"amount"`
	Available           int    `json:"available"`
	TotalPrice          int    `json:"total_price"`
	Unit                string `json:"unit"`
}

type OrderListInsertItem struct {
	Id                  int                      `json:"id"`
	DataOrder           string                   `json:"data_order"`
	PublicProcurementId int                      `json:"public_procurement_id"`
	SupplierId          int                      `json:"supplier_id"`
	Articles            []OrderArticleInsertItem `json:"articles"`
}

type OrderArticleInsertItem struct {
	Id     int `json:"id"`
	Amount int `json:"amount"`
}

type OrderReceiveItem struct {
	OrderId       int    `json:"id"`
	DateSystem    string `json:"date_system"`
	InvoiceDate   string `json:"invoice_date"`
	TotalPrice    int    `json:"total_price"`
	InvoiceNumber string `json:"invoice_number"`
}

type OrderAssetMovementItem struct {
	OrderId            int `json:"id"`
	OrganizationUnitId int `json:"organization_unit_id"`
	UserProfileId      int `json:"user_profile_id"`
}
