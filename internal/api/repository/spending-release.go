package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"context"
)

func (repo *MicroserviceRepository) CreateSpendingRelease(ctx context.Context, spendingRelease *structs.SpendingReleaseInsert) (*structs.SpendingRelease, error) {
	res := dto.GetSpendingReleaseResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.SpendingReleaseInsert, spendingRelease, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetSpendingReleaseOverview(ctx context.Context, input *dto.SpendingReleaseOverviewFilterDTO) ([]dto.SpendingReleaseOverviewItem, error) {
	res := dto.GetSpendingReleaseOverviewResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.SpendingReleaseOverview, input, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
