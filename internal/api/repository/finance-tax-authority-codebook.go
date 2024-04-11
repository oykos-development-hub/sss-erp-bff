package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateTaxAuthorityCodebook(data *structs.TaxAuthorityCodebook) (*structs.TaxAuthorityCodebook, error) {
	res := &dto.GetTaxAuthorityCodebookResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.TaxAuthorityCodebook, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteTaxAuthorityCodebook(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.TaxAuthorityCodebook+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateTaxAuthorityCodebook(id int, data *structs.TaxAuthorityCodebook) (*structs.TaxAuthorityCodebook, error) {
	res := &dto.GetTaxAuthorityCodebookResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.TaxAuthorityCodebook+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeactivateTaxAuthorityCodebook(id int, active bool) error {
	data := structs.TaxAuthorityCodebook{
		Active: active,
	}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.DeactivateTaxAuthorityCodebook+"/"+strconv.Itoa(id), data, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetTaxAuthorityCodebooks(input dto.TaxAuthorityCodebookFilter) (*dto.GetTaxAuthorityCodebooksResponseMS, error) {
	res := &dto.GetTaxAuthorityCodebooksResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.TaxAuthorityCodebook, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetTaxAuthorityCodebookByID(id int) (*structs.TaxAuthorityCodebook, error) {
	res := &dto.GetTaxAuthorityCodebookResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.TaxAuthorityCodebook+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
