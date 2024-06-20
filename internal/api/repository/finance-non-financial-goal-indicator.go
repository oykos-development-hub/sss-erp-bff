package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateNonFinancialGoalIndicator(ctx context.Context, nonFinancialGoalIndicator *structs.NonFinancialGoalIndicatorItem) (*structs.NonFinancialGoalIndicatorItem, error) {
	res := &dto.GetNonFinancialGoalIndicatorResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.NonFinancialGoalIndicator, nonFinancialGoalIndicator, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateNonFinancialGoalIndicator(ctx context.Context, id int, nonFinancialGoalIndicator *structs.NonFinancialGoalIndicatorItem) (*structs.NonFinancialGoalIndicatorItem, error) {
	res := &dto.GetNonFinancialGoalIndicatorResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.NonFinancialGoalIndicator+"/"+strconv.Itoa(id), nonFinancialGoalIndicator, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteNonFinancialGoalIndicator(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.NonFinancialGoalIndicator+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetNonFinancialGoalIndicatorList(input *dto.GetNonFinancialGoalIndicatorListInputMS) ([]structs.NonFinancialGoalIndicatorItem, error) {
	res := &dto.GetNonFinancialGoalIndicatorListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.NonFinancialGoalIndicator, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetNonFinancialGoalIndicator(id int) (*structs.NonFinancialGoalIndicatorItem, error) {
	res := &dto.GetNonFinancialGoalIndicatorResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.NonFinancialGoalIndicator+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}
