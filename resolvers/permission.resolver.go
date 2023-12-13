package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"
)

var PermissionsUpdateResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data dto.PermissionsUpdate
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	_, err := syncPermissions(data.RoleID, data.PermissionList)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	permissions, err := getPermissionList(data.RoleID)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	permissionsTree := buildTree(permissions)

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    permissionsTree,
	}, nil
}

var PermissionsForRoleResolver = func(params graphql.ResolveParams) (interface{}, error) {
	roleID := params.Args["role_id"].(int)

	permissions, err := getPermissionList(roleID)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	permissionsTree := buildTree(permissions)

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    permissionsTree,
	}, nil
}

func buildTree(permissions []structs.Permissions) *dto.PermissionNode {
	nodes := make(map[int]*dto.PermissionNode)
	var tree *dto.PermissionNode

	for _, p := range permissions {
		nodes[p.ID] = &dto.PermissionNode{
			ID:       p.ID,
			Title:    p.Title,
			Route:    p.Route,
			ParentID: p.ParentId,
			Create:   p.Create,
			Read:     p.Read,
			Update:   p.Update,
			Delete:   p.Delete,
			Name:     p.Title,
			Children: nil,
		}
	}

	for _, node := range nodes {
		if node.ParentID != nil {
			parent, ok := nodes[*node.ParentID]
			if ok {
				parent.Children = append(parent.Children, node)
			}
		} else {
			tree = node
		}
	}

	return tree
}

func syncPermissions(roleID int, input []*structs.RolePermission) ([]structs.RolePermission, error) {
	res := &dto.GetInsertRolesPermissionListResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ROLES_ENDPOINT+"/"+strconv.Itoa(roleID)+"/permissions/sync", input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getPermissionList(roleID int) ([]structs.Permissions, error) {
	res := &dto.GetPermissionListForRoleResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ROLES_ENDPOINT+"/"+strconv.Itoa(roleID)+"/permissions", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
