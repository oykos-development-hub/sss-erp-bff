package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateBudgetRequest(ctx context.Context, budgetItem *structs.BudgetRequest) (*structs.BudgetRequest, error) {
	res := &dto.GetBudgetRequestResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.BudgetRequest, budgetItem, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateBudgetRequest(ctx context.Context, item *structs.BudgetRequest) (*structs.BudgetRequest, error) {
	res := &dto.GetBudgetRequestResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.BudgetRequest+"/"+strconv.Itoa(item.ID), item, res, header)
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
		return budgetReq, errors.NewNotFoundError("budget not found")
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
