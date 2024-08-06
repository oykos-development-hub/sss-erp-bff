package resolvers

import (
	"bff/config"
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	latestVersion, err := r.Repo.GetLatestVersionOfAccounts()
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{Version: &financialBudget.AccountVersion})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	accountResItemlist, err := buildAccountItemResponseItemList(accounts.Data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    financialBudgetOveriew,
	}, nil
}

func (r *Resolver) GetFinancialBudgetDetails(ctx context.Context, budgetID, unitID int, summary bool) (*dto.FinancialBudgetOverviewResponse, error) {

	var filledAccounts []structs.FilledFinanceBudget
	var donationFinancialBudgetRequest *structs.BudgetRequest
	var currentFinancialBudgetRequest *structs.BudgetRequest

	financialBudget, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get financial budget by budget id")
	}

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{Version: &financialBudget.AccountVersion})
	if err != nil {
		return nil, errors.Wrap(err, "repo get account items")
	}

	currentFinancialType := structs.RequestTypeCurrentFinancial
	if summary {
		filledAccounts, err = r.Repo.GetFinancialFilledSummary(budgetID, currentFinancialType)
		if err != nil {
			return nil, errors.Wrap(err, "repo get one budget request")
		}
	} else {
		currentFinancialBudgetRequest, err = r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
			OrganizationUnitID: &unitID,
			BudgetID:           &budgetID,
			RequestType:        &currentFinancialType,
		})
		if err != nil {
			return nil, errors.Wrap(err, "repo get one budget request")
		}
		filledAccounts, err = r.Repo.GetFilledFinancialBudgetList(&dto.FilledFinancialBudgetInputMS{BudgetRequestID: currentFinancialBudgetRequest.ID})
		if err != nil {
			return nil, errors.Wrap(err, "repo get filled financial budget list")
		}
	}

	currentFinancialRequestResList, err := buildFilledRequestData(accounts.Data, filledAccounts)
	if err != nil {
		return nil, errors.Wrap(err, "build filled request data")
	}

	donationFinancialType := structs.RequestTypeDonationFinancial
	if summary {
		filledAccounts, err = r.Repo.GetFinancialFilledSummary(budgetID, donationFinancialType)
		if err != nil {
			return nil, errors.Wrap(err, "get financial filled summary")
		}
	} else {
		donationFinancialBudgetRequest, err = r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
			OrganizationUnitID: &unitID,
			BudgetID:           &budgetID,
			RequestType:        &donationFinancialType,
		})
		if err != nil {
			return nil, errors.Wrap(err, "repo get one budget request")
		}

		filledAccounts, err = r.Repo.GetFilledFinancialBudgetList(&dto.FilledFinancialBudgetInputMS{BudgetRequestID: donationFinancialBudgetRequest.ID})
		if err != nil {
			return nil, errors.Wrap(err, "build filled request data")
		}
	}
	donationFinancialRequestResList, err := buildFilledRequestData(accounts.Data, filledAccounts)
	if err != nil {
		return nil, errors.Wrap(err, "build filled request data")
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
			return nil, errors.Wrap(err, "repo get one budget request")
		}

		financialBudgetOveriew.RequestID = financialRequest.ID
		financialBudgetOveriew.OfficialComment = financialRequest.Comment
		status, err := r.buildBudgetRequestStatus(ctx, donationFinancialBudgetRequest.Status)
		if err != nil {
			return nil, errors.Wrap(err, "repo build budget request status")
		}
		financialBudgetOveriew.Status = *status

		financialBudgetOveriew.CurrentRequestID = currentFinancialBudgetRequest.ID
		financialBudgetOveriew.CurrentBudgetComment = currentFinancialBudgetRequest.Comment
		status, err = r.buildBudgetRequestStatus(ctx, donationFinancialBudgetRequest.Status)
		if err != nil {
			return nil, errors.Wrap(err, "repo build budget request status")
		}
		financialBudgetOveriew.CurrentBudgetStatus = *status
		financialBudgetOveriew.DonationRequestID = donationFinancialBudgetRequest.ID
		financialBudgetOveriew.CurrentBudgetComment = donationFinancialBudgetRequest.Comment
		status, err = r.buildBudgetRequestStatus(ctx, donationFinancialBudgetRequest.Status)
		if err != nil {
			return nil, errors.Wrap(err, "repo build budget request status")
		}
		financialBudgetOveriew.DonationBudgetStatus = *status
	}

	return financialBudgetOveriew, nil
}

