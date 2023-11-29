package dto

import "bff/structs"

type GetOrderListsResponseMS struct {
	Data  []structs.OrderListItem `json:"data"`
	Total int                     `json:"total"`
}

type GetOrderListResponseMS struct {
	Data structs.OrderListItem `json:"data"`
}

type GetOrderListInput struct {
	Page                *int    `json:"page"`
	Size                *int    `json:"size"`
	SupplierID          *int    `json:"supplier_id"`
	Search              *string `json:"search"`
	Status              *string `json:"status"`
	Year                *string `json:"year"`
	PublicProcurementID *int    `json:"public_procurement_id"`
	OrganizationUnitId  *int    `json:"organization_unit_id"`
}

type OrderListOverviewResponse struct {
	Id                  int                                    `json:"id"`
	DateOrder           string                                 `json:"date_order" validate:"required"`
	TotalBruto          float32                                `json:"total_bruto"`
	TotalNeto           float32                                `json:"total_neto"`
	PublicProcurementID int                                    `json:"public_procurement_id"`
	SupplierID          int                                    `json:"supplier_id"`
	Status              string                                 `json:"status"`
	DateSystem          *string                                `json:"date_system"`
	InvoiceDate         *string                                `json:"invoice_date"`
	InvoiceNumber       string                                 `json:"invoice_number"`
	OrganizationUnitID  int                                    `json:"organization_unit_id"`
	OfficeID            int                                    `json:"office_id"`
	RecipientUserID     *int                                   `json:"recipient_user_id"`
	Description         *string                                `json:"description"`
	IsUsed              bool                                   `json:"is_used"`
	CreatedAt           string                                 `json:"created_at"`
	UpdatedAt           string                                 `json:"updated_at"`
	GroupOfArticles     *DropdownSimple                        `json:"group_of_articles"`
	PublicProcurement   *DropdownSimple                        `json:"public_procurement"`
	Supplier            *DropdownSimple                        `json:"supplier"`
	RecipientUser       *DropdownSimple                        `json:"recipient_user"`
	Office              *DropdownSimple                        `json:"office"`
	Articles            *[]DropdownProcurementAvailableArticle `json:"articles"`
	OrderFile           FileDropdownSimple                     `json:"order_file"`
	ReceiveFile         FileDropdownSimple                     `json:"receive_file"`
	MovementFile        FileDropdownSimple                     `json:"movement_file"`
}

type StockFilter struct {
	Page               *int    `json:"page"`
	Size               *int    `json:"size"`
	Title              *string `json:"title"`
	Year               *string `json:"year"`
	Description        *string `json:"description"`
	OrganizationUnitID *int    `json:"organization_unit_id"`
}

type GetStockResponseMS struct {
	Data  []structs.StockArticle `json:"data"`
	Total int                    `json:"total"`
}

type GetSingleStockResponseMS struct {
	Data structs.StockArticle `json:"data"`
}

type MovementFilter struct {
	Page            *int `json:"page"`
	Size            *int `json:"size"`
	OfficeID        *int `json:"office_id"`
	RecipientUserID *int `json:"recipient_user_id"`
}

type GetMovementResponseMS struct {
	Data  []structs.Movement `json:"data"`
	Total int                `json:"total"`
}

type GetSingleMovementResponseMS struct {
	Data structs.Movement `json:"data"`
}

type GetSingleMovementArticleResponseMS struct {
	Data MovementArticle `json:"data"`
}

type MovementResponse struct {
	ID            int            `json:"id"`
	Description   string         `json:"description"`
	Office        DropdownSimple `json:"office"`
	RecipientUser DropdownSimple `json:"recipient_user"`
	DateOrder     string         `json:"date_order"`
}

type MovementDetailsResponse struct {
	ID            int                `json:"id"`
	Description   string             `json:"description"`
	Office        DropdownSimple     `json:"office"`
	RecipientUser DropdownSimple     `json:"recipient_user"`
	DateOrder     string             `json:"date_order"`
	Articles      []ArticlesDropdown `json:"articles"`
	File          FileDropdownSimple `json:"file"`
}

type ArticlesDropdown struct {
	ID          int    `json:"id"`
	ArticleID   int    `json:"article_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Year        string `json:"year"`
}

type MovementArticle struct {
	ID                 int    `json:"id"`
	Year               string `json:"year"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	Amount             int    `json:"amount"`
	MovementID         int    `json:"movement_id"`
	OrganizationUnitID int    `json:"organization_unit_id"`
}

type GetMovementArticleResponseMS struct {
	Data  []MovementArticle `json:"data"`
	Total int               `json:"total"`
}

type MovementArticlesFilter struct {
	MovementID *int `json:"movement_id"`
}
