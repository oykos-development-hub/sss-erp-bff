package structs

type OrderListItem struct {
	Id                  int     `json:"id"`
	DateOrder           string  `json:"date_order"`
	TotalPrice          float32 `json:"total_price"`
	PublicProcurementId int     `json:"public_procurement_id"`
	SupplierId          *int    `json:"supplier_id"`
	Status              string  `json:"status"`
	DateSystem          *string `json:"date_system"`
	InvoiceDate         *string `json:"invoice_date"`
	InvoiceNumber       *string `json:"invoice_number"`
	OrganizationUnitId  int     `json:"organization_unit_id"`
	OfficeId            *int    `json:"office_id"`
	RecipientUserId     *int    `json:"recipient_user_id"`
	Description         *string `json:"description"`
	IsUsed              bool    `json:"is_used"`
	File                []int   `json:"file"`
}

type OrderProcurementArticleItem struct {
	Id        int `json:"id"`
	OrderId   int `json:"order_id"`
	ArticleId int `json:"article_id"`
	Amount    int `json:"amount"`
}

type OrderArticleItem struct {
	Id                  int     `json:"id"`
	BudgetIndentId      int     `json:"budget_indent_id"`
	PublicProcurementId int     `json:"public_procurement_id"`
	Title               string  `json:"title"`
	Description         string  `json:"description"`
	NetPrice            float32 `json:"net_price"`
	VatPercentage       string  `json:"vat_percentage"`
	Manufacturer        string  `json:"manufacturer"`
	Amount              int     `json:"amount"`
	Available           int     `json:"available"`
	TotalPrice          float32 `json:"total_price"`
	Price               float32 `json:"price"`
	Unit                string  `json:"unit"`
}

type OrderListInsertItem struct {
	Id                  int                      `json:"id"`
	DateOrder           string                   `json:"date_order"`
	PublicProcurementId int                      `json:"public_procurement_id"`
	SupplierId          int                      `json:"supplier_id"`
	Articles            []OrderArticleInsertItem `json:"articles"`
	IsUsed              bool                     `json:"is_used"`
	File                []int                    `json:"file"`
}

type OrderArticleInsertItem struct {
	Id     int `json:"id"`
	Amount int `json:"amount"`
}

type OrderReceiveItem struct {
	OrderId       int     `json:"order_id"`
	DateSystem    string  `json:"date_system"`
	InvoiceDate   string  `json:"invoice_date"`
	InvoiceNumber string  `json:"invoice_number"`
	Description   *string `json:"description"`
}

type OrderAssetMovementItem struct {
	OrderId         int `json:"order_id"`
	OfficeId        int `json:"office_id"`
	RecipientUserId int `json:"recipient_user_id"`
}