func buildFilledRequestData(accounts []*structs.AccountItem, filledAccounts []structs.FilledFinanceBudget) (dto.AccountWithFilledFinanceBudgetResponseList, error) {
	var filledAccountResItemList dto.AccountWithFilledFinanceBudgetResponseList

	for _, account := range accounts {
		filledAccountItem, err := buildAccountWithFilledFinanceResponseItem(account, filledAccounts)
		if err != nil {
			return filledAccountResItemList, errors.Wrap(err, "build account with filled finance response item")
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	financialBudget, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	financialBudget.AccountVersion = latestVersion

	_, err = r.Repo.UpdateFinancialBudget(params.Context, financialBudget)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	financialBudgetRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID:     &budgetID,
		RequestTypes: []structs.RequestType{structs.RequestTypeCurrentFinancial, structs.RequestTypeDonationFinancial},
	})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, request := range financialBudgetRequests {
		filledRequestData, err := r.Repo.GetFilledFinancialBudgetList(&dto.FilledFinancialBudgetInputMS{BudgetRequestID: request.ID})
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		filledDataForDelete := make(map[int]bool)
		for _, filledData := range filledRequestData {
			filledDataForDelete[filledData.ID] = true
		}

		for _, filledData := range filledRequestData {
			oldAccount, err := r.Repo.GetAccountItemByID(filledData.AccountID)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			newAccount, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{
				Version:      &latestVersion,
				SerialNumber: &oldAccount.SerialNumber,
				Title:        &oldAccount.Title,
			})
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
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
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
				filledDataForDelete[filledData.ID] = false
			}
		}
		for id, delete := range filledDataForDelete {
			if delete {
				err := r.Repo.DeleteFilledFinancialBudgetData(params.Context, id)
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}
	}

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{Version: &financialBudget.AccountVersion})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	accountResItemlist, err := buildAccountItemResponseItemList(accounts.Data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var resItemList []dto.FilledFinancialBudgetResItem

	for _, data := range items {
		itemID := data.ID
		data.BudgetRequestID = requestID

		var accounts []int
		accounts = append(accounts, data.AccountID)
		item, err := r.Repo.GetFilledFinancialBudgetList(&dto.FilledFinancialBudgetInputMS{
			BudgetRequestID: data.BudgetRequestID,
			Accounts:        accounts,
		})

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		if len(item) > 0 {
			itemID = item[0].ID
		}

		if itemID != 0 {
			item, err := r.Repo.UpdateFilledFinancialBudget(params.Context, itemID, &data)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			resItem, err := buildFilledFinancialBudgetResItem(r.Repo, *item)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			resItemList = append(resItemList, *resItem)
		} else {
			item, err := r.Repo.FillFinancialBudget(params.Context, &data)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			resItem, err := buildFilledFinancialBudgetResItem(r.Repo, *item)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			resItemList = append(resItemList, *resItem)
		}
	}

	request.Status = structs.BudgetRequestFilledStatus
	request.Comment = params.Args["comment"].(string)
	_, err = r.Repo.UpdateBudgetRequest(params.Context, request)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	//update parent financial status to sent if all financial sub requests (current, donation) are filled
	financialSubRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		ParentID: request.ParentID,
	})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		financialRequest.Status = structs.BudgetRequestFilledStatus
		_, err = r.Repo.UpdateBudgetRequest(params.Context, financialRequest)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		financialAndNonFinancialRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
			ParentID: financialRequest.ParentID,
		})
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			generalRequest.Status = structs.BudgetRequestFilledStatus
			_, err = r.Repo.UpdateBudgetRequest(params.Context, generalRequest)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var resItemList []dto.FilledFinancialBudgetResItem

	for _, data := range items {
		_, err := r.Repo.FillActualFinancialBudget(params.Context, data.ID, data.Actual, data.Type, requestID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	request.Status = structs.BudgetRequestCompletedActualStatus
	_, err = r.Repo.UpdateBudgetRequest(params.Context, request)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	financialRequest, err := r.Repo.GetBudgetRequest(*request.ParentID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	financialRequest.Status = structs.BudgetRequestCompletedActualStatus
	_, err = r.Repo.UpdateBudgetRequest(params.Context, financialRequest)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	generalRequest, err := r.Repo.GetBudgetRequest(*financialRequest.ParentID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	generalRequest.Status = structs.BudgetRequestCompletedActualStatus
	_, err = r.Repo.UpdateBudgetRequest(params.Context, generalRequest)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	generallReqType := structs.RequestTypeGeneral
	budgetGeneralRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID:    &financialRequest.BudgetID,
		RequestType: &generallReqType,
	})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		budget.Status = structs.BudgetCompletedActualStatus
	}

	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	targetUsers, err := r.Repo.GetUsersByPermission(config.FinanceBudget, config.OperationRead)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	employees, err := GetEmployeesOfOrganizationUnit(r.Repo, generalRequest.OrganizationUnitID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, targetUser := range targetUsers {
		for _, employee := range employees {
			if targetUser.ID != loggedInUser.ID && employee.UserAccountID == targetUser.ID {
				_, err := r.NotificationsService.CreateNotification(&structs.Notifications{
					Content:     "Unesen je tekući budžet.",
					Module:      "Finansije",
					FromUserID:  loggedInUser.ID,
					ToUserID:    targetUser.ID,
					FromContent: "Službenik za budžet",
					Path:        "/finance/budget/current",
					IsRead:      false,
				})
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
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
		return nil, errors.Wrap(err, "repo get budget request")
	}

	organizationUnit, err := r.GetOrganizationUnitByID(budgetRequest.OrganizationUnitID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get organization unit by id")
	}
	resItem.OrganizationUnit = dto.DropdownSimple{ID: organizationUnit.ID, Title: organizationUnit.Title}

	account, err := r.GetAccountItemByID(filledFinancialBudget.AccountID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get account item by id")
	}
	resItem.Account = dto.DropdownSimple{ID: account.ID, Title: account.Title}

	return resItem, nil
}
