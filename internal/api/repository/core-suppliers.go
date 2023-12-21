package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateSupplier(supplier *structs.Suppliers) (*structs.Suppliers, error) {
	res := &dto.GetSupplierResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.SUPPLIERS, supplier, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateSupplier(id int, supplier *structs.Suppliers) (*structs.Suppliers, error) {
	res := &dto.GetSupplierResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.SUPPLIERS+"/"+strconv.Itoa(id), supplier, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteSupplier(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Core.SUPPLIERS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetSupplier(id int) (*structs.Suppliers, error) {
	res := &dto.GetSupplierResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.SUPPLIERS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetSupplierList(input *dto.GetSupplierInputMS) (*dto.GetSupplierListResponseMS, error) {
	res := &dto.GetSupplierListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.SUPPLIERS, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
