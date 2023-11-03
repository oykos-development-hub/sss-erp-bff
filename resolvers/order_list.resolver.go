package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

func GetProcurementArticles(context context.Context, publicProcurementId int) ([]structs.OrderArticleItem, error) {
	items := []structs.OrderArticleItem{}
	itemsMap := make(map[int]structs.OrderArticleItem)

	relatedContractsResponse, err := getProcurementContractsList(&dto.GetProcurementContractsInput{
		ProcurementID: &publicProcurementId,
	})
	if err != nil {
		return nil, err
	}

	for _, contract := range relatedContractsResponse.Data {
		if err := processContract(context, &items, itemsMap, contract); err != nil {
			return nil, err
		}
	}

	return items, nil
}

func processContract(context context.Context, items *[]structs.OrderArticleItem, itemsMap map[int]structs.OrderArticleItem, contract *structs.PublicProcurementContract) error {
	relatedContractArticlesResponse, err := getProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
		ContractID: &contract.Id,
	})
	if err != nil {
		return err
	}

	for _, contractArticle := range relatedContractArticlesResponse.Data {
		if err := processContractArticle(context, items, itemsMap, contractArticle); err != nil {
			return err
		}
	}

	return nil
}

func processContractArticle(context context.Context, items *[]structs.OrderArticleItem, itemsMap map[int]structs.OrderArticleItem, contractArticle *structs.PublicProcurementContractArticle) error {
	organizationUnitID, _ := context.Value(config.OrganizationUnitIDKey).(*int)

	relatedPublicProcurementArticle, err := getProcurementArticle(contractArticle.PublicProcurementArticleId)
	if err != nil {
		return err
	}

	resProcurementArticle, err := buildProcurementArticleResponseItem(context, relatedPublicProcurementArticle, organizationUnitID)
	if err != nil {
		return err
	}

	amount := resProcurementArticle.Amount
	// all organization units
	if organizationUnitID != nil && *organizationUnitID == 0 {
		amount = resProcurementArticle.TotalAmount
	}

	overageList, err := getProcurementContractArticleOverageList(&dto.GetProcurementContractArticleOverageInput{
		ContractArticleID:  &contractArticle.Id,
		OrganizationUnitID: organizationUnitID,
	})
	if err != nil {
		return err
	}

	overageTotal := 0
	for _, item := range overageList {
		overageTotal += item.Amount
	}

	if existingItem, exists := itemsMap[contractArticle.PublicProcurementArticleId]; exists {
		// Update the existing item
		existingItem.Amount += amount
		existingItem.Available += amount + overageTotal
		existingItem.TotalPrice += contractArticle.GrossValue
	} else {
		// Add new item
		newItem := structs.OrderArticleItem{
			Id:            relatedPublicProcurementArticle.Id,
			Description:   relatedPublicProcurementArticle.Description,
			Title:         relatedPublicProcurementArticle.Title,
			NetPrice:      relatedPublicProcurementArticle.NetPrice,
			VatPercentage: relatedPublicProcurementArticle.VatPercentage,
			Amount:        amount,
			Available:     amount + overageTotal,
			TotalPrice:    contractArticle.GrossValue,
			Unit:          "kom",
			Manufacturer:  relatedPublicProcurementArticle.Manufacturer,
		}

		*items = append(*items, newItem)
		itemsMap[newItem.Id] = newItem
	}

	return nil
}

var OrderListOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []dto.OrderListOverviewResponse
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	supplierId := params.Args["supplier_id"]
	publicProcurementID := params.Args["public_procurement_id"]
	status, statusOK := params.Args["status"].(string)
	search, searchOk := params.Args["search"].(string)

	if id != nil && shared.IsInteger(id) && id != 0 {
		orderList, err := getOrderListById(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}

		orderListItem, err := buildOrderListResponseItem(params.Context, orderList)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		items = []dto.OrderListOverviewResponse{*orderListItem}
		total = 1
	} else {
		input := dto.GetOrderListInput{}

		organizationUnitID, unitOK := params.Context.Value(config.OrganizationUnitIDKey).(*int)

		if unitOK && organizationUnitID != nil {
			input.OrganizationUnitId = organizationUnitID
		}
		if shared.IsInteger(page) && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if shared.IsInteger(size) && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if shared.IsInteger(supplierId) && supplierId.(int) > 0 {
			supplierId := supplierId.(int)
			input.SupplierID = &supplierId
		}
		if shared.IsInteger(publicProcurementID) && publicProcurementID.(int) > 0 {
			publicProcurementID := publicProcurementID.(int)
			input.PublicProcurementID = &publicProcurementID
		}
		if statusOK && status != "" {
			input.Status = &status
		}
		if searchOk && search != "" {
			input.Search = &search
		}

		orderLists, err := getOrderLists(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, orderList := range orderLists.Data {
			orderListItem, err := buildOrderListResponseItem(params.Context, &orderList)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			items = append(items, *orderListItem)
		}
		total = orderLists.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
	}, nil
}

var OrderListInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OrderListInsertItem
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	itemId := data.Id

	listInsertItem, err := buildOrderListInsertItem(params.Context, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		listInsertItem.Status = "Updated"
		res, err := updateOrderListItem(itemId, listInsertItem)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		if len(data.Articles) > 0 {
			err := deleteOrderArticles(itemId)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			err = createOrderListProcurementArticles(res.Id, data)
			if err != nil {
				return shared.HandleAPIError(err)
			}
		}

		item, err := buildOrderListResponseItem(params.Context, res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		listInsertItem.Status = "Created"
		res, err := createOrderListItem(listInsertItem)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		err = createOrderListProcurementArticles(res.Id, data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		item, err := buildOrderListResponseItem(params.Context, res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

var OrderProcurementAvailableResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []structs.OrderArticleItem
		total int
	)
	publicProcurementID, ok := params.Args["public_procurement_id"].(int)
	if !ok || publicProcurementID <= 0 {
		return shared.ErrorResponse("You must pass the item procurement id"), nil
	}

	ctx := params.Context

	if params.Args["organization_unit_id"] != nil {
		organizationUnitID := params.Args["organization_unit_id"].(int)
		ctx = context.WithValue(ctx, config.OrganizationUnitIDKey, &organizationUnitID)
	}

	articles, err := GetProcurementArticles(ctx, publicProcurementID)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, item := range articles {
		currentArticle := item // work with a copy to avoid modifying the range variable
		currentArticle.Price = currentArticle.NetPrice
		getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
			ArticleID: &currentArticle.Id,
		}

		relatedOrderProcurementArticleResponse, err := getOrderProcurementArticles(&getOrderProcurementArticleInput)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		if relatedOrderProcurementArticleResponse.Total > 0 {
			for _, orderArticle := range relatedOrderProcurementArticleResponse.Data {
				// if article is used in another order, deduct the amount to get Available articles
				currentArticle.TotalPrice = currentArticle.TotalPrice * float32(currentArticle.Available-orderArticle.Amount/currentArticle.Available)
				currentArticle.TotalPrice = currentArticle.TotalPrice * float32(currentArticle.Available-orderArticle.Amount/currentArticle.Available)
				currentArticle.Available -= orderArticle.Amount
			}
		}
		items = append(items, currentArticle)
	}

	total = len(items)

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
	}, nil
}

var OrderListAssetMovementResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OrderAssetMovementItem

	dataBytes, _ := json.Marshal(params.Args["data"])
	_ = json.Unmarshal(dataBytes, &data)

	orderList, err := getOrderListById(data.OrderId)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	orderList.Status = "Movement"
	orderList.RecipientUserId = &data.RecipientUserId
	orderList.OfficeId = &data.OfficeId

	_, err = updateOrderListItem(data.OrderId, orderList)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You Asset Movement this order!",
	}, nil
}

