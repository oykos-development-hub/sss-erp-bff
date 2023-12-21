package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

// processContractArticle refactored to only take context and contractArticle, and return an OrderArticleItem.
func processContractArticle(r repository.MicroserviceRepositoryInterface, ctx context.Context, contractArticle *structs.PublicProcurementContractArticle) (structs.OrderArticleItem, error) {
	organizationUnitID, _ := ctx.Value(config.OrganizationUnitIDKey).(*int)

	// Get the related public procurement article details.
	relatedPublicProcurementArticle, err := r.GetProcurementArticle(contractArticle.PublicProcurementArticleId)
	if err != nil {
		return structs.OrderArticleItem{}, err
	}

	// Build response item based on the related article and possibly organization unit.
	resProcurementArticle, err := buildProcurementArticleResponseItem(r, ctx, relatedPublicProcurementArticle, organizationUnitID)
	if err != nil {
		return structs.OrderArticleItem{}, err
	}

	// Determine the amount based on the organization unit.
	amount := resProcurementArticle.Amount
	if organizationUnitID != nil && *organizationUnitID == 0 {
		amount = resProcurementArticle.TotalAmount
	}

	// Get overages for the contract article.
	overageList, err := r.GetProcurementContractArticleOverageList(&dto.GetProcurementContractArticleOverageInput{
		ContractArticleID:  &contractArticle.Id,
		OrganizationUnitID: organizationUnitID,
	})
	if err != nil {
		return structs.OrderArticleItem{}, err
	}

	// Calculate the total overage amount.
	overageTotal := 0
	for _, item := range overageList {
		overageTotal += item.Amount
	}

	// Build the new item with the calculated amounts.
	newItem := structs.OrderArticleItem{
		Id:             relatedPublicProcurementArticle.Id,
		Description:    relatedPublicProcurementArticle.Description,
		Title:          relatedPublicProcurementArticle.Title,
		NetPrice:       relatedPublicProcurementArticle.NetPrice,
		VatPercentage:  relatedPublicProcurementArticle.VatPercentage,
		Amount:         amount,
		Available:      amount + overageTotal,
		TotalPrice:     contractArticle.GrossValue,
		Unit:           "kom",
		Manufacturer:   relatedPublicProcurementArticle.Manufacturer,
		VisibilityType: relatedPublicProcurementArticle.VisibilityType,
	}

	return newItem, nil
}

// GetProcurementArticles simplified to utilize the refactored processContractArticle function.
func GetProcurementArticles(r repository.MicroserviceRepositoryInterface, ctx context.Context, publicProcurementId int) ([]structs.OrderArticleItem, error) {
	var items []structs.OrderArticleItem
	itemsMap := make(map[int]structs.OrderArticleItem)

	// Get related contracts.
	relatedContractsResponse, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{
		ProcurementID: &publicProcurementId,
	})
	if err != nil {
		return nil, err
	}

	// Process each contract.
	for _, contract := range relatedContractsResponse.Data {
		relatedContractArticlesResponse, err := r.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
			ContractID: &contract.Id,
		})
		if err != nil {
			return nil, err
		}

		// Process each contract article.
		for _, contractArticle := range relatedContractArticlesResponse.Data {
			newItem, err := processContractArticle(r, ctx, contractArticle)
			if err != nil {
				return nil, err
			}

			if existingItem, exists := itemsMap[newItem.Id]; exists {
				// Update the existing item in the map if it exists.
				existingItem.Amount += newItem.Amount
				existingItem.Available += newItem.Available
				existingItem.TotalPrice += newItem.TotalPrice
				itemsMap[newItem.Id] = existingItem
			} else {
				// Add new item to the map and slice.
				itemsMap[newItem.Id] = newItem
			}
		}
	}

	// Convert map to slice.
	items = make([]structs.OrderArticleItem, 0, len(itemsMap))
	for _, item := range itemsMap {
		items = append(items, item)
	}

	return items, nil
}

