package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	errors "bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) ExternalReallocationOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		ExternalReallocation, err := r.Repo.GetExternalReallocationByID(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		res, err := buildExternalReallocation(*ExternalReallocation, r)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.ExternalReallocationResponse{res},
			Total:   1,
		}, nil
	}

	input := dto.ExternalReallocationFilter{}
	if value, ok := params.Args["page"].(int); ok && value != 0 {
		input.Page = &value
	}

	if value, ok := params.Args["size"].(int); ok && value != 0 {
		input.Size = &value
	}

	if value, ok := params.Args["source_organization_unit_id"].(int); ok && value != 0 {
		input.SourceOrganizationUnitID = &value
	}

	if value, ok := params.Args["destination_organization_unit_id"].(int); ok && value != 0 {
		input.DestinationOrganizationUnitID = &value
	}

	if value, ok := params.Args["organization_unit_id"].(int); ok && value != 0 {
		input.OrganizationUnitID = &value
	}

	if value, ok := params.Args["status"].(string); ok && value != "" {
		input.Status = &value
	}

	if value, ok := params.Args["requested_by"].(int); ok && value != 0 {
		input.RequestedBy = &value
	}

	if value, ok := params.Args["budget_id"].(int); ok && value != 0 {
		input.BudgetID = &value
	}

	items, total, err := r.Repo.GetExternalReallocationList(input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var resItems []dto.ExternalReallocationResponse
	for _, item := range items {
		resItem, err := buildExternalReallocation(item, r)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		resItems = append(resItems, *resItem)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   resItems,
		Total:   total,
	}, nil
}

func (r *Resolver) ExternalReallocationInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.ExternalReallocation
	response := dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if data.SourceOrganizationUnitID == 0 {

		organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitID == nil {
			return errors.HandleAPPError(fmt.Errorf("user does not have organization unit assigned"))
		}

		data.SourceOrganizationUnitID = *organizationUnitID

	}

	if data.RequestedBy == 0 {
		userProfile, ok := params.Context.Value(config.LoggedInProfileKey).(*structs.UserProfiles)
		if !ok || userProfile == nil {
			return errors.HandleAPPError(fmt.Errorf("error during checking user profile id"))
		}

		data.RequestedBy = userProfile.ID
	}

	var item *structs.ExternalReallocation

	item, err = r.Repo.CreateExternalReallocation(params.Context, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	singleItem, err := buildExternalReallocation(*item, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Item = *singleItem

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	targetUsers, _ := r.Repo.GetUsersByPermission(config.FinanceBudget, config.OperationRead)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	employees, _ := GetEmployeesOfOrganizationUnit(r.Repo, singleItem.SourceOrganizationUnit.ID)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	for _, targetUser := range targetUsers {
		for _, employee := range employees {
			if targetUser.ID != loggedInUser.ID && employee.UserAccountID == targetUser.ID {
				_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
					Content:     "Poslat je novi zahtjev za eksterno preusmjerenje od strane OJ " + singleItem.DestinationOrganizationUnit.Title + ".",
					Module:      "Finansije",
					FromUserID:  loggedInUser.ID,
					ToUserID:    targetUser.ID,
					FromContent: "Službenik za budžet",
					Path:        fmt.Sprintf("/finance/budget/requests/%d", singleItem.ID),
					IsRead:      false,
				})
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}
	}

	return response, nil
}

func (r *Resolver) ExternalReallocationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteExternalReallocation(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) ExternalReallocationOUAcceptResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.ExternalReallocation
	response := dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if data.AcceptedBy == 0 {
		userProfile, ok := params.Context.Value(config.LoggedInProfileKey).(*structs.UserProfiles)
		if !ok || userProfile == nil {
			return errors.HandleAPPError(fmt.Errorf("error during checking user profile id"))
		}

		data.AcceptedBy = userProfile.ID
	}

	var item *structs.ExternalReallocation

	item, err = r.Repo.AcceptOUExternalReallocation(params.Context, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	singleItem, err := buildExternalReallocation(*item, r)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Item = *singleItem

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	targetUsers, _ := r.Repo.GetUsersByPermission(config.FinanceBudget, config.OperationFullAccess)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	for _, targetUser := range targetUsers {
		if targetUser.ID != loggedInUser.ID {
			_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
				Content:     "Poslat je novi zahtjev za eksterno preusmjerenje od strane OJ " + singleItem.DestinationOrganizationUnit.Title + ".",
				Module:      "Finansije",
				FromUserID:  loggedInUser.ID,
				ToUserID:    targetUser.ID,
				FromContent: "Službenik za budžet",
				Path:        fmt.Sprintf("/finance/budget/requests/%d", singleItem.ID),
				IsRead:      false,
			})
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

		}
	}

	return response, nil
}