var RecipientUsersResolver = func(params graphql.ResolveParams) (interface{}, error) {
	organizationUnitID, _ := params.Context.Value(config.OrganizationUnitIDKey).(*int)

	var userProfileDropdownList []*dto.DropdownSimple

	if organizationUnitID == nil {
		return dto.Response{
			Message: "User has no organization unit assigned!",
			Status:  "success",
			Items:   userProfileDropdownList,
			Total:   0,
		}, nil
	}

	employees, err := getEmployeesOfOrganizationUnit(*organizationUnitID)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	for _, employee := range employees {
		userProfileDropdownList = append(userProfileDropdownList, &dto.DropdownSimple{
			Id:    employee.Id,
			Title: employee.GetFullName(),
		})
	}

	return dto.Response{
		Message: "Here's the list you asked for!",
		Status:  "success",
		Items:   userProfileDropdownList,
		Total:   len(userProfileDropdownList),
	}, nil
}

func getEmployeesOfOrganizationUnit(id int) ([]*structs.UserProfiles, error) {
	var userProfileList []*structs.UserProfiles
	active := 2
	systematizationsRes, err := getSystematizations(&dto.GetSystematizationsInput{Active: &active, OrganizationUnitID: &id})
	if err != nil {
		return nil, err
	}
	if len(systematizationsRes.Data) == 0 {
		return nil, errors.New("no systematization")
	}
	systematization := systematizationsRes.Data[0]
	jobPositionsInOrganizationUnit, err := getJobPositionsInOrganizationUnits(&dto.GetJobPositionInOrganizationUnitsInput{SystematizationID: &systematization.Id})
	if err != nil {
		return nil, err
	}

	for _, jobPosition := range jobPositionsInOrganizationUnit.Data {
		employeesByJobPosition, err := getEmployeesInOrganizationUnitList(&dto.GetEmployeesInOrganizationUnitInput{PositionInOrganizationUnit: &jobPosition.Id})
		if err != nil {
			return nil, err
		}

		for _, employee := range employeesByJobPosition {
			userProfile, err := getUserProfileById(employee.UserProfileId)
			if err != nil {
				return nil, err
			}
			userProfileList = append(userProfileList, userProfile)
		}
	}

	return userProfileList, nil
}

func deleteOrderArticles(itemId int) error {
	orderProcurementArticlesResponse, err := getOrderProcurementArticles(&dto.GetOrderProcurementArticleInput{
		OrderID: &itemId,
	})
	if err != nil {
		return err
	}

	for _, orderProcurementArticle := range orderProcurementArticlesResponse.Data {
		err = deleteOrderProcurementArticle(orderProcurementArticle.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

var OrderListDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteOrderArticles(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	err = deleteOrderList(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

var OrderListReceiveResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OrderReceiveItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	_ = json.Unmarshal(dataBytes, &data)

	orderList, err := getOrderListById(data.OrderId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	orderList.Status = "Receive"
	orderList.InvoiceNumber = &data.InvoiceNumber
	orderList.DateSystem = &data.DateSystem
	orderList.InvoiceDate = &data.InvoiceDate
	if data.Description != nil {
		orderList.Description = data.Description
	}

	_, err = updateOrderListItem(data.OrderId, orderList)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You received this order!",
	}, nil
}

var OrderListReceiveDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	orderList, err := getOrderListById(id)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	orderList.Status = "Created"
	orderList.DateSystem = nil
	orderList.InvoiceDate = nil
	orderList.InvoiceNumber = nil
	orderList.OfficeId = nil
	orderList.RecipientUserId = nil
	orderList.Description = nil

	_, err = updateOrderListItem(id, orderList)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted received order!",
	}, nil
}

var OrderListAssetMovementDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	orderList, err := getOrderListById(id)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	orderList.Status = "Received"
	orderList.RecipientUserId = nil
	orderList.OfficeId = nil

	_, err = updateOrderListItem(id, orderList)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You Asset Movement this order!",
	}, nil
}

