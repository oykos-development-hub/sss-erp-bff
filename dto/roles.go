package dto

import "bff/structs"

type RoleDetails struct {
	Role  structs.Roles          `json:"role"`
	Users []structs.UserAccounts `json:"users"`
}

type GetRoleResponseMS struct {
	Data structs.Roles `json:"data"`
}

type GeRoleListResponseMS struct {
	Data  []structs.Roles `json:"data"`
	Total int             `json:"total"`
}
