package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"reflect"
)

func PopulateResolutionItemProperties(resolutions []interface{}, id int, year string) []interface{} {
	var items []interface{}

	for _, item := range resolutions {
		// # Resolution item
		var mergedItem map[string]interface{}

		if reflect.TypeOf(item).Kind() == reflect.Map {
			mergedItem = item.(map[string]interface{})
		} else {
			mergedItem = shared.WriteStructToInterface(item)
		}

		var availableSlotsJudgesTotal = 0
		var judgesNumberTotal = 0
		var judgePresidentsNumberTotal = 0

		if shared.IsInteger(id) && id > 0 && id != mergedItem["id"] {
			continue
		}
		if shared.IsString(year) && len(year) > 0 && year != mergedItem["year"] {
			continue
		}

		// # Related Resolution item
		var relatedResolutionItems = shared.FetchByProperty(
			"judge_resolution_item",
			"ResolutionId",
			mergedItem["id"],
		)
		var mergedResolutionItems []interface{}

		if relatedResolutionItems != nil && len(relatedResolutionItems) > 0 {
			for _, resolutionItem := range relatedResolutionItems {
				var mergedResolutionItem = shared.WriteStructToInterface(resolutionItem)
				var availableSlotsJudges = mergedResolutionItem["available_slots_judges"].(int)
				var employeesNumber = 0
				var judgesNumber = 0
				var judgesPresidentNumber = 0

				availableSlotsJudgesTotal = availableSlotsJudgesTotal + availableSlotsJudges

				// # Related Organization Unit
				var relatedOrganizationUnit = shared.FetchByProperty(
					"organization_unit",
					"Id",
					mergedResolutionItem["organization_unit_id"],
				)

				if relatedOrganizationUnit != nil && len(relatedOrganizationUnit) > 0 {
					for _, organizationUnit := range relatedOrganizationUnit {
						var mergedOrganizationUnit = shared.WriteStructToInterface(organizationUnit)

						organizationUnitField := make(map[string]interface{})

						organizationUnitField["title"] = mergedOrganizationUnit["title"]
						organizationUnitField["id"] = mergedOrganizationUnit["id"]

						mergedResolutionItem["organization_unit"] = organizationUnitField

						// # Job Positions in Organization Unit
						var jobPositionsInOrganizationUnit = shared.FetchByProperty(
							"job_positions_in_organization_units",
							"ParentOrganizationUnitId",
							mergedOrganizationUnit["id"],
						)

						if jobPositionsInOrganizationUnit != nil && len(jobPositionsInOrganizationUnit) > 0 {
							for _, jobPositionInOrganizationUnit := range jobPositionsInOrganizationUnit {
								var jobPositionData = shared.WriteStructToInterface(jobPositionInOrganizationUnit)

								// # Related Job Position
								var relatedJobPosition = shared.FetchByProperty(
									"job_positions",
									"Id",
									jobPositionData["job_position_id"],
								)

								if relatedJobPosition != nil && len(relatedJobPosition) > 0 {
									for _, jobPositionItem := range relatedJobPosition {
										var jobPosition = shared.WriteStructToInterface(jobPositionItem)
										// # Employees for Job Position
										var employeesInOrganizationUnit = shared.FetchByProperty(
											"employees_in_organization_units",
											"PositionInOrganizationUnitId",
											jobPositionData["id"],
										)

										employeesNumber = employeesNumber + len(employeesInOrganizationUnit)

										if jobPosition["is_judge"].(bool) {
											// # Judges for Job Position
											var judgesInOrganizationUnit = shared.FetchByProperty(
												"employees_in_organization_units",
												"PositionInOrganizationUnitId",
												jobPositionData["id"],
											)

											judgesNumber = judgesNumber + len(judgesInOrganizationUnit)
										} else if jobPosition["is_judge_president"].(bool) {
											// # Judge Presidents for Job Position
											var judgePresidentsInOrganizationUnit = shared.FetchByProperty(
												"employees_in_organization_units",
												"PositionInOrganizationUnitId",
												jobPositionData["id"],
											)

											judgesPresidentNumber = judgesPresidentNumber + len(judgePresidentsInOrganizationUnit)
										}
									}
								}
							}
						}
					}
				}

				mergedResolutionItem["number_of_employees"] = employeesNumber
				mergedResolutionItem["number_of_judges"] = judgesNumber
				mergedResolutionItem["number_of_presidents"] = judgesPresidentNumber
				mergedResolutionItem["number_of_relocated_judges"] = 0
				mergedResolutionItem["number_of_suspended_judges"] = 0

				judgesNumberTotal = judgesNumberTotal + judgesNumber
				judgePresidentsNumberTotal = judgePresidentsNumberTotal + judgesPresidentNumber

				mergedResolutionItems = append(mergedResolutionItems, mergedResolutionItem)
			}
		}

		mergedItem["available_slots_judges"] = availableSlotsJudgesTotal
		mergedItem["number_of_judges"] = judgesNumberTotal
		mergedItem["items"] = mergedResolutionItems

		items = append(items, mergedItem)
	}

	return items
}

