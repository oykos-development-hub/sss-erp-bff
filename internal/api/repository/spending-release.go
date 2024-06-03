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
