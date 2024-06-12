package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"fmt"
	"strconv"
)

func (repo *MicroserviceRepository) CreateActivity(ctx context.Context, activity *structs.ActivitiesItem) (*structs.ActivitiesItem, error) {
	res := &dto.GetFinanceActivityResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Activity, activity, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateActivity(ctx context.Context, id int, activity *structs.ActivitiesItem) (*structs.ActivitiesItem, error) {
	res := &dto.GetFinanceActivityResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Activity+"/"+strconv.Itoa(id), activity, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteActivity(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Activity+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetActivityList(input *dto.GetFinanceActivityListInputMS) ([]structs.ActivitiesItem, error) {
	res := &dto.GetFinanceActivityListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Activity, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetActivityByUnit(unitID int) (*structs.ActivitiesItem, error) {
	res := &dto.GetFinanceActivityListResponseMS{}
	input := dto.GetFinanceActivityListInputMS{OrganizationUnitID: &unitID}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Activity, input, res)
	if err != nil {
		return nil, err
	}

	if len(res.Data) == 0 {
		return nil, fmt.Errorf("cannot find activity for unit id: %d", unitID)
	}

	return &res.Data[0], nil
}

func (repo *MicroserviceRepository) GetActivity(id int) (*structs.ActivitiesItem, error) {
	res := &dto.GetFinanceActivityResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Activity+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
