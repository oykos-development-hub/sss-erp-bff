package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func PopulateBasicInventoryDispatchItemProperties(basicInventoryDispatchItems []interface{}, id int, typeDispatch string, sourceUserProfileId int, status ...interface{}) []interface{} {
	var items []interface{}
	for _, item := range basicInventoryDispatchItems {
		var mergedItem = shared.WriteStructToInterface(item)
		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
			// if return for revers
		} else if mergedItem["target_organization_unit_id"] == mergedItem["source_organization_unit_id"] {
			continue
		}

		// Filtering by status
		if shared.IsString(typeDispatch) && len(typeDispatch) > 0 && typeDispatch != mergedItem["type"] {
			continue
		}

		// Filtering by sourceUserProfileId
		if shared.IsInteger(sourceUserProfileId) && sourceUserProfileId != 0 && id != mergedItem["source_user_profile_id"] {
			continue
		}

		// Filtering by sourceUserProfileId
		if len(status) > 0 && status[0] != nil && status[0] != mergedItem["is_accepted"] && mergedItem["type"] == "revers" {
			continue
		}

		if shared.IsInteger(mergedItem["target_user_profile_id"]) && mergedItem["target_user_profile_id"].(int) > 0 {
			var relatedInventoryUserProfile = shared.FetchByProperty(
				"user_profile",
				"Id",
				mergedItem["target_user_profile_id"],
			)

			// Populating User Profile data
			if len(relatedInventoryUserProfile) > 0 {
				var relatedUserProfile = shared.WriteStructToInterface(relatedInventoryUserProfile[0])

				mergedItem["target_user_profile"] = map[string]interface{}{
					"title": relatedUserProfile["first_name"].(string) + " " + relatedUserProfile["last_name"].(string),
					"id":    relatedUserProfile["id"],
				}
			}
		}

		if shared.IsInteger(mergedItem["source_user_profile_id"]) && mergedItem["source_user_profile_id"].(int) > 0 {
			var relatedInventoryUserProfile = shared.FetchByProperty(
				"user_profile",
				"Id",
				mergedItem["source_user_profile_id"],
			)

			// Populating User Profile data
			if len(relatedInventoryUserProfile) > 0 {
				var relatedUserProfile = shared.WriteStructToInterface(relatedInventoryUserProfile[0])

				mergedItem["source_user_profile"] = map[string]interface{}{
					"title": relatedUserProfile["first_name"].(string) + " " + relatedUserProfile["last_name"].(string),
					"id":    relatedUserProfile["id"],
				}
			}
		}

		if mergedItem["source_organization_unit_id"].(int) > 0 {
			var relatedOfficesOrganizationUnit = shared.FetchByProperty(
				"organization_unit",
				"Id",
				mergedItem["source_organization_unit_id"],
			)
			if len(relatedOfficesOrganizationUnit) > 0 {
				var relatedOrganizationUnit = shared.WriteStructToInterface(relatedOfficesOrganizationUnit[0])

				mergedItem["source_organization_unit"] = map[string]interface{}{
					"title": relatedOrganizationUnit["title"],
					"id":    relatedOrganizationUnit["id"],
				}
			}
		}

		if mergedItem["target_organization_unit_id"].(int) > 0 {
			var relatedOfficesOrganizationUnit = shared.FetchByProperty(
				"organization_unit",
				"Id",
				mergedItem["target_organization_unit_id"],
			)
			if len(relatedOfficesOrganizationUnit) > 0 {
				var relatedOrganizationUnit = shared.WriteStructToInterface(relatedOfficesOrganizationUnit[0])

				mergedItem["target_organization_unit"] = map[string]interface{}{
					"title": relatedOrganizationUnit["title"],
					"id":    relatedOrganizationUnit["id"],
				}
			}
		}

		if shared.IsInteger(mergedItem["office_id"]) && mergedItem["office_id"].(int) > 0 {
			var relatedInventoryOffice = shared.FetchByProperty(
				"offices_of_organization_units",
				"Id",
				mergedItem["office_id"],
			)
			if len(relatedInventoryOffice) > 0 {
				var relatedOffice = shared.WriteStructToInterface(relatedInventoryOffice[0])

				mergedItem["office"] = map[string]interface{}{
					"title": relatedOffice["title"],
					"id":    relatedOffice["id"],
				}
			}
		}

		if mergedItem["target_organization_unit_id"] != mergedItem["source_organization_unit_id"] {
			if mergedItem["type"] == "revers" {
				mergedItem["type"] = "Revers"
			}

			if mergedItem["type"] == "return-revers" {
				mergedItem["type"] = "Return revers"
			}
		}

		if mergedItem["target_organization_unit_id"] == mergedItem["source_organization_unit_id"] && shared.IsInteger(mergedItem["office_id"]) && mergedItem["office_id"].(int) > 0 && mergedItem["type"] == "allocation" {
			mergedItem["type"] = "Allocation"
		} else if mergedItem["target_user_profile_id"].(int) == 0 && mergedItem["type"] == "return" {
			mergedItem["type"] = "Return"
		}

		if shared.IsInteger(id) && id != 0 && id == mergedItem["id"] {
			BasicInventoryDispatchItemsType := &structs.BasicInventoryDispatchItemsItem{}
			basicInventoryDispatchItemsData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_dispatch_items.json", BasicInventoryDispatchItemsType)

			if err != nil {
				fmt.Printf("Fetching Basic Inventory Dispatch Items failed because of this error - %s.\n", err)
			}

			basicInventoryDispatchItemsData = shared.FindByProperty(basicInventoryDispatchItemsData, "DispatchId", id)

			BasicInventoryType := &structs.BasicInventoryInsertItem{}
			basicInventoryData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_items.json", BasicInventoryType)

			if err != nil {
				fmt.Printf("Fetching Basic Inventory Details failed because of this error - %s.\n", err)
			}
			mergedItem["inventory"] = []interface{}{}
			if len(basicInventoryDispatchItemsData) > 0 {
				for _, item := range basicInventoryDispatchItemsData {
					if m, ok := item.(*structs.BasicInventoryDispatchItemsItem); ok {
						dataInventory := shared.FindByProperty(basicInventoryData, "Id", m.InventoryId)

						if len(dataInventory) > 0 {
							inventory, ok := mergedItem["inventory"].([]interface{})
							if ok {
								mergedItem["inventory"] = append(inventory, dataInventory[0])
							}
						}
					}
				}
			}
		}

		items = append(items, mergedItem)
	}
	return items
}

var BasicInventoryDispatchOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var typeDispatch string
	var sourceUserProfileId int

	page := params.Args["page"]
	size := params.Args["size"]
	id := params.Args["id"]

	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	if params.Args["type"] == nil {
		typeDispatch = ""
	} else {
		typeDispatch = params.Args["type"].(string)
	}

	if params.Args["source_user_profile_id"] == nil {
		sourceUserProfileId = 0
	} else {
		sourceUserProfileId = params.Args["source_user_profile_id"].(int)
	}

	BasicInventoryDispatchType := &structs.BasicInventoryDispatchItem{}
	basicInventoryDispatchData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_dispatch.json", BasicInventoryDispatchType)

	if err != nil {
		fmt.Printf("Fetching Job Tenders failed because of this error - %s.\n", err)
		return nil, err
	}

	// Populate data for each Basic Inventory Depreciation Types
	items = PopulateBasicInventoryDispatchItemProperties(basicInventoryDispatchData, id.(int), typeDispatch, sourceUserProfileId, params.Args["status"])

	total = len(items)

	// Filtering by Pagination params
	if shared.IsInteger(page) && page != 0 && shared.IsInteger(size) && size != 0 {
		items = shared.Pagination(items, page.(int), size.(int))
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"total":   total,
		"items":   items,
	}, nil

}

var BasicInventoryDispatchInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BasicInventoryDispatchItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BasicInventoryDispatchType := &structs.BasicInventoryDispatchItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	targetOrganizationUnitId := data.TargetOrganizationUnitId

	basicInventoryDispatchData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_dispatch.json", BasicInventoryDispatchType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Dispatch failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		basicInventoryDispatchData = shared.FilterByProperty(basicInventoryDispatchData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
		BasicInventoryDispatchItemsType := &structs.BasicInventoryDispatchItemsItem{}
		basicInventoryDispatchItemsData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_dispatch_items.json", BasicInventoryDispatchItemsType)

		if err != nil {
			fmt.Printf("Fetching Basic Inventory Dispatch Items failed because of this error - %s.\n", err)
		}

		BasicInventoryType := &structs.BasicInventoryInsertItem{}
		basicInventoryData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_items.json", BasicInventoryType)

		if err != nil {
			fmt.Printf("Fetching Basic Inventory Details failed because of this error - %s.\n", err)
		}

		var updatedDataItems = append(basicInventoryDispatchItemsData)
		if len(data.InventoryId) > 0 {
			for _, item := range data.InventoryId {
				basicInventoryDispatchItem := structs.BasicInventoryDispatchItemsItem{
					Id:          shared.GetRandomNumber(),
					InventoryId: item,
					DispatchId:  data.Id,
				}
				updatedDataItems = append(updatedDataItems, basicInventoryDispatchItem)
				updateItemInventory := shared.FindByProperty(basicInventoryData, "Id", item)

				if m, ok := updateItemInventory[0].(*structs.BasicInventoryInsertItem); ok {
					m.TargetOrganizationUnitId = targetOrganizationUnitId
				}

				basicInventoryData = shared.FilterByProperty(basicInventoryData, "Id", item)
				basicInventoryData = append(basicInventoryData, updateItemInventory[0])
			}
		}

		_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_dispatch_items.json"), updatedDataItems)
		_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_items.json"), basicInventoryData)
	}

	data.InventoryId = nil
	var updatedData = append(basicInventoryDispatchData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_dispatch.json"), updatedData)

	if data.SourceOrganizationUnitId == data.TargetOrganizationUnitId {
		var targetInventoryUserId int
		var targetInventoryId = data.InventoryId[0]
		if data.OfficeId > 0 {
			targetInventoryUserId = data.TargetOrganizationUnitId
		} else {
			targetInventoryUserId = 0
		}
		BasicInventoryType := &structs.BasicInventoryInsertItem{}
		basicInventoryData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_items.json", BasicInventoryType)
		if err != nil {
			fmt.Printf("Fetching Basic Inventory Details failed because of this error - %s.\n", err)
		}

		targetInventory := shared.FindByProperty(basicInventoryData, "Id", targetInventoryId)
		for _, item := range targetInventory {
			if m, ok := item.(*structs.BasicInventoryInsertItem); ok {
				m.TargetUserProfileId = targetInventoryUserId
			}
		}

		basicInventoryData = shared.FilterByProperty(basicInventoryData, "Id", itemId)

		var updatedData = append(basicInventoryData, targetInventory)

		_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_items.json"), updatedData)
	}

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = PopulateBasicInventoryDispatchItemProperties(sliceData, itemId, "", 0)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var BasicInventoryDispatchDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	BasicInventoryDispatchType := &structs.BasicInventoryDispatchItem{}
	basicInventoryDispatchData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_dispatch.json", BasicInventoryDispatchType)

	if err != nil {
		fmt.Printf("Fetching Inventory Dispatch Delete failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		basicInventoryDispatchData = shared.FilterByProperty(basicInventoryDispatchData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_dispatch.json"), basicInventoryDispatchData)

	BasicInventoryDispatchItemsType := &structs.BasicInventoryDispatchItemsItem{}
	basicInventoryDispatchItemsData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_dispatch_items.json", BasicInventoryDispatchItemsType)

	removeItemsDispatchItemsData := shared.FindByProperty(basicInventoryDispatchItemsData, "DispatchId", itemId)

	if len(removeItemsDispatchItemsData) > 0 {
		BasicInventoryType := &structs.BasicInventoryInsertItem{}
		basicInventoryData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_items.json", BasicInventoryType)

		if err != nil {
			fmt.Printf("Fetching Basic Inventory Details failed because of this error - %s.\n", err)
		}

		for _, item := range removeItemsDispatchItemsData {
			if i, ok := item.(*structs.BasicInventoryDispatchItemsItem); ok {
				updateItemInventory := shared.FindByProperty(basicInventoryData, "Id", i.InventoryId)

				if m, ok := updateItemInventory[0].(*structs.BasicInventoryInsertItem); ok {
					m.TargetOrganizationUnitId = 0
				}

				basicInventoryData = shared.FilterByProperty(basicInventoryData, "Id", i.InventoryId)
				basicInventoryData = append(basicInventoryData, updateItemInventory[0])
			}

		}

		_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_items.json"), basicInventoryData)
	}

	if err != nil {
		fmt.Printf("Fetching Inventory Dispatch Delete failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		basicInventoryDispatchItemsData = shared.FilterByProperty(basicInventoryDispatchItemsData, "DispatchId", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_dispatch_items.json"), basicInventoryDispatchItemsData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

var BasicInventoryDispatchAcceptResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["dispatch_id"].(int)
	targetUserId := params.Args["target_user_id"].(int)
	var targetOrganizationUnitId int

	BasicInventoryDispatchType := &structs.BasicInventoryDispatchItem{}

	if !shared.IsInteger(itemId) && itemId == 0 {
		fmt.Printf("Dispatch id is empty")
		return nil, nil
	}
	if !shared.IsInteger(targetUserId) && targetUserId == 0 {
		fmt.Printf("Target user id is empty")
		return nil, nil
	}

	basicInventoryDispatchData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_dispatch.json", BasicInventoryDispatchType)

	if err != nil {
		fmt.Printf("Fetching Inventory Dispatch failed because of this error - %s.\n", err)
	}

	item := shared.FindByProperty(basicInventoryDispatchData, "Id", itemId, true)

	for _, i := range item {
		if m, ok := i.(*structs.BasicInventoryDispatchItem); ok {
			m.IsAccepted = true
			targetOrganizationUnitId = m.TargetOrganizationUnitId
			if shared.IsInteger(targetUserId) && targetUserId != 0 {
				m.TargetUserProfileId = targetUserId
			}
		}
	}

	basicInventoryDispatchData = shared.FilterByProperty(basicInventoryDispatchData, "Id", itemId)
	updateData := append(basicInventoryDispatchData, item[0])

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_dispatch.json"), updateData)

	BasicInventoryDispatchItemsType := &structs.BasicInventoryDispatchItemsItem{}
	basicInventoryDispatchItemsData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_dispatch_items.json", BasicInventoryDispatchItemsType)
	if err != nil {
		fmt.Printf("Fetching Basic Inventory Dispatch Items failed because of this error - %s.\n", err)
	}
	basicInventoryDispatchItemsData = shared.FindByProperty(basicInventoryDispatchItemsData, "DispatchId", itemId)

	BasicInventoryType := &structs.BasicInventoryInsertItem{}

	basicInventoryData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_items.json", BasicInventoryType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Details failed because of this error - %s.\n", err)
	}

	for _, item := range basicInventoryData {
		if m, ok := item.(*structs.BasicInventoryInsertItem); ok {
			for _, i := range basicInventoryDispatchItemsData {
				if n, ok := i.(*structs.BasicInventoryDispatchItemsItem); ok {
					if m.Id == n.InventoryId {
						m.TargetOrganizationUnitId = targetOrganizationUnitId
						break
					}
				}
			}
		}
	}
	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_items.json"), basicInventoryData)
	return map[string]interface{}{
		"status":  "success",
		"message": "You Accept this item!",
	}, nil
}
