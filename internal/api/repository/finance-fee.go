package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFee(ctx context.Context, item *structs.Fee) (*structs.Fee, error) {
	res := &dto.GetFeeResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Fee, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFee(id int) (*structs.Fee, error) {
	res := &dto.GetFeeResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Fee+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFeeList(input *dto.GetFeeListInputMS) ([]structs.Fee, int, error) {
	res := &dto.GetFeeListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Fee, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFee(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Fee+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateFee(ctx context.Context, item *structs.Fee) (*structs.Fee, error) {
	res := &dto.GetFeeResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Fee+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
