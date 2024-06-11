package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) GetSystematizationByID(id int) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Systematization+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetSystematizations(input *dto.GetSystematizationsInput) (*dto.GetSystematizationsResponseMS, error) {
	res := &dto.GetSystematizationsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Systematization, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) UpdateSystematization(ctx context.Context, id int, data *structs.Systematization) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.Systematization+"/"+strconv.Itoa(id), data, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateSystematization(ctx context.Context, data *structs.Systematization) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.Systematization, data, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteSystematization(ctx context.Context, id int) error {
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.Systematization+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}
