package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateProcurementContract(ctx context.Context, resolution *structs.PublicProcurementContract) (*structs.PublicProcurementContract, error) {
	res := &dto.GetProcurementContractResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.Contracts, resolution, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementContract(ctx context.Context, id int, resolution *structs.PublicProcurementContract) (*structs.PublicProcurementContract, error) {
	res := &dto.GetProcurementContractResponseMS{}

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.Contracts+"/"+strconv.Itoa(id), resolution, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteProcurementContract(ctx context.Context, id int) error {

	header := make(map[string]string)

	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Procurements.Contracts+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) GetProcurementContract(id int) (*structs.PublicProcurementContract, error) {
	res := &dto.GetProcurementContractResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.Contracts+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcurementContractsList(input *dto.GetProcurementContractsInput) (*dto.GetProcurementContractListResponseMS, error) {
	res := &dto.GetProcurementContractListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.Contracts, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetOrganizationUnitArticlesList(input dto.GetProcurementOrganizationUnitArticleListInputDTO) ([]dto.GetPublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleListResponse{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.OrganizationUnitArticle, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetOrganizationUnitArticleByID(id int) (*dto.GetPublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleResponse{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.OrganizationUnitArticle+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateProcurementContractArticle(article *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error) {
	res := &dto.GetProcurementContractArticleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.ContractArticle, article, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementContractArticle(id int, article *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error) {
	res := &dto.GetProcurementContractArticleResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.ContractArticle+"/"+strconv.Itoa(id), article, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcurementContractArticlesList(input *dto.GetProcurementContractArticlesInput) (*dto.GetProcurementContractArticlesListResponseMS, error) {
	res := &dto.GetProcurementContractArticlesListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.ContractArticle, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res, nil
}

func (repo *MicroserviceRepository) CreateProcurementContractArticleOverage(articleOverage *structs.PublicProcurementContractArticleOverage) (*structs.PublicProcurementContractArticleOverage, error) {
	res := &dto.GetProcurementContractArticleOverageResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.ContractArticleOverage, articleOverage, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementContractArticleOverage(id int, articleOverage *structs.PublicProcurementContractArticleOverage) (*structs.PublicProcurementContractArticleOverage, error) {
	res := &dto.GetProcurementContractArticleOverageResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.ContractArticleOverage+"/"+strconv.Itoa(id), articleOverage, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcurementContractArticleOverageList(input *dto.GetProcurementContractArticleOverageInput) ([]*structs.PublicProcurementContractArticleOverage, error) {
	res := &dto.GetProcurementContractArticleOverageListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.ContractArticleOverage, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteProcurementContractArticleOverage(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Procurements.ContractArticleOverage+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}
