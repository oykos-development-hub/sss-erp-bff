package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
)

func PopulateOrderListItemProperties(OrderList []interface{}, id int, supplierId int, status string, search string, publicProcurementId int) []interface{} {
	var items []interface{}

	for _, item := range OrderList {

		var mergedItem = shared.WriteStructToInterface(item)

		// Filtering by ID
		if shared.IsInteger(id) && id != 0 && id != mergedItem["id"] {
			continue
		}

		// if pass only id item
		if shared.IsInteger(id) && id != 0 && id == mergedItem["id"] && shared.IsInteger(publicProcurementId) && publicProcurementId == 0 {
			publicProcurementId = mergedItem["public_procurement_id"].(int)
		}

		// Filtering by supplierId
		if shared.IsInteger(supplierId) && supplierId != 0 && supplierId != mergedItem["supplier_id"] {
			continue
		}

		// Filtering by status
		if shared.IsString(status) && len(status) > 0 && status != mergedItem["status"] {
			continue
		}

		// Filtering by status
		if shared.IsString(search) && len(search) > 0 && !shared.StringContains(mergedItem["serial_number"].(string), search) {
			continue
		}

		if shared.IsInteger(mergedItem["supplier_id"]) && mergedItem["supplier_id"] != 0 {
			var relatedSuppliers = shared.FetchByProperty(
				"suppliers",
				"Id",
				mergedItem["supplier_id"],
			)

			if len(relatedSuppliers) > 0 {
				for _, supplierData := range relatedSuppliers {
					var supplier = shared.WriteStructToInterface(supplierData)

					mergedItem["supplier"] = map[string]interface{}{
						"id":    supplier["id"],
						"title": supplier["title"],
					}
				}
			}
		}

		if shared.IsInteger(mergedItem["public_procurement_id"]) && mergedItem["public_procurement_id"] != 0 {
			var relatedProcurement = shared.FetchByProperty(
				"procurement",
				"Id",
				mergedItem["public_procurement_id"],
			)

			if len(relatedProcurement) > 0 {
				for _, procurementData := range relatedProcurement {
					var procurement = shared.WriteStructToInterface(procurementData)

					mergedItem["public_procurement"] = map[string]interface{}{
						"id":    procurement["id"],
						"title": procurement["title"],
					}
				}
			}
		}

		if shared.IsInteger(mergedItem["organization_unit_id"]) && mergedItem["organization_unit_id"] != 0 {
			var relatedOrganizationUnit = shared.FetchByProperty(
				"organization_unit",
				"Id",
				mergedItem["organization_unit_id"],
			)

			if len(relatedOrganizationUnit) > 0 {
				for _, organizationUnitData := range relatedOrganizationUnit {
					var organizationUnit = shared.WriteStructToInterface(organizationUnitData)

					mergedItem["organization_unit"] = map[string]interface{}{
						"id":    organizationUnit["id"],
						"title": organizationUnit["title"],
					}
				}
			}
		}

		if shared.IsInteger(mergedItem["user_profile_id"]) && mergedItem["user_profile_id"] != 0 {
			var relatedUserProfile = shared.FetchByProperty(
				"user_profile",
				"Id",
				mergedItem["user_profile_id"],
			)

			if len(relatedUserProfile) > 0 {
				for _, userProfileData := range relatedUserProfile {
					var userProfile = shared.WriteStructToInterface(userProfileData)

					mergedItem["user_profile"] = map[string]interface{}{
						"id":    userProfile["id"],
						"title": userProfile["first_name"].(string) + " " + userProfile["last_name"].(string),
					}
				}
			}
		}

		if shared.IsInteger(publicProcurementId) && publicProcurementId > 0 {

			// get all with publicProcurementId in public_procurement_contract.json
			mergedItem["articles"] = GetProcurementArticles(publicProcurementId)
			// check articles exist in order_procurement_article
			if articles, ok := mergedItem["articles"].([]interface{}); ok {
				for _, item := range articles {
					if article, ok := item.(structs.OrderArticleItem); ok {
						var relatedOrderProcurementArticle = shared.FetchByProperty(
							"order_procurement_article",
							"ArticleId",
							article.Id,
						)
						if len(relatedOrderProcurementArticle) > 0 {
							for _, itemOrderArticle := range relatedOrderProcurementArticle {
								//if article use in other order, deduct amount to get Available articles
								if orderArticle, ok := itemOrderArticle.(*structs.OrderProcurementArticleItem); ok {
									article.TotalPrice = article.TotalPrice * (article.Amount - orderArticle.Amount) / article.Amount
									article.Amount = orderArticle.Amount
									articles = shared.FilterByProperty(articles, "Id", article.Id)
									mergedItem["articles"] = append(articles, article)
								}
							}
						}
					}
				}
			}

		}
		items = append(items, mergedItem)
	}

	return items
}

