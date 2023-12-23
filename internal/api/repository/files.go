package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetFileByID(id int) (*structs.File, error) {
	res := &dto.GetFileResponsePom{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Files.Files+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data.Data, nil
}

func (repo *MicroserviceRepository) DeleteFile(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.Files.Files+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
