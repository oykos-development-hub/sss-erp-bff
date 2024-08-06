package dto

import "bff/structs"

type GetUserAccountListResponseMS struct {
	Data  []structs.UserAccounts `json:"data"`
	Total int                    `json:"total"`
}

type GetUserAccountResponseMS struct {
	Data structs.UserAccounts `json:"data"`
}

type GetUserAccountListInput struct {
	Page      *int    `json:"page"`
	Size      *int    `json:"size"`
	IsActive  *bool   `json:"is_active"`
	RoleID    *int    `json:"role_id"`
	Email     *string `json:"email"`
	ProfileID *int    `json:"profile_id"`
}

type DeactivateUserAccount struct {
	Active bool `json:"active"`
}
