package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) DeleteAccount(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Core.Account+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetAccountItems(filters *dto.GetAccountsFilter) (*dto.GetAccountItemListResponseMS, error) {
	res := &dto.GetAccountItemListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Account, filters, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (repo *MicroserviceRepository) GetAccountItemByID(id int) (*structs.AccountItem, error) {
	res := &dto.GetAccountItemResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Account+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateAccountItem(accountItem *structs.AccountItem) (*structs.AccountItem, error) {
	res := &dto.GetAccountItemResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.Account, accountItem, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateAccountItem(id int, accountItem *structs.AccountItem) (*structs.AccountItem, error) {
	res := &dto.GetAccountItemResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.Account+"/"+strconv.Itoa(id), accountItem, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
