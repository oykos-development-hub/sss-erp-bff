package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) DeleteOrganizationUnits(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.ORGANIZATION_UNITS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetOrganizationUnits(input *dto.GetOrganizationUnitsInput) (*dto.GetOrganizationUnitsResponseMS, error) {
	res := &dto.GetOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.ORGANIZATION_UNITS, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetOrganizationUnitById(id int) (*structs.OrganizationUnits, error) {
	res := &dto.GetOrganizationUnitResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.ORGANIZATION_UNITS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateOrganizationUnits(id int, data *structs.OrganizationUnits) (*dto.GetOrganizationUnitResponseMS, error) {
	res := &dto.GetOrganizationUnitResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.ORGANIZATION_UNITS+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) CreateOrganizationUnits(data *structs.OrganizationUnits) (*dto.GetOrganizationUnitResponseMS, error) {
	res := &dto.GetOrganizationUnitResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.ORGANIZATION_UNITS, data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetOrganizationUnitIdByUserProfile(id int) (*int, error) {
	employeesInOrganizationUnit, err := repo.GetEmployeesInOrganizationUnitsByProfileId(id)
	if err != nil {
		return nil, err
	}

	if employeesInOrganizationUnit == nil {
		return nil, nil
	}

	jobPositionInOrganizationUnit, err := repo.GetJobPositionsInOrganizationUnitsById(employeesInOrganizationUnit.PositionInOrganizationUnitId)
	if err != nil {
		return nil, err
	}

	systematization, err := repo.GetSystematizationById(jobPositionInOrganizationUnit.SystematizationId)
	if err != nil {
		return nil, err
	}

	return &systematization.OrganizationUnitId, nil
}
