package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"errors"
	"strconv"
)

func (repo *MicroserviceRepository) GetLog(entity config.Module, id int) (*structs.Logs, error) {
	res := &dto.GetLogResponseMS{}

	switch entity {
	case config.ModuleCore:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Logs+strconv.Itoa(id), nil, res)
		if err != nil {
			return nil, err
		}

		return &res.Data, nil
	case config.ModuleHR:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Logs+strconv.Itoa(id), nil, res)
		if err != nil {
			return nil, err
		}

		return &res.Data, nil
	case config.ModuleProcurements:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.Logs+strconv.Itoa(id), nil, res)
		if err != nil {
			return nil, err
		}

		return &res.Data, nil
	case config.ModuleAccounting:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.Logs+strconv.Itoa(id), nil, res)
		if err != nil {
			return nil, err
		}

		return &res.Data, nil
	case config.ModuleInventory:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Logs+strconv.Itoa(id), nil, res)
		if err != nil {
			return nil, err
		}

		return &res.Data, nil
	case config.ModuleFinance:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Logs+strconv.Itoa(id), nil, res)
		if err != nil {
			return nil, err
		}

		return &res.Data, nil
	}

	return nil, errors.New("invalid module")

}

func (repo *MicroserviceRepository) GetLogs(filter dto.LogFilterDTO) ([]structs.Logs, int, error) {
	res := &dto.GetLogResponseListMS{}

	switch filter.Module {
	case config.ModuleCore:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Logs, filter, res)
		if err != nil {
			return nil, 0, err
		}

		return res.Data, res.Total, nil
	case config.ModuleHR:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Logs, filter, res)
		if err != nil {
			return nil, 0, err
		}

		return res.Data, res.Total, nil
	case config.ModuleProcurements:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.Procurements.Logs, filter, res)
		if err != nil {
			return nil, 0, err
		}

		return res.Data, res.Total, nil
	case config.ModuleAccounting:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.Accounting.Logs, filter, res)
		if err != nil {
			return nil, 0, err
		}

		return res.Data, res.Total, nil
	case config.ModuleInventory:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.Inventory.Logs, filter, res)
		if err != nil {
			return nil, 0, err
		}

		return res.Data, res.Total, nil
	case config.ModuleFinance:
		_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Logs, filter, res)
		if err != nil {
			return nil, 0, err
		}

		return res.Data, res.Total, nil
	}

	return nil, 0, errors.New("invalid module")
}
