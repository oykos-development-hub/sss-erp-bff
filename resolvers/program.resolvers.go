package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func ProgramItemProperties(basicInventoryItems []interface{}, id int, typeProgram string) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}

		// Filtering by program
		if len(typeProgram) > 0 && typeProgram == "program" && mergedItem["parent_id"].(int) > 0 {
			continue
		}

		// Filtering by activities
		if len(typeProgram) > 0 && typeProgram == "activities" && mergedItem["organization_unit_id"].(int) == 0 {
			continue
		}

		// Filtering by subprogram
		if len(typeProgram) > 0 && typeProgram == "subprogram" && mergedItem["organization_unit_id"].(int) > 0 || mergedItem["parent_id"].(int) == 0 {
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

var ProgramOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	var typeProgram string
	var search string

	if params.Args["id"].(int) > 0 {
		id = params.Args["id"].(int)
	}

	if params.Args["type"] == nil {
		typeProgram = ""
	} else {
		typeProgram = params.Args["type"].(string)
	}

	if params.Args["search"] == nil {
		search = ""
	} else {
		search = params.Args["search"].(string)
	}

	page := params.Args["page"]
	size := params.Args["size"]

	ProgramType := &structs.ProgramItem{}
	ProgramData, err := shared.ReadJson(shared.GetDataRoot()+"/program.json", ProgramType)

	if err != nil {
		fmt.Printf("Fetching Program failed because of this error - %s.\n", err)
	}
	// Filtering by search
	if search != "" && shared.IsString(search) {
		ProgramData = shared.FindByProperty(ProgramData, "Title", search, true)
	}
	// Populate data for each Basic Inventory Real Estates
	items = ProgramItemProperties(ProgramData, id, typeProgram)

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
	var populatedData = ProgramItemProperties(sliceData, itemId, "")

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
		ProgramData = DeleteProgramChildren(itemId.(int), ProgramData)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/program.json"), ProgramData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

func DeleteProgramChildren(itemId int, data []interface{}) []interface{} {
	toDelete := map[int]bool{itemId: true}
	prevLen := 0

	for len(toDelete) != prevLen {
		prevLen = len(toDelete)
		for _, item := range data {
			if item, ok := item.(*structs.ProgramItem); ok {
				id := item.Id
				parentId := item.ParentId
				if _, exists := toDelete[parentId]; exists {
					toDelete[id] = true
				}
			}
		}
	}

	var result []interface{}
	for _, item := range data {
		if item, ok := item.(*structs.ProgramItem); ok {
			id := item.Id
			if _, exists := toDelete[id]; !exists {
				result = append(result, item)
			}
		}
	}

	return result
}
