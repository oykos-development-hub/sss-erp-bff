package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateNonFinancialGoal(ctx context.Context, nonFinancialGoal *structs.NonFinancialGoalItem) (*structs.NonFinancialGoalItem, error) {
	res := &dto.GetNonFinancialGoalResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.NonFinancialGoal, nonFinancialGoal, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateNonFinancialGoal(ctx context.Context, id int, nonFinancialGoal *structs.NonFinancialGoalItem) (*structs.NonFinancialGoalItem, error) {
	res := &dto.GetNonFinancialGoalResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.NonFinancialGoal+"/"+strconv.Itoa(id), nonFinancialGoal, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteNonFinancialGoal(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.NonFinancialGoal+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetNonFinancialGoalList(input *dto.GetNonFinancialGoalListInputMS) ([]structs.NonFinancialGoalItem, error) {
	res := &dto.GetNonFinancialGoalListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.NonFinancialGoal, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetNonFinancialGoal(id int) (*structs.NonFinancialGoalItem, error) {
	res := &dto.GetNonFinancialGoalResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.NonFinancialGoal+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}