func updateOrderListItem(id int, orderListItem *structs.OrderListItem) (*structs.OrderListItem, error) {
	res := &dto.GetOrderListResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.ORDER_LISTS_ENDPOINT+"/"+strconv.Itoa(id), orderListItem, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createOrderListItem(orderListItem *structs.OrderListItem) (*structs.OrderListItem, error) {
	res := &dto.GetOrderListResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ORDER_LISTS_ENDPOINT, orderListItem, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createOrderProcurementArticle(orderProcurementArticleItem *structs.OrderProcurementArticleItem) (*structs.OrderProcurementArticleItem, error) {
	res := &dto.GetOrderProcurementArticleResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ORDER_PROCUREMENT_ARTICLES_ENDPOINT, orderProcurementArticleItem, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteOrderProcurementArticle(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.ORDER_PROCUREMENT_ARTICLES_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func deleteOrderList(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.ORDER_LISTS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func createOrderListProcurementArticles(orderListId int, data structs.OrderListInsertItem) error {
	for _, article := range data.Articles {
		newArticle := structs.OrderProcurementArticleItem{
			OrderId:   orderListId,
			ArticleId: article.Id,
			Amount:    article.Amount,
		}
		_, err := createOrderProcurementArticle(&newArticle)
		if err != nil {
			return err
		}
	}
	return nil
}

func buildOrderListInsertItem(context context.Context, item *structs.OrderListInsertItem) (*structs.OrderListItem, error) {
	currentTime := time.Now().UTC()
	timeString := currentTime.Format("2006-01-02T15:04:05Z07:00")

	var newItem *structs.OrderListItem

	totalPrice := float32(0.0)
	var supplierId *int

	if len(item.Articles) > 0 {
		articleMap := make(map[int]structs.OrderArticleInsertItem)
		for _, article := range item.Articles {
			articleMap[article.Id] = article
		}

		relatedContractsResponse, err := getProcurementContractsList(&dto.GetProcurementContractsInput{
			ProcurementID: &item.PublicProcurementId,
		})
		if err != nil {
			return nil, err
		}

		for _, contract := range relatedContractsResponse.Data {
			supplierId = &contract.SupplierId
			relatedContractArticlesResponse, err := getProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
				ContractID: &contract.Id,
			})
			if err != nil {
				return nil, err
			}

			for _, contractArticle := range relatedContractArticlesResponse.Data {
				if article, exists := articleMap[contractArticle.PublicProcurementArticleId]; exists {
					totalPrice += (contractArticle.GrossValue) * float32(article.Amount)
				}
			}
		}
	}

	newItem = &structs.OrderListItem{
		Id:                  item.Id,
		DateOrder:           timeString,
		TotalPrice:          totalPrice,
		PublicProcurementId: item.PublicProcurementId,
		SupplierId:          supplierId,
	}

	// Getting organizationUnitId from job position
	loggedInProfile, _ := context.Value(config.LoggedInProfileKey).(*structs.UserProfiles)
	organizationUnitID, unitOK := context.Value(config.OrganizationUnitIDKey).(*int)

	newItem.RecipientUserId = &loggedInProfile.Id

	if unitOK && organizationUnitID != nil {
		newItem.OrganizationUnitId = *organizationUnitID
	}

	return newItem, nil
}

func getOrderProcurementArticles(input *dto.GetOrderProcurementArticleInput) (*dto.GetOrderProcurementArticlesResponseMS, error) {
	res := &dto.GetOrderProcurementArticlesResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ORDER_PROCUREMENT_ARTICLES_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func buildOrderListResponseItem(context context.Context, item *structs.OrderListItem) (*dto.OrderListOverviewResponse, error) {
	totalBruto := float32(0.0)
	totalNeto := float32(0.0)

	procurementItem, err := getProcurementItem(item.PublicProcurementId)
	if err != nil {
		return nil, err
	}

	office := &dto.DropdownSimple{}
	if item.OfficeId != nil {
		officeItem, _ := getDropdownSettingById(*item.OfficeId)
		office.Title = officeItem.Title
		office.Id = officeItem.Id
	}

	// getting articles and total price
	articles := []dto.DropdownProcurementAvailableArticle{}
	getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
		OrderID: &item.Id,
	}
	relatedOrderProcurementArticle, err := getOrderProcurementArticles(&getOrderProcurementArticleInput)
	if err != nil {
		return nil, err
	}

	publicProcurementArticles, err := GetProcurementArticles(context, item.PublicProcurementId)
	if err != nil {
		return nil, err
	}

	publicProcurementArticlesMap := make(map[int]structs.OrderArticleItem)
	for _, article := range publicProcurementArticles {
		publicProcurementArticlesMap[article.Id] = article
	}

	for _, itemOrderArticle := range relatedOrderProcurementArticle.Data {
		if article, exists := publicProcurementArticlesMap[itemOrderArticle.ArticleId]; exists {
			/*articleUnitPrice := article.TotalPrice / float32(article.Amount)
			articleTotalPrice := articleUnitPrice * float32(itemOrderArticle.Amount)
			articleVat, _ := strconv.ParseFloat(article.VatPercentage, 32)
			articleVat32 := float32(articleVat)
			vatPrice := articleTotalPrice * articleVat32 / 100
			totalPrice += articleTotalPrice
			totalNeto += vatPrice*/
			articleVat, _ := strconv.ParseFloat(article.VatPercentage, 32)
			articleVat32 := float32(articleVat)
			articleUnitPrice := article.NetPrice + article.NetPrice*articleVat32/100
			articleTotalPrice := articleUnitPrice * float32(itemOrderArticle.Amount)
			totalBruto += articleTotalPrice
			vat := articleTotalPrice * (100 - articleVat32) / 100
			totalNeto += vat

			articles = append(articles, dto.DropdownProcurementAvailableArticle{
				Id:           itemOrderArticle.Id,
				Title:        article.Title,
				Manufacturer: article.Manufacturer,
				Description:  article.Description,
				Unit:         article.Unit,
				Available:    article.Available,
				Amount:       itemOrderArticle.Amount,
				TotalPrice:   articleTotalPrice,
				Price:        articleUnitPrice,
			})
		}
	}

	res := dto.OrderListOverviewResponse{
		Id:                  item.Id,
		DateOrder:           (string)(item.DateOrder),
		TotalNeto:           totalNeto,
		TotalBruto:          totalBruto,
		PublicProcurementID: procurementItem.Id,
		PublicProcurement: &dto.DropdownSimple{
			Id:    procurementItem.Id,
			Title: procurementItem.Title,
		},
		DateSystem:         item.DateSystem,
		InvoiceDate:        item.InvoiceDate,
		OrganizationUnitID: item.OrganizationUnitId,
		OfficeID:           office.Id,
		Office:             office,
		Description:        item.Description,
		Status:             item.Status,
		Articles:           &articles,
	}

	if item.RecipientUserId != nil {
		userProfile, err := getUserProfileById(*item.RecipientUserId)
		if err != nil {
			return nil, err
		}
		res.RecipientUser = &dto.DropdownSimple{
			Id:    userProfile.Id,
			Title: userProfile.GetFullName(),
		}
		res.RecipientUserID = item.RecipientUserId
	}

	if item.SupplierId != nil {
		supplier, err := getSupplier(*item.SupplierId)
		if err != nil {
			return nil, err
		}
		res.SupplierID = supplier.Id
		res.Supplier = &dto.DropdownSimple{
			Id:    supplier.Id,
			Title: supplier.Title,
		}
	}

	if item.InvoiceNumber != nil {
		res.InvoiceNumber = *item.InvoiceNumber
	}

	return &res, nil
}

func getOrderListById(id int) (*structs.OrderListItem, error) {
	res := &dto.GetOrderListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ORDER_LISTS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getOrderLists(input *dto.GetOrderListInput) (*dto.GetOrderListsResponseMS, error) {
	res := &dto.GetOrderListsResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ORDER_LISTS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
