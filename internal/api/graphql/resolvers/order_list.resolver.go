package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/internal/api/repository"
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
func processContractArticle(ctx context.Context, r repository.MicroserviceRepositoryInterface, contractArticle *structs.PublicProcurementContractArticle) (structs.OrderArticleItem, error) {
	organizationUnitID, _ := ctx.Value(config.OrganizationUnitIDKey).(*int)

	// Get the related public procurement article details.
	relatedPublicProcurementArticle, err := r.GetProcurementArticle(contractArticle.PublicProcurementArticleID)
	if err != nil {
		return structs.OrderArticleItem{}, err
	}

	// Build response item based on the related article and possibly organization unit.
	resProcurementArticle, err := buildProcurementArticleResponseItem(ctx, r, relatedPublicProcurementArticle, organizationUnitID)
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
		ContractArticleID:  &contractArticle.ID,
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
		ID:             relatedPublicProcurementArticle.ID,
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
func GetProcurementArticles(ctx context.Context, r repository.MicroserviceRepositoryInterface, publicProcurementID int) ([]structs.OrderArticleItem, error) {
	var items []structs.OrderArticleItem
	itemsMap := make(map[int]structs.OrderArticleItem)

	// Get related contracts.
	relatedContractsResponse, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{
		ProcurementID: &publicProcurementID,
	})
	if err != nil {
		return nil, err
	}

	// Process each contract.
	for _, contract := range relatedContractsResponse.Data {
		relatedContractArticlesResponse, err := r.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
			ContractID: &contract.ID,
		})
		if err != nil {
			return nil, err
		}

		// Process each contract article.
		for _, contractArticle := range relatedContractArticlesResponse.Data {
			newItem, err := processContractArticle(ctx, r, contractArticle)
			if err != nil {
				return nil, err
			}

			if existingItem, exists := itemsMap[newItem.ID]; exists {
				// Update the existing item in the map if it exists.
				existingItem.Amount += newItem.Amount
				existingItem.Available += newItem.Available
				existingItem.TotalPrice += newItem.TotalPrice
				itemsMap[newItem.ID] = existingItem
			} else {
				// Add new item to the map and slice.
				itemsMap[newItem.ID] = newItem
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
	supplierID := params.Args["supplier_id"]
	publicProcurementID := params.Args["public_procurement_id"]
	status, statusOK := params.Args["status"].(string)
	search, searchOk := params.Args["search"].(string)
	activePlan, _ := params.Args["active_plan"].(bool)
	financeOverview, _ := params.Args["finance_overview"].(bool)
	year, yearOK := params.Args["year"].(string)
	sortByTotalPrice, sortByTotalPriceOK := params.Args["sort_by_total_price"].(string)
	sortByDateOrder, sortByDateOrderOK := params.Args["sort_by_date_order"].(string)

	if id != nil && id != 0 {
		orderList, err := r.Repo.GetOrderListByID(id.(int))
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		orderListItem, err := buildOrderListResponseItem(params.Context, r.Repo, orderList)
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
				item, _ := buildProcurementPlanResponseItem(params.Context, r.Repo, plan, nil, &dto.GetProcurementItemListInputMS{})

				if item.Status == dto.PlanStatusPostBudgetClosed {
					if len(item.Items) > 0 {
						for _, procurement := range item.Items {
							inputOrderList.PublicProcurementID = &procurement.ID
							orderLists, err := r.Repo.GetOrderLists(&inputOrderList)
							if err != nil {
								return apierrors.HandleAPIError(err)
							}
							for _, orderList := range orderLists.Data {
								if orderList.IsUsed {
									continue
								}
								orderListItem, err := buildOrderListResponseItem(params.Context, r.Repo, &orderList)
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
			input.OrganizationUnitID = organizationUnitID
		}
		if page != nil && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if size != nil && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if supplierID != nil && supplierID.(int) > 0 {
			supplierID := supplierID.(int)
			input.SupplierID = &supplierID
		}
		if publicProcurementID != nil && publicProcurementID.(int) > 0 {
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

		if financeOverview {
			input.FinanceOverview = &financeOverview
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
			orderListItem, err := buildOrderListResponseItem(params.Context, r.Repo, &orderList)
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

	itemID := data.ID

	if data.PassedToFinance && data.DateOrder == "" {
		err := r.Repo.SendOrderListToFinance(data.ID)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		response.Status = "success"
		response.Message = "You passed to finance this item!"
		return response, nil
	}

	listInsertItem, err := buildOrderListInsertItem(params.Context, r.Repo, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if itemID != 0 {
		listInsertItem.Status = "Updated"
		res, err := r.Repo.UpdateOrderListItem(itemID, listInsertItem)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		if len(data.Articles) > 0 {
			err := deleteOrderArticles(r.Repo, itemID)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
			err = r.Repo.CreateOrderListProcurementArticles(res.ID, data)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
		}

		item, err := buildOrderListResponseItem(params.Context, r.Repo, res)
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

		err = r.Repo.CreateOrderListProcurementArticles(res.ID, data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		item, err := buildOrderListResponseItem(params.Context, r.Repo, res)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		if item.IsProFormaInvoice && item.ProFormaInvoiceDate != nil {
			proFormaInvoiceDate, _ := parseDate(*item.ProFormaInvoiceDate)

			invoice := structs.Invoice{
				ProFormaInvoiceNumber: item.ProFormaInvoiceNumber,
				ProFormaInvoiceDate:   proFormaInvoiceDate,
				Status:                "waiting",
				Type:                  "invoice",
				SupplierID:            item.SupplierID,
				OrderID:               item.ID,
				OrganizationUnitID:    item.OrganizationUnitID,
				FileID:                item.OrderFile.ID,
			}

			insertedItem, err := r.Repo.CreateInvoice(&invoice)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			for _, article := range *item.Articles {
				vatPercentage, _ := strconv.Atoi(article.VatPercentage)

				invoiceArticle := structs.InvoiceArticles{
					Title:         article.Title,
					NetPrice:      float64(article.NetPrice),
					VatPercentage: vatPercentage,
					Description:   article.Description,
					InvoiceID:     insertedItem.ID,
					AccountID:     item.Account.ID,
				}

				_, err = r.Repo.CreateInvoiceArticle(&invoiceArticle)

				if err != nil {
					return apierrors.HandleAPIError(err)
				}
			}
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
		return apierrors.HandleAPIError(errors.New("you must pass the item procurement id"))
	}

	visibilityType := params.Args["visibility_type"]

	ctx := params.Context
	var organizationUnitID int

	if params.Args["organization_unit_id"] != nil && params.Args["organization_unit_id"].(int) != 0 {
		organizationUnitID = params.Args["organization_unit_id"].(int)
		ctx = context.WithValue(ctx, config.OrganizationUnitIDKey, &organizationUnitID)
	} else {
		/*organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitID == nil {
			return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
		}

		organizationUnitID = *organizationUnitID*/
		organizationUnitID = 0
	}

	articles, err := GetProcurementArticles(ctx, r.Repo, publicProcurementID)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	for _, item := range articles {
		if visibilityType != nil && visibilityType.(int) > 0 && visibilityType.(int) != int(item.VisibilityType) {
			continue
		}
		processedArticle, err := ProcessOrderArticleItem(r.Repo, item, organizationUnitID)
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
func ProcessOrderArticleItem(r repository.MicroserviceRepositoryInterface, article structs.OrderArticleItem, organizationUnitID int) (structs.OrderArticleItem, error) {
	var err error
	currentArticle := article // work with a copy to avoid modifying the original

	articleVat, _ := strconv.ParseFloat(article.VatPercentage, 32)
	articleVat32 := float32(articleVat)
	currentArticle.Price = article.NetPrice + article.NetPrice*(articleVat32/100)

	getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
		ArticleID: &currentArticle.ID,
	}

	relatedOrderProcurementArticleResponse, err := r.GetOrderProcurementArticles(&getOrderProcurementArticleInput)
	if err != nil {
		return currentArticle, err
	}

	if relatedOrderProcurementArticleResponse.Total > 0 {
		for _, orderArticle := range relatedOrderProcurementArticleResponse.Data {
			order, err := r.GetOrderListByID(orderArticle.OrderID)

			if err != nil {
				return currentArticle, err
			}

			if organizationUnitID > 0 && order.OrganizationUnitID == organizationUnitID {
				// if article is used in another order, deduct the amount to get Available articles
				currentArticle.TotalPrice *= float32(currentArticle.Available-orderArticle.Amount) / float32(currentArticle.Available)
				currentArticle.Available -= orderArticle.Amount
			}
		}
	}

	if organizationUnitID == 0 {
		articles, err := r.GetOrganizationUnitArticlesList(dto.GetProcurementOrganizationUnitArticleListInputDTO{
			ArticleID: &currentArticle.ID,
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
		ArticleID:          &article.ID,
		OrganizationUnitID: &organizationUnitID,
	})

	if err != nil {
		return currentArticle, err
	}

	currentArticle.Available -= len(articleInventory.Data)

	overages, err := r.GetProcurementContractArticleOverageList(&dto.GetProcurementContractArticleOverageInput{
		ContractArticleID:  &article.ID,
		OrganizationUnitID: &organizationUnitID})

	if err != nil {
		return currentArticle, err
	}

	for _, overage := range overages {
		currentArticle.Available += overage.Amount
	}

	return currentArticle, nil
}

func parseDate(dateString string) (time.Time, error) {
	var layouts = []string{
		"02.01.2006",                         // dd.mm.yyyy
		"02.01.2006.",                        // dd.mm.yyyy.
		"02/01/2006",                         // dd/mm/yyyy
		"02-01-2006",                         // dd-mm-yyyy
		"2006-01-02",                         // yyyy-mm-dd
		"01/02/2006",                         // mm/dd/yyyy
		"02.01.06",                           // dd.mm.yy
		"02.01.06.",                          // dd.mm.yy.
		"02/01/06",                           // dd/mm/yy
		"01/02/06",                           // mm/dd/yy
		"01-02-06",                           // mm-dd-yy
		"02-01-06",                           // dd-mm-yy
		"06-01-02",                           // yy-mm-dd
		"01/02/06",                           // mm/dd/yy
		"Jan 02, 2006",                       // Mon dd, yyyy
		"January 2, 2006",                    // Month dd, yyyy
		"02 Jan 2006",                        // dd Mon yyyy
		"2006-01-02T15:04:05Z07:00",          // ISO 8601
		"2006-01-02T15:04:05Z",               // ISO
		"Mon Jan 02 15:04:05 -0700 MST 2006", // Go standard format
		"Jan _2 15:04:05",                    // Without leading zeros in day
		"2006-01-02 15:04:05",                // YYYY-MM-DD HH:MM:SS
		"02 Jan 06 15:04 MST",                // YY instead of YYYY
		"02 Jan 2006 15:04",                  // dd Mon yyyy HH:MM
		"Mon, 02 Jan 2006 15:04:05 MST",      // Day, dd Mon yyyy HH:MM:SS
		"Mon, 02 Jan 2006 15:04:05 -0700",    // Day, dd Mon yyyy HH:MM:SS Timezone Offset
		"Monday, 02-Jan-06 15:04:05 MST",     // Full Day, dd-Mon-yy HH:MM:SS
		"Monday, 02-Jan-06 15:04:05 -0700",   // Full Day, dd-Mon-yy HH:MM:SS Timezone Offset
		"3:04 PM",                            // Short time
		"Jan 02, 2006 at 3:04pm",             // Date with time
	}

	var date time.Time
	var err error

	for _, layout := range layouts {
		date, err = time.Parse(layout, dateString)
		if err == nil {
			return date, nil
		}
	}

	numberOfDays, err := strconv.Atoi(dateString)

	if err == nil {
		startDate := time.Date(1899, time.December, 31, 0, 0, 0, 0, time.UTC)
		daysDuration := time.Duration(numberOfDays) * 24 * time.Hour
		return startDate.Add(daysDuration), nil
	}

	return date, fmt.Errorf("date format is not valid: %s", dateString)
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
			ID:    employee.ID,
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
	jobPositionsInOrganizationUnit, err := r.GetJobPositionsInOrganizationUnits(&dto.GetJobPositionInOrganizationUnitsInput{SystematizationID: &systematization.ID})
	if err != nil {
		return nil, err
	}

	for _, jobPosition := range jobPositionsInOrganizationUnit.Data {
		employeesByJobPosition, err := r.GetEmployeesInOrganizationUnitList(&dto.GetEmployeesInOrganizationUnitInput{PositionInOrganizationUnit: &jobPosition.ID})
		if err != nil {
			return nil, err
		}

		for _, employee := range employeesByJobPosition {
			userProfile, err := r.GetUserProfileByID(employee.UserProfileID)
			if err != nil {
				return nil, err
			}
			userProfileList = append(userProfileList, userProfile)
		}
	}

	return userProfileList, nil
}

func deleteOrderArticles(r repository.MicroserviceRepositoryInterface, itemID int) error {
	orderProcurementArticlesResponse, err := r.GetOrderProcurementArticles(&dto.GetOrderProcurementArticleInput{
		OrderID: &itemID,
	})
	if err != nil {
		return err
	}

	for _, orderProcurementArticle := range orderProcurementArticlesResponse.Data {
		err = r.DeleteOrderProcurementArticle(orderProcurementArticle.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Resolver) OrderListDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	orderList, err := r.Repo.GetOrderListByID(itemID)

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

	err = deleteOrderArticles(r.Repo, itemID)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var defaultString string
	if *orderList.InvoiceNumber == defaultString {
		invoice, total, err := r.Repo.GetInvoiceList(&dto.GetInvoiceListInputMS{
			OrderID: &itemID,
		})

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		if total > 0 {
			err = r.Repo.DeleteInvoice(invoice[0].ID)

			if err != nil {
				return apierrors.HandleAPIError(err)
			}
		}
	}

	err = r.Repo.DeleteOrderList(itemID)
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

	orderList, err := r.Repo.GetOrderListByID(data.OrderID)
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
		if (orderList.GroupOfArticlesID != nil && *orderList.GroupOfArticlesID != 0) || orderList.IsProFormaInvoice {
			for _, article := range data.Articles {
				orderArticle, err := r.Repo.GetOrderProcurementArticleByID(article.ID)

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
				OrderID: &orderList.ID,
			})

			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			for _, orderArticle := range orderArticles.Data {
				currentArticle, err := r.Repo.GetProcurementArticle(orderArticle.ArticleID)

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

	_, err = r.Repo.UpdateOrderListItem(data.OrderID, orderList)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if orderList.IsProFormaInvoice && orderList.ProFormaInvoiceNumber != "" {
		invoice, total, err := r.Repo.GetInvoiceList(&dto.GetInvoiceListInputMS{
			OrderID: &orderList.ID,
		})

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		if total > 0 {
			invoiceDate, _ := parseDate(*orderList.InvoiceDate)

			newInvoice := structs.Invoice{
				ID:                     invoice[0].ID,
				InvoiceNumber:          *orderList.InvoiceNumber,
				Status:                 invoice[0].Status,
				Type:                   invoice[0].Type,
				TypeOfSubject:          invoice[0].TypeOfSubject,
				TypeOfContract:         invoice[0].TypeOfContract,
				SourceOfFunding:        invoice[0].SourceOfFunding,
				Supplier:               invoice[0].Supplier,
				GrossPrice:             invoice[0].GrossPrice,
				VATPrice:               invoice[0].VATPrice,
				SupplierID:             invoice[0].SupplierID,
				OrderID:                invoice[0].OrderID,
				OrganizationUnitID:     invoice[0].OrganizationUnitID,
				ActivityID:             invoice[0].ActivityID,
				TaxAuthorityCodebookID: invoice[0].TaxAuthorityCodebookID,
				DateOfInvoice:          invoiceDate,
				ReceiptDate:            invoice[0].ReceiptDate,
				DateOfPayment:          invoice[0].DateOfPayment,
				DateOfStart:            invoice[0].DateOfStart,
				SSSInvoiceReceiptDate:  invoice[0].SSSInvoiceReceiptDate,
				FileID:                 invoice[0].FileID,
				BankAccount:            invoice[0].BankAccount,
				Description:            invoice[0].Description,
				ProFormaInvoiceDate:    invoice[0].ProFormaInvoiceDate,
				ProFormaInvoiceNumber:  invoice[0].ProFormaInvoiceNumber,
				Articles:               invoice[0].Articles,
				AdditionalExpenses:     invoice[0].AdditionalExpenses,
				CreatedAt:              invoice[0].CreatedAt,
				UpdatedAt:              invoice[0].UpdatedAt,
			}

			_, err = r.Repo.UpdateInvoice(&newInvoice)

			if err != nil {
				return apierrors.HandleAPIError(err)
			}

		}

	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You received this order!",
	}, nil
}

func (r *Resolver) OrderListReceiveDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	orderList, err := r.Repo.GetOrderListByID(id)
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
	orderList.OfficeID = nil
	orderList.RecipientUserID = nil
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

func buildOrderListInsertItem(context context.Context, r repository.MicroserviceRepositoryInterface, item *structs.OrderListInsertItem) (*structs.OrderListItem, error) {
	currentTime := time.Now().UTC()
	timeString := currentTime.Format("2006-01-02T15:04:05Z07:00")

	var newItem *structs.OrderListItem

	totalPrice := float32(0.0)
	var supplierID *int

	if len(item.Articles) > 0 {
		articleMap := make(map[int]structs.OrderArticleInsertItem)
		for _, article := range item.Articles {
			articleMap[article.ID] = article
		}

		relatedContractsResponse, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{
			ProcurementID: &item.PublicProcurementID,
		})
		if err != nil {
			return nil, err
		}

		for _, contract := range relatedContractsResponse.Data {
			supplierID = &contract.SupplierID
			relatedContractArticlesResponse, err := r.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
				ContractID: &contract.ID,
			})
			if err != nil {
				return nil, err
			}

			for _, contractArticle := range relatedContractArticlesResponse.Data {
				if article, exists := articleMap[contractArticle.PublicProcurementArticleID]; exists {
					totalPrice += (contractArticle.GrossValue) * float32(article.Amount)
				}
			}
		}
	}

	if item.PublicProcurementID == 0 {
		supplierID = &item.SupplierID
	} else {
		item, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &item.PublicProcurementID})

		if err != nil {
			return nil, err
		}
		if len(item.Data) > 0 {
			supplierID = &item.Data[0].SupplierID
		}

	}

	newItem = &structs.OrderListItem{
		ID:                    item.ID,
		DateOrder:             timeString,
		TotalPrice:            totalPrice,
		PublicProcurementID:   &item.PublicProcurementID,
		GroupOfArticlesID:     &item.GroupOfArticlesID,
		SupplierID:            supplierID,
		OrderFile:             &item.OrderFile,
		PassedToFinance:       item.PassedToFinance,
		UsedInFinance:         item.UsedInFinance,
		IsProFormaInvoice:     item.IsProFormaInvoice,
		ProFormaInvoiceNumber: item.ProFormaInvoiceNumber,
		AccountID:             &item.AccountID,
	}

	if item.ProFormaInvoiceDate != nil {
		newItem.ProFormaInvoiceDate = item.ProFormaInvoiceDate
	}

	// Getting organizationUnitID from job position
	loggedInProfile, _ := context.Value(config.LoggedInProfileKey).(*structs.UserProfiles)
	organizationUnitID, unitOK := context.Value(config.OrganizationUnitIDKey).(*int)

	newItem.RecipientUserID = &loggedInProfile.ID

	if unitOK && organizationUnitID != nil {
		newItem.OrganizationUnitID = *organizationUnitID
	}

	return newItem, nil
}

func buildOrderListResponseItem(context context.Context, r repository.MicroserviceRepositoryInterface, item *structs.OrderListItem) (*dto.OrderListOverviewResponse, error) {

	var res dto.OrderListOverviewResponse
	var procurementDropdown dto.DropdownSimple
	var supplierDropdown dto.DropdownSimple
	articles := []dto.DropdownProcurementAvailableArticle{}
	totalBruto := float32(0.0)
	totalNeto := float32(0.0)
	zero := 0

	if item.PublicProcurementID != nil && *item.PublicProcurementID != zero {
		procurementItem, err := r.GetProcurementItem(*item.PublicProcurementID)
		if err != nil {
			return nil, err
		}

		procurementDropdown.ID = procurementItem.ID
		procurementDropdown.Title = procurementItem.Title

		contract, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &procurementItem.ID})
		if err != nil {
			return nil, err
		}

		supplier, err := r.GetSupplier(contract.Data[0].SupplierID)

		if err != nil {
			return nil, err
		}

		supplierDropdown = dto.DropdownSimple{
			ID:    supplier.ID,
			Title: supplier.Title,
		}

		// getting articles and total price
		getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
			OrderID: &item.ID,
		}
		relatedOrderProcurementArticle, err := r.GetOrderProcurementArticles(&getOrderProcurementArticleInput)
		if err != nil {
			return nil, err
		}

		publicProcurementArticles, err := GetProcurementArticles(context, r, *item.PublicProcurementID)
		if err != nil {
			return nil, err
		}

		publicProcurementArticlesMap := make(map[int]structs.OrderArticleItem)
		for _, article := range publicProcurementArticles {
			publicProcurementArticlesMap[article.ID] = article
		}

		for _, itemOrderArticle := range relatedOrderProcurementArticle.Data {
			if article, exists := publicProcurementArticlesMap[itemOrderArticle.ArticleID]; exists {
				articleVat, _ := strconv.ParseFloat(article.VatPercentage, 32)
				articleVat32 := float32(articleVat)
				articleUnitPrice := article.NetPrice + article.NetPrice*articleVat32/100
				articleTotalPrice := articleUnitPrice * float32(itemOrderArticle.Amount)
				totalBruto += articleTotalPrice
				vat := articleTotalPrice * (100 - articleVat32) / 100
				totalNeto += vat

				articles = append(articles, dto.DropdownProcurementAvailableArticle{
					ID:            itemOrderArticle.ID,
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
			OrderID: &item.ID,
		}
		relatedOrderProcurementArticle, err := r.GetOrderProcurementArticles(&getOrderProcurementArticleInput)
		if err != nil {
			return nil, err
		}

		for _, article := range relatedOrderProcurementArticle.Data {
			articles = append(articles, dto.DropdownProcurementAvailableArticle{
				ID:            article.ID,
				Title:         article.Title,
				Description:   article.Description,
				Amount:        article.Amount,
				NetPrice:      article.NetPrice,
				VatPercentage: strconv.Itoa(article.VatPercentage),
			})
		}
	}

	office := &dto.DropdownSimple{}
	if item.OfficeID != nil && *item.OfficeID > zero {
		officeItem, _ := r.GetDropdownSettingByID(*item.OfficeID)
		office.Title = officeItem.Title
		office.ID = officeItem.ID
	}

	defaultFile := dto.FileDropdownSimple{
		ID:   0,
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
		orderFile.ID = *item.OrderFile
		orderFile.Name = file.Name
		orderFile.Type = *file.Type
	}

	for _, fileID := range item.ReceiveFile {
		file, err := r.GetFileByID(fileID)
		if err != nil {
			return nil, err
		}

		receiveFile = append(receiveFile, dto.FileDropdownSimple{
			ID:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		})
	}

	if item.MovementFile != nil && *item.MovementFile != zero {
		file, err := r.GetFileByID(*item.MovementFile)
		if err != nil {
			return nil, err
		}
		movementFile.ID = file.ID
		movementFile.Name = file.Name
		movementFile.Type = *file.Type
	}

	var groupOfArticles dto.DropdownSimple
	if item.GroupOfArticlesID != nil && *item.GroupOfArticlesID != zero {
		getGroupOfArticles, err := r.GetDropdownSettingByID(*item.GroupOfArticlesID)

		if err != nil {
			return nil, err
		}
		groupOfArticles.ID = getGroupOfArticles.ID
		groupOfArticles.Title = getGroupOfArticles.Title
	}

	var account dto.DropdownSimple
	if item.AccountID != nil && *item.AccountID != zero {
		getAccount, err := r.GetAccountItemByID(*item.AccountID)

		if err != nil {
			return nil, err
		}
		account.ID = getAccount.ID
		account.Title = getAccount.Title
	}

	res = dto.OrderListOverviewResponse{
		ID:                    item.ID,
		DateOrder:             (string)(item.DateOrder),
		TotalNeto:             totalNeto,
		TotalBruto:            totalBruto,
		PublicProcurementID:   procurementDropdown.ID,
		PublicProcurement:     &procurementDropdown,
		DateSystem:            item.DateSystem,
		InvoiceDate:           item.InvoiceDate,
		OrganizationUnitID:    item.OrganizationUnitID,
		OfficeID:              office.ID,
		Office:                office,
		Description:           item.Description,
		Status:                item.Status,
		Supplier:              &supplierDropdown,
		GroupOfArticles:       &groupOfArticles,
		Articles:              &articles,
		PassedToFinance:       item.PassedToFinance,
		UsedInFinance:         item.UsedInFinance,
		OrderFile:             orderFile,
		ReceiveFile:           receiveFile,
		MovementFile:          movementFile,
		IsProFormaInvoice:     item.IsProFormaInvoice,
		ProFormaInvoiceDate:   item.ProFormaInvoiceDate,
		ProFormaInvoiceNumber: item.ProFormaInvoiceNumber,
		Account:               &account,
	}

	if item.RecipientUserID != nil && *item.RecipientUserID > 0 {
		userProfile, err := r.GetUserProfileByID(*item.RecipientUserID)
		if err != nil {
			return nil, err
		}
		res.RecipientUser = &dto.DropdownSimple{
			ID:    userProfile.ID,
			Title: userProfile.GetFullName(),
		}
		res.RecipientUserID = item.RecipientUserID
	}

	if item.SupplierID != nil && *item.SupplierID != 0 {
		supplier, err := r.GetSupplier(*item.SupplierID)
		if err != nil {
			return nil, err
		}
		res.SupplierID = supplier.ID
		res.Supplier = &dto.DropdownSimple{
			ID:    supplier.ID,
			Title: supplier.Title,
		}
	}

	if item.InvoiceNumber != nil {
		res.InvoiceNumber = *item.InvoiceNumber
	}

	return &res, nil
}
