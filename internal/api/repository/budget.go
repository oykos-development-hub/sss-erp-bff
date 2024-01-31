package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateBudget(budgetItem *structs.Budget) (*structs.Budget, error) {
	res := &dto.GetBudgetResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Budget, budgetItem, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateBudget(id int, budgetItem *structs.Budget) (*structs.Budget, error) {
	res := &dto.GetBudgetResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Budget+"/"+strconv.Itoa(id), budgetItem, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetBudgetList(input *dto.GetBudgetListInputMS) ([]structs.Budget, error) {
	res := &dto.GetBudgetListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Budget, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetBudget(id int) (*structs.Budget, error) {
	res := &dto.GetBudgetResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Budget+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteBudget(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Budget+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateFinancialBudget(financialBudget *structs.FinancialBudget) (*structs.FinancialBudget, error) {
	res := &dto.GetFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FinancialBudget, financialBudget, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFinancialBudgetByBudgetID(id int) (*structs.FinancialBudget, error) {
	res := &dto.GetFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Budget+"/"+strconv.Itoa(id)+"/financial", nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateLimitsForFinancialBudget(financialBudgetLimit *structs.FinancialBudgetLimit) (*structs.FinancialBudgetLimit, error) {
	res := &dto.GetFinancialBudgetLimitResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FinancialBudgetLimit, financialBudgetLimit, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
