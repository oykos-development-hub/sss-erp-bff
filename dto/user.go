package dto

import "bff/structs"

type GetUsersResponseMS struct {
	Data  []structs.UserAccounts `json:"data"`
	Total int                    `json:"total"`
}

type GetUserAccountResponseMS struct {
	Data structs.UserAccounts `json:"data"`
}

type GetUsersInput struct {
	Page *int `json:"page"`
	Size *int `json:"size"`
}
