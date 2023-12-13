package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
)

var StockOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		input    dto.StockFilter
		articles []structs.StockArticle
		Total    int
	)

	size := params.Args["size"]
	page := params.Args["page"]
	search, searchOk := params.Args["title"].(string)
	date, dateOk := params.Args["date"].(string)
	sortByYear, sortByYearOk := params.Args["sort_by_year"].(string)
	sortByAmount, sortByAmountOK := params.Args["sort_by_amount"].(string)

	if shared.IsInteger(page) && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}
	if shared.IsInteger(size) && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}

	if searchOk && search != "" {
		input.Title = &search
	}

	if sortByAmountOK && sortByAmount != "" {
		input.SortByAmount = &sortByAmount
	}

	if sortByYearOk && sortByYear != "" {
		input.SortByYear = &sortByYear
	}

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return shared.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	if dateOk && date != "" {
		statusReceive := "Receive"
		orders, err := getOrderLists(&dto.GetOrderListInput{
			DateSystem:         &date,
			Status:             &statusReceive,
			OrganizationUnitId: organizationUnitID,
		})

		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, order := range orders.Data {
			orderArticles, err := getOrderProcurementArticles(&dto.GetOrderProcurementArticleInput{OrderID: &order.Id})
			if err != nil {
				return shared.HandleAPIError(err)
			}
			for _, article := range orderArticles.Data {
				flag := false

				if article.ArticleId != 0 {
					currentArticle, err := getProcurementArticle(article.ArticleId)

					if err != nil {
						return shared.HandleAPIError(err)
					}

					article.Title = currentArticle.Title
					article.Description = currentArticle.Description
					article.NetPrice = currentArticle.NetPrice
					vatPercentageInt, err := strconv.Atoi(currentArticle.VatPercentage)

					if err != nil {
						return nil, err
					}

					article.VatPercentage = vatPercentageInt
				}

				for i := 0; i < len(articles); i++ {
					if article.Title == articles[i].Title && article.Year == articles[i].Year {
						articles[i].Amount += article.Amount
						flag = true
						break
					}
				}

				if !flag {
					newArticle := structs.StockArticle{
						Title:         article.Title,
						Description:   article.Description,
						Year:          article.Year,
						Amount:        article.Amount,
						ID:            article.Id,
						NetPrice:      article.NetPrice,
						VatPercentage: article.VatPercentage,
					}
					articles = append(articles, newArticle)
				}
			}
		}

		movementArticles, err := getMovementArticleList(dto.OveralSpendingFilter{
			EndDate:            &date,
			OrganizationUnitID: organizationUnitID,
		})

		if err != nil {
			return shared.HandleAPIError(err)
		}

		for i := 0; i < len(movementArticles); i++ {
			for j := 0; j < len(articles); j++ {
				if movementArticles[i].Title == articles[j].Title && movementArticles[i].Year == articles[j].Year && officeInOrgUnit(movementArticles[i].OfficeID, *organizationUnitID) {
					articles[j].Amount -= movementArticles[i].Amount
				}
			}
		}

	} else {
		input.OrganizationUnitID = organizationUnitID

		articleList, total, err := getStock(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		Total = *total

		for _, article := range articleList {
			if article.Amount > 0 {
				articles = append(articles, article)
			}
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   Total,
		Items:   articles,
	}, nil
}

var MovementOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		input dto.MovementFilter
	)

	size := params.Args["size"]
	page := params.Args["page"]
	officeID, officeOk := params.Args["office_id"].(int)
	userID, userOk := params.Args["recipient_user_id"].(int)
	sortByDateOrder, sortByDateOrderOK := params.Args["sort_by_date_order"].(string)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return shared.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	input.OrganizationUnitID = organizationUnitID

	if shared.IsInteger(page) && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}
	if shared.IsInteger(size) && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}

	if officeOk && officeID != 0 {
		input.OfficeID = &officeID
	}

	if userOk && userID != 0 {
		input.RecipientUserID = &userID
	}

	if sortByDateOrderOK && sortByDateOrder != "" {
		input.SortByDateOrder = &sortByDateOrder
	}

	movementList, total, err := getMovements(&input)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	var response []dto.MovementResponse

	for _, movement := range movementList {
		var item dto.MovementResponse
		item.ID = movement.ID
		item.DateOrder = movement.DateOrder
		item.Description = movement.Description

		officeItem, _ := getDropdownSettingById(movement.OfficeID)

		if err != nil {
			return shared.HandleAPIError(err)
		}

		item.Office.Title = officeItem.Title
		item.Office.Id = officeItem.Id

		userItem, _ := getUserProfileById(movement.RecipientUserID)

		if err != nil {
			return shared.HandleAPIError(err)
		}

		item.RecipientUser.Title = userItem.FirstName + " " + userItem.LastName
		item.RecipientUser.Id = userItem.Id

		articles, err := getMovementArticles(item.ID)
		if err != nil {
			return nil, err
		}

		var movementArticles []dto.MovementArticle
		for _, article := range articles {
			stockArticle, err := getStockByID(article.StockID)

			if err != nil {
				return nil, err
			}

			movementArticle := dto.MovementArticle{
				Title:         stockArticle.Title,
				Description:   stockArticle.Description,
				Year:          stockArticle.Year,
				Amount:        article.Amount,
				ID:            stockArticle.ID,
				NetPrice:      stockArticle.NetPrice,
				VatPercentage: stockArticle.VatPercentage,
			}

			movementArticles = append(movementArticles, movementArticle)
		}

		item.Articles = movementArticles

		response = append(response, item)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   *total,
		Items:   response,
	}, nil
}

var MovementDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]

	response, err := buildMovementDetailsResponse(id.(int))

	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

var MovementArticlesResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var response []string
	title, titleOK := params.Args["title"].(string)
	var titleFilter *string

	if titleOK && title != "" {
		titleFilter = &title
	} else {
		titleFilter = nil
	}

	articles, err := getMovementArticleList(dto.OveralSpendingFilter{
		Title: titleFilter,
	})
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, article := range articles {
		str := article.Year + " " + article.Title
		found := false
		for _, element := range response {
			if strings.Contains(element, str) {
				found = true
				break
			}
		}
		if !found {
			response = append(response, str)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

var OrderListAssetMovementResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OrderAssetMovementItem

	dataBytes, _ := json.Marshal(params.Args["data"])
	_ = json.Unmarshal(dataBytes, &data)

	if data.ID == 0 {
		organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitID == nil {
			return shared.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
		}

		data.OrganizationUnitID = *organizationUnitID

		movement, err := createMovements(data)

		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, article := range data.Articles {
			item := dto.MovementArticle{
				StockID:    article.ID,
				MovementID: movement.ID,
				Amount:     article.Quantity,
			}

			if err != nil {
				return shared.HandleAPIError(err)
			}

			_, err = createMovementArticle(item)

			if err != nil {
				return shared.HandleAPIError(err)
			}

			stockArticle, err := getStockByID(article.ID)

			if err != nil {
				return shared.HandleAPIError(err)
			}

			stockArticle.Amount -= article.Quantity

			err = updateStock(*stockArticle)

			if err != nil {
				return shared.HandleAPIError(err)
			}
		}
	} else {
		_, err := updateMovements(data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You created movement!",
	}, nil
}

func buildMovementDetailsResponse(id int) (*dto.MovementDetailsResponse, error) {
	var item dto.MovementDetailsResponse

	movement, err := getMovementByID(id)
	if err != nil {
		return nil, err
	}

	item.ID = movement.ID
	item.DateOrder = movement.DateOrder
	item.Description = movement.Description

	officeItem, _ := getDropdownSettingById(movement.OfficeID)

	if err != nil {
		return nil, err
	}

	item.Office.Title = officeItem.Title
	item.Office.Id = officeItem.Id

	userItem, _ := getUserProfileById(movement.RecipientUserID)

	if err != nil {
		return nil, err
	}

	item.RecipientUser.Title = userItem.FirstName + " " + userItem.LastName
	item.RecipientUser.Id = userItem.Id

	articles, err := getMovementArticles(item.ID)
	if err != nil {
		return nil, err
	}

	if movement.FileID != 0 {
		file, err := getFileByID(movement.FileID)

		if err != nil {
			return nil, err
		}
		item.File.Id = file.ID
		item.File.Name = file.Name
		item.File.Type = *file.Type
	}

	var movementArticles []dto.MovementArticle
	for _, article := range articles {
		stockArticle, err := getStockByID(article.StockID)

		if err != nil {
			return nil, err
		}

		movementArticle := dto.MovementArticle{
			Title:         stockArticle.Title,
			Description:   stockArticle.Description,
			Year:          stockArticle.Year,
			Amount:        article.Amount,
			ID:            stockArticle.ID,
			NetPrice:      stockArticle.NetPrice,
			VatPercentage: stockArticle.VatPercentage,
		}

		movementArticles = append(movementArticles, movementArticle)
	}

	item.Articles = movementArticles
	return &item, nil

}

var MovementDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	orderList, err := getMovementByID(id)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	if orderList.FileID != 0 {
		err := deleteFile(orderList.FileID)

		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	articles, err := getMovementArticles(id)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return shared.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	for _, article := range articles {
		stock, err := getStockByID(article.StockID)

		if err != nil {
			return shared.HandleAPIError(err)
		}

		article.OrganizationUnitID = *organizationUnitID

		stock.Amount += article.Amount
		err = updateStock(*stock)
		if err != nil {
			return shared.HandleAPIError(err)
		}

	}

	err = deleteMovement(id)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted movement!",
	}, nil
}

func getStock(input *dto.StockFilter) ([]structs.StockArticle, *int, error) {
	res := &dto.GetStockResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.STOCK_ENDPOINT, input, res)
	if err != nil {
		return nil, nil, err
	}

	return res.Data, &res.Total, nil
}

