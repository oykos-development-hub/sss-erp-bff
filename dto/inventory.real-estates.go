package dto

import "bff/structs"

type GetInventoryRealEstateListInputMS struct {
	Page *int `json:"page"`
	Size *int `json:"size"`
}

type GetInventoryRealEstateResponseMS struct {
	Data structs.BasicInventoryRealEstatesItem `json:"data"`
}

type GetInventoryRealEstateListResponseMS struct {
	Data  []*structs.BasicInventoryRealEstatesItem `json:"data"`
	Total int                                      `json:"total"`
}

type GetMyInventoryRealEstateResponseMS struct {
	Data structs.BasicInventoryRealEstatesItem `json:"data"`
}
