package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateSalary(ctx context.Context, item *structs.Salary) (*structs.Salary, error) {
	res := &dto.GetSalaryResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Salary, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateSalary(ctx context.Context, item *structs.Salary) (*structs.Salary, error) {
	res := &dto.GetSalaryResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Salary+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetSalaryByID(id int) (*structs.Salary, error) {
	res := &dto.GetSalaryResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Salary+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetSalaryList(filter dto.SalaryFilter) ([]structs.Salary, int, error) {
	res := &dto.GetSalaryListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Salary, filter, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteSalary(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Salary+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}
