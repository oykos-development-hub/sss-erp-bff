package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetStock(input *dto.StockFilter) ([]structs.StockArticle, *int, error) {
	res := &dto.GetStockResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.Stock, input, res)
	if err != nil {
		return nil, nil, err
	}

	return res.Data, &res.Total, nil
}

func (repo *MicroserviceRepository) GetStockByID(id int) (*structs.StockArticle, error) {
	res := &dto.GetSingleStockResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.Stock+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetMovements(input *dto.MovementFilter) ([]structs.Movement, *int, error) {
	res := &dto.GetMovementResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.Movements, input, res)
	if err != nil {
		return nil, nil, err
	}

	return res.Data, &res.Total, nil
}

func (repo *MicroserviceRepository) CreateMovements(input structs.OrderAssetMovementItem) (*structs.Movement, error) {
	res := &dto.GetSingleMovementResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.Movements, input, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateStock(input dto.MovementArticle) error {
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.Stock, input, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateStock(input structs.StockArticle) error {
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Accounting.Stock+"/"+strconv.Itoa(input.ID), input, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateMovementArticle(input dto.MovementArticle) (*dto.MovementArticle, error) {
	res := &dto.GetSingleMovementArticleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Accounting.MovementArticles, input, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteMovement(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Accounting.Movements+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateMovements(input structs.OrderAssetMovementItem) (*structs.Movement, error) {
	res := &dto.GetSingleMovementResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Accounting.Movements+"/"+strconv.Itoa(input.ID), input, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetMovementByID(id int) (*structs.Movement, error) {
	res := &dto.GetSingleMovementResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.Movements+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return res.Data, nil
}
