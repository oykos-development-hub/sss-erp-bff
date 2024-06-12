package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreatePaymentOrder(ctx context.Context, item *structs.PaymentOrder) (*structs.PaymentOrder, error) {
	res := &dto.GetPaymentOrderResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.PaymentOrder, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdatePaymentOrder(ctx context.Context, item *structs.PaymentOrder) (*structs.PaymentOrder, error) {
	res := &dto.GetPaymentOrderResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.PaymentOrder+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetPaymentOrderByID(id int) (*structs.PaymentOrder, error) {
	res := &dto.GetPaymentOrderResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.PaymentOrder+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetPaymentOrderList(filter dto.PaymentOrderFilter) ([]structs.PaymentOrder, int, error) {
	res := &dto.GetPaymentOrderListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.PaymentOrder, filter, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) GetAllObligations(input dto.ObligationsFilter) ([]dto.Obligation, int, error) {
	res := &dto.GetObligationsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetObligation, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeletePaymentOrder(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.PaymentOrder+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) PayPaymentOrder(ctx context.Context, input structs.PaymentOrder) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.PayPaymentOrder+"/"+strconv.Itoa(input.ID), input, nil, header)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) CancelPaymentOrder(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.CancelPaymentOrder+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}
	return nil
}
