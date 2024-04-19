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
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) BudgetOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.BudgetResponseItem
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		budget, err := r.Repo.GetBudget(id)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		budgetResItem, err := buildBudgetResponseItem(params.Context, r.Repo, *budget, nil)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		if budgetResItem != nil {
			items = append(items, budgetResItem)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   items,
			Total:   len(items),
		}, nil
	}

	input := dto.GetBudgetListInputMS{}
	if budgetType, ok := params.Args["budget_type"].(int); ok && budgetType != 0 {
		input.BudgetType = &budgetType
	}
	if year, ok := params.Args["year"].(int); ok && year != 0 {
		input.Year = &year
	}
	if status, ok := params.Args["status"].(int); ok && status != 0 {
		input.Status = &status
	}

	budgets, err := r.Repo.GetBudgetList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	items, err = buildBudgetResponseItemList(params.Context, r.Repo, budgets)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   len(budgets),
	}, nil
}

func buildBudgetStatus(ctx context.Context, r repository.MicroserviceRepositoryInterface, budget structs.Budget) (*dto.DropdownSimple, error) {
	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

	switch budget.Status {
	case structs.BudgetCreatedStatus:
		return &dto.DropdownSimple{
			ID:    int(structs.BudgetCreatedStatus),
			Title: string(dto.BudgetCreatedStatus),
		}, nil
	case structs.BudgetClosedStatus:
		return &dto.DropdownSimple{
			ID:    int(structs.BudgetClosedStatus),
			Title: string(dto.BudgetClosedStatus),
		}, nil
	}

	if loggedInUser.RoleID == structs.UserRoleManagerOJ {
		status, err := buildBudgetStatusForManager(ctx, r, budget)
		if err != nil {
			return nil, err
		}
		return status, nil
	} else if budget.Status == structs.BudgetSentStatus {
		return &dto.DropdownSimple{
			ID:    int(structs.OfficialBudgetSentStatus),
			Title: string(dto.OfficialBudgetSentStatus),
		}, nil
	}

	return nil, fmt.Errorf("budget with id: %d has incorrect status: %d", budget.ID, budget.Status)
}

func buildBudgetStatusForManager(ctx context.Context, r repository.MicroserviceRepositoryInterface, budget structs.Budget) (*dto.DropdownSimple, error) {
	var status dto.DropdownSimple
	managerUnitID, ok := ctx.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || managerUnitID == nil {
		return &status, fmt.Errorf("user does not have organization unit assigned")
	}

	requests, err := r.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: managerUnitID,
		BudgetID:           budget.ID,
	})
	if err != nil {
		return &status, nil
	}

	allRequestsOnReview := true
	for _, request := range requests {
		if request.Status != structs.BudgetRequestSentOnReviewStatus {
			allRequestsOnReview = false
			break
		}
	}

	if allRequestsOnReview {
		status = dto.DropdownSimple{
			ID:    int(structs.ManagerBudgetOnReviewStatus),
			Title: string(dto.ManagerBudgetOnReviewStatus),
		}
	} else {
		status = dto.DropdownSimple{
			ID:    int(structs.ManagerBudgetProcessStatus),
			Title: string(dto.ManagerBudgetProcessStatus),
		}
	}

	return &status, nil
}

func buildBudgetResponseItemList(ctx context.Context, r repository.MicroserviceRepositoryInterface, budgetList []structs.Budget) (budgetResItemList []*dto.BudgetResponseItem, err error) {
	for _, budget := range budgetList {
		budgetResponseItem, err := buildBudgetResponseItem(ctx, r, budget, nil)
		if err != nil {
			return nil, err
		}
		if budgetResponseItem != nil {
			budgetResItemList = append(budgetResItemList, budgetResponseItem)
		}
	}

	return
}

