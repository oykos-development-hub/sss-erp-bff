package dto

import "bff/structs"

type GetProcurementItemResponseMS struct {
	Data structs.PublicProcurementItem `json:"data"`
}

type GetProcurementItemListResponseMS struct {
	Data []*structs.PublicProcurementItem `json:"data"`
}

type ProcurementItemResponseItem struct {
	Id                int     `json:"id"`
	BudgetIndentId    int     `json:"budget_indent_id"`
	PlanId            int     `json:"plan_id"`
	IsOpenProcurement bool    `json:"is_open_procurement"`
	Title             string  `json:"title"`
	ArticleType       string  `json:"article_type"`
	Status            *string `json:"status"`
	SerialNumber      *string `json:"serial_number"`
	DateOfPublishing  *string `json:"date_of_publishing"`
	DateOfAwarding    *string `json:"date_of_awarding"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
	FileId            *int    `json:"file_id"`
}
