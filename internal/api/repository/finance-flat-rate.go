package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFlatRate(ctx context.Context, item *structs.FlatRate) (*structs.FlatRate, error) {
	res := &dto.GetFlatRateResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FlatRate, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFlatRate(id int) (*structs.FlatRate, error) {
	res := &dto.GetFlatRateResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FlatRate+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFlatRateList(input *dto.GetFlatRateListInputMS) ([]structs.FlatRate, int, error) {
	res := &dto.GetFlatRateListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FlatRate, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFlatRate(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FlatRate+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateFlatRate(ctx context.Context, item *structs.FlatRate) (*structs.FlatRate, error) {
	res := &dto.GetFlatRateResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FlatRate+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
