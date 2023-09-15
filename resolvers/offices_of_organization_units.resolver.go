package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func PopulateOfficesOfOrganizationUnitItemProperties(basicInventoryItems []interface{}, id int, organizationUnitId int, search string) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}

		// Filtering by organizationUnitId
		if shared.IsInteger(organizationUnitId) && organizationUnitId != 0 && organizationUnitId != mergedItem["organization_unit_id"] {
			continue
		}

		// Filtering by status
		if shared.IsString(search) && len(search) > 0 && !shared.StringContains(mergedItem["title"].(string), search) {
			continue
		}

		if mergedItem["organization_unit_id"].(int) > 0 {
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
		} else {
			continue
		}

		items = append(items, mergedItem)
	}

	return items
}

var OfficesOfOrganizationUnitOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	var organizationUnitId int
	var search string

	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	if params.Args["organization_unit_id"] == nil {
		organizationUnitId = 0
	} else {
		organizationUnitId = params.Args["organization_unit_id"].(int)
	}

	if params.Args["search"] == nil {
		search = ""
	} else {
		search = params.Args["search"].(string)
	}

	OfficesOfOrganizationUnitType := &structs.OfficesOfOrganizationUnitItem{}
	OfficesOfOrganizationUnitData, err := shared.ReadJson(shared.GetDataRoot()+"/offices_of_organization_units.json", OfficesOfOrganizationUnitType)

	if err != nil {
		fmt.Printf("Fetching Job Tenders failed because of this error - %s.\n", err)
	}

	// Populate data for each Basic Inventory Depreciation Types
	items = PopulateOfficesOfOrganizationUnitItemProperties(OfficesOfOrganizationUnitData, id, organizationUnitId, search)

	total = len(items)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"total":   total,
		"items":   items,
	}, nil
}

var OfficesOfOrganizationUnitInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.OfficesOfOrganizationUnitItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	OfficesOfOrganizationUnitType := &structs.OfficesOfOrganizationUnitItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	OfficesOfOrganizationUnitData, err := shared.ReadJson(shared.GetDataRoot()+"/offices_of_organization_units.json", OfficesOfOrganizationUnitType)

	if err != nil {
		fmt.Printf("Fetching Offices Of Organization Unit failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		OfficesOfOrganizationUnitData = shared.FilterByProperty(OfficesOfOrganizationUnitData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(OfficesOfOrganizationUnitData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/offices_of_organization_units.json"), updatedData)

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = PopulateOfficesOfOrganizationUnitItemProperties(sliceData, itemId, 0, "")

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var OfficesOfOrganizationUnitDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	OfficesOfOrganizationUnitType := &structs.OfficesOfOrganizationUnitItem{}
	OfficesOfOrganizationUnitData, err := shared.ReadJson(shared.GetDataRoot()+"/offices_of_organization_units.json", OfficesOfOrganizationUnitType)

	if err != nil {
		fmt.Printf("Fetching Inventory Depreciation Types Delete failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		OfficesOfOrganizationUnitData = shared.FilterByProperty(OfficesOfOrganizationUnitData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/offices_of_organization_units.json"), OfficesOfOrganizationUnitData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
