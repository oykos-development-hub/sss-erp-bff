package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) NonFinancialBudgetOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		activity, err := r.Repo.GetActivity(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*structs.ActivitiesItem{activity},
			Total:   1,
		}, nil
	}

	input := dto.GetNonFinancialBudgetListInputMS{}
	if organizationUnitID, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitID != 0 {
		input.OrganizationUnitID = &organizationUnitID
	}
	if budgetID, ok := params.Args["budget_id"].(int); ok && budgetID != 0 {
		input.BudgetID = &budgetID
	}
	if search, ok := params.Args["search"].(string); ok {
		input.Search = &search
	}

	nonFinancialBudgets, err := r.Repo.GetNonFinancialBudgetList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	nonFinancialBudgetResItemList, err := buildNonFinancialBudgetResItemList(r.Repo, nonFinancialBudgets)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   nonFinancialBudgetResItemList,
		Total:   len(nonFinancialBudgetResItemList),
	}, nil
}

func buildNonFinancialBudgetResItemList(r repository.MicroserviceRepositoryInterface, nonFinancialBudgetList []structs.NonFinancialBudgetItem) (nonFinancialBudgetResItemList []*dto.NonFinancialBudgetResItem, err error) {
	for _, nonFinancialBudget := range nonFinancialBudgetList {
		nonFinancialBudgetResItem, err := buildNonFinancialBudgetResItem(r, nonFinancialBudget)
		if err != nil {
			return nil, err
		}
		nonFinancialBudgetResItemList = append(nonFinancialBudgetResItemList, nonFinancialBudgetResItem)
	}

	return
}

func buildNonFinancialBudgetResItem(r repository.MicroserviceRepositoryInterface, nonFinancialBudget structs.NonFinancialBudgetItem) (*dto.NonFinancialBudgetResItem, error) {
	resItem := &dto.NonFinancialBudgetResItem{
		ID:                      nonFinancialBudget.ID,
		ImplContactFullName:     nonFinancialBudget.ImplContactFullName,
		ImplContactWorkingPlace: nonFinancialBudget.ImplContactWorkingPlace,
		ImplContactPhone:        nonFinancialBudget.ImplContactPhone,
		ImplContactEmail:        nonFinancialBudget.ImplContactEmail,
		ContactFullName:         nonFinancialBudget.ContactFullName,
		ContactWorkingPlace:     nonFinancialBudget.ContactWorkingPlace,
		ContactPhone:            nonFinancialBudget.ContactPhone,
		ContactEmail:            nonFinancialBudget.ContactEmail,
	}

	organizationUnit, err := r.GetOrganizationUnitByID(nonFinancialBudget.OrganizationUnitID)
	if err != nil {
		return nil, err
	}
	resItem.OrganizationUnit = dto.DropdownSimple{ID: organizationUnit.ID, Title: organizationUnit.Title}

	activityRequest, err := buildActivityRequestResItem(r, resItem)
	if err != nil {
		return nil, err
	}

	resItem.ActivityRequest = *activityRequest

	return resItem, nil
}

func buildActivityRequestResItem(r repository.MicroserviceRepositoryInterface, nonFinancialBudget *dto.NonFinancialBudgetResItem) (*dto.ActivityRequestResItem, error) {
	activities, err := r.GetActivityList(&dto.GetFinanceActivityListInputMS{OrganizationUnitID: &nonFinancialBudget.OrganizationUnit.ID})
	if err != nil {
		return nil, err
	}
	activity := activities[0]

	activityResItem, err := buildActivityResItem(r, activity)
	if err != nil {
		return nil, err
	}

	goalList, err := r.GetNonFinancialGoalList(&dto.GetNonFinancialGoalListInputMS{
		ActivityID:           &activity.ID,
		NonFinancialBudgetID: &nonFinancialBudget.ID,
	})
	if err != nil {
		return nil, err
	}

	goalResItemList, err := buildActivityGoalRequestResItemList(r, goalList)
	if err != nil {
		return nil, err
	}

	activityRequest := &dto.ActivityRequestResItem{
		ID:               activityResItem.ID,
		SubProgram:       activityResItem.SubProgram,
		OrganizationUnit: activityResItem.OrganizationUnit,
		Title:            activityResItem.Title,
		Description:      activityResItem.Description,
		Code:             activityResItem.Code,
		Goals:            goalResItemList,
	}

	return activityRequest, nil
}

func buildActivityGoalRequestResItemList(r repository.MicroserviceRepositoryInterface, goals []structs.NonFinancialGoalItem) (goalsRequestResItemList []*dto.ActivityGoalRequestResItem, err error) {
	for _, goal := range goals {
		goalRequestResItem, err := buildGoalRequestResItem(r, goal)
		if err != nil {
			return nil, err
		}
		goalsRequestResItemList = append(goalsRequestResItemList, goalRequestResItem)
	}

	return
}

func buildGoalRequestResItem(r repository.MicroserviceRepositoryInterface, goal structs.NonFinancialGoalItem) (*dto.ActivityGoalRequestResItem, error) {
	resItem := &dto.ActivityGoalRequestResItem{
		ID:          goal.ID,
		Title:       goal.Title,
		Description: goal.Description,
	}

	indicators, err := r.GetNonFinancialGoalIndicatorList(&dto.GetNonFinancialGoalIndicatorListInputMS{
		GoalID: &goal.ID,
	})
	if err != nil {
		return nil, err
	}

	indicatorResItemList, err := buildActivityGoalIndicatorRequestResItemList(r, indicators)
	if err != nil {
		return nil, err
	}

	resItem.Indicators = indicatorResItemList

	return resItem, nil
}

