package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateExternalReallocation(ctx context.Context, item *structs.ExternalReallocation) (*structs.ExternalReallocation, error) {
	res := &dto.GetExternalReallocationSingleResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.ExternalReallocation, item, res, header)
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

func (repo *MicroserviceRepository) DeleteExternalReallocation(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.ExternalReallocation+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) AcceptOUExternalReallocation(ctx context.Context, item *structs.ExternalReallocation) (*structs.ExternalReallocation, error) {
	res := &dto.GetExternalReallocationSingleResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.AcceptOUExternalReallocation, item, res, header)
	if err != nil {
		return nil, err
	}
	return res.Data, nil
}

func (repo *MicroserviceRepository) RejectOUExternalReallocation(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.RejectOUExternalReallocation+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) AcceptSSSExternalReallocation(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.AcceptSSSExternalReallocation+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) RejectSSSExternalReallocation(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.RejectSSSExternalReallocation+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}
