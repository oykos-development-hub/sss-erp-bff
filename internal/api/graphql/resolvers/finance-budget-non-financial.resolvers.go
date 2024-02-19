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

	input := dto.GetNonFinancialBudgetListInputMS{}
	if organizationUnitID, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitID != 0 {
		request, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{OrganizationUnitID: &organizationUnitID})
		if err != nil {
			return errors.HandleAPIError(err)
		}
		requestIDList := []int{request.ID}
		input.RequestIDList = &requestIDList
	}
	if budgetID, ok := params.Args["budget_id"].(int); ok && budgetID != 0 {
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

	var item *structs.NonFinancialBudgetItem

	itemID := data.ID

	if itemID != 0 {
		item, err = r.Repo.UpdateNonFinancialBudget(itemID, &data.NonFinancialBudgetItem)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
	} else {
		item, err = r.Repo.CreateNonFinancialBudget(&data.NonFinancialBudgetItem)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
	}

	for _, goal := range data.Goals {
		insertGoalItem := goal.NonFinancialGoalItem
		insertGoalItem.NonFinancialBudgetID = item.ID

		createdGoal, err := UpsertNonFinancialGoal(r.Repo, insertGoalItem)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		for _, indicator := range goal.Indicators {
			indicator.GoalID = createdGoal.ID
			_, err := UpsertNonFinancialGoalIndicator(r.Repo, indicator)
			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
	}

	request, err := r.Repo.GetBudgetRequest(data.NonFinancialBudgetItem.RequestID)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	request.Status = structs.BudgetRequestFinishedStatus
	_, err = r.Repo.UpdateBudgetRequest(request)
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

func UpsertNonFinancialGoal(r repository.MicroserviceRepositoryInterface, data structs.NonFinancialGoalItem) (*structs.NonFinancialGoalItem, error) {
	if data.ID == 0 {
		item, err := r.CreateNonFinancialGoal(&data)
		if err != nil {
			return nil, err
		}

		return item, nil
	}

	item, err := r.UpdateNonFinancialGoal(data.ID, &data)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func UpsertNonFinancialGoalIndicator(r repository.MicroserviceRepositoryInterface, data structs.NonFinancialGoalIndicatorItem) (*structs.NonFinancialGoalIndicatorItem, error) {
	if data.ID == 0 {
		item, err := r.CreateNonFinancialGoalIndicator(&data)
		if err != nil {
			return nil, err
		}

		return item, nil
	}

	item, err := r.UpdateNonFinancialGoalIndicator(data.ID, &data)
	if err != nil {
		return nil, err
	}

	return item, nil
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
