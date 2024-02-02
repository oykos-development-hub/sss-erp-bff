package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateNonFinancialBudget(nonFinancialBudget *structs.NonFinancialBudgetItem) (*structs.NonFinancialBudgetItem, error) {
	res := &dto.GetNonFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.NonFinancialBudget, nonFinancialBudget, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateNonFinancialBudget(id int, nonFinancialBudget *structs.NonFinancialBudgetItem) (*structs.NonFinancialBudgetItem, error) {
	res := &dto.GetNonFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.NonFinancialBudget+"/"+strconv.Itoa(id), nonFinancialBudget, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteNonFinancialBudget(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.NonFinancialBudget+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetNonFinancialBudgetList(input *dto.GetNonFinancialBudgetListInputMS) ([]structs.NonFinancialBudgetItem, error) {
	res := &dto.GetNonFinancialBudgetListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.NonFinancialBudget, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetNonFinancialBudget(id int) (*structs.NonFinancialBudgetItem, error) {
	res := &dto.GetNonFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.NonFinancialBudget+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
