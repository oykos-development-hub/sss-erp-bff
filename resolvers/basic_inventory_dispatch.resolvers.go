package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

var BasicInventoryDispatchOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.InventoryDispatchResponse
	var filter dto.InventoryDispatchFilter

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

		response.Message = "You created this item!"
		items, err = buildInventoryDispatchResponse(itemRes)

		if err != nil {
			return shared.HandleAPIError(err)
		}

	}

	response.Item = items
	return response, nil
}

var BasicInventoryDispatchDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteInventoryDispatch(itemId)
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

	var authToken = params.Context.Value(config.TokenKey).(string)

	loggedInProfile, err := getLoggedInUserProfile(authToken)
	if err != nil {
		return shared.HandleAPIError(err)
	}

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

	_, err = updateDispatchItem(dispatch.Id, dispatch)
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
		/*if err != nil {
			// return nil, err
		}*/
		if user != nil {
			sourceUserDropdown = dto.DropdownSimple{Id: user.Id, Title: user.FirstName + " " + user.LastName}
		}
	}

	targetUserDropdown := dto.DropdownSimple{}
	if item.TargetUserProfileId != 0 {
		user, _ := getUserProfileById(item.TargetUserProfileId)
		/*if err != nil {
			 return nil, err
		}*/
		if user != nil {
			targetUserDropdown = dto.DropdownSimple{Id: user.Id, Title: user.FirstName + " " + user.LastName}
		}
	}

	sourceOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.SourceOrganizationUnitId != 0 {
		sourceOrganizationUnit, _ := getOrganizationUnitById(item.SourceOrganizationUnitId)
		/*if err != nil {
			// return nil, err
		}*/
		if sourceOrganizationUnit != nil {
			sourceOrganizationUnitDropdown = dto.DropdownSimple{Id: sourceOrganizationUnit.Id, Title: sourceOrganizationUnit.Title}
		}
	}

	targetOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.TargetOrganizationUnitId != 0 {
		targetOrganizationUnit, _ := getOrganizationUnitById(item.TargetOrganizationUnitId)
		/*if err != nil {
			// return nil, err
		}*/
		if targetOrganizationUnit != nil {
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

			item := dto.BasicInventoryResponseItem{
				Id:              itemInventory.Id,
				Type:            itemInventory.Type,
				InventoryNumber: itemInventory.InventoryNumber,
				Title:           itemInventory.Title,
				GrossPrice:      itemInventory.GrossPrice,
				SerialNumber:    itemInventory.SerialNumber,
			}
			inventoryItems = append(inventoryItems, item)
		}
	}

	res := dto.InventoryDispatchResponse{
		ID:                     item.Id,
		Type:                   item.Type,
		SerialNumber:           item.SerialNumber,
		Office:                 settingDropdownOfficeId,
		SourceUserProfile:      sourceUserDropdown,
		TargetUserProfile:      targetUserDropdown,
		SourceOrganizationUnit: sourceOrganizationUnitDropdown,
		TargetOrganizationUnit: targetOrganizationUnitDropdown,
		InventoryType:          item.InventoryType,
		Inventory:              inventoryItems,
		CreatedAt:              item.CreatedAt,
		UpdatedAt:              item.UpdatedAt,
		DispatchDescription:    item.DispatchDescription,
		FileId:                 item.FileId,
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
