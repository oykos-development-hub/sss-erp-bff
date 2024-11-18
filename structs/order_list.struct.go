package structs

import "time"

type OrderListItem struct {
	ID                    int        `json:"id"`
	DateOrder             string     `json:"date_order"`
	TotalPrice            float64    `json:"total_price"`
	PublicProcurementID   *int       `json:"public_procurement_id"`
	GroupOfArticlesID     *int       `json:"group_of_articles_id"`
	SupplierID            *int       `json:"supplier_id"`
	AccountID             *int       `json:"account_id"`
	Status                string     `json:"status"`
	PassedToFinance       bool       `json:"passed_to_finance"`
	UsedInFinance         bool       `json:"used_in_finance"`
	IsProFormaInvoice     bool       `json:"is_pro_forma_invoice"`
	ProFormaInvoiceDate   *string    `json:"pro_forma_invoice_date"`
	ProFormaInvoiceNumber string     `json:"pro_forma_invoice_number"`
	DateSystem            *string    `json:"date_system"`
	InvoiceDate           *string    `json:"invoice_date"`
	InvoiceNumber         *string    `json:"invoice_number"`
	OrganizationUnitID    int        `json:"organization_unit_id"`
	OfficeID              *int       `json:"office_id"`
	RecipientUserID       *int       `json:"recipient_user_id"`
	Description           *string    `json:"description"`
	IsUsed                bool       `json:"is_used"`
	OrderFile             *int       `json:"order_file"`
	ReceiveFile           []int      `json:"receive_file"`
	MovementFile          *int       `json:"movement_file"`
	DeliveryDate          *time.Time `json:"delivery_date"`
	DeliveryNumber        string     `json:"delivery_number"`
	DeliveryFileID        *int       `json:"delivery_file_id"`
}

type OrderProcurementArticleItem struct {
	ID            int     `json:"id"`
	OrderID       int     `json:"order_id"`
	ArticleID     int     `json:"article_id"`
	Year          string  `json:"year"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Amount        int     `json:"amount"`
	NetPrice      float64 `json:"net_price"`
	VatPercentage int     `json:"vat_percentage"`
}

type OrderArticleItem struct {
	ID                  int            `json:"id"`
	BudgetIndentID      int            `json:"budget_indent_id"`
	PublicProcurementID int            `json:"public_procurement_id"`
	Title               string         `json:"title"`
	Description         string         `json:"description"`
	NetPrice            float64        `json:"net_price"`
	VatPercentage       string         `json:"vat_percentage"`
	Manufacturer        string         `json:"manufacturer"`
	Amount              int            `json:"amount"`
	ConsumedAmount      int            `json:"consumed_amount"`
	Available           int            `json:"available"`
	TotalPrice          float64        `json:"total_price"`
	Price               float64        `json:"price"`
	Unit                string         `json:"unit"`
	VisibilityType      VisibilityType `json:"visibility_type"`
}

type OrderListInsertItem struct {
	ID                    int                      `json:"id"`
	DateOrder             string                   `json:"date_order"`
	PublicProcurementID   int                      `json:"public_procurement_id"`
	GroupOfArticlesID     int                      `json:"group_of_articles_id"`
	SupplierID            int                      `json:"supplier_id"`
	AccountID             int                      `json:"account_id"`
	Articles              []OrderArticleInsertItem `json:"articles"`
	IsUsed                bool                     `json:"is_used"`
	OrderFile             int                      `json:"order_file"`
	PassedToFinance       bool                     `json:"passed_to_finance"`
	UsedInFinance         bool                     `json:"used_in_finance"`
	IsProFormaInvoice     bool                     `json:"is_pro_forma_invoice"`
	ProFormaInvoiceDate   *string                  `json:"pro_forma_invoice_date"`
	ProFormaInvoiceNumber string                   `json:"pro_forma_invoice_number"`
}

type OrderArticleInsertItem struct {
	ID            int     `json:"id"`
	Amount        int     `json:"amount"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	NetPrice      float64 `json:"net_price"`
	VatPercentage int     `json:"vat_percentage"`
}

type OrderReceiveItem struct {
	OrderID        int                      `json:"order_id"`
	DateSystem     string                   `json:"date_system"`
	InvoiceDate    *string                  `json:"invoice_date"`
	InvoiceNumber  *string                  `json:"invoice_number"`
	Description    *string                  `json:"description"`
	ReceiveFile    []int                    `json:"receive_file"`
	DeliveryDate   *time.Time               `json:"delivery_date"`
	DeliveryNumber string                   `json:"delivery_number"`
	DeliveryFileID *int                     `json:"delivery_file_id"`
	Articles       []OrderArticleInsertItem `json:"articles"`
}

type OrderAssetMovementItem struct {
	ID                 int            `json:"id"`
	DateOrder          string         `json:"date_order"`
	OfficeID           int            `json:"office_id"`
	RecipientUserID    int            `json:"recipient_user_id"`
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
	NetPrice           float64 `json:"net_price"`
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
