package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFlatRatePayment(item *structs.FlatRatePayment) (*structs.FlatRatePayment, error) {
	res := &dto.GetFlatRatePaymentResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FlatRatePayment, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFlatRatePayment(id int) (*structs.FlatRatePayment, error) {
	res := &dto.GetFlatRatePaymentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FlatRatePayment+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFlatRatePaymentList(input *dto.GetFlatRatePaymentListInputMS) ([]structs.FlatRatePayment, int, error) {
	res := &dto.GetFlatRatePaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FlatRatePayment, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFlatRatePayment(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FlatRatePayment+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateFlatRatePayment(item *structs.FlatRatePayment) (*structs.FlatRatePayment, error) {
	res := &dto.GetFlatRatePaymentResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FlatRatePayment+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
