package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
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

func (repo *MicroserviceRepository) UpdateJobPositions(ctx context.Context, id int, data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error) {
	res := &dto.GetJobPositionResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JobPositions+"/"+strconv.Itoa(id), data, res, header)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) CreateJobPositions(ctx context.Context, data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error) {
	res := &dto.GetJobPositionResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JobPositions, data, res, header)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) DeleteJobPositions(ctx context.Context, id int) error {

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JobPositions+"/"+strconv.Itoa(id), nil, nil, header)
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
