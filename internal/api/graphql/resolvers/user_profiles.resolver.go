package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserProfilesOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []dto.UserProfileOverviewResponse
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"].(int)
	size := params.Args["size"].(int)
	organizationUnitID := params.Args["organization_unit_id"]
	jobPositionId := params.Args["job_position_id"]
	isActive, isActiveOk := params.Args["is_active"].(bool)
	name, nameOk := params.Args["name"].(string)

	if id != nil && shared.IsInteger(id) && id != 0 {
		user, err := r.Repo.GetUserProfileById(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resItem, err := buildUserProfileOverviewResponse(r.Repo, user)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		items = []dto.UserProfileOverviewResponse{*resItem}
		total = 1
	} else {
		input := dto.GetUserProfilesInput{}
		profiles, err := r.Repo.GetUserProfiles(&input)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		total = len(profiles)
		for _, userProfile := range profiles {
			resItem, err := buildUserProfileOverviewResponse(r.Repo, userProfile)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			if isActiveOk &&
				resItem.Active != isActive {
				total--
				continue
			}
			if nameOk && name != "" && !strings.Contains(strings.ToLower(resItem.FirstName), strings.ToLower(name)) && !strings.Contains(strings.ToLower(resItem.LastName), strings.ToLower(name)) {
				total--
				continue
			}
			if shared.IsInteger(organizationUnitID) && organizationUnitID.(int) > 0 &&
				resItem.OrganizationUnit.Id != organizationUnitID {
				total--
				continue
			}
			if shared.IsInteger(jobPositionId) && jobPositionId.(int) > 0 &&
				resItem.JobPosition.Id != jobPositionId {
				total--
				continue
			}

			items = append(items, *resItem)

		}
	}

	paginatedItems, _ := shared.Paginate(items, page, size)

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   paginatedItems,
		Total:   total,
	}, nil

}

func (r *Resolver) UserProfileContractsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]

	contracts, err := r.Repo.GetEmployeeContracts(id.(int), nil)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	contractResponseItems, err := buildContractResponseItemList(r.Repo, contracts)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   contractResponseItems,
	}, nil
}

func buildUserProfileOverviewResponse(
	r repository.MicroserviceRepositoryInterface,
	profile *structs.UserProfiles,
) (*dto.UserProfileOverviewResponse, error) {
	var (
		organizationUnitDropdown structs.SettingsDropdown
		jobPositionDropdown      structs.SettingsDropdown
		isJudge, isPresident     bool
	)
	account, err := r.GetUserAccountById(profile.UserAccountId)
	if err != nil {
		return nil, err
	}

	role, err := r.GetRole(account.RoleId)
	if err != nil {
		return nil, err
	}

	employeesInOrganizationUnit, _ := r.GetEmployeesInOrganizationUnitsByProfileId(profile.Id)

	if employeesInOrganizationUnit != nil {
		jobPositionInOrganizationUnit, err := r.GetJobPositionsInOrganizationUnitsById(employeesInOrganizationUnit.PositionInOrganizationUnitId)
		if err != nil {
			return nil, err
		}

		jobPosition, err := r.GetJobPositionById(jobPositionInOrganizationUnit.JobPositionId)
		if err != nil {
			return nil, err
		}
		jobPositionDropdown.Id = jobPosition.Id
		jobPositionDropdown.Title = jobPosition.Title

		systematization, _ := r.GetSystematizationById(jobPositionInOrganizationUnit.SystematizationId)

		organizationUnit, err := r.GetOrganizationUnitById(systematization.OrganizationUnitId)
		if err != nil {
			return nil, err
		}
		organizationUnitDropdown.Id = organizationUnit.Id
		organizationUnitDropdown.Title = organizationUnit.Title
	}

	contract, err := r.GetEmployeeContracts(profile.Id, nil)

	if err != nil {
		return nil, err
	}

	if len(contract) > 0 {
		orgUnit, err := r.GetOrganizationUnitById(contract[0].OrganizationUnitID)
		if err != nil {
			return nil, err
		}
		organizationUnitDropdown.Id = orgUnit.Id
		organizationUnitDropdown.Title = orgUnit.Title

	}

	active := true
	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}
	resolution, _ := r.GetJudgeResolutionList(&input)

	if len(resolution.Data) > 0 {

		judgeResolutionOrganizationUnit, _, err := r.GetJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
			UserProfileId: &profile.Id,
			ResolutionId:  &resolution.Data[0].Id,
		})

		if err != nil {
			return nil, err
		}

		if len(judgeResolutionOrganizationUnit) > 0 {
			isJudge = true
			isPresident = judgeResolutionOrganizationUnit[0].IsPresident
		}
	}

	if isJudge {
		filter := dto.JudgeResolutionsOrganizationUnitInput{
			UserProfileId: &profile.Id,
		}
		judge, _, err := r.GetJudgeResolutionOrganizationUnit(&filter)

		if err != nil {
			return nil, err
		}

		orgUnit, err := r.GetOrganizationUnitById(judge[0].OrganizationUnitId)

		if err != nil {
			return nil, err
		}

		organizationUnitDropdown.Id = orgUnit.Id
		organizationUnitDropdown.Title = orgUnit.Title
	}

	return &dto.UserProfileOverviewResponse{
		ID:          profile.Id,
		FirstName:   profile.FirstName,
		LastName:    profile.LastName,
		DateOfBirth: profile.DateOfBirth,
		Email:       account.Email,
		Phone:       account.Phone,
		Active:      account.Active,
		IsJudge:     isJudge,
		IsPresident: isPresident,
		Role: structs.SettingsDropdown{
			Id:    role.Id,
			Title: role.Title,
		},
		OrganizationUnit: organizationUnitDropdown,
		JobPosition:      jobPositionDropdown,
		CreatedAt:        profile.CreatedAt,
		UpdatedAt:        profile.UpdatedAt,
	}, nil
}

