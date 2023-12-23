package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetProcurementOULimitList(input *dto.GetProcurementOULimitListInputMS) ([]*structs.PublicProcurementLimit, error) {
	res := &dto.GetProcurementOULimitListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.OULimits, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateProcurementOULimit(limit *structs.PublicProcurementLimit) (*structs.PublicProcurementLimit, error) {
	res := &dto.GetProcurementOULimitResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.OULimits, limit, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementOULimit(id int, limit *structs.PublicProcurementLimit) (*structs.PublicProcurementLimit, error) {
	res := &dto.GetProcurementOULimitResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.OULimits+"/"+strconv.Itoa(id), limit, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
