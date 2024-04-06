package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateFixedDeposit(item *structs.FixedDeposit) (*structs.FixedDeposit, error) {
	res := &dto.GetFixedDepositResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDeposit, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateFixedDeposit(item *structs.FixedDeposit) (*structs.FixedDeposit, error) {
	res := &dto.GetFixedDepositResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDeposit+"/"+strconv.Itoa(item.ID), item, res)
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

func (repo *MicroserviceRepository) DeleteFixedDeposit(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDeposit+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateFixedDepositItem(item *structs.FixedDepositItem) error {
	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDepositItem, item, nil)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) UpdateFixedDepositItem(item *structs.FixedDepositItem) error {
	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDepositItem+"/"+strconv.Itoa(item.ID), item, nil)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) DeleteFixedDepositItem(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDepositItem+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateFixedDepositDispatch(item *structs.FixedDepositDispatch) error {
	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDepositDispatch, item, nil)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) UpdateFixedDepositDispatch(item *structs.FixedDepositDispatch) error {
	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDepositDispatch+"/"+strconv.Itoa(item.ID), item, nil)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) DeleteFixedDepositDispatch(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDepositDispatch+"/"+strconv.Itoa(id), nil, nil)
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