func (r *Resolver) ExternalReallocationOURejectResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.RejectOUExternalReallocation(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	externalReallocation, err := r.Repo.GetExternalReallocationByID(itemID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	organizationUnit, err := r.Repo.GetOrganizationUnitByID(externalReallocation.SourceOrganizationUnitID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	targetUsers, _ := r.Repo.GetUsersByPermission(config.FinanceBudget, config.OperationRead)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	employees, _ := GetEmployeesOfOrganizationUnit(r.Repo, externalReallocation.DestinationOrganizationUnitID)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	for _, targetUser := range targetUsers {
		for _, employee := range employees {
			if targetUser.ID != loggedInUser.ID && employee.UserAccountID == targetUser.ID {
				_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
					Content:     "Vaš zahtjev za eksterno preusmjerenje je odbijen od strane OJ. " + organizationUnit.Title,
					Module:      "Finansije",
					FromUserID:  loggedInUser.ID,
					ToUserID:    targetUser.ID,
					FromContent: "Službenik za budžet",
					Path:        fmt.Sprintf("/finance/budget/external-reallocation/%d", externalReallocation.ID),
					IsRead:      false,
				})
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You reject this item!",
	}, nil
}

func (r *Resolver) ExternalReallocationSSSAcceptResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.AcceptSSSExternalReallocation(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	targetUsers, _ := r.Repo.GetUsersByPermission(config.FinanceBudget, config.OperationRead)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	externalReallocation, err := r.Repo.GetExternalReallocationByID(itemID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	employees, _ := GetEmployeesOfOrganizationUnit(r.Repo, externalReallocation.DestinationOrganizationUnitID)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	sourceEmployees, _ := GetEmployeesOfOrganizationUnit(r.Repo, externalReallocation.SourceOrganizationUnitID)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	employees = append(employees, sourceEmployees...)

	for _, targetUser := range targetUsers {
		for _, employee := range employees {
			if targetUser.ID != loggedInUser.ID && employee.UserAccountID == targetUser.ID {
				_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
					Content:     "Vaš zahtjev za eksterno preusmjerenje je prihvaćen od strane SSS. ",
					Module:      "Finansije",
					FromUserID:  loggedInUser.ID,
					ToUserID:    targetUser.ID,
					FromContent: "Službenik za budžet",
					Path:        fmt.Sprintf("/finance/budget/external-reallocation/%d", externalReallocation.ID),
					IsRead:      false,
				})
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You accept this item!",
	}, nil
}

func (r *Resolver) ExternalReallocationSSSRejectResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.RejectSSSExternalReallocation(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	externalReallocation, err := r.Repo.GetExternalReallocationByID(itemID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	targetUsers, _ := r.Repo.GetUsersByPermission(config.FinanceBudget, config.OperationRead)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	employees, _ := GetEmployeesOfOrganizationUnit(r.Repo, externalReallocation.DestinationOrganizationUnitID)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	sourceEmployees, _ := GetEmployeesOfOrganizationUnit(r.Repo, externalReallocation.SourceOrganizationUnitID)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	employees = append(employees, sourceEmployees...)

	for _, targetUser := range targetUsers {
		for _, employee := range employees {
			if targetUser.ID != loggedInUser.ID && employee.UserAccountID == targetUser.ID {
				_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
					Content:     "Vaš zahtjev za eksterno preusmjerenje je odbijen od strane SSS. ",
					Module:      "Finansije",
					FromUserID:  loggedInUser.ID,
					ToUserID:    targetUser.ID,
					FromContent: "Službenik za budžet",
					Path:        fmt.Sprintf("/finance/budget/current/external-reallocation/%d", externalReallocation.ID),
					IsRead:      false,
				})
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You reject this item!",
	}, nil
}

func buildExternalReallocation(item structs.ExternalReallocation, r *Resolver) (*dto.ExternalReallocationResponse, error) {

	response := dto.ExternalReallocationResponse{
		ID:                      item.ID,
		Title:                   item.Title,
		DateOfRequest:           item.DateOfRequest,
		Status:                  item.Status,
		DateOfActionDestOrgUnit: item.DateOfActionDestOrgUnit,
		DateOfActionSSS:         item.DateOfActionSSS,
		CreatedAt:               item.CreatedAt,
		UpdatedAt:               item.UpdatedAt,
	}

	if item.SourceOrganizationUnitID != 0 {
		value, _ := r.Repo.GetOrganizationUnitByID(item.SourceOrganizationUnitID)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}*/

		if value != nil {

			dropdown := dto.DropdownSimple{
				ID:    value.ID,
				Title: value.Title,
			}

			response.SourceOrganizationUnit = dropdown
		}
	}

	if item.DestinationOrganizationUnitID != 0 {
		value, _ := r.Repo.GetOrganizationUnitByID(item.DestinationOrganizationUnitID)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}*/

		if value != nil {

			dropdown := dto.DropdownSimple{
				ID:    value.ID,
				Title: value.Title,
			}

			response.DestinationOrganizationUnit = dropdown
		}
	}

	if item.RequestedBy != 0 {
		value, _ := r.Repo.GetUserProfileByID(item.RequestedBy)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}*/

		if value != nil {

			dropdown := dto.DropdownSimple{
				ID:    value.ID,
				Title: value.FirstName + " " + value.LastName,
			}

			response.RequestedBy = dropdown
		}
	}

	if item.AcceptedBy != 0 {
		value, _ := r.Repo.GetUserProfileByID(item.AcceptedBy)
		/*
			if err != nil {
				return nil, errors.Wrap(err, "repo get user profile by id")
			}
		*/

		if value != nil {
			dropdown := dto.DropdownSimple{
				ID:    value.ID,
				Title: value.FirstName + " " + value.LastName,
			}

			response.AcceptedBy = dropdown
		}
	}

	if item.FileID != 0 {
		value, _ := r.Repo.GetFileByID(item.FileID)

		/*
			if err != nil {
				return nil, errors.Wrap(err, "repo get file by id")
			}
		*/
		if value != nil {
			dropdown := dto.FileDropdownSimple{
				ID:   value.ID,
				Name: value.Name,
				Type: *value.Type,
			}

			response.File = dropdown
		}
	}

	if item.DestinationOrgUnitFileID != 0 {
		value, _ := r.Repo.GetFileByID(item.DestinationOrgUnitFileID)
		/*
			if err != nil {
				return nil, errors.Wrap(err, "repo get file by id")
			}
		*/
		if value != nil {
			dropdown := dto.FileDropdownSimple{
				ID:   value.ID,
				Name: value.Name,
				Type: *value.Type,
			}

			response.DestinationOrgUnitFile = dropdown
		}
	}

	if item.SSSFileID != 0 {
		value, _ := r.Repo.GetFileByID(item.SSSFileID)
		/*
			if err != nil {
				return nil, errors.Wrap(err, "repo get file by id")
			}
		*/
		if value != nil {
			dropdown := dto.FileDropdownSimple{
				ID:   value.ID,
				Name: value.Name,
				Type: *value.Type,
			}

			response.SSSFile = dropdown
		}
	}

	if item.BudgetID != 0 {
		value, err := r.Repo.GetBudget(item.BudgetID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get budget")
		}

		dropdown := dto.DropdownSimple{
			ID:    value.ID,
			Title: strconv.Itoa(value.Year),
		}

		response.Budget = dropdown
	}

	for _, orderItem := range item.Items {
		builtItem, err := buildExternalReallocationItem(orderItem, r)

		if err != nil {
			return nil, errors.Wrap(err, "repo build external reallocation item")
		}

		response.Items = append(response.Items, *builtItem)
	}

	return &response, nil
}

func buildExternalReallocationItem(item structs.ExternalReallocationItem, r *Resolver) (*dto.ExternalReallocationItemResponse, error) {
	response := dto.ExternalReallocationItemResponse{
		ID:     item.ID,
		Amount: item.Amount,
	}

	if item.SourceAccountID != 0 {
		value, _ := r.Repo.GetAccountItemByID(item.SourceAccountID)
		/*
			if err != nil {
				return nil, errors.Wrap(err, "repo get account item by id")
			}
		*/
		if value != nil {
			dropdown := dto.DropdownSimple{
				ID:    value.ID,
				Title: value.Title,
			}

			response.SourceAccount = dropdown
		}
	}

	if item.DestinationAccountID != 0 {
		value, _ := r.Repo.GetAccountItemByID(item.DestinationAccountID)

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get account item by id")
		}*/

		if value != nil {

			dropdown := dto.DropdownSimple{
				ID:    value.ID,
				Title: value.Title,
			}

			response.DestinationAccount = dropdown
		}
	}

	return &response, nil
}
