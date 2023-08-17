package dto

import "bff/structs"

type GetOrderProcurementArticlesResponseMS struct {
	Data  []structs.OrderProcurementArticleItem `json:"data"`
	Total int                                   `json:"total"`
}

type GetOrderProcurementArticleResponseMS struct {
	Data structs.OrderProcurementArticleItem `json:"data"`
}

type GetOrderProcurementArticleInput struct {
	ArticleID *int `json:"article_id"`
	OrderID   *int `json:"order_id"`
}

type OrderProcurementArticleOverviewResponse struct {
	Id                   int                                    `json:"id"`
	DateOrder            string                                 `json:"date_order" validate:"required"`
	TotalPrice           int                                    `json:"total_price"`
	PublicProcurementID  int                                    `json:"public_procurement_id"`
	SupplierID           int                                    `json:"supplier_id"`
	Status               string                                 `json:"status"`
	DateSystem           string                                 `json:"date_system"`
	InvoiceDate          string                                 `json:"invoice_date"`
	InvoiceNumber        int                                    `json:"invoice_number"`
	OrganizationUnitID   int                                    `json:"organization_unit_id"`
	OfficeID             int                                    `json:"office_id"`
	RecipientUserID      int                                    `json:"recipient_user_id"`
	DescriptionRecipient *string                                `json:"description_recipient"`
	CreatedAt            string                                 `json:"created_at"`
	UpdatedAt            string                                 `json:"updated_at"`
	PublicProcurement    *DropdownSimple                        `json:"public_procurement"`
	Supplier             *DropdownSimple                        `json:"supplier"`
	Articles             *[]DropdownProcurementAvailableArticle `json:"articles"`
}
