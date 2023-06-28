package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

var SuppliersOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var total int
	var search string
	if params.Args["search"] == nil {
		search = ""
	} else {
		search = params.Args["search"].(string)
	}
	var id int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}
	page := params.Args["page"]
	size := params.Args["size"]

	var items []interface{}
	var suppliers = shared.FetchByProperty(
		"suppliers",
		"",
		"",
	)

	if len(suppliers) > 0 {
		for _, supplierItem := range suppliers {
			var supplier = shared.WriteStructToInterface(supplierItem)

			if len(search) > 0 && !shared.StringContains(supplier["title"].(string), search) {
				continue
			}
			if id > 0 && id != supplier["id"].(int) {
				continue
			}

			items = append(items, supplier)
		}
	} else {
		items = suppliers
	}

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

var SuppliersInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.Suppliers
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	var items = shared.FetchByProperty(
		"suppliers",
		"",
		"",
	)

	if shared.IsInteger(itemId) && itemId != 0 {
		items = shared.FilterByProperty(items, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(items, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/suppliers.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   []interface{}{data},
	}, nil
}

var SuppliersDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]

	var items = shared.FetchByProperty(
		"suppliers",
		"",
		"",
	)

	if shared.IsInteger(itemId) && itemId != 0 {
		items = shared.FilterByProperty(items, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/suppliers.json"), items)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
