package dto

import (
	"bff/structs"
)

type GetProcurementContractArticleResponseMS struct {
	Data structs.PublicProcurementContractArticle `json:"data"`
}

type GetProcurementContractArticlesInput struct {
	ContractID *int `json:"contract_id"`
}

type GetProcurementContractArticlesListResponseMS struct {
	Data  []*structs.PublicProcurementContractArticle `json:"data"`
	Total int                                         `json:"total"`
}

type ProcurementContractArticlesResponseItem struct {
	Id         int                        `json:"id"`
	Article    DropdownProcurementArticle `json:"public_procurement_article"`
	Contract   DropdownSimple             `json:"contract"`
	Amount     int                        `json:"amount" validate:"required"`
	NetValue   float32                    `json:"net_value"`
	GrossValue float32                    `json:"gross_value"`
	CreatedAt  string                     `json:"created_at"`
	UpdatedAt  string                     `json:"updated_at"`
}
