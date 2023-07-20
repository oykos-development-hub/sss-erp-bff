package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/graphql-go/graphql"
)

type bySerialNumber []*structs.AccountItem

func (a bySerialNumber) Len() int           { return len(a) }
func (a bySerialNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySerialNumber) Less(i, j int) bool { return a[i].SerialNumber < a[j].SerialNumber }

func AccountItemProperties(basicInventoryItems []interface{}, id int) []interface{} {
	var items []interface{}

	for _, item := range basicInventoryItems {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}

		items = append(items, mergedItem)
	}

	return items
}

var AccountOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	var tree bool
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	if params.Args["tree"] == nil {
		tree = false
	} else {
		tree = params.Args["tree"].(bool)
	}

	page := params.Args["page"]
	size := params.Args["size"]

	AccountType := &structs.AccountItem{}
	AccountData, err := shared.ReadJson(shared.GetDataRoot()+"/account.json", AccountType)

	if err != nil {
		fmt.Printf("Fetching Account failed because of this error - %s.\n", err)
	}

	// Populate data for each Basic Inventory Real Estates
	items = AccountItemProperties(AccountData, id)

	total = len(items)

	if id != 0 {
		return map[string]interface{}{
			"status":  "success",
			"message": "Here's the list you asked for!",
			"total":   total,
			"items":   items,
		}, nil
	}

	if tree == true {
		items, err = CreateTree(AccountData)
		if err != nil {
			fmt.Printf("Fetching Account failed because of this error - %s.\n", err)
		}

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

	} else {
		childMap := make(map[int][]*structs.AccountItem)
		for _, node := range AccountData {
			if nodeMap, ok := node.(*structs.AccountItem); ok {
				childMap[nodeMap.ParentId] = append(childMap[nodeMap.ParentId], nodeMap)
			}
		}

		for _, children := range childMap {
			sort.Sort(bySerialNumber(children))
		}

		var result []interface{}

		appendChildren(&result, childMap, 0)

		if err != nil {
			fmt.Printf("Fetching Account failed because of this error - %s.\n", err)
		}

		// Filtering by Pagination params
		if shared.IsInteger(page) && page != 0 && shared.IsInteger(size) && size != 0 {
			items = shared.Pagination(items, page.(int), size.(int))
		}

		return map[string]interface{}{
			"status":  "success",
			"message": "Here's the list you asked for!",
			"total":   total,
			"items":   result,
		}, nil
	}
}

var AccountInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.AccountItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	AccountItemType := &structs.AccountItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id

	AccountData, err := shared.ReadJson(shared.GetDataRoot()+"/account.json", AccountItemType)

	if err != nil {
		fmt.Printf("Fetching Account failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		AccountData = shared.FilterByProperty(AccountData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(AccountData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/account.json"), updatedData)

	sliceData := []interface{}{data}

	// Populate data for each Basic Inventory
	var populatedData = AccountItemProperties(sliceData, data.Id)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"items":   populatedData[0],
	}, nil
}

var AccountDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	AccountItemType := &structs.AccountItem{}
	AccountData, err := shared.ReadJson(shared.GetDataRoot()+"/account.json", AccountItemType)

	if err != nil {
		fmt.Printf("Fetching Account failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		AccountData = shared.FilterByProperty(AccountData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/account.json"), AccountData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

func CreateTree(nodes []interface{}) ([]interface{}, error) {
	mappedNodes := map[int]*structs.AccountItemNode{}
	for _, nodeInterface := range nodes {
		if nodeMap, ok := nodeInterface.(*structs.AccountItem); ok {
			node := &structs.AccountItemNode{
				Id:           nodeMap.Id,
				ParentId:     nodeMap.ParentId,
				SerialNumber: nodeMap.SerialNumber,
				Title:        nodeMap.Title,
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

func appendChildren(result *[]interface{}, childMap map[int][]*structs.AccountItem, parentId int) {
	children, exist := childMap[parentId]
	if !exist {
		return
	}

	for _, child := range children {
		*result = append(*result, child)
		appendChildren(result, childMap, child.Id)
	}
}
