package dto

import "bff/structs"

type PermissionsUpdate struct {
	PermissionList []*structs.RolePermission `json:"permissions_list"`
	RoleID         int                       `json:"role_id"`
}

type GetPermissionListForRoleResponseMS struct {
	Data []structs.Permissions `json:"data"`
}

type GetInsertRolesPermissionListResponseMS struct {
	Data []structs.RolePermission `json:"data"`
}

type PermissionNode struct {
	ID       int               `json:"id"`
	Title    string            `json:"title"`
	Name     string            `json:"name"`
	Route    string            `json:"route"`
	ParentID *int              `json:"parent_id"`
	Create   bool              `json:"create"`
	Read     bool              `json:"read"`
	Update   bool              `json:"update"`
	Delete   bool              `json:"delete"`
	Children []*PermissionNode `json:"children"`
}
