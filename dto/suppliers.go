package dto

import "bff/structs"

type GetSupplierResponseMS struct {
	Data structs.Suppliers `json:"data"`
}

type GetSupplierInputMS struct {
	Entity *string `json:"entity"`
	Search *string `json:"search"`
	Page   *int    `json:"page"`
	Size   *int    `json:"size"`
}

type GetSupplierListResponseMS struct {
	Data  []*structs.Suppliers `json:"data"`
	Total int                  `json:"total"`
}
