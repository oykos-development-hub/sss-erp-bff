package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"context"
	"encoding/json"
	goerrors "errors"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) buildNonFinancialBudgetDetails(ctx context.Context, request *structs.BudgetRequest) (*dto.NonFinancialBudgetResItem, error) {
	resItem := &dto.NonFinancialBudgetResItem{}

	resItem.Status = buildBudgetRequestStatus(ctx, request.Status)
	resItem.RequestID = request.ID

	activity, err := r.Repo.GetActivityByUnit(request.OrganizationUnitID)
	if err != nil {
		return nil, err
	}
	activityResItem, err := buildActivityResItem(r.Repo, *activity)
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
	}

	resItem.ActivityRequest = *activityRequest

	nonFinancialBudget, err := r.Repo.GetNonFinancialBudgetByRequestID(request.ID)
	if err != nil {
		var appErr *errors.AppError
		if goerrors.As(err, &appErr) {
			if goerrors.Is(appErr.Err, errors.ErrNonFinancialBudgetNotFound) {
				return resItem, nil
			}

			return nil, errors.Wrap(err, "GetNonFinancialBudgetDetails")
		}
	}

	resItem.ID = nonFinancialBudget.ID
	resItem.ImplContactFullName = nonFinancialBudget.ImplContactFullName
	resItem.ImplContactWorkingPlace = nonFinancialBudget.ImplContactWorkingPlace
	resItem.ImplContactPhone = nonFinancialBudget.ImplContactPhone
	resItem.ImplContactPhone = nonFinancialBudget.ImplContactPhone
	resItem.ImplContactEmail = nonFinancialBudget.ImplContactEmail
	resItem.ContactFullName = nonFinancialBudget.ContactFullName
	resItem.ContactWorkingPlace = nonFinancialBudget.ContactWorkingPlace
	resItem.ContactPhone = nonFinancialBudget.ContactPhone
	resItem.ContactEmail = nonFinancialBudget.ContactEmail

	goalList, err := r.Repo.GetNonFinancialGoalList(&dto.GetNonFinancialGoalListInputMS{
		ActivityID:           &activity.ID,
		NonFinancialBudgetID: &nonFinancialBudget.ID,
	})
	if err != nil {
		return nil, err
	}
	goalResItemList, err := buildActivityGoalRequestResItemList(r.Repo, goalList)
	if err != nil {
		return nil, err
	}

	resItem.ActivityRequest.Goals = goalResItemList

	return resItem, nil
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

	_, err = r.upsertNonFinancialBudget(data)
	if err != nil {
		return nil, err
	}

	request, err := r.Repo.GetBudgetRequest(data.RequestID)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "NonFinancialBudgetInsertResolver"))
	}

	request.Status = structs.BudgetRequestFilledStatus
	_, err = r.Repo.UpdateBudgetRequest(request)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "NonFinancialBudgetInsertResolver"))
	}

	resItem, err := r.buildNonFinancialBudgetDetails(params.Context, request)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "NonFinancialBudgetInsertResolver"))
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
