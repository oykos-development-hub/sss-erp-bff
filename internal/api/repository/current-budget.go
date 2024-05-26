package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"context"
)

func (repo *MicroserviceRepository) CreateCurrentBudget(ctx context.Context, currentBudget *structs.CurrentBudget) (*structs.CurrentBudget, error) {
	res := dto.GetCurrentBudgetResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.CurrentBudget, currentBudget, &res)
	if err != nil {
		return res.Data, err
	}

	return res.Data, nil
}
