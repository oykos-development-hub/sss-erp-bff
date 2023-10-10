package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

func buildAccountItemResponseItem(item *structs.AccountItem) (*dto.AccountItemResponseItem, error) {
	return &dto.AccountItemResponseItem{
		ID:           item.Id,
		SerialNumber: item.SerialNumber,
		Title:        item.Title,
		ParentId:     item.ParentId,
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

var AccountOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
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

	accounts, err := getAccountItems(&accountFilters)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	accountResItemlist, err := buildAccountItemResponseItemList(accounts.Data)
	if err != nil {
		return shared.HandleAPIError(err)
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
			if determineLevel(account, accountsMap) == level {
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

func determineLevel(account *dto.AccountItemResponseItem, accountsMap map[int]*dto.AccountItemResponseItem) int {
	level := 0
	for account.ParentId != nil {
		account = accountsMap[*account.ParentId]
		if account == nil {
			break
		}
		level++
	}
	return level
}

var AccountInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.AccountItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateAccountItem(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildAccountItemResponseItem(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := createAccountItem(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildAccountItemResponseItem(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You created this item!"
	}

	return response, nil
}

var AccountDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteAccount(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func deleteAccount(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.ACCOUNT_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getAccountItems(filters *dto.GetAccountsFilter) (*dto.GetAccountItemListResponseMS, error) {
	res := &dto.GetAccountItemListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ACCOUNT_ENDPOINT, filters, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getAccountItemById(id int) (*structs.AccountItem, error) {
	res := &dto.GetAccountItemResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ACCOUNT_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func createAccountItem(accountItem *structs.AccountItem) (*structs.AccountItem, error) {
	res := &dto.GetAccountItemResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ACCOUNT_ENDPOINT, accountItem, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func updateAccountItem(id int, accountItem *structs.AccountItem) (*structs.AccountItem, error) {
	res := &dto.GetAccountItemResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.ACCOUNT_ENDPOINT+"/"+strconv.Itoa(id), accountItem, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
