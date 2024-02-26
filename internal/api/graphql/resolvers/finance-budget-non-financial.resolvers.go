package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) NonFinancialBudgetOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if requestID, ok := params.Args["request_id"].(int); ok && requestID != 0 {
		request, err := r.Repo.GetBudgetRequest(requestID)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		nonFinancialBudget, err := r.Repo.GetNonFinancialBudget(request.BudgetID)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		nonFinancialBudgetResItem, err := buildNonFinancialBudgetResItem(r.Repo, *nonFinancialBudget)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.NonFinancialBudgetResItem{nonFinancialBudgetResItem},
			Total:   1,
		}, nil
	}

	budgetID, ok := params.Args["budget_id"].(int)
	if !ok || budgetID == 0 {
		return errors.HandleAPIError(fmt.Errorf("budget_id is required"))
	}

	input := dto.GetNonFinancialBudgetListInputMS{}
	if organizationUnitID, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitID != 0 {
		request, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{OrganizationUnitID: &organizationUnitID, BudgetID: budgetID})
		if err != nil {
			return errors.HandleAPIError(err)
		}
		requestIDList := []int{request.ID}
		input.RequestIDList = &requestIDList
	} else {
		requests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{BudgetID: budgetID})
		if err != nil {
			return errors.HandleAPIError(err)
		}
		var requestIDs []int
		for _, request := range requests {
			requestIDs = append(requestIDs, request.ID)
		}
		input.RequestIDList = &requestIDs
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

	request, err := r.GetBudgetRequest(nonFinancialBudget.RequestID)
	if err != nil {
		return nil, err
	}
	requestResItem, err := buildBudgetRequestResponseItem(r, request)
	if err != nil {
		return nil, err
	}
	resItem.Request = *requestResItem
	resItem.Status = string(requestResItem.Status)

	activityRequest, err := buildActivityRequestResItem(r, resItem)
	if err != nil {
		return nil, err
	}

	resItem.ActivityRequest = *activityRequest

	return resItem, nil
}

func buildActivityRequestResItem(r repository.MicroserviceRepositoryInterface, nonFinancialBudget *dto.NonFinancialBudgetResItem) (*dto.ActivityRequestResItem, error) {
	activity, err := r.GetActivityByUnit(nonFinancialBudget.Request.OrganizationUnit.ID)
	if err != nil {
		return nil, err
	}
	activityResItem, err := buildActivityResItem(r, *activity)
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

func (r *Resolver) NonFinancialBudgetInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data dto.CreateNonFinancialBudget
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	requestType := structs.NonFinancialRequestType
	request, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &data.OrganizationUnitID,
		BudgetID:           data.BudgetID,
		RequestType:        &requestType,
	})
	if err != nil {
		return nil, err
	}
	data.RequestID = request.ID

	item, err := r.upsertNonFinancialBudget(data)
	if err != nil {
		return nil, err
	}

	request.Status = structs.BudgetRequestFilledStatus
	_, err = r.Repo.UpdateBudgetRequest(&request)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	resItem, err := buildNonFinancialBudgetResItem(r.Repo, *item)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	response.Item = resItem

	return response, nil
}

// upsertNonFinancialBudget processes the creation or update of a non-financial budget.
func (r *Resolver) upsertNonFinancialBudget(data dto.CreateNonFinancialBudget) (*structs.NonFinancialBudgetItem, error) {
	var item *structs.NonFinancialBudgetItem
	var err error

	if data.ID != 0 {
		item, err = r.Repo.UpdateNonFinancialBudget(data.ID, &data.NonFinancialBudgetItem)
		if err != nil {
			return nil, err
		}
	} else {
		item, err = r.Repo.CreateNonFinancialBudget(&data.NonFinancialBudgetItem)
		if err != nil {
			return nil, err
		}
	}

	if err := r.upsertNonFinancialGoals(data.Goals, item.ID); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return item, nil
}

// upsertNonFinancialGoals processes the goals and indicators for a budget.
func (r *Resolver) upsertNonFinancialGoals(goalsData []dto.CreateNonGinancialGoal, budgetID int) error {
	goalsToDelete := make(map[int]bool)

	goals, err := r.Repo.GetNonFinancialGoalList(&dto.GetNonFinancialGoalListInputMS{NonFinancialBudgetID: &budgetID})
	if err != nil {
		return err
	}
	for _, goal := range goals {
		goalsToDelete[goal.ID] = true
	}

	for _, goal := range goalsData {
		insertGoalItem := goal.NonFinancialGoalItem
		insertGoalItem.NonFinancialBudgetID = budgetID

		if insertGoalItem.ID != 0 {
			updatedGoal, err := r.Repo.UpdateNonFinancialGoal(insertGoalItem.ID, &insertGoalItem)
			if err != nil {
				return err
			}

			err = r.updateIndicators(updatedGoal.ID, goal.Indicators)
			if err != nil {
				return err
			}

			goalsToDelete[goal.ID] = false
		} else {
			createdGoal, err := r.Repo.CreateNonFinancialGoal(&insertGoalItem)
			if err != nil {
				return err
			}

			err = r.createIndicators(createdGoal.ID, goal.Indicators)
			if err != nil {
				return err
			}
		}
	}

	for id, delete := range goalsToDelete {
		if delete {
			err := r.Repo.DeleteNonFinancialGoal(id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Resolver) updateIndicators(goalID int, indicators []structs.NonFinancialGoalIndicatorItem) error {
	indicatorsToDelete := make(map[int]bool)

	existingIndicators, err := r.Repo.GetNonFinancialGoalIndicatorList(&dto.GetNonFinancialGoalIndicatorListInputMS{GoalID: &goalID})
	if err != nil {
		return err
	}

	for _, indicator := range existingIndicators {
		indicatorsToDelete[indicator.ID] = true
	}

	for _, indicator := range indicators {
		indicator.GoalID = goalID

		if indicator.ID != 0 {
			_, err := r.Repo.UpdateNonFinancialGoalIndicator(indicator.ID, &indicator)
			if err != nil {
				return err
			}

			indicatorsToDelete[indicator.ID] = false
		} else {
			_, err := r.Repo.CreateNonFinancialGoalIndicator(&indicator)
			if err != nil {
				return err
			}
		}
	}

	for id, delete := range indicatorsToDelete {
		if delete {
			err := r.Repo.DeleteNonFinancialGoalIndicator(id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Resolver) createIndicators(goalID int, indicators []structs.NonFinancialGoalIndicatorItem) error {
	for _, indicator := range indicators {
		indicator.GoalID = goalID

		_, err := r.Repo.CreateNonFinancialGoalIndicator(&indicator)
		if err != nil {
			return err
		}

	}

	return nil
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
