package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserAccountsOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []structs.UserAccounts
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	isActive, isActiveOk := params.Args["is_active"].(bool)
	email, emailOk := params.Args["email"].(string)

	if id != nil && id != 0 {
		user, err := r.Repo.GetUserAccountByID(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}
		items = []structs.UserAccounts{*user}
		total = 1
	} else {
		input := dto.GetUserAccountListInput{}
		if page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if isActiveOk {
			input.IsActive = &isActive
		}
		if emailOk && email != "" {
			input.Email = &email
		}

		res, err := r.Repo.GetUserAccounts(&input)
		if err != nil {
			return errors.HandleAPIError(err)
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

func (r *Resolver) UserAccountBasicInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.UserAccounts
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	if itemID != 0 {
		var dataUpdate structs.UserAccounts
		dataBytes, _ := json.Marshal(params.Args["data"])
		_ = json.Unmarshal(dataBytes, &dataUpdate)

		userResponse, err := r.Repo.UpdateUserAccount(itemID, dataUpdate)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.ResponseSingle{
			Status:  "success",
			Message: "You updated this item!",
			Item:    userResponse,
		}, nil
	}
	userResponse, err := r.Repo.CreateUserAccount(data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
		Item:    userResponse,
	}, nil
}

func (r *Resolver) UserAccountDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	if id == 0 {
		return errors.ErrorResponse("You must pass the id"), nil
	}
	user, _ := r.Repo.GetUserAccountByID(id.(int))
	user.Active = false
	_, err := r.Repo.UpdateUserAccount(id.(int), *user)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deactivated this user!",
	}, nil
}
