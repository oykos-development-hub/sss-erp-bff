package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func ActivitiesItemProperties(basicInventoryItems []interface{}, id int, organizationUnitID int) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if id != 0 && id != mergedItem["id"] {
			continue
		}

		// Filtering by ID
		if organizationUnitID != 0 && organizationUnitID != mergedItem["organization_unit_id"] {
			continue
		}

		if shared.IsInteger(mergedItem["subroutine_id"]) && mergedItem["subroutine_id"].(int) > 0 {
			var relatedActivity = shared.FetchByProperty(
				"program",
				"ID",
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
				"ID",
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

func (r *Resolver) ActivitiesOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	var organizationUnitID int
	if params.Args["id"].(int) > 0 {
		id = params.Args["id"].(int)
	}

	if params.Args["organization_unit_id"] == nil {
		organizationUnitID = 0
	} else {
		organizationUnitID = params.Args["organization_unit_id"].(int)
	}

	page := params.Args["page"]
	size := params.Args["size"]

	ActivityType := &structs.ActivitiesItem{}
	ActivitiesData, err := shared.ReadJSON(shared.GetDataRoot()+"/activities.json", ActivityType)

	if err != nil {
		fmt.Printf("Fetching Activities failed because of this error - %s.\n", err)
	}

	if params.Args["search"] != nil {
		ActivitiesData = shared.FindByProperty(ActivitiesData, "Title", params.Args["search"].(string), true)
	}
	// Populate data for each Basic Inventory Real Estates
	items = ActivitiesItemProperties(ActivitiesData, id, organizationUnitID)

	total = len(items)

	// Filtering by Pagination params
	if page != 0 && size != 0 {
		items = shared.Pagination(items, page.(int), size.(int))
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"total":   total,
		"items":   items,
	}, nil
}

func (r *Resolver) ActivityInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.ActivitiesItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	ActivityItemType := &structs.ActivitiesItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	ActivityData, err := shared.ReadJSON(shared.GetDataRoot()+"/activities.json", ActivityItemType)

	if err != nil {
		fmt.Printf("Fetching Activities failed because of this error - %s.\n", err)
	}

	if itemID != 0 {
		ActivityData = shared.FilterByProperty(ActivityData, "ID", itemID)
	} else {
		data.ID = shared.GetRandomNumber()
	}

	var updatedData = append(ActivityData, data)

	_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/activities.json"), updatedData)

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = ActivitiesItemProperties(sliceData, itemID, 0)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   populatedData[0],
	}, nil
}

func (r *Resolver) ActivityDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemID := params.Args["id"]
	ActivityItemType := &structs.ActivitiesItem{}
	ActivitiesData, err := shared.ReadJSON(shared.GetDataRoot()+"/activities.json", ActivityItemType)

	if err != nil {
		fmt.Printf("Fetching Activities failed because of this error - %s.\n", err)
	}

	if itemID != 0 {
		ActivitiesData = shared.FilterByProperty(ActivitiesData, "ID", itemID)
	}

	_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/activities.json"), ActivitiesData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
