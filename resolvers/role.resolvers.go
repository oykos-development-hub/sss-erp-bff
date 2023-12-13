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

var RoleDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	roleID := params.Args["id"].(int)

	role, err := getRole(structs.UserRole(roleID))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	users, err := GetUserAccounts(&dto.GetUserAccountListInput{RoleID: (*structs.UserRole)(&roleID)})
	if err != nil {
		return shared.HandleAPIError(err)
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

var RoleOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	roleList, err := getRoleList()
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   roleList,
	}, nil
}

var RolesInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Roles
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		roleRes, err := updateRole(itemId, data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		return dto.ResponseSingle{
			Status:  "success",
			Message: "You updated this item!",
			Item:    roleRes,
		}, nil
	} else {
		roleRes, err := createRole(data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		return dto.ResponseSingle{
			Status:  "success",
			Message: "You created this item!",
			Item:    roleRes,
		}, nil
	}
}

func getRole(id structs.UserRole) (*structs.Roles, error) {
	res := &dto.GetRoleResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ROLES_ENDPOINT+"/"+strconv.Itoa(int(id)), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getRoleList() ([]structs.Roles, error) {
	res := &dto.GeRoleListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ROLES_ENDPOINT, nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func updateRole(id int, data structs.Roles) (*structs.Roles, error) {
	res := &dto.GetRoleResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.ROLES_ENDPOINT+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createRole(data structs.Roles) (*structs.Roles, error) {
	res := &dto.GetRoleResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ROLES_ENDPOINT, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