func (r *Resolver) OrderListOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []dto.OrderListOverviewResponse
		total int
		price float32
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	supplierId := params.Args["supplier_id"]
	publicProcurementID := params.Args["public_procurement_id"]
	status, statusOK := params.Args["status"].(string)
	search, searchOk := params.Args["search"].(string)
	activePlan, _ := params.Args["active_plan"].(bool)
	year, yearOK := params.Args["year"].(string)
	sortByTotalPrice, sortByTotalPriceOK := params.Args["sort_by_total_price"].(string)
	sortByDateOrder, sortByDateOrderOK := params.Args["sort_by_date_order"].(string)

	if id != nil && shared.IsInteger(id) && id != 0 {
		orderList, err := r.Repo.GetOrderListById(id.(int))
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		orderListItem, err := buildOrderListResponseItem(r.Repo, params.Context, orderList)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		items = []dto.OrderListOverviewResponse{*orderListItem}
		total = 1
	} else if activePlan {
		inputPlans := dto.GetProcurementPlansInput{}
		plans, err := r.Repo.GetProcurementPlanList(&inputPlans)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		currentYear := time.Now().Year()
		inputOrderList := dto.GetOrderListInput{}
		for _, plan := range plans {

			if plan.Year <= strconv.Itoa(currentYear) {
				item, _ := buildProcurementPlanResponseItem(r.Repo, params.Context, plan, nil, &dto.GetProcurementItemListInputMS{})

				if item.Status == dto.PlanStatusPostBudgetClosed {
					if len(item.Items) > 0 {
						for _, procurement := range item.Items {
							inputOrderList.PublicProcurementID = &procurement.Id
							orderLists, err := r.Repo.GetOrderLists(&inputOrderList)
							if err != nil {
								return apierrors.HandleAPIError(err)
							}
							for _, orderList := range orderLists.Data {
								if orderList.IsUsed {
									continue
								}
								orderListItem, err := buildOrderListResponseItem(r.Repo, params.Context, &orderList)
								if err != nil {
									return apierrors.HandleAPIError(err)
								}

								items = append(items, *orderListItem)
							}
						}
					}

					break
				}
			}
		}
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

		if yearOK && year != "" {
			input.Year = &year
		}

		if sortByDateOrderOK && sortByDateOrder != "" {
			input.SortByDateOrder = &sortByDateOrder
		}

		if sortByTotalPriceOK && sortByTotalPrice != "" {
			input.SortByTotalPrice = &sortByTotalPrice
		}

		orderLists, err := r.Repo.GetOrderLists(&input)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		for _, orderList := range orderLists.Data {
			orderListItem, err := buildOrderListResponseItem(r.Repo, params.Context, &orderList)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
			items = append(items, *orderListItem)
			price += orderListItem.TotalBruto
		}
		total = orderLists.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
		Price:   price,
	}, nil
}

func (r *Resolver) OrderListInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OrderListInsertItem
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	itemId := data.Id

	listInsertItem, err := buildOrderListInsertItem(r.Repo, params.Context, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		listInsertItem.Status = "Updated"
		res, err := r.Repo.UpdateOrderListItem(itemId, listInsertItem)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		if len(data.Articles) > 0 {
			err := deleteOrderArticles(r.Repo, itemId)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
			err = r.Repo.CreateOrderListProcurementArticles(res.Id, data)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
		}

		item, err := buildOrderListResponseItem(r.Repo, params.Context, res)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		listInsertItem.Status = "Created"
		listInsertItem.IsUsed = false
		res, err := r.Repo.CreateOrderListItem(listInsertItem)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		err = r.Repo.CreateOrderListProcurementArticles(res.Id, data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		item, err := buildOrderListResponseItem(r.Repo, params.Context, res)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func (r *Resolver) OrderProcurementAvailableResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []structs.OrderArticleItem
		total int
	)
	publicProcurementID, ok := params.Args["public_procurement_id"].(int)
	if !ok || publicProcurementID <= 0 {
		return apierrors.HandleAPIError(errors.New("You must pass the item procurement id"))
	}

	visibilityType := params.Args["visibility_type"]

	ctx := params.Context
	var organizationUnitID int

	if params.Args["organization_unit_id"] != nil && params.Args["organization_unit_id"].(int) != 0 {
		organizationUnitID = params.Args["organization_unit_id"].(int)
		ctx = context.WithValue(ctx, config.OrganizationUnitIDKey, &organizationUnitID)
	} else {
		/*organizationUnitId, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitId == nil {
			return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
		}

		organizationUnitID = *organizationUnitId*/
		organizationUnitID = 0
	}

	articles, err := GetProcurementArticles(r.Repo, ctx, publicProcurementID)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	for _, item := range articles {
		if visibilityType != nil && visibilityType.(int) > 0 && visibilityType.(int) != int(item.VisibilityType) {
			continue
		}
		processedArticle, err := ProcessOrderArticleItem(r.Repo, ctx, item, organizationUnitID)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		items = append(items, processedArticle)

	}

	total = len(items)

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
	}, nil
}

