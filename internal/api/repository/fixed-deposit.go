package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"context"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFixedDeposit(ctx context.Context, item *structs.FixedDeposit) (*structs.FixedDeposit, error) {
	res := &dto.GetFixedDepositResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDeposit, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateFixedDeposit(ctx context.Context, item *structs.FixedDeposit) (*structs.FixedDeposit, error) {
	res := &dto.GetFixedDepositResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDeposit+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFixedDepositByID(id int) (*structs.FixedDeposit, error) {
	res := &dto.GetFixedDepositResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FixedDeposit+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFixedDepositList(filter dto.FixedDepositFilter) ([]structs.FixedDeposit, int, error) {
	res := &dto.GetFixedDepositListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FixedDeposit, filter, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFixedDeposit(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDeposit+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateFixedDepositWill(ctx context.Context, item *structs.FixedDepositWill) (*structs.FixedDepositWill, error) {
	res := &dto.GetFixedDepositWillResponseMS{}
	item.Status = "Depozit"

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDepositWill, item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateFixedDepositWill(ctx context.Context, item *structs.FixedDepositWill) (*structs.FixedDepositWill, error) {
	res := &dto.GetFixedDepositWillResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDepositWill+"/"+strconv.Itoa(item.ID), item, res, header)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFixedDepositWillByID(id int) (*structs.FixedDepositWill, error) {
	res := &dto.GetFixedDepositWillResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FixedDepositWill+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetFixedDepositWillList(filter dto.FixedDepositWillFilter) ([]structs.FixedDepositWill, int, error) {
	res := &dto.GetFixedDepositWillListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FixedDepositWill, filter, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteFixedDepositWill(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDepositWill+"/"+strconv.Itoa(id), nil, nil, header)

	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateFixedDepositItem(ctx context.Context, item *structs.FixedDepositItem) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDepositItem, item, nil, header)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) UpdateFixedDepositItem(ctx context.Context, item *structs.FixedDepositItem) error {
	//res := &dto.GetFixedDepositItemResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDepositItem+"/"+strconv.Itoa(item.ID), item, nil, header)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) DeleteFixedDepositItem(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDepositItem+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateFixedDepositDispatch(ctx context.Context, item *structs.FixedDepositDispatch) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDepositDispatch, item, nil, header)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) UpdateFixedDepositDispatch(ctx context.Context, item *structs.FixedDepositDispatch) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDepositDispatch+"/"+strconv.Itoa(item.ID), item, nil, header)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) DeleteFixedDepositDispatch(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDepositDispatch+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateFixedDepositJudge(item *structs.FixedDepositJudge) error {

	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDepositJudge, item, nil)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) UpdateFixedDepositJudge(item *structs.FixedDepositJudge) error {
	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDepositJudge+"/"+strconv.Itoa(item.ID), item, nil)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) DeleteFixedDepositJudge(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDepositJudge+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateFixedDepositWillDispatch(ctx context.Context, item *structs.FixedDepositWillDispatch) error {
	//res := &dto.GetFixedDepositItemResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	will, err := repo.GetFixedDepositWillByID(item.WillID)

	if err != nil {
		return err
	}

	if will.Status == "U radu" {
		will.Status = "Depozit"
	} else {
		will.Status = "U radu"
	}

	_, err = repo.UpdateFixedDepositWill(ctx, will)

	if err != nil {
		return nil
	}

	_, err = makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDepositWillDispatch, item, nil, header)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) UpdateFixedDepositWillDispatch(ctx context.Context, item *structs.FixedDepositWillDispatch) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDepositWillDispatch+"/"+strconv.Itoa(item.ID), item, nil, header)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) DeleteFixedDepositWillDispatch(ctx context.Context, id int) error {

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	item, err := repo.GetFixedDepositWillDispatchByID(id)

	if err != nil {
		return err
	}

	will, err := repo.GetFixedDepositWillByID(item.WillID)

	if err != nil {
		return err
	}

	if will.Status == "U radu" {
		will.Status = "Depozit"
	} else {
		will.Status = "U radu"
	}

	_, err = repo.UpdateFixedDepositWill(ctx, will)

	if err != nil {
		return nil
	}

	_, err = makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDepositWillDispatch+"/"+strconv.Itoa(id), nil, nil, header)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetFixedDepositWillDispatchByID(id int) (*structs.FixedDepositWillDispatch, error) {

	res := &dto.GetFixedDepositWillDispatchResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.FixedDepositWillDispatch+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
