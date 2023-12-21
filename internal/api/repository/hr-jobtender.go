package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"strconv"
)

func (repo *MicroserviceRepository) GetTenderTypeList(input *dto.GetJobTenderTypeInputMS) ([]*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JOB_TENDER_TYPES, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetTenderType(id int) (*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JOB_TENDER_TYPES+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteJobTenderType(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JOB_TENDER_TYPES+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateJobTenderType(jobTender *structs.JobTenderTypes) (*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JOB_TENDER_TYPES, jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateJobTenderType(id int, jobTender *structs.JobTenderTypes) (*structs.JobTenderTypes, error) {
	res := &dto.GetJobTenderTypeResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JOB_TENDER_TYPES+"/"+strconv.Itoa(id), jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateJobTender(jobTender *structs.JobTenders) (*structs.JobTenders, error) {
	res := &dto.GetJobTenderResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JOB_TENDERS, jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateJobTender(id int, jobTender *structs.JobTenders) (*structs.JobTenders, error) {
	res := &dto.GetJobTenderResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JOB_TENDERS+"/"+strconv.Itoa(id), jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetJobTender(id int) (*structs.JobTenders, error) {
	res := &dto.GetJobTenderResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JOB_TENDERS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetJobTenderList() ([]*structs.JobTenders, error) {
	res := &dto.GetJobTenderListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JOB_TENDERS, nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteJobTender(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JOB_TENDERS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateJobTenderApplication(jobTender *structs.JobTenderApplications) (*structs.JobTenderApplications, error) {
	res := &dto.GetJobTenderApplicationResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.JOB_TENDER_APPLICATIONS, jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateJobTenderApplication(id int, jobTender *structs.JobTenderApplications) (*structs.JobTenderApplications, error) {
	currentTenderApplication, _ := repo.GetTenderApplication(id)
	if currentTenderApplication.Status != "Izabran" && jobTender.Status == "Izabran" {
		applications, _ := repo.GetTenderApplicationList(&dto.GetJobTenderApplicationsInput{JobTenderID: &currentTenderApplication.JobTenderId})
		for _, application := range applications.Data {
			if currentTenderApplication.Id != application.Id {
				application.Status = "Nije izabran"
				_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JOB_TENDER_APPLICATIONS+"/"+strconv.Itoa(application.Id), application, nil)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	res := &dto.GetJobTenderApplicationResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.JOB_TENDER_APPLICATIONS+"/"+strconv.Itoa(id), jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteJobTenderApplication(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.JOB_TENDER_APPLICATIONS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetTenderApplication(id int) (*structs.JobTenderApplications, error) {
	res := &dto.GetJobTenderApplicationResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JOB_TENDER_APPLICATIONS+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetTenderApplicationList(input *dto.GetJobTenderApplicationsInput) (*dto.GetJobTenderApplicationListResponseMS, error) {
	res := &dto.GetJobTenderApplicationListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.JOB_TENDER_APPLICATIONS, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
