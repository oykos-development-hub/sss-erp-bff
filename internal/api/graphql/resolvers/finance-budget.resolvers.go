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
	"slices"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/shopspring/decimal"
)

func (r *Resolver) BudgetOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.BudgetResponseItem
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		budget, err := r.Repo.GetBudget(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		budgetResItem, err := buildBudgetResponseItem(params.Context, r.Repo, *budget)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	items, err = buildBudgetResponseItemList(params.Context, r.Repo, budgets)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		budgetResponseItem, err := buildBudgetResponseItem(ctx, r, budget)
		if err != nil {
			return nil, errors.Wrap(err, "build budget response item")
		}
		if budgetResponseItem != nil {
			budgetResItemList = append(budgetResItemList, budgetResponseItem)
		}
	}

	return
}

func buildBudgetResponseItem(ctx context.Context, r repository.MicroserviceRepositoryInterface, budget structs.Budget) (*dto.BudgetResponseItem, error) {
	limits, err := r.GetBudgetLimits(budget.ID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get budget limits")
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

				return nil, errors.Wrap(err, "repo get one budget request")
			}
			return nil, errors.Wrap(err, "repo get one budget request")
		}
		status = buildBudgetRequestStatus(ctx, req.Status)
	} else {
		status = dto.DropdownSimple{
			ID:    int(budget.Status),
			Title: string(dto.GetBudgetStatus(budget.Status)),
		}
	}

	generalRequestType := structs.RequestTypeGeneral
	requests, err := r.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID:    &budget.ID,
		RequestType: &generalRequestType,
		Statuses: []structs.BudgetRequestStatus{
			structs.BudgetRequestSentOnReviewStatus,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "get budget request list")
	}

	item := &dto.BudgetResponseItem{
		ID:                          budget.ID,
		Year:                        budget.Year,
		BudgetType:                  budget.BudgetType,
		Status:                      status,
		Limits:                      limits,
		NumberOfRequestsForOfficial: len(requests),
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if data.ID == 0 {
		budgetToCreate := structs.Budget{
			Year:       data.Year,
			BudgetType: data.BudgetType,
			Status:     structs.BudgetCreatedStatus,
		}
		budget, err := r.Repo.CreateBudget(params.Context, &budgetToCreate)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		accountLatestVersion, err := r.Repo.GetLatestVersionOfAccounts()
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		_, err = r.Repo.CreateFinancialBudget(params.Context, &structs.FinancialBudget{
			AccountVersion: accountLatestVersion,
			BudgetID:       budget.ID,
		})
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, limit := range data.Limits {
			limit.BudgetID = budget.ID
			_, err := r.Repo.CreateBudgetLimit(&limit)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}

		resItem, err := buildBudgetResponseItem(params.Context, r.Repo, *budget)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Item = resItem

		return response, nil
	}

	budget, err := r.Repo.GetBudget(data.ID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	limits, err := r.Repo.GetBudgetLimits(budget.ID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			limitsForDelete[limit.ID] = false
		} else {
			_, err := r.Repo.CreateBudgetLimit(&limit)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}
	}

	for id, delete := range limitsForDelete {
		if delete {
			err := r.Repo.DeleteBudgetLimit(id)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}
	}

	resItem, err := buildBudgetResponseItem(params.Context, r.Repo, *budget)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Item = resItem

	return response, nil
}

