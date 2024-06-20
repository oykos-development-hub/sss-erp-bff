package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateTaxAuthorityCodebook(ctx context.Context, data *structs.TaxAuthorityCodebook) (*structs.TaxAuthorityCodebook, error) {
	res := &dto.GetTaxAuthorityCodebookResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.TaxAuthorityCodebook, data, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteTaxAuthorityCodebook(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.TaxAuthorityCodebook+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateTaxAuthorityCodebook(ctx context.Context, id int, data *structs.TaxAuthorityCodebook) (*structs.TaxAuthorityCodebook, error) {
	res := &dto.GetTaxAuthorityCodebookResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.TaxAuthorityCodebook+"/"+strconv.Itoa(id), data, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeactivateTaxAuthorityCodebook(ctx context.Context, id int, active bool) error {
	data := structs.TaxAuthorityCodebook{
		Active: active,
	}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.DeactivateTaxAuthorityCodebook+"/"+strconv.Itoa(id), data, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetTaxAuthorityCodebooks(input dto.TaxAuthorityCodebookFilter) (*dto.GetTaxAuthorityCodebooksResponseMS, error) {
	res := &dto.GetTaxAuthorityCodebooksResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.TaxAuthorityCodebook, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetTaxAuthorityCodebookByID(id int) (*structs.TaxAuthorityCodebook, error) {
	res := &dto.GetTaxAuthorityCodebookResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.TaxAuthorityCodebook+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}
