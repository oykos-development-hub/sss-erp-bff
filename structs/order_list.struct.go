package structs

type OrderListItem struct {
	Id                  int     `json:"id"`
	DateOrder           string  `json:"date_order"`
	TotalPrice          float32 `json:"total_price"`
	PublicProcurementId *int    `json:"public_procurement_id"`
	GroupOfArticlesID   *int    `json:"group_of_articles_id"`
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
	OrderFile           *int    `json:"order_file"`
	ReceiveFile         *int    `json:"receive_file"`
	MovementFile        *int    `json:"movement_file"`
}

type OrderProcurementArticleItem struct {
	Id            int     `json:"id"`
	OrderId       int     `json:"order_id"`
	ArticleId     int     `json:"article_id"`
	Year          string  `json:"year"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Amount        int     `json:"amount"`
	NetPrice      float32 `json:"net_price"`
	VatPercentage int     `json:"vat_percentage"`
}

type OrderArticleItem struct {
	Id                  int            `json:"id"`
	BudgetIndentId      int            `json:"budget_indent_id"`
	PublicProcurementId int            `json:"public_procurement_id"`
	Title               string         `json:"title"`
	Description         string         `json:"description"`
	NetPrice            float32        `json:"net_price"`
	VatPercentage       string         `json:"vat_percentage"`
	Manufacturer        string         `json:"manufacturer"`
	Amount              int            `json:"amount"`
	Available           int            `json:"available"`
	TotalPrice          float32        `json:"total_price"`
	Price               float32        `json:"price"`
	Unit                string         `json:"unit"`
	VisibilityType      VisibilityType `json:"visibility_type"`
}

type OrderListInsertItem struct {
	Id                  int                      `json:"id"`
	DateOrder           string                   `json:"date_order"`
	PublicProcurementId int                      `json:"public_procurement_id"`
	GroupOfArticlesID   int                      `json:"group_of_articles_id"`
	SupplierId          int                      `json:"supplier_id"`
	Articles            []OrderArticleInsertItem `json:"articles"`
	IsUsed              bool                     `json:"is_used"`
	OrderFile           int                      `json:"order_file"`
}

type OrderArticleInsertItem struct {
	Id            int     `json:"id"`
	Amount        int     `json:"amount"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	NetPrice      float32 `json:"net_price"`
	VatPercentage int     `json:"vat_percentage"`
}

type OrderReceiveItem struct {
	OrderId       int                      `json:"order_id"`
	DateSystem    string                   `json:"date_system"`
	InvoiceDate   *string                  `json:"invoice_date"`
	InvoiceNumber *string                  `json:"invoice_number"`
	Description   *string                  `json:"description"`
	ReceiveFile   *int                     `json:"receive_file"`
	Articles      []OrderArticleInsertItem `json:"articles"`
}

type OrderAssetMovementItem struct {
	ID                 int            `json:"id"`
	DateOrder          string         `json:"date_order"`
	OfficeId           int            `json:"office_id"`
	RecipientUserId    int            `json:"recipient_user_id"`
	FileID             int            `json:"file_id"`
	Description        string         `json:"description"`
	OrganizationUnitID int            `json:"organization_unit_id"`
	Articles           []StockArticle `json:"articles"`
}

type StockArticle struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	ArticleID          int     `json:"article_id"`
	Description        string  `json:"description"`
	Exception          bool    `json:"exception"`
	OrganizationUnitID int     `json:"organization_unit_id"`
	NetPrice           float32 `json:"net_price"`
	VatPercentage      int     `json:"vat_percentage"`
	Year               string  `json:"year"`
	Amount             int     `json:"amount"`
	Quantity           int     `json:"quantity"`
}

type Movement struct {
	ID              int    `json:"id"`
	Description     string `json:"description"`
	OfficeID        int    `json:"office_id"`
	RecipientUserID int    `json:"recipient_user_id"`
	DateOrder       string `json:"date_order"`
	FileID          int    `json:"file_id"`
}