func (r *Resolver) UserProfileBasicResolver(params graphql.ResolveParams) (interface{}, error) {
	profileId := params.Args["user_profile_id"]

	profile, err := r.Repo.GetUserProfileById(profileId.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	res, err := buildUserProfileBasicResponse(r.Repo, profile)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Item:    res,
	}, nil
}

func (r *Resolver) UserProfileBasicInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var userAccountData structs.UserAccounts
	var userProfileData structs.UserProfiles
	var activeContract dto.MutateUserProfileActiveContract

	var userAccountRes *structs.UserAccounts
	var userProfileRes *structs.UserProfiles

	dataBytes, _ := json.Marshal(params.Args["data"])

	userAccountData.Id = userProfileData.UserAccountId

	err = json.Unmarshal(dataBytes, &userAccountData)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return errors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &userProfileData)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return errors.HandleAPIError(err)
	}
	err = json.Unmarshal(dataBytes, &activeContract)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return errors.HandleAPIError(err)
	}

	active := true
	inactive := false

	if activeContract.Contract != nil {
		userProfileData.ActiveContract = &active
	} else {
		userProfileData.ActiveContract = &inactive
	}

	userAccountRes, err = r.Repo.CreateUserAccount(userAccountData)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	userProfileData.UserAccountId = userAccountRes.Id
	userProfileRes, err = r.Repo.CreateUserProfile(userProfileData)
	if err != nil {
		_ = r.Repo.DeleteUserAccount(userAccountRes.Id)
		return errors.HandleAPIError(err)
	}

	if activeContract.Contract != nil {
		activeContract.Contract.UserProfileId = userProfileRes.Id
		_, err := r.Repo.CreateEmployeeContract(activeContract.Contract)
		if err != nil {
			_ = r.Repo.DeleteUserAccount(userAccountRes.Id)
			_ = r.Repo.DeleteUserProfile(userProfileRes.Id)
			return errors.HandleAPIError(err)
		}

		if activeContract.Contract.JobPositionInOrganizationUnitID > 0 {
			input := &structs.EmployeesInOrganizationUnits{
				PositionInOrganizationUnitId: activeContract.Contract.JobPositionInOrganizationUnitID,
				UserProfileId:                userProfileRes.Id,
			}
			_, err = r.Repo.CreateEmployeesInOrganizationUnits(input)
			if err != nil {
				_ = r.Repo.DeleteUserAccount(userAccountRes.Id)
				_ = r.Repo.DeleteUserProfile(userProfileRes.Id)
				return errors.HandleAPIError(err)
			}
		}

	}

	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}
	resolution, _ := r.Repo.GetJudgeResolutionList(&input)

	if len(resolution.Data) > 0 {
		if userProfileData.IsJudge {
			inputCreate := dto.JudgeResolutionsOrganizationUnitItem{
				UserProfileId:      userProfileRes.Id,
				OrganizationUnitId: activeContract.Contract.OrganizationUnitID,
				ResolutionId:       resolution.Data[0].Id,
				IsPresident:        userProfileData.IsPresident,
			}
			_, err := r.Repo.CreateJudgeResolutionOrganizationUnit(&inputCreate)
			if err != nil {
				_ = r.Repo.DeleteUserAccount(userAccountRes.Id)
				_ = r.Repo.DeleteUserProfile(userProfileRes.Id)
				return errors.HandleAPIError(err)
			}
		}
	}

	res, err := buildUserProfileBasicResponse(r.Repo, userProfileRes)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
		Item:    res,
	}, nil
}

