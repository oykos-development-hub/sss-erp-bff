package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateTemplate(ctx context.Context, item *structs.Template) error {
	res := &dto.GetTemplateResponseMS{}
	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.Templates, item, res, header)

	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	isParent := true

	organizationUnits, err := repo.GetOrganizationUnits(&dto.GetOrganizationUnitsInput{
		IsParent: &isParent,
	})

	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	templateID := res.Data.ID

	for _, orgUnit := range organizationUnits.Data {
		newItem := structs.Template{
			OrganizationUnitID: orgUnit.ID,
			TemplateID:         templateID,
			FileID:             item.FileID,
		}

		_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.TemplateItems, newItem, res, header)

		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateTemplate(ctx context.Context, item *structs.Template) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.Templates+"/"+strconv.Itoa(item.ID), item, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}
	return nil
}

func (repo *MicroserviceRepository) UpdateTemplateItem(ctx context.Context, item *structs.Template) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.TemplateItems+"/"+strconv.Itoa(item.ID), item, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}
	return nil
}

func (repo *MicroserviceRepository) UpdateCustomerSupport(ctx context.Context, item *structs.CustomerSupport) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.CustomerSupport+"/"+strconv.Itoa(item.ID), item, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}
	return nil
}

func (repo *MicroserviceRepository) GetCustomerSupport(id int) (*structs.CustomerSupport, error) {
	res := &dto.GetCustomerResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.CustomerSupport+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetTemplateList(input dto.TemplateFilter) ([]structs.Template, int, error) {
	res := &dto.GetTemplateResponseListMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.TemplateItems, input, res)
	if err != nil {
		return nil, 0, errors.Wrap(err, "make api request")
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) GetTemplateByID(id int) (*structs.Template, error) {
	res := &dto.GetTemplateResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Templates+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteTemplate(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Core.Templates+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetListOfParameters() ([]structs.ListOfParameters, error) {
	res := &dto.GetListOfParametersResponseListMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.ListOfParameters, nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}
