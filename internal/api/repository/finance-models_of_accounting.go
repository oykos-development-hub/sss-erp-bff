package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) UpdateModelsOfAccounting(ctx context.Context, item *structs.ModelsOfAccounting) (*structs.ModelsOfAccounting, error) {
	res := &dto.GetModelsOfAccountingResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.ModelsOfAccounting+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetModelsOfAccountingByID(id int) (*structs.ModelsOfAccounting, error) {
	res := &dto.GetModelsOfAccountingResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.ModelsOfAccounting+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetModelsOfAccountingList(filter dto.ModelsOfAccountingFilter) ([]structs.ModelsOfAccounting, int, error) {
	res := &dto.GetModelsOfAccountingListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.ModelsOfAccounting, filter, res)
	if err != nil {
		return nil, 0, errors.Wrap(err, "make api request")
	}

	return res.Data, res.Total, nil
}
