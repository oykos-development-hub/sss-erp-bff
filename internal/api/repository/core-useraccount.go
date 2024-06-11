package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) GetUserAccounts(input *dto.GetUserAccountListInput) (*dto.GetUserAccountListResponseMS, error) {
	res := &dto.GetUserAccountListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.UserAccounts, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) UpdateUserAccount(ctx context.Context, userID int, user structs.UserAccounts) (*structs.UserAccounts, error) {
	header := make(map[string]string)
	header["UserID"] = ctx.Value(config.LoggedInAccountKey).(string)

	res := &dto.GetUserAccountResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.UserAccounts+"/"+strconv.Itoa(userID), user, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetUserAccountByID(id int) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.UserAccounts+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateUserAccount(ctx context.Context, user structs.UserAccounts) (*structs.UserAccounts, error) {
	header := make(map[string]string)
	header["UserID"] = ctx.Value(config.LoggedInAccountKey).(string)

	res := &dto.GetUserAccountResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.UserAccounts, user, res, header)

	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeactivateUserAccount(ctx context.Context, userID int) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	user := dto.DeactivateUserAccount{
		Active: false,
	}

	header := make(map[string]string)
	header["UserID"] = ctx.Value(config.LoggedInAccountKey).(string)
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.UserAccounts+"/"+strconv.Itoa(userID), user, res, header)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteUserAccount(ctx context.Context, id int) error {
	header := make(map[string]string)
	header["UserID"] = ctx.Value(config.LoggedInAccountKey).(string)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Core.UserAccounts+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}
