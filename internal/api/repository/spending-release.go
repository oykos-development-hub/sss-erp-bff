package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"fmt"
)

func (repo *MicroserviceRepository) CreateSpendingRelease(ctx context.Context, spendingReleaseList []structs.SpendingReleaseInsert, budgetID, unitID int) ([]structs.SpendingRelease, error) {
	res := dto.GetSpendingReleaseListResponseMS{}
	_, err := makeAPIRequest("POST", fmt.Sprintf(repo.Config.Microservices.Finance.SpendingReleaseInsert, budgetID, unitID), spendingReleaseList, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetSpendingReleaseOverview(ctx context.Context, input *dto.SpendingReleaseOverviewFilterDTO) ([]dto.SpendingReleaseOverviewItem, error) {
	res := dto.GetSpendingReleaseOverviewResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.SpendingReleaseOverview, input, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetSpendingReleaseList(ctx context.Context, input *dto.GetSpendingReleaseListInput) ([]structs.SpendingRelease, error) {
	res := dto.GetSpendingReleaseListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.SpendingReleaseList, input, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteSpendingRelease(ctx context.Context, id int) error {
	_, err := makeAPIRequest("DELETE", fmt.Sprintf(repo.Config.Microservices.Finance.SpendingReleaseDelete, id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