func getStockByID(id int) (*structs.StockArticle, error) {
	res := &dto.GetSingleStockResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.STOCK_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getMovements(input *dto.MovementFilter) ([]structs.Movement, *int, error) {
	res := &dto.GetMovementResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.MOVEMENTS_ENDPOINT, input, res)
	if err != nil {
		return nil, nil, err
	}

	return res.Data, &res.Total, nil
}

func createMovements(input structs.OrderAssetMovementItem) (*structs.Movement, error) {
	res := &dto.GetSingleMovementResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.MOVEMENTS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createStock(input dto.MovementArticle) error {
	_, err := shared.MakeAPIRequest("POST", config.STOCK_ENDPOINT, input, nil)
	if err != nil {
		return err
	}

	return nil
}

func updateStock(input structs.StockArticle) error {
	_, err := shared.MakeAPIRequest("PUT", config.STOCK_ENDPOINT+"/"+strconv.Itoa(input.ID), input, nil)
	if err != nil {
		return err
	}

	return nil
}

func createMovementArticle(input dto.MovementArticle) (*dto.MovementArticle, error) {
	res := &dto.GetSingleMovementArticleResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.MOVEMENT_ARTICLES_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteMovement(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.MOVEMENTS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func updateMovements(input structs.OrderAssetMovementItem) (*structs.Movement, error) {
	res := &dto.GetSingleMovementResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.MOVEMENTS_ENDPOINT+"/"+strconv.Itoa(input.ID), input, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getMovementByID(id int) (*structs.Movement, error) {
	res := &dto.GetSingleMovementResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.MOVEMENTS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getMovementArticles(id int) ([]dto.MovementArticle, error) {
	input := dto.MovementArticlesFilter{
		MovementID: &id,
	}
	res := &dto.GetMovementArticleResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.MOVEMENT_ARTICLES_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func officeInOrgUnit(OfficeID int, OrganizationUnitID int) bool {
	orgUnitString := strconv.Itoa(OrganizationUnitID)

	res, err := getOfficeDropdownSettings(&dto.GetOfficesOfOrganizationInput{
		Value: &orgUnitString,
	})

	if err != nil {
		return false
	}

	for _, office := range res.Data {
		if office.Id == OfficeID {
			return true
		}
	}
	return false
}
