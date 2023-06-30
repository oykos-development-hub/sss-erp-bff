package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func PopulateBasicInventoryItemProperties(basicInventoryItems []interface{}, organizationUnitId int, id int, typeParam string, classTypeId int, officeId int, search string, sourceType string, depreciationTypeId int) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}

		// Filtering by classTypeId
		if shared.IsInteger(classTypeId) && classTypeId != 0 && classTypeId != mergedItem["class_type_id"] {
			continue
		}

		// Filtering by officeId
		if shared.IsInteger(officeId) && officeId != 0 && officeId != mergedItem["office_id"] {
			continue
		}

		// Filtering by depreciationTypeId
		if shared.IsInteger(depreciationTypeId) && depreciationTypeId != 0 && depreciationTypeId != mergedItem["depreciation_type_id"] {
			continue
		}

		// Filtering by organizationUnitId
		if shared.IsInteger(organizationUnitId) && organizationUnitId != 0 && organizationUnitId != mergedItem["target_organization_unit_id"] && organizationUnitId != mergedItem["organization_unit_id"] {
			continue
		}

		// Filtering by Type
		if len(typeParam) > 0 && typeParam != mergedItem["type"] {
			continue
		}

		if mergedItem["type"] == "immovable" {
			if mergedItem["organization_unit_id"] == mergedItem["target_organization_unit_id"] || organizationUnitId == mergedItem["organization_unit_id"] {
				mergedItem["source_type"] = "NS1"
			} else {
				mergedItem["source_type"] = "NS2"
			}
			if len(sourceType) > 0 && sourceType != mergedItem["source_type"] {
				continue
			}
		}

		if mergedItem["type"] == "movable" {
			if mergedItem["organization_unit_id"] == mergedItem["target_organization_unit_id"] || organizationUnitId == mergedItem["organization_unit_id"] {
				mergedItem["source_type"] = "PS1"
			} else {
				mergedItem["source_type"] = "PS2"
			}

			if len(sourceType) > 0 && sourceType != mergedItem["source_type"] {
				continue
			}
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
		}

		if mergedItem["target_organization_unit_id"].(int) > 0 {
			var relatedOfficesOrganizationUnit = shared.FetchByProperty(
				"organization_unit",
				"Id",
				mergedItem["target_organization_unit_id"],
			)
			if len(relatedOfficesOrganizationUnit) > 0 {
				var relatedOrganizationUnit = shared.WriteStructToInterface(relatedOfficesOrganizationUnit[0])

				mergedItem["target_organization_unit"] = map[string]interface{}{
					"title": relatedOrganizationUnit["title"],
					"id":    relatedOrganizationUnit["id"],
				}
			}
		}

		if shared.IsInteger(mergedItem["class_type_id"]) && mergedItem["class_type_id"].(int) > 0 {
			var relatedInventoryClassType = shared.FetchByProperty(
				"inventory_class_type",
				"Id",
				mergedItem["class_type_id"],
			)

			// Populating User Profile data
			if len(relatedInventoryClassType) > 0 {
				var relatedClassType = shared.WriteStructToInterface(relatedInventoryClassType[0])

				mergedItem["class_type"] = map[string]interface{}{
					"title": relatedClassType["title"],
					"id":    relatedClassType["id"],
				}
			}
		}

		if shared.IsInteger(mergedItem["target_user_profile_id"]) && mergedItem["target_user_profile_id"].(int) > 0 {
			var relatedInventoryUserProfile = shared.FetchByProperty(
				"user_profile",
				"Id",
				mergedItem["target_user_profile_id"],
			)

			// Populating User Profile data
			if len(relatedInventoryUserProfile) > 0 {
				var relatedUserProfile = shared.WriteStructToInterface(relatedInventoryUserProfile[0])

				mergedItem["target_user_profile"] = map[string]interface{}{
					"title": relatedUserProfile["first_name"].(string) + " " + relatedUserProfile["last_name"].(string),
					"id":    relatedUserProfile["id"],
				}
			}
		}

		if shared.IsInteger(mergedItem["supplier_id"]) && mergedItem["supplier_id"].(int) > 0 {
			var relatedInventorySupplier = shared.FetchByProperty(
				"suppliers",
				"Id",
				mergedItem["supplier_id"],
			)
			if len(relatedInventorySupplier) > 0 {
				var relatedSupplier = shared.WriteStructToInterface(relatedInventorySupplier[0])

				mergedItem["supplier"] = map[string]interface{}{
					"title": relatedSupplier["title"],
					"id":    relatedSupplier["id"],
				}
			}
		}

		if shared.IsInteger(mergedItem["depreciation_type_id"]) && mergedItem["depreciation_type_id"].(int) > 0 {
			var relatedInventoryDepreciationType = shared.FetchByProperty(
				"basic_inventory_depreciation_types",
				"Id",
				mergedItem["depreciation_type_id"],
			)
			if len(relatedInventoryDepreciationType) > 0 {
				var relatedDepreciationType = shared.WriteStructToInterface(relatedInventoryDepreciationType[0])

				mergedItem["depreciation_type"] = map[string]interface{}{
					"title": relatedDepreciationType["title"],
					"id":    relatedDepreciationType["id"],
				}
			}
		}

		if mergedItem["type"].(string) == "immovable" && shared.IsInteger(mergedItem["real_estate_id"]) && mergedItem["real_estate_id"].(int) > 0 {
			var relatedInventoryRealEstate = shared.FetchByProperty(
				"basic_inventory_real_estates",
				"Id",
				mergedItem["real_estate_id"],
			)
			if len(relatedInventoryRealEstate) > 0 {
				var relatedRealEstate = shared.WriteStructToInterface(relatedInventoryRealEstate[0])

				mergedItem["real_estate"] = map[string]interface{}{
					"id":                         relatedRealEstate["id"],
					"type_id":                    relatedRealEstate["type_id"],
					"square_area":                relatedRealEstate["square_area"],
					"land_serial_number":         relatedRealEstate["land_serial_number"],
					"estate_serial_number":       relatedRealEstate["estate_serial_number"],
					"ownership_type":             relatedRealEstate["ownership_type"],
					"ownership_scope":            relatedRealEstate["ownership_scope"],
					"ownership_investment_scope": relatedRealEstate["ownership_investment_scope"],
					"limitations_description":    relatedRealEstate["limitations_description"],
					"property_document":          relatedRealEstate["property_document"],
					"limitation_id":              relatedRealEstate["limitation_id"],
					"document":                   relatedRealEstate["document"],
					"file_id":                    relatedRealEstate["file_id"],
				}
			}
		}

		if shared.IsInteger(mergedItem["office_id"]) && mergedItem["office_id"].(int) > 0 {
			var relatedInventoryOffice = shared.FetchByProperty(
				"offices_of_organization_units",
				"Id",
				mergedItem["office_id"],
			)
			if len(relatedInventoryOffice) > 0 {
				var relatedOffice = shared.WriteStructToInterface(relatedInventoryOffice[0])

				mergedItem["office"] = map[string]interface{}{
					"title": relatedOffice["title"],
					"id":    relatedOffice["id"],
				}
			}
		}

		if shared.IsInteger(id) && id != 0 && id == mergedItem["id"] {

			BasicInventoryAssessmentsType := &structs.BasicInventoryAssessmentsTypesItem{}
			basicInventoryAssessmentsData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_assessments.json", BasicInventoryAssessmentsType)

			if err != nil {
				fmt.Printf("Fetching Basic Inventory Assessments failed because of this error - %s.\n", err)
			}

			mergedItem["assessments"] = PopulateBasicInventoryAssessmentsItemProperties(basicInventoryAssessmentsData, 0, id)

			BasicInventoryDispatchType := &structs.BasicInventoryDispatchItem{}
			basicInventoryDispatchData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_dispatch.json", BasicInventoryDispatchType)

			if err != nil {
				fmt.Printf("Fetching Job Tenders failed because of this error - %s.\n", err)
			}

			BasicInventoryDispatchItemsType := &structs.BasicInventoryDispatchItemsItem{}
			basicInventoryDispatchItemsData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_dispatch_items.json", BasicInventoryDispatchItemsType)

			if err != nil {
				fmt.Printf("Fetching Basic Inventory Dispatch Items failed because of this error - %s.\n", err)
			}

			basicInventoryDispatchItemsData = shared.FindByProperty(basicInventoryDispatchItemsData, "InventoryId", id)
			mergedItem["movements"] = []interface{}{}
			if len(basicInventoryDispatchItemsData) > 0 {
				for _, i := range basicInventoryDispatchItemsData {
					if m, ok := i.(*structs.BasicInventoryDispatchItemsItem); ok {
						basicInventoryDispatchItem := PopulateBasicInventoryDispatchItemProperties(basicInventoryDispatchData, m.DispatchId, "", 0, organizationUnitId)
						if len(basicInventoryDispatchItem) > 0 {
							movements, ok := mergedItem["movements"].([]interface{})
							if ok {
								mergedItem["movements"] = append(movements, basicInventoryDispatchItem[0])
							}
						}

					}
				}
			}

		}

		items = append(items, mergedItem)
	}

	return items
}

var BasicInventoryOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var organizationUnitId int
	var id int
	var classTypeId int
	var officeId int
	var typeParam string
	var search string
	var sourceType string
	var depreciationTypeId int

	var authToken = params.Context.Value("token").(string)

	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	if params.Args["class_type_id"] == nil {
		classTypeId = 0
	} else {
		classTypeId = params.Args["class_type_id"].(int)
	}

	if params.Args["office_id"] == nil {
		officeId = 0
	} else {
		officeId = params.Args["office_id"].(int)
	}

	if params.Args["type"] == nil {
		typeParam = ""
	} else {
		typeParam = params.Args["type"].(string)
	}

	if params.Args["search"] == nil {
		search = ""
	} else {
		search = params.Args["search"].(string)
	}

	if params.Args["source_type"] == nil {
		sourceType = ""
	} else {
		sourceType = params.Args["source_type"].(string)
	}

	if params.Args["depreciation_type_id"] == nil {
		depreciationTypeId = 0
	} else {
		depreciationTypeId = params.Args["depreciation_type_id"].(int)
	}

	if authToken == "sss" {
		organizationUnitId = 1
	} else {
		organizationUnitId = 2
	}

	// if params.Args["organization_unit_id"] == nil {
	// 	organizationUnitId = 0
	// } else {
	// 	organizationUnitId = params.Args["organization_unit_id"].(int)
	// }

	page := params.Args["page"]
	size := params.Args["size"]

	BasicInventoryType := &structs.BasicInventoryItem{}
	BasicInventoryData, BasicInventoryDataErr := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_items.json", BasicInventoryType)

	if BasicInventoryDataErr != nil {
		fmt.Printf("Fetching Basic Inventory failed because of this error - %s.\n", BasicInventoryDataErr)
	}

	// Filtering by search
	if search != "" && shared.IsString(search) {
		BasicInventoryData = shared.FindByProperty(BasicInventoryData, "Title", search, true)
	}

	// Populate data for each Basic Inventory
	items = PopulateBasicInventoryItemProperties(BasicInventoryData, organizationUnitId, id, typeParam, classTypeId, officeId, search, sourceType, depreciationTypeId)

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

var BasicInventoryDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var item []interface{}
	var id int
	var organizationUnitId int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	var authToken = params.Context.Value("token").(string)
	if authToken == "sss" {
		organizationUnitId = 1
	} else {
		organizationUnitId = 2
	}

	// if params.Args["organization_unit_id"] == nil {
	// 	organizationUnitId = 0
	// } else {
	// 	organizationUnitId = params.Args["organization_unit_id"].(int)
	// }

	BasicInventoryDetailsType := &structs.BasicInventoryDetailsItem{}
	BasicInventoryDetailsData, BasicInventoryDataErr := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_items.json", BasicInventoryDetailsType)

	if BasicInventoryDataErr != nil {
		fmt.Printf("Fetching Basic Inventory Details failed because of this error - %s.\n", BasicInventoryDataErr)
	}

	// Populate data for each Basic Inventory
	item = PopulateBasicInventoryItemProperties(BasicInventoryDetailsData, organizationUnitId, id, "", 0, 0, "", "", 0)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"item":    item,
	}, nil
}

var BasicInventoryInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var dataArray []structs.BasicInventoryInsertItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BasicInventoryType := &structs.BasicInventoryInsertItem{}
	var organizationUnitId int
	var results []interface{}
	_ = json.Unmarshal(dataBytes, &dataArray)

	var authToken = params.Context.Value("token").(string)
	if authToken == "sss" {
		organizationUnitId = 1
	} else {
		organizationUnitId = 2
	}

	if len(dataArray) > 0 {
		for _, data := range dataArray {
			itemId := data.Id
			data.OrganizationUnitId = organizationUnitId

			basicInventoryData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_items.json", BasicInventoryType)

			if err != nil {
				fmt.Printf("Fetching Basic Inventory Details failed because of this error - %s.\n", err)
			}

			if shared.IsInteger(itemId) && itemId != 0 {
				basicInventoryData = shared.FilterByProperty(basicInventoryData, "Id", itemId)
			} else {
				data.Id = shared.GetRandomNumber()
			}
			if data.Type == "immovable" && data.RealEstate != nil && shared.IsString(data.RealEstate.TypeId) && data.RealEstate.TypeId != "" {
				RealEstateType := &structs.BasicInventoryRealEstatesItem{}
				realEstateData, err := shared.ReadJson("http://localhost:8080/mocked-data/basic_inventory_real_estates.json", RealEstateType)

				if err != nil {
					fmt.Printf("Fetching Basic Inventory Real Estates failed because of this error - %s.\n", err)
				}

				realEstateItemId := data.RealEstate.Id

				if shared.IsInteger(realEstateItemId) && realEstateItemId != 0 {
					realEstateData = shared.FilterByProperty(realEstateData, "Id", realEstateItemId)
				} else {
					data.RealEstate.Id = shared.GetRandomNumber()
				}

				data.RealEstateId = data.RealEstate.Id
				var updatedDataRealEstate = append(realEstateData, data.RealEstate)
				_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_real_estates.json"), updatedDataRealEstate)

			}

			data.RealEstate = nil

			var updatedData = append(basicInventoryData, data)

			// Populate data for each Basic Inventory
			sliceData := []interface{}{data}
			var populatedData = PopulateBasicInventoryItemProperties(sliceData, data.OrganizationUnitId, itemId, "", 0, 0, "", "", 0)
			results = append(results, populatedData[0])

			_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/basic_inventory_items.json"), updatedData)

		}
	}
	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    results,
	}, nil
}
