package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"fmt"
	"strconv"
)

func (repo *MicroserviceRepository) CreateSpendingRelease(ctx context.Context, spendingReleaseList []structs.SpendingReleaseInsert, budgetID, unitID int) ([]structs.SpendingRelease, error) {
	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	loggedInProfile, _ := ctx.Value(config.LoggedInProfileKey).(*structs.UserProfiles)

	spendingReleaseListToInsert := make([]structs.SpendingReleaseInsert, len(spendingReleaseList))
	for i, spendingRelease := range spendingReleaseList {
		spendingRelease.Username = loggedInProfile.GetFullName()
		spendingReleaseListToInsert[i] = spendingRelease
	}

	res := dto.GetSpendingReleaseListResponseMS{}
	_, err := makeAPIRequest("POST", fmt.Sprintf(repo.Config.Microservices.Finance.SpendingReleaseInsert, budgetID, unitID), spendingReleaseListToInsert, &res, header)
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

func (repo *MicroserviceRepository) DeleteSpendingRelease(ctx context.Context, input *dto.DeleteSpendingReleaseInput) error {
	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.SpendingReleaseDelete, input, nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}
