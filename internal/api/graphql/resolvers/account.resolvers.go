package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
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
			return nil, errors.Wrap(err, "build account item response item")
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
	if version, ok := params.Args["version"].(int); ok && version != 0 {
		accountFilters.Version = &version
	}
	leaf := params.Args["leaf"].(bool)
	accountFilters.Leaf = leaf

	accounts, err := r.Repo.GetAccountItems(&accountFilters)
	if err != nil {
		return errors.HandleAPPError(err)
	}
	accountResItemlist, err := buildAccountItemResponseItemList(accounts.Data)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	if tree, ok := params.Args["tree"].(bool); ok && tree {
		accountResItemlist, err = CreateTree(accountResItemlist)
		if err != nil {
			return errors.HandleAPPError(err)
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

	version := 0
	if len(accounts.Data) > 0 {
		version = accounts.Data[0].Version
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   accounts.Total,
		Items:   accountResItemlist,
		Version: version,
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
	var data []structs.AccountItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.Response{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	res, err := r.Repo.CreateAccountItemList(params.Context, data)
	if err != nil {
		return errors.HandleAPPError(err)
	}
	items, err := buildAccountItemResponseItemList(res)
	if err != nil {
		return errors.HandleAPPError(err)
	}
	response.Items = items
	response.Message = "You created a new version of account items!"

	return response, nil
}

func (r *Resolver) AccountDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteAccount(params.Context, itemID)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
