package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetJobPositionByID(id int) (*structs.JobPositions, error) {
	res := &dto.GetJobPositionResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JobPositions+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetJobPositions(input *dto.GetJobPositionsInput) (*dto.GetJobPositionsResponseMS, error) {
	res := &dto.GetJobPositionsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JobPositions, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) UpdateJobPositions(id int, data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error) {
	res := &dto.GetJobPositionResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JobPositions+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) CreateJobPositions(data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error) {
	res := &dto.GetJobPositionResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JobPositions, data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) DeleteJobPositions(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JobPositions+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateJobPositionsInOrganizationUnits(data *structs.JobPositionsInOrganizationUnits) (*dto.GetJobPositionInOrganizationUnitsResponseMS, error) {
	res := &dto.GetJobPositionInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JobPositionInOrganizationUnits, data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) UpdateJobPositionsInOrganizationUnits(data *structs.JobPositionsInOrganizationUnits) (*dto.GetJobPositionInOrganizationUnitsResponseMS, error) {
	res := &dto.GetJobPositionInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JobPositionInOrganizationUnits+"/"+strconv.Itoa(data.ID), data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetJobPositionsInOrganizationUnitsByID(id int) (*structs.JobPositionsInOrganizationUnits, error) {
	res := &dto.GetJobPositionInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JobPositionInOrganizationUnits+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteJobPositionsInOrganizationUnits(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JobPositionInOrganizationUnits+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetJobPositionsInOrganizationUnits(input *dto.GetJobPositionInOrganizationUnitsInput) (*dto.GetJobPositionsInOrganizationUnitsResponseMS, error) {
	res := &dto.GetJobPositionsInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JobPositionInOrganizationUnits, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) CreateEmployeesInOrganizationUnits(data *structs.EmployeesInOrganizationUnits) (*structs.EmployeesInOrganizationUnits, error) {
	res := &dto.GetEmployeesInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.EmployeesInOrganizationUnits, data, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteEmployeeInOrganizationUnit(jobPositionInOrganizationUnitID int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.EmployeesInOrganizationUnits+"/"+strconv.Itoa(jobPositionInOrganizationUnitID), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) DeleteEmployeeInOrganizationUnitByID(jobPositionInOrganizationUnitID int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.EmployeesInOrganizationUnitByID+"/"+strconv.Itoa(jobPositionInOrganizationUnitID), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
