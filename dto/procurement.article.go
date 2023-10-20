package dto

import "bff/structs"

type GetProcurementArticleResponseMS struct {
	Data structs.PublicProcurementArticle `json:"data"`
}

type GetProcurementArticleListResponseMS struct {
	Data []*structs.PublicProcurementArticle `json:"data"`
}

type GetProcurementArticleListInputMS struct {
	ItemID *int `json:"public_procurement_id"`
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
	CreatedAt         string         `json:"created_at"`
	UpdatedAt         string         `json:"updated_at"`
}
