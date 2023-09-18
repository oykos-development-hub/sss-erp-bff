package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/graphql-go/graphql"
)

func BudgetAccountItemProperties(basicInventoryItems []interface{}, budgetId int) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by budget ID
		if shared.IsInteger(budgetId) && budgetId != 0 && budgetId != mergedItem["budget_id"] {
			continue
		}

		items = append(items, mergedItem)
	}

	return items
}

var BudgetAccountOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var accounts []interface{}
	var budgetId int
	var activityId int
	if params.Args["budget_id"] == nil {
		budgetId = 0
	} else {
		budgetId = params.Args["budget_id"].(int)
	}

	if params.Args["activity_id"] == nil {
		activityId = 0
	} else {
		activityId = params.Args["activity_id"].(int)
	}

	AccountType := &structs.AccountItem{}
	AccountData, err := shared.ReadJson(shared.GetDataRoot()+"/account.json", AccountType)

	if err != nil {
		fmt.Printf("Fetching account_budget_activity failed because of this error - %s.\n", err)
	}

	// Populate data for each Basic Inventory Real Estates
	accounts = AccountItemProperties(AccountData, 0)

	accounts, err = CreateTree(AccountData, budgetId, activityId)
	for _, account := range accounts {
		if item, ok := account.(*structs.AccountItemNode); ok {
			updateParentValues(item)
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"items":   accounts,
	}, nil
}

var BudgetAccountInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var dataArray []structs.AccountBudgetActivityItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	AccountBudgetActivityItemType := &structs.AccountBudgetActivityItem{}

	_ = json.Unmarshal(dataBytes, &dataArray)

	if len(dataArray) > 0 {
		for _, data := range dataArray {
			itemId := data.Id

			AccountBudgetActivityData, err := shared.ReadJson(shared.GetDataRoot()+"/account_budget_activity.json", AccountBudgetActivityItemType)

			if err != nil {
				fmt.Printf("Fetching Account Budget Activity failed because of this error - %s.\n", err)
			}

			if shared.IsInteger(itemId) && itemId != 0 {
				AccountBudgetActivityData = shared.FilterByProperty(AccountBudgetActivityData, "Id", itemId)
			} else {
				data.Id = shared.GetRandomNumber()
			}

			var updatedData = append(AccountBudgetActivityData, data)

			_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/account_budget_activity.json"), updatedData)
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You account budget activity this item!",
	}, nil
}

func CreateTree(nodes []interface{}, budgetId int, activityId int) ([]interface{}, error) {
	var accountBudgetActivity = shared.FetchByProperty(
		"account_budget_activity",
		"BudgetId",
		budgetId,
	)

	if activityId > 0 {
		accountBudgetActivity = shared.FindByProperty(accountBudgetActivity, "ActivityId", activityId)
	}

	mappedNodes := map[int]*structs.AccountItemNode{}
	for _, nodeInterface := range nodes {
		if nodeMap, ok := nodeInterface.(*structs.AccountItem); ok {

			node := &structs.AccountItemNode{
				Id:           nodeMap.Id,
				ParentId:     nodeMap.ParentId,
				SerialNumber: nodeMap.SerialNumber,
				Title:        nodeMap.Title,
			}
			var values = shared.FindByProperty(accountBudgetActivity, "AccountId", nodeMap.Id)
			if len(values) > 0 {
				for _, value := range values {
					if accountValue, ok := value.(*structs.AccountBudgetActivityItem); ok {
						node.ValueNextYear = accountValue.ValueNextYear
						node.ValueAfterNextYear = accountValue.ValueAfterNextYear
						node.ValueCurrentYear = accountValue.ValueCurrentYear
					}
				}
			}
			mappedNodes[nodeMap.Id] = node
		}
	}

	sortedNodes := make([]*structs.AccountItemNode, 0, len(mappedNodes))
	for _, node := range mappedNodes {
		sortedNodes = append(sortedNodes, node)
	}
	sort.Slice(sortedNodes, func(i, j int) bool {
		return sortedNodes[i].SerialNumber < sortedNodes[j].SerialNumber
	})

	var rootNodes []interface{}
	for _, node := range sortedNodes {
		if parentNode, ok := mappedNodes[node.ParentId]; ok {
			parentNode.Children = append(parentNode.Children, node)
		} else {
			rootNodes = append(rootNodes, node)
		}
	}

	if len(rootNodes) == 0 {
		return nil, fmt.Errorf("no root node found")
	}

	return rootNodes, nil
}

func updateParentValues(node *structs.AccountItemNode) {
	if len(node.Children) == 0 {
		return
	}

	sumCurrentYear, sumNextYear, sumAfterNextYear := 0, 0, 0

	for _, child := range node.Children {
		updateParentValues(child)

		sumCurrentYear += child.ValueCurrentYear
		sumNextYear += child.ValueNextYear
		sumAfterNextYear += child.ValueAfterNextYear
	}

	node.ValueCurrentYear = sumCurrentYear
	node.ValueNextYear = sumNextYear
	node.ValueAfterNextYear = sumAfterNextYear
}
