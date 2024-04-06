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

func (repo *MicroserviceRepository) CreateFixedDepositWill(item *structs.FixedDepositWill) (*structs.FixedDepositWill, error) {
	res := &dto.GetFixedDepositWillResponseMS{}
	item.Status = "Depozit"
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDepositWill, item, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateFixedDepositWill(item *structs.FixedDepositWill) (*structs.FixedDepositWill, error) {
	res := &dto.GetFixedDepositWillResponseMS{}

	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDepositWill+"/"+strconv.Itoa(item.ID), item, res)
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

func (repo *MicroserviceRepository) DeleteFixedDepositWill(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDepositWill+"/"+strconv.Itoa(id), nil, nil)
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

func (repo *MicroserviceRepository) CreateFixedDepositWillDispatch(item *structs.FixedDepositWillDispatch) error {
	//res := &dto.GetFixedDepositItemResponseMS{}

	will, err := repo.GetFixedDepositWillByID(item.WillID)

	if err != nil {
		return err
	}

	if will.Status == "U toku" {
		will.Status = "Depozit"
	} else {
		will.Status = "U toku"
	}

	_, err = repo.UpdateFixedDepositWill(will)

	if err != nil {
		return nil
	}

	_, err = makeAPIRequest("POST", repo.Config.Microservices.Finance.FixedDepositWillDispatch, item, nil)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) UpdateFixedDepositWillDispatch(item *structs.FixedDepositWillDispatch) error {
	//res := &dto.GetFixedDepositItemResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.FixedDepositWillDispatch+"/"+strconv.Itoa(item.ID), item, nil)
	if err != nil {
		return err
	}
	return nil
}

func (repo *MicroserviceRepository) DeleteFixedDepositWillDispatch(id int) error {
	item, err := repo.GetFixedDepositWillDispatchByID(id)

	if err != nil {
		return err
	}

	will, err := repo.GetFixedDepositWillByID(item.WillID)

	if err != nil {
		return err
	}

	if will.Status == "U toku" {
		will.Status = "Depozit"
	} else {
		will.Status = "U toku"
	}

	_, err = repo.UpdateFixedDepositWill(will)

	if err != nil {
		return nil
	}

	_, err = makeAPIRequest("DELETE", repo.Config.Microservices.Finance.FixedDepositWillDispatch+"/"+strconv.Itoa(id), nil, nil)
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
