package dto

import "bff/structs"

type GetUserAccountListResponseMS struct {
	Data  []structs.UserAccounts `json:"data"`
	Total int                    `json:"total"`
}

type GetUserAccountResponseMS struct {
	Data structs.UserAccounts `json:"data"`
}

type GetUserAccountRoleResponseMS struct {
	Data structs.UserAccountRoles `json:"data"`
}

type GetUserAccountListInput struct {
	Page     *int    `json:"page"`
	Size     *int    `json:"size"`
	IsActive *bool   `json:"is_active"`
	Email    *string `json:"email"`
}
