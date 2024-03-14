package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFeePayment(item *structs.FeePayment) (*structs.FeePayment, error) {
	res := &dto.GetFeePaymentResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FeePayment, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFeePayment(id int) (*structs.FeePayment, error) {
	res := &dto.GetFeePaymentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FeePayment+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFeePaymentList(input *dto.GetFeePaymentListInputMS) ([]structs.FeePayment, int, error) {
	res := &dto.GetFeePaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FeePayment, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFeePayment(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FeePayment+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateFeePayment(item *structs.FeePayment) (*structs.FeePayment, error) {
	res := &dto.GetFeePaymentResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FeePayment+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
