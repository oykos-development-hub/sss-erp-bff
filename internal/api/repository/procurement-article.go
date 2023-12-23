package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateProcurementArticle(article *structs.PublicProcurementArticle) (*structs.PublicProcurementArticle, error) {
	res := &dto.GetProcurementArticleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.Articles, article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementArticle(id int, article *structs.PublicProcurementArticle) (*structs.PublicProcurementArticle, error) {
	res := &dto.GetProcurementArticleResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.Articles+"/"+strconv.Itoa(id), article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteProcurementArticle(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Procurements.Articles+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetProcurementArticlesList(input *dto.GetProcurementArticleListInputMS) ([]*structs.PublicProcurementArticle, error) {
	res := &dto.GetProcurementArticleListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.Articles, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetProcurementArticle(id int) (*structs.PublicProcurementArticle, error) {
	res := &dto.GetProcurementArticleResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.Articles+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateProcurementOUArticle(article *structs.PublicProcurementOrganizationUnitArticle) (*structs.PublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.OrganizationUnitArticle, article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementOUArticle(id int, article *structs.PublicProcurementOrganizationUnitArticle) (*structs.PublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.OrganizationUnitArticle+"/"+strconv.Itoa(id), article, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcurementOUArticleList(input *dto.GetProcurementOrganizationUnitArticleListInputDTO) ([]*structs.PublicProcurementOrganizationUnitArticle, error) {
	res := &dto.GetOrganizationUnitArticleListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.OrganizationUnitArticle, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