func buildBudgetResponseItem(ctx context.Context, r repository.MicroserviceRepositoryInterface, budget structs.Budget, organizationUnitID *int) (*dto.BudgetResponseItem, error) {
	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	if loggedInUser.RoleID == structs.UserRoleManagerOJ && (budget.Status == structs.BudgetCreatedStatus || budget.Status == structs.BudgetClosedStatus) {
		return nil, nil
	}

	status, err := buildBudgetStatus(ctx, r, budget)
	if err != nil {
		return nil, err
	}

	limits, err := r.GetBudgetLimits(budget.ID)
	if err != nil {
		return nil, err
	}

	item := &dto.BudgetResponseItem{
		ID:         budget.ID,
		Year:       budget.Year,
		BudgetType: budget.BudgetType,
		Status:     *status,
		Limits:     limits,
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

	if data.ID == 0 {
		budgetToCreate := structs.Budget{
			Year:       data.Year,
			BudgetType: data.BudgetType,
			Status:     structs.BudgetCreatedStatus,
		}
		budget, err := r.Repo.CreateBudget(&budgetToCreate)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		accountLatestVersion, err := r.Repo.GetLatestVersionOfAccounts()
		if err != nil {
			return errors.HandleAPIError(err)
		}

		_, err = r.Repo.CreateFinancialBudget(&structs.FinancialBudget{
			AccountVersion: accountLatestVersion,
			BudgetID:       budget.ID,
		})
		if err != nil {
			return errors.HandleAPIError(err)
		}

		for _, limit := range data.Limits {
			limit.BudgetID = budget.ID
			_, err := r.Repo.CreateBudgetLimit(&limit)
			if err != nil {
				return errors.HandleAPIError(err)
			}
		}

		resItem, err := buildBudgetResponseItem(params.Context, r.Repo, *budget, nil)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Item = resItem

		return response, nil
	}

	budget, err := r.Repo.GetBudget(data.ID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	limits, err := r.Repo.GetBudgetLimits(budget.ID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	limitsForDelete := make(map[int]bool)
	for _, limit := range limits {
		limitsForDelete[limit.ID] = true
	}

	for _, limit := range data.Limits {
		limit.BudgetID = budget.ID
		if limit.ID != 0 {
			_, err := r.Repo.UpdateBudgetLimit(&limit)

			if err != nil {
				return errors.HandleAPIError(err)
			}
			limitsForDelete[limit.ID] = false
		} else {
			_, err := r.Repo.CreateBudgetLimit(&limit)

			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
	}

	for id, delete := range limitsForDelete {
		if delete {
			err := r.Repo.DeleteBudgetLimit(id)
			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
	}

	resItem, err := buildBudgetResponseItem(params.Context, r.Repo, *budget, nil)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	response.Item = resItem

	return response, nil
}

func (r *Resolver) BudgetSendResolver(params graphql.ResolveParams) (interface{}, error) {
	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	if loggedInUser.RoleID != structs.UserRoleOfficialForFinanceBudget && loggedInUser.RoleID != structs.UserRoleAdmin {
		return errors.HandleAPIError(fmt.Errorf("forbidden"))
	}

	budgetID := params.Args["id"].(int)

	budget, err := r.Repo.GetBudget(budgetID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	if budget.Status == structs.BudgetSentStatus {
		return dto.ResponseSingle{
			Message: "Budget is already sent",
			Status:  "success",
		}, nil
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
			Status:             structs.BudgetRequestSentStatus,
		}
		_, err := r.Repo.CreateBudgetRequest(currentFinancialRequestToCreate)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		donationFinancialRequestToCreate := &structs.BudgetRequest{
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.DonationFinancialRequestType,
			Status:             structs.BudgetRequestSentStatus,
		}
		_, err = r.Repo.CreateBudgetRequest(donationFinancialRequestToCreate)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		nonFinancialRequestToCreate := &structs.BudgetRequest{
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.NonFinancialRequestType,
			Status:             structs.BudgetRequestSentStatus,
		}
		_, err = r.Repo.CreateBudgetRequest(nonFinancialRequestToCreate)
		if err != nil {
			return errors.HandleAPIError(err)
		}
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

func (r *Resolver) BudgetSendOnReviewResolver(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)

	requests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &unitID,
		BudgetID:           budgetID,
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, request := range requests {
		if request.Status != structs.BudgetRequestFilledStatus {
			return errors.HandleAPPError(errors.NewBadRequestError("resolvers.BudgetSendOnReviewResolver: request with id %d is not filled", request.ID))
		}
	}

	for _, request := range requests {
		request.Status = structs.BudgetRequestSentOnReviewStatus
		_, err := r.Repo.UpdateBudgetRequest(&request)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	budget, err := r.Repo.GetBudget(budgetID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	resItem, err := buildBudgetResponseItem(params.Context, r.Repo, *budget, nil)
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

func (r *Resolver) BudgetDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

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

func buildBudgetRequestStatus(budgetRequests *structs.BudgetRequest) (dto.BudgetRequestStatus, error) {
	if budgetRequests.Status == structs.BudgetRequestSentStatus {
		return dto.FinancialBudgetTakeActionStatus, nil
	} else if budgetRequests.Status == structs.BudgetRequestFilledStatus || budgetRequests.Status == structs.BudgetRequestSentOnReviewStatus {
		return dto.FinancialBudgetFinishedStatus, nil
	}

	return "", fmt.Errorf("could not determine status of request budget")
}

func buildBudgetRequestResponseItem(r repository.MicroserviceRepositoryInterface, request *structs.BudgetRequest) (*dto.BudgetRequestResponseItem, error) {
	organizationUnit, err := r.GetOrganizationUnitByID(request.OrganizationUnitID)
	if err != nil {
		return nil, err
	}

	status, err := buildBudgetRequestStatus(request)
	if err != nil {
		return nil, err
	}

	item := &dto.BudgetRequestResponseItem{
		ID:               request.ID,
		OrganizationUnit: dto.DropdownSimple{ID: organizationUnit.ID, Title: organizationUnit.Title},
		BudgetID:         request.BudgetID,
		Status:           status,
		RequestType:      dto.GetRequestType(request.RequestType),
	}

	return item, nil
}

func (r *Resolver) BudgetRequestsOfficialResolver(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)

	requests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID: budgetID,
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	unitRequests := make(map[int]dto.BudgetRequestOfficialOverview)
	totalOnReview := 0

	for _, request := range requests {
		if _, exists := unitRequests[request.OrganizationUnitID]; !exists {
			var receiveDate *time.Time
			if request.Status == structs.BudgetRequestSentOnReviewStatus {
				receiveDate = &request.UpdatedAt
				totalOnReview++
			}
			unitRequests[request.OrganizationUnitID] = dto.BudgetRequestOfficialOverview{
				UnitID:      request.OrganizationUnitID,
				Status:      request.Status.StatusForOfficial(),
				ReceiveDate: receiveDate,
			}
		}
	}

	unitRequestsList := make([]dto.BudgetRequestOfficialOverview, 0, len(unitRequests))
	for _, item := range unitRequests {
		unitRequestsList = append(unitRequestsList, item)
	}

	return dto.Response{
		Message: "Budget requests",
		Status:  "success",
		Items:   unitRequestsList,
		Total:   totalOnReview,
	}, nil
}
