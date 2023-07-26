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

var UserAccountsOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []structs.UserAccounts
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	isActive, isActiveOk := params.Args["is_active"].(bool)
	email, emailOk := params.Args["email"].(string)

	if id != nil && shared.IsInteger(id) && id != 0 {
		user, err := GetUserAccountById(id.(int))
		if err != nil {
			return dto.Response{
				Status:  "error",
				Message: err.Error(),
			}, nil
		}
		items = []structs.UserAccounts{*user}
		total = 1
	} else {
		input := dto.GetUserAccountListInput{}
		if shared.IsInteger(page) && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if shared.IsInteger(size) && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if isActiveOk {
			input.IsActive = &isActive
		}
		if emailOk && email != "" {
			input.Email = &email
		}

		res, err := GetUserAccounts(&input)
		if err != nil {
			return dto.Response{
				Status:  "error",
				Message: err.Error(),
			}, nil
		}
		items = res.Data
		total = res.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
	}, nil
}

var UserAccountBasicInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.UserAccounts
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	if shared.IsInteger(itemId) && itemId != 0 {
		var dataUpdate structs.UserAccounts
		dataBytes, _ := json.Marshal(params.Args["data"])
		_ = json.Unmarshal(dataBytes, &dataUpdate)

		userResponse, err := UpdateUserAccount(itemId, dataUpdate)
		if err != nil {
			return dto.Response{
				Status:  "error",
				Message: err.Error(),
			}, nil
		}

		return dto.ResponseSingle{
			Status:  "success",
			Message: "You updated this item!",
			Item:    userResponse,
		}, nil
	} else {
		userResponse, err := CreateUserAccount(data)
		if err != nil {
			return dto.Response{
				Status:  "error",
				Message: err.Error(),
			}, nil
		}

		return dto.ResponseSingle{
			Status:  "success",
			Message: "You created this item!",
			Item:    userResponse,
		}, nil
	}
}

var UserAccountDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	if !shared.IsInteger(id) || id == 0 {
		return shared.ErrorResponse("You must pass the id"), nil
	}
	user, _ := GetUserAccountById(id.(int))
	user.Active = false
	_, err := UpdateUserAccount(id.(int), *user)
	if err != nil {
		return dto.Response{
			Status:  "error",
			Message: err.Error(),
		}, nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deactivated this user!",
	}, nil
}

func GetUserAccounts(input *dto.GetUserAccountListInput) (*dto.GetUserAccountListResponseMS, error) {
	res := &dto.GetUserAccountListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_ACCOUNTS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateUserAccount(userID int, user structs.UserAccounts) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.USER_ACCOUNTS_ENDPOINT+"/"+strconv.Itoa(userID), user, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func GetUserAccountById(id int) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_ACCOUNTS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func CreateUserAccount(user structs.UserAccounts) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.USER_ACCOUNTS_ENDPOINT, user, res)

	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getRole(id int) (*structs.UserAccountRoles, error) {
	res := &dto.GetUserAccountRoleResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ROLES_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
