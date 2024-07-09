package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
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
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetSpendingReleaseOverview(ctx context.Context, input *dto.SpendingReleaseOverviewFilterDTO) ([]dto.SpendingReleaseOverviewItem, error) {
	res := dto.GetSpendingReleaseOverviewResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.SpendingReleaseOverview, input, &res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetSpendingReleaseList(ctx context.Context, input *dto.GetSpendingReleaseListInput) ([]structs.SpendingRelease, error) {
	res := dto.GetSpendingReleaseListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.SpendingReleaseList, input, &res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteSpendingRelease(ctx context.Context, input *dto.DeleteSpendingReleaseInput) error {
	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.SpendingReleaseDelete, input, nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) CreateSpendingReleaseRequest(ctx context.Context, spendingRelease structs.SpendingReleaseRequest) error {

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.SpendingReleaseRequest, spendingRelease, nil, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetSpendingReleaseRequests(filter dto.SpendingReleaseOverviewRequestFilter) ([]structs.SpendingReleaseRequest, error) {
	res := dto.GetSpendingReleaseRequestListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.SpendingReleaseRequest, filter, &res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) SpendingReleaseAcceptSSS(id int, fileID int) error {

	item := structs.SpendingReleaseRequest{
		SSSFileID: fileID,
	}

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.AcceptSpendingReleaseRequest+"/"+strconv.Itoa(id), item, nil, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}
