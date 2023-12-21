package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
	"time"
)

func (repo *MicroserviceRepository) UpdateOrderListItem(id int, orderListItem *structs.OrderListItem) (*structs.OrderListItem, error) {
	res := &dto.GetOrderListResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Accounting.ORDER_LISTS+"/"+strconv.Itoa(id), orderListItem, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateOrderListItem(orderListItem *structs.OrderListItem) (*structs.OrderListItem, error) {
	res := &dto.GetOrderListResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.ORDER_LISTS, orderListItem, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateOrderProcurementArticle(orderProcurementArticleItem *structs.OrderProcurementArticleItem) (*structs.OrderProcurementArticleItem, error) {
	res := &dto.GetOrderProcurementArticleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.ORDER_PROCUREMENT_ARTICLES, orderProcurementArticleItem, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteOrderProcurementArticle(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Accounting.ORDER_PROCUREMENT_ARTICLES+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) DeleteOrderList(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Accounting.ORDER_LISTS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateOrderListProcurementArticles(orderListId int, data structs.OrderListInsertItem) error {
	for _, article := range data.Articles {
		newArticle := structs.OrderProcurementArticleItem{
			Amount:  article.Amount,
			OrderId: orderListId,
		}
		if article.Id != 0 {
			newArticle.ArticleId = article.Id
			article, err := repo.GetProcurementArticle(article.Id)
			if err != nil {
				return err
			}

			procurement, err := repo.GetProcurementItem(article.PublicProcurementId)
			if err != nil {
				return err
			}

			plan, err := repo.GetProcurementPlan(procurement.PlanId)

			if err != nil {
				return err
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
			return err
		}
	}
	return nil
}

func (repo *MicroserviceRepository) GetOrderProcurementArticles(input *dto.GetOrderProcurementArticleInput) (*dto.GetOrderProcurementArticlesResponseMS, error) {
	res := &dto.GetOrderProcurementArticlesResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.ORDER_PROCUREMENT_ARTICLES, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetOrderListById(id int) (*structs.OrderListItem, error) {
	res := &dto.GetOrderListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.ORDER_LISTS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetOrderLists(input *dto.GetOrderListInput) (*dto.GetOrderListsResponseMS, error) {
	res := &dto.GetOrderListsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.ORDER_LISTS, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetOrderProcurementArticleByID(id int) (*structs.OrderProcurementArticleItem, error) {
	res := &dto.GetOrderProcurementArticleResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.ORDER_PROCUREMENT_ARTICLES+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateOrderProcurementArticle(item *structs.OrderProcurementArticleItem) (*structs.OrderProcurementArticleItem, error) {
	res := &dto.GetOrderProcurementArticleResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Accounting.ORDER_PROCUREMENT_ARTICLES+"/"+strconv.Itoa(item.Id), item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) AddOnStock(stock []structs.StockArticle, article structs.OrderProcurementArticleItem, organizationUnitID int) error {

	if len(stock) == 0 {
		input := dto.MovementArticle{
			Amount:             article.Amount,
			Year:               article.Year,
			Description:        article.Description,
			Title:              article.Title,
			NetPrice:           article.NetPrice,
			VatPercentage:      article.VatPercentage,
			OrganizationUnitID: organizationUnitID,
			Exception:          true,
		}

		err := repo.CreateStock(input)
		if err != nil {
			return err
		}
	} else {
		stock[0].Amount += article.Amount
		err := repo.UpdateStock(stock[0])
		if err != nil {
			return err
		}
	}
	return nil
}
