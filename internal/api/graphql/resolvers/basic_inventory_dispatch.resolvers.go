package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/internal/api/websockets/notifications"
	"bff/shared"
	"bff/structs"
	"context"
	"encoding/json"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) BasicInventoryDispatchOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.InventoryDispatchResponse
	var filter dto.InventoryDispatchFilter
	organizationUnitID, _ := params.Context.Value(config.OrganizationUnitIDKey).(*int)

	if organizationUnitID != nil {
		filter.OrganizationUnitID = organizationUnitID
	}

	if id, ok := params.Args["id"].(int); ok && id != 0 {
		filter.ID = &id
	}

	if typeFilter, ok := params.Args["type"].(string); ok && typeFilter != "" {
		filter.Type = &typeFilter
	}

	if accepted, ok := params.Args["accepted"].(bool); ok {
		filter.Accepted = &accepted
	}

	if source, ok := params.Args["source_organization_unit_id"].(int); ok && source != 0 {
		filter.SourceOrganizationUnitID = &source
	}

	if page, ok := params.Args["page"].(int); ok && page != 0 {
		filter.Page = &page
	}

	if size, ok := params.Args["size"].(int); ok && size != 0 {
		filter.Size = &size
	}

	if inventory_type, ok := params.Args["inventory_type"].(string); ok && inventory_type != "" {
		filter.InventoryType = &inventory_type
	}

	data, err := r.Repo.GetAllInventoryDispatches(filter)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, item := range data.Data {
		resItem, err := buildInventoryDispatchResponse(r.Repo, item)
		items = append(items, resItem)

		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   data.Total,
		Items:   items,
	}, nil
}

func (r *Resolver) BasicInventoryDispatchInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.BasicInventoryDispatchItem
	var items *dto.InventoryDispatchResponse

	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	if shared.IsInteger(data.Id) && data.Id != 0 {
		itemRes, err := r.Repo.UpdateDispatchItem(data.Id, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"

		items, err = buildInventoryDispatchResponse(r.Repo, itemRes)

		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		itemRes, err := r.Repo.CreateDispatchItem(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		err = sendInventoryDispatchNotification(params.Context, r.Repo, r.NotificationsService, itemRes.SourceOrganizationUnitId, itemRes.TargetOrganizationUnitId)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		items, err = buildInventoryDispatchResponse(r.Repo, itemRes)

		if err != nil {
			return errors.HandleAPIError(err)
		}

	}

	response.Item = items
	return response, nil
}

func sendInventoryDispatchNotification(
	ctx context.Context,
	r repository.MicroserviceRepositoryInterface,
	notificationService *notifications.Websockets,
	sourceOrganizationUnitID int,
	targetOrganziationUnitID int,
) error {
	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	sourceOrganizationUnit, _ := r.GetOrganizationUnitById(sourceOrganizationUnitID)
	employees, _ := GetEmployeesOfOrganizationUnit(r, targetOrganziationUnitID)

	for _, employee := range employees {
		userAccount, _ := r.GetUserAccountById(employee.UserAccountId)
		if userAccount.RoleId == structs.UserRoleManagerOJ {
			_, err := notificationService.CreateNotification(&structs.Notifications{
				Content:     "Kreiran je revers. Potrebno je da ga odobrite ili odbijete.",
				Module:      "Osnovna sredstva",
				FromUserID:  loggedInUser.Id,
				ToUserID:    userAccount.Id,
				FromContent: "Menadžer organizacione jedinice - " + sourceOrganizationUnit.Title,
				Path:        "/inventory/movable-inventory/receive-inventory",
				Data:        nil,
				IsRead:      false,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func createInventoryDispatchApprovalnotification(
	ctx context.Context,
	r repository.MicroserviceRepositoryInterface,
	notificationService *notifications.Websockets,
	sourceOrganizationUnitID int,
	targetOrganziationUnitID int,
	isAccepted bool,
) error {
	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	targetOrganizationUnit, _ := r.GetOrganizationUnitById(targetOrganziationUnitID)
	employees, _ := GetEmployeesOfOrganizationUnit(r, sourceOrganizationUnitID)

	var content string

	if isAccepted {
		content = "Revers je prihvaćen."
	} else {
		content = "Revers je obijen."
	}

	for _, employee := range employees {
		userAccount, _ := r.GetUserAccountById(employee.UserAccountId)
		if userAccount.RoleId == structs.UserRoleManagerOJ {
			_, err := notificationService.CreateNotification(&structs.Notifications{
				Content:     content,
				Module:      "Osnovna sredstva",
				FromUserID:  loggedInUser.Id,
				ToUserID:    userAccount.Id,
				FromContent: "Menadžer organizacione jedinice - " + targetOrganizationUnit.Title,
				Path:        "/inventory/movable-inventory",
				Data:        nil,
				IsRead:      false,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Resolver) BasicInventoryDispatchDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	dispatch, err := r.Repo.GetDispatchItemByID(itemId)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	if dispatch.FileId != 0 {
		err := r.Repo.DeleteFile(dispatch.FileId)

		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	err = r.Repo.DeleteInventoryDispatch(itemId)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	err = createInventoryDispatchApprovalnotification(params.Context, r.Repo, r.NotificationsService, dispatch.SourceOrganizationUnitId, dispatch.TargetOrganizationUnitId, false)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) BasicInventoryDispatchAcceptResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["dispatch_id"].(int)

	loggedInProfile, _ := params.Context.Value(config.LoggedInProfileKey).(*structs.UserProfiles)

	dispatch, err := r.Repo.GetDispatchItemByID(id)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	dispatch.IsAccepted = true
	dispatch.TargetUserProfileId = loggedInProfile.Id

	filter := dto.DispatchInventoryItemFilter{
		DispatchID: &dispatch.Id,
	}

	itemDispatchList, _ := r.Repo.GetMyInventoryDispatchesItems(&filter)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, itemDispatch := range itemDispatchList {
		dispatch.InventoryId = append(dispatch.InventoryId, itemDispatch.InventoryId)

		item, err := r.Repo.GetInventoryItem(itemDispatch.InventoryId)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		item.TargetOrganizationUnitId = dispatch.TargetOrganizationUnitId

		_, err = r.Repo.UpdateInventoryItem(item.Id, item)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}
	currentDate := time.Now()
	dispatch.Date = currentDate.Format("2006-01-02T15:04:05.999999Z07:00")
	_, err = r.Repo.UpdateDispatchItem(dispatch.Id, dispatch)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	err = createInventoryDispatchApprovalnotification(params.Context, r.Repo, r.NotificationsService, dispatch.SourceOrganizationUnitId, dispatch.TargetOrganizationUnitId, true)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You accepted this item!",
	}, nil

}

func buildInventoryDispatchResponse(repo repository.MicroserviceRepositoryInterface, item *structs.BasicInventoryDispatchItem) (*dto.InventoryDispatchResponse, error) {
	settings, err := repo.GetDropdownSettingById(item.OfficeId)
	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok && apiErr.StatusCode != 404 {
			return nil, err
		}
	}

	settingDropdownOfficeId := dto.DropdownSimple{}
	if settings != nil {
		settingDropdownOfficeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
	}

	sourceUserDropdown := dto.DropdownSimple{}
	if item.SourceUserProfileId != 0 {
		user, _ := repo.GetUserProfileById(item.SourceUserProfileId)

		if user != nil {
			sourceUserDropdown = dto.DropdownSimple{Id: user.Id, Title: user.FirstName + " " + user.LastName}
		}
	}

	targetUserDropdown := dto.DropdownSimple{}
	if item.TargetUserProfileId != 0 {
		user, _ := repo.GetUserProfileById(item.TargetUserProfileId)

		if user != nil {
			targetUserDropdown = dto.DropdownSimple{Id: user.Id, Title: user.FirstName + " " + user.LastName}
		}
	}

	sourceOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.SourceOrganizationUnitId != 0 {
		sourceOrganizationUnit, _ := repo.GetOrganizationUnitById(item.SourceOrganizationUnitId)

		if sourceOrganizationUnit != nil {
			sourceOrganizationUnitDropdown = dto.DropdownSimple{Id: sourceOrganizationUnit.Id, Title: sourceOrganizationUnit.Title}
		}
	}
	city := ""
	targetOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.TargetOrganizationUnitId != 0 {
		targetOrganizationUnit, _ := repo.GetOrganizationUnitById(item.TargetOrganizationUnitId)

		if targetOrganizationUnit != nil {
			city = targetOrganizationUnit.City
			targetOrganizationUnitDropdown = dto.DropdownSimple{Id: targetOrganizationUnit.Id, Title: targetOrganizationUnit.Title}
		}
	}

	filter := dto.DispatchInventoryItemFilter{
		DispatchID: &item.Id,
	}

	dispatchItems, _ := repo.GetMyInventoryDispatchesItems(&filter)
	/*if err != nil {
		// return nil, err
	}*/

	inventoryItems := []dto.BasicInventoryResponseItem{}

	if dispatchItems != nil {
		for i := 0; i < len(dispatchItems); i++ {
			itemInventory, err := repo.GetInventoryItem(dispatchItems[i].InventoryId)
			if err != nil {
				return nil, err
			}

			itemArr := dto.BasicInventoryResponseItem{
				Id:              itemInventory.Id,
				Type:            itemInventory.Type,
				InventoryNumber: itemInventory.InventoryNumber,
				Title:           itemInventory.Title,
				GrossPrice:      itemInventory.GrossPrice,
				SerialNumber:    itemInventory.SerialNumber,
				Location:        itemInventory.Location,
			}
			inventoryItems = append(inventoryItems, itemArr)
		}
	}

	var fileDropdown dto.FileDropdownSimple

	if item.FileId != 0 {
		file, err := repo.GetFileByID(item.FileId)

		if err != nil {
			return nil, err
		}

		fileDropdown.Id = file.ID
		fileDropdown.Name = file.Name
		fileDropdown.Type = *file.Type
	}

	res := dto.InventoryDispatchResponse{
		ID:                     item.Id,
		DispatchID:             item.DispatchID,
		Type:                   item.Type,
		SerialNumber:           item.SerialNumber,
		Office:                 settingDropdownOfficeId,
		SourceUserProfile:      sourceUserDropdown,
		IsAccepted:             item.IsAccepted,
		TargetUserProfile:      targetUserDropdown,
		SourceOrganizationUnit: sourceOrganizationUnitDropdown,
		TargetOrganizationUnit: targetOrganizationUnitDropdown,
		InventoryType:          item.InventoryType,
		Inventory:              inventoryItems,
		Date:                   item.Date,
		City:                   city,
		CreatedAt:              item.CreatedAt,
		UpdatedAt:              item.UpdatedAt,
		DispatchDescription:    item.DispatchDescription,
		File:                   fileDropdown,
	}

	return &res, nil
}
