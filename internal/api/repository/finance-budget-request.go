package repository

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateBudgetRequest(budgetItem *structs.BudgetRequest) (*structs.BudgetRequest, error) {
	res := &dto.GetBudgetRequestResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.BudgetRequest, budgetItem, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateBudgetRequest(item *structs.BudgetRequest) (*structs.BudgetRequest, error) {
	res := &dto.GetBudgetRequestResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.BudgetRequest+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetBudgetRequestList(input *dto.GetBudgetRequestListInputMS) ([]structs.BudgetRequest, error) {
	res := &dto.GetBudgetRequestListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.BudgetRequest, input, res)
	if err != nil {
		return nil, errors.WrapInternalServerError(err, "repo.GetBudgetRequestList")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetOneBudgetRequest(input *dto.GetBudgetRequestListInputMS) (budgetReq *structs.BudgetRequest, err error) {
	res := &dto.GetBudgetRequestListResponseMS{}
	_, err = makeAPIRequest("GET", repo.Config.Microservices.Finance.BudgetRequest, input, res)
	if err != nil {
		return budgetReq, errors.WrapInternalServerError(err, "repo.GetOneBudgetRequest")
	}

	if len(res.Data) == 0 {
		return budgetReq, errors.NewInternalServerError("repo.GetOneBudgetRequest: budget not sent")
	}

	return &res.Data[0], nil
}

func (repo *MicroserviceRepository) GetBudgetRequest(id int) (*structs.BudgetRequest, error) {
	res := &dto.GetBudgetRequestResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.BudgetRequest+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
