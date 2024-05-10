package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"context"
	"encoding/json"
	goerrors "errors"
	"fmt"
	"strconv"
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
		Total:   len(items),
	}, nil
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
	limits, err := r.GetBudgetLimits(budget.ID)
	if err != nil {
		return nil, err
	}

	var status dto.DropdownSimple

	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	if loggedInUser.RoleID == structs.UserRoleManagerOJ {
		unitID, _ := ctx.Value(config.OrganizationUnitIDKey).(*int)
		generalRequestType := structs.RequestTypeGeneral
		req, err := r.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
			OrganizationUnitID: unitID,
			BudgetID:           &budget.ID,
			RequestType:        &generalRequestType,
		})
		if err != nil {
			var appErr *errors.AppError
			if goerrors.As(err, &appErr) {
				if appErr.Code == errors.NotFoundCode {
					return nil, nil
				}

				return nil, errors.Wrap(err, "buildBudgetResponseItem")
			}
			return nil, errors.Wrap(err, "buildBudgetResponseItem")
		}
		status = buildBudgetRequestStatus(ctx, req.Status)
	} else {
		status = dto.DropdownSimple{
			ID:    int(budget.Status),
			Title: string(dto.GetBudgetStatus(budget.Status)),
		}
	}

	item := &dto.BudgetResponseItem{
		ID:         budget.ID,
		Year:       budget.Year,
		BudgetType: budget.BudgetType,
		Status:     status,
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

	isParent := true
	organizationUnitList, err := r.Repo.GetOrganizationUnits(&dto.GetOrganizationUnitsInput{IsParent: &isParent})
	if err != nil {
		return errors.HandleAPIError(err)
	}
	for _, organizationUnit := range organizationUnitList.Data {
		generalRequestToCreate := &structs.BudgetRequest{
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.RequestTypeGeneral,
			Status:             structs.BudgetRequestSentStatus,
		}
		generalRequest, err := r.Repo.CreateBudgetRequest(generalRequestToCreate)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetSendResolver: error creating general request type"))
		}

		nonFinancialRequestToCreate := &structs.BudgetRequest{
			ParentID:           &generalRequest.ID,
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.RequestTypeNonFinancial,
			Status:             structs.BudgetRequestSentStatus,
		}
		_, err = r.Repo.CreateBudgetRequest(nonFinancialRequestToCreate)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		financialRequestToCreate := &structs.BudgetRequest{
			ParentID:           &generalRequest.ID,
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.RequestTypeFinancial,
			Status:             structs.BudgetRequestSentStatus,
		}
		financialRequest, err := r.Repo.CreateBudgetRequest(financialRequestToCreate)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetSendResolver: error creating financial request type"))
		}

		currentFinancialRequestToCreate := &structs.BudgetRequest{
			ParentID:           &financialRequest.ID,
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.RequestTypeCurrentFinancial,
			Status:             structs.BudgetRequestSentStatus,
		}
		_, err = r.Repo.CreateBudgetRequest(currentFinancialRequestToCreate)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		donationFinancialRequestToCreate := &structs.BudgetRequest{
			ParentID:           &financialRequest.ID,
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.RequestTypeDonationFinancial,
			Status:             structs.BudgetRequestSentStatus,
		}
		_, err = r.Repo.CreateBudgetRequest(donationFinancialRequestToCreate)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	budget.Status = structs.BudgetSentStatus
	updatedBudget, err := r.Repo.UpdateBudget(budget)
	if err != nil {
		return errors.HandleAPIError(err)
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
	reqID := params.Args["request_id"].(int)
	req, err := r.Repo.GetBudgetRequest(reqID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	if req.RequestType != structs.RequestTypeGeneral {
		return errors.HandleAPPError(errors.NewBadRequestError(
			"request must be of type %s, provided request is of type: %s",
			dto.GetRequestType(structs.RequestTypeGeneral),
			dto.GetRequestType(req.RequestType),
		))
	}

	if req.Status != structs.BudgetRequestFilledStatus {
		return errors.HandleAPPError(errors.NewBadRequestError("budget is not filled"))
	}

	requests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &req.OrganizationUnitID,
		BudgetID:           &req.BudgetID,
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, request := range requests {
		request.Status = structs.BudgetRequestSentOnReviewStatus
		_, err := r.Repo.UpdateBudgetRequest(&request)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	return dto.ResponseSingle{
		Message: "Budget sent on review",
		Status:  "success",
	}, nil
}

func (r *Resolver) BudgetRequestRejectResolver(params graphql.ResolveParams) (interface{}, error) {
	requestID := params.Args["request_id"].(int)
	request, err := r.Repo.GetBudgetRequest(requestID)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestRejectResolver"))
	}

	if request.RequestType != structs.RequestTypeFinancial && request.RequestType != structs.RequestTypeNonFinancial {
		return errors.HandleAPPError(errors.NewBadRequestError(
			"request must be one of types: %s, %s, provided request is of type: %s",
			dto.GetRequestType(structs.RequestTypeFinancial),
			dto.GetRequestType(structs.RequestTypeNonFinancial),
			dto.GetRequestType(request.RequestType),
		))
	}

	switch request.RequestType {
	case structs.RequestTypeFinancial:
		err := r.rejectFinancialRequest(request)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestRejectResolver"))
		}
	case structs.RequestTypeNonFinancial:
		err := r.rejectNonFinancialRequest(request)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestRejectResolver"))
		}
	default:
		return errors.HandleAPPError(errors.NewInternalServerError("request type: %d with id: %d must be of type financial or non-financial", request.RequestType, request.ID))
	}

	generalRequest, err := r.Repo.GetBudgetRequest(*request.ParentID)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "BudgetRequestAcceptResolver"))
	}

	generalRequest.Status = structs.BudgetRequestRejectedStatus
	_, err = r.Repo.UpdateBudgetRequest(generalRequest)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "BudgetRequestAcceptResolver"))
	}

	return dto.ResponseSingle{
		Message: "Budget rejected and sent to managers again",
		Status:  "success",
	}, nil
}

