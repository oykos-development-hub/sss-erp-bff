package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateDepositPayment(ctx context.Context, item *structs.DepositPayment) (*structs.DepositPayment, error) {
	res := &dto.GetDepositPaymentResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.DepositPayment, item, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateDepositPayment(ctx context.Context, item *structs.DepositPayment) (*structs.DepositPayment, error) {
	res := &dto.GetDepositPaymentResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.DepositPayment+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetDepositPaymentByID(id int) (*structs.DepositPayment, error) {
	res := &dto.GetDepositPaymentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPayment+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetDepositPaymentList(filter dto.DepositPaymentFilter) ([]structs.DepositPayment, int, error) {
	res := &dto.GetDepositPaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPayment, filter, res)
	if err != nil {
		return nil, 0, errors.Wrap(err, "make api request")
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) GetInitialState(filter dto.DepositInitialStateFilter) ([]structs.DepositPayment, error) {
	res := &dto.GetDepositPaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetInitialState, filter, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetDepositPaymentCaseNumber(caseNumber string, bankAccount string) (*structs.DepositPayment, error) {
	res := &dto.GetDepositPaymentResponseMS{}
	filter := dto.DepositPaymentFilter{
		CaseuNumber:       &caseNumber,
		SourceBankAccount: &bankAccount,
	}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPaymentCaseNumber, filter, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteDepositPayment(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.DepositPayment+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetCaseNumber(organizationUnitID int, bankAccount string) ([]structs.DepositPayment, error) {
	res := &dto.GetDepositPaymentListResponseMS{}
	filter := dto.DepositPaymentFilter{
		OrganizationUnitID: &organizationUnitID,
		SourceBankAccount:  &bankAccount,
	}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.GetDepositPaymentCaseNumber, filter, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}
