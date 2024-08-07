package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) PermissionsUpdateResolver(params graphql.ResolveParams) (interface{}, error) {
	var data dto.PermissionsUpdate
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	_, err := r.Repo.SyncPermissions(data.RoleID, data.PermissionList)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	permissions, err := r.Repo.GetPermissionList(data.RoleID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	permissionsTree := buildTree(permissions)

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    permissionsTree,
	}, nil
}

func (r *Resolver) PermissionsForRoleResolver(params graphql.ResolveParams) (interface{}, error) {
	roleID, roleIDOK := params.Args["role_id"].(int)

	var permissions []structs.Permissions
	var err error

	if !roleIDOK {
		permissions, err = r.Repo.GetPermissionList(1)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		for i := 0; i < len(permissions); i++ {
			permissions[i].Create = false
			permissions[i].Update = false
			permissions[i].Read = false
			permissions[i].Delete = false
		}
	} else {
		permissions, err = r.Repo.GetPermissionList(roleID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
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
			ParentID: p.ParentID,
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
