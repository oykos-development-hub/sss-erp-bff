package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
)

func PopulateUserAccountProperties(userAccounts []interface{}, filters ...interface{}) []interface{} {
	var id int
	var isActive interface{}
	var email string

	switch len(filters) {
	case 1:
		id = filters[0].(int)
	case 2:
		id = filters[0].(int)
		isActive = filters[1]
	case 3:
		id = filters[0].(int)
		isActive = filters[1]
		email = filters[2].(string)
	}

	var items []interface{}

	for _, userAccount := range userAccounts {
		var user = shared.WriteStructToInterface(userAccount)

		if shared.IsInteger(id) && id > 0 && id != user["id"] {
			continue
		}
		if isActive != nil && isActive != user["active"] {
			continue
		}
		if shared.IsString(email) && len(email) > 0 && !shared.StringContains(user["email"].(string), email) {
			continue
		}

		var relatedRole = shared.FetchByProperty(
			"roles",
			"Id",
			user["role_id"],
		)

		if relatedRole != nil && len(relatedRole) > 0 {
			for _, role := range relatedRole {
				var roleData = shared.WriteStructToInterface(role)
				roleItem := make(map[string]interface{})

				roleItem["title"] = roleData["title"]
				roleItem["id"] = roleData["id"]

				user["role"] = roleItem
			}
		}

		items = append(items, user)
	}

	return items
}

var UserAccountsOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var total int
	var itemId int
	if params.Args["id"] == nil {
		itemId = 0
	} else {
		itemId = params.Args["id"].(int)
	}
	var isActive = params.Args["is_active"]
	page := params.Args["page"]
	size := params.Args["size"]
	email := params.Args["email"]

	if !shared.IsString(email) {
		email = ""
	}

	var items []interface{}
	var userAccounts = shared.FetchByProperty(
		"user_account",
		"",
		"",
	)

	if userAccounts != nil && len(userAccounts) > 0 {
		items = PopulateUserAccountProperties(userAccounts, itemId, isActive, email)
	}

	total = len(items)

	// Filtering by Pagination params
	if shared.IsInteger(page) && page != 0 && shared.IsInteger(size) && size != 0 {
		items = shared.Pagination(items, page.(int), size.(int))
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"total":   total,
		"items":   items,
	}, nil
}

var UserAccountBasicInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.UserAccounts
	dataBytes, _ := json.Marshal(params.Args["data"])
	UserAccountType := &structs.UserAccounts{}

	json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	userAccountData, userAccountDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_accounts.json", UserAccountType)

	if userAccountDataErr != nil {
		fmt.Printf("Fetching User Accounts failed because of this error - %s.\n", userAccountDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		userAccountData = shared.FilterByProperty(userAccountData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(userAccountData, data)

	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_accounts.json"), updatedData)

	var populatedItems = PopulateUserAccountProperties([]interface{}{data})

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   populatedItems,
	}, nil
}

var UserAccountDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	UserAccountType := &structs.UserAccounts{}
	userAccountData, userAccountDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_accounts.json", UserAccountType)

	if userAccountDataErr != nil {
		fmt.Printf("Fetching User Accounts failed because of this error - %s.\n", userAccountDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		userAccountData = shared.FilterByProperty(userAccountData, "Id", itemId)
	}

	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_accounts.json"), userAccountData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
