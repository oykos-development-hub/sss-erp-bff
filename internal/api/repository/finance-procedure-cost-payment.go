package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateProcedureCostPayment(ctx context.Context, item *structs.ProcedureCostPayment) (*structs.ProcedureCostPayment, error) {
	res := &dto.GetProcedureCostPaymentResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.ProcedureCostPayment, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcedureCostPayment(id int) (*structs.ProcedureCostPayment, error) {
	res := &dto.GetProcedureCostPaymentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.ProcedureCostPayment+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcedureCostPaymentList(input *dto.GetProcedureCostPaymentListInputMS) ([]structs.ProcedureCostPayment, int, error) {
	res := &dto.GetProcedureCostPaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.ProcedureCostPayment, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteProcedureCostPayment(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.ProcedureCostPayment+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateProcedureCostPayment(ctx context.Context, item *structs.ProcedureCostPayment) (*structs.ProcedureCostPayment, error) {
	res := &dto.GetProcedureCostPaymentResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.ProcedureCostPayment+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