func (r *Resolver) UserProfileUpdateResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var userProfileData structs.UserProfiles
	var activeContract dto.MutateUserProfileActiveContract

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &userProfileData)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	err = json.Unmarshal(dataBytes, &activeContract)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	active := true
	inactive := false
	if activeContract.Contract != nil {
		userProfileData.ActiveContract = &active
		activeContract.Contract.UserProfileId = userProfileData.Id
		if activeContract.Contract.Id == 0 {
			_, err = r.Repo.CreateEmployeeContract(activeContract.Contract)
			if err != nil {
				return errors.HandleAPIError(err)
			}
		} else {
			_, err = r.Repo.UpdateEmployeeContract(activeContract.Contract.Id, activeContract.Contract)
			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
		if activeContract.Contract.JobPositionInOrganizationUnitID > -1 {
			var myBool int = 2
			var check bool = true
			inputSys := dto.GetSystematizationsInput{}
			inputSys.OrganizationUnitID = &activeContract.Contract.OrganizationUnitID
			inputSys.Active = &myBool

			systematizationsResponse, err := r.Repo.GetSystematizations(&inputSys)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			if len(systematizationsResponse.Data) > 0 {
				for _, systematization := range systematizationsResponse.Data {
					inputJpbPos := dto.GetJobPositionInOrganizationUnitsInput{
						SystematizationID: &systematization.Id,
					}
					jobPositionsInOrganizationUnits, err := r.Repo.GetJobPositionsInOrganizationUnits(&inputJpbPos)
					if err != nil {
						return errors.HandleAPIError(err)
					}
					if len(jobPositionsInOrganizationUnits.Data) > 0 {
						for _, job := range jobPositionsInOrganizationUnits.Data {

							input := dto.GetEmployeesInOrganizationUnitInput{
								PositionInOrganizationUnit: &job.Id,
								UserProfileId:              &userProfileData.Id,
							}
							employeesInOrganizationUnit, _ := r.Repo.GetEmployeesInOrganizationUnitList(&input)
							if len(employeesInOrganizationUnit) > 0 {
								for _, emp := range employeesInOrganizationUnit {
									if emp.UserProfileId == userProfileData.Id {
										if emp.PositionInOrganizationUnitId == activeContract.Contract.JobPositionInOrganizationUnitID {
											check = false
										} else {
											err := r.Repo.DeleteEmployeeInOrganizationUnit(emp.Id)
											if err != nil {
												return errors.HandleAPIError(err)
											}
										}
									}
								}
							}

						}
					}
				}
			}
			if activeContract.Contract.JobPositionInOrganizationUnitID > 0 && check {
				input := &structs.EmployeesInOrganizationUnits{
					PositionInOrganizationUnitId: activeContract.Contract.JobPositionInOrganizationUnitID,
					UserProfileId:                userProfileData.Id,
				}
				_, err = r.Repo.CreateEmployeesInOrganizationUnits(input)
				if err != nil {
					return errors.HandleAPIError(err)
				}
			}
		}
	} else {
		userProfileData.ActiveContract = &inactive
	}

	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}
	resolution, _ := r.Repo.GetJudgeResolutionList(&input)

	if len(resolution.Data) > 0 {
		judgeResolutionOrganizationUnit, _, _ := r.Repo.GetJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
			OrganizationUnitId: &activeContract.Contract.OrganizationUnitID,
			UserProfileId:      &userProfileData.Id,
			ResolutionId:       &resolution.Data[0].Id,
		})

		if len(judgeResolutionOrganizationUnit) > 0 {
			if userProfileData.IsJudge {
				inputUpdate := dto.JudgeResolutionsOrganizationUnitItem{
					Id:                 judgeResolutionOrganizationUnit[0].Id,
					UserProfileId:      userProfileData.Id,
					OrganizationUnitId: activeContract.Contract.OrganizationUnitID,
					ResolutionId:       resolution.Data[0].Id,
					IsPresident:        userProfileData.IsPresident,
				}
				_, err := r.Repo.UpdateJudgeResolutionOrganizationUnit(&inputUpdate)
				if err != nil {
					return errors.HandleAPIError(err)
				}
			} else {
				err := r.Repo.DeleteJJudgeResolutionOrganizationUnit(judgeResolutionOrganizationUnit[0].Id)
				if err != nil {
					return errors.HandleAPIError(err)
				}
			}
		}
		if len(judgeResolutionOrganizationUnit) == 0 && userProfileData.IsJudge {
			inputCreate := dto.JudgeResolutionsOrganizationUnitItem{
				UserProfileId:      userProfileData.Id,
				OrganizationUnitId: activeContract.Contract.OrganizationUnitID,
				ResolutionId:       resolution.Data[0].Id,
				IsPresident:        userProfileData.IsPresident,
			}
			_, err := r.Repo.CreateJudgeResolutionOrganizationUnit(&inputCreate)
			if err != nil {
				return errors.HandleAPIError(err)
			}
		}

	}

	userProfileRes, err := r.Repo.UpdateUserProfile(userProfileData.Id, userProfileData)
	if err != nil {
		fmt.Printf("Creating the user profile failed because of this error - %s.\n", err)
		return errors.ErrorResponse("Error creating the user profile data"), nil
	}

	res, err := buildUserProfileBasicResponse(r.Repo, userProfileRes)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    res,
	}, nil
}

