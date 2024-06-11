package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) DeleteOrganizationUnits(ctx context.Context, id int) error {
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.OrganizationUnits+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetOrganizationUnits(input *dto.GetOrganizationUnitsInput) (*dto.GetOrganizationUnitsResponseMS, error) {
	res := &dto.GetOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.OrganizationUnits, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetOrganizationUnitByID(id int) (*structs.OrganizationUnits, error) {
	res := &dto.GetOrganizationUnitResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.OrganizationUnits+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateOrganizationUnits(ctx context.Context, id int, data *structs.OrganizationUnits) (*dto.GetOrganizationUnitResponseMS, error) {
	res := &dto.GetOrganizationUnitResponseMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.OrganizationUnits+"/"+strconv.Itoa(id), data, res, header)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) CreateOrganizationUnits(ctx context.Context, data *structs.OrganizationUnits) (*dto.GetOrganizationUnitResponseMS, error) {
	res := &dto.GetOrganizationUnitResponseMS{}
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.OrganizationUnits, data, res, header)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetOrganizationUnitIDByUserProfile(id int) (*int, error) {
	employeesInOrganizationUnit, err := repo.GetEmployeesInOrganizationUnitsByProfileID(id)
	if err != nil {
		return nil, err
	}

	if employeesInOrganizationUnit == nil {
		return nil, nil
	}

	jobPositionInOrganizationUnit, err := repo.GetJobPositionsInOrganizationUnitsByID(employeesInOrganizationUnit.PositionInOrganizationUnitID)
	if err != nil {
		return nil, err
	}

	systematization, err := repo.GetSystematizationByID(jobPositionInOrganizationUnit.SystematizationID)
	if err != nil {
		return nil, err
	}

	return &systematization.OrganizationUnitID, nil
}