func (r *Resolver) BudgetSendResolver(params graphql.ResolveParams) (interface{}, error) {
	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	if loggedInUser.RoleID != structs.UserRoleOfficialForFinanceBudget && loggedInUser.RoleID != structs.UserRoleAdmin {
		return errors.HandleAPPError(fmt.Errorf("forbidden"))
	}

	budgetID := params.Args["id"].(int)

	budget, err := r.Repo.GetBudget(budgetID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	for _, organizationUnit := range organizationUnitList.Data {
		generalRequestToCreate := &structs.BudgetRequest{
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.RequestTypeGeneral,
			Status:             structs.BudgetRequestSentStatus,
		}
		generalRequest, err := r.Repo.CreateBudgetRequest(params.Context, generalRequestToCreate)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		nonFinancialRequestToCreate := &structs.BudgetRequest{
			ParentID:           &generalRequest.ID,
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.RequestTypeNonFinancial,
			Status:             structs.BudgetRequestSentStatus,
		}
		_, err = r.Repo.CreateBudgetRequest(params.Context, nonFinancialRequestToCreate)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		financialRequestToCreate := &structs.BudgetRequest{
			ParentID:           &generalRequest.ID,
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.RequestTypeFinancial,
			Status:             structs.BudgetRequestSentStatus,
		}
		financialRequest, err := r.Repo.CreateBudgetRequest(params.Context, financialRequestToCreate)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		currentFinancialRequestToCreate := &structs.BudgetRequest{
			ParentID:           &financialRequest.ID,
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.RequestTypeCurrentFinancial,
			Status:             structs.BudgetRequestSentStatus,
		}
		_, err = r.Repo.CreateBudgetRequest(params.Context, currentFinancialRequestToCreate)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		donationFinancialRequestToCreate := &structs.BudgetRequest{
			ParentID:           &financialRequest.ID,
			OrganizationUnitID: organizationUnit.ID,
			BudgetID:           budgetID,
			RequestType:        structs.RequestTypeDonationFinancial,
			Status:             structs.BudgetRequestSentStatus,
		}
		_, err = r.Repo.CreateBudgetRequest(params.Context, donationFinancialRequestToCreate)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	budget.Status = structs.BudgetSentStatus
	updatedBudget, err := r.Repo.UpdateBudget(params.Context, budget)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	resItem, err := buildBudgetResponseItem(params.Context, r.Repo, *updatedBudget)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, request := range requests {
		request.Status = structs.BudgetRequestSentOnReviewStatus
		_, err := r.Repo.UpdateBudgetRequest(params.Context, &request)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if request.RequestType != structs.RequestTypeFinancial && request.RequestType != structs.RequestTypeNonFinancial {
		return errors.HandleAPPError(errors.NewBadRequestError(
			"request must be one of types: %s, %s, provided request is of type: %s",
			dto.GetRequestType(structs.RequestTypeFinancial),
			dto.GetRequestType(structs.RequestTypeNonFinancial),
			dto.GetRequestType(request.RequestType),
		))
	}

	request.Comment = params.Args["comment"].(string)
	_, err = r.Repo.UpdateBudgetRequest(params.Context, request)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	switch request.RequestType {
	case structs.RequestTypeFinancial:
		err := r.rejectFinancialRequest(params.Context, request)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestRejectResolver"))
		}
	case structs.RequestTypeNonFinancial:
		err := r.rejectNonFinancialRequest(params.Context, request)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "BudgetRequestRejectResolver"))
		}
	default:
		return errors.HandleAPPError(errors.NewInternalServerError("request type: %d with id: %d must be of type financial or non-financial", request.RequestType, request.ID))
	}

	generalRequest, err := r.Repo.GetBudgetRequest(*request.ParentID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	generalRequest.Status = structs.BudgetRequestRejectedStatus
	_, err = r.Repo.UpdateBudgetRequest(params.Context, generalRequest)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Message: "Budget rejected and sent to managers again",
		Status:  "success",
	}, nil
}

func (r *Resolver) rejectFinancialRequest(ctx context.Context, request *structs.BudgetRequest) error {
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
		return errors.Wrap(err, "repo get budget request list")
	}

	for _, req := range requests {
		req.Status = structs.BudgetRequestRejectedStatus
		_, err := r.Repo.UpdateBudgetRequest(ctx, &req)
		if err != nil {
			return errors.Wrap(err, "repo update budget request")
		}
	}

	reqTypeNonFinancial := structs.RequestTypeNonFinancial
	nonFinancialRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &request.OrganizationUnitID,
		BudgetID:           &request.BudgetID,
		RequestType:        &reqTypeNonFinancial,
	})
	if err != nil {
		return errors.Wrap(err, "repo get one budget request ")
	}

	if nonFinancialRequest.Status != structs.BudgetRequestRejectedStatus {
		nonFinancialRequest.Status = structs.BudgetRequestFilledStatus
		_, err = r.Repo.UpdateBudgetRequest(ctx, nonFinancialRequest)
		if err != nil {
			return errors.Wrap(err, "repo update budget request")
		}
	}

	return nil
}

