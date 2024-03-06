package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFinePayment(item *structs.FinePayment) (*structs.FinePayment, error) {
	res := &dto.GetFinePaymentResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FinePayment, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFinePayment(id int) (*structs.FinePayment, error) {
	res := &dto.GetFinePaymentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FinePayment+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFinePaymentList(input *dto.GetFinePaymentListInputMS) ([]structs.FinePayment, int, error) {
	res := &dto.GetFinePaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FinePayment, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFinePayment(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FinePayment+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateFinePayment(item *structs.FinePayment) (*structs.FinePayment, error) {
	res := &dto.GetFinePaymentResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FinePayment+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
