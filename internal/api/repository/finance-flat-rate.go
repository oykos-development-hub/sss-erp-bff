package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFlatRate(item *structs.FlatRate) (*structs.FlatRate, error) {
	res := &dto.GetFlatRateResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FlatRate, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFlatRate(id int) (*structs.FlatRate, error) {
	res := &dto.GetFlatRateResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FlatRate+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFlatRateList(input *dto.GetFlatRateListInputMS) ([]structs.FlatRate, int, error) {
	res := &dto.GetFlatRateListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FlatRate, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFlatRate(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FlatRate+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateFlatRate(item *structs.FlatRate) (*structs.FlatRate, error) {
	res := &dto.GetFlatRateResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FlatRate+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
