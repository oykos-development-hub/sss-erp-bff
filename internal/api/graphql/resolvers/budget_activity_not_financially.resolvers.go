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
			"ID",
			mergedItem["request_id"].(int),
		)

		if len(requestData) > 0 {
			for _, requestItem := range requestData {
				if request, ok := requestItem.(*structs.RequestBudgetType); ok {

					var activitiesData = shared.FetchByProperty(
						"program",
						"ID",
						request.ActivityID,
					)
					var allProgramsNotFinanciallyData = shared.FetchByProperty(
						"budget_program_not_financially",
						"BudgetNotFinanciallyID",
						mergedItem["id"].(int),
					)
					activitiesNotFinanciallyData := shared.FindByProperty(allProgramsNotFinanciallyData, "ProgramID", request.ActivityID)
					if len(activitiesData) > 0 {
						for _, activityData := range activitiesData {

							if activity, ok := activityData.(*structs.ProgramItem); ok {
								var activitiesNotFinancially = shared.WriteStructToInterface(activitiesNotFinanciallyData[0])
								var goalsNotFinanciallyData = shared.FetchByProperty(
									"budget_goals_not_financially",
									"BudgetProgramID",
									activitiesNotFinancially["id"],
								)
								mergedItem["activity"] = map[string]interface{}{
									"id":          activity.ID,
									"title":       activity.Title,
									"code":        activitiesNotFinancially["code"],
									"description": activitiesNotFinancially["description"],
									"goals":       goalsNotFinanciallyData,
								}

								var subroutinesData = shared.FetchByProperty(
									"program",
									"ID",
									activity.ParentID,
								)

								subroutineNotFinanciallyData := shared.FindByProperty(allProgramsNotFinanciallyData, "ProgramID", activity.ParentID)

								if len(subroutinesData) > 0 {
									for _, subroutineData := range subroutinesData {

										if subprogram, ok := subroutineData.(*structs.ProgramItem); ok {
											var subroutineNotFinancially = shared.WriteStructToInterface(subroutineNotFinanciallyData[0])
											var goalsNotFinanciallyData = shared.FetchByProperty(
												"budget_goals_not_financially",
												"BudgetProgramID",
												subroutineNotFinancially["id"],
											)
											mergedItem["subprogram"] = map[string]interface{}{
												"id":          subprogram.ID,
												"title":       subprogram.Title,
												"code":        subroutineNotFinancially["code"],
												"description": subroutineNotFinancially["description"],
												"goals":       goalsNotFinanciallyData,
											}

											var programsData = shared.FetchByProperty(
												"program",
												"ID",
												subprogram.ParentID,
											)

											programNotFinanciallyData := shared.FindByProperty(allProgramsNotFinanciallyData, "ProgramID", subprogram.ParentID)

											if len(programsData) > 0 {
												for _, programData := range programsData {

													if program, ok := programData.(*structs.ProgramItem); ok {
														var programNotFinancially = shared.WriteStructToInterface(programNotFinanciallyData[0])
														var goalsNotFinanciallyData = shared.FetchByProperty(
															"budget_goals_not_financially",
															"BudgetProgramID",
															programNotFinancially["id"],
														)
														mergedItem["program"] = map[string]interface{}{
															"id":          program.ID,
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

func BudgetActivityNotFinanciallyItemInductorProperties(notFinanciallyInductorItems []interface{}) []interface{} {
	var items []interface{}
	for _, item := range notFinanciallyInductorItems {
		var mergedItem = shared.WriteStructToInterface(item)

		items = append(items, mergedItem)
	}
	return items
}

func (r *Resolver) BudgetActivityNotFinanciallyOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	id := params.Args["request_id"].(int)

	BudgetActivityNotFinanciallyType := &structs.BudgetActivityNotFinanciallyItem{}
	BudgetActivityNotFinanciallyData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget_activity_not_financially.json", BudgetActivityNotFinanciallyType)

	if err != nil {
		fmt.Printf("Fetching account_budget_activity failed because of this error - %s.\n", err)
	}
	BudgetActivityNotFinanciallyData = shared.FindByProperty(BudgetActivityNotFinanciallyData, "RequestID", id)
	// Populate data for each Basic Inventory Real Estates
	items = BudgetActivityNotFinanciallyItemProperties(BudgetActivityNotFinanciallyData)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"item":    items[0],
	}, nil
}

func (r *Resolver) BudgetActivityNotFinanciallyInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgetActivityNotFinanciallyItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgetActivityNotFinanciallyType := &structs.BudgetActivityNotFinanciallyItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	BudgetActivityNotFinanciallyData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget_activity_not_financially.json", BudgetActivityNotFinanciallyType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
	}

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgetActivityNotFinanciallyItemInductorProperties(sliceData)

	if itemID != 0 {
		BudgetActivityNotFinanciallyData = shared.FilterByProperty(BudgetActivityNotFinanciallyData, "ID", itemID)
	} else {
		data.ID = shared.GetRandomNumber()
	}

	var updatedData = append(BudgetActivityNotFinanciallyData, data)

	_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/budget_activity_not_financially.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

func (r *Resolver) BudgetActivityNotFinanciallyProgramInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgetActivityNotFinanciallyProgramItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgetActivityNotFinanciallyProgramType := &structs.BudgetActivityNotFinanciallyProgramItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	BudgetActivityNotFinanciallyProgramData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget_program_not_financially.json", BudgetActivityNotFinanciallyProgramType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
	}

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgetActivityNotFinanciallyItemInductorProperties(sliceData)

	if itemID != 0 {
		BudgetActivityNotFinanciallyProgramData = shared.FilterByProperty(BudgetActivityNotFinanciallyProgramData, "ID", itemID)
	} else {
		data.ID = shared.GetRandomNumber()
	}

	var updatedData = append(BudgetActivityNotFinanciallyProgramData, data)

	_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/budget_program_not_financially.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

func (r *Resolver) BudgetActivityNotFinanciallyGoalsInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgetActivityNotFinanciallyGoalsItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgetActivityNotFinanciallyProgramType := &structs.BudgetActivityNotFinanciallyGoalsItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	BudgetActivityNotFinanciallyGoalsData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget_goals_not_financially.json", BudgetActivityNotFinanciallyProgramType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
	}

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgetActivityNotFinanciallyItemInductorProperties(sliceData)

	if itemID != 0 {
		BudgetActivityNotFinanciallyGoalsData = shared.FilterByProperty(BudgetActivityNotFinanciallyGoalsData, "ID", itemID)
	} else {
		data.ID = shared.GetRandomNumber()
	}

	var updatedData = append(BudgetActivityNotFinanciallyGoalsData, data)

	_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/budget_goals_not_financially.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

func (r *Resolver) BudgetActivityNotFinanciallyInductorResolver(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	goalsID := params.Args["goals_id"].(int)

	BudgetActivityNotFinanciallyIndicatorType := &structs.BudgetActivityNotFinanciallyIndicatorItem{}
	BudgetActivityNotFinanciallyIndicatorData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget_indicator_not_financially.json", BudgetActivityNotFinanciallyIndicatorType)

	if err != nil {
		fmt.Printf("Fetching account_budget_activity failed because of this error - %s.\n", err)
	}
	BudgetActivityNotFinanciallyIndicatorData = shared.FindByProperty(BudgetActivityNotFinanciallyIndicatorData, "GoalsID", goalsID)
	// Populate data for each Basic Inventory Real Estates
	items = BudgetActivityNotFinanciallyItemInductorProperties(BudgetActivityNotFinanciallyIndicatorData)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"items":   items,
	}, nil
}

func (r *Resolver) BudgetActivityNotFinanciallyInductorInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.BudgetActivityNotFinanciallyIndicatorItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	BudgetActivityNotFinanciallyIndicatorType := &structs.BudgetActivityNotFinanciallyIndicatorItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	BudgetActivityNotFinanciallyIndicatorData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget_indicator_not_financially.json", BudgetActivityNotFinanciallyIndicatorType)

	if err != nil {
		fmt.Printf("Fetching Basic Inventory Depreciation Types failed because of this error - %s.\n", err)
	}

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = BudgetActivityNotFinanciallyItemInductorProperties(sliceData)

	if itemID != 0 {
		BudgetActivityNotFinanciallyIndicatorData = shared.FilterByProperty(BudgetActivityNotFinanciallyIndicatorData, "ID", itemID)
	} else {
		data.ID = shared.GetRandomNumber()
	}

	var updatedData = append(BudgetActivityNotFinanciallyIndicatorData, data)

	_ = shared.WriteJSON(shared.FormatPath(projectRoot+"/mocked-data/budget_indicator_not_financially.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

func (r *Resolver) CheckBudgetActivityNotFinanciallyIsDoneResolver(params graphql.ResolveParams) (interface{}, error) {
	check := true
	id := params.Args["id"].(int)

	BudgetActivityNotFinanciallyType := &structs.BudgetActivityNotFinanciallyItem{}
	BudgetActivityNotFinanciallyData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget_activity_not_financially.json", BudgetActivityNotFinanciallyType)

	if err != nil {
		fmt.Printf("Fetching account_budget_activity failed because of this error - %s.\n", err)
	}
	BudgetActivityNotFinanciallyData = shared.FindByProperty(BudgetActivityNotFinanciallyData, "ID", id)
	// Populate data for each Basic Inventory Real Estates
	if len(BudgetActivityNotFinanciallyData) > 0 {
		for _, item := range BudgetActivityNotFinanciallyData {
			var mergedItem = shared.WriteStructToInterface(item)

			var allProgramsNotFinanciallyData = shared.FetchByProperty(
				"budget_program_not_financially",
				"BudgetNotFinanciallyID",
				mergedItem["id"].(int),
			)
			if len(allProgramsNotFinanciallyData) == 3 {
				for _, programData := range allProgramsNotFinanciallyData {
					if program, ok := programData.(*structs.BudgetActivityNotFinanciallyProgramItem); ok {
						var goalsNotFinanciallyData = shared.FetchByProperty(
							"budget_goals_not_financially",
							"BudgetProgramID",
							program.ID,
						)
						if len(goalsNotFinanciallyData) > 0 {
							for _, goalData := range goalsNotFinanciallyData {
								if goal, ok := goalData.(*structs.BudgetActivityNotFinanciallyGoalsItem); ok {
									//budget_indicator_not_financially
									var indicatorNotFinanciallyData = shared.FetchByProperty(
										"budget_indicator_not_financially",
										"BudgetProgramID",
										goal.ID,
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
