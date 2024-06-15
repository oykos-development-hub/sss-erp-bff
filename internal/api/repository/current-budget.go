package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateCurrentBudget(ctx context.Context, currentBudget *structs.CurrentBudget) (*structs.CurrentBudget, error) {
	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	res := dto.GetCurrentBudgetResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.CurrentBudget, currentBudget, &res)
	if err != nil {
		return res.Data, err
	}

	return res.Data, nil
}
