package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateSalary(item *structs.Salary) (*structs.Salary, error) {
	res := &dto.GetSalaryResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Salary, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateSalary(item *structs.Salary) (*structs.Salary, error) {
	res := &dto.GetSalaryResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Salary+"/"+strconv.Itoa(item.ID), item, res)
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

func (repo *MicroserviceRepository) DeleteSalary(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Salary+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
