package resolvers

import (
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func BudgetItemProperties(basicInventoryItems []interface{}, id int, typeBudget string, year string, status string) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}
		// Filtering by type
		if shared.IsString(typeBudget) && len(typeBudget) > 0 && typeBudget != mergedItem["type"] {
			continue
		}
		// Filtering by year
		if shared.IsString(year) && len(year) > 0 && year != mergedItem["year"] {
			continue
		}
		// Filtering by status
		if shared.IsString(status) && len(status) > 0 && status != mergedItem["status"] {
			continue
		}

		items = append(items, mergedItem)
	}

	return items
}

var BudgetOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	var status string
	var year string
	var typeBudget string

	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	if params.Args["status"] == nil {
		status = ""
	} else {
		status = params.Args["status"].(string)
	}

	if params.Args["year"] == nil {
		year = ""
	} else {
		year = params.Args["year"].(string)
	}

	if params.Args["type_budget"] == nil {
		typeBudget = ""
	} else {
		typeBudget = params.Args["type_budget"].(string)
	}

	page := params.Args["page"]
	size := params.Args["size"]

	BudgetType := &structs.BudgetItem{}
	BudgetData, err := shared.ReadJson(shared.GetDataRoot()+"/budget.json", BudgetType)

	if err != nil {
		fmt.Printf("Fetching Budget failed because of this error - %s.\n", err)
	}

	// Populate data for each Basic Inventory Real Estates
	items = BudgetItemProperties(BudgetData, id, typeBudget, year, status)

	total = len(items)

	// Filtering by Pagination params
	if shared.IsInteger(page) && page != 0 && shared.IsInteger(size) && size != 0 {
		items = shared.Pagination(items, page.(int), size.(int))
	}

	return dto.Response{
		Status:  "success",
		Message: "You fetched items!",
		Items:   items,
		Total:   total,
	}, nil
}

var BudgetInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgetItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgetItemType := &structs.BudgetItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	BudgetData, err := shared.ReadJson(shared.GetDataRoot()+"/budget.json", BudgetItemType)

	if err != nil {
		fmt.Printf("Fetching Budget failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		BudgetData = shared.FilterByProperty(BudgetData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
		data.Status = "kreiran"
	}

	var updatedData = append(BudgetData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/budget.json"), updatedData)

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgetItemProperties(sliceData, itemId, "", "", "")

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   populatedData[0],
	}, nil
}

var BudgetSendResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]

	BudgetItemType := &structs.BudgetItem{}
	BudgetData, err := shared.ReadJson(shared.GetDataRoot()+"/budget.json", BudgetItemType)

	if err != nil {
		fmt.Printf("Fetching Budget failed because of this error - %s.\n", err)
		return nil, err
	}
	budget := shared.FindByProperty(BudgetData, "Id", itemId)
	if len(budget) == 0 {
		fmt.Printf("Fetching Budget failed because of this error - %s.\n", err)
		return nil, err
	}
	BudgetData = shared.FilterByProperty(BudgetData, "Id", itemId)
	newItem := structs.BudgetItem{}

	for _, item := range budget {
		if updateBudget, ok := item.(*structs.BudgetItem); ok {
			newItem.Id = updateBudget.Id
			newItem.Year = updateBudget.Year
			newItem.Source = updateBudget.Source
			newItem.Type = updateBudget.Type
			newItem.Status = "poslat"
		}
	}

	var updatedData = append(BudgetData, newItem)
	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/budget.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You send budget to OJ!",
	}, nil
}

var BudgetDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	BudgetItemType := &structs.BudgetItem{}
	BudgetData, err := shared.ReadJson(shared.GetDataRoot()+"/budget.json", BudgetItemType)

	if err != nil {
		fmt.Printf("Fetching Budget failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		BudgetData = shared.FilterByProperty(BudgetData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/budget.json"), BudgetData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
