package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetUserAccounts(input *dto.GetUserAccountListInput) (*dto.GetUserAccountListResponseMS, error) {
	res := &dto.GetUserAccountListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.USER_ACCOUNTS, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) UpdateUserAccount(userID int, user structs.UserAccounts) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.USER_ACCOUNTS+"/"+strconv.Itoa(userID), user, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetUserAccountById(id int) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.USER_ACCOUNTS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateUserAccount(user structs.UserAccounts) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.USER_ACCOUNTS, user, res)

	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeactivateUserAccount(userID int) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	user := dto.DeactivateUserAccount{
		Active: false,
	}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.USER_ACCOUNTS+"/"+strconv.Itoa(userID), user, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteUserAccount(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Core.USER_ACCOUNTS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