// ProcessOrderArticleItem processes a single order article item to calculate its available amount and total price
func ProcessOrderArticleItem(r repository.MicroserviceRepositoryInterface, ctx context.Context, article structs.OrderArticleItem, organizationUnitID int) (structs.OrderArticleItem, error) {
	var err error
	currentArticle := article // work with a copy to avoid modifying the original

	articleVat, _ := strconv.ParseFloat(article.VatPercentage, 32)
	articleVat32 := float32(articleVat)
	currentArticle.Price = article.NetPrice + article.NetPrice*(articleVat32/100)

	getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
		ArticleID: &currentArticle.Id,
	}

	relatedOrderProcurementArticleResponse, err := r.GetOrderProcurementArticles(&getOrderProcurementArticleInput)
	if err != nil {
		return currentArticle, err
	}

	if relatedOrderProcurementArticleResponse.Total > 0 {
		for _, orderArticle := range relatedOrderProcurementArticleResponse.Data {
			order, err := r.GetOrderListById(orderArticle.OrderId)

			if err != nil {
				return currentArticle, err
			}

			if organizationUnitID > 0 && order.OrganizationUnitId == organizationUnitID {
				// if article is used in another order, deduct the amount to get Available articles
				currentArticle.TotalPrice *= float32(currentArticle.Available-orderArticle.Amount) / float32(currentArticle.Available)
				currentArticle.Available -= orderArticle.Amount
			}
		}
	}

	if organizationUnitID == 0 {
		articles, err := r.GetOrganizationUnitArticlesList(dto.GetProcurementOrganizationUnitArticleListInputDTO{
			ArticleID: &currentArticle.Id,
		})
		if err != nil {
			return currentArticle, nil
		}
		amount := 0
		for _, article := range articles {
			amount += article.Amount
		}

		for _, article := range relatedOrderProcurementArticleResponse.Data {
			amount -= article.Amount
		}
		currentArticle.Available = amount
	}

	articleInventory, err := r.GetAllInventoryItem(dto.InventoryItemFilter{
		ArticleId:          &article.Id,
		OrganizationUnitID: &organizationUnitID,
	})

	if err != nil {
		return currentArticle, err
	}

	currentArticle.Available -= len(articleInventory.Data)

	overages, err := r.GetProcurementContractArticleOverageList(&dto.GetProcurementContractArticleOverageInput{
		ContractArticleID:  &article.Id,
		OrganizationUnitID: &organizationUnitID})

	if err != nil {
		return currentArticle, err
	}

	for _, overage := range overages {
		currentArticle.Available += overage.Amount
	}

	return currentArticle, nil
}

