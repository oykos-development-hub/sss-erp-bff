package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"fmt"
	"strconv"
)

func (repo *MicroserviceRepository) CreateActivity(activity *structs.ActivitiesItem) (*structs.ActivitiesItem, error) {
	res := &dto.GetFinanceActivityResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Activity, activity, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateActivity(id int, activity *structs.ActivitiesItem) (*structs.ActivitiesItem, error) {
	res := &dto.GetFinanceActivityResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Activity+"/"+strconv.Itoa(id), activity, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteActivity(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Activity+"/"+strconv.Itoa(id), nil, nil)
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
