package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateNonFinancialGoal(nonFinancialGoal *structs.NonFinancialGoalItem) (*structs.NonFinancialGoalItem, error) {
	res := &dto.GetNonFinancialGoalResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.NonFinancialGoal, nonFinancialGoal, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateNonFinancialGoal(id int, nonFinancialGoal *structs.NonFinancialGoalItem) (*structs.NonFinancialGoalItem, error) {
	res := &dto.GetNonFinancialGoalResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.NonFinancialGoal+"/"+strconv.Itoa(id), nonFinancialGoal, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteNonFinancialGoal(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.NonFinancialGoal+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetNonFinancialGoalList(input *dto.GetNonFinancialGoalListInputMS) ([]structs.NonFinancialGoalItem, error) {
	res := &dto.GetNonFinancialGoalListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.NonFinancialGoal, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetNonFinancialGoal(id int) (*structs.NonFinancialGoalItem, error) {
	res := &dto.GetNonFinancialGoalResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.NonFinancialGoal+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
