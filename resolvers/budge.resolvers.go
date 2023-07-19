package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func BudgeItemProperties(basicInventoryItems []interface{}, id int) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}

		if shared.IsInteger(mergedItem["activity_id"]) && mergedItem["activity_id"].(int) > 0 {
			var relatedBudge = shared.FetchByProperty(
				"activities",
				"Id",
				mergedItem["activity_id"],
			)
			if len(relatedBudge) > 0 {
				var relatedBudge = shared.WriteStructToInterface(relatedBudge[0])

				mergedItem["activity"] = map[string]interface{}{
					"title": relatedBudge["title"],
					"id":    relatedBudge["id"],
				}
			}
		}

		items = append(items, mergedItem)
	}

	return items
}

var BudgeOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
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

	BudgeType := &structs.BudgeItem{}
	BudgeData, err := shared.ReadJson(shared.GetDataRoot()+"/activities.json", BudgeType)

	if err != nil {
		fmt.Printf("Fetching Budge failed because of this error - %s.\n", err)
	}

	// Populate data for each Basic Inventory Real Estates
	items = BudgeItemProperties(BudgeData, id)

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

var BudgeInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgeItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgeItemType := &structs.BudgeItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	BudgeData, err := shared.ReadJson(shared.GetDataRoot()+"/activities.json", BudgeItemType)

	if err != nil {
		fmt.Printf("Fetching Budge failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		BudgeData = shared.FilterByProperty(BudgeData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(BudgeData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/activities.json"), updatedData)

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgeItemProperties(sliceData, itemId)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var BudgeDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	BudgeItemType := &structs.BudgeItem{}
	BudgeData, err := shared.ReadJson(shared.GetDataRoot()+"/activities.json", BudgeItemType)

	if err != nil {
		fmt.Printf("Fetching Budge failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		BudgeData = shared.FilterByProperty(BudgeData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/activities.json"), BudgeData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
