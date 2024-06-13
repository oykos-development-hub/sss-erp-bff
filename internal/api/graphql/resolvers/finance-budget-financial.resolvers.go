package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"context"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) FinancialBudgetDetails(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)

	financialBudget, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	latestVersion, err := r.Repo.GetLatestVersionOfAccounts()
	if err != nil {
		return errors.HandleAPIError(err)
	}

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{Version: &financialBudget.AccountVersion})
	if err != nil {
		return errors.HandleAPIError(err)
	}
	accountResItemlist, err := buildAccountItemResponseItemList(accounts.Data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item: dto.FinancialBudgetDetails{
			Version:       financialBudget.AccountVersion,
			LatestVersion: latestVersion,
			Accounts:      accountResItemlist,
		},
	}, nil
}

func (r *Resolver) FinancialBudgetOverview(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	unitID := params.Args["organization_unit_id"].(int)

	financialBudgetOveriew, err := r.GetFinancialBudgetDetails(params.Context, budgetID, unitID, false)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "FinancialBudgetOverview"))
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    financialBudgetOveriew,
	}, nil
}

func (r *Resolver) GetFinancialBudgetDetails(ctx context.Context, budgetID, unitID int, summary bool) (*dto.FinancialBudgetOverviewResponse, error) {
	var response dto.FinancialBudgetOverviewResponse

	var filledAccounts []structs.FilledFinanceBudget
	var donationFinancialBudgetRequest *structs.BudgetRequest
	var currentFinancialBudgetRequest *structs.BudgetRequest

	financialBudget, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails: error getting financial budget")
	}
	response.AccountVersion = financialBudget.AccountVersion

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{Version: &financialBudget.AccountVersion})
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails: error getting account items")
	}

	currentFinancialType := structs.RequestTypeCurrentFinancial
	if summary {
		filledAccounts, err = r.Repo.GetFinancialFilledSummary(budgetID, currentFinancialType)
		if err != nil {
			return nil, errors.Wrap(err, "GetFinancialBudgetDetails: error getting current financial budget request")
		}
	} else {
		currentFinancialBudgetRequest, err = r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
			OrganizationUnitID: &unitID,
			BudgetID:           &budgetID,
			RequestType:        &currentFinancialType,
		})
		if err != nil {
			return nil, errors.Wrap(err, "GetFinancialBudgetDetails: error getting current financial budget request")
		}
		filledAccounts, err = r.Repo.GetFilledFinancialBudgetList(&dto.FilledFinancialBudgetInputMS{BudgetRequestID: currentFinancialBudgetRequest.ID})
		if err != nil {
			return nil, errors.Wrap(err, "buildFilledRequestData")
		}
	}

	currentFinancialRequestResList, err := buildFilledRequestData(accounts.Data, filledAccounts)
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails")
	}

	donationFinancialType := structs.RequestTypeDonationFinancial
	if summary {
		filledAccounts, err = r.Repo.GetFinancialFilledSummary(budgetID, donationFinancialType)
		if err != nil {
			return nil, errors.Wrap(err, "GetFinancialBudgetDetails: error getting current financial budget request")
		}
	} else {
		donationFinancialBudgetRequest, err = r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
			OrganizationUnitID: &unitID,
			BudgetID:           &budgetID,
			RequestType:        &donationFinancialType,
		})
		if err != nil {
			return nil, errors.Wrap(err, "GetFinancialBudgetDetails")
		}

		filledAccounts, err = r.Repo.GetFilledFinancialBudgetList(&dto.FilledFinancialBudgetInputMS{BudgetRequestID: donationFinancialBudgetRequest.ID})
		if err != nil {
			return nil, errors.Wrap(err, "buildFilledRequestData")
		}
	}
	donationFinancialRequestResList, err := buildFilledRequestData(accounts.Data, filledAccounts)
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails")
	}

	financialBudgetOveriew := &dto.FinancialBudgetOverviewResponse{
		AccountVersion:                 financialBudget.AccountVersion,
		CurrentAccountsWithFilledData:  currentFinancialRequestResList.CreateTree(),
		DonationAccountsWithFilledData: donationFinancialRequestResList.CreateTree(),
	}

	if !summary {
		financialRequestType := structs.RequestTypeFinancial
		financialRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
			OrganizationUnitID: &unitID,
			BudgetID:           &budgetID,
			RequestType:        &financialRequestType,
		})
		if err != nil {
			return nil, errors.Wrap(err, "GetFinancialBudgetDetails")
		}

		financialBudgetOveriew.RequestID = financialRequest.ID
		financialBudgetOveriew.OfficialComment = financialRequest.Comment
		financialBudgetOveriew.Status = buildBudgetRequestStatus(ctx, financialRequest.Status)

		financialBudgetOveriew.CurrentRequestID = currentFinancialBudgetRequest.ID
		financialBudgetOveriew.CurrentBudgetComment = currentFinancialBudgetRequest.Comment
		financialBudgetOveriew.CurrentBudgetStatus = buildBudgetRequestStatus(ctx, currentFinancialBudgetRequest.Status)

		financialBudgetOveriew.DonationRequestID = donationFinancialBudgetRequest.ID
		financialBudgetOveriew.CurrentBudgetComment = donationFinancialBudgetRequest.Comment
		financialBudgetOveriew.DonationBudgetStatus = buildBudgetRequestStatus(ctx, donationFinancialBudgetRequest.Status)

	}

	return financialBudgetOveriew, nil
}

