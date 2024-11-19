package dto

import (
	"bff/structs"
)

type GetProcurementContractArticleResponseMS struct {
	Data structs.PublicProcurementContractArticle `json:"data"`
}

type GetProcurementContractArticlesInput struct {
	ContractID *int `json:"contract_id"`
	ArticleID  *int `json:"article_id"`
	ID         *int `json:"id"`
}

type GetProcurementContractArticlesListResponseMS struct {
	Data  []*structs.PublicProcurementContractArticle `json:"data"`
	Total int                                         `json:"total"`
}

type ProcurementContractArticlesResponseItem struct {
	ID           int                                                `json:"id"`
	Article      DropdownProcurementArticle                         `json:"public_procurement_article"`
	Contract     DropdownSimple                                     `json:"contract"`
	Amount       int                                                `json:"amount"`
	UsedArticles int                                                `json:"used_articles"`
	OverageList  []*structs.PublicProcurementContractArticleOverage `json:"overages"`
	OverageTotal int                                                `json:"overage_total"`
	NetValue     float64                                            `json:"net_value"`
	GrossValue   float64                                            `json:"gross_value"`
	CreatedAt    string                                             `json:"created_at"`
	UpdatedAt    string                                             `json:"updated_at"`
}

type GetProcurementContractArticleOverageResponseMS struct {
	Data structs.PublicProcurementContractArticleOverage `json:"data"`
}

type GetProcurementContractArticleOverageInput struct {
	ContractArticleID  *int `json:"article_id"`
	OrganizationUnitID *int `json:"organization_unit_id"`
}

type GetProcurementContractArticleOverageListResponseMS struct {
	Data []*structs.PublicProcurementContractArticleOverage `json:"data"`
}
