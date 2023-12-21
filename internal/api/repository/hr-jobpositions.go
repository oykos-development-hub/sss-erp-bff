package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetJobPositionById(id int) (*structs.JobPositions, error) {
	res := &dto.GetJobPositionResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JOB_POSITIONS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetJobPositions(input *dto.GetJobPositionsInput) (*dto.GetJobPositionsResponseMS, error) {
	res := &dto.GetJobPositionsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JOB_POSITIONS, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) UpdateJobPositions(id int, data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error) {
	res := &dto.GetJobPositionResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JOB_POSITIONS+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) CreateJobPositions(data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error) {
	res := &dto.GetJobPositionResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JOB_POSITIONS, data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) DeleteJobPositions(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JOB_POSITIONS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateJobPositionsInOrganizationUnits(data *structs.JobPositionsInOrganizationUnits) (*dto.GetJobPositionInOrganizationUnitsResponseMS, error) {
	res := &dto.GetJobPositionInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JOB_POSITIONS_IN_ORGANIZATION_UNITS, data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) UpdateJobPositionsInOrganizationUnits(data *structs.JobPositionsInOrganizationUnits) (*dto.GetJobPositionInOrganizationUnitsResponseMS, error) {
	res := &dto.GetJobPositionInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JOB_POSITIONS_IN_ORGANIZATION_UNITS+"/"+strconv.Itoa(data.Id), data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetJobPositionsInOrganizationUnitsById(id int) (*structs.JobPositionsInOrganizationUnits, error) {
	res := &dto.GetJobPositionInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JOB_POSITIONS_IN_ORGANIZATION_UNITS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteJobPositionsInOrganizationUnits(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JOB_POSITIONS_IN_ORGANIZATION_UNITS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetJobPositionsInOrganizationUnits(input *dto.GetJobPositionInOrganizationUnitsInput) (*dto.GetJobPositionsInOrganizationUnitsResponseMS, error) {
	res := &dto.GetJobPositionsInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JOB_POSITIONS_IN_ORGANIZATION_UNITS, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) CreateEmployeesInOrganizationUnits(data *structs.EmployeesInOrganizationUnits) (*structs.EmployeesInOrganizationUnits, error) {
	res := &dto.GetEmployeesInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.EMPLOYEES_IN_ORGANIZATION_UNITS, data, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteEmployeeInOrganizationUnit(jobPositionInOrganizationUnitId int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.EMPLOYEES_IN_ORGANIZATION_UNITS+"/"+strconv.Itoa(jobPositionInOrganizationUnitId), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) DeleteEmployeeInOrganizationUnitByID(jobPositionInOrganizationUnitId int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.EMPLOYEES_IN_ORGANIZATION_UNITS_BY_ID+"/"+strconv.Itoa(jobPositionInOrganizationUnitId), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
