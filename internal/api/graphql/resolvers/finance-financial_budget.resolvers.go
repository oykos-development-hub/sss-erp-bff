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

func (r *Resolver) FinancialBudgetOverview(params graphql.ResolveParams) (interface{}, error) {
	budgetID := params.Args["budget_id"].(int)
	var response dto.FinancialBudgetOverviewResponse

	financialBudget, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	response.AccountVersion = financialBudget.AccountVersion

	accounts, err := r.Repo.GetAccountItems(&dto.GetAccountsFilter{Version: &financialBudget.AccountVersion})
	if err != nil {
		return errors.HandleAPIError(err)
	}
	accountResItemlist, _ := buildAccountItemResponseItemList(accounts.Data)

	if params.Args["organization_unit_id"] == nil || params.Args["organization_unit_id"] == 0 {
		return dto.ResponseSingle{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Item:    response,
		}, nil
	}

	unitID := params.Args["organization_unit_id"].(int)

	currentFinancialType := structs.CurrentFinancialRequestType
	currentFinancialBudgetRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &unitID,
		BudgetID:           budgetID,
		RequestType:        &currentFinancialType,
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}
	currentFinancialRequestResList, err := buildFilledRequestData(r.Repo, accountResItemlist, currentFinancialBudgetRequest.ID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	donationFinancialType := structs.DonationFinancialRequestType
	donationFinancialBudgetRequest, err := r.Repo.GetOneBudgetRequest(&dto.GetBudgetRequestListInputMS{
		OrganizationUnitID: &unitID,
		BudgetID:           budgetID,
		RequestType:        &donationFinancialType,
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}
	donationFinancialRequestResList, err := buildFilledRequestData(r.Repo, accountResItemlist, donationFinancialBudgetRequest.ID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	status, err := buildFinancialBudgetStatus(currentFinancialBudgetRequest, donationFinancialBudgetRequest)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	financialBudgetOveriew := &dto.FinancialBudgetOverviewResponse{
		AccountVersion:                 financialBudget.AccountVersion,
		CurrentAccountsWithFilledData:  currentFinancialRequestResList.CreateTree(),
		CurrentRequestID:               currentFinancialBudgetRequest.ID,
		DonationAccountsWithFilledData: donationFinancialRequestResList.CreateTree(),
		DonationRequestID:              donationFinancialBudgetRequest.ID,
		Status:                         status,
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    financialBudgetOveriew,
	}, nil
}

func buildFilledRequestData(r repository.MicroserviceRepositoryInterface, accounts []*dto.AccountItemResponseItem, requestID int) (dto.AccountWithFilledFinanceBudgetResponseList, error) {
	var filledAccountResItemList dto.AccountWithFilledFinanceBudgetResponseList

	filledAccounts, err := r.GetFilledFinancialBudgetList(requestID)
	if err != nil {
		return filledAccountResItemList, err
	}

	for _, account := range accounts {
		filledAccountItem, err := buildAccountWithFilledFinanceResponseItem(account, filledAccounts)
		if err != nil {
			return filledAccountResItemList, err
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

func buildFinancialBudgetStatus(currentFinancialRequests structs.BudgetRequest, donationFinancialRequests structs.BudgetRequest) (dto.FinancialBudgetStatus, error) {
	currentFinancialStatus, err := buildBudgetRequestStatus(currentFinancialRequests)
	if err != nil {
		return "", err
	}
	donationFinancialStatus, err := buildBudgetRequestStatus(donationFinancialRequests)
	if err != nil {
		return "", err
	}

	if currentFinancialStatus == dto.FinancialBudgetFinishedStatus && donationFinancialStatus == dto.FinancialBudgetFinishedStatus {
		return dto.FinancialBudgetFinishedStatus, nil
	}

	return dto.FinancialBudgetTakeActionStatus, err
}

func buildBudgetRequestStatus(financialBudgetRequests structs.BudgetRequest) (dto.FinancialBudgetStatus, error) {
	if financialBudgetRequests.Status == structs.BudgetRequestSentStatus {
		return dto.FinancialBudgetTakeActionStatus, nil
	} else if financialBudgetRequests.Status == structs.BudgetRequestFinishedStatus {
		return dto.FinancialBudgetFinishedStatus, nil
	}

	return "", fmt.Errorf("could not determine status of request budget")
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
	request.Status = structs.BudgetRequestFinishedStatus
	_, err = r.Repo.UpdateBudgetRequest(request)
	if err != nil {
		return errors.HandleAPIError(err)
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