func (r *Resolver) rejectNonFinancialRequest(ctx context.Context, request *structs.BudgetRequest) error {
	request.Status = structs.BudgetRequestRejectedStatus
	_, err := r.Repo.UpdateBudgetRequest(ctx, request)
	if err != nil {
		return errors.Wrap(err, "repo update budget request")
	}

	financialRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &request.OrganizationUnitID,
		BudgetID:           &request.BudgetID,
		RequestTypes: []structs.RequestType{
			structs.RequestTypeCurrentFinancial,
			structs.RequestTypeDonationFinancial,
			structs.RequestTypeFinancial,
		},
	})
	if err != nil {
		return errors.Wrap(err, "repo get budget request list")
	}

	for _, req := range financialRequests {
		if req.Status != structs.BudgetRequestRejectedStatus {
			req.Status = structs.BudgetRequestFilledStatus
			_, err := r.Repo.UpdateBudgetRequest(ctx, &req)
			if err != nil {
				return errors.Wrap(err, "repo update budget request")
			}
		}
	}

	return nil
}

func (r *Resolver) BudgetRequestAcceptResolver(params graphql.ResolveParams) (interface{}, error) {
	requestID := params.Args["request_id"].(int)
	request, err := r.Repo.GetBudgetRequest(requestID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		err := r.acceptFinancialRequest(params.Context, request)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	case structs.RequestTypeNonFinancial:
		err := r.acceptNonFinancialRequest(params.Context, request)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	default:
		return errors.HandleAPPError(errors.NewInternalServerError("request type: %d with id: %d must be of type financial or non-financial", request.RequestType, request.ID))
	}

	siblingRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		ParentID: request.ParentID,
	})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		generalRequest.Status = structs.BudgetRequestAcceptedStatus

		_, err = r.Repo.UpdateBudgetRequest(params.Context, generalRequest)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	generalReqType := structs.RequestTypeGeneral
	generalBudgets, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID:    &request.BudgetID,
		RequestType: &generalReqType,
	})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		budget.Status = structs.BudgetAcceptedStatus
		_, err = r.Repo.UpdateBudget(params.Context, budget)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, req := range financialRequests {
			req.Status = structs.BudgetRequestWaitingForActual
			_, err = r.Repo.UpdateBudgetRequest(params.Context, &req)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}
	}

	return dto.ResponseSingle{
		Message: "Budget accepted successfuly",
		Status:  "success",
	}, nil
}

func (r *Resolver) acceptFinancialRequest(ctx context.Context, request *structs.BudgetRequest) error {
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
		return errors.Wrap(err, "repo get budget request list")
	}

	for _, req := range requests {
		req.Status = structs.BudgetRequestAcceptedStatus

		if req.RequestType == structs.RequestTypeFinancial {
			req.Comment = ""
		}

		_, err := r.Repo.UpdateBudgetRequest(ctx, &req)
		if err != nil {
			return errors.Wrap(err, "repo update budget request")
		}
	}

	return nil
}

func (r *Resolver) acceptNonFinancialRequest(ctx context.Context, request *structs.BudgetRequest) error {
	request.Status = structs.BudgetRequestAcceptedStatus
	request.Comment = ""
	_, err := r.Repo.UpdateBudgetRequest(ctx, request)
	if err != nil {
		return errors.Wrap(err, "repo update budget request")
	}

	return nil
}

