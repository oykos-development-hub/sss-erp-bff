package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"fmt"
	"strconv"
	"time"
)

func (repo *MicroserviceRepository) UpdateOrderListItem(ctx context.Context, id int, orderListItem *structs.OrderListItem) (*structs.OrderListItem, error) {
	res := &dto.GetOrderListResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Accounting.OrderLists+"/"+strconv.Itoa(id), orderListItem, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) SendOrderListToFinance(ctx context.Context, id int) error {

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Accounting.OrderListSendToFinance+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) CreateOrderListItem(ctx context.Context, orderListItem *structs.OrderListItem) (*structs.OrderListItem, error) {
	res := &dto.GetOrderListResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.OrderLists, orderListItem, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateOrderProcurementArticle(orderProcurementArticleItem *structs.OrderProcurementArticleItem) (*structs.OrderProcurementArticleItem, error) {
	res := &dto.GetOrderProcurementArticleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.OrderProcurementArticles, orderProcurementArticleItem, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteOrderProcurementArticle(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Accounting.OrderProcurementArticles+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) DeleteOrderList(ctx context.Context, id int) error {
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Accounting.OrderLists+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) CreateOrderListProcurementArticles(orderListID int, data structs.OrderListInsertItem) error {
	for _, article := range data.Articles {
		newArticle := structs.OrderProcurementArticleItem{
			Amount:  article.Amount,
			OrderID: orderListID,
		}
		if article.ID != 0 {
			newArticle.ArticleID = article.ID
			article, err := repo.GetProcurementArticle(article.ID)
			if err != nil {
				return errors.Wrap(err, "repo get procurement article")
			}

			procurement, err := repo.GetProcurementItem(article.PublicProcurementID)
			if err != nil {
				return errors.Wrap(err, "repo get procurement item")
			}

			plan, err := repo.GetProcurementPlan(procurement.PlanID)

			if err != nil {
				return errors.Wrap(err, "repo get procurement plan")
			}

			newArticle.Year = plan.Year

		} else {
			newArticle.Title = article.Title
			newArticle.Description = article.Description
			newArticle.NetPrice = article.NetPrice
			newArticle.VatPercentage = article.VatPercentage
			newArticle.Amount = article.Amount
			newArticle.Year = strconv.Itoa(time.Now().Year())
		}

		_, err := repo.CreateOrderProcurementArticle(&newArticle)
		if err != nil {
			return errors.Wrap(err, "repo create order procurement article")
		}
	}
	return nil
}

func (repo *MicroserviceRepository) GetOrderProcurementArticles(input *dto.GetOrderProcurementArticleInput) (*dto.GetOrderProcurementArticlesResponseMS, error) {
	res := &dto.GetOrderProcurementArticlesResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.OrderProcurementArticles, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetOrderListByID(id int) (*structs.OrderListItem, error) {
	res := &dto.GetOrderListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.OrderLists+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetOrderLists(input *dto.GetOrderListInput) (*dto.GetOrderListsResponseMS, error) {
	res := &dto.GetOrderListsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.OrderLists, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetOrderProcurementArticleByID(id int) (*structs.OrderProcurementArticleItem, error) {
	res := &dto.GetOrderProcurementArticleResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.OrderProcurementArticles+"/"+strconv.Itoa(id), nil, res)
	fmt.Println(repo.Config.Microservices.Accounting.OrderProcurementArticles + "/" + strconv.Itoa(id))
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateOrderProcurementArticle(item *structs.OrderProcurementArticleItem) (*structs.OrderProcurementArticleItem, error) {
	res := &dto.GetOrderProcurementArticleResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Accounting.OrderProcurementArticles+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) AddOnStock(stock []structs.StockArticle, article structs.OrderProcurementArticleItem, organizationUnitID int, exception bool) error {

	var id int
	var err error

	if len(stock) == 0 {
		input := dto.MovementArticle{
			Amount:             article.Amount,
			Year:               article.Year,
			Description:        article.Description,
			Title:              article.Title,
			NetPrice:           article.NetPrice,
			VatPercentage:      article.VatPercentage,
			OrganizationUnitID: organizationUnitID,
			Exception:          exception,
		}

		id, err = repo.CreateStock(input)
		if err != nil {
			return errors.Wrap(err, "repo create stock")
		}
	} else {
		id = stock[0].ID
		stock[0].Amount += article.Amount
		err := repo.UpdateStock(stock[0])
		if err != nil {
			return errors.Wrap(err, "repo update stock")
		}
	}

	err = repo.CreateStockOrderArticle(dto.StockOrderArticle{
		StockID:   id,
		ArticleID: article.ID,
	})

	if err != nil {
		return errors.Wrap(err, "repo create stock order article")
	}

	return nil
}
