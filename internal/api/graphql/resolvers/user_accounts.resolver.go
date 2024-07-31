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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		if user.RoleID != nil {
			role, err := r.Repo.GetRole(*user.RoleID)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			user.Role = *role
		}
		items = []structs.UserAccounts{*user}
		total = 1
	} else {
		input := dto.GetUserAccountListInput{}
		if page != nil && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if size != nil && size.(int) > 0 {
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		for i := 0; i < len(res.Data); i++ {
			if res.Data[i].RoleID != nil {
				role, err := r.Repo.GetRole(*res.Data[i].RoleID)
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
				res.Data[i].Role = *role
			}
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

		userResponse, err := r.Repo.UpdateUserAccount(params.Context, itemID, dataUpdate)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if userResponse.RoleID != nil {
			role, err := r.Repo.GetRole(*userResponse.RoleID)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			userResponse.Role.ID = role.ID
			userResponse.Role.Title = role.Title
		}

		return dto.ResponseSingle{
			Status:  "success",
			Message: "You updated this item!",
			Item:    userResponse,
		}, nil
	}
	/*userResponse, err := r.Repo.CreateUserAccount(params.Context, data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	return dto.ResponseSingle{
		Status:  "failed",
		Message: "You can not create this item!",
		//Item:    userResponse,
	}, nil
}

func (r *Resolver) UserAccountDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	if id == 0 {
		return errors.ErrorResponse("You must pass the id"), nil
	}
	user, _ := r.Repo.GetUserAccountByID(id.(int))
	user.Active = false
	_, err := r.Repo.UpdateUserAccount(params.Context, id.(int), *user)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deactivated this user!",
	}, nil
}
