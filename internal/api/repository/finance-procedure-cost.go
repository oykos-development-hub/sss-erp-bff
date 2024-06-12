package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateProcedureCost(ctx context.Context, item *structs.ProcedureCost) (*structs.ProcedureCost, error) {
	res := &dto.GetProcedureCostResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.ProcedureCost, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcedureCost(id int) (*structs.ProcedureCost, error) {
	res := &dto.GetProcedureCostResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.ProcedureCost+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcedureCostList(input *dto.GetProcedureCostListInputMS) ([]structs.ProcedureCost, int, error) {
	res := &dto.GetProcedureCostListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.ProcedureCost, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteProcedureCost(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.ProcedureCost+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateProcedureCost(ctx context.Context, item *structs.ProcedureCost) (*structs.ProcedureCost, error) {
	res := &dto.GetProcedureCostResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.ProcedureCost+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
