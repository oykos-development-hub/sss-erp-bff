package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"
)

var StockOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		input dto.StockFilter
	)

	size := params.Args["size"]
	page := params.Args["page"]
	search, searchOk := params.Args["title"].(string)

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

	articleList, total, err := getStock(&input)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   *total,
		Items:   articleList,
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

		/*if err != nil {
			return shared.HandleAPIError(err)
		}*/

		item.Office.Title = officeItem.Title
		item.Office.Id = officeItem.Id

		userItem, _ := getUserProfileById(movement.RecipientUserID)

		/*if err != nil {
			return shared.HandleAPIError(err)
		}*/

		item.RecipientUser.Title = userItem.FirstName + " " + userItem.LastName
		item.RecipientUser.Id = userItem.Id
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

var OrderListAssetMovementResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OrderAssetMovementItem

	dataBytes, _ := json.Marshal(params.Args["data"])
	_ = json.Unmarshal(dataBytes, &data)

	if data.ID == 0 {
		movement, err := createMovements(data)

		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, article := range data.Articles {
			var item dto.MovementArticle
			item.Amount = article.Amount
			item.Description = article.Description
			item.Title = article.Title
			item.MovementID = movement.ID

			_, err := createMovementArticle(item)

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

	var movementArticles []dto.ArticlesDropdown
	for _, article := range articles {
		var movementArticle dto.ArticlesDropdown

		movementArticle.Title = article.Title
		movementArticle.Description = article.Description
		movementArticle.Amount = article.Amount

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

	for _, article := range articles {
		stock, _ := getStockByName(article.Title)

		if stock == nil {
			err = createStock(article)

			if err != nil {
				return shared.HandleAPIError(err)
			}
		} else {
			stock.Amount += article.Amount
			err := updateStock(*stock)
			if err != nil {
				return shared.HandleAPIError(err)
			}
		}

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

func getStockByName(name string) (*structs.StockArticle, error) {
	input := dto.StockFilter{
		Title: &name,
	}
	res := &dto.GetStockResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.STOCK_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	if len(res.Data) != 0 {
		return &res.Data[0], nil
	}

	return nil, nil
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
