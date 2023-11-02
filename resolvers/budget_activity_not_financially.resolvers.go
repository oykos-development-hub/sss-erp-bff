package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func BudgetActivityNotFinanciallyItemProperties(notFinanciallyItems []interface{}) []interface{} {
	var items []interface{}

	for _, item := range notFinanciallyItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by budget ID
		if mergedItem["id"].(int) == 0 {
			continue
		}

		var requestData = shared.FetchByProperty(
			"request_budget",
			"Id",
			mergedItem["request_id"].(int),
		)

		if len(requestData) > 0 {
			for _, requestItem := range requestData {
				if request, ok := requestItem.(*structs.RequestBudgetType); ok {

					var activitiesData = shared.FetchByProperty(
						"program",
						"Id",
						request.ActivityId,
					)
					var allProgramsNotFinanciallyData = shared.FetchByProperty(
						"budget_program_not_financially",
						"BudgetNotFinanciallyId",
						mergedItem["id"].(int),
					)
					activitiesNotFinanciallyData := shared.FindByProperty(allProgramsNotFinanciallyData, "ProgramId", request.ActivityId)
					if len(activitiesData) > 0 {
						for _, activityData := range activitiesData {

							if activity, ok := activityData.(*structs.ProgramItem); ok {
								var activitiesNotFinancially = shared.WriteStructToInterface(activitiesNotFinanciallyData[0])
								var goalsNotFinanciallyData = shared.FetchByProperty(
									"budget_goals_not_financially",
									"BudgetProgramId",
									activitiesNotFinancially["id"],
								)
								mergedItem["activity"] = map[string]interface{}{
									"id":          activity.Id,
									"title":       activity.Title,
									"code":        activitiesNotFinancially["code"],
									"description": activitiesNotFinancially["description"],
									"goals":       goalsNotFinanciallyData,
								}

								var subroutinesData = shared.FetchByProperty(
									"program",
									"Id",
									activity.ParentId,
								)

								subroutineNotFinanciallyData := shared.FindByProperty(allProgramsNotFinanciallyData, "ProgramId", activity.ParentId)

								if len(subroutinesData) > 0 {
									for _, subroutineData := range subroutinesData {

										if subprogram, ok := subroutineData.(*structs.ProgramItem); ok {
											var subroutineNotFinancially = shared.WriteStructToInterface(subroutineNotFinanciallyData[0])
											var goalsNotFinanciallyData = shared.FetchByProperty(
												"budget_goals_not_financially",
												"BudgetProgramId",
												subroutineNotFinancially["id"],
											)
											mergedItem["subprogram"] = map[string]interface{}{
												"id":          subprogram.Id,
												"title":       subprogram.Title,
												"code":        subroutineNotFinancially["code"],
												"description": subroutineNotFinancially["description"],
												"goals":       goalsNotFinanciallyData,
											}

											var programsData = shared.FetchByProperty(
												"program",
												"Id",
												subprogram.ParentId,
											)

											programNotFinanciallyData := shared.FindByProperty(allProgramsNotFinanciallyData, "ProgramId", subprogram.ParentId)

											if len(programsData) > 0 {
												for _, programData := range programsData {

													if program, ok := programData.(*structs.ProgramItem); ok {
														var programNotFinancially = shared.WriteStructToInterface(programNotFinanciallyData[0])
														var goalsNotFinanciallyData = shared.FetchByProperty(
															"budget_goals_not_financially",
															"BudgetProgramId",
															programNotFinancially["id"],
														)
														mergedItem["program"] = map[string]interface{}{
															"id":          program.Id,
															"title":       program.Title,
															"code":        programNotFinancially["code"],
															"description": programNotFinancially["description"],
															"goals":       goalsNotFinanciallyData,
														}
													}
												}
											}
										}
									}
								}
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

func BudgetActivityNotFinanciallyItemInductorProperties(notFinanciallyInductorItems []interface{}, id int) []interface{} {
	var items []interface{}
	for _, item := range notFinanciallyInductorItems {
		var mergedItem = shared.WriteStructToInterface(item)

		items = append(items, mergedItem)
	}
	return items
}

var BudgetActivityNotFinanciallyOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var id int

	if params.Args["request_id"] == nil {
		fmt.Printf("Id request is important")
		return nil, nil
	} else {
		id = params.Args["request_id"].(int)
	}

	BudgetActivityNotFinanciallyType := &structs.BudgetActivityNotFinanciallyItem{}
	BudgetActivityNotFinanciallyData, err := shared.ReadJson(shared.GetDataRoot()+"/budget_activity_not_financially.json", BudgetActivityNotFinanciallyType)

	if err != nil {
		fmt.Printf("Fetching account_budget_activity failed because of this error - %s.\n", err)
	}
	BudgetActivityNotFinanciallyData = shared.FindByProperty(BudgetActivityNotFinanciallyData, "RequestId", id)
	// Populate data for each Basic Inventory Real Estates
	items = BudgetActivityNotFinanciallyItemProperties(BudgetActivityNotFinanciallyData)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"item":    items[0],
	}, nil
}

var BudgetActivityNotFinanciallyInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgetActivityNotFinanciallyItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgetActivityNotFinanciallyType := &structs.BudgetActivityNotFinanciallyItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	BudgetActivityNotFinanciallyData, err := shared.ReadJson(shared.GetDataRoot()+"/budget_activity_not_financially.json", BudgetActivityNotFinanciallyType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
	}

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgetActivityNotFinanciallyItemInductorProperties(sliceData, itemId)

	if shared.IsInteger(itemId) && itemId != 0 {
		BudgetActivityNotFinanciallyData = shared.FilterByProperty(BudgetActivityNotFinanciallyData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(BudgetActivityNotFinanciallyData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/budget_activity_not_financially.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var BudgetActivityNotFinanciallyProgramInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgetActivityNotFinanciallyProgramItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgetActivityNotFinanciallyProgramType := &structs.BudgetActivityNotFinanciallyProgramItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	BudgetActivityNotFinanciallyProgramData, err := shared.ReadJson(shared.GetDataRoot()+"/budget_program_not_financially.json", BudgetActivityNotFinanciallyProgramType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
	}

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgetActivityNotFinanciallyItemInductorProperties(sliceData, itemId)

	if shared.IsInteger(itemId) && itemId != 0 {
		BudgetActivityNotFinanciallyProgramData = shared.FilterByProperty(BudgetActivityNotFinanciallyProgramData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(BudgetActivityNotFinanciallyProgramData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/budget_program_not_financially.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var BudgetActivityNotFinanciallyGoalsInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgetActivityNotFinanciallyGoalsItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgetActivityNotFinanciallyProgramType := &structs.BudgetActivityNotFinanciallyGoalsItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	BudgetActivityNotFinanciallyGoalsData, err := shared.ReadJson(shared.GetDataRoot()+"/budget_goals_not_financially.json", BudgetActivityNotFinanciallyProgramType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
	}

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgetActivityNotFinanciallyItemInductorProperties(sliceData, itemId)

	if shared.IsInteger(itemId) && itemId != 0 {
		BudgetActivityNotFinanciallyGoalsData = shared.FilterByProperty(BudgetActivityNotFinanciallyGoalsData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(BudgetActivityNotFinanciallyGoalsData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/budget_goals_not_financially.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var BudgetActivityNotFinanciallyInductorResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var goalsId int
	var id int

	if params.Args["goals_id"] == nil {
		fmt.Printf("Id goals request is important")
		return nil, nil
	} else {
		goalsId = params.Args["goals_id"].(int)
	}

	if params.Args["id"] == nil {
		return nil, nil
	} else {
		id = params.Args["id"].(int)
	}

	BudgetActivityNotFinanciallyIndicatorType := &structs.BudgetActivityNotFinanciallyIndicatorItem{}
	BudgetActivityNotFinanciallyIndicatorData, err := shared.ReadJson(shared.GetDataRoot()+"/budget_indicator_not_financially.json", BudgetActivityNotFinanciallyIndicatorType)

	if err != nil {
		fmt.Printf("Fetching account_budget_activity failed because of this error - %s.\n", err)
	}
	BudgetActivityNotFinanciallyIndicatorData = shared.FindByProperty(BudgetActivityNotFinanciallyIndicatorData, "GoalsId", goalsId)
	// Populate data for each Basic Inventory Real Estates
	items = BudgetActivityNotFinanciallyItemInductorProperties(BudgetActivityNotFinanciallyIndicatorData, id)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"items":   items,
	}, nil
}

var BudgetActivityNotFinanciallyInductorInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgetActivityNotFinanciallyIndicatorItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgetActivityNotFinanciallyIndicatorType := &structs.BudgetActivityNotFinanciallyIndicatorItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	BudgetActivityNotFinanciallyIndicatorData, err := shared.ReadJson(shared.GetDataRoot()+"/budget_indicator_not_financially.json", BudgetActivityNotFinanciallyIndicatorType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
	}

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgetActivityNotFinanciallyItemInductorProperties(sliceData, itemId)

	if shared.IsInteger(itemId) && itemId != 0 {
		BudgetActivityNotFinanciallyIndicatorData = shared.FilterByProperty(BudgetActivityNotFinanciallyIndicatorData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(BudgetActivityNotFinanciallyIndicatorData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/budget_indicator_not_financially.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var CheckBudgetActivityNotFinanciallyIsDoneResolver = func(params graphql.ResolveParams) (interface{}, error) {
	check := true
	var id int

	if params.Args["id"] == nil {
		fmt.Printf("Id request is important")
		return nil, nil
	} else {
		id = params.Args["id"].(int)
	}

	BudgetActivityNotFinanciallyType := &structs.BudgetActivityNotFinanciallyItem{}
	BudgetActivityNotFinanciallyData, err := shared.ReadJson(shared.GetDataRoot()+"/budget_activity_not_financially.json", BudgetActivityNotFinanciallyType)

	if err != nil {
		fmt.Printf("Fetching account_budget_activity failed because of this error - %s.\n", err)
	}
	BudgetActivityNotFinanciallyData = shared.FindByProperty(BudgetActivityNotFinanciallyData, "Id", id)
	// Populate data for each Basic Inventory Real Estates
	if len(BudgetActivityNotFinanciallyData) > 0 {
		for _, item := range BudgetActivityNotFinanciallyData {
			var mergedItem = shared.WriteStructToInterface(item)

			var allProgramsNotFinanciallyData = shared.FetchByProperty(
				"budget_program_not_financially",
				"BudgetNotFinanciallyId",
				mergedItem["id"].(int),
			)
			if len(allProgramsNotFinanciallyData) == 3 {
				for _, programData := range allProgramsNotFinanciallyData {
					if program, ok := programData.(*structs.BudgetActivityNotFinanciallyProgramItem); ok {
						var goalsNotFinanciallyData = shared.FetchByProperty(
							"budget_goals_not_financially",
							"BudgetProgramId",
							program.Id,
						)
						if len(goalsNotFinanciallyData) > 0 {
							for _, goalData := range goalsNotFinanciallyData {
								if goal, ok := goalData.(*structs.BudgetActivityNotFinanciallyGoalsItem); ok {
									//budget_indicator_not_financially
									var indicatorNotFinanciallyData = shared.FetchByProperty(
										"budget_indicator_not_financially",
										"BudgetProgramId",
										goal.Id,
									)

									if len(indicatorNotFinanciallyData) == 0 {
										check = false
										break
									}
								}
							}
						} else {
							check = false
							break
						}
					}
				}
			} else {
				check = false
			}

		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You check this item!",
		"item":    check,
	}, nil
}
