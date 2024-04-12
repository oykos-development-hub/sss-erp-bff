package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateDepositPayment(item *structs.DepositPayment) (*structs.DepositPayment, error) {
	res := &dto.GetDepositPaymentResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.DepositPayment, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateDepositPayment(item *structs.DepositPayment) (*structs.DepositPayment, error) {
	res := &dto.GetDepositPaymentResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.DepositPayment+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetDepositPaymentByID(id int) (*structs.DepositPayment, error) {
	res := &dto.GetDepositPaymentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPayment+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetDepositPaymentList(filter dto.DepositPaymentFilter) ([]structs.DepositPayment, int, error) {
	res := &dto.GetDepositPaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.DepositPayment, filter, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteDepositPayment(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.DepositPayment+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
