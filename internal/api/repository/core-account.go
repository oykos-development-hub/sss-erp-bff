package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) DeleteAccount(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Core.Account+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetAccountItems(filters *dto.GetAccountsFilter) (*dto.GetAccountItemListResponseMS, error) {
	res := &dto.GetAccountItemListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Account, filters, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return res, nil
}

func (repo *MicroserviceRepository) GetLatestVersionOfAccounts() (int, error) {
	accountItems, err := repo.GetAccountItems(nil)
	if err != nil {
		return -1, err
	}

	if accountItems.Total == 0 {
		return -1, errors.New("there is no accounts in the system")
	}

	return accountItems.Data[0].Version, nil
}

func (repo *MicroserviceRepository) GetAccountItemByID(id int) (*structs.AccountItem, error) {
	res := &dto.GetAccountItemResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Account+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateAccountItemList(ctx context.Context, accountItemList []structs.AccountItem) ([]*structs.AccountItem, error) {
	res := &dto.InsertAccountItemListResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.Account, accountItemList, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateAccountItem(ctx context.Context, id int, accountItem *structs.AccountItem) (*structs.AccountItem, error) {
	res := &dto.GetAccountItemResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.Account+"/"+strconv.Itoa(id), accountItem, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}
