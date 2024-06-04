package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) StockOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
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

	if page != nil && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}
	if size != nil && size.(int) > 0 {
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
		return errors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	if dateOk && date != "" {
		statusReceive := "Receive"
		orders, err := r.Repo.GetOrderLists(&dto.GetOrderListInput{
			DateSystem:         &date,
			Status:             &statusReceive,
			OrganizationUnitID: organizationUnitID,
		})

		if err != nil {
			return errors.HandleAPIError(err)
		}

		for _, order := range orders.Data {
			orderArticles, err := r.Repo.GetOrderProcurementArticles(&dto.GetOrderProcurementArticleInput{OrderID: &order.ID})
			if err != nil {
				return errors.HandleAPIError(err)
			}
			for _, article := range orderArticles.Data {
				flag := false

				if article.ArticleID != 0 {
					currentArticle, err := r.Repo.GetProcurementArticle(article.ArticleID)

					if err != nil {
						return errors.HandleAPIError(err)
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
						ID:            article.ID,
						NetPrice:      article.NetPrice,
						VatPercentage: article.VatPercentage,
					}
					articles = append(articles, newArticle)
				}
			}
		}

		movementArticles, err := r.Repo.GetMovementArticleList(dto.OveralSpendingFilter{
			EndDate:            &date,
			OrganizationUnitID: organizationUnitID,
		})

		if err != nil {
			return errors.HandleAPIError(err)
		}

		for i := 0; i < len(movementArticles); i++ {
			for j := 0; j < len(articles); j++ {
				if movementArticles[i].Title == articles[j].Title && movementArticles[i].Year == articles[j].Year && officeInOrgUnit(r.Repo, movementArticles[i].OfficeID, *organizationUnitID) {
					articles[j].Amount -= movementArticles[i].Amount
				}
			}
		}

	} else {
		input.OrganizationUnitID = organizationUnitID

		articleList, total, err := r.Repo.GetStock(&input)
		if err != nil {
			return errors.HandleAPIError(err)
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

func (r *Resolver) MovementOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		input dto.MovementFilter
	)

	size := params.Args["size"]
	page := params.Args["page"]
	OfficeID, officeOk := params.Args["office_id"].(int)
	userID, userOk := params.Args["recipient_user_id"].(int)
	sortByDateOrder, sortByDateOrderOK := params.Args["sort_by_date_order"].(string)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return errors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	input.OrganizationUnitID = organizationUnitID

	if page != nil && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}
	if size != nil && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}

	if officeOk && OfficeID != 0 {
		input.OfficeID = &OfficeID
	}

	if userOk && userID != 0 {
		input.RecipientUserID = &userID
	}

	if sortByDateOrderOK && sortByDateOrder != "" {
		input.SortByDateOrder = &sortByDateOrder
	}

	movementList, total, err := r.Repo.GetMovements(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	var response []dto.MovementResponse

	for _, movement := range movementList {
		var item dto.MovementResponse
		item.ID = movement.ID
		item.DateOrder = movement.DateOrder
		item.Description = movement.Description

		officeItem, err := r.Repo.GetDropdownSettingByID(movement.OfficeID)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		item.Office.Title = officeItem.Title
		item.Office.ID = officeItem.ID

		userItem, err := r.Repo.GetUserProfileByID(movement.RecipientUserID)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		item.RecipientUser.Title = userItem.FirstName + " " + userItem.LastName
		item.RecipientUser.ID = userItem.ID

		articles, err := r.Repo.GetMovementArticles(item.ID)
		if err != nil {
			return nil, err
		}

		var movementArticles []dto.MovementArticle
		for _, article := range articles {
			stockArticle, err := r.Repo.GetStockByID(article.StockID)

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

func (r *Resolver) MovementDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]

	response, err := buildMovementDetailsResponse(r.Repo, id.(int))

	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

func (r *Resolver) MovementArticlesResolver(params graphql.ResolveParams) (interface{}, error) {
	var response []string
	title, titleOK := params.Args["title"].(string)
	var titleFilter *string

	if titleOK && title != "" {
		titleFilter = &title
	} else {
		titleFilter = nil
	}

	articles, err := r.Repo.GetMovementArticleList(dto.OveralSpendingFilter{
		Title: titleFilter,
	})
	if err != nil {
		return errors.HandleAPIError(err)
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

func (r *Resolver) OrderListAssetMovementResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.OrderAssetMovementItem

	dataBytes, _ := json.Marshal(params.Args["data"])
	_ = json.Unmarshal(dataBytes, &data)

	if data.ID == 0 {
		organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitID == nil {
			return errors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
		}

		data.OrganizationUnitID = *organizationUnitID

		movement, err := r.Repo.CreateMovements(data)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		for _, article := range data.Articles {
			item := dto.MovementArticle{
				StockID:    article.ID,
				MovementID: movement.ID,
				Amount:     article.Quantity,
			}

			if err != nil {
				return errors.HandleAPIError(err)
			}

			_, err = r.Repo.CreateMovementArticle(item)

			if err != nil {
				return errors.HandleAPIError(err)
			}

			stockArticle, err := r.Repo.GetStockByID(article.ID)

			if err != nil {
				return errors.HandleAPIError(err)
			}

			stockArticle.Amount -= article.Quantity

			err = r.Repo.UpdateStock(*stockArticle)

			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
	} else {
		_, err := r.Repo.UpdateMovements(data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You created movement!",
	}, nil
}

func buildMovementDetailsResponse(r repository.MicroserviceRepositoryInterface, id int) (*dto.MovementDetailsResponse, error) {
	var item dto.MovementDetailsResponse

	movement, err := r.GetMovementByID(id)
	if err != nil {
		return nil, err
	}

	item.ID = movement.ID
	item.DateOrder = movement.DateOrder
	item.Description = movement.Description

	officeItem, err := r.GetDropdownSettingByID(movement.OfficeID)

	if err != nil {
		return nil, err
	}

	item.Office.Title = officeItem.Title
	item.Office.ID = officeItem.ID

	userItem, err := r.GetUserProfileByID(movement.RecipientUserID)

	if err != nil {
		return nil, err
	}

	item.RecipientUser.Title = userItem.FirstName + " " + userItem.LastName
	item.RecipientUser.ID = userItem.ID

	articles, err := r.GetMovementArticles(item.ID)
	if err != nil {
		return nil, err
	}

	if movement.FileID != 0 {
		file, err := r.GetFileByID(movement.FileID)

		if err != nil {
			return nil, err
		}
		item.File.ID = file.ID
		item.File.Name = file.Name
		item.File.Type = *file.Type
	}

	var movementArticles []dto.MovementArticle
	for _, article := range articles {
		stockArticle, err := r.GetStockByID(article.StockID)

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

func (r *Resolver) MovementDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	orderList, err := r.Repo.GetMovementByID(id)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	if orderList.FileID != 0 {
		err := r.Repo.DeleteFile(orderList.FileID)

		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	articles, err := r.Repo.GetMovementArticles(id)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return errors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	for _, article := range articles {
		stock, err := r.Repo.GetStockByID(article.StockID)

		if err != nil {
			return errors.HandleAPIError(err)
		}

		article.OrganizationUnitID = *organizationUnitID

		stock.Amount += article.Amount
		err = r.Repo.UpdateStock(*stock)
		if err != nil {
			return errors.HandleAPIError(err)
		}

	}

	err = r.Repo.DeleteMovement(id)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted movement!",
	}, nil
}

func officeInOrgUnit(r repository.MicroserviceRepositoryInterface, OfficeID int, OrganizationUnitID int) bool {
	orgUnitString := strconv.Itoa(OrganizationUnitID)

	res, err := r.GetOfficeDropdownSettings(&dto.GetOfficesOfOrganizationInput{
		Value: &orgUnitString,
	})

	if err != nil {
		return false
	}

	for _, office := range res.Data {
		if office.ID == OfficeID {
			return true
		}
	}
	return false
}