func buildFilledRequestData(accounts []*structs.AccountItem, filledAccounts []structs.FilledFinanceBudget) (dto.AccountWithFilledFinanceBudgetResponseList, error) {
	var filledAccountResItemList dto.AccountWithFilledFinanceBudgetResponseList

	for _, account := range accounts {
		filledAccountItem, err := buildAccountWithFilledFinanceResponseItem(account, filledAccounts)
		if err != nil {
			return filledAccountResItemList, errors.Wrap(err, "buildFilledRequestData")
		}
		filledAccountResItemList.Data = append(filledAccountResItemList.Data, filledAccountItem)
	}

	return filledAccountResItemList, nil
}

func buildAccountWithFilledFinanceResponseItem(item *structs.AccountItem, filledAccounts []structs.FilledFinanceBudget) (*dto.AccountWithFilledFinanceBudget, error) {
	resItem := &dto.AccountWithFilledFinanceBudget{
		ID:           item.ID,
		Title:        item.Title,
		ParentID:     item.ParentID,
		SerialNumber: item.SerialNumber,
	}

	for _, filledAccount := range filledAccounts {
		if filledAccount.AccountID == item.ID {
			resItem.FilledFinanceBudget = &filledAccount
			break
		}
	}

	return resItem, nil
}

func (r *Resolver) FinancialBudgetVersionUpdate(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)

	latestVersion, err := r.Repo.GetLatestVersionOfAccounts()
	if err != nil {
		return errors.HandleAPIError(err)
	}

	financialBudget, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	financialBudget.AccountVersion = latestVersion

	_, err = r.Repo.UpdateFinancialBudget(params.Context, financialBudget)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	financialBudgetRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID:     &budgetID,
		RequestTypes: []structs.RequestType{structs.RequestTypeCurrentFinancial, structs.RequestTypeDonationFinancial},
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, request := range financialBudgetRequests {
		filledRequestData, err := r.Repo.GetFilledFinancialBudgetList(&dto.FilledFinancialBudgetInputMS{BudgetRequestID: request.ID})
		if err != nil {
			return errors.HandleAPIError(err)
		}

		filledDataForDelete := make(map[int]bool)
		for _, filledData := range filledRequestData {
			filledDataForDelete[filledData.ID] = true
		}

		for _, filledData := range filledRequestData {
			oldAccount, err := r.Repo.GetAccountItemByID(filledData.AccountID)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			newAccount, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{
				Version:      &latestVersion,
				SerialNumber: &oldAccount.SerialNumber,
				Title:        &oldAccount.Title,
			})
			if err != nil {
				return errors.HandleAPIError(err)
			}

			if len(newAccount.Data) > 0 {
				_, err := r.Repo.UpdateFilledFinancialBudget(params.Context, filledData.ID, &structs.FilledFinanceBudget{
					BudgetRequestID: filledData.BudgetRequestID,
					AccountID:       newAccount.Data[0].ID,
					CurrentYear:     filledData.CurrentYear,
					NextYear:        filledData.NextYear,
					YearAfterNext:   filledData.YearAfterNext,
					Description:     filledData.Description,
				})
				if err != nil {
					return errors.HandleAPIError(err)
				}
				filledDataForDelete[filledData.ID] = false
			}
		}
		for id, delete := range filledDataForDelete {
			if delete {
				err := r.Repo.DeleteFilledFinancialBudgetData(params.Context, id)
				if err != nil {
					return errors.HandleAPIError(err)
				}
			}
		}
	}

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{Version: &financialBudget.AccountVersion})
	if err != nil {
		return errors.HandleAPIError(err)
	}
	accountResItemlist, err := buildAccountItemResponseItemList(accounts.Data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item: dto.FinancialBudgetDetails{
			Version:       financialBudget.AccountVersion,
			LatestVersion: latestVersion,
			Accounts:      accountResItemlist,
		},
	}, nil
}