func (r *Resolver) rejectFinancialRequest(request *structs.BudgetRequest) error {
	requests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &request.OrganizationUnitID,
		BudgetID:           &request.BudgetID,
		RequestTypes: []structs.RequestType{
			structs.RequestTypeCurrentFinancial,
			structs.RequestTypeDonationFinancial,
			structs.RequestTypeFinancial,
		},
	})
	if err != nil {
		return errors.Wrap(err, "rejectFinancialRequest: error getting financial budget requests")
	}

	for _, req := range requests {
		req.Status = structs.BudgetRequestRejectedStatus
		_, err := r.Repo.UpdateBudgetRequest(&req)
		if err != nil {
			return errors.Wrap(err, "rejectFinancialRequest: error updating financial request")
		}
	}

	return nil
}

func (r *Resolver) rejectNonFinancialRequest(request *structs.BudgetRequest) error {
	request.Status = structs.BudgetRequestRejectedStatus
	_, err := r.Repo.UpdateBudgetRequest(request)
	if err != nil {
		return errors.Wrap(err, "acceptFinancialRequest: error updating non-financial request")
	}

	return nil
}

func (r *Resolver) BudgetRequestAcceptResolver(params graphql.ResolveParams) (interface{}, error) {
	requestID := params.Args["request_id"].(int)
	request, err := r.Repo.GetBudgetRequest(requestID)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetAcceptResolver"))
	}

	if request.RequestType != structs.RequestTypeFinancial && request.RequestType != structs.RequestTypeNonFinancial {
		return errors.HandleAPPError(errors.NewBadRequestError(
			"request must be one of types: %s, %s, provided request is of type: %s",
			dto.GetRequestType(structs.RequestTypeFinancial),
			dto.GetRequestType(structs.RequestTypeNonFinancial),
			dto.GetRequestType(request.RequestType),
		))
	}

	switch request.RequestType {
	case structs.RequestTypeFinancial:
		err := r.acceptFinancialRequest(request)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestAcceptResolver"))
		}
	case structs.RequestTypeNonFinancial:
		err := r.acceptNonFinancialRequest(request)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestAcceptResolver"))
		}
	default:
		return errors.HandleAPPError(errors.NewInternalServerError("request type: %d with id: %d must be of type financial or non-financial", request.RequestType, request.ID))
	}

	siblingRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		ParentID: request.ParentID,
	})
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestAcceptResolver: error updating parent budget request"))
	}

	allAccepted := true
	for _, req := range siblingRequests {
		if req.Status != structs.BudgetRequestAcceptedStatus {
			allAccepted = false
			break
		}
	}
	if allAccepted {
		generalRequest, err := r.Repo.GetBudgetRequest(*request.ParentID)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestAcceptResolver: error getting parent financial request"))
		}
		generalRequest.Status = structs.BudgetRequestAcceptedStatus
		_, err = r.Repo.UpdateBudgetRequest(generalRequest)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestAcceptResolver: error updating parent budget request"))
		}
	}

	generalReqType := structs.RequestTypeGeneral
	generalBudgets, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID:    &request.BudgetID,
		RequestType: &generalReqType,
	})
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestAcceptResolver"))
	}
	allGeneralAccepted := true
	for _, req := range generalBudgets {
		if req.Status != structs.BudgetRequestAcceptedStatus {
			allGeneralAccepted = false
			break
		}
	}
	if allGeneralAccepted {
		budget, err := r.Repo.GetBudget(request.BudgetID)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestAcceptResolver"))
		}

		budget.Status = structs.BudgetAcceptedStatus
		_, err = r.Repo.UpdateBudget(budget)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestAcceptResolver"))
		}

		financialRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
			OrganizationUnitID: &request.OrganizationUnitID,
			BudgetID:           &request.BudgetID,
			RequestTypes: []structs.RequestType{
				structs.RequestTypeGeneral,
				structs.RequestTypeFinancial,
				structs.RequestTypeCurrentFinancial,
				structs.RequestTypeDonationFinancial,
			},
		})
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestAcceptResolver"))
		}

		for _, req := range financialRequests {
			req.Status = structs.BudgetRequestWaitingForActual
			_, err = r.Repo.UpdateBudgetRequest(&req)
			if err != nil {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestAcceptResolver"))
			}
		}
	}

	return dto.ResponseSingle{
		Message: "Budget accepted successfuly",
		Status:  "success",
	}, nil
}

