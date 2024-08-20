package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) LogsOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	moduleStr, _ := params.Args["module"].(string)

	module := config.Module(moduleStr)

	if id, ok := params.Args["id"].(int); ok && id != 0 {
		log, err := r.Repo.GetLog(module, id)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		responseItem, err := buildLogItem(r, *log)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.LogResponse{responseItem},
			Total:   1,
		}, nil
	}

	input := dto.LogFilterDTO{}
	input.Module = module
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["user_id"].(int); ok && value != 0 {
		input.UserID = &value
	}

	if value, ok := params.Args["item_id"].(int); ok && value != 0 {
		input.ItemID = &value
	}

	if value, ok := params.Args["search"].(string); ok && value != "" {
		input.Search = &value
	}

	if value, ok := params.Args["operation"].(string); ok && value != "" {
		input.Operation = &value
	}

	if value, ok := params.Args["entity"].(string); ok && value != "" {
		input.Entity = &value
	}

	items, total, err := r.Repo.GetLogs(input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var resItems []dto.LogResponse
	for _, item := range items {
		resItem, err := buildLogItem(r, item)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		resItems = append(resItems, *resItem)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   int(total),
		Items:   resItems,
	}, nil
}

func (r *Resolver) ErrorLogsOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	moduleStr, _ := params.Args["module"].(string)

	module := config.Module(moduleStr)

	if id, ok := params.Args["id"].(int); ok && id != 0 {
		log, err := r.Repo.GetErrorLog(module, id)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []structs.ErrorLogs{*log},
			Total:   1,
		}, nil
	}

	input := dto.ErrorLogFilterDTO{}
	input.Module = module
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["entity"].(string); ok && value != "" {
		input.Entity = &value
	}

	if value, ok := params.Args["date_of_start"].(string); ok && value != "" {
		date, err := parseDate(value)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		input.DateOfStart = &date
	}

	if value, ok := params.Args["date_of_end"].(string); ok && value != "" {
		date, err := parseDate(value)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		input.DateOfEnd = &date
	}

	items, total, err := r.Repo.GetErrorLogs(input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   int(total),
		Items:   items,
	}, nil
}

func buildLogItem(r *Resolver, log structs.Logs) (*dto.LogResponse, error) {
	response := dto.LogResponse{
		ID:        log.ID,
		ChangedAt: log.ChangedAt,
		ItemID:    log.ItemID,
		Operation: log.Operation,
		Entity:    log.Entity,
		OldState:  log.OldState,
		NewState:  log.NewState,
	}

	if log.UserID != 0 {
		user, _ := r.Repo.GetUserAccountByID(log.UserID)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get user account by id")
		}*/

		if user != nil {

			response.User = dto.DropdownSimple{
				ID:    user.ID,
				Title: user.Email,
			}
		}

		userProfile, _ := r.Repo.GetUserProfileByUserAccountID(log.UserID)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by user account id")
		}*/

		if userProfile != nil {
			response.UserProfile = dto.DropdownSimple{
				ID:    userProfile.ID,
				Title: userProfile.FirstName + " " + userProfile.LastName,
			}
		}
	}

	return &response, nil
}
