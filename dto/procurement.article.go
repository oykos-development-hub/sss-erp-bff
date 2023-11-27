package dto

import "bff/structs"

type GetProcurementArticleResponseMS struct {
	Data structs.PublicProcurementArticle `json:"data"`
}

type GetProcurementArticleListResponseMS struct {
	Data []*structs.PublicProcurementArticle `json:"data"`
}

type GetProcurementArticleListInputMS struct {
	ItemID      *int    `json:"public_procurement_id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Year        *string `json:"year"`
}

type ProcurementArticleResponseItem struct {
	Id                int            `json:"id"`
	PublicProcurement DropdownSimple `json:"public_procurement"`
	Title             string         `json:"active"`
	Description       string         `json:"year"`
	NetPrice          float32        `json:"net_price"`
	VATPercentage     *string        `json:"vat_percentage"`
	Manufacturer      string         `json:"manufacturer"`
	Amount            int            `json:"amount"`
	TotalAmount       int            `json:"total_amount"`
	GrossPrice        float32        `json:"gross_price"`
	VisibilityType    int            `json:"visibility_type"`
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         string         `json:"updated_at"`
}
