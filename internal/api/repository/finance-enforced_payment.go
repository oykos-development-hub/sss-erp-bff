package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateEnforcedPayment(item *structs.EnforcedPayment) (*structs.EnforcedPayment, error) {
	res := &dto.GetEnforcedPaymentResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.EnforcedPayment, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateEnforcedPayment(item *structs.EnforcedPayment) (*structs.EnforcedPayment, error) {
	res := &dto.GetEnforcedPaymentResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.EnforcedPayment+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetEnforcedPaymentByID(id int) (*structs.EnforcedPayment, error) {
	res := &dto.GetEnforcedPaymentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.EnforcedPayment+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetEnforcedPaymentList(filter dto.EnforcedPaymentFilter) ([]structs.EnforcedPayment, int, error) {
	res := &dto.GetEnforcedPaymentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.EnforcedPayment, filter, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) ReturnEnforcedPayment(input structs.EnforcedPayment) error {
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.ReturnEnforcedPayment+"/"+strconv.Itoa(input.ID), input, nil)
	if err != nil {
		return err
	}
	return nil
}
