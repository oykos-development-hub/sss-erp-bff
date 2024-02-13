package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) FinancialBudgetOverview(params graphql.ResolveParams) (interface{}, error) {
	var status dto.FinancialBudgetStatus
	budgetID := params.Args["budget_id"].(int)
	financialBudget, err := r.Repo.GetFinancialBudgetByBudgetID(budgetID)
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

	unitID := params.Args["organization_unit_id"].(int)

	filledAccounts, err := r.Repo.GetFilledFinancialBudgetList(unitID, financialBudget.ID)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	if len(filledAccounts) > 0 {
		status = dto.FinancialBudgetFinishedStatus
	} else {
		status = dto.FinancialBudgetTakeActionStatus
	}

	var filledAccountResItemList dto.AccountWithFilledFinanceBudgetResponseList
	for _, account := range accountResItemlist {
		filledAccountItem, err := buildAccountWithFilledFinanceResponseItem(account, filledAccounts)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		filledAccountResItemList.Data = append(filledAccountResItemList.Data, filledAccountItem)
	}

	resItem := filledAccountResItemList.CreateTree()
	if err != nil {
		return errors.HandleAPIError(err)
	}

	financialBudgetOveriew := &dto.FinancialBudgetOverviewResponse{
		AccountVersion:         financialBudget.AccountVersion,
		AccountsWithFilledData: resItem,
		Status:                 status,
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    financialBudgetOveriew,
	}, nil
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

func (r *Resolver) FinancialBudgetFillResolver(params graphql.ResolveParams) (interface{}, error) {
	var items []structs.FilledFinanceBudget

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &items)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	var resItemList []dto.FilledFinancialBudgetResItem

	for _, data := range items {
		itemID := data.ID

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

	return dto.Response{
		Status:  "success",
		Message: "Here is the list of updated items.",
		Items:   resItemList,
	}, nil
}

func buildFilledFinancialBudgetResItem(r repository.MicroserviceRepositoryInterface, filledFinancialBudget structs.FilledFinanceBudget) (*dto.FilledFinancialBudgetResItem, error) {
	resItem := &dto.FilledFinancialBudgetResItem{
		ID:                filledFinancialBudget.ID,
		CurrentYear:       filledFinancialBudget.CurrentYear,
		NextYear:          filledFinancialBudget.NextYear,
		YearAfterNext:     filledFinancialBudget.YearAfterNext,
		Description:       filledFinancialBudget.Description,
		FinancialBudgetID: filledFinancialBudget.FinanceBudgetID,
	}

	organizationUnit, err := r.GetOrganizationUnitByID(filledFinancialBudget.OrganizationUnitID)
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
