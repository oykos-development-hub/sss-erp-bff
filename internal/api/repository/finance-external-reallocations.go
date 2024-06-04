package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateExternalReallocation(item *structs.ExternalReallocation) (*structs.ExternalReallocation, error) {
	res := &dto.GetExternalReallocationSingleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.ExternalReallocation, item, res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (repo *MicroserviceRepository) GetExternalReallocationByID(id int) (*structs.ExternalReallocation, error) {
	res := &dto.GetExternalReallocationSingleResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.ExternalReallocation+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetExternalReallocationList(filter dto.ExternalReallocationFilter) ([]structs.ExternalReallocation, int, error) {
	res := &dto.GetExternalReallocationResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.ExternalReallocation, filter, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteExternalReallocation(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.ExternalReallocation+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) AcceptOUExternalReallocation(item *structs.ExternalReallocation) (*structs.ExternalReallocation, error) {
	res := &dto.GetExternalReallocationSingleResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.AcceptOUExternalReallocation, item, res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (repo *MicroserviceRepository) RejectOUExternalReallocation(id int) error {
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.RejectOUExternalReallocation+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
