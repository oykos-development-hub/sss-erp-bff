package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateEnforcedPayment(ctx context.Context, item *structs.EnforcedPayment) (*structs.EnforcedPayment, error) {
	res := &dto.GetEnforcedPaymentResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.EnforcedPayment, item, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateEnforcedPayment(ctx context.Context, item *structs.EnforcedPayment) (*structs.EnforcedPayment, error) {
	res := &dto.GetEnforcedPaymentResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.EnforcedPayment+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetEnforcedPaymentByID(id int) (*structs.EnforcedPayment, error) {
	res := &dto.GetEnforcedPaymentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.EnforcedPayment+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetEnforcedPaymentList(filter dto.EnforcedPaymentFilter) ([]structs.EnforcedPayment, int, error) {
	res := &dto.GetEnforcedPaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.EnforcedPayment, filter, res)
	if err != nil {
		return nil, 0, errors.Wrap(err, "make api request")
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) ReturnEnforcedPayment(ctx context.Context, input structs.EnforcedPayment) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.ReturnEnforcedPayment+"/"+strconv.Itoa(input.ID), input, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}
	return nil
}
