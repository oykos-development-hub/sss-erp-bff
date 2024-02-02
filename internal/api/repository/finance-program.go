package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) CreateProgram(program *structs.ProgramItem) (*structs.ProgramItem, error) {
	res := &dto.GetFinanceProgramResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Finance.Program, program, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateProgram(id int, program *structs.ProgramItem) (*structs.ProgramItem, error) {
	res := &dto.GetFinanceProgramResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Finance.Program+"/"+strconv.Itoa(id), program, res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteProgram(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Finance.Program+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetProgramList(input *dto.GetFinanceProgramListInputMS) ([]structs.ProgramItem, error) {
	res := &dto.GetFinanceProgramListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Program, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetProgram(id int) (*structs.ProgramItem, error) {
	res := &dto.GetFinanceProgramResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Finance.Program+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
