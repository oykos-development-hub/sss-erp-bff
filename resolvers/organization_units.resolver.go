package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func PopulateOrganizationUnitItemProperties(organizationUnits []interface{}) []interface{} {
	var items []interface{}

	for _, item := range organizationUnits {

		var mergedItem = shared.WriteStructToInterface(item)
		// Fetching children Organization Units
		mergedItem["children"] = shared.FetchByProperty("organization_unit", "ParentId", mergedItem["id"])

		items = append(items, mergedItem)
	}

	return items
}

var OrganizationUnitsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	search := params.Args["search"]
	OrganizationUnitType := &structs.OrganizationUnits{}
	organizationUnitsData, organizationUnitDataErr := shared.ReadJson("http://localhost:8080/mocked-data/organization_units.json", OrganizationUnitType)
	var organizationUnitData []interface{}

	if organizationUnitDataErr != nil {
		fmt.Printf("Fetching organization units failed because of this error - %s.\n", organizationUnitDataErr)
	}

	organizationUnitData = organizationUnitsData

	if id != nil && shared.IsInteger(id) && id != 0 {
		organizationUnitData = shared.FindByProperty(organizationUnitData, "Id", id)
	}

	if search != nil && shared.IsString(search) {
		organizationUnitData = shared.FindByProperty(organizationUnitData, "Title", search, true)
	}

	organizationUnitData = PopulateOrganizationUnitItemProperties(organizationUnitData)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"items":   organizationUnitData,
	}, nil
}

var OrganizationUnitInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.OrganizationUnits
	dataBytes, _ := json.Marshal(params.Args["data"])
	OrganizationUnitType := &structs.OrganizationUnits{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	organizationUnitData, organizationUnitDataErr := shared.ReadJson("http://localhost:8080/mocked-data/organization_units.json", OrganizationUnitType)

	if organizationUnitDataErr != nil {
		fmt.Printf("Fetching organization units failed because of this error - %s.\n", organizationUnitDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		organizationUnitData = shared.FilterByProperty(organizationUnitData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(organizationUnitData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/organization_units.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var OrganizationUnitDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	OrganizationUnitType := &structs.OrganizationUnits{}
	organizationUnitData, organizationUnitDataErr := shared.ReadJson("http://localhost:8080/mocked-data/organization_units.json", OrganizationUnitType)

	if organizationUnitDataErr != nil {
		fmt.Printf("Fetching organization units failed because of this error - %s.\n", organizationUnitDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		organizationUnitData = shared.FilterByProperty(organizationUnitData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/organization_units.json"), organizationUnitData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
