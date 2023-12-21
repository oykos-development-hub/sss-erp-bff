package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetSystematizationById(id int) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.SYSTEMATIZATIONS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetSystematizations(input *dto.GetSystematizationsInput) (*dto.GetSystematizationsResponseMS, error) {
	res := &dto.GetSystematizationsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.SYSTEMATIZATIONS, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) UpdateSystematization(id int, data *structs.Systematization) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.SYSTEMATIZATIONS+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateSystematization(data *structs.Systematization) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.SYSTEMATIZATIONS, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteSystematization(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.SYSTEMATIZATIONS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
