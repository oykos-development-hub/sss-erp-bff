package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"context"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) BudgetOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		budget, err := r.Repo.GetBudget(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		budgetResItem, err := buildBudgetResponseItem(params.Context, r.Repo, *budget, nil)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []*dto.BudgetResponseItem{budgetResItem},
			Total:   1,
		}, nil
	}

	input := dto.GetBudgetListInputMS{}
	if budgetType, ok := params.Args["budget_type"].(int); ok && budgetType != 0 {
		input.BudgetType = &budgetType
	}
	if year, ok := params.Args["year"].(int); ok && year != 0 {
		input.Year = &year
	}
	if status, ok := params.Args["status"].(string); ok && status != "" {
		input.Status = &status
	}

	budgets, err := r.Repo.GetBudgetList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	budgetResItem, err := buildBudgetResponseItemList(params.Context, r.Repo, budgets)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   budgetResItem,
		Total:   len(budgets),
	}, nil
}

func buildBudgetStatus(ctx context.Context, budget structs.Budget) (dto.BudgetStatus, error) {
	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

	switch budget.Status {
	case structs.BudgetCreatedStatus:
		return dto.BudgetCreatedStatus, nil
	case structs.BudgetClosedStatus:
		return dto.BudgetClosedStatus, nil
	}

	if loggedInUser.RoleID == structs.UserRoleManagerOJ {
		switch budget.Status {
		case structs.BudgetSentStatus:
			return dto.ManagerBudgetSentStatus, nil
		case structs.BudgetOnReviewStatus:
			return dto.ManagerBudgetOnReviewStatus, nil
		}
	} else {
		switch budget.Status {
		case structs.BudgetSentStatus:
			return dto.OfficialBudgetSentStatus, nil
		case structs.BudgetOnReviewStatus:
			return dto.OfficialBudgetOnReviewStatus, nil
		}
	}

	return "", fmt.Errorf("budget with id: %d has incorrect status: %d", budget.ID, budget.Status)
}

func buildBudgetResponseItemList(ctx context.Context, r repository.MicroserviceRepositoryInterface, budgetList []structs.Budget) (budgetResItemList []*dto.BudgetResponseItem, err error) {
	for _, budget := range budgetList {
		budgetResponseItem, err := buildBudgetResponseItem(ctx, r, budget, nil)
		if err != nil {
			return nil, err
		}
		budgetResItemList = append(budgetResItemList, budgetResponseItem)
	}

	return
}

func buildBudgetResponseItem(ctx context.Context, r repository.MicroserviceRepositoryInterface, budget structs.Budget, organizationUnitID *int) (*dto.BudgetResponseItem, error) {
	status, err := buildBudgetStatus(ctx, budget)
	if err != nil {
		return nil, err
	}

	item := &dto.BudgetResponseItem{
		ID:         budget.ID,
		Year:       budget.Year,
		BudgetType: budget.BudgetType,
		Status:     status,
	}

	return item, nil
}

func (r *Resolver) BudgetInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data dto.CreateBudget
	response := dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
	}

	dataBytes, err := json.Marshal(params.Args["data"])
	if err != nil {
		return errors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	budgetToCreate := structs.Budget{
		Year:       data.Year,
		BudgetType: data.BudgetType,
		Status:     structs.BudgetCreatedStatus,
	}
	item, err := r.Repo.CreateBudget(&budgetToCreate)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	accountLatestVersion, err := r.Repo.GetLatestVersionOfAccounts()
	if err != nil {
		return errors.HandleAPIError(err)
	}

	financialBudget, err := r.Repo.CreateFinancialBudget(&structs.FinancialBudget{
		AccountVersion: accountLatestVersion,
		BudgetID:       item.ID,
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, limitData := range data.Limits {
		_, err = r.Repo.CreateLimitsForFinancialBudget(&structs.FinancialBudgetLimit{
			OrganizationUnitID: limitData.OrganizationUnitID,
			FinancialBudgetID:  financialBudget.ID,
			Limit:              limitData.Limit,
		})
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	resItem, err := buildBudgetResponseItem(params.Context, r.Repo, *item, nil)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	response.Item = resItem

	return response, nil
}

func (r *Resolver) BudgetSendResolver(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["id"].(int)

	budget, err := r.Repo.GetBudget(budgetID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	budget.Status = structs.BudgetSentStatus
	updatedBudget, err := r.Repo.UpdateBudget(budget)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	isParent := true
	organizationUnitList, err := r.Repo.GetOrganizationUnits(&dto.GetOrganizationUnitsInput{IsParent: &isParent})
	if err != nil {
		return errors.HandleAPIError(err)
	}
	for _, organizationUnit := range organizationUnitList.Data {
		currentFinancialRequestToCreate := &structs.BudgetRequest{
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.CurrentFinancialRequestType,
			Status:             structs.BudgetRequestCreatedStatus,
		}
		r.Repo.CreateBudgetRequest(currentFinancialRequestToCreate)

		donationFinancialRequestToCreate := &structs.BudgetRequest{
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.DonationFinancialRequestType,
			Status:             structs.BudgetRequestCreatedStatus,
		}
		r.Repo.CreateBudgetRequest(donationFinancialRequestToCreate)

		nonFinancialRequestToCreate := &structs.BudgetRequest{
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.NonFinancialRequestType,
			Status:             structs.BudgetRequestCreatedStatus,
		}
		r.Repo.CreateBudgetRequest(nonFinancialRequestToCreate)
	}

	resItem, err := buildBudgetResponseItem(params.Context, r.Repo, *updatedBudget, nil)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Message: "Budget sent successfuly",
		Status:  "success",
		Item:    resItem,
	}, nil
}

func (r *Resolver) BudgetDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteBudget(itemID)
	if err != nil {
		fmt.Printf("Deleting budget item failed because of this error - %s.\n", err)
		return fmt.Errorf("error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
