package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreatePaymentOrder(item *structs.PaymentOrder) (*structs.PaymentOrder, error) {
	res := &dto.GetPaymentOrderResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.PaymentOrder, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdatePaymentOrder(item *structs.PaymentOrder) (*structs.PaymentOrder, error) {
	res := &dto.GetPaymentOrderResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.PaymentOrder+"/"+strconv.Itoa(item.ID), item, res)
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

func (repo *MicroserviceRepository) DeletePaymentOrder(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.PaymentOrder+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

/*
func (repo *MicroserviceRepository) GetPaymentAdditionalExpenses(input *dto.PaymentAdditionalExpensesListInputMS) ([]structs.PaymentAdditionalExpenses, int, error) {
	res := &dto.GetPaymentAdditionalExpensesListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.PaymentAdditionalExpenses, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}
*/
/*

func (repo *MicroserviceRepository) PayPaymentOrder(input structs.PaymentOrder) error {
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.PayPaymentOrder+"/"+strconv.Itoa(input.ID), input, nil)
	if err != nil {
		return err
	}
	return nil
}
*/
