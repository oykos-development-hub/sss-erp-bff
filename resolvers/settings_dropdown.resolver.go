package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

var SettingsDropdownResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	entity := params.Args["entity"]
	search := params.Args["search"]
	SettingsDropdownType := &structs.SettingsDropdown{}
	var items []interface{}
	var status = "error"

	if entity.(string) == "inventory_class_type" {
		entity = "settings_dropdown_options"
	}

	if shared.IsString(entity) && len(entity.(string)) > 0 {
		SettingsDropdownData, SettingsDropdownDataErr := shared.ReadJson("http://localhost:8080/mocked-data/"+entity.(string)+".json", SettingsDropdownType)

		if SettingsDropdownDataErr != nil {
			fmt.Printf("Fetching "+entity.(string)+" failed because of this error - %s.\n", SettingsDropdownDataErr)
		}

		if id != nil && shared.IsInteger(id) && id != 0 {
			SettingsDropdownData = shared.FindByProperty(SettingsDropdownData, "Id", id)
		}

		if search != nil && shared.IsString(search) {
			SettingsDropdownData = shared.FindByProperty(SettingsDropdownData, "Title", search, true)
		}

		status = "success"
		items = SettingsDropdownData
	}

	return map[string]interface{}{
		"status":  status,
		"message": "Here's the list you asked for!",
		"items":   items,
	}, nil
}

var SettingsDropdownInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.SettingsDropdown
	var item interface{}
	var status = "error"
	entity := params.Args["entity"]
	dataBytes, _ := json.Marshal(params.Args["data"])
	SettingsDropdownType := &structs.SettingsDropdown{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	if entity.(string) == "inventory_class_type" {
		entity = "settings_dropdown_options"
	}

	if shared.IsString(entity) && len(entity.(string)) > 0 {
		SettingsDropdownData, SettingsDropdownDataErr := shared.ReadJson("http://localhost:8080/mocked-data/"+entity.(string)+".json", SettingsDropdownType)

		if SettingsDropdownDataErr != nil {
			fmt.Printf("Fetching "+entity.(string)+" failed because of this error - %s.\n", SettingsDropdownDataErr)
		}

		if shared.IsInteger(itemId) && itemId != 0 {
			SettingsDropdownData = shared.FilterByProperty(SettingsDropdownData, "Id", itemId)
		} else {
			data.Id = shared.GetRandomNumber()
		}

		var updatedData = append(SettingsDropdownData, data)

		_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/"+entity.(string)+".json"), updatedData)

		status = "success"
		item = data
	}

	return map[string]interface{}{
		"status":  status,
		"message": "You updated this item!",
		"item":    item,
	}, nil
}

var SettingsDropdownDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var status = "error"
	itemId := params.Args["id"]
	entity := params.Args["entity"]
	SettingsDropdownType := &structs.SettingsDropdown{}

	if entity.(string) == "inventory_class_type" {
		entity = "settings_dropdown_options"
	}

	if shared.IsString(entity) && len(entity.(string)) > 0 {
		SettingsDropdownData, SettingsDropdownDataErr := shared.ReadJson("http://localhost:8080/mocked-data/"+entity.(string)+".json", SettingsDropdownType)

		if SettingsDropdownDataErr != nil {
			fmt.Printf("Fetching "+entity.(string)+" failed because of this error - %s.\n", SettingsDropdownDataErr)
		}

		if shared.IsInteger(itemId) && itemId != 0 {
			SettingsDropdownData = shared.FilterByProperty(SettingsDropdownData, "Id", itemId)
		}

		_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/"+entity.(string)+".json"), SettingsDropdownData)

		status = "success"
	}

	return map[string]interface{}{
		"status":  status,
		"message": "You deleted this item!",
	}, nil
}
