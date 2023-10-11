package dto

import (
	"bff/structs"
)

type ProcurementItemWithOrganizationUnitArticleResponseItem struct {
	Id                int                                               `json:"id"`
	BudgetIndent      DropdownSimple                                    `json:"budget_indent"`
	Plan              DropdownSimple                                    `json:"plan"`
	IsOpenProcurement bool                                              `json:"is_open_procurement"`
	Title             string                                            `json:"title"`
	ArticleType       string                                            `json:"article_type"`
	Status            structs.ProcurementStatus                         `json:"status"`
	SerialNumber      *string                                           `json:"serial_number"`
	DateOfPublishing  *string                                           `json:"date_of_publishing"`
	DateOfAwarding    *string                                           `json:"date_of_awarding"`
	FileID            *int                                              `json:"file_id"`
	Articles          []*ProcurementOrganizationUnitArticleResponseItem `json:"articles"`
	CreatedAt         string                                            `json:"created_at"`
	UpdatedAt         string                                            `json:"updated_at"`
}

type ProcurementOrganizationUnitArticleResponseItem struct {
	Id                  int                            `json:"id"`
	Article             ProcurementArticleResponseItem `json:"public_procurement_article"`
	OrganizationUnit    DropdownSimple                 `json:"organization_unit"`
	Amount              int                            `json:"amount"`
	Status              structs.ArticleStatus          `json:"status"`
	IsRejected          bool                           `json:"is_rejected"`
	RejectedDescription string                         `json:"rejected_description"`
	CreatedAt           string                         `json:"created_at"`
	UpdatedAt           string                         `json:"updated_at"`
}

type GetProcurementOrganizationUnitArticleListInputDTO struct {
	OrganizationUnitID *int `json:"organization_unit_id" validate:"omitempty"`
	ArticleID          *int `json:"article_id" validate:"omitempty"`
}

type GetOrganizationUnitArticleResponseMS struct {
	Data structs.PublicProcurementOrganizationUnitArticle `json:"data"`
}

type GetOrganizationUnitArticleListResponseMS struct {
	Data []*structs.PublicProcurementOrganizationUnitArticle `json:"data"`
}
