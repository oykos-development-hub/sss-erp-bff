package repository

import (
	"bff/internal/api/dto"
	"bff/structs"
	"fmt"
	"strconv"
)

func (repo *MicroserviceRepository) GetEmployeeContracts(employeeID int, input *dto.GetEmployeeContracts) ([]*structs.Contracts, error) {
	res := &dto.GetEmployeeContractListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(employeeID)+"/contracts", input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteEmployeeContract(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.EmployeeContracts+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) CreateUserProfile(user structs.UserProfiles) (*structs.UserProfiles, error) {
	res := &dto.GetUserProfileResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.UserProfiles, user, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetUserProfiles(input *dto.GetUserProfilesInput) ([]*structs.UserProfiles, error) {
	res := &dto.GetUserProfileListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetUserProfileByUserAccountID(accountID int) (*structs.UserProfiles, error) {
	input := &dto.GetUserProfilesInput{AccountID: &accountID}
	res := &dto.GetUserProfileListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles, input, res)
	if err != nil {
		return nil, err
	}
	if res.Total != 1 {
		return nil, fmt.Errorf("user profile not created for user account with ID %d", accountID)
	}

	return res.Data[0], nil
}

func (repo *MicroserviceRepository) GetUserProfileByID(id int) (*structs.UserProfiles, error) {
	res := &dto.GetUserProfileResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteUserProfile(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateUserProfile(userID int, user structs.UserProfiles) (*structs.UserProfiles, error) {
	res := &dto.GetUserProfileResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(userID), user, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetEmployeesInOrganizationUnitsByProfileID(profileID int) (*structs.EmployeesInOrganizationUnits, error) {
	res := &dto.GetEmployeesInOrganizationUnitsResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(profileID)+"/employee-in-organization-unit", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetEmployeesInOrganizationUnitList(input *dto.GetEmployeesInOrganizationUnitInput) ([]*structs.EmployeesInOrganizationUnits, error) {
	res := &dto.GetEmployeesInOrganizationUnitsListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.EmployeesInOrganizationUnits, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateEmployeeContract(id int, contract *structs.Contracts) (*structs.Contracts, error) {
	res := &dto.GetUserProfileContractResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.EmployeeContracts+"/"+strconv.Itoa(id), contract, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateEmployeeContract(contract *structs.Contracts) (*structs.Contracts, error) {
	res := &dto.GetUserProfileContractResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.EmployeeContracts, contract, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateEmployeeEducation(education *structs.Education) (*structs.Education, error) {
	res := &dto.GetEmployeeEducationResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.EmployeeEducations, education, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateEmployeeEducation(id int, education *structs.Education) (*structs.Education, error) {
	res := &dto.GetEmployeeEducationResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.EmployeeEducations+"/"+strconv.Itoa(id), education, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteEmployeeEducation(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.EmployeeEducations+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetEmployeeEducations(input dto.EducationInput) ([]structs.Education, error) {
	res := &dto.GetEmployeeEducationListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.EmployeeEducations, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateExperience(id int, contract *structs.Experience) (*structs.Experience, error) {
	res := &dto.ExperienceItemResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.EmployeeExperiences+"/"+strconv.Itoa(id), contract, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateExperience(contract *structs.Experience) (*structs.Experience, error) {
	res := &dto.ExperienceItemResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.EmployeeExperiences, contract, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteExperience(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.EmployeeExperiences+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetEmployeeExperiences(employeeID int) ([]*structs.Experience, error) {
	res := &dto.GetEmployeeExperienceListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(employeeID)+"/experiences", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateEmployeeFamilyMember(familyMember *structs.Family) (*structs.Family, error) {
	res := &dto.GetEmployeeFamilyMemberResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.EmployeeFamilyMembers, familyMember, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateEmployeeFamilyMember(id int, education *structs.Family) (*structs.Family, error) {
	res := &dto.GetEmployeeFamilyMemberResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.EmployeeFamilyMembers+"/"+strconv.Itoa(id), education, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteEmployeeFamilyMember(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.EmployeeFamilyMembers+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetEmployeeFamilyMembers(employeeID int) ([]*structs.Family, error) {
	res := &dto.GetEmployeeFamilyMemberListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(employeeID)+"/family-members", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteAbsentType(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.AbsentType+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateAbsentType(id int, absent *structs.AbsentType) (*structs.AbsentType, error) {
	res := &dto.GetAbsentTypeResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.AbsentType+"/"+strconv.Itoa(id), absent, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateAbsentType(absent *structs.AbsentType) (*structs.AbsentType, error) {
	res := &dto.GetAbsentTypeResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.AbsentType, absent, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetAbsentTypes() (*dto.GetAbsentTypeListResponseMS, error) {
	res := &dto.GetAbsentTypeListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.AbsentType, nil, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (repo *MicroserviceRepository) GetAbsentTypeByID(absentTypeID int) (*structs.AbsentType, error) {
	res := &dto.GetAbsentTypeResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.AbsentType+"/"+strconv.Itoa(absentTypeID), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateAbsent(absent *structs.Absent) (*structs.Absent, error) {
	res := &dto.GetAbsentResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.EmployeeAbsents, absent, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateAbsent(id int, absent *structs.Absent) (*structs.Absent, error) {
	res := &dto.GetAbsentResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.EmployeeAbsents+"/"+strconv.Itoa(id), absent, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteAbsent(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.EmployeeAbsents+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetEmployeeAbsents(userProfileID int, input *dto.EmployeeAbsentsInput) ([]*structs.Absent, error) {
	res := &dto.GetAbsentListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(userProfileID)+"/absents", input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetAbsentByID(absentID int) (*structs.Absent, error) {
	res := &dto.GetAbsentResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.EmployeeAbsents+"/"+strconv.Itoa(absentID), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetEmployeeEvaluations(userProfileID int) ([]*structs.Evaluation, error) {
	res := &dto.GetEmployeeEvaluationListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(userProfileID)+"/evaluations", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetEvaluation(evaulationID int) (*structs.Evaluation, error) {
	res := &dto.GetEvaluationResponse{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Evaluations+"/"+strconv.Itoa(evaulationID), nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateEmployeeEvaluation(id int, evaluation *structs.Evaluation) (*structs.Evaluation, error) {
	res := dto.GetEvaluationResponse{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.Evaluations+"/"+strconv.Itoa(id), evaluation, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateEmployeeEvaluation(evaluation *structs.Evaluation) (*structs.Evaluation, error) {
	res := dto.GetEvaluationResponse{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.Evaluations, evaluation, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteEvaluation(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.Evaluations+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetEmployeeForeigners(userProfileID int) ([]*structs.Foreigners, error) {
	res := dto.GetEmployeeForeignersListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(userProfileID)+"/foreigners", nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateEmployeeForeigner(id int, foreigner *structs.Foreigners) (*structs.Foreigners, error) {
	res := dto.GetEmployeeForeignersResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.Foreigners+"/"+strconv.Itoa(id), foreigner, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateEmployeeForeigner(foreigner *structs.Foreigners) (*structs.Foreigners, error) {
	res := dto.GetEmployeeForeignersResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.Foreigners, foreigner, &res)
	//foreigners
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteForeigner(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.Foreigners+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetEmployeeResolutions(employeeID int, input *dto.EmployeeResolutionListInput) ([]*structs.Resolution, error) {
	res := &dto.GetResolutionListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(employeeID)+"/resolutions", input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) GetEmployeeResolution(id int) (*structs.Resolution, error) {
	res := &dto.GetResolutionResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.Resolutions+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) UpdateResolution(id int, resolution *structs.Resolution) (*structs.Resolution, error) {

	res := &dto.GetResolutionResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.Resolutions+"/"+strconv.Itoa(id), resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateResolution(resolution *structs.Resolution) (*structs.Resolution, error) {
	res := &dto.GetResolutionResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.Resolutions, resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) DeleteResolution(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.Resolutions+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) GetEmployeeSalaryParams(userProfileID int) ([]*structs.SalaryParams, error) {
	res := &dto.GetEmployeeSalaryParamsListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.HR.UserProfiles+"/"+strconv.Itoa(userProfileID)+"/salaries", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) CreateEmployeeSalaryParams(salaries *structs.SalaryParams) (*structs.SalaryParams, error) {
	res := dto.GetEmployeeSalaryParamsResponseMS{}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.HR.Salaries, salaries, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) DeleteSalaryParams(id int) error {
	_, err := makeAPIRequest("DELETE", repo.Config.Microservices.HR.Salaries+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MicroserviceRepository) UpdateEmployeeSalaryParams(id int, salaries *structs.SalaryParams) (*structs.SalaryParams, error) {
	res := dto.GetEmployeeSalaryParamsResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.HR.Salaries+"/"+strconv.Itoa(id), salaries, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
