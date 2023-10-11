package dto

import "bff/structs"

type GetProcurementItemResponseMS struct {
	Data structs.PublicProcurementItem `json:"data"`
}

type GetProcurementItemListResponseMS struct {
	Data []*structs.PublicProcurementItem `json:"data"`
}

type GetProcurementItemListInputMS struct {
	Page   *int `json:"page"`
	Size   *int `json:"size"`
	PlanID *int `json:"plan_id"`
}

type ProcurementItemResponseItem struct {
	Id                int                               `json:"id"`
	BudgetIndent      DropdownBudgetIndent              `json:"budget_indent"`
	Plan              DropdownSimple                    `json:"plan_id"`
	IsOpenProcurement bool                              `json:"is_open_procurement"`
	Title             string                            `json:"title"`
	ArticleType       string                            `json:"article_type"`
	Status            structs.ProcurementStatus         `json:"status"`
	SerialNumber      *string                           `json:"serial_number"`
	DateOfPublishing  *string                           `json:"date_of_publishing"`
	DateOfAwarding    *string                           `json:"date_of_awarding"`
	FileId            *int                              `json:"file_id"`
	Articles          []*ProcurementArticleResponseItem `json:"articles"`
	CreatedAt         string                            `json:"created_at"`
	UpdatedAt         string                            `json:"updated_at"`
}
