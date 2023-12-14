package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"bff/websocketmanager"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

var BasicInventoryDispatchOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
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

	data, err := getAllInventoryDispatches(filter)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, item := range data.Data {
		resItem, err := buildInventoryDispatchResponse(item)
		items = append(items, resItem)

		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   data.Total,
		Items:   items,
	}, nil
}

var BasicInventoryDispatchInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.BasicInventoryDispatchItem
	var items *dto.InventoryDispatchResponse

	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	if shared.IsInteger(data.Id) && data.Id != 0 {
		itemRes, err := updateDispatchItem(data.Id, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You updated this item!"

		items, err = buildInventoryDispatchResponse(itemRes)

		if err != nil {
			return shared.HandleAPIError(err)
		}
	} else {
		itemRes, err := createDispatchItem(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		err = sendInventoryDispatchNotification(params.Context, itemRes.SourceOrganizationUnitId, itemRes.TargetOrganizationUnitId)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		items, err = buildInventoryDispatchResponse(itemRes)

		if err != nil {
			return shared.HandleAPIError(err)
		}

	}

	response.Item = items
	return response, nil
}

func sendInventoryDispatchNotification(ctx context.Context, sourceOrganizationUnitID int, targetOrganziationUnitID int) error {
	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	sourceOrganizationUnit, _ := getOrganizationUnitById(sourceOrganizationUnitID)
	employees, _ := getEmployeesOfOrganizationUnit(targetOrganziationUnitID)

	for _, employee := range employees {
		userAccount, _ := GetUserAccountById(employee.UserAccountId)
		if userAccount.RoleId == structs.UserRoleManagerOJ {
			_, err := websocketmanager.CreateNotification(&structs.Notifications{
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

func sendInventoryDispatchAcceptNotification(ctx context.Context, sourceOrganizationUnitID int, targetOrganziationUnitID int) error {
	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	targetOrganizationUnit, _ := getOrganizationUnitById(targetOrganziationUnitID)
	employees, _ := getEmployeesOfOrganizationUnit(sourceOrganizationUnitID)

	for _, employee := range employees {
		userAccount, _ := GetUserAccountById(employee.UserAccountId)
		if userAccount.RoleId == structs.UserRoleManagerOJ {
			_, err := websocketmanager.CreateNotification(&structs.Notifications{
				Content:     "Revers je prihvaćen.",
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

var BasicInventoryDispatchDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	dispatch, err := getDispatchItemByID(itemId)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	if dispatch.FileId != 0 {
		err := deleteFile(dispatch.FileId)

		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	err = deleteInventoryDispatch(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

var BasicInventoryDispatchAcceptResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["dispatch_id"].(int)

	loggedInProfile, _ := params.Context.Value(config.LoggedInProfileKey).(*structs.UserProfiles)

	dispatch, err := getDispatchItemByID(id)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	dispatch.IsAccepted = true
	dispatch.TargetUserProfileId = loggedInProfile.Id

	filter := dto.DispatchInventoryItemFilter{
		DispatchID: &dispatch.Id,
	}

	itemDispatchList, _ := getMyInventoryDispatchesItems(&filter)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, itemDispatch := range itemDispatchList {
		dispatch.InventoryId = append(dispatch.InventoryId, itemDispatch.InventoryId)

		item, err := getInventoryItem(itemDispatch.InventoryId)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		item.TargetOrganizationUnitId = dispatch.TargetOrganizationUnitId

		_, err = updateInventoryItem(item.Id, item)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	}
	currentDate := time.Now()
	dispatch.Date = currentDate.Format("2006-01-02T15:04:05.999999Z07:00")
	_, err = updateDispatchItem(dispatch.Id, dispatch)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	err = sendInventoryDispatchAcceptNotification(params.Context, dispatch.SourceOrganizationUnitId, dispatch.TargetOrganizationUnitId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You accepted this item!",
	}, nil

}

func getAllInventoryDispatches(filter dto.InventoryDispatchFilter) (*dto.GetAllBasicInventoryDispatches, error) {
	res := &dto.GetAllBasicInventoryDispatches{}
	_, err := shared.MakeAPIRequest("GET", config.INVENTORY_DISPATCH_ENDOPOINT, filter, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func buildInventoryDispatchResponse(item *structs.BasicInventoryDispatchItem) (*dto.InventoryDispatchResponse, error) {
	settings, err := getDropdownSettingById(item.OfficeId)
	if err != nil {
		if apiErr, ok := err.(*shared.APIError); ok && apiErr.StatusCode != 404 {
			return nil, err
		}
	}

	settingDropdownOfficeId := dto.DropdownSimple{}
	if settings != nil {
		settingDropdownOfficeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
	}

	sourceUserDropdown := dto.DropdownSimple{}
	if item.SourceUserProfileId != 0 {
		user, _ := getUserProfileById(item.SourceUserProfileId)

		if user != nil {
			sourceUserDropdown = dto.DropdownSimple{Id: user.Id, Title: user.FirstName + " " + user.LastName}
		}
	}

	targetUserDropdown := dto.DropdownSimple{}
	if item.TargetUserProfileId != 0 {
		user, _ := getUserProfileById(item.TargetUserProfileId)

		if user != nil {
			targetUserDropdown = dto.DropdownSimple{Id: user.Id, Title: user.FirstName + " " + user.LastName}
		}
	}

	sourceOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.SourceOrganizationUnitId != 0 {
		sourceOrganizationUnit, _ := getOrganizationUnitById(item.SourceOrganizationUnitId)

		if sourceOrganizationUnit != nil {
			sourceOrganizationUnitDropdown = dto.DropdownSimple{Id: sourceOrganizationUnit.Id, Title: sourceOrganizationUnit.Title}
		}
	}
	city := ""
	targetOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.TargetOrganizationUnitId != 0 {
		targetOrganizationUnit, _ := getOrganizationUnitById(item.TargetOrganizationUnitId)

		if targetOrganizationUnit != nil {
			city = targetOrganizationUnit.City
			targetOrganizationUnitDropdown = dto.DropdownSimple{Id: targetOrganizationUnit.Id, Title: targetOrganizationUnit.Title}
		}
	}

	filter := dto.DispatchInventoryItemFilter{
		DispatchID: &item.Id,
	}

	dispatchItems, _ := getMyInventoryDispatchesItems(&filter)
	/*if err != nil {
		// return nil, err
	}*/

	inventoryItems := []dto.BasicInventoryResponseItem{}

	if dispatchItems != nil {
		for i := 0; i < len(dispatchItems); i++ {
			itemInventory, err := getInventoryItem(dispatchItems[i].InventoryId)
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
		file, err := getFileByID(item.FileId)

		if err != nil {
			return nil, err
		}

		fileDropdown.Id = file.ID
		fileDropdown.Name = file.Name
		fileDropdown.Type = *file.Type
	}

	res := dto.InventoryDispatchResponse{
		ID:                     item.Id,
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

func getMyInventoryDispatchesItems(filter *dto.DispatchInventoryItemFilter) ([]*structs.BasicInventoryDispatchItemsItem, error) {
	res := &dto.GetAllBasicInventoryDispatchItems{}

	_, err := shared.MakeAPIRequest("GET", config.INVENTORY_DISPATCH_ITEMS_ENDOPOINT, filter, res)

	if err != nil {
		fmt.Printf("Fetching Inventory items failed because of this error - %s.\n", err)
		return nil, err
	}

	return res.Data, nil
}

func createDispatchItem(item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error) {
	res := dto.GetBasicInventoryDispatch{}

	_, err := shared.MakeAPIRequest("POST", config.INVENTORY_DISPATCH_ENDOPOINT, item, &res)
	if err != nil {
		return nil, err
	}

	if item.InventoryId != nil {
		for i := 0; i < len(item.InventoryId); i++ {
			itemDispatch := structs.BasicInventoryDispatchItemsItem{
				InventoryId: item.InventoryId[i],
				DispatchId:  res.Data.Id,
			}

			if item.Type != "revers" {
				inventory, err := getInventoryItem(item.InventoryId[i])
				if err != nil {
					return nil, err
				}

				targetOrganizationUnitID := 0
				targetUserProfileID := 0
				officeID := 0

				if item.Type == "allocation" {
					targetOrganizationUnitID = item.TargetOrganizationUnitId
					targetUserProfileID = item.TargetUserProfileId
					officeID = item.OfficeId
				}
				if item.Type == "return" {
					targetOrganizationUnitID = item.TargetOrganizationUnitId
				}

				inventory.TargetOrganizationUnitId = targetOrganizationUnitID
				inventory.TargetUserProfileId = targetUserProfileID
				inventory.OfficeId = officeID

				_, err = updateInventoryItem(inventory.Id, inventory)
				if err != nil {
					return nil, err
				}

			}
			_, err := shared.MakeAPIRequest("POST", config.INVENTORY_DISPATCH_ITEMS_ENDOPOINT, itemDispatch, nil)
			if err != nil {
				return nil, err
			}
		}
	}

	return res.Data, nil
}

func updateDispatchItem(id int, item *structs.BasicInventoryDispatchItem) (*structs.BasicInventoryDispatchItem, error) {
	dispatch := dto.GetBasicInventoryDispatch{}

	_, err := shared.MakeAPIRequest("PUT", config.INVENTORY_DISPATCH_ENDOPOINT+"/"+strconv.Itoa(id), item, &dispatch)
	if err != nil {
		return nil, err
	}

	if item.InventoryId != nil {
		filter := dto.DispatchInventoryItemFilter{
			DispatchID: &item.Id,
		}

		dispatchItems, _ := getMyInventoryDispatchesItems(&filter)

		for _, dispatch := range dispatchItems {
			_, err := shared.MakeAPIRequest("DELETE", config.INVENTORY_DISPATCH_ITEMS_ENDOPOINT+"/"+strconv.Itoa(dispatch.Id), nil, nil)

			if err != nil {
				return nil, err
			}
		}

		for _, inventoryID := range item.InventoryId {
			itemDispatch := structs.BasicInventoryDispatchItemsItem{
				InventoryId: inventoryID,
				DispatchId:  dispatch.Data.Id,
			}

			_, err := shared.MakeAPIRequest("POST", config.INVENTORY_DISPATCH_ITEMS_ENDOPOINT, itemDispatch, nil)
			if err != nil {
				return nil, err
			}
		}
	}

	return dispatch.Data, nil
}

func deleteInventoryDispatch(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.INVENTORY_DISPATCH_ENDOPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getDispatchItemByID(id int) (*structs.BasicInventoryDispatchItem, error) {
	res := dto.GetBasicInventoryDispatch{}

	_, err := shared.MakeAPIRequest("GET", config.INVENTORY_DISPATCH_ENDOPOINT+"/"+strconv.Itoa(id), nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
