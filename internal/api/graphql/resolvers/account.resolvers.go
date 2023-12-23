package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"
	"unicode"

	"github.com/graphql-go/graphql"
)

func buildAccountItemResponseItem(item *structs.AccountItem) (*dto.AccountItemResponseItem, error) {
	return &dto.AccountItemResponseItem{
		ID:           item.ID,
		SerialNumber: item.SerialNumber,
		Title:        item.Title,
		ParentID:     item.ParentID,
	}, nil
}

func buildAccountItemResponseItemList(accountItems []*structs.AccountItem) ([]*dto.AccountItemResponseItem, error) {
	var responseItems []*dto.AccountItemResponseItem
	for _, item := range accountItems {
		resItem, err := buildAccountItemResponseItem(item)
		if err != nil {
			return nil, err
		}
		responseItems = append(responseItems, resItem)
	}

	return responseItems, nil
}

func (r *Resolver) AccountOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var accountFilters dto.GetAccountsFilter
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		accountFilters.ID = &id
	}
	if search, ok := params.Args["search"].(string); ok && search != "" {
		accountFilters.Search = &search
	}
	if page, ok := params.Args["page"].(int); ok && page != 0 {
		accountFilters.Page = &page
	}
	if size, ok := params.Args["size"].(int); ok && size != 0 {
		accountFilters.Size = &size
	}

	accounts, err := r.Repo.GetAccountItems(&accountFilters)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	accountResItemlist, err := buildAccountItemResponseItemList(accounts.Data)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	if tree, ok := params.Args["tree"].(bool); ok && tree {
		accountResItemlist, err = CreateTree(accountResItemlist)
		if err != nil {
			fmt.Printf("Fetching Account failed because of this error - %s.\n", err)
		}

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Total:   accounts.Total,
			Items:   accountResItemlist,
		}, nil
	}
	if level, ok := params.Args["level"].(int); ok {
		accountsMap := make(map[int]*dto.AccountItemResponseItem)
		for _, account := range accountResItemlist {
			accountsMap[account.ID] = account
		}

		var filteredAccounts []*dto.AccountItemResponseItem
		for _, account := range accountResItemlist {
			if determineLevel(account.SerialNumber) == level {
				filteredAccounts = append(filteredAccounts, account)
			}
		}
		accountResItemlist = filteredAccounts
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   accounts.Total,
		Items:   accountResItemlist,
	}, nil
}

func determineLevel(serialNumber string) int {
	count := 0

	for _, char := range serialNumber {
		if unicode.IsDigit(char) {
			count++
		}
	}

	return count - 1
}

func (r *Resolver) AccountInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.AccountItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		res, err := r.Repo.UpdateAccountItem(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		item, err := buildAccountItemResponseItem(res)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateAccountItem(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		item, err := buildAccountItemResponseItem(res)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) AccountDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteAccount(itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
