package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) UpdateJudgeNorm(id int, norm *structs.JudgeNorms) (*structs.JudgeNorms, error) {
	res := &dto.GetJudgeNormResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JudgeNorm+"/"+strconv.Itoa(id), norm, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateJudgeNorm(norm *structs.JudgeNorms) (*structs.JudgeNorms, error) {
	res := &dto.GetJudgeNormResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JudgeNorm, norm, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteJudgeNorm(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JudgeNorm+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateJudgeResolutionItems(id int, item *structs.JudgeResolutionItems) (*structs.JudgeResolutionItems, error) {
	res := &dto.GetJudgeResolutionItemResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JudgeResolutionItems+"/"+strconv.Itoa(id), item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateJudgeResolutionItems(item *structs.JudgeResolutionItems) (*structs.JudgeResolutionItems, error) {
	res := &dto.GetJudgeResolutionItemResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JudgeResolutionItems, item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetJudgeResolutionItemsList(input *dto.GetJudgeResolutionItemListInputMS) ([]*structs.JudgeResolutionItems, error) {
	res := &dto.GetJudgeResolutionItemListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JudgeResolutionItems, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateJudgeResolutions(id int, resolution *structs.JudgeResolutions) (*structs.JudgeResolutions, error) {
	res := &dto.GetJudgeResolutionResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JudgeResolutions+"/"+strconv.Itoa(id), resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateJudgeResolutions(resolution *structs.JudgeResolutions) (*structs.JudgeResolutions, error) {
	res := &dto.GetJudgeResolutionResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JudgeResolutions, resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteJudgeResolution(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JudgeResolutions+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetJudgeResolutionList(input *dto.GetJudgeResolutionListInputMS) (*dto.GetJudgeResolutionListResponseMS, error) {
	res := &dto.GetJudgeResolutionListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JudgeResolutions, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetJudgeResolution(id int) (*structs.JudgeResolutions, error) {
	res := &dto.GetJudgeResolutionResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JudgeResolutions+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetJudgeNormListByEmployee(userProfileID int) ([]structs.JudgeNorms, error) {
	res := &dto.GetEmployeeNormListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(userProfileID)+"/norms", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitItem) (*dto.JudgeResolutionsOrganizationUnitItem, error) {
	res := &dto.GetJudgeResolutionsOrganizationUnitResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.Judges, input, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitItem) (*dto.JudgeResolutionsOrganizationUnitItem, error) {
	res := &dto.GetJudgeResolutionsOrganizationUnitResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.Judges+"/"+strconv.Itoa(input.ID), input, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitInput) ([]dto.JudgeResolutionsOrganizationUnitItem, int, error) {
	res := &dto.GetJudgeResolutionsOrganizationUnitListMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Judges, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func (repo *MicroserviceRepository) DeleteJJudgeResolutionOrganizationUnit(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.Judges+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