func (r *Resolver) BudgetDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteBudget(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	budgetResItem, err := buildBudgetResponseItem(params.Context, r.Repo, *budget)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
	financialReqType := structs.RequestTypeCurrentFinancial

	requests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID:    &budgetID,
		RequestType: &generalReqType,
	})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	unitRequestsList := make([]dto.BudgetRequestOfficialOverview, 0, len(requests))
	totalOnReview := 0

	limits, err := r.Repo.GetBudgetLimits(budgetID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	mapLimits := make(map[int]decimal.Decimal)

	for _, limit := range limits {
		mapLimits[limit.OrganizationUnitID] = decimal.NewFromInt(int64(limit.Limit))
	}

	for _, request := range requests {
		var receiveDate *time.Time

		if request.Status == structs.BudgetRequestSentOnReviewStatus {
			receiveDate = &request.UpdatedAt
			totalOnReview++
		}

		unit, err := r.Repo.GetOrganizationUnitByID(request.OrganizationUnitID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		resItem := dto.BudgetRequestOfficialOverview{
			Unit: dto.DropdownOUSimple{
				ID:    unit.ID,
				Title: unit.Title,
			},
			Status:      string(dto.RequestStatusForOfficial(request.Status)),
			ReceiveDate: receiveDate,
		}

		if slices.Contains([]structs.BudgetRequestStatus{
			structs.BudgetRequestSentOnReviewStatus,
			structs.BudgetRequestAcceptedStatus,
			structs.BudgetRequestWaitingForActual,
			structs.BudgetRequestCompletedActualStatus,
		}, request.Status) {
			financialRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
				BudgetID:           &budgetID,
				OrganizationUnitID: &request.OrganizationUnitID,
				RequestType:        &financialReqType,
			})
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			financialBudget, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			topMostAccounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{
				Version: &financialBudget.AccountVersion,
				TopMost: true,
			})
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			topMostAccountIDList := make([]int, 0, topMostAccounts.Total)
			for _, account := range topMostAccounts.Data {
				topMostAccountIDList = append(topMostAccountIDList, account.ID)
			}

			filledData, err := r.Repo.GetFilledFinancialBudgetList(&dto.FilledFinancialBudgetInputMS{
				BudgetRequestID: financialRequest.ID,
				Accounts:        topMostAccountIDList,
			})
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			total := decimal.Zero
			for _, data := range filledData {
				total = total.Add(data.CurrentYear)
			}

			resItem.Total = total
		}

		resItem.Limit = mapLimits[resItem.Unit.ID]

		unitRequestsList = append(unitRequestsList, resItem)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	nonFinancialRequestType := structs.RequestTypeNonFinancial
	nonFinancialRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &unitID,
		BudgetID:           &budgetID,
		RequestType:        &nonFinancialRequestType,
	})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	nonFinancialDetails, err := r.buildNonFinancialBudgetDetails(params.Context, nonFinancialRequest)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	generalRequest, err := r.Repo.GetBudgetRequest(*nonFinancialRequest.ParentID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	budget, err := r.Repo.GetBudget(budgetID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	limit, err := r.Repo.GetBudgetUnitLimit(budgetID, unitID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Message: "Budget requests",
		Status:  "success",
		Item: &dto.BudgetRequestsDetails{
			Limit: limit,
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

func (r *Resolver) NonFinancialBudgetOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	unitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok {
		return errors.HandleAPPError(errors.NewBadRequestError("error getting logged in unit"))
	}

	year := params.Args["year"].(int)

	//TODO: after planning budget is done on FE, add status filter Done
	filter := dto.GetBudgetListInputMS{}
	if year != 0 {
		filter.Year = &year
	}
	budgets, err := r.Repo.GetBudgetList(&filter)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "repo get budget list"))
	}

	var nonFinancialData []dto.NonFinancialBudgetResItem

	for _, budget := range budgets {
		nonFinancialRequestType := structs.RequestTypeNonFinancial
		nonFinancialRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
			OrganizationUnitID: unitID,
			BudgetID:           &budget.ID,
			RequestType:        &nonFinancialRequestType,
		})
		if err != nil {
			if errors.IsErr(err, errors.NotFoundCode) {
				continue
			}

			return errors.HandleAPPError(errors.WrapInternalServerError(err, "get one budget request"))
		}

		nonFinancialDetails, err := r.buildNonFinancialBudgetDetails(params.Context, nonFinancialRequest)
		if err != nil {
			return errors.HandleAPPError(errors.Wrap(err, "build non financial budget details"))
		}

		nonFinancialDetails.Year = budget.Year

		nonFinancialData = append(nonFinancialData, *nonFinancialDetails)
	}

	return dto.Response{
		Message: "Budget requests",
		Status:  "success",
		Items:   nonFinancialData,
	}, nil
}

func (r *Resolver) FinancialBudgetSummary(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["unit_id"].(int)

	financialDetails, err := r.GetFinancialBudgetDetails(params.Context, budgetID, unitID, unitID == 0)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the data you asked for!",
		Item:    financialDetails,
	}, nil
}

