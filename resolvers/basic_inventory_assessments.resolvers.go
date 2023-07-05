package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func PopulateBasicInventoryAssessmentsItemProperties(basicInventoryAssessmentsItems []interface{}, id int, idInventory ...int) []interface{} {
	var items []interface{}
	for _, item := range basicInventoryAssessmentsItems {
		var mergedItem = shared.WriteStructToInterface(item)
		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}
		if len(idInventory) > 0 && shared.IsInteger(idInventory[0]) && idInventory[0] != 0 && idInventory[0] != mergedItem["inventory_id"] {
			continue
		} else {
			if shared.IsInteger(mergedItem["depreciation_type_id"]) && mergedItem["depreciation_type_id"].(int) > 0 {
				var relatedInventoryDepreciationType = shared.FetchByProperty(
					"basic_inventory_depreciation_types",
					"Id",
					mergedItem["depreciation_type_id"],
				)
				if len(relatedInventoryDepreciationType) > 0 {
					var relatedDepreciationType = shared.WriteStructToInterface(relatedInventoryDepreciationType[0])

					mergedItem["depreciation_type"] = map[string]interface{}{
						"title": relatedDepreciationType["title"],
						"id":    relatedDepreciationType["id"],
					}
				}
			}

			if shared.IsInteger(mergedItem["user_profile_id"]) && mergedItem["user_profile_id"].(int) > 0 {
				var userProfile = shared.FetchByProperty(
					"user_profile",
					"Id",
					mergedItem["user_profile_id"],
				)

				if len(userProfile) > 0 {
					var userProfileInterface = shared.WriteStructToInterface(userProfile[0])

					mergedItem["user_profile"] = map[string]interface{}{
						"title": userProfileInterface["first_name"].(string) + " " + userProfileInterface["last_name"].(string),
						"id":    userProfileInterface["id"],
					}
				}
			}

		}

		items = append(items, mergedItem)
	}
	return items
}

var BasicInventoryAssessmentsInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BasicInventoryAssessmentsTypesItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BasicInventoryAssessmentsType := &structs.BasicInventoryAssessmentsTypesItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	data.Active = true

	basicInventoryAssessmentsData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_assessments.json", BasicInventoryAssessmentsType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Assessments failed because of this error - %s.\n", err)
	}

	for _, item := range basicInventoryAssessmentsData {
		if n, ok := item.(*structs.BasicInventoryAssessmentsTypesItem); ok {
			n.Active = false
		}
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		basicInventoryAssessmentsData = shared.FilterByProperty(basicInventoryAssessmentsData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(basicInventoryAssessmentsData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_assessments.json"), updatedData)

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = PopulateBasicInventoryAssessmentsItemProperties(sliceData, itemId, data.InventoryId)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var BasicInventoryAssessmentDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	BasicInventoryAssessmentsType := &structs.BasicInventoryAssessmentsTypesItem{}
	basicInventoryAssessmentsData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_assessments.json", BasicInventoryAssessmentsType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Assessment Delete failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		basicInventoryAssessmentsData = shared.FilterByProperty(basicInventoryAssessmentsData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_assessments.json"), basicInventoryAssessmentsData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
