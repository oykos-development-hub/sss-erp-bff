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

func (repo *MicroserviceRepository) UpdateBudget(item *structs.Budget) (*structs.Budget, error) {
	res := &dto.GetBudgetResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Budget+"/"+strconv.Itoa(item.ID), item, res)
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

func (repo *MicroserviceRepository) UpdateFinancialBudget(financialBudget *structs.FinancialBudget) (*structs.FinancialBudget, error) {
	res := &dto.GetFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FinancialBudget+"/"+strconv.Itoa(financialBudget.ID), financialBudget, res)
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

func (repo *MicroserviceRepository) CreateBudgetLimit(budgetLimit *structs.FinancialBudgetLimit) (*structs.FinancialBudgetLimit, error) {
	res := &dto.GetFinancialBudgetLimitResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FinancialBudgetLimit, budgetLimit, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateBudgetLimit(budgetLimit *structs.FinancialBudgetLimit) (*structs.FinancialBudgetLimit, error) {
	res := &dto.GetFinancialBudgetLimitResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FinancialBudgetLimit+"/"+strconv.Itoa(budgetLimit.ID), budgetLimit, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteBudgetLimit(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FinancialBudgetLimit+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetBudgetLimits(budgetID int) ([]structs.FinancialBudgetLimit, error) {
	input := dto.GetFinancialBudgetListInputMS{
		BudgetID: budgetID,
	}
	res := &dto.GetFinancialBudgetLimitListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FinancialBudgetLimit, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetFilledFinancialBudgetList(requestID int) ([]structs.FilledFinanceBudget, error) {
	input := &dto.FilledFinancialBudgetInputMS{
		BudgetRequestID: requestID,
	}
	res := &dto.GetFilledFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FilledFinancialBudget, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetFinancialBudgetByID(id int) (*structs.FinancialBudget, error) {
	res := &dto.GetFinancialBudgetResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FinancialBudget+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) FillFinancialBudget(data *structs.FilledFinanceBudget) (*structs.FilledFinanceBudget, error) {
	res := &dto.GetFilledFinancialBudgetResponseItemMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FilledFinancialBudget, data, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateFilledFinancialBudget(id int, data *structs.FilledFinanceBudget) (*structs.FilledFinanceBudget, error) {
	res := &dto.GetFilledFinancialBudgetResponseItemMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FilledFinancialBudget+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteFilledFinancialBudgetData(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FilledFinancialBudget+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
