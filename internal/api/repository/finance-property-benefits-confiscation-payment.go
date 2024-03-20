package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreatePropBenConfPayment(item *structs.PropBenConfPayment) (*structs.PropBenConfPayment, error) {
	res := &dto.GetPropBenConfPaymentResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.PropBenConfPayment, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetPropBenConfPayment(id int) (*structs.PropBenConfPayment, error) {
	res := &dto.GetPropBenConfPaymentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.PropBenConfPayment+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetPropBenConfPaymentList(input *dto.GetPropBenConfPaymentListInputMS) ([]structs.PropBenConfPayment, int, error) {
	res := &dto.GetPropBenConfPaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.PropBenConfPayment, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeletePropBenConfPayment(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.PropBenConfPayment+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdatePropBenConfPayment(item *structs.PropBenConfPayment) (*structs.PropBenConfPayment, error) {
	res := &dto.GetPropBenConfPaymentResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.PropBenConfPayment+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
