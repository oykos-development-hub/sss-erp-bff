package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/shared"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) RoleDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	roleID := params.Args["id"].(int)

	role, err := r.Repo.GetRole(structs.UserRole(roleID))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	users, err := r.Repo.GetUserAccounts(&dto.GetUserAccountListInput{RoleID: (*structs.UserRole)(&roleID)})
	if err != nil {
		return errors.HandleAPIError(err)
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

func (r *Resolver) RoleOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	roleList, err := r.Repo.GetRoleList()
	if err != nil {
		return errors.HandleAPIError(err)
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

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		roleRes, err := r.Repo.UpdateRole(itemId, data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.ResponseSingle{
			Status:  "success",
			Message: "You updated this item!",
			Item:    roleRes,
		}, nil
	} else {
		roleRes, err := r.Repo.CreateRole(data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.ResponseSingle{
			Status:  "success",
			Message: "You created this item!",
			Item:    roleRes,
		}, nil
	}
}