func GetProcurementArticles(publicProcurementId int) []interface{} {
	items := []interface{}{}
	var relatedPublicProcurementContract = shared.FetchByProperty(
		"public_procurement_contract",
		"PublicProcurementId",
		publicProcurementId,
	)
	// check public_procurement_contract
	if len(relatedPublicProcurementContract) > 0 {
		// for public_procurement_contract
		for _, contract := range relatedPublicProcurementContract {
			if contract, ok := contract.(*structs.PublicProcurementContract); ok {
				// get all articles_contract from public_procurement_contract_articles with public_procurement_contract_id
				var relatedPublicProcurementContractArticles = shared.FetchByProperty(
					"public_procurement_contract_articles",
					"PublicProcurementContractId",
					contract.Id,
				)
				if len(relatedPublicProcurementContractArticles) > 0 {
					// for public_procurement_contract_articles
					for _, contractArticles := range relatedPublicProcurementContractArticles {
						if contractArticles, ok := contractArticles.(*structs.PublicProcurementContractArticle); ok {
							// get article with id
							var relatedPublicProcurementArticles = shared.FetchByProperty(
								"public_procurement_articles",
								"Id",
								contractArticles.PublicProcurementArticleId,
							)
							if len(relatedPublicProcurementArticles) > 0 {
								// check is exist articles in mergedItem["articles"]

								var existArticle = shared.FindByProperty(items, "Id", contractArticles.PublicProcurementArticleId)
								// if exist current item sum amount
								if len(existArticle) > 0 {
									items = shared.FilterByProperty(items, "Id", contractArticles.PublicProcurementArticleId)
									for _, itemExistArticle := range existArticle {
										if articleExist, ok := itemExistArticle.(*structs.PublicProcurementArticle); ok {
											num, _ := strconv.Atoi(contractArticles.Amount)

											numberGrossValue := strings.Split(contractArticles.GrossValue, ".")[0]
											total_price_number, _ := strconv.Atoi(numberGrossValue)
											newItem := structs.OrderArticleItem{
												Id:            articleExist.Id,
												Description:   articleExist.Description,
												Title:         articleExist.Title,
												NetPrice:      articleExist.NetPrice,
												VatPercentage: articleExist.VatPercentage,
												Amount:        num,
												Available:     num,
												TotalPrice:    total_price_number,
												Unit:          "kom",
											}
											items = append(items, newItem)
										}
									}
								} else {
									// in not exist append item in array
									for _, article := range relatedPublicProcurementArticles {
										if articleExist, ok := article.(*structs.PublicProcurementArticle); ok {
											num, _ := strconv.Atoi(contractArticles.Amount)

											numberGrossValue := strings.Split(contractArticles.GrossValue, ".")[0]
											total_price_number, _ := strconv.Atoi(numberGrossValue)
											newItem := structs.OrderArticleItem{
												Id:            articleExist.Id,
												Description:   articleExist.Description,
												Title:         articleExist.Title,
												NetPrice:      articleExist.NetPrice,
												VatPercentage: articleExist.VatPercentage,
												Amount:        num,
												Available:     num,
												TotalPrice:    total_price_number,
												Unit:          "kom",
											}

											items = append(items, newItem)
										}
									}
								}

							}
						}
					}
				}
			}
		}
	}
	return items
}

var OrderListOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var id int
	var supplierId int
	var status string
	var search string
	var publicProcurementId int
	if params.Args["id"] == nil {
		id = 0
	} else {
		id = params.Args["id"].(int)
	}

	if params.Args["supplier_id"] == nil {
		supplierId = 0
	} else {
		supplierId = params.Args["supplier_id"].(int)
	}

	if params.Args["status"] == nil {
		status = ""
	} else {
		status = params.Args["status"].(string)
	}

	if params.Args["search"] == nil {
		search = ""
	} else {
		search = params.Args["search"].(string)
	}

	if params.Args["public_procurement_id"] == nil {
		publicProcurementId = 0
	} else {
		publicProcurementId = params.Args["public_procurement_id"].(int)
	}

	page := params.Args["page"]
	size := params.Args["size"]

	OrderListType := &structs.OrderListItem{}
	OrderListData, err := shared.ReadJson("http://localhost:8080/mocked-data/order_list.json", OrderListType)

	if err != nil {
		fmt.Printf("Fetching Order List failed because of this error - %s.\n", err)
	}

	// Populate data for each Order List with
	items = PopulateOrderListItemProperties(OrderListData, id, supplierId, status, search, publicProcurementId)

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

var OrderListInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.OrderListInsertItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	OrderListItemType := &structs.OrderListItem{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	orderListData, err := shared.ReadJson("http://localhost:8080/mocked-data/order_list.json", OrderListItemType)

	if err != nil {
		fmt.Printf("Fetching Order List failed because of this error - %s.\n", err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		orderListData = shared.FilterByProperty(orderListData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	totalPrice := 0

	var interfaceArticles []interface{}

	OrderProcurementArticleItemType := &structs.OrderProcurementArticleItem{}
	orderProcurementArticleData, err := shared.ReadJson("http://localhost:8080/mocked-data/order_procurement_article.json", OrderProcurementArticleItemType)

	if err != nil {
		fmt.Printf("Fetching Order List failed because of this error - %s.\n", err)
	}

	for _, article := range data.Articles {
		interfaceArticles = append(interfaceArticles, article)

		newItem := structs.OrderProcurementArticleItem{
			Id:        shared.GetRandomNumber(),
			OrderId:   data.Id,
			ArticleId: article.Id,
			Amount:    article.Amount,
		}

		orderProcurementArticleData = append(orderProcurementArticleData, newItem)
	}
	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/order_procurement_article.json"), orderProcurementArticleData)

	if len(interfaceArticles) > 0 {
		var relatedPublicProcurementContract = shared.FetchByProperty(
			"public_procurement_contract",
			"PublicProcurementId",
			data.PublicProcurementId,
		)
		// check public_procurement_contract
		if len(relatedPublicProcurementContract) > 0 {
			for _, contract := range relatedPublicProcurementContract {
				if contract, ok := contract.(*structs.PublicProcurementContract); ok {
					var relatedPublicProcurementContractArticles = shared.FetchByProperty(
						"public_procurement_contract_articles",
						"PublicProcurementContractId",
						contract.Id,
					)
					if len(relatedPublicProcurementContractArticles) > 0 {
						// for public_procurement_contract_articles
						for _, contractArticles := range relatedPublicProcurementContractArticles {
							if contractArticles, ok := contractArticles.(*structs.PublicProcurementContractArticle); ok {
								articles := shared.FindByProperty(interfaceArticles, "Id", contractArticles.Id)
								if len(articles) > 0 {
									for _, article := range articles {
										if a, ok := article.(*structs.OrderArticleInsertItem); ok {
											numberGrossValue := strings.Split(contractArticles.GrossValue, ".")[0]
											numGrossValue, _ := strconv.Atoi(numberGrossValue)
											numberAmount := strings.Split(contractArticles.Amount, ".")[0]
											numAmount, _ := strconv.Atoi(numberAmount)
											totalPrice = totalPrice + numGrossValue/numAmount*a.Amount
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	newItem := structs.OrderListItem{
		Id:                  data.Id,
		DataOrder:           data.DataOrder,
		TotalPrice:          totalPrice,
		PublicProcurementId: data.PublicProcurementId,
		SupplierId:          data.SupplierId,
		Status:              "Created",
	}

	var updatedData = append(orderListData, newItem)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/order_list.json"), updatedData)

	sliceData := []interface{}{data}

	// Populate data for each Order List
	var populatedData = PopulateOrderListItemProperties(sliceData, itemId, 0, "", "", 0)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    populatedData[0],
	}, nil
}

var OrderProcurementAvailableResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []interface{}
	var total int
	var publicProcurementId int

	if params.Args["public_procurement_id"] == nil {
		publicProcurementId = 0
	} else {
		publicProcurementId = params.Args["public_procurement_id"].(int)
	}

	items = GetProcurementArticles(publicProcurementId)

	for _, item := range items {
		if article, ok := item.(structs.OrderArticleItem); ok {
			var relatedOrderProcurementArticle = shared.FetchByProperty(
				"order_procurement_article",
				"ArticleId",
				article.Id,
			)
			if len(relatedOrderProcurementArticle) > 0 {
				for _, itemOrderArticle := range relatedOrderProcurementArticle {
					//if article use in other order, deduct amount to get Available articles
					if orderArticle, ok := itemOrderArticle.(*structs.OrderProcurementArticleItem); ok {
						article.TotalPrice = article.TotalPrice * (article.Available - orderArticle.Amount) / article.Available
						article.Available = article.Available - orderArticle.Amount
						items = shared.FilterByProperty(items, "Id", article.Id)
						if article.Amount > 0 {
							items = append(items, article)
						}
					}
				}
			}
		}
	}

	total = len(items)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"total":   total,
		"items":   items,
	}, nil
}

var OrderListReceiveResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.OrderReceiveItem

	dataBytes, _ := json.Marshal(params.Args["data"])
	OrderListType := &structs.OrderListItem{}

	_ = json.Unmarshal(dataBytes, &data)

	orderListData, err := shared.ReadJson("http://localhost:8080/mocked-data/order_list.json", OrderListType)

	if err != nil {
		fmt.Printf("Fetching Order List failed because of this error - %s.\n", err)
	}

	order := shared.FindByProperty(orderListData, "Id", data.OrderId)
	orderListData = shared.FilterByProperty(orderListData, "Id", data.OrderId)
	newItem := structs.OrderListItem{}
	for _, item := range order {
		if updateOrder, ok := item.(structs.OrderListItem); ok {
			newItem.Id = updateOrder.Id
			newItem.DataOrder = updateOrder.DataOrder
			newItem.TotalPrice = updateOrder.TotalPrice
			newItem.PublicProcurementId = updateOrder.PublicProcurementId
			newItem.SupplierId = updateOrder.SupplierId
			newItem.Status = "Received"
			newItem.DateSystem = updateOrder.DateSystem
			newItem.InvoiceDate = updateOrder.InvoiceDate
			newItem.InvoiceNumber = updateOrder.InvoiceNumber
		}
	}

	var updatedData = append(orderListData, newItem)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/order_list.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You Receive this order!",
	}, nil
}

var OrderListAssetMovementResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.OrderAssetMovementItem

	dataBytes, _ := json.Marshal(params.Args["data"])
	OrderListType := &structs.OrderListItem{}

	_ = json.Unmarshal(dataBytes, &data)

	orderListData, err := shared.ReadJson("http://localhost:8080/mocked-data/order_list.json", OrderListType)

	if err != nil {
		fmt.Printf("Fetching Order List failed because of this error - %s.\n", err)
	}

	order := shared.FindByProperty(orderListData, "Id", data.OrderId)
	orderListData = shared.FilterByProperty(orderListData, "Id", data.OrderId)
	newItem := structs.OrderListItem{}
	for _, item := range order {
		if updateOrder, ok := item.(structs.OrderListItem); ok {
			newItem.Id = updateOrder.Id
			newItem.DataOrder = updateOrder.DataOrder
			newItem.TotalPrice = updateOrder.TotalPrice
			newItem.PublicProcurementId = updateOrder.PublicProcurementId
			newItem.SupplierId = updateOrder.SupplierId
			newItem.Status = "Movement"
			newItem.DateSystem = updateOrder.DateSystem
			newItem.InvoiceDate = updateOrder.InvoiceDate
			newItem.InvoiceNumber = updateOrder.InvoiceNumber
			newItem.OrganizationUnitId = data.OrganizationUnitId
			newItem.UserProfileId = data.UserProfileId
		}
	}

	var updatedData = append(orderListData, newItem)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/order_list.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You Asset Movement this order!",
	}, nil
}
