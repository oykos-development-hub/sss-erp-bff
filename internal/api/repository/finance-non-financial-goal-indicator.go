package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateNonFinancialGoalIndicator(nonFinancialGoalIndicator *structs.NonFinancialGoalIndicatorItem) (*structs.NonFinancialGoalIndicatorItem, error) {
	res := &dto.GetNonFinancialGoalIndicatorResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.NonFinancialGoalIndicator, nonFinancialGoalIndicator, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateNonFinancialGoalIndicator(id int, nonFinancialGoalIndicator *structs.NonFinancialGoalIndicatorItem) (*structs.NonFinancialGoalIndicatorItem, error) {
	res := &dto.GetNonFinancialGoalIndicatorResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.NonFinancialGoalIndicator+"/"+strconv.Itoa(id), nonFinancialGoalIndicator, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteNonFinancialGoalIndicator(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.NonFinancialGoalIndicator+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetNonFinancialGoalIndicatorList(input *dto.GetNonFinancialGoalIndicatorListInputMS) ([]structs.NonFinancialGoalIndicatorItem, error) {
	res := &dto.GetNonFinancialGoalIndicatorListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.NonFinancialGoalIndicator, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetNonFinancialGoalIndicator(id int) (*structs.NonFinancialGoalIndicatorItem, error) {
	res := &dto.GetNonFinancialGoalIndicatorResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.NonFinancialGoalIndicator+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