func (r *Resolver) RecipientUsersResolver(params graphql.ResolveParams) (interface{}, error) {
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

	employees, err := GetEmployeesOfOrganizationUnit(r.Repo, *organizationUnitID)
	if err != nil {
		return apierrors.HandleAPIError(err)
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

func GetEmployeesOfOrganizationUnit(r repository.MicroserviceRepositoryInterface, id int) ([]*structs.UserProfiles, error) {
	var userProfileList []*structs.UserProfiles
	active := 2
	systematizationsRes, err := r.GetSystematizations(&dto.GetSystematizationsInput{Active: &active, OrganizationUnitID: &id})
	if err != nil {
		return nil, err
	}
	if len(systematizationsRes.Data) == 0 {
		return nil, errors.New("no systematization")
	}
	systematization := systematizationsRes.Data[0]
	jobPositionsInOrganizationUnit, err := r.GetJobPositionsInOrganizationUnits(&dto.GetJobPositionInOrganizationUnitsInput{SystematizationID: &systematization.Id})
	if err != nil {
		return nil, err
	}

	for _, jobPosition := range jobPositionsInOrganizationUnit.Data {
		employeesByJobPosition, err := r.GetEmployeesInOrganizationUnitList(&dto.GetEmployeesInOrganizationUnitInput{PositionInOrganizationUnit: &jobPosition.Id})
		if err != nil {
			return nil, err
		}

		for _, employee := range employeesByJobPosition {
			userProfile, err := r.GetUserProfileById(employee.UserProfileId)
			if err != nil {
				return nil, err
			}
			userProfileList = append(userProfileList, userProfile)
		}
	}

	return userProfileList, nil
}

func deleteOrderArticles(r repository.MicroserviceRepositoryInterface, itemId int) error {
	orderProcurementArticlesResponse, err := r.GetOrderProcurementArticles(&dto.GetOrderProcurementArticleInput{
		OrderID: &itemId,
	})
	if err != nil {
		return err
	}

	for _, orderProcurementArticle := range orderProcurementArticlesResponse.Data {
		err = r.DeleteOrderProcurementArticle(orderProcurementArticle.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Resolver) OrderListDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	orderList, err := r.Repo.GetOrderListById(itemId)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if orderList.OrderFile != nil && *orderList.OrderFile != 0 {
		err := r.Repo.DeleteFile(*orderList.OrderFile)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	}

	for _, fileID := range orderList.ReceiveFile {
		err := r.Repo.DeleteFile(fileID)
		if err != nil {
			return nil, err
		}
	}

	if orderList.MovementFile != nil && *orderList.MovementFile != 0 {
		err := r.Repo.DeleteFile(*orderList.MovementFile)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	}

	err = deleteOrderArticles(r.Repo, itemId)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	err = r.Repo.DeleteOrderList(itemId)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

func (r *Resolver) OrderListReceiveResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OrderReceiveItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	_ = json.Unmarshal(dataBytes, &data)

	orderList, err := r.Repo.GetOrderListById(data.OrderId)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	status := orderList.Status

	orderList.Status = "Receive"
	orderList.InvoiceNumber = data.InvoiceNumber
	orderList.DateSystem = &data.DateSystem
	orderList.InvoiceDate = data.InvoiceDate
	orderList.ReceiveFile = data.ReceiveFile
	if data.Description != nil {
		orderList.Description = data.Description
	}

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}
	if status != "Receive" {
		if orderList.GroupOfArticlesID != nil && *orderList.GroupOfArticlesID != 0 {
			for _, article := range data.Articles {
				orderArticle, err := r.Repo.GetOrderProcurementArticleByID(article.Id)

				if err != nil {
					return apierrors.HandleAPIError(err)
				}

				orderArticle.NetPrice = article.NetPrice
				orderArticle.VatPercentage = article.VatPercentage

				_, err = r.Repo.UpdateOrderProcurementArticle(orderArticle)

				if err != nil {
					return apierrors.HandleAPIError(err)
				}

				stock, _, _ := r.Repo.GetStock(&dto.StockFilter{
					Year:               &orderArticle.Year,
					Title:              &orderArticle.Title,
					Description:        &orderArticle.Description,
					NetPrice:           &orderArticle.NetPrice,
					VatPercentage:      &orderArticle.VatPercentage,
					OrganizationUnitID: organizationUnitID})

				err = r.Repo.AddOnStock(stock, *orderArticle, *organizationUnitID)

				if err != nil {
					return apierrors.HandleAPIError(err)
				}
			}
		} else {
			orderArticles, err := r.Repo.GetOrderProcurementArticles(&dto.GetOrderProcurementArticleInput{
				OrderID: &orderList.Id,
			})

			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			for _, orderArticle := range orderArticles.Data {
				currentArticle, err := r.Repo.GetProcurementArticle(orderArticle.ArticleId)

				if err != nil {
					return apierrors.HandleAPIError(err)
				}

				vatPercentageInt, err := strconv.Atoi(currentArticle.VatPercentage)

				if err != nil {
					return apierrors.HandleAPIError(err)
				}

				stockArticle, _, err := r.Repo.GetStock(&dto.StockFilter{
					Title:              &currentArticle.Title,
					Year:               &orderArticle.Year,
					OrganizationUnitID: organizationUnitID,
				})

				if err != nil {
					return apierrors.HandleAPIError(err)
				}

				orderArticle.Title = currentArticle.Title
				orderArticle.Description = currentArticle.Description
				orderArticle.NetPrice = currentArticle.NetPrice
				orderArticle.VatPercentage = vatPercentageInt

				err = r.Repo.AddOnStock(stockArticle, orderArticle, *organizationUnitID)

				if err != nil {
					return apierrors.HandleAPIError(err)
				}
			}
		}
	}

	_, err = r.Repo.UpdateOrderListItem(data.OrderId, orderList)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You received this order!",
	}, nil
}

func (r *Resolver) OrderListReceiveDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	orderList, err := r.Repo.GetOrderListById(id)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if orderList.MovementFile != nil && *orderList.MovementFile != 0 {
		err := r.Repo.DeleteFile(*orderList.MovementFile)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	}

	for _, fileID := range orderList.ReceiveFile {
		err := r.Repo.DeleteFile(fileID)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	}

	orderList.Status = "Created"
	orderList.DateSystem = nil
	orderList.InvoiceDate = nil
	orderList.InvoiceNumber = nil
	orderList.OfficeId = nil
	orderList.RecipientUserId = nil
	orderList.Description = nil
	orderList.MovementFile = nil
	orderList.ReceiveFile = nil

	_, err = r.Repo.UpdateOrderListItem(id, orderList)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted received order!",
	}, nil
}

