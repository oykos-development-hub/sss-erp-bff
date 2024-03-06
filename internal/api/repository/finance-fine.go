package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFine(item *structs.Fine) (*structs.Fine, error) {
	res := &dto.GetFineResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Fine, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFine(id int) (*structs.Fine, error) {
	res := &dto.GetFineResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Fine+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFineList(input *dto.GetFineListInputMS) ([]structs.Fine, int, error) {
	res := &dto.GetFineListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Fine, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFine(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Fine+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateFine(item *structs.Fine) (*structs.Fine, error) {
	res := &dto.GetFineResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Fine+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
