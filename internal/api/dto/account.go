package dto

import "bff/structs"

type GetAccountsFilter struct {
	ID     *int    `json:"id"`
	Search *string `json:"search"`
	Page   *int    `json:"page"`
	Size   *int    `json:"size"`
}

type AccountItemResponseItem struct {
	ID           int                        `json:"id"`
	SerialNumber string                     `json:"serial_number"`
	Title        string                     `json:"title"`
	ParentId     *int                       `json:"parent_id"`
	Children     []*AccountItemResponseItem `json:"children"`
}

type GetAccountItemResponseMS struct {
	Data structs.AccountItem `json:"data"`
}

type GetAccountItemListResponseMS struct {
	Data  []*structs.AccountItem `json:"data"`
	Total int                    `json:"total"`
}