func buildOrderListInsertItem(r repository.MicroserviceRepositoryInterface, context context.Context, item *structs.OrderListInsertItem) (*structs.OrderListItem, error) {
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

		relatedContractsResponse, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{
			ProcurementID: &item.PublicProcurementId,
		})
		if err != nil {
			return nil, err
		}

		for _, contract := range relatedContractsResponse.Data {
			supplierId = &contract.SupplierId
			relatedContractArticlesResponse, err := r.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
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

	if item.PublicProcurementId == 0 {
		supplierId = &item.SupplierId
	}

	newItem = &structs.OrderListItem{
		Id:                  item.Id,
		DateOrder:           timeString,
		TotalPrice:          totalPrice,
		PublicProcurementId: &item.PublicProcurementId,
		GroupOfArticlesID:   &item.GroupOfArticlesID,
		SupplierId:          supplierId,
		OrderFile:           &item.OrderFile,
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

func buildOrderListResponseItem(r repository.MicroserviceRepositoryInterface, context context.Context, item *structs.OrderListItem) (*dto.OrderListOverviewResponse, error) {

	var res dto.OrderListOverviewResponse
	var procurementDropdown dto.DropdownSimple
	var supplierDropdown dto.DropdownSimple
	articles := []dto.DropdownProcurementAvailableArticle{}
	totalBruto := float32(0.0)
	totalNeto := float32(0.0)
	zero := 0

	if item.PublicProcurementId != nil && *item.PublicProcurementId != zero {
		procurementItem, err := r.GetProcurementItem(*item.PublicProcurementId)
		if err != nil {
			return nil, err
		}

		procurementDropdown.Id = procurementItem.Id
		procurementDropdown.Title = procurementItem.Title

		contract, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &procurementItem.Id})
		if err != nil {
			return nil, err
		}

		supplier, err := r.GetSupplier(contract.Data[0].SupplierId)

		if err != nil {
			return nil, err
		}

		supplierDropdown = dto.DropdownSimple{
			Id:    supplier.Id,
			Title: supplier.Title,
		}

		// getting articles and total price
		getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
			OrderID: &item.Id,
		}
		relatedOrderProcurementArticle, err := r.GetOrderProcurementArticles(&getOrderProcurementArticleInput)
		if err != nil {
			return nil, err
		}

		publicProcurementArticles, err := GetProcurementArticles(r, context, *item.PublicProcurementId)
		if err != nil {
			return nil, err
		}

		publicProcurementArticlesMap := make(map[int]structs.OrderArticleItem)
		for _, article := range publicProcurementArticles {
			publicProcurementArticlesMap[article.Id] = article
		}

		for _, itemOrderArticle := range relatedOrderProcurementArticle.Data {
			if article, exists := publicProcurementArticlesMap[itemOrderArticle.ArticleId]; exists {
				articleVat, _ := strconv.ParseFloat(article.VatPercentage, 32)
				articleVat32 := float32(articleVat)
				articleUnitPrice := article.NetPrice + article.NetPrice*articleVat32/100
				articleTotalPrice := articleUnitPrice * float32(itemOrderArticle.Amount)
				totalBruto += articleTotalPrice
				vat := articleTotalPrice * (100 - articleVat32) / 100
				totalNeto += vat

				articles = append(articles, dto.DropdownProcurementAvailableArticle{
					Id:            itemOrderArticle.Id,
					Title:         article.Title,
					Manufacturer:  article.Manufacturer,
					Description:   article.Description,
					Unit:          article.Unit,
					Available:     article.Available,
					Amount:        itemOrderArticle.Amount,
					TotalPrice:    articleTotalPrice,
					Price:         articleUnitPrice,
					NetPrice:      article.NetPrice,
					VatPercentage: article.VatPercentage,
				})
			}
		}
	} else {
		getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
			OrderID: &item.Id,
		}
		relatedOrderProcurementArticle, err := r.GetOrderProcurementArticles(&getOrderProcurementArticleInput)
		if err != nil {
			return nil, err
		}

		for _, article := range relatedOrderProcurementArticle.Data {
			articles = append(articles, dto.DropdownProcurementAvailableArticle{
				Id:            article.Id,
				Title:         article.Title,
				Description:   article.Description,
				Amount:        article.Amount,
				NetPrice:      article.NetPrice,
				VatPercentage: strconv.Itoa(article.VatPercentage),
			})
		}
	}

	office := &dto.DropdownSimple{}
	if item.OfficeId != nil && *item.OfficeId > zero {
		officeItem, _ := r.GetDropdownSettingById(*item.OfficeId)
		office.Title = officeItem.Title
		office.Id = officeItem.Id
	}

	defaultFile := dto.FileDropdownSimple{
		Id:   0,
		Name: "",
		Type: "",
	}

	orderFile := defaultFile
	receiveFile := []dto.FileDropdownSimple{}
	movementFile := defaultFile

	if item.OrderFile != nil && *item.OrderFile != zero {
		file, err := r.GetFileByID(*item.OrderFile)
		if err != nil {
			return nil, err
		}
		orderFile.Id = *item.OrderFile
		orderFile.Name = file.Name
		orderFile.Type = *file.Type
	}

	for _, fileID := range item.ReceiveFile {
		file, err := r.GetFileByID(fileID)
		if err != nil {
			return nil, err
		}

		receiveFile = append(receiveFile, dto.FileDropdownSimple{
			Id:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		})
	}

	if item.MovementFile != nil && *item.MovementFile != zero {
		file, err := r.GetFileByID(*item.MovementFile)
		if err != nil {
			return nil, err
		}
		movementFile.Id = file.ID
		movementFile.Name = file.Name
		movementFile.Type = *file.Type
	}

	var groupOfArticles dto.DropdownSimple
	if item.GroupOfArticlesID != nil && *item.GroupOfArticlesID != zero {
		getGroupOfArticles, err := r.GetDropdownSettingById(*item.GroupOfArticlesID)

		if err != nil {
			return nil, err
		}
		groupOfArticles.Id = getGroupOfArticles.Id
		groupOfArticles.Title = getGroupOfArticles.Title
	}

	res = dto.OrderListOverviewResponse{
		Id:                  item.Id,
		DateOrder:           (string)(item.DateOrder),
		TotalNeto:           totalNeto,
		TotalBruto:          totalBruto,
		PublicProcurementID: procurementDropdown.Id,
		PublicProcurement:   &procurementDropdown,
		DateSystem:          item.DateSystem,
		InvoiceDate:         item.InvoiceDate,
		OrganizationUnitID:  item.OrganizationUnitId,
		OfficeID:            office.Id,
		Office:              office,
		Description:         item.Description,
		Status:              item.Status,
		Supplier:            &supplierDropdown,
		GroupOfArticles:     &groupOfArticles,
		Articles:            &articles,
		OrderFile:           orderFile,
		ReceiveFile:         receiveFile,
		MovementFile:        movementFile,
	}

	if item.RecipientUserId != nil && *item.RecipientUserId > 0 {
		userProfile, err := r.GetUserProfileById(*item.RecipientUserId)
		if err != nil {
			return nil, err
		}
		res.RecipientUser = &dto.DropdownSimple{
			Id:    userProfile.Id,
			Title: userProfile.GetFullName(),
		}
		res.RecipientUserID = item.RecipientUserId
	}

	if item.SupplierId != nil && *item.SupplierId != 0 {
		supplier, err := r.GetSupplier(*item.SupplierId)
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
