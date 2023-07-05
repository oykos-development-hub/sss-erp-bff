package resolvers

import (
	"bff/shared"
	"bff/structs"
	"fmt"

	"github.com/graphql-go/graphql"
)

func PopulateBasicInventoryRealEstatesItemProperties(basicInventoryItems []interface{}, id int) []interface{} {
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

var BasicInventoryRealEstatesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
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

	BasicInventoryRealEstatesType := &structs.BasicInventoryRealEstatesItem{}
	BasicInventoryRealEstatesData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_real_estates.json", BasicInventoryRealEstatesType)

	if err != nil {
		fmt.Printf("Fetching Job Tenders failed because of this error - %s.\n", err)
	}

	// Populate data for each Basic Inventory Real Estates
	items = PopulateBasicInventoryRealEstatesItemProperties(BasicInventoryRealEstatesData, id)

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
