package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateProcurementItem(item *structs.PublicProcurementItem) (*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Procurements.ITEMS, item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProcurementItem(id int, item *structs.PublicProcurementItem) (*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Procurements.ITEMS+"/"+strconv.Itoa(id), item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteProcurementItem(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Procurements.ITEMS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetProcurementItem(id int) (*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.ITEMS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetProcurementItemList(input *dto.GetProcurementItemListInputMS) ([]*structs.PublicProcurementItem, error) {
	res := &dto.GetProcurementItemListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.ITEMS, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
