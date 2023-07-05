package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func PopulateBasicInventoryDepreciationTypesItemProperties(basicInventoryItems []interface{}, id int) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}

		items = append(items, mergedItem)
	}

	return items
}

var BasicInventoryDepreciationTypesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	page := params.Args["page"]
	size := params.Args["size"]

	BasicInventoryDepreciationTypesType := &structs.BasicInventoryDepreciationTypesItem{}
	BasicInventoryDepreciationTypesData, err := shared.ReadJson(shared.GetDataRoot()+"/basic_inventory_depreciation_types.json", BasicInventoryDepreciationTypesType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
	}

	// Populate data for each Basic Inventory Depreciation Types
	items = PopulateBasicInventoryDepreciationTypesItemProperties(BasicInventoryDepreciationTypesData, id)

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

var BasicInventoryDepreciationTypesInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BasicInventoryDepreciationTypesItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BasicInventoryDepreciationTypesType := &structs.BasicInventoryDepreciationTypesItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	BasicInventoryDepreciationTypesData, err := shared.ReadJson(shared.GetDataRoot()+"/basic_inventory_depreciation_types.json", BasicInventoryDepreciationTypesType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
	}

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = PopulateBasicInventoryDepreciationTypesItemProperties(sliceData, itemId)

	if shared.IsInteger(itemId) && itemId != 0 {
		BasicInventoryDepreciationTypesData = shared.FilterByProperty(BasicInventoryDepreciationTypesData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(BasicInventoryDepreciationTypesData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_depreciation_types.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var BasicInventoryDepreciationTypesDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	BasicInventoryDepreciationTypesType := &structs.BasicInventoryDepreciationTypesItem{}
	BasicInventoryDepreciationTypesData, err := shared.ReadJson(shared.GetDataRoot()+"/basic_inventory_depreciation_types.json", BasicInventoryDepreciationTypesType)

	if err != nil {
		fmt.Printf("Fetching Inventory Depreciation Types Delete failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		BasicInventoryDepreciationTypesData = shared.FilterByProperty(BasicInventoryDepreciationTypesData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_depreciation_types.json"), BasicInventoryDepreciationTypesData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
