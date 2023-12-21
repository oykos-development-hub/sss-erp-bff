package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
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

func (r *Resolver) BudgetAccountOverviewResolver(params graphql.ResolveParams) (interface{}, error) {

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

	accountsTree, err := CreateTree(accountResItemlist)
	if err != nil {
		fmt.Printf("Create tree error - %s.\n", err)
	}
	// for _, account := range accountsTree {
	// 	if item, ok := account.(*structs.AccountItemNode); ok {
	// 		updateParentValues(item)
	// 	}
	// }

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"items":   accountsTree,
	}, nil
}

func (r *Resolver) BudgetAccountInsertResolver(params graphql.ResolveParams) (interface{}, error) {
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

func CreateTree(nodes []*dto.AccountItemResponseItem) ([]*dto.AccountItemResponseItem, error) {
	mappedNodes := make(map[int]*dto.AccountItemResponseItem, len(nodes))
	var rootNodes []*dto.AccountItemResponseItem

	// Create map and identify root nodes
	for _, node := range nodes {
		mappedNodes[node.ID] = node
		if node.ParentId == nil {
			rootNodes = append(rootNodes, node)
		}
	}

	// Populate children for each node
	for _, node := range nodes {
		if node.ParentId != nil {
			if parentNode, exists := mappedNodes[*node.ParentId]; exists {
				parentNode.Children = append(parentNode.Children, node)
			}
		}
	}

	// Sort the root nodes based on SerialNumber
	sort.Slice(rootNodes, func(i, j int) bool {
		return rootNodes[i].SerialNumber < rootNodes[j].SerialNumber
	})

	if len(rootNodes) == 0 {
		return nil, fmt.Errorf("no root node found")
	}

	return rootNodes, nil
}

//Zakomentarisano zbog pipeline-a
/*
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
*/
