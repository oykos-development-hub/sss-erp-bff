package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateDropdownSettings(data *structs.SettingsDropdown) (*structs.SettingsDropdown, error) {
	res := &dto.GetDropdownTypeResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.SETTINGS, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteDropdownSettings(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Core.SETTINGS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateDropdownSettings(id int, data *structs.SettingsDropdown) (*structs.SettingsDropdown, error) {
	res := &dto.GetDropdownTypeResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.SETTINGS+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetDropdownSettings(input *dto.GetSettingsInput) (*dto.GetDropdownTypesResponseMS, error) {
	res := &dto.GetDropdownTypesResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.SETTINGS, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetDropdownSettingById(id int) (*structs.SettingsDropdown, error) {
	res := &dto.GetDropdownTypeResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.SETTINGS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetOfficeDropdownSettings(input *dto.GetOfficesOfOrganizationInput) (*dto.GetDropdownTypesResponseMS, error) {
	res := &dto.GetDropdownTypesResponseMS{}
	input.Entity = config.OfficeTypes
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.SETTINGS, &input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