func (r *Resolver) UserProfileContractInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var data structs.Contracts

	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return errors.ErrorResponse("Error updating user profile contract data"), nil
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		item, err := r.Repo.UpdateEmployeeContract(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		contractResponseItem, err := buildContractResponseItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = contractResponseItem
	} else {
		item, err := r.Repo.CreateEmployeeContract(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		contractResponseItem, err := buildContractResponseItem(r.Repo, *item)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = contractResponseItem
	}

	return response, nil
}

func (r *Resolver) UserProfileContractDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := r.Repo.DeleteEmployeeContract(itemId.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildEducationResItem(r repository.MicroserviceRepositoryInterface, education structs.Education) (*dto.Education, error) {
	educationType, err := r.GetDropdownSettingById(education.TypeId)
	if err != nil {
		return nil, err
	}
	educationResItem := &dto.Education{
		Id:                  education.Id,
		Title:               education.Title,
		Description:         education.Description,
		Price:               education.Price,
		DateOfStart:         education.DateOfStart,
		DateOfEnd:           education.DateOfEnd,
		DateOfCertification: education.DateOfCertification,
		AcademicTitle:       education.AcademicTitle,
		CertificateIssuer:   education.CertificateIssuer,
		Score:               education.Score,
		CreatedAt:           education.CreatedAt,
		UpdatedAt:           education.UpdatedAt,
		FileId:              education.FileId,
		UserProfileId:       education.UserProfileId,
		ExpertiseLevel:      education.ExpertiseLevel,
	}

	educationResItem.Type = dto.DropdownSimple{Id: educationType.Id, Title: educationType.Title}

	return educationResItem, nil
}

func buildEducationResItemList(r repository.MicroserviceRepositoryInterface, educations []structs.Education) (educationResItemList []*dto.Education, err error) {
	for _, education := range educations {
		educationResItem, err := buildEducationResItem(r, education)
		if err != nil {
			return nil, err
		}
		educationResItemList = append(educationResItemList, educationResItem)
	}
	return
}

func (r *Resolver) UserProfileEducationResolver(params graphql.ResolveParams) (interface{}, error) {
	userProfileID := params.Args["user_profile_id"].(int)
	educationType := params.Args["education_type"].(string)
	var responseItemList []*dto.Education

	educationTypes, err := r.Repo.GetDropdownSettings(&dto.GetSettingsInput{
		Entity: educationType,
	})
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, educationType := range educationTypes.Data {
		educations, err := r.Repo.GetEmployeeEducations(dto.EducationInput{
			UserProfileID: userProfileID,
			TypeID:        &educationType.Id,
		})
		if err != nil {
			return errors.HandleAPIError(err)
		}
		educationResItemList, err := buildEducationResItemList(r.Repo, educations)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		responseItemList = append(responseItemList, educationResItemList...)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   responseItemList,
	}, nil
}

func (r *Resolver) UserProfileEducationInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		data              structs.Education
		err               error
		employeeEducation *structs.Education
	)

	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}
	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		employeeEducation, err = r.Repo.UpdateEmployeeEducation(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
	} else {
		employeeEducation, err = r.Repo.CreateEmployeeEducation(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
	}

	responseItem, err := buildEducationResItem(r.Repo, *employeeEducation)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	response.Item = responseItem

	return response, nil
}

