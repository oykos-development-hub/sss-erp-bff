package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateProcurementContract(resolution *structs.PublicProcurementContract) (*structs.PublicProcurementContract, error) {
	res := &dto.GetProcurementContractResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.Contracts, resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementContract(id int, resolution *structs.PublicProcurementContract) (*structs.PublicProcurementContract, error) {
	res := &dto.GetProcurementContractResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.Contracts+"/"+strconv.Itoa(id), resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteProcurementContract(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Procurements.Contracts+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetProcurementContract(id int) (*structs.PublicProcurementContract, error) {
	res := &dto.GetProcurementContractResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.Contracts+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcurementContractsList(input *dto.GetProcurementContractsInput) (*dto.GetProcurementContractListResponseMS, error) {
	res := &dto.GetProcurementContractListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.Contracts, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetOrganizationUnitArticlesList(input dto.GetProcurementOrganizationUnitArticleListInputDTO) ([]dto.GetPublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleListResponse{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.OrganizationUnitArticle, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetOrganizationUnitArticleByID(id int) (*dto.GetPublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleResponse{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.OrganizationUnitArticle+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateProcurementContractArticle(article *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error) {
	res := &dto.GetProcurementContractArticleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.ContractArticle, article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementContractArticle(id int, article *structs.PublicProcurementContractArticle) (*structs.PublicProcurementContractArticle, error) {
	res := &dto.GetProcurementContractArticleResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.ContractArticle+"/"+strconv.Itoa(id), article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcurementContractArticlesList(input *dto.GetProcurementContractArticlesInput) (*dto.GetProcurementContractArticlesListResponseMS, error) {
	res := &dto.GetProcurementContractArticlesListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.ContractArticle, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) CreateProcurementContractArticleOverage(articleOverage *structs.PublicProcurementContractArticleOverage) (*structs.PublicProcurementContractArticleOverage, error) {
	res := &dto.GetProcurementContractArticleOverageResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.ContractArticleOverage, articleOverage, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementContractArticleOverage(id int, articleOverage *structs.PublicProcurementContractArticleOverage) (*structs.PublicProcurementContractArticleOverage, error) {
	res := &dto.GetProcurementContractArticleOverageResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.ContractArticleOverage+"/"+strconv.Itoa(id), articleOverage, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcurementContractArticleOverageList(input *dto.GetProcurementContractArticleOverageInput) ([]*structs.PublicProcurementContractArticleOverage, error) {
	res := &dto.GetProcurementContractArticleOverageListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.ContractArticleOverage, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteProcurementContractArticleOverage(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Procurements.ContractArticleOverage+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
