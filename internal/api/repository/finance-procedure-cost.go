package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateProcedureCost(item *structs.ProcedureCost) (*structs.ProcedureCost, error) {
	res := &dto.GetProcedureCostResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.ProcedureCost, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcedureCost(id int) (*structs.ProcedureCost, error) {
	res := &dto.GetProcedureCostResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.ProcedureCost+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcedureCostList(input *dto.GetProcedureCostListInputMS) ([]structs.ProcedureCost, int, error) {
	res := &dto.GetProcedureCostListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.ProcedureCost, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteProcedureCost(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.ProcedureCost+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateProcedureCost(item *structs.ProcedureCost) (*structs.ProcedureCost, error) {
	res := &dto.GetProcedureCostResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.ProcedureCost+"/"+strconv.Itoa(item.ID), item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