func (r *Resolver) UserProfileEducationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := r.Repo.DeleteEmployeeEducation(itemId.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) UserProfileExperienceResolver(params graphql.ResolveParams) (interface{}, error) {
	userProfileID := params.Args["user_profile_id"]

	experiences, err := r.Repo.GetEmployeeExperiences(userProfileID.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   experiences,
	}, nil
}

func (r *Resolver) UserProfileExperienceInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var data structs.Experience

	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return errors.ErrorResponse("Error updating experience data"), nil
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		item, err := r.Repo.UpdateExperience(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := r.Repo.CreateExperience(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func (r *Resolver) UserProfileExperienceDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := r.Repo.DeleteExperience(itemId.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) UserProfileFamilyResolver(params graphql.ResolveParams) (interface{}, error) {
	userProfileID := params.Args["user_profile_id"].(int)

	familyMembers, err := r.Repo.GetEmployeeFamilyMembers(userProfileID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   familyMembers,
	}, nil
}

func (r *Resolver) UserProfileFamilyInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Family
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := r.Repo.UpdateEmployeeFamilyMember(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = res
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateEmployeeFamilyMember(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = res
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) UserProfileFamilyDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := r.Repo.DeleteEmployeeFamilyMember(itemId.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildContractResponseItemList(r repository.MicroserviceRepositoryInterface, contracts []*structs.Contracts) (contractResponseItemList []*dto.Contract, err error) {
	for _, contract := range contracts {
		contractResItem, err := buildContractResponseItem(r, *contract)
		if err != nil {
			return nil, err
		}
		contractResponseItemList = append(contractResponseItemList, contractResItem)
	}
	return
}

func buildContractResponseItem(r repository.MicroserviceRepositoryInterface, contract structs.Contracts) (*dto.Contract, error) {
	responseContract := &dto.Contract{
		Id:                 contract.Id,
		Title:              contract.Title,
		Abbreviation:       contract.Abbreviation,
		Description:        contract.Description,
		NumberOfConference: contract.NumberOfConference,
		Active:             contract.Active,
		SerialNumber:       contract.SerialNumber,
		NetSalary:          contract.NetSalary,
		GrossSalary:        contract.GrossSalary,
		BankAccount:        contract.BankAccount,
		BankName:           contract.BankName,
		DateOfSignature:    contract.DateOfSignature,
		DateOfStart:        contract.DateOfStart,
		DateOfEnd:          contract.DateOfEnd,
		DateOfEligibility:  contract.DateOfEligibility,
		CreatedAt:          contract.CreatedAt,
		UpdatedAt:          contract.UpdatedAt,
		FileId:             contract.FileId,
	}

	contractType, err := r.GetDropdownSettingById(contract.ContractTypeId)
	if err != nil {
		return nil, err
	}
	responseContract.ContractType = dto.DropdownSimple{Id: contractType.Id, Title: contractType.Title}

	userProfile, err := r.GetUserProfileById(contract.UserProfileId)
	if err != nil {
		return nil, err
	}
	responseContract.UserProfile = dto.DropdownSimple{Id: userProfile.Id, Title: userProfile.GetFullName()}

	organizationUnit, err := r.GetOrganizationUnitById(contract.OrganizationUnitID)
	if err != nil {
		return nil, err
	}
	responseContract.OrganizationUnit = dto.DropdownSimple{Id: organizationUnit.Id, Title: organizationUnit.Title}

	if contract.Active {
		if contract.OrganizationUnitDepartmentID != nil {
			department, err := r.GetOrganizationUnitById(*contract.OrganizationUnitDepartmentID)
			if err != nil {
				return nil, err
			}
			responseContract.Department = &dto.DropdownSimple{Id: department.Id, Title: department.Title}
		}

		var myBool int = 2
		inputSys := dto.GetSystematizationsInput{}
		inputSys.OrganizationUnitID = &contract.OrganizationUnitID
		inputSys.Active = &myBool

		systematizationsResponse, err := r.GetSystematizations(&inputSys)
		if err != nil {
			return nil, err
		}
		var jobPositionInOU structs.JobPositionsInOrganizationUnits
		if len(systematizationsResponse.Data) > 0 {
			for _, systematization := range systematizationsResponse.Data {

				inputJpbPos := dto.GetJobPositionInOrganizationUnitsInput{
					SystematizationID: &systematization.Id,
				}
				jobPositionsInOrganizationUnits, err := r.GetJobPositionsInOrganizationUnits(&inputJpbPos)
				if err != nil {
					return nil, err
				}
				if len(jobPositionsInOrganizationUnits.Data) > 0 {
					for _, job := range jobPositionsInOrganizationUnits.Data {
						input := dto.GetEmployeesInOrganizationUnitInput{
							PositionInOrganizationUnit: &job.Id,
							UserProfileId:              &userProfile.Id,
						}
						employeesInOrganizationUnit, _ := r.GetEmployeesInOrganizationUnitList(&input)
						if len(employeesInOrganizationUnit) > 0 && employeesInOrganizationUnit[0].UserProfileId == userProfile.Id {
							jobPositionInOU = job
						}
					}
				}
			}
		}
		if jobPositionInOU.JobPositionId > 0 {
			jobPosition, err := r.GetJobPositionById(jobPositionInOU.JobPositionId)
			if err != nil {
				return nil, err
			}
			responseContract.JobPositionInOrganizationUnit = dto.DropdownSimple{Id: jobPositionInOU.Id, Title: jobPosition.Title}
		}
	}

	return responseContract, nil
}

func buildUserProfileBasicResponse(
	r repository.MicroserviceRepositoryInterface,
	profile *structs.UserProfiles,
) (*dto.UserProfileBasicResponse, error) {
	account, err := r.GetUserAccountById(profile.UserAccountId)

	if err != nil {
		return nil, err
	}

	var (
		jobPosition                     *structs.JobPositions
		organizationUnit                *structs.OrganizationUnits
		jobPositionInOrganizationUnitID int
	)

	employeesInOrganizationUnit, _ := r.GetEmployeesInOrganizationUnitsByProfileId(profile.Id)
	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok && apiErr.StatusCode != 404 {
			return nil, err
		}
	}

	if employeesInOrganizationUnit != nil {
		jobPositionInOrganizationUnit, err := r.GetJobPositionsInOrganizationUnitsById(employeesInOrganizationUnit.PositionInOrganizationUnitId)
		if err != nil {
			return nil, err
		}
		jobPositionInOrganizationUnitID = jobPositionInOrganizationUnit.Id

		jobPosition, err = r.GetJobPositionById(jobPositionInOrganizationUnit.JobPositionId)
		if err != nil {
			return nil, err
		}

		organizationUnit, err = r.GetOrganizationUnitById(jobPositionInOrganizationUnit.ParentOrganizationUnitId)
		if err != nil {
			return nil, err
		}
	}

	userProfileResItem := &dto.UserProfileBasicResponse{
		ID:                            profile.Id,
		FirstName:                     profile.FirstName,
		LastName:                      profile.LastName,
		DateOfBirth:                   profile.DateOfBirth,
		BirthLastName:                 profile.BirthLastName,
		CountryOfBirth:                profile.CountryOfBirth,
		CityOfBirth:                   profile.CityOfBirth,
		Nationality:                   profile.Nationality,
		Citizenship:                   profile.Citizenship,
		Address:                       profile.Address,
		FatherName:                    profile.FatherName,
		MotherName:                    profile.MotherName,
		MotherBirthLastName:           profile.MotherBirthLastName,
		BankAccount:                   profile.BankAccount,
		BankName:                      profile.BankName,
		PersonalID:                    profile.PersonalID,
		OfficialPersonalID:            profile.OfficialPersonalId,
		OfficialPersonalDocNumber:     profile.OfficialPersonalDocumentNumber,
		OfficialPersonalDocIssuer:     profile.OfficialPersonalDocumentIssuer,
		Gender:                        profile.Gender,
		SingleParent:                  profile.SingleParent,
		HousingDone:                   profile.HousingDone,
		HousingDescription:            profile.HousingDescription,
		IsPresident:                   false,
		IsJudge:                       false,
		MaritalStatus:                 profile.MaritalStatus,
		DateOfTakingOath:              profile.DateOfTakingOath,
		DateOfBecomingJudge:           profile.DateOfBecomingJudge,
		Email:                         account.Email,
		Phone:                         account.Phone,
		OrganizationUnit:              organizationUnit,
		JobPosition:                   jobPosition,
		JobPositionInOrganizationUnit: jobPositionInOrganizationUnitID,
		NationalMinority:              profile.NationalMinority,
	}
	active := true
	contracts, err := r.GetEmployeeContracts(profile.Id, nil)
	if err != nil {
		return nil, err
	}
	if len(contracts) > 0 {

		contractResponseItem, err := buildContractResponseItem(r, *contracts[0])
		if err != nil {
			return nil, err
		}
		userProfileResItem.Contract = contractResponseItem

		// need check user is judge or president
		if contractResponseItem.OrganizationUnit.Id > 0 {

			input := dto.GetJudgeResolutionListInputMS{
				Active: &active,
			}

			resolution, _ := r.GetJudgeResolutionList(&input)

			if len(resolution.Data) > 0 {
				resolutionId := resolution.Data[0].Id
				judgeResolutionOrganizationUnit, _, err := r.GetJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
					OrganizationUnitId: &contractResponseItem.OrganizationUnit.Id,
					ResolutionId:       &resolutionId,
					UserProfileId:      &profile.Id,
				})
				if err != nil {
					return nil, err
				}

				if len(judgeResolutionOrganizationUnit) > 0 {
					for _, judge := range judgeResolutionOrganizationUnit {
						userProfileResItem.IsPresident = judge.IsPresident
						userProfileResItem.IsJudge = true
					}
				}
			}
		}

	}

	return userProfileResItem, nil
}
