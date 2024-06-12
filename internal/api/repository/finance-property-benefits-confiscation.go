package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreatePropBenConf(ctx context.Context, item *structs.PropBenConf) (*structs.PropBenConf, error) {
	res := &dto.GetPropBenConfResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.PropBenConf, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetPropBenConf(id int) (*structs.PropBenConf, error) {
	res := &dto.GetPropBenConfResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.PropBenConf+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetPropBenConfList(input *dto.GetPropBenConfListInputMS) ([]structs.PropBenConf, int, error) {
	res := &dto.GetPropBenConfListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.PropBenConf, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeletePropBenConf(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.PropBenConf+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdatePropBenConf(ctx context.Context, item *structs.PropBenConf) (*structs.PropBenConf, error) {
	res := &dto.GetPropBenConfResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.PropBenConf+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