func (r *Resolver) CurrentBudgetOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	organizationUnitID, _ := params.Args["organization_unit_id"].(int)

	var items []*dto.CurrentBudgetAccounts
	var donationItems []*dto.CurrentBudgetAccounts

	currentBudgetItems, err := r.Repo.GetCurrentBudgetByOrganizationUnit(organizationUnitID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	accountVersion := 0

	if len(currentBudgetItems) > 0 && currentBudgetItems[0].AccountID != 0 {
		accountID, err := r.Repo.GetAccountItemByID(currentBudgetItems[0].AccountID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		accountVersion = accountID.Version
	}

	if accountVersion != 0 {
		items, err = buildCurrentBudget(r, accountVersion, currentBudgetItems)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		donationItems, err = buildDonationBudget(r, accountVersion, currentBudgetItems)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	response := dto.CurrentBudgetAccountsResponse{
		CurrentAccounts:  items,
		DonationAccounts: donationItems,
		Version:          accountVersion,
	}

	if len(currentBudgetItems) > 0 && currentBudgetItems[0].BudgetID != 0 {
		response.BudgetID = currentBudgetItems[0].BudgetID
	}

	unitIDList, err := r.Repo.GetCurrentBudgetUnitList(params.Context)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	unitsList := make([]dto.DropdownOUSimple, len(unitIDList))
	for i, unitID := range unitIDList {
		unit, err := r.Repo.GetOrganizationUnitByID(unitID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		unitsList[i] = dto.DropdownOUSimple{
			ID:    unit.ID,
			Title: unit.Title,
		}
	}

	response.Units = unitsList

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

func buildCurrentBudget(r *Resolver, accountVersion int, currentBudgetItems []structs.CurrentBudget) ([]*dto.CurrentBudgetAccounts, error) {

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{
		Version: &accountVersion,
	})

	if err != nil {
		return nil, errors.Wrap(err, "repo get account items")
	}

	accountMap := make(map[int]*dto.CurrentBudgetAccounts)

	// Inicijalizacija mape konta
	for _, account := range accounts.Data {
		accountMap[account.ID] = &dto.CurrentBudgetAccounts{
			ID:           account.ID,
			Title:        account.Title,
			ParentID:     account.ParentID,
			SerialNumber: account.SerialNumber,
			Children:     []*dto.CurrentBudgetAccounts{},
		}
	}

	// Popunjavanje mape
	for _, budget := range currentBudgetItems {
		if accountNode, exists := accountMap[budget.AccountID]; exists && budget.Type == 1 {
			accountNode.FilledFinanceBudget = dto.CurrentBudgetResponse{
				InititalActual: budget.InitialActual,
				Actual:         budget.Actual,
				Balance:        budget.Balance,
				CurrentAmount:  budget.CurrentAmount,
				BudgetID:       budget.BudgetID,
			}
		}
	}

	// Kreiranje stabla
	var rootNodes []*dto.CurrentBudgetAccounts
	for _, account := range accounts.Data {
		if account.ParentID == nil {
			rootNodes = append(rootNodes, accountMap[account.ID])
		} else if parentAccount, exists := accountMap[*account.ParentID]; exists {
			parentAccount.Children = append(parentAccount.Children, accountMap[account.ID])
		}
	}

	return rootNodes, nil
}

func buildDonationBudget(r *Resolver, accountVersion int, currentBudgetItems []structs.CurrentBudget) ([]*dto.CurrentBudgetAccounts, error) {

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{
		Version: &accountVersion,
	})

	if err != nil {
		return nil, errors.Wrap(err, "repo get account items")
	}

	accountMap := make(map[int]*dto.CurrentBudgetAccounts)

	// Inicijalizacija mape konta
	for _, account := range accounts.Data {
		accountMap[account.ID] = &dto.CurrentBudgetAccounts{
			ID:           account.ID,
			Title:        account.Title,
			ParentID:     account.ParentID,
			SerialNumber: account.SerialNumber,
			Children:     []*dto.CurrentBudgetAccounts{},
		}
	}

	// Popunjavanje mape
	for _, budget := range currentBudgetItems {
		if accountNode, exists := accountMap[budget.AccountID]; exists && budget.Type == 2 {
			accountNode.FilledFinanceBudget = dto.CurrentBudgetResponse{
				InititalActual: budget.InitialActual,
				Actual:         budget.Actual,
				Balance:        budget.Balance,
				CurrentAmount:  budget.CurrentAmount,
				BudgetID:       budget.BudgetID,
			}
		}
	}

	// Kreiranje stabla
	var rootNodes []*dto.CurrentBudgetAccounts
	for _, account := range accounts.Data {
		if account.ParentID == nil {
			rootNodes = append(rootNodes, accountMap[account.ID])
		} else if parentAccount, exists := accountMap[*account.ParentID]; exists {
			parentAccount.Children = append(parentAccount.Children, accountMap[account.ID])
		}
	}

	return rootNodes, nil
}
