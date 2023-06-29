package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func PopulateSystematizationItemProperties(systematizations []interface{}, filters ...int) []interface{} {
	var items []interface{}
	var id, organizationUnitId int

	switch len(filters) {
	case 1:
		id = filters[0]
	case 2:
		id = filters[0]
		organizationUnitId = filters[1]
	}

	for _, item := range systematizations {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}
		// Filtering by Organization Unit
		if shared.IsInteger(organizationUnitId) && organizationUnitId != 0 && organizationUnitId != mergedItem["organization_unit_id"] {
			continue
		}
		// # Related user Profile
		var relatedUserProfile = shared.FetchByProperty(
			"user_profile",
			"Id",
			mergedItem["user_profile_id"],
		)

		if len(relatedUserProfile) > 0 {
			var userProfile = shared.WriteStructToInterface(relatedUserProfile[0])

			mergedItem["user_profile"] = map[string]interface{}{
				"title": userProfile["first_name"].(string) + " " + userProfile["last_name"].(string),
				"id":    userProfile["id"],
			}
		}

		// # Related Organization Unit
		var relatedOrganizationUnit = shared.FetchByProperty(
			"organization_unit",
			"Id",
			mergedItem["organization_unit_id"],
		)

		if len(relatedOrganizationUnit) > 0 {
			var organizationUnit = shared.WriteStructToInterface(relatedOrganizationUnit[0])
			var populatedOrganizationUnitData = PopulateOrganizationUnitItemProperties(relatedOrganizationUnit)
			var sectors []interface{}

			if len(populatedOrganizationUnitData) > 0 {
				for _, unit := range populatedOrganizationUnitData {
					var unitData = unit.(map[string]interface{})
					var childrenData = unitData["children"].([]interface{})

					if unitData["id"] == organizationUnit["id"] && len(childrenData) > 0 {
						// # Related Organization Unit sectors
						for _, sector := range childrenData {
							var sectorItem = shared.WriteStructToInterface(sector)
							var sectorJobPositions []interface{}
							// # Related Job Positions in sector
							var relatedJobPositions = shared.FetchByProperty(
								"job_positions_in_organization_units",
								"ParentOrganizationUnitId",
								sectorItem["id"],
							)

							if len(relatedJobPositions) > 0 {
								for _, jobPosition := range relatedJobPositions {
									var jobPositionItem = shared.WriteStructToInterface(jobPosition)
									var jobPositionDetails = shared.FetchByProperty(
										"job_position",
										"Id",
										jobPositionItem["job_position_id"],
									)
									// # Related Employees for Job Position
									var relatedEmployeeJobPositions = shared.FetchByProperty(
										"employees_in_organization_units",
										"PositionInOrganizationUnitId",
										jobPositionItem["id"],
									)
									var josPositionEmployees []interface{}

									if len(jobPositionDetails) > 0 {
										for _, jobPositionDetailsItem := range jobPositionDetails {
											var jobPositionDetailsData = shared.WriteStructToInterface(jobPositionDetailsItem)

											jobPositionItem["job_position"] = map[string]interface{}{
												"title": jobPositionDetailsData["title"],
												"id":    jobPositionDetailsData["id"],
											}
										}
									}

									if len(relatedEmployeeJobPositions) > 0 {
										for _, employeeJobPosition := range relatedEmployeeJobPositions {
											var employeeJobPositionData = shared.WriteStructToInterface(employeeJobPosition)
											// # Related User Account
											var relatedUserAccount = shared.FetchByProperty(
												"user_account",
												"Id",
												employeeJobPositionData["user_account_id"],
											)
											if len(relatedUserAccount) > 0 {
												for _, userAccount := range relatedUserAccount {
													var userAccountData = shared.WriteStructToInterface(userAccount)
													// # Related User Profile
													var relatedUserProfile = shared.FetchByProperty(
														"user_profile",
														"UserAccountId",
														userAccountData["id"],
													)
													if len(relatedUserProfile) > 0 {
														for _, userProfile := range relatedUserProfile {
															var userProfileData = shared.WriteStructToInterface(userProfile)
															var employeeItem = map[string]interface{}{
																"title":           userProfileData["first_name"].(string) + " " + userProfileData["last_name"].(string),
																"user_profile_id": userProfileData["id"],
																"id":              employeeJobPositionData["id"],
															}
															josPositionEmployees = append(josPositionEmployees, employeeItem)
														}
													}
												}
											}
										}
									}

									jobPositionItem["employees"] = josPositionEmployees

									sectorJobPositions = append(sectorJobPositions, jobPositionItem)
								}
							}

							sectorItem["job_positions"] = sectorJobPositions

							sectors = append(sectors, sectorItem)
						}
					}
				}
			}

			mergedItem["organization_unit"] = map[string]interface{}{
				"title": organizationUnit["title"],
				"id":    organizationUnit["id"],
			}
			mergedItem["sectors"] = sectors
		}

		items = append(items, mergedItem)
	}

	return items
}

var SystematizationsOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}
	var organizationUnitId int
	if params.Args["organization_unit_id"] == nil {
		organizationUnitId = 0
	} else {
		organizationUnitId = params.Args["organization_unit_id"].(int)
	}
	page := params.Args["page"]
	size := params.Args["size"]

	SystematizationsType := &structs.Systematization{}
	SystematizationsData, SystematizationsDataErr := shared.ReadJson("http://localhost:8080/mocked-data/systematizations.json", SystematizationsType)

	if SystematizationsDataErr != nil {
		fmt.Printf("Fetching Systematizations failed because of this error - %s.\n", SystematizationsDataErr)
	}

	// Populate data for each Systematization with Organization Unit, User Profile, Sectors and Job Positions
	items = PopulateSystematizationItemProperties(SystematizationsData, id, organizationUnitId)

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

var SystematizationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	SystematizationType := &structs.Systematization{}
	SystematizationData, SystematizationDataErr := shared.ReadJson("http://localhost:8080/mocked-data/systematizations.json", SystematizationType)

	var id int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	if SystematizationDataErr != nil {
		fmt.Printf("Fetching Systematizations failed because of this error - %s.\n", SystematizationDataErr)
	}

	// Populate data for each Systematization with Organization Unit, Job Positions and Related Employees
	var items = PopulateSystematizationItemProperties(SystematizationData, id)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"items":   items,
	}, nil
}

var SystematizationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.Systematization
	dataBytes, _ := json.Marshal(params.Args["data"])
	SystematizationType := &structs.Systematization{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	systematizationData, systematizationDataErr := shared.ReadJson("http://localhost:8080/mocked-data/systematizations.json", SystematizationType)

	if systematizationDataErr != nil {
		fmt.Printf("Fetching Systematization failed because of this error - %s.\n", systematizationDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		systematizationData = shared.FilterByProperty(systematizationData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	sliceData := []interface{}{data}
	// Populate data for each Systematization with Organization Unit, Job Positions and Related Employees
	var populatedData = PopulateSystematizationItemProperties(sliceData, data.Id)
	var updatedData = append(systematizationData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/systematizations.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   populatedData,
	}, nil
}

var SystematizationDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	SystematizationType := &structs.Systematization{}
	systematizationData, systematizationDataErr := shared.ReadJson("http://localhost:8080/mocked-data/systematizations.json", SystematizationType)

	if systematizationDataErr != nil {
		fmt.Printf("Fetching User Profile's Systematization failed because of this error - %s.\n", systematizationDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		systematizationData = shared.FilterByProperty(systematizationData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/systematizations.json"), systematizationData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
