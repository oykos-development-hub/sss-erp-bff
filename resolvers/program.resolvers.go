package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func ProgramItemProperties(basicInventoryItems []interface{}, id int, program ...interface{}) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}

		// Filtering by program
		if len(program) > 0 && program[0] != nil && program[0] == true && mergedItem["parent_id"].(int) > 0 {
			continue
		}

		// Filtering by subprogram
		if len(program) > 0 && program[0] != nil && program[0] == false && mergedItem["parent_id"].(int) == 0 {
			continue
		}

		if shared.IsInteger(mergedItem["parent_id"]) && mergedItem["parent_id"].(int) > 0 {
			var relatedProgram = shared.FetchByProperty(
				"program",
				"Id",
				mergedItem["parent_id"],
			)
			if len(relatedProgram) > 0 {
				var relatedProgram = shared.WriteStructToInterface(relatedProgram[0])

				mergedItem["parent"] = map[string]interface{}{
					"title": relatedProgram["title"],
					"id":    relatedProgram["id"],
				}
			}
		}

		items = append(items, mergedItem)
	}

	return items
}

var ProgramOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	var program = params.Args["program"]
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	page := params.Args["page"]
	size := params.Args["size"]

	ProgramType := &structs.ProgramItem{}
	ProgramData, err := shared.ReadJson(shared.GetDataRoot()+"/program.json", ProgramType)

	if err != nil {
		fmt.Printf("Fetching Program failed because of this error - %s.\n", err)
	}

	if params.Args["search"] != nil {
		ProgramData = shared.FindByProperty(ProgramData, "Title", params.Args["search"].(string), true)
	}
	// Populate data for each Basic Inventory Real Estates
	items = ProgramItemProperties(ProgramData, id, program)

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

var ProgramInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.ProgramItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	ProgramItemType := &structs.ProgramItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	ProgramData, err := shared.ReadJson(shared.GetDataRoot()+"/program.json", ProgramItemType)

	if err != nil {
		fmt.Printf("Fetching Program failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		ProgramData = shared.FilterByProperty(ProgramData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(ProgramData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/program.json"), updatedData)

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = ProgramItemProperties(sliceData, itemId)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   populatedData[0],
	}, nil
}

var ProgramDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	ProgramItemType := &structs.ProgramItem{}
	ProgramData, err := shared.ReadJson(shared.GetDataRoot()+"/program.json", ProgramItemType)

	if err != nil {
		fmt.Printf("Fetching Program failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		ProgramData = shared.FilterByProperty(ProgramData, "Id", itemId)
		ProgramData = shared.FilterByProperty(ProgramData, "ParentId", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/program.json"), ProgramData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
