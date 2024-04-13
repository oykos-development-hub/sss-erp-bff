package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateDepositPaymentOrder(item *structs.DepositPaymentOrder) (*structs.DepositPaymentOrder, error) {
	res := &dto.GetDepositPaymentOrderResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.DepositPaymentOrder, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateDepositPaymentOrder(item *structs.DepositPaymentOrder) (*structs.DepositPaymentOrder, error) {
	res := &dto.GetDepositPaymentOrderResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.DepositPaymentOrder+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetDepositPaymentOrderByID(id int) (*structs.DepositPaymentOrder, error) {
	res := &dto.GetDepositPaymentOrderResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPaymentOrder+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetDepositPaymentOrderList(filter dto.DepositPaymentOrderFilter) ([]structs.DepositPaymentOrder, int, error) {
	res := &dto.GetDepositPaymentOrderListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPaymentOrder, filter, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteDepositPaymentOrder(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.DepositPaymentOrder+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetDepositPaymentAdditionalExpenses(input *dto.DepositPaymentAdditionalExpensesListInputMS) ([]structs.DepositPaymentAdditionalExpenses, int, error) {
	res := &dto.GetDepositPaymentAdditionalExpensesListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPaymentAdditionalExpenses, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) PayDepositPaymentOrder(input structs.DepositPaymentOrder) error {
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.PayDepositPaymentOrder+"/"+strconv.Itoa(input.ID), input, nil)
	if err != nil {
		return err
	}
	return nil
}