func PopulateJudgeItemProperties(judgeJobPositions []interface{}, isPresident bool, filters ...interface{}) map[int]interface{} {
	var userProfileId, organizationUnitId int
	var search string

	judges := make(map[int]interface{})

	switch len(filters) {
	case 1:
		userProfileId = filters[0].(int)
	case 2:
		userProfileId = filters[0].(int)
		organizationUnitId = filters[1].(int)
	case 3:
		userProfileId = filters[0].(int)
		organizationUnitId = filters[1].(int)
		search = filters[2].(string)
	}

	for _, jobPosition := range judgeJobPositions {
		var position = shared.WriteStructToInterface(jobPosition)
		var relatedJobPositions = shared.FetchByProperty(
			"job_positions_in_organization_units",
			"JobPositionId",
			position["id"],
		)

		if relatedJobPositions != nil && len(relatedJobPositions) > 0 {
			for _, relatedJobPosition := range relatedJobPositions {
				var relatedPosition = shared.WriteStructToInterface(relatedJobPosition)
				var employeeJobPositions = shared.FetchByProperty(
					"employees_in_organization_units",
					"PositionInOrganizationUnitId",
					relatedPosition["id"],
				)
				var relatedOrganizationUnits = shared.FetchByProperty(
					"organization_unit",
					"Id",
					relatedPosition["parent_organization_unit_id"],
				)
				var relatedOrganizationUnit map[string]interface{}

				if relatedOrganizationUnits != nil && len(relatedOrganizationUnits) > 0 {
					relatedOrganizationUnit = shared.WriteStructToInterface(relatedOrganizationUnits[0])

					if shared.IsInteger(organizationUnitId) && organizationUnitId > 0 && organizationUnitId != relatedOrganizationUnit["id"] {
						continue
					}
				} else if shared.IsInteger(organizationUnitId) && organizationUnitId > 0 {
					continue
				}

				if employeeJobPositions != nil && len(employeeJobPositions) > 0 {
					for _, employeeJobPosition := range employeeJobPositions {
						var employeePosition = shared.WriteStructToInterface(employeeJobPosition)
						var userAccounts = shared.FetchByProperty(
							"user_account",
							"Id",
							employeePosition["user_account_id"],
						)

						if userAccounts != nil && len(userAccounts) > 0 {
							for _, userAccount := range userAccounts {
								var account = shared.WriteStructToInterface(userAccount)
								var userProfiles = shared.FetchByProperty(
									"user_profile",
									"UserAccountId",
									account["id"],
								)

								if userProfiles != nil && len(userProfiles) > 0 {
									for _, userProfile := range userProfiles {
										var profile = shared.WriteStructToInterface(userProfile)

										if shared.IsInteger(userProfileId) && userProfileId > 0 && userProfileId != profile["id"] {
											continue
										}
										if shared.IsString(search) && len(search) > 0 {
											UserProfileName := profile["first_name"].(string) + profile["last_name"].(string)

											if !shared.StringContains(UserProfileName, search) {
												continue
											}
										}

										organizationUnitField := make(map[string]interface{})
										jobPositionField := make(map[string]interface{})
										judgeItem := make(map[string]interface{})

										var judgeNorms = shared.FetchByProperty(
											"judge_norms",
											"UserProfileId",
											profile["id"],
										)

										if judgeNorms != nil && len(judgeNorms) > 0 {
											judgeItem["norms"] = judgeNorms
										}

										jobPositionField["title"] = position["title"]
										jobPositionField["id"] = position["id"]

										organizationUnitField["title"] = relatedOrganizationUnit["title"]
										organizationUnitField["id"] = relatedOrganizationUnit["id"]

										judgeItem["id"] = profile["id"]
										judgeItem["is_judge_president"] = isPresident
										judgeItem["first_name"] = profile["first_name"]
										judgeItem["last_name"] = profile["last_name"]
										judgeItem["evaluation"] = ""
										judgeItem["created_at"] = profile["created_at"]
										judgeItem["updated_at"] = profile["updated_at"]
										judgeItem["folder_id"] = account["folder_id"]
										judgeItem["organization_unit"] = organizationUnitField
										judgeItem["job_position"] = jobPositionField

										judges[profile["id"].(int)] = judgeItem
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return judges
}

var JudgesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var total int
	var userProfileId int
	if params.Args["user_profile_id"] == nil {
		userProfileId = 0
	} else {
		userProfileId = params.Args["user_profile_id"].(int)
	}
	var organizationUnitId int
	if params.Args["organization_unit_id"] == nil {
		organizationUnitId = 0
	} else {
		organizationUnitId = params.Args["organization_unit_id"].(int)
	}
	page := params.Args["page"]
	size := params.Args["size"]
	search := params.Args["search"]

	if !shared.IsString(search) {
		search = ""
	}

	everyone := make(map[int]interface{})
	judges := make(map[int]interface{})
	judgePresidents := make(map[int]interface{})

	var judgeJobPositions = shared.FetchByProperty(
		"job_position",
		"IsJudge",
		true,
	)
	var judgePresidentJobPositions = shared.FetchByProperty(
		"job_position",
		"IsJudgePresident",
		true,
	)

	if judgeJobPositions != nil && len(judgeJobPositions) > 0 {
		judges = PopulateJudgeItemProperties(judgeJobPositions, false, userProfileId, organizationUnitId, search)
	}
	if judgePresidentJobPositions != nil && len(judgePresidentJobPositions) > 0 {
		judgePresidents = PopulateJudgeItemProperties(judgePresidentJobPositions, true, userProfileId, organizationUnitId, search)
	}

	for key, item := range judges {
		everyone[key] = item
	}
	for key, item := range judgePresidents {
		everyone[key] = item
	}

	items := make([]interface{}, 0, len(everyone))
	for _, value := range everyone {
		items = append(items, value)
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

var JudgeNormInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.JudgeNorms
	dataBytes, _ := json.Marshal(params.Args["data"])
	JudgeNormType := &structs.JudgeNorms{}

	json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	judgeNormData, judgeNormDataErr := shared.ReadJson("http://localhost:8080/mocked-data/judge_norms.json", JudgeNormType)

	if judgeNormDataErr != nil {
		fmt.Printf("Fetching Judge Norms failed because of this error - %s.\n", judgeNormDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		judgeNormData = shared.FilterByProperty(judgeNormData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(judgeNormData, data)

	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/judge_norms.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var JudgeNormDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	JudgeNormType := &structs.JudgeNorms{}
	judgeNormData, judgeNormDataErr := shared.ReadJson("http://localhost:8080/mocked-data/judge_norms.json", JudgeNormType)

	if judgeNormDataErr != nil {
		fmt.Printf("Fetching Judge Norm failed because of this error - %s.\n", judgeNormDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		judgeNormData = shared.FilterByProperty(judgeNormData, "Id", itemId)
	}

	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/judge_norms.json"), judgeNormData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

var JudgeResolutionsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var total int
	var items []interface{}
	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	year := params.Args["year"]

	if !shared.IsString(year) {
		year = ""
	}
	if !shared.IsInteger(id) {
		id = 0
	}

	JudgeResolutionsType := &structs.JudgeResolutions{}
	JudgeResolutionsData, JudgeResolutionsDataErr := shared.ReadJson("http://localhost:8080/mocked-data/judge_resolutions.json", JudgeResolutionsType)

	if JudgeResolutionsDataErr != nil {
		fmt.Printf("Fetching Judge Resolutions failed because of this error - %s.\n", JudgeResolutionsDataErr)
	}

	items = PopulateResolutionItemProperties(JudgeResolutionsData, id.(int), year.(string))

	total = len(JudgeResolutionsData)

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

var JudgeResolutionInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.JudgeResolutionData
	dataBytes, _ := json.Marshal(params.Args["data"])
	JudgeResolutionType := &structs.JudgeResolutions{}
	JudgeResolutionItemType := &structs.JudgeResolutionItems{}

	json.Unmarshal(dataBytes, &data)

	var resolutionItems = data.Items
	itemId := data.Id
	judgeResolutionData, judgeResolutionDataErr := shared.ReadJson("http://localhost:8080/mocked-data/judge_resolutions.json", JudgeResolutionType)
	judgeResolutionItemsData, judgeResolutionItemsDataErr := shared.ReadJson("http://localhost:8080/mocked-data/judge_resolution_items.json", JudgeResolutionItemType)

	if judgeResolutionDataErr != nil {
		fmt.Printf("Fetching Judge Resolutions failed because of this error - %s.\n", judgeResolutionDataErr)
	}
	if judgeResolutionItemsDataErr != nil {
		fmt.Printf("Fetching Judge Resolution Items failed because of this error - %s.\n", judgeResolutionItemsDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		judgeResolutionData = shared.FilterByProperty(judgeResolutionData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var parsedData = shared.WriteStructToInterface(data)

	if resolutionItems != nil && len(resolutionItems) > 0 {
		for _, resolutionItem := range resolutionItems {
			var resolutionItemData = shared.WriteStructToInterface(resolutionItem)

			resolutionItemData["resolution_id"] = parsedData["id"]

			if shared.IsInteger(resolutionItemData["id"]) && resolutionItemData["id"] != 0 {
				judgeResolutionItemsData = shared.FilterByProperty(judgeResolutionItemsData, "Id", resolutionItemData["id"])
			} else {
				resolutionItemData["id"] = shared.GetRandomNumber()
			}

			judgeResolutionItemsData = append(judgeResolutionItemsData, resolutionItemData)
		}
	}

	delete(parsedData, "items")

	var updatedData = append(judgeResolutionData, parsedData)

	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/judge_resolutions.json"), updatedData)
	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/judge_resolution_items.json"), judgeResolutionItemsData)

	parsedData["items"] = judgeResolutionItemsData

	var populatedItems = PopulateResolutionItemProperties([]interface{}{parsedData}, parsedData["id"].(int), "")

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   populatedItems,
	}, nil
}

var JudgeResolutionDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	JudgeResolutionType := &structs.JudgeResolutions{}
	JudgeResolutionItemType := &structs.JudgeResolutionItems{}

	judgeResolutionData, judgeResolutionDataErr := shared.ReadJson("http://localhost:8080/mocked-data/judge_resolutions.json", JudgeResolutionType)
	judgeResolutionItemsData, judgeResolutionItemsDataErr := shared.ReadJson("http://localhost:8080/mocked-data/judge_resolution_items.json", JudgeResolutionItemType)

	if judgeResolutionDataErr != nil {
		fmt.Printf("Fetching Resolutions failed because of this error - %s.\n", judgeResolutionDataErr)
	}
	if judgeResolutionItemsDataErr != nil {
		fmt.Printf("Fetching Resolution Items failed because of this error - %s.\n", judgeResolutionItemsDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		judgeResolutionData = shared.FilterByProperty(judgeResolutionData, "Id", itemId)
		judgeResolutionItemsData = shared.FilterByProperty(judgeResolutionItemsData, "ResolutionId", itemId)
	}

	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/judge_resolutions.json"), judgeResolutionData)
	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/judge_resolution_items.json"), judgeResolutionItemsData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
