package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateProcurementPlan(resolution *structs.PublicProcurementPlan) (*structs.PublicProcurementPlan, error) {
	res := &dto.GetProcurementPlanResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.Plans, resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementPlan(id int, resolution *structs.PublicProcurementPlan) (*structs.PublicProcurementPlan, error) {
	res := &dto.GetProcurementPlanResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.Plans+"/"+strconv.Itoa(id), resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteProcurementPlan(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Procurements.Plans+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetProcurementPlan(id int) (*structs.PublicProcurementPlan, error) {
	res := &dto.GetProcurementPlanResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.Plans+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcurementPlanList(input *dto.GetProcurementPlansInput) ([]*structs.PublicProcurementPlan, error) {
	res := &dto.GetProcurementPlanListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.Plans, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
