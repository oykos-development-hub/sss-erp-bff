package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFeePayment(ctx context.Context, item *structs.FeePayment) (*structs.FeePayment, error) {
	res := &dto.GetFeePaymentResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FeePayment, item, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFeePayment(id int) (*structs.FeePayment, error) {
	res := &dto.GetFeePaymentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FeePayment+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFeePaymentList(input *dto.GetFeePaymentListInputMS) ([]structs.FeePayment, int, error) {
	res := &dto.GetFeePaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FeePayment, input, res)
	if err != nil {
		return nil, 0, errors.Wrap(err, "make api request")
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFeePayment(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FeePayment+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateFeePayment(ctx context.Context, item *structs.FeePayment) (*structs.FeePayment, error) {
	res := &dto.GetFeePaymentResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FeePayment+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}
