package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/internal/api/websockets/notifications"
	"bff/structs"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

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

	if intentoryType, ok := params.Args["inventory_type"].(string); ok && intentoryType != "" {
		filter.InventoryType = &intentoryType
	}

	data, err := r.Repo.GetAllInventoryDispatches(filter)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	for _, item := range data.Data {
		resItem, err := buildInventoryDispatchResponse(r.Repo, item, *organizationUnitID)
		items = append(items, resItem)

		if err != nil {
			return apierrors.HandleAPIError(err)
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
		return apierrors.HandleAPIError(err)
	}

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	if data.ID != 0 {
		itemRes, err := r.Repo.UpdateDispatchItem(data.ID, &data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		items, err = buildInventoryDispatchResponse(r.Repo, itemRes, *organizationUnitID)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	} else {
		if data.Type != "convert" {
			itemRes, err := r.Repo.CreateDispatchItem(&data)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			if itemRes.Type == "return-revers" {
				// return revers is always one by one item.
				item, err := r.Repo.GetInventoryItem(data.InventoryID[0])
				if err != nil {
					return apierrors.HandleAPIError(err)
				}
				err = sendInventoryReturnReversDispatchNotification(params.Context, r.Repo, r.NotificationsService, itemRes.SourceOrganizationUnitID, item.OrganizationUnitID, item.ID)
				if err != nil {
					return apierrors.HandleAPIError(err)
				}

			} else if itemRes.Type == "revers" {
				err = sendInventoryDispatchNotification(params.Context, r.Repo, r.NotificationsService, itemRes.SourceOrganizationUnitID, itemRes.TargetOrganizationUnitID)
				if err != nil {
					return apierrors.HandleAPIError(err)
				}
			}

			response.Message = "You created this item!"
			items, err = buildInventoryDispatchResponse(r.Repo, itemRes, *organizationUnitID)

			if err != nil {
				return apierrors.HandleAPIError(err)
			}
		} else {
			organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
			if !ok || organizationUnitID == nil {
				return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
			}

			input := structs.BasicInventoryDispatchItem{
				Type:                     data.Type,
				SourceUserProfileID:      data.SourceUserProfileID,
				SourceOrganizationUnitID: *organizationUnitID,
				Date:                     data.Date,
				InventoryID:              data.InventoryID,
				FileID:                   data.FileID,
			}

			_, err := r.Repo.CreateDispatchItem(&input)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			for _, itemID := range data.InventoryID {
				item, err := r.Repo.GetInventoryItem(itemID)

				if err != nil {
					return apierrors.HandleAPIError(err)
				}

				item.IsExternalDonation = false
				item.DonationFiles = append(item.DonationFiles, data.DonationFiles...)

				_, err = r.Repo.UpdateInventoryItem(itemID, item)

				if err != nil {
					return apierrors.HandleAPIError(err)
				}
			}
		}
	}

	response.Item = items
	return response, nil
}

func sendInventoryReturnReversDispatchNotification(
	ctx context.Context,
	r repository.MicroserviceRepositoryInterface,
	notificationService *notifications.Websockets,
	sourceOrganizationUnitID int,
	targetOrganziationUnitID int,
	itemID int,
) error {
	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	sourceOrganizationUnit, _ := r.GetOrganizationUnitByID(sourceOrganizationUnitID)
	employees, _ := GetEmployeesOfOrganizationUnit(r, targetOrganziationUnitID)

	for _, employee := range employees {
		userAccount, _ := r.GetUserAccountByID(employee.UserAccountID)
		if userAccount.RoleID == structs.UserRoleManagerOJ {
			_, err := notificationService.CreateNotification(&structs.Notifications{
				Content:     "Izvršen je povrat sredstva.",
				Module:      "Osnovna sredstva",
				FromUserID:  loggedInUser.ID,
				ToUserID:    userAccount.ID,
				FromContent: "Menadžer organizacione jedinice - " + sourceOrganizationUnit.Title,
				Path:        "/inventory/movable-inventory/" + strconv.Itoa(itemID),
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

func sendInventoryDispatchNotification(
	ctx context.Context,
	r repository.MicroserviceRepositoryInterface,
	notificationService *notifications.Websockets,
	sourceOrganizationUnitID int,
	targetOrganziationUnitID int,
) error {
	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	sourceOrganizationUnit, _ := r.GetOrganizationUnitByID(sourceOrganizationUnitID)
	employees, _ := GetEmployeesOfOrganizationUnit(r, targetOrganziationUnitID)

	for _, employee := range employees {
		userAccount, _ := r.GetUserAccountByID(employee.UserAccountID)
		if userAccount.RoleID == structs.UserRoleManagerOJ {
			_, err := notificationService.CreateNotification(&structs.Notifications{
				Content:     "Kreiran je revers. Potrebno je da ga odobrite ili odbijete.",
				Module:      "Osnovna sredstva",
				FromUserID:  loggedInUser.ID,
				ToUserID:    userAccount.ID,
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
	targetOrganizationUnit, _ := r.GetOrganizationUnitByID(targetOrganziationUnitID)
	employees, _ := GetEmployeesOfOrganizationUnit(r, sourceOrganizationUnitID)

	var content string

	if isAccepted {
		content = "Revers je prihvaćen."
	} else {
		content = "Revers je obijen."
	}

	for _, employee := range employees {
		userAccount, _ := r.GetUserAccountByID(employee.UserAccountID)
		if userAccount.RoleID == structs.UserRoleManagerOJ {
			_, err := notificationService.CreateNotification(&structs.Notifications{
				Content:     content,
				Module:      "Osnovna sredstva",
				FromUserID:  loggedInUser.ID,
				ToUserID:    userAccount.ID,
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
	itemID := params.Args["id"].(int)

	dispatch, err := r.Repo.GetDispatchItemByID(itemID)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if dispatch.FileID != 0 {
		err := r.Repo.DeleteFile(dispatch.FileID)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	}

	err = r.Repo.DeleteInventoryDispatch(itemID)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	err = createInventoryDispatchApprovalnotification(params.Context, r.Repo, r.NotificationsService, dispatch.SourceOrganizationUnitID, dispatch.TargetOrganizationUnitID, false)
	if err != nil {
		return apierrors.HandleAPIError(err)
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
		return apierrors.HandleAPIError(err)
	}

	dispatch.IsAccepted = true
	dispatch.TargetUserProfileID = loggedInProfile.ID

	filter := dto.DispatchInventoryItemFilter{
		DispatchID: &dispatch.ID,
	}

	itemDispatchList, _ := r.Repo.GetMyInventoryDispatchesItems(&filter)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	for _, itemDispatch := range itemDispatchList {
		dispatch.InventoryID = append(dispatch.InventoryID, itemDispatch.InventoryID)

		item, err := r.Repo.GetInventoryItem(itemDispatch.InventoryID)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		if dispatch.Type == "revers" {
			item.TargetOrganizationUnitID = dispatch.TargetOrganizationUnitID

			if item.TargetOrganizationUnitID != 0 {
				_, err = r.Repo.UpdateInventoryItem(item.ID, item)
				if err != nil {
					return apierrors.HandleAPIError(err)
				}
			}
		}

		if dispatch.Type == "return-revers" {
			item.TargetOrganizationUnitID = 0

			_, err = r.Repo.UpdateInventoryItem(item.ID, item)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
		}

	}
	//currentDate := time.Now()
	//dispatch.Date = currentDate.Format("2006-01-02T15:04:05.999999Z07:00")
	_, err = r.Repo.UpdateDispatchItem(dispatch.ID, dispatch)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	err = createInventoryDispatchApprovalnotification(params.Context, r.Repo, r.NotificationsService, dispatch.SourceOrganizationUnitID, dispatch.TargetOrganizationUnitID, true)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You accepted this item!",
	}, nil

}

func buildInventoryDispatchResponse(repo repository.MicroserviceRepositoryInterface, item *structs.BasicInventoryDispatchItem, organizationUnitID int) (*dto.InventoryDispatchResponse, error) {
	settings, err := repo.GetDropdownSettingByID(item.OfficeID)
	if err != nil {
		if apiErr, ok := err.(*apierrors.APIError); ok && apiErr.StatusCode != 404 {
			return nil, err
		}
	}

	settingDropdownOfficeID := dto.DropdownSimple{}
	if settings != nil {
		settingDropdownOfficeID = dto.DropdownSimple{ID: settings.ID, Title: settings.Title}
	}

	sourceUserDropdown := dto.DropdownSimple{}
	if item.SourceUserProfileID != 0 && item.Type != "revers" {
		user, _ := repo.GetUserProfileByID(item.SourceUserProfileID)

		if user != nil {
			sourceUserDropdown = dto.DropdownSimple{ID: user.ID, Title: user.FirstName + " " + user.LastName}
		}
	}

	targetUserDropdown := dto.DropdownSimple{}
	if item.TargetUserProfileID != 0 && item.Type != "revers" {
		user, _ := repo.GetUserProfileByID(item.TargetUserProfileID)

		if user != nil {
			targetUserDropdown = dto.DropdownSimple{ID: user.ID, Title: user.FirstName + " " + user.LastName}
		}
	}

	sourceOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.SourceOrganizationUnitID != 0 {
		sourceOrganizationUnit, _ := repo.GetOrganizationUnitByID(item.SourceOrganizationUnitID)

		if sourceOrganizationUnit != nil {
			sourceOrganizationUnitDropdown = dto.DropdownSimple{ID: sourceOrganizationUnit.ID, Title: sourceOrganizationUnit.Title}
		}
	}
	city := ""
	targetOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.TargetOrganizationUnitID != 0 {
		targetOrganizationUnit, _ := repo.GetOrganizationUnitByID(item.TargetOrganizationUnitID)

		if targetOrganizationUnit != nil {
			city = targetOrganizationUnit.City
			targetOrganizationUnitDropdown = dto.DropdownSimple{ID: targetOrganizationUnit.ID, Title: targetOrganizationUnit.Title}
		}
	}

	filter := dto.DispatchInventoryItemFilter{
		DispatchID: &item.ID,
	}

	dispatchItems, _ := repo.GetMyInventoryDispatchesItems(&filter)
	/*if err != nil {
		// return nil, err
	}*/

	inventoryItems := []dto.BasicInventoryResponseItem{}

	if dispatchItems != nil {
		for i := 0; i < len(dispatchItems); i++ {
			itemInventory, err := repo.GetInventoryItem(dispatchItems[i].InventoryID)
			if err != nil {
				return nil, err
			}

			itemArr := dto.BasicInventoryResponseItem{
				ID:              itemInventory.ID,
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

	if item.FileID != 0 {
		file, err := repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, err
		}

		fileDropdown.ID = file.ID
		fileDropdown.Name = file.Name
		fileDropdown.Type = *file.Type
	}

	res := dto.InventoryDispatchResponse{
		ID:                     item.ID,
		DispatchID:             item.DispatchID,
		Type:                   item.Type,
		SerialNumber:           item.SerialNumber,
		Office:                 settingDropdownOfficeID,
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
