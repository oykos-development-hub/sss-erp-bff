package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateDepositPaymentOrder(ctx context.Context, item *structs.DepositPaymentOrder) (*structs.DepositPaymentOrder, error) {
	res := &dto.GetDepositPaymentOrderResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.DepositPaymentOrder, item, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateDepositPaymentOrder(ctx context.Context, item *structs.DepositPaymentOrder) (*structs.DepositPaymentOrder, error) {
	res := &dto.GetDepositPaymentOrderResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.DepositPaymentOrder+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetDepositPaymentOrderByID(id int) (*structs.DepositPaymentOrder, error) {
	res := &dto.GetDepositPaymentOrderResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPaymentOrder+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetDepositPaymentOrderList(filter dto.DepositPaymentOrderFilter) ([]structs.DepositPaymentOrder, int, error) {
	res := &dto.GetDepositPaymentOrderListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPaymentOrder, filter, res)
	if err != nil {
		return nil, 0, errors.Wrap(err, "make api request")
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteDepositPaymentOrder(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.DepositPaymentOrder+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetDepositPaymentAdditionalExpenses(input *dto.DepositPaymentAdditionalExpensesListInputMS) ([]structs.DepositPaymentAdditionalExpenses, int, error) {
	res := &dto.GetDepositPaymentAdditionalExpensesListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPaymentAdditionalExpenses, input, res)
	if err != nil {
		return nil, 0, errors.Wrap(err, "make api request")
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) PayDepositPaymentOrder(ctx context.Context, input structs.DepositPaymentOrder) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.PayDepositPaymentOrder+"/"+strconv.Itoa(input.ID), input, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}
	return nil
}
