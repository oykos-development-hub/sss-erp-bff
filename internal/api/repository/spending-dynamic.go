package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

func (repo *MicroserviceRepository) GetSpendingDynamic(BudgetID, unitID int, input *dto.GetSpendingDynamicHistoryInput) ([]dto.SpendingDynamicDTO, error) {
	res := dto.GetSpendingDynamicListResponseMS{}
	_, err := makeAPIRequest("GET", fmt.Sprintf(repo.Config.Microservices.Finance.SpendingDynamicGet, BudgetID, unitID), input, &res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetSpendingDynamicHistory(BudgetID, unitID int) ([]dto.SpendingDynamicHistoryDTO, error) {
	res := dto.GetSpendingDynamicHistoryResponseMS{}
	_, err := makeAPIRequest("GET", fmt.Sprintf(repo.Config.Microservices.Finance.SpendingDynamicGetHistory, BudgetID, unitID), nil, &res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetSpendingDynamicActual(BudgetID, unitID, accountID int) (decimal.NullDecimal, error) {
	res := dto.GetSpendingDynamicActualResponseMS{}
	_, err := makeAPIRequest("GET", fmt.Sprintf(repo.Config.Microservices.Finance.SpendingDynamicActual, BudgetID, unitID, accountID), nil, &res)
	if err != nil {
		return res.Data, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateSpendingDynamic(ctx context.Context, budgetID, unitID int, spendingDynamicList []structs.SpendingDynamicInsert) ([]dto.SpendingDynamicDTO, error) {
	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	loggedInProfile, _ := ctx.Value(config.LoggedInProfileKey).(*structs.UserProfiles)

	spendingDynamicListToInsert := make([]structs.SpendingDynamicInsert, len(spendingDynamicList))
	for i, spendingDynamic := range spendingDynamicList {
		spendingDynamic.Username = loggedInProfile.GetFullName()
		spendingDynamicListToInsert[i] = spendingDynamic
	}

	res := dto.GetSpendingDynamicListResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Budget+"/"+strconv.Itoa(budgetID)+"/units/"+strconv.Itoa(unitID)+"/spending-dynamics", spendingDynamicListToInsert, &res, header)
	if err != nil {
		return res.Data, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}