func (r *Resolver) FinancialBudgetFillResolver(params graphql.ResolveParams) (interface{}, error) {
	requestID := params.Args["request_id"].(int)

	request, err := r.Repo.GetBudgetRequest(requestID)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	if request.RequestType != structs.RequestTypeCurrentFinancial && request.RequestType != structs.RequestTypeDonationFinancial {
		return errors.HandleAPPError(errors.NewBadRequestError(
			"request must be one of types: %s, %s, provided request is of type: %s",
			dto.GetRequestType(structs.RequestTypeCurrentFinancial),
			dto.GetRequestType(structs.RequestTypeDonationFinancial),
			dto.GetRequestType(request.RequestType),
		))
	}

	var items []structs.FilledFinanceBudget

	dataBytes, _ := json.Marshal(params.Args["data"])
	err = json.Unmarshal(dataBytes, &items)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	var resItemList []dto.FilledFinancialBudgetResItem

	for _, data := range items {
		itemID := data.ID
		data.BudgetRequestID = requestID
		if itemID != 0 {
			item, err := r.Repo.UpdateFilledFinancialBudget(params.Context, itemID, &data)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			resItem, err := buildFilledFinancialBudgetResItem(r.Repo, *item)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			resItemList = append(resItemList, *resItem)
		} else {
			item, err := r.Repo.FillFinancialBudget(params.Context, &data)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			resItem, err := buildFilledFinancialBudgetResItem(r.Repo, *item)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			resItemList = append(resItemList, *resItem)
		}
	}

	request.Status = structs.BudgetRequestFilledStatus
	request.Comment = params.Args["comment"].(string)
	_, err = r.Repo.UpdateBudgetRequest(params.Context, request)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	//update parent financial status to sent if all financial sub requests (current, donation) are filled
	financialSubRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		ParentID: request.ParentID,
	})
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillResolver: error getting financial child requests"))
	}

	allFinancialFilled := true
	for _, financialChildRequest := range financialSubRequests {
		if financialChildRequest.Status != structs.BudgetRequestFilledStatus {
			allFinancialFilled = false
			break
		}
	}

	if allFinancialFilled {
		financialRequest, err := r.Repo.GetBudgetRequest(*request.ParentID)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillResolver: error getting parent financial request"))
		}
		financialRequest.Status = structs.BudgetRequestFilledStatus
		_, err = r.Repo.UpdateBudgetRequest(params.Context, financialRequest)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillResolver: error updating parent budget request"))
		}

		financialAndNonFinancialRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
			ParentID: financialRequest.ParentID,
		})
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillResolver: error updating parent budget request"))
		}

		allFilled := true
		for _, financialChildRequest := range financialAndNonFinancialRequests {
			if financialChildRequest.Status != structs.BudgetRequestFilledStatus {
				allFilled = false
				break
			}
		}
		if allFilled {
			generalRequest, err := r.Repo.GetBudgetRequest(*financialRequest.ParentID)
			if err != nil {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillResolver: error getting parent financial request"))
			}
			generalRequest.Status = structs.BudgetRequestFilledStatus
			_, err = r.Repo.UpdateBudgetRequest(params.Context, generalRequest)
			if err != nil {
				return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillResolver: error updating parent budget request"))
			}
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here is the list of updated items.",
		Items:   resItemList,
	}, nil
}

