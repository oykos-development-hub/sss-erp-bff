package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateInternalReallocation(ctx context.Context, item *structs.InternalReallocation) (*structs.InternalReallocation, error) {
	res := &dto.GetInternalReallocationSingleResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.InternalReallocation, item, res, header)
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

func (repo *MicroserviceRepository) DeleteInternalReallocation(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.InternalReallocation+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}
