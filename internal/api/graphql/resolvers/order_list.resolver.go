package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

// processContractArticle refactored to only take context and contractArticle, and return an OrderArticleItem.
func processContractArticle(ctx context.Context, r *Resolver, contractArticle *structs.PublicProcurementContractArticle) (structs.OrderArticleItem, error) {
	organizationUnitID, _ := ctx.Value(config.OrganizationUnitIDKey).(*int)

	var resProcurementArticle *dto.ProcurementArticleResponseItem
	var err error

	// Get the related public procurement article details.
	relatedPublicProcurementArticle, _ := r.Repo.GetProcurementArticle(contractArticle.PublicProcurementArticleID)
	/*if err != nil {
		return structs.OrderArticleItem{}, errors.Wrap(err, "repo get procurement article")
	}*/

	if relatedPublicProcurementArticle != nil {

		// Build response item based on the related article and possibly organization unit.
		resProcurementArticle, err = buildProcurementArticleResponseItem(ctx, r, relatedPublicProcurementArticle, organizationUnitID)
		if err != nil {
			return structs.OrderArticleItem{}, errors.Wrap(err, "build procurement article response item")
		}
	}
	// Determine the amount based on the organization unit.
	amount := resProcurementArticle.Amount
	if organizationUnitID != nil && *organizationUnitID == 0 {
		amount = resProcurementArticle.TotalAmount
	}

	// Get overages for the contract article.
	overageList, _ := r.Repo.GetProcurementContractArticleOverageList(&dto.GetProcurementContractArticleOverageInput{
		ContractArticleID:  &contractArticle.ID,
		OrganizationUnitID: organizationUnitID,
	})
	/*if err != nil {
		return structs.OrderArticleItem{}, errors.Wrap(err, "repo get procurement contract article overage list")
	}*/

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
func GetProcurementArticles(ctx context.Context, r *Resolver, publicProcurementID int) ([]structs.OrderArticleItem, error) {
	var items []structs.OrderArticleItem
	itemsMap := make(map[int]structs.OrderArticleItem)

	// Get related contracts.
	relatedContractsResponse, _ := r.Repo.GetProcurementContractsList(&dto.GetProcurementContractsInput{
		ProcurementID: &publicProcurementID,
	})
	/*if err != nil {
		return nil, errors.Wrap(err, "repo get procurement contracts list")
	}*/

	// Process each contract.
	for _, contract := range relatedContractsResponse.Data {
		relatedContractArticlesResponse, err := r.Repo.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
			ContractID: &contract.ID,
		})
		if err != nil {
			return nil, errors.Wrap(err, "repo get procurement contract artcliles list")
		}

		// Process each contract article.
		for _, contractArticle := range relatedContractArticlesResponse.Data {
			newItem, err := processContractArticle(ctx, r, contractArticle)
			if err != nil {
				return nil, errors.Wrap(err, "process contact article")
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
		price float64
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		orderListItem, err := buildOrderListResponseItem(params.Context, r, orderList)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		items = []dto.OrderListOverviewResponse{*orderListItem}
		total = 1
	} else if activePlan {
		inputPlans := dto.GetProcurementPlansInput{}
		plans, _ := r.Repo.GetProcurementPlanList(&inputPlans)
		/*if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}*/
		currentYear := time.Now().Year()
		inputOrderList := dto.GetOrderListInput{}
		for _, plan := range plans {

			if plan.Year <= strconv.Itoa(currentYear) {
				item, _ := buildProcurementPlanResponseItem(params.Context, r, plan, nil, &dto.GetProcurementItemListInputMS{})

				if item.Status == dto.PlanStatusPostBudgetClosed {
					if len(item.Items) > 0 {
						for _, procurement := range item.Items {
							inputOrderList.PublicProcurementID = &procurement.ID
							orderLists, err := r.Repo.GetOrderLists(&inputOrderList)
							if err != nil {
								_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
								return errors.HandleAPPError(err)
							}
							for _, orderList := range orderLists.Data {
								if orderList.IsUsed {
									continue
								}
								orderListItem, err := buildOrderListResponseItem(params.Context, r, &orderList)
								if err != nil {
									_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
									return errors.HandleAPPError(err)
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		for _, orderList := range orderLists.Data {
			orderListItem, err := buildOrderListResponseItem(params.Context, r, &orderList)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
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
	var item *dto.OrderListOverviewResponse
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	itemID := data.ID

	listInsertItem, err := buildOrderListInsertItem(params.Context, r.Repo, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if itemID != 0 {
		res, err := r.Repo.UpdateOrderListItem(params.Context, itemID, listInsertItem)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if len(data.Articles) > 0 {
			err := deleteOrderArticles(r.Repo, itemID)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			err = r.Repo.CreateOrderListProcurementArticles(res.ID, data)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}

		item, err = buildOrderListResponseItem(params.Context, r, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		invoices, _, err := r.Repo.GetInvoiceList(&dto.GetInvoiceListInputMS{OrderID: &item.ID})
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if len(invoices) == 1 {
			if invoices[0].ProFormaInvoiceFileID != item.OrderFile.ID {
				invoices[0].ProFormaInvoiceFileID = item.OrderFile.ID
				_, err := r.Repo.UpdateInvoice(params.Context, &invoices[0])

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		listInsertItem.Status = "Created"
		listInsertItem.IsUsed = false

		for _, item := range data.Articles {

			articles, err := GetProcurementArticles(params.Context, r, *listInsertItem.PublicProcurementID)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			for _, article := range articles {
				if item.ID == article.ID {
					processedArticle, err := ProcessOrderArticleItem(r.Repo, article, listInsertItem.OrganizationUnitID)
					if err != nil {
						_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
						return errors.HandleAPPError(err)
					}

					if processedArticle.Available-item.Amount < 0 {
						err = errors.New("there are not available articles")
						_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
						return errors.HandleAPPError(err)
					}
				}
			}

		}

		res, err := r.Repo.CreateOrderListItem(params.Context, listInsertItem)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		err = r.Repo.CreateOrderListProcurementArticles(res.ID, data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		item, err = buildOrderListResponseItem(params.Context, r, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func (r *Resolver) PassOrderListToFinance(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	orderListBE, err := r.Repo.GetOrderListByID(id)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	orderList, err := buildOrderListResponseItem(params.Context, r, orderListBE)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	err = r.Repo.SendOrderListToFinance(params.Context, id)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var proFormaInvoiceDatePtr *time.Time
	if orderList.ProFormaInvoiceDate != nil {
		proFormaInvoiceDate, _ := parseDate(*orderList.ProFormaInvoiceDate)
		proFormaInvoiceDatePtr = &proFormaInvoiceDate
	}

	var receiptDatePtr *time.Time
	if orderList.DateSystem != nil {
		receiptDate, _ := parseDate(*orderList.DateSystem)
		receiptDatePtr = &receiptDate
	}

	var invoiceDatePtr *time.Time
	if orderList.InvoiceDate != nil {
		invoiceDate, _ := parseDate(*orderList.InvoiceDate)
		invoiceDatePtr = &invoiceDate
	}

	invoice := structs.Invoice{
		ProFormaInvoiceNumber: orderList.ProFormaInvoiceNumber,
		ProFormaInvoiceDate:   proFormaInvoiceDatePtr,
		Status:                "Nepotpun",
		Type:                  "invoices",
		SupplierID:            orderList.SupplierID,
		OrderID:               orderList.ID,
		OrganizationUnitID:    orderList.OrganizationUnitID,
		FileID:                orderList.OrderFile.ID,
		Registred:             false,
		InvoiceNumber:         orderList.InvoiceNumber,
		ReceiptDate:           receiptDatePtr,
		DateOfInvoice:         invoiceDatePtr,
	}

	if len(orderList.ReceiveFile) > 0 {
		invoice.ProFormaInvoiceFileID = orderList.ReceiveFile[0].ID
	}

	insertedItem, err := r.Repo.CreateInvoice(params.Context, &invoice)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, article := range *orderList.Articles {
		vatPercentage, _ := strconv.Atoi(article.VatPercentage)

		invoiceArticle := structs.InvoiceArticles{
			Title:         article.Title,
			NetPrice:      float64(article.NetPrice),
			VatPercentage: vatPercentage,
			Description:   article.Description,
			InvoiceID:     insertedItem.ID,
			Amount:        article.Amount,
		}

		_, err = r.Repo.CreateInvoiceArticle(&invoiceArticle)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}
	return dto.Response{
		Status: "success",
	}, nil
}

func (r *Resolver) OrderProcurementAvailableResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []structs.OrderArticleItem
		total int
	)
	publicProcurementID, ok := params.Args["public_procurement_id"].(int)
	if !ok || publicProcurementID <= 0 {
		return errors.HandleAPPError(errors.New("you must pass the item procurement id"))
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
			return errors.HandleAPPError(fmt.Errorf("user does not have organization unit assigned"))
		}

		organizationUnitID = *organizationUnitID*/
		organizationUnitID = 0
	}

	articles, _ := GetProcurementArticles(ctx, r, publicProcurementID)
	/*if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/

	for _, item := range articles {
		if visibilityType != nil && visibilityType.(int) > 0 && visibilityType.(int) != int(item.VisibilityType) {
			continue
		}
		processedArticle, err := ProcessOrderArticleItem(r.Repo, item, organizationUnitID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if processedArticle.Available > 0 {
			items = append(items, processedArticle)
		}
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
	var consumedAmount int
	currentArticle := article // work with a copy to avoid modifying the original

	articleVat, _ := strconv.ParseFloat(article.VatPercentage, 32)
	articleVat32 := float64(articleVat)
	currentArticle.Price = article.NetPrice + article.NetPrice*(articleVat32/100)

	getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
		ArticleID: &currentArticle.ID,
	}

	relatedOrderProcurementArticleResponse, _ := r.GetOrderProcurementArticles(&getOrderProcurementArticleInput)
	/*if err != nil {
		return currentArticle, errors.Wrap(err, "repo get order procurement articles")
	}*/

	if relatedOrderProcurementArticleResponse.Total > 0 {
		for _, orderArticle := range relatedOrderProcurementArticleResponse.Data {
			order, err := r.GetOrderListByID(orderArticle.OrderID)

			if err != nil {
				return currentArticle, errors.Wrap(err, "repo get order list by id")
			}

			if organizationUnitID > 0 && order.OrganizationUnitID == organizationUnitID {
				// if article is used in another order, deduct the amount to get Available articles
				currentArticle.TotalPrice *= float64(currentArticle.Available-orderArticle.Amount) / float64(currentArticle.Available)
				consumedAmount += orderArticle.Amount
			}
		}
	}

	if organizationUnitID == 0 {
		articles, err := r.GetOrganizationUnitArticlesList(dto.GetProcurementOrganizationUnitArticleListInputDTO{
			ArticleID: &currentArticle.ID,
		})
		if err != nil {
			return currentArticle, errors.Wrap(err, "repo get organization unit articles list")
		}
		amount := 0
		for _, article := range articles {
			amount += article.Amount
		}

		for _, article := range relatedOrderProcurementArticleResponse.Data {
			amount -= article.Amount
			consumedAmount += article.Amount
		}
		currentArticle.Available = amount
		currentArticle.ConsumedAmount = consumedAmount
	}

	input := dto.InventoryItemFilter{
		ArticleID: &article.ID,
	}

	if organizationUnitID != 0 {
		input.OrganizationUnitID = &organizationUnitID
	}

	articleInventory, _ := r.GetAllInventoryItem(input)

	/*if err != nil {
		return currentArticle, errors.Wrap(err, "repo get all inventory item")
	}*/

	numberOfArticles := len(articleInventory.Data)

	for _, article := range articleInventory.Data {
		if organizationUnitID != 0 && article.TargetOrganizationUnitID != 0 && article.TargetOrganizationUnitID != organizationUnitID {
			numberOfArticles--
		}
	}

	consumedAmount += numberOfArticles

	currentArticle.ConsumedAmount = consumedAmount
	currentArticle.Available -= consumedAmount

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
		"2006-01-02T15:04:05.000Z",           // ISO 8601 sa milisekundama
		"2006-01-02T15:04:05.000",            // ISO 8601 sa milisekundama bez Z
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

	employees, _ := GetEmployeesOfOrganizationUnit(r.Repo, *organizationUnitID)
	/*	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}*/
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
		return nil, errors.Wrap(err, "repo get systematization")
	}
	if len(systematizationsRes.Data) == 0 {
		return nil, errors.Wrap(errors.New("no systematization"), "repo get systematization")
	}
	systematization := systematizationsRes.Data[0]
	jobPositionsInOrganizationUnit, err := r.GetJobPositionsInOrganizationUnits(&dto.GetJobPositionInOrganizationUnitsInput{SystematizationID: &systematization.ID})
	if err != nil {
		return nil, errors.Wrap(err, "repo get job positions in organization units")
	}

	for _, jobPosition := range jobPositionsInOrganizationUnit.Data {
		employeesByJobPosition, err := r.GetEmployeesInOrganizationUnitList(&dto.GetEmployeesInOrganizationUnitInput{PositionInOrganizationUnit: &jobPosition.ID})
		if err != nil {
			return nil, errors.Wrap(err, "repo get employees in organization unit list")
		}

		for _, employee := range employeesByJobPosition {
			userProfile, err := r.GetUserProfileByID(employee.UserProfileID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get user profile by id")
			}
			userProfileList = append(userProfileList, userProfile)
		}
	}

	activeBool := true
	input := dto.GetJudgeResolutionListInputMS{
		Active: &activeBool,
	}

	resolution, err := r.GetJudgeResolutionList(&input)

	if err != nil {
		return nil, errors.Wrap(err, "repo get resolution list")
	}

	if len(resolution.Data) > 0 {

		filter := dto.JudgeResolutionsOrganizationUnitInput{
			ResolutionID:       &resolution.Data[0].ID,
			OrganizationUnitID: &id,
		}

		judges, _, err := r.GetJudgeResolutionOrganizationUnit(&filter)

		if err != nil {
			return nil, errors.Wrap(err, "repo get judge resolution organization unit")
		}

		if len(judges) > 0 {
			for _, item := range judges {
				userProfile, err := r.GetUserProfileByID(item.UserProfileID)

				if err != nil {
					return nil, errors.Wrap(err, "repo get user profile by id")
				}

				userProfileList = append(userProfileList, userProfile)
			}
		}

	}

	seen := make(map[int]bool)
	var uniqueProfiles []*structs.UserProfiles

	for _, profile := range userProfileList {
		if !seen[profile.ID] {
			seen[profile.ID] = true
			uniqueProfiles = append(uniqueProfiles, profile)
		}
	}

	return uniqueProfiles, nil
}

func deleteOrderArticles(r repository.MicroserviceRepositoryInterface, itemID int) error {
	orderProcurementArticlesResponse, err := r.GetOrderProcurementArticles(&dto.GetOrderProcurementArticleInput{
		OrderID: &itemID,
	})
	if err != nil {
		return errors.Wrap(err, "repo get order procurement articles")
	}

	for _, orderProcurementArticle := range orderProcurementArticlesResponse.Data {
		err = r.DeleteOrderProcurementArticle(orderProcurementArticle.ID)
		if err != nil {
			return errors.Wrap(err, "repo delete order procurement article")
		}
	}
	return nil
}

func (r *Resolver) OrderListDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	orderList, err := r.Repo.GetOrderListByID(itemID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if orderList.OrderFile != nil && *orderList.OrderFile != 0 {
		err := r.Repo.DeleteFile(*orderList.OrderFile)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	for _, fileID := range orderList.ReceiveFile {
		if fileID != 0 {
			err := r.Repo.DeleteFile(fileID)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}
	}

	if orderList.MovementFile != nil && *orderList.MovementFile != 0 {
		err := r.Repo.DeleteFile(*orderList.MovementFile)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	err = deleteOrderArticles(r.Repo, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	invoice, total, err := r.Repo.GetInvoiceList(&dto.GetInvoiceListInputMS{
		OrderID: &itemID,
	})

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if total > 0 {
		err = r.Repo.DeleteInvoice(params.Context, invoice[0].ID)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	err = r.Repo.DeleteOrderList(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	status := orderList.Status

	orderList.Status = "Receive"
	orderList.InvoiceNumber = data.InvoiceNumber
	orderList.DateSystem = &data.DateSystem
	orderList.InvoiceDate = data.InvoiceDate
	orderList.ReceiveFile = data.ReceiveFile
	orderList.DeliveryDate = data.DeliveryDate
	orderList.DeliveryNumber = data.DeliveryNumber
	orderList.DeliveryFileID = data.DeliveryFileID
	orderList.ReceiveFile = append(orderList.ReceiveFile, data.ReceiveFile...)

	if data.Description != nil {
		orderList.Description = data.Description
	}

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return errors.HandleAPPError(fmt.Errorf("user does not have organization unit assigned"))
	}
	var articles []structs.OrderArticleInsertItem
	if len(data.Articles) == 0 {
		articlesBE, err := r.Repo.GetOrderProcurementArticles(&dto.GetOrderProcurementArticleInput{OrderID: &orderList.ID})

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, article := range articlesBE.Data {
			articles = append(articles, structs.OrderArticleInsertItem{
				ID:            article.ID,
				Amount:        article.Amount,
				Title:         article.Title,
				Description:   article.Description,
				NetPrice:      article.NetPrice,
				VatPercentage: article.VatPercentage,
			})
		}
	} else {
		articles = data.Articles

		for _, article := range articles {

			oldArticle, err := r.Repo.GetOrderProcurementArticleByID(article.ID)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			_, err = r.Repo.UpdateOrderProcurementArticle(&structs.OrderProcurementArticleItem{
				ID:            article.ID,
				OrderID:       oldArticle.OrderID,
				ArticleID:     oldArticle.ArticleID,
				Year:          oldArticle.Year,
				Title:         oldArticle.Title,
				Description:   oldArticle.Description,
				Amount:        oldArticle.Amount,
				NetPrice:      article.NetPrice,
				VatPercentage: article.VatPercentage,
			})

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}
	}

	if status != "Receive" {
		if (orderList.GroupOfArticlesID != nil && *orderList.GroupOfArticlesID != 0) || orderList.IsProFormaInvoice {
			for _, article := range articles {
				orderArticle, err := r.Repo.GetOrderProcurementArticleByID(article.ID)

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}

				orderArticle.NetPrice = article.NetPrice
				orderArticle.VatPercentage = article.VatPercentage

				_, err = r.Repo.UpdateOrderProcurementArticle(orderArticle)

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}

				stock, _, _ := r.Repo.GetStock(&dto.StockFilter{
					Year:               &orderArticle.Year,
					Title:              &orderArticle.Title,
					Description:        &orderArticle.Description,
					NetPrice:           &orderArticle.NetPrice,
					VatPercentage:      &orderArticle.VatPercentage,
					OrganizationUnitID: organizationUnitID})

				err = r.Repo.AddOnStock(stock, *orderArticle, *organizationUnitID, true)

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		} else {
			orderArticles, err := r.Repo.GetOrderProcurementArticles(&dto.GetOrderProcurementArticleInput{
				OrderID: &orderList.ID,
			})

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			for _, orderArticle := range orderArticles.Data {
				currentArticle, err := r.Repo.GetProcurementArticle(orderArticle.ArticleID)

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}

				vatPercentageInt, err := strconv.Atoi(currentArticle.VatPercentage)

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}

				stockArticle, _, err := r.Repo.GetStock(&dto.StockFilter{
					Title:              &currentArticle.Title,
					Year:               &orderArticle.Year,
					OrganizationUnitID: organizationUnitID,
				})

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}

				orderArticle.Title = currentArticle.Title
				orderArticle.Description = currentArticle.Description
				orderArticle.NetPrice = currentArticle.NetPrice
				orderArticle.VatPercentage = vatPercentageInt

				err = r.Repo.AddOnStock(stockArticle, orderArticle, *organizationUnitID, false)

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}
	}

	_, err = r.Repo.UpdateOrderListItem(params.Context, data.OrderID, orderList)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	invoices, _, err := r.Repo.GetInvoiceList(&dto.GetInvoiceListInputMS{
		OrderID: &data.OrderID,
	})

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if len(invoices) > 0 {
		if data.InvoiceNumber != nil && data.InvoiceDate != nil {
			invoices[0].InvoiceNumber = *data.InvoiceNumber
			dateOfInvoice, err := parseDate(*data.InvoiceDate)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			invoices[0].DateOfInvoice = &dateOfInvoice

			if len(orderList.ReceiveFile) > 0 {
				invoices[0].FileID = orderList.ReceiveFile[0]
			}

			_, err = r.Repo.UpdateInvoice(params.Context, &invoices[0])
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if orderList.MovementFile != nil && *orderList.MovementFile != 0 {
		err := r.Repo.DeleteFile(*orderList.MovementFile)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	for _, fileID := range orderList.ReceiveFile {
		err := r.Repo.DeleteFile(fileID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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

	_, err = r.Repo.UpdateOrderListItem(params.Context, id, orderList)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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

	totalPrice := float64(0.0)
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
			return nil, errors.Wrap(err, "repo get procurement contracts list")
		}

		for _, contract := range relatedContractsResponse.Data {
			supplierID = &contract.SupplierID
			relatedContractArticlesResponse, err := r.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
				ContractID: &contract.ID,
			})
			if err != nil {
				return nil, errors.Wrap(err, "repo get contract articles list")
			}

			for _, contractArticle := range relatedContractArticlesResponse.Data {
				if article, exists := articleMap[contractArticle.PublicProcurementArticleID]; exists {
					totalPrice += (contractArticle.GrossValue) * float64(article.Amount)
				}
			}
		}
	}

	if item.PublicProcurementID == 0 {
		supplierID = &item.SupplierID
	} else {
		item, err := r.GetProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &item.PublicProcurementID})

		if err != nil {
			return nil, errors.Wrap(err, "repo get procurement contracts list")
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

func buildOrderListResponseItem(context context.Context, r *Resolver, item *structs.OrderListItem) (*dto.OrderListOverviewResponse, error) {

	var res dto.OrderListOverviewResponse
	var procurementDropdown dto.DropdownSimple
	var supplierDropdown dto.DropdownSimple
	articles := []dto.DropdownProcurementAvailableArticle{}
	totalBruto := float64(0.0)
	totalNeto := float64(0.0)
	zero := 0

	if item.PublicProcurementID != nil && *item.PublicProcurementID != zero {
		procurementItem, _ := r.Repo.GetProcurementItem(*item.PublicProcurementID)
		/*if err != nil {
			return nil, errors.Wrap(err, "repo get procurement item")
		}*/

		if procurementItem != nil {
			procurementDropdown.ID = procurementItem.ID
			procurementDropdown.Title = procurementItem.Title

			contract, _ := r.Repo.GetProcurementContractsList(&dto.GetProcurementContractsInput{ProcurementID: &procurementItem.ID})
			/*if err != nil {
				return nil, errors.Wrap(err, "repo get procurement contracts list")
			}*/

			if len(contract.Data) > 0 {
				supplier, _ := r.Repo.GetSupplier(contract.Data[0].SupplierID)

				/*if err != nil {
					return nil, errors.Wrap(err, "repo get supplier")
				}*/

				if supplier != nil {
					supplierDropdown = dto.DropdownSimple{
						ID:    supplier.ID,
						Title: supplier.Title,
					}
				}

				// getting articles and total price
				getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
					OrderID: &item.ID,
				}
				relatedOrderProcurementArticle, _ := r.Repo.GetOrderProcurementArticles(&getOrderProcurementArticleInput)
				/*if err != nil {
					return nil, errors.Wrap(err, "repo get order procurement articles")
				}*/

				publicProcurementArticles, _ := GetProcurementArticles(context, r, *item.PublicProcurementID)
				/*if err != nil {
					return nil, errors.Wrap(err, "repo get procurement articles")
				}*/

				publicProcurementArticlesMap := make(map[int]structs.OrderArticleItem)
				for _, article := range publicProcurementArticles {
					publicProcurementArticlesMap[article.ID] = article
				}

				for _, itemOrderArticle := range relatedOrderProcurementArticle.Data {
					if article, exists := publicProcurementArticlesMap[itemOrderArticle.ArticleID]; exists {
						articleVat, _ := strconv.ParseFloat(article.VatPercentage, 32)
						articleVat32 := float64(articleVat)
						articleUnitPrice := article.NetPrice + article.NetPrice*articleVat32/100
						articleTotalPrice := articleUnitPrice * float64(itemOrderArticle.Amount)
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
			}
		}
	} else {
		getOrderProcurementArticleInput := dto.GetOrderProcurementArticleInput{
			OrderID: &item.ID,
		}
		relatedOrderProcurementArticle, err := r.Repo.GetOrderProcurementArticles(&getOrderProcurementArticleInput)
		if err != nil {
			return nil, errors.Wrap(err, "repo get order procurement articles")
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
			totalNeto += article.NetPrice * float64(article.Amount)
			totalBruto += (article.NetPrice + article.NetPrice*float64(article.VatPercentage/100)) * float64(article.Amount)
		}
	}

	office := &dto.DropdownSimple{}
	if item.OfficeID != nil && *item.OfficeID > zero {
		officeItem, _ := r.Repo.GetDropdownSettingByID(*item.OfficeID)
		if officeItem != nil {
			office.Title = officeItem.Title
			office.ID = officeItem.ID
		}
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
		file, _ := r.Repo.GetFileByID(*item.OrderFile)
		/*if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}*/
		if file != nil {
			orderFile.ID = *item.OrderFile
			orderFile.Name = file.Name
			orderFile.Type = *file.Type
		}
	}

	for _, fileID := range item.ReceiveFile {

		if fileID == 0 {
			continue
		}

		file, _ := r.Repo.GetFileByID(fileID)
		/*if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}*/

		if file != nil {

			receiveFile = append(receiveFile, dto.FileDropdownSimple{
				ID:   file.ID,
				Name: file.Name,
				Type: *file.Type,
			})
		}
	}

	if item.MovementFile != nil && *item.MovementFile != zero {
		file, _ := r.Repo.GetFileByID(*item.MovementFile)
		/*if err != nil {
			return nil, errors.Wrap(err, "repo get movement file")
		}*/

		if file != nil {
			movementFile.ID = file.ID
			movementFile.Name = file.Name
			movementFile.Type = *file.Type
		}
	}

	var groupOfArticles dto.DropdownSimple
	if item.GroupOfArticlesID != nil && *item.GroupOfArticlesID != zero {
		getGroupOfArticles, _ := r.Repo.GetDropdownSettingByID(*item.GroupOfArticlesID)
		/*
			if err != nil {
				return nil, errors.Wrap(err, "repo get dropdown setting by id")
			}*/

		if getGroupOfArticles != nil {
			groupOfArticles.ID = getGroupOfArticles.ID
			groupOfArticles.Title = getGroupOfArticles.Title
		}
	}
	var account dto.DropdownSimple
	if item.AccountID != nil && *item.AccountID != zero {
		getAccount, _ := r.Repo.GetAccountItemByID(*item.AccountID)

		/*		if err != nil {
				return nil, errors.Wrap(err, "repo get account item by id")
			}*/
		if getAccount != nil {
			account.ID = getAccount.ID
			account.Title = getAccount.Title
		}
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
		DeliveryDate:          item.DeliveryDate,
		DeliveryNumber:        item.DeliveryNumber,
	}

	if res.ProFormaInvoiceDate != nil && *res.ProFormaInvoiceDate == "0001-01-01T00:00:00Z" {
		res.ProFormaInvoiceDate = nil
	}

	if item.RecipientUserID != nil && *item.RecipientUserID > 0 {
		userProfile, _ := r.Repo.GetUserProfileByID(*item.RecipientUserID)
		/*if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}*/

		if userProfile != nil {
			res.RecipientUser = &dto.DropdownSimple{
				ID:    userProfile.ID,
				Title: userProfile.GetFullName(),
			}
			res.RecipientUserID = item.RecipientUserID
		}
	}

	if item.SupplierID != nil && *item.SupplierID != 0 {
		supplier, _ := r.Repo.GetSupplier(*item.SupplierID)
		/*if err != nil {
			return nil, errors.Wrap(err, "repo get supplier")
		}*/

		if supplier != nil {
			res.SupplierID = supplier.ID
			res.Supplier = &dto.DropdownSimple{
				ID:    supplier.ID,
				Title: supplier.Title,
			}
		}
	}

	if item.InvoiceNumber != nil {
		res.InvoiceNumber = *item.InvoiceNumber
	}

	if item.DeliveryFileID != nil && *item.DeliveryFileID != 0 {
		file, _ := r.Repo.GetFileByID(*item.DeliveryFileID)
		/*if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}*/

		if file != nil {

			deliveryFile := dto.FileDropdownSimple{
				ID:   file.ID,
				Name: file.Name,
				Type: *file.Type,
			}

			res.DeliveryFile = &deliveryFile
		}
	}

	return &res, nil
}
