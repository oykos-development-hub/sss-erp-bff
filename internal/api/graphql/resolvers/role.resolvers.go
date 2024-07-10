package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) RoleDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	roleID := params.Args["id"].(int)

	role, err := r.Repo.GetRole(structs.UserRole(roleID))
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	users, err := r.Repo.GetUserAccounts(&dto.GetUserAccountListInput{RoleID: (*structs.UserRole)(&roleID)})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	responseItem := dto.RoleDetails{
		Role:  *role,
		Users: users.Data,
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    responseItem,
	}, nil
}

func (r *Resolver) RoleOverviewResolver(_ graphql.ResolveParams) (interface{}, error) {
	roleList, err := r.Repo.GetRoleList()
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   roleList,
	}, nil
}

func (r *Resolver) RolesInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Roles
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	if itemID != 0 {
		roleRes, err := r.Repo.UpdateRole(params.Context, itemID, data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		return dto.ResponseSingle{
			Status:  "success",
			Message: "You updated this item!",
			Item:    roleRes,
		}, nil
	}
	roleRes, err := r.Repo.CreateRole(params.Context, data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
		Item:    roleRes,
	}, nil
}