func (r *Resolver) acceptFinancialRequest(request *structs.BudgetRequest) error {
	requests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &request.OrganizationUnitID,
		BudgetID:           &request.BudgetID,
		RequestTypes: []structs.RequestType{
			structs.RequestTypeCurrentFinancial,
			structs.RequestTypeDonationFinancial,
			structs.RequestTypeFinancial,
		},
	})
	if err != nil {
		return errors.Wrap(err, "acceptFinancialRequest: error getting budget requests")
	}

	for _, req := range requests {
		req.Status = structs.BudgetRequestAcceptedStatus
		_, err := r.Repo.UpdateBudgetRequest(&req)
		if err != nil {
			return errors.Wrap(err, "acceptFinancialRequest: error updating financial request")
		}
	}

	return nil
}

func (r *Resolver) acceptNonFinancialRequest(request *structs.BudgetRequest) error {
	request.Status = structs.BudgetRequestAcceptedStatus
	_, err := r.Repo.UpdateBudgetRequest(request)
	if err != nil {
		return errors.Wrap(err, "acceptFinancialRequest: error updating non-financial request")
	}

	return nil
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

func buildBudgetRequestStatus(ctx context.Context, s structs.BudgetRequestStatus) dto.DropdownSimple {
	loggedInUser := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

	if loggedInUser.RoleID == structs.UserRoleOfficialForFinanceBudget {
		return dto.DropdownSimple{
			ID:    int(s),
			Title: string(dto.RequestStatusForOfficial(s)),
		}
	}

	return dto.DropdownSimple{
		ID:    int(s),
		Title: string(dto.RequestStatusForManager(s)),
	}
}

func (r *Resolver) BudgetRequestsOfficialResolver(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)

	generalReqType := structs.RequestTypeGeneral
	requests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID:    &budgetID,
		RequestType: &generalReqType,
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	unitRequestsList := make([]dto.BudgetRequestOfficialOverview, 0, len(requests))
	totalOnReview := 0

	for _, request := range requests {
		var receiveDate *time.Time
		if request.Status == structs.BudgetRequestSentOnReviewStatus {
			receiveDate = &request.UpdatedAt
			totalOnReview++
		}
		unit, err := r.Repo.GetOrganizationUnitByID(request.OrganizationUnitID)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestsOfficialResolver: unit not found"))
		}
		unitRequestsList = append(unitRequestsList, dto.BudgetRequestOfficialOverview{
			Unit: dto.DropdownOUSimple{
				ID:    unit.ID,
				Title: unit.Title,
			},
			Status:      string(dto.RequestStatusForOfficial(request.Status)),
			ReceiveDate: receiveDate,
		})
	}

	return dto.Response{
		Message: "Budget requests",
		Status:  "success",
		Items:   unitRequestsList,
		Total:   totalOnReview,
	}, nil
}

func (r *Resolver) BudgetRequestsDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)

	financialDetails, err := r.GetFinancialBudgetDetails(params.Context, budgetID, unitID, false)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting financial details"))
	}

	nonFinancialRequestType := structs.RequestTypeNonFinancial
	nonFinancialRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &unitID,
		BudgetID:           &budgetID,
		RequestType:        &nonFinancialRequestType,
	})
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting non financial request"))
	}

	nonFinancialDetails, err := r.buildNonFinancialBudgetDetails(params.Context, nonFinancialRequest)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "Error getting non financial data"))
	}

	generalRequest, err := r.Repo.GetBudgetRequest(*nonFinancialRequest.ParentID)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "Error getting general req data"))
	}

	budget, err := r.Repo.GetBudget(budgetID)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "Error getting budget data"))
	}

	return dto.ResponseSingle{
		Message: "Budget requests",
		Status:  "success",
		Item: &dto.BudgetRequestsDetails{
			Budget: dto.DropdownSimple{
				ID:    budget.ID,
				Title: strconv.Itoa(budget.Year),
			},
			Status:                    buildBudgetRequestStatus(params.Context, generalRequest.Status),
			RequestID:                 generalRequest.ID,
			FinancialBudgetDetails:    financialDetails,
			NonFinancialBudgetDetails: nonFinancialDetails,
		},
	}, nil
}

func (r *Resolver) FinancialBudgetSummary(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)

	financialDetails, err := r.GetFinancialBudgetDetails(params.Context, budgetID, unitID, unitID == 0)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "Error getting financial details"))
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the data you asked for!",
		Item:    financialDetails,
	}, nil
}
