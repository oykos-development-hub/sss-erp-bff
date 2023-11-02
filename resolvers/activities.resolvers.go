package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func ActivitiesItemProperties(basicInventoryItems []interface{}, id int, organizationUnitId int) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}

		// Filtering by ID
		if shared.IsInteger(organizationUnitId) && organizationUnitId != 0 && organizationUnitId != mergedItem["organization_unit_id"] {
			continue
		}

		if shared.IsInteger(mergedItem["subroutine_id"]) && mergedItem["subroutine_id"].(int) > 0 {
			var relatedActivity = shared.FetchByProperty(
				"program",
				"Id",
				mergedItem["subroutine_id"],
			)
			if len(relatedActivity) > 0 {
				var relatedActivity = shared.WriteStructToInterface(relatedActivity[0])

				mergedItem["subroutine"] = map[string]interface{}{
					"title": relatedActivity["title"],
					"id":    relatedActivity["id"],
				}
			}
		}

		if shared.IsInteger(mergedItem["organization_unit_id"]) && mergedItem["organization_unit_id"].(int) > 0 {
			var relatedOfficesOrganizationUnit = shared.FetchByProperty(
				"organization_unit",
				"Id",
				mergedItem["organization_unit_id"],
			)
			if len(relatedOfficesOrganizationUnit) > 0 {
				var relatedOrganizationUnit = shared.WriteStructToInterface(relatedOfficesOrganizationUnit[0])

				mergedItem["organization_unit"] = map[string]interface{}{
					"title": relatedOrganizationUnit["title"],
					"id":    relatedOrganizationUnit["id"],
				}
			}
		}

		items = append(items, mergedItem)
	}

	return items
}

var ActivitiesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	var organizationUnitId int
	if params.Args["id"].(int) > 0 {
		id = params.Args["id"].(int)
	}

	if params.Args["organization_unit_id"] == nil {
		organizationUnitId = 0
	} else {
		organizationUnitId = params.Args["organization_unit_id"].(int)
	}

	page := params.Args["page"]
	size := params.Args["size"]

	ActivityType := &structs.ActivitiesItem{}
	ActivitiesData, err := shared.ReadJson(shared.GetDataRoot()+"/activities.json", ActivityType)

	if err != nil {
		fmt.Printf("Fetching Activities failed because of this error - %s.\n", err)
	}

	if params.Args["search"] != nil {
		ActivitiesData = shared.FindByProperty(ActivitiesData, "Title", params.Args["search"].(string), true)
	}
	// Populate data for each Basic Inventory Real Estates
	items = ActivitiesItemProperties(ActivitiesData, id, organizationUnitId)

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

var ActivityInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.ActivitiesItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	ActivityItemType := &structs.ActivitiesItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	ActivityData, err := shared.ReadJson(shared.GetDataRoot()+"/activities.json", ActivityItemType)

	if err != nil {
		fmt.Printf("Fetching Activities failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		ActivityData = shared.FilterByProperty(ActivityData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(ActivityData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/activities.json"), updatedData)

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = ActivitiesItemProperties(sliceData, itemId, 0)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   populatedData[0],
	}, nil
}

var ActivityDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	ActivityItemType := &structs.ActivitiesItem{}
	ActivitiesData, err := shared.ReadJson(shared.GetDataRoot()+"/activities.json", ActivityItemType)

	if err != nil {
		fmt.Printf("Fetching Activities failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		ActivitiesData = shared.FilterByProperty(ActivitiesData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/activities.json"), ActivitiesData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
