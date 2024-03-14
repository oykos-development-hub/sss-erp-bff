package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFee(item *structs.Fee) (*structs.Fee, error) {
	res := &dto.GetFeeResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Fee, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFee(id int) (*structs.Fee, error) {
	res := &dto.GetFeeResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Fee+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFeeList(input *dto.GetFeeListInputMS) ([]structs.Fee, int, error) {
	res := &dto.GetFeeListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Fee, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFee(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Fee+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateFee(item *structs.Fee) (*structs.Fee, error) {
	res := &dto.GetFeeResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Fee+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