func (r *Resolver) FinancialBudgetFillActualResolver(params graphql.ResolveParams) (interface{}, error) {
	requestID := params.Args["request_id"].(int)

	request, err := r.Repo.GetBudgetRequest(requestID)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	if request.RequestType != structs.RequestTypeCurrentFinancial {
		return errors.HandleAPPError(errors.NewBadRequestError(
			"request must be if type: %s, provided request is of type: %s",
			dto.GetRequestType(structs.RequestTypeCurrentFinancial),
			dto.GetRequestType(request.RequestType),
		))
	}

	var items []dto.FillActualFinanceBudgetInput

	dataBytes, _ := json.Marshal(params.Args["data"])
	err = json.Unmarshal(dataBytes, &items)
	if err != nil {
		return errors.HandleAPPError(errors.WrapBadRequestError(err, "FinancialBudgetFillActualResolver"))
	}

	var resItemList []dto.FilledFinancialBudgetResItem

	for _, data := range items {
		item, err := r.Repo.FillActualFinancialBudget(params.Context, data.ID, data.Actual)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillActualResolver"))
		}
		resItem, err := buildFilledFinancialBudgetResItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillActualResolver"))
		}

		resItemList = append(resItemList, *resItem)
	}

	request.Status = structs.BudgetRequestCompletedActualStatus
	_, err = r.Repo.UpdateBudgetRequest(params.Context, request)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillActualResolver"))
	}

	financialRequest, err := r.Repo.GetBudgetRequest(*request.ParentID)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillActualResolver: error getting parent financial request"))
	}
	financialRequest.Status = structs.BudgetRequestCompletedActualStatus
	_, err = r.Repo.UpdateBudgetRequest(params.Context, financialRequest)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillActualResolver: error updating parent budget request"))
	}

	generalRequest, err := r.Repo.GetBudgetRequest(*financialRequest.ParentID)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillActualResolver: error getting parent financial request"))
	}
	generalRequest.Status = structs.BudgetRequestCompletedActualStatus
	_, err = r.Repo.UpdateBudgetRequest(params.Context, generalRequest)
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillActualResolver: error updating parent budget request"))
	}

	generallReqType := structs.RequestTypeGeneral
	budgetGeneralRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID:    &financialRequest.BudgetID,
		RequestType: &generallReqType,
	})
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillActualResolver: error updating parent budget request"))
	}

	allGeneralRequestsCompleted := true
	for _, generalReq := range budgetGeneralRequests {
		if generalReq.Status != structs.BudgetRequestCompletedActualStatus {
			allGeneralRequestsCompleted = false
			break
		}
	}

	if allGeneralRequestsCompleted {
		budget, err := r.Repo.GetBudget(request.BudgetID)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillActualResolver"))
		}
		budget.Status = structs.BudgetCompletedActualStatus

		accounts, err := r.Repo.GetAccountItems(nil)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "get all accounts"))
		}

		for _, req := range budgetGeneralRequests {
			for _, account := range accounts.Data {
				actual, err := r.Repo.GetSpendingDynamicActual(req.BudgetID, req.OrganizationUnitID, account.ID)
				if err != nil {
					return errors.HandleAPPError(errors.WrapInternalServerError(err, "get actual"))
				}

				_, err = r.Repo.CreateCurrentBudget(params.Context, &structs.CurrentBudget{
					BudgetID:      req.BudgetID,
					UnitID:        req.OrganizationUnitID,
					AccountID:     account.ID,
					InitialActual: actual.Decimal,
					Actual:        actual.Decimal,
				})
				if err != nil {
					return errors.HandleAPPError(errors.WrapInternalServerError(err, "create current budgets"))
				}
			}
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here is the list of updated items.",
		Items:   resItemList,
	}, nil
}

func buildFilledFinancialBudgetResItem(r repository.MicroserviceRepositoryInterface, filledFinancialBudget structs.FilledFinanceBudget) (*dto.FilledFinancialBudgetResItem, error) {
	resItem := &dto.FilledFinancialBudgetResItem{
		ID:            filledFinancialBudget.ID,
		CurrentYear:   filledFinancialBudget.CurrentYear,
		NextYear:      filledFinancialBudget.NextYear,
		YearAfterNext: filledFinancialBudget.YearAfterNext,
		Actual:        filledFinancialBudget.Actual,
		Description:   filledFinancialBudget.Description,
		RequestID:     filledFinancialBudget.BudgetRequestID,
	}

	budgetRequest, err := r.GetBudgetRequest(filledFinancialBudget.BudgetRequestID)
	if err != nil {
		return nil, err
	}

	organizationUnit, err := r.GetOrganizationUnitByID(budgetRequest.OrganizationUnitID)
	if err != nil {
		return nil, err
	}
	resItem.OrganizationUnit = dto.DropdownSimple{ID: organizationUnit.ID, Title: organizationUnit.Title}

	account, err := r.GetAccountItemByID(filledFinancialBudget.AccountID)
	if err != nil {
		return nil, err
	}
	resItem.Account = dto.DropdownSimple{ID: account.ID, Title: account.Title}

	return resItem, nil
}
