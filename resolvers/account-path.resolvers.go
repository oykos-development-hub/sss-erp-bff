package resolvers

// import (
// 	"bff/shared"
// 	"bff/structs"
// 	"encoding/json"
// 	"fmt"
// 	"sort"

// 	"github.com/graphql-go/graphql"
// )

// type bySerialNumber []*structs.AccountItemPath

// func (a bySerialNumber) Len() int           { return len(a) }
// func (a bySerialNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a bySerialNumber) Less(i, j int) bool { return a[i].SerialNumber < a[j].SerialNumber }

// var AccountOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
// 	var items []interface{}
// 	var total int
// 	var id int
// 	if params.Args["id"] == nil {
// 		id = 0
// 	} else {
// 		id = params.Args["id"].(int)
// 	}

// 	page := params.Args["page"]
// 	size := params.Args["size"]

// 	AccountType := &structs.AccountItemPath{}
// 	AccountData, err := shared.ReadJson(shared.GetDataRoot()+"/account.json", AccountType)

// 	if err != nil {
// 		fmt.Printf("Fetching Account failed because of this error - %s.\n", err)
// 	}

// 	// Populate data for each Basic Inventory Real Estates
// 	items = AccountItemProperties(AccountData, id)

// 	total = len(items)

// 	childMap := make(map[int][]*structs.AccountItemPath)
// 	for _, node := range AccountData {
// 		if nodeMap, ok := node.(*structs.AccountItemPath); ok {
// 			childMap[nodeMap.ParentId] = append(childMap[nodeMap.ParentId], nodeMap)
// 		}
// 	}

// 	for _, children := range childMap {
// 		sort.Sort(bySerialNumber(children))
// 	}

// 	var result []interface{}

// 	appendChildren(&result, childMap, 0, "")

// 	if err != nil {
// 		fmt.Printf("Fetching Account failed because of this error - %s.\n", err)
// 	}

// 	// Filtering by Pagination params
// 	if shared.IsInteger(page) && page != 0 && shared.IsInteger(size) && size != 0 {
// 		items = shared.Pagination(items, page.(int), size.(int))
// 	}

// 	return map[string]interface{}{
// 		"status":  "success",
// 		"message": "Here's the list you asked for!",
// 		"total":   total,
// 		"items":   result,
// 	}, nil
// }

// func appendChildren(result *[]interface{}, childMap map[int][]*structs.AccountItemPath, parentId int, parentPath string) {
// 	children, exist := childMap[parentId]
// 	if !exist {
// 		return
// 	}

// 	for _, child := range children {
// 		child.Path = parentPath + "/" + child.SerialNumber
// 		*result = append(*result, child)
// 		appendChildren(result, childMap, child.Id, child.Path)
// 	}
// }
