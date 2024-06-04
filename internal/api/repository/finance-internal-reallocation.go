package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateInternalReallocation(item *structs.InternalReallocation) (*structs.InternalReallocation, error) {
	res := &dto.GetInternalReallocationSingleResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.InternalReallocation, item, res)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (repo *MicroserviceRepository) GetInternalReallocationByID(id int) (*structs.InternalReallocation, error) {
	res := &dto.GetInternalReallocationSingleResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.InternalReallocation+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetInternalReallocationList(filter dto.InternalReallocationFilter) ([]structs.InternalReallocation, int, error) {
	res := &dto.GetInternalReallocationResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.InternalReallocation, filter, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteInternalReallocation(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.InternalReallocation+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
