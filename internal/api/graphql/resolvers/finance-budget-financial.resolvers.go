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

	financialBudgetOveriew, err := r.GetFinancialBudgetDetails(params.Context, budgetID, unitID)
	if err != nil {
		return errors.HandleAPPError(errors.Wrap(err, "FinancialBudgetOverview"))
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    financialBudgetOveriew,
	}, nil
}

func (r *Resolver) GetFinancialBudgetDetails(ctx context.Context, budgetID, unitID int) (*dto.FinancialBudgetOverviewResponse, error) {
	var response dto.FinancialBudgetOverviewResponse

	financialBudget, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails: error getting financial budget")
	}
	response.AccountVersion = financialBudget.AccountVersion

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{Version: &financialBudget.AccountVersion})
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails: error getting account items")
	}
	accountResItemlist, _ := buildAccountItemResponseItemList(accounts.Data)

	currentFinancialType := structs.CurrentFinancialRequestType
	currentFinancialBudgetRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &unitID,
		BudgetID:           budgetID,
		RequestType:        &currentFinancialType,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails: error getting current financial budget request")
	}
	currentFinancialRequestResList, err := buildFilledRequestData(r.Repo, accountResItemlist, currentFinancialBudgetRequest.ID)
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails")
	}

	donationFinancialType := structs.DonationFinancialRequestType
	donationFinancialBudgetRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &unitID,
		BudgetID:           budgetID,
		RequestType:        &donationFinancialType,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails")
	}
	donationFinancialRequestResList, err := buildFilledRequestData(r.Repo, accountResItemlist, donationFinancialBudgetRequest.ID)
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails")
	}

	financialRequestType := structs.FinancialRequestType
	financialParentRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &unitID,
		BudgetID:           budgetID,
		RequestType:        &financialRequestType,
	})
	if err != nil {
		return nil, errors.Wrap(err, "GetFinancialBudgetDetails")
	}

	financialBudgetOveriew := &dto.FinancialBudgetOverviewResponse{
		RequestID:                      financialParentRequest.ID,
		AccountVersion:                 financialBudget.AccountVersion,
		CurrentAccountsWithFilledData:  currentFinancialRequestResList.CreateTree(),
		CurrentRequestID:               currentFinancialBudgetRequest.ID,
		DonationAccountsWithFilledData: donationFinancialRequestResList.CreateTree(),
		DonationRequestID:              donationFinancialBudgetRequest.ID,
		Status:                         buildBudgetRequestStatus(ctx, financialParentRequest.Status),
		DonationBudgetStatus:           buildBudgetRequestStatus(ctx, donationFinancialBudgetRequest.Status),
		CurrentBudgetStatus:            buildBudgetRequestStatus(ctx, currentFinancialBudgetRequest.Status),
		OfficialComment:                financialParentRequest.Comment,
		CurrentBudgetComment:           currentFinancialBudgetRequest.Comment,
		DonationBudgetComment:          donationFinancialBudgetRequest.Comment,
	}
	return financialBudgetOveriew, nil
}

func buildFilledRequestData(r repository.MicroserviceRepositoryInterface, accounts []*dto.AccountItemResponseItem, requestID int) (dto.AccountWithFilledFinanceBudgetResponseList, error) {
	var filledAccountResItemList dto.AccountWithFilledFinanceBudgetResponseList

	filledAccounts, err := r.GetFilledFinancialBudgetList(requestID)
	if err != nil {
		return filledAccountResItemList, errors.Wrap(err, "buildFilledRequestData")
	}

	for _, account := range accounts {
		filledAccountItem, err := buildAccountWithFilledFinanceResponseItem(account, filledAccounts)
		if err != nil {
			return filledAccountResItemList, errors.Wrap(err, "buildFilledRequestData")
		}
		filledAccountResItemList.Data = append(filledAccountResItemList.Data, filledAccountItem)
	}

	return filledAccountResItemList, nil
}

func buildAccountWithFilledFinanceResponseItem(item *dto.AccountItemResponseItem, filledAccounts []structs.FilledFinanceBudget) (*dto.AccountWithFilledFinanceBudget, error) {
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

	_, err = r.Repo.UpdateFinancialBudget(financialBudget)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	financialBudgetRequests, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		BudgetID:     budgetID,
		RequestTypes: []structs.RequestType{structs.CurrentFinancialRequestType, structs.DonationFinancialRequestType},
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, request := range financialBudgetRequests {
		filledRequestData, err := r.Repo.GetFilledFinancialBudgetList(request.ID)
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
				_, err := r.Repo.UpdateFilledFinancialBudget(filledData.ID, &structs.FilledFinanceBudget{
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
				err := r.Repo.DeleteFilledFinancialBudgetData(id)
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

	var items []structs.FilledFinanceBudget

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &items)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	var resItemList []dto.FilledFinancialBudgetResItem

	for _, data := range items {
		itemID := data.ID
		data.BudgetRequestID = requestID
		if itemID != 0 {
			item, err := r.Repo.UpdateFilledFinancialBudget(itemID, &data)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			resItem, err := buildFilledFinancialBudgetResItem(r.Repo, *item)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			resItemList = append(resItemList, *resItem)
		} else {
			item, err := r.Repo.FillFinancialBudget(&data)
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

	request, err := r.Repo.GetBudgetRequest(requestID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	request.Status = structs.BudgetRequestFilledStatus
	request.Comment = params.Args["comment"].(string)

	_, err = r.Repo.UpdateBudgetRequest(request)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	//update parent financial status to sent if all financial chidrens are filled
	financialChildRequestList, err := r.Repo.GetBudgetRequestList(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &request.OrganizationUnitID,
		BudgetID:           request.BudgetID,
		RequestTypes: []structs.RequestType{
			structs.CurrentFinancialRequestType,
			structs.DonationFinancialRequestType,
		},
	})
	if err != nil {
		return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillResolver: error getting financial child requests"))
	}

	allFilled := true
	for _, financialChildRequest := range financialChildRequestList {
		if financialChildRequest.Status != structs.BudgetRequestFilledStatus {
			allFilled = false
			break
		}
	}

	parentFinancialRequest := structs.FinancialRequestType
	if allFilled {
		parentFinancialRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
			OrganizationUnitID: &request.OrganizationUnitID,
			BudgetID:           request.BudgetID,
			RequestType:        &parentFinancialRequest,
		})
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillResolver: error getting parent financial request"))
		}
		parentFinancialRequest.Status = structs.BudgetRequestFilledStatus
		_, err = r.Repo.UpdateBudgetRequest(parentFinancialRequest)
		if err != nil {
			return errors.HandleAPPError(errors.WrapInternalServerError(err, "FinancialBudgetFillResolver: error updating parent budget request"))
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