func buildActivityGoalIndicatorRequestResItemList(r repository.MicroserviceRepositoryInterface, indicators []structs.NonFinancialGoalIndicatorItem) (goalsRequestResItemList []*dto.BudgetActivityGoalIndicatorResItem, err error) {
	for _, indicator := range indicators {
		goalIndicatorRequestResItem, err := buildGoalIndicatorRequestResItem(r, indicator)
		if err != nil {
			return nil, err
		}
		goalsRequestResItemList = append(goalsRequestResItemList, goalIndicatorRequestResItem)
	}

	return
}

func buildGoalIndicatorRequestResItem(r repository.MicroserviceRepositoryInterface, indicator structs.NonFinancialGoalIndicatorItem) (*dto.BudgetActivityGoalIndicatorResItem, error) {
	resItem := &dto.BudgetActivityGoalIndicatorResItem{
		ID:                       indicator.ID,
		PerformanceIndicatorCode: indicator.PerformanceIndicatorCode,
		IndicatorSource:          indicator.IndicatorSource,
		BaseYear:                 indicator.BaseYear,
		GenderEquality:           indicator.GenderEquality,
		BaseValue:                indicator.BaseValue,
		SourceOfInformation:      indicator.SourceOfInformation,
		UnitOfMeasure:            indicator.UnitOfMeasure,
		IndicatorDescription:     indicator.IndicatorDescription,
		PlannedValue1:            indicator.PlannedValue1,
		RevisedValue1:            indicator.RevisedValue1,
		AchievedValue1:           indicator.AchievedValue1,
		PlannedValue2:            indicator.PlannedValue2,
		RevisedValue2:            indicator.RevisedValue2,
		AchievedValue2:           indicator.AchievedValue2,
		PlannedValue3:            indicator.PlannedValue3,
		RevisedValue3:            indicator.RevisedValue3,
		AchievedValue3:           indicator.AchievedValue3,
	}

	return resItem, nil
}

func (r *Resolver) BudgetActivityNotFinanciallyInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.NonFinancialBudgetItem
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	itemID := data.ID

	if itemID != 0 {
		item, err := r.Repo.UpdateNonFinancialBudget(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		resItem, err := buildNonFinancialBudgetResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		item, err := r.Repo.CreateNonFinancialBudget(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		resItem, err := buildNonFinancialBudgetResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = resItem
	}

	return response, nil
}

func (r *Resolver) NonFinancialGoalInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.NonFinancialGoalItem
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	itemID := data.ID

	if itemID != 0 {
		item, err := r.Repo.UpdateNonFinancialGoal(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		resItem, err := buildGoalRequestResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		item, err := r.Repo.CreateNonFinancialGoal(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		resItem, err := buildGoalRequestResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = resItem
	}

	return response, nil
}

func (r *Resolver) NonFinancialGoalIndicatorInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.NonFinancialGoalIndicatorItem
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	itemID := data.ID

	if itemID != 0 {
		item, err := r.Repo.UpdateNonFinancialGoalIndicator(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		resItem, err := buildGoalIndicatorRequestResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		item, err := r.Repo.CreateNonFinancialGoalIndicator(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		resItem, err := buildGoalIndicatorRequestResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = resItem
	}

	return response, nil
}

func (r *Resolver) CheckBudgetActivityNotFinanciallyIsDoneResolver(params graphql.ResolveParams) (interface{}, error) {
	// check := true
	// id := params.Args["id"].(int)

	// BudgetActivityNotFinanciallyType := &structs.BudgetActivityNotFinanciallyItem{}
	// BudgetActivityNotFinanciallyData, err := shared.ReadJSON(shared.GetDataRoot()+"/budget_activity_not_financially.json", BudgetActivityNotFinanciallyType)

	// if err != nil {
	// 	fmt.Printf("Fetching account_budget_activity failed because of this error - %s.\n", err)
	// }
	// BudgetActivityNotFinanciallyData = shared.FindByProperty(BudgetActivityNotFinanciallyData, "ID", id)
	// // Populate data for each Basic Inventory Real Estates
	// if len(BudgetActivityNotFinanciallyData) > 0 {
	// 	for _, item := range BudgetActivityNotFinanciallyData {
	// 		var mergedItem = shared.WriteStructToInterface(item)

	// 		var allProgramsNotFinanciallyData = shared.FetchByProperty(
	// 			"budget_program_not_financially",
	// 			"BudgetNotFinanciallyID",
	// 			mergedItem["id"].(int),
	// 		)
	// 		if len(allProgramsNotFinanciallyData) == 3 {
	// 			for _, programData := range allProgramsNotFinanciallyData {
	// 				if program, ok := programData.(*structs.BudgetActivityNotFinanciallyProgramItem); ok {
	// 					var goalsNotFinanciallyData = shared.FetchByProperty(
	// 						"budget_goals_not_financially",
	// 						"BudgetProgramID",
	// 						program.ID,
	// 					)
	// 					if len(goalsNotFinanciallyData) > 0 {
	// 						for _, goalData := range goalsNotFinanciallyData {
	// 							if goal, ok := goalData.(*structs.BudgetActivityNotFinanciallyGoalsItem); ok {
	// 								//budget_indicator_not_financially
	// 								var indicatorNotFinanciallyData = shared.FetchByProperty(
	// 									"budget_indicator_not_financially",
	// 									"BudgetProgramID",
	// 									goal.ID,
	// 								)

	// 								if len(indicatorNotFinanciallyData) == 0 {
	// 									check = false
	// 									break
	// 								}
	// 							}
	// 						}
	// 					} else {
	// 						check = false
	// 						break
	// 					}
	// 				}
	// 			}
	// 		} else {
	// 			check = false
	// 		}

	// 	}
	// }

	// return map[string]interface{}{
	// 	"status":  "success",
	// 	"message": "You check this item!",
	// 	"item":    check,
	// }, nil
	return nil, nil

}
