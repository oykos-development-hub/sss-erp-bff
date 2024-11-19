package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) GetStock(input *dto.StockFilter) ([]structs.StockArticle, *int, error) {
	res := &dto.GetStockResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.Stock, input, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "make api request")
	}

	return res.Data, &res.Total, nil
}

func (repo *MicroserviceRepository) GetStockReport(input *dto.StockFilter) ([]structs.StockArticle, error) {
	res := &dto.GetStockResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.StockReport, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetStockByID(id int) (*structs.StockArticle, error) {
	res := &dto.GetSingleStockResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.Stock+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetMovements(input *dto.MovementFilter) ([]structs.Movement, *int, error) {
	res := &dto.GetMovementResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.Movements, input, res)
	if err != nil {
		return nil, nil, errors.Wrap(err, "make api request")
	}

	return res.Data, &res.Total, nil
}

func (repo *MicroserviceRepository) CreateMovements(ctx context.Context, input structs.OrderAssetMovementItem) (*structs.Movement, error) {
	res := &dto.GetSingleMovementResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.Movements, input, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateStock(input dto.MovementArticle) (int, error) {

	res := &dto.GetSingleStockResponseMS{}

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.Stock, input, res)
	if err != nil {
		return 0, errors.Wrap(err, "make api request")
	}

	return res.Data.ID, nil
}

func (repo *MicroserviceRepository) CreateStockOrderArticle(input dto.StockOrderArticle) error {
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.StockOrderArticle, input, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateStock(input structs.StockArticle) error {
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Accounting.Stock+"/"+strconv.Itoa(input.ID), input, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) CreateMovementArticle(input dto.MovementArticle) (*dto.MovementArticle, error) {
	res := &dto.GetSingleMovementArticleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.MovementArticles, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteMovement(ctx context.Context, id int) error {

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Accounting.Movements+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateMovements(ctx context.Context, input structs.OrderAssetMovementItem) (*structs.Movement, error) {
	res := &dto.GetSingleMovementResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Accounting.Movements+"/"+strconv.Itoa(input.ID), input, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetMovementByID(id int) (*structs.Movement, error) {
	res := &dto.GetSingleMovementResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.Movements+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetMovementArticles(id int) ([]dto.MovementArticle, error) {
	input := dto.MovementArticlesFilter{
		MovementID: &id,
	}
	res := &dto.GetMovementArticleResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.MovementArticles, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}
