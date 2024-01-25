package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetTenderTypeList(input *dto.GetJobTenderTypeInputMS) ([]*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JobTenderTypes, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetTenderType(id int) (*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JobTenderTypes+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteJobTenderType(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JobTenderTypes+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateJobTenderType(jobTender *structs.JobTenderTypes) (*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JobTenderTypes, jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateJobTenderType(id int, jobTender *structs.JobTenderTypes) (*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JobTenderTypes+"/"+strconv.Itoa(id), jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateJobTender(jobTender *structs.JobTenders) (*structs.JobTenders, error) {
	res := &dto.GetJobTenderResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JobTenders, jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateJobTender(id int, jobTender *structs.JobTenders) (*structs.JobTenders, error) {
	res := &dto.GetJobTenderResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JobTenders+"/"+strconv.Itoa(id), jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetJobTender(id int) (*structs.JobTenders, error) {
	res := &dto.GetJobTenderResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JobTenders+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetJobTenderList() ([]*structs.JobTenders, error) {
	res := &dto.GetJobTenderListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JobTenders, nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteJobTender(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JobTenders+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateJobTenderApplication(jobTender *structs.JobTenderApplications) (*structs.JobTenderApplications, error) {
	res := &dto.GetJobTenderApplicationResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JobTenderApplications, jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateJobTenderApplication(id int, jobTender *structs.JobTenderApplications) (*structs.JobTenderApplications, error) {
	currentTenderApplication, _ := repo.GetTenderApplication(id)
	if currentTenderApplication.Status != "Izabran" && jobTender.Status == "Izabran" {
		applications, _ := repo.GetTenderApplicationList(&dto.GetJobTenderApplicationsInput{JobTenderID: &currentTenderApplication.JobTenderID})
		jobTender, _ := repo.GetJobTender(currentTenderApplication.JobTenderID)

		count := 0
		for _, tenderApp := range applications.Data {
			if tenderApp.Status == "Izabran" {
				count++
			}
		}

		if count == jobTender.NumberOfVacantSeats {
			for _, application := range applications.Data {
				if currentTenderApplication.ID != application.ID && application.Status != "Izabran" {
					application.Status = "Nije izabran"
					_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JobTenderApplications+"/"+strconv.Itoa(application.ID), application, nil)
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}
	res := &dto.GetJobTenderApplicationResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JobTenderApplications+"/"+strconv.Itoa(id), jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteJobTenderApplication(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JobTenderApplications+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetTenderApplication(id int) (*structs.JobTenderApplications, error) {
	res := &dto.GetJobTenderApplicationResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JobTenderApplications+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetTenderApplicationList(input *dto.GetJobTenderApplicationsInput) (*dto.GetJobTenderApplicationListResponseMS, error) {
	res := &dto.GetJobTenderApplicationListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JobTenderApplications, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
