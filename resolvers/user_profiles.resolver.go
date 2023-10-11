package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
)

var UserProfilesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
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
		user, err := getUserProfileById(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resItem, err := buildUserProfileOverviewResponse(user)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		items = []dto.UserProfileOverviewResponse{*resItem}
		total = 1
	} else {
		input := dto.GetUserProfilesInput{}
		profiles, err := getUserProfiles(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		total = len(profiles)
		for _, userProfile := range profiles {
			resItem, err := buildUserProfileOverviewResponse(userProfile)
			if err != nil {
				return shared.HandleAPIError(err)
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

var UserProfileContractsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]

	contracts, err := getEmployeeContracts(id.(int), nil)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	contractResponseItems, err := buildContractResponseItemList(contracts)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   contractResponseItems,
	}, nil
}

func buildUserProfileOverviewResponse(
	profile *structs.UserProfiles,
) (*dto.UserProfileOverviewResponse, error) {
	var (
		organizationUnitDropdown structs.SettingsDropdown
		jobPositionDropdown      structs.SettingsDropdown
		isJudge, isPresident     bool
	)
	account, err := GetUserAccountById(profile.UserAccountId)
	if err != nil {
		return nil, err
	}

	role, err := getRole(account.RoleId)
	if err != nil {
		return nil, err
	}

	employeesInOrganizationUnit, _ := getEmployeesInOrganizationUnitsByProfileId(profile.Id)

	if employeesInOrganizationUnit != nil {
		jobPositionInOrganizationUnit, err := getJobPositionsInOrganizationUnitsById(employeesInOrganizationUnit.PositionInOrganizationUnitId)
		if err != nil {
			return nil, err
		}

		jobPosition, err := getJobPositionById(jobPositionInOrganizationUnit.JobPositionId)
		if err != nil {
			return nil, err
		}
		jobPositionDropdown.Id = jobPosition.Id
		jobPositionDropdown.Title = jobPosition.Title

		systematization, _ := getSystematizationById(jobPositionInOrganizationUnit.SystematizationId)

		organizationUnit, err := getOrganizationUnitById(systematization.OrganizationUnitId)
		if err != nil {
			return nil, err
		}
		organizationUnitDropdown.Id = organizationUnit.Id
		organizationUnitDropdown.Title = organizationUnit.Title
	}

	active := true
	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}
	resolution, _ := getJudgeResolutionList(&input)

	if len(resolution.Data) > 0 {

		judgeResolutionOrganizationUnit, _, err := getJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
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
		judge, _, err := getJudgeResolutionOrganizationUnit(&filter)

		if err != nil {
			return nil, err
		}

		orgUnit, err := getOrganizationUnitById(judge[0].OrganizationUnitId)

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

var UserProfileBasicResolver = func(params graphql.ResolveParams) (interface{}, error) {
	profileId := params.Args["user_profile_id"]

	profile, err := getUserProfileById(profileId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	res, err := buildUserProfileBasicResponse(profile)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Item:    res,
	}, nil
}

var UserProfileBasicInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
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
		return shared.ErrorResponse("Error updating settings data"), nil
	}
	err = json.Unmarshal(dataBytes, &userProfileData)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating settings data"), nil
	}
	err = json.Unmarshal(dataBytes, &activeContract)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating settings data"), nil
	}

	active := true
	inactive := false

	if activeContract.Contract != nil {
		userProfileData.ActiveContract = &active
	} else {
		userProfileData.ActiveContract = &inactive
	}

	userAccountRes, err = CreateUserAccount(userAccountData)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	userProfileData.UserAccountId = userAccountRes.Id
	userProfileRes, err = createUserProfile(userProfileData)
	if err != nil {
		_ = DeleteUserAccount(userAccountRes.Id)
		return shared.HandleAPIError(err)
	}

	if activeContract.Contract != nil {
		activeContract.Contract.UserProfileId = userProfileRes.Id
		_, err := createEmployeeContract(activeContract.Contract)
		if err != nil {
			_ = DeleteUserAccount(userAccountRes.Id)
			_ = DeleteUserProfile(userProfileRes.Id)
			return shared.HandleAPIError(err)
		}

		if activeContract.Contract.JobPositionInOrganizationUnitID > 0 {
			input := &structs.EmployeesInOrganizationUnits{
				PositionInOrganizationUnitId: activeContract.Contract.JobPositionInOrganizationUnitID,
				UserProfileId:                userProfileRes.Id,
			}
			_, err = createEmployeesInOrganizationUnits(input)
			if err != nil {
				_ = DeleteUserAccount(userAccountRes.Id)
				_ = DeleteUserProfile(userProfileRes.Id)
				return shared.HandleAPIError(err)
			}
		}

	}

	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}
	resolution, _ := getJudgeResolutionList(&input)

	if len(resolution.Data) > 0 {
		if userProfileData.IsJudge {
			inputCreate := dto.JudgeResolutionsOrganizationUnitItem{
				UserProfileId:      userProfileRes.Id,
				OrganizationUnitId: activeContract.Contract.OrganizationUnitID,
				ResolutionId:       resolution.Data[0].Id,
				IsPresident:        userProfileData.IsPresident,
			}
			_, err := createJudgeResolutionOrganizationUnit(&inputCreate)
			if err != nil {
				_ = DeleteUserAccount(userAccountRes.Id)
				_ = DeleteUserProfile(userProfileRes.Id)
				return shared.HandleAPIError(err)
			}
		}
	}

	res, err := buildUserProfileBasicResponse(userProfileRes)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You created this item!",
		Item:    res,
	}, nil
}

var UserProfileUpdateResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var userProfileData structs.UserProfiles
	var activeContract dto.MutateUserProfileActiveContract

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &userProfileData)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	err = json.Unmarshal(dataBytes, &activeContract)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	active := true
	inactive := false
	if activeContract.Contract != nil {
		userProfileData.ActiveContract = &active
		activeContract.Contract.UserProfileId = userProfileData.Id
		if activeContract.Contract.Id == 0 {
			_, err = createEmployeeContract(activeContract.Contract)
			if err != nil {
				return shared.HandleAPIError(err)
			}
		} else {
			_, err = updateEmployeeContract(activeContract.Contract.Id, activeContract.Contract)
			if err != nil {
				return shared.HandleAPIError(err)
			}
		}
		if activeContract.Contract.JobPositionInOrganizationUnitID > -1 {
			var myBool int = 2
			var check bool = true
			inputSys := dto.GetSystematizationsInput{}
			inputSys.OrganizationUnitID = &activeContract.Contract.OrganizationUnitID
			inputSys.Active = &myBool

			systematizationsResponse, err := getSystematizations(&inputSys)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			if len(systematizationsResponse.Data) > 0 {
				for _, systematization := range systematizationsResponse.Data {
					inputJpbPos := dto.GetJobPositionInOrganizationUnitsInput{
						SystematizationID: &systematization.Id,
					}
					jobPositionsInOrganizationUnits, err := getJobPositionsInOrganizationUnits(&inputJpbPos)
					if err != nil {
						return shared.HandleAPIError(err)
					}
					if len(jobPositionsInOrganizationUnits.Data) > 0 {
						for _, job := range jobPositionsInOrganizationUnits.Data {

							input := dto.GetEmployeesInOrganizationUnitInput{
								PositionInOrganizationUnit: &job.Id,
								UserProfileId:              &userProfileData.Id,
							}
							employeesInOrganizationUnit, _ := getEmployeesInOrganizationUnitList(&input)
							if len(employeesInOrganizationUnit) > 0 {
								for _, emp := range employeesInOrganizationUnit {
									if emp.UserProfileId == userProfileData.Id {
										if emp.PositionInOrganizationUnitId == activeContract.Contract.JobPositionInOrganizationUnitID {
											check = false
										} else {
											err := deleteEmployeeInOrganizationUnit(emp.Id)
											if err != nil {
												return shared.HandleAPIError(err)
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
				_, err = createEmployeesInOrganizationUnits(input)
				if err != nil {
					return shared.HandleAPIError(err)
				}
			}
		}
	} else {
		userProfileData.ActiveContract = &inactive
	}

	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}
	resolution, _ := getJudgeResolutionList(&input)

	if len(resolution.Data) > 0 {
		judgeResolutionOrganizationUnit, _, _ := getJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
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
				_, err := updateJudgeResolutionOrganizationUnit(&inputUpdate)
				if err != nil {
					return shared.HandleAPIError(err)
				}
			} else {
				err := deleteJJudgeResolutionOrganizationUnit(judgeResolutionOrganizationUnit[0].Id)
				if err != nil {
					return shared.HandleAPIError(err)
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
			_, err := createJudgeResolutionOrganizationUnit(&inputCreate)
			if err != nil {
				return shared.HandleAPIError(err)
			}
		}

	}

	userProfileRes, err := updateUserProfile(userProfileData.Id, userProfileData)
	if err != nil {
		fmt.Printf("Creating the user profile failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error creating the user profile data"), nil
	}

	res, err := buildUserProfileBasicResponse(userProfileRes)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    res,
	}, nil
}

var UserProfileContractInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var data structs.Contracts

	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating user profile contract data"), nil
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		item, err := updateEmployeeContract(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		contractResponseItem, err := buildContractResponseItem(*item)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = contractResponseItem
	} else {
		item, err := createEmployeeContract(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		contractResponseItem, err := buildContractResponseItem(*item)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Message = "You created this item!"
		response.Item = contractResponseItem
	}

	return response, nil
}

var UserProfileContractDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := deleteEmployeeContract(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildEducationResItem(education structs.Education) (*dto.Education, error) {
	educationType, err := getDropdownSettingById(education.TypeId)
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

func buildEducationResItemList(educations []structs.Education) (educationResItemList []*dto.Education, err error) {
	for _, education := range educations {
		educationResItem, err := buildEducationResItem(education)
		if err != nil {
			return nil, err
		}
		educationResItemList = append(educationResItemList, educationResItem)
	}
	return
}

var UserProfileEducationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	userProfileID := params.Args["user_profile_id"].(int)
	educationType := params.Args["education_type"].(string)
	var responseItemList []*dto.Education

	educationTypes, err := getDropdownSettings(&dto.GetSettingsInput{
		Entity: educationType,
	})
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, educationType := range educationTypes.Data {
		educations, err := getEmployeeEducations(dto.EducationInput{
			UserProfileID: userProfileID,
			TypeID:        &educationType.Id,
		})
		if err != nil {
			return shared.HandleAPIError(err)
		}
		educationResItemList, err := buildEducationResItemList(educations)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		responseItemList = append(responseItemList, educationResItemList...)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   responseItemList,
	}, nil
}

var UserProfileEducationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
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
		employeeEducation, err = updateEmployeeEducation(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
	} else {
		employeeEducation, err = createEmployeeEducation(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You created this item!"
	}

	responseItem, err := buildEducationResItem(*employeeEducation)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	response.Item = responseItem

	return response, nil
}

var UserProfileEducationDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := deleteEmployeeEducation(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

var UserProfileExperienceResolver = func(params graphql.ResolveParams) (interface{}, error) {
	userProfileID := params.Args["user_profile_id"]

	experiences, err := getEmployeeExperiences(userProfileID.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   experiences,
	}, nil
}

var UserProfileExperienceInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var data structs.Experience

	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating experience data"), nil
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		item, err := updateExperience(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := createExperience(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

var UserProfileExperienceDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := deleteExperience(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

var UserProfileFamilyResolver = func(params graphql.ResolveParams) (interface{}, error) {
	userProfileID := params.Args["user_profile_id"].(int)

	familyMembers, err := getEmployeeFamilyMembers(userProfileID)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   familyMembers,
	}, nil
}

var UserProfileFamilyInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Family
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateEmployeeFamilyMember(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = res
		response.Message = "You updated this item!"
	} else {
		res, err := createEmployeeFamilyMember(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = res
		response.Message = "You created this item!"
	}

	return response, nil
}

var UserProfileFamilyDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := deleteEmployeeFamilyMember(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildContractResponseItemList(contracts []*structs.Contracts) (contractResponseItemList []*dto.Contract, err error) {
	for _, contract := range contracts {
		contractResItem, err := buildContractResponseItem(*contract)
		if err != nil {
			return nil, err
		}
		contractResponseItemList = append(contractResponseItemList, contractResItem)
	}
	return
}

func buildContractResponseItem(contract structs.Contracts) (*dto.Contract, error) {
	responseContract := &dto.Contract{
		Id:                contract.Id,
		Title:             contract.Title,
		Abbreviation:      contract.Abbreviation,
		Description:       contract.Description,
		Active:            contract.Active,
		SerialNumber:      contract.SerialNumber,
		NetSalary:         contract.NetSalary,
		GrossSalary:       contract.GrossSalary,
		BankAccount:       contract.BankAccount,
		BankName:          contract.BankName,
		DateOfSignature:   contract.DateOfSignature,
		DateOfStart:       contract.DateOfStart,
		DateOfEnd:         contract.DateOfEnd,
		DateOfEligibility: contract.DateOfEligibility,
		CreatedAt:         contract.CreatedAt,
		UpdatedAt:         contract.UpdatedAt,
		FileId:            contract.FileId,
	}

	contractType, err := getDropdownSettingById(contract.ContractTypeId)
	if err != nil {
		return nil, err
	}
	responseContract.ContractType = dto.DropdownSimple{Id: contractType.Id, Title: contractType.Title}

	userProfile, err := getUserProfileById(contract.UserProfileId)
	if err != nil {
		return nil, err
	}
	responseContract.UserProfile = dto.DropdownSimple{Id: userProfile.Id, Title: userProfile.GetFullName()}

	organizationUnit, err := getOrganizationUnitById(contract.OrganizationUnitID)
	if err != nil {
		return nil, err
	}
	responseContract.OrganizationUnit = dto.DropdownSimple{Id: organizationUnit.Id, Title: organizationUnit.Title}

	if contract.Active {
		if contract.OrganizationUnitDepartmentID != nil {
			department, err := getOrganizationUnitById(*contract.OrganizationUnitDepartmentID)
			if err != nil {
				return nil, err
			}
			responseContract.Department = &dto.DropdownSimple{Id: department.Id, Title: department.Title}
		}

		var myBool int = 2
		inputSys := dto.GetSystematizationsInput{}
		inputSys.OrganizationUnitID = &contract.OrganizationUnitID
		inputSys.Active = &myBool

		systematizationsResponse, err := getSystematizations(&inputSys)
		if err != nil {
			return nil, err
		}
		var jobPositionInOU structs.JobPositionsInOrganizationUnits
		if len(systematizationsResponse.Data) > 0 {
			for _, systematization := range systematizationsResponse.Data {

				inputJpbPos := dto.GetJobPositionInOrganizationUnitsInput{
					SystematizationID: &systematization.Id,
				}
				jobPositionsInOrganizationUnits, err := getJobPositionsInOrganizationUnits(&inputJpbPos)
				if err != nil {
					return nil, err
				}
				if len(jobPositionsInOrganizationUnits.Data) > 0 {
					for _, job := range jobPositionsInOrganizationUnits.Data {
						input := dto.GetEmployeesInOrganizationUnitInput{
							PositionInOrganizationUnit: &job.Id,
							UserProfileId:              &userProfile.Id,
						}
						employeesInOrganizationUnit, _ := getEmployeesInOrganizationUnitList(&input)
						if len(employeesInOrganizationUnit) > 0 && employeesInOrganizationUnit[0].UserProfileId == userProfile.Id {
							jobPositionInOU = job
						}
					}
				}
			}
		}
		if jobPositionInOU.JobPositionId > 0 {
			jobPosition, err := getJobPositionById(jobPositionInOU.JobPositionId)
			if err != nil {
				return nil, err
			}
			responseContract.JobPositionInOrganizationUnit = dto.DropdownSimple{Id: jobPositionInOU.Id, Title: jobPosition.Title}
		}
	}

	return responseContract, nil
}

func buildUserProfileBasicResponse(
	profile *structs.UserProfiles,
) (*dto.UserProfileBasicResponse, error) {
	account, err := GetUserAccountById(profile.UserAccountId)

	if err != nil {
		return nil, err
	}

	var (
		jobPosition                     *structs.JobPositions
		organizationUnit                *structs.OrganizationUnits
		jobPositionInOrganizationUnitID int
	)

	employeesInOrganizationUnit, _ := getEmployeesInOrganizationUnitsByProfileId(profile.Id)
	if err != nil {
		if apiErr, ok := err.(*shared.APIError); ok && apiErr.StatusCode != 404 {
			return nil, err
		}
	}

	if employeesInOrganizationUnit != nil {
		jobPositionInOrganizationUnit, err := getJobPositionsInOrganizationUnitsById(employeesInOrganizationUnit.PositionInOrganizationUnitId)
		if err != nil {
			return nil, err
		}
		jobPositionInOrganizationUnitID = jobPositionInOrganizationUnit.Id

		jobPosition, err = getJobPositionById(jobPositionInOrganizationUnit.JobPositionId)
		if err != nil {
			return nil, err
		}

		organizationUnit, err = getOrganizationUnitById(jobPositionInOrganizationUnit.ParentOrganizationUnitId)
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
	contracts, err := getEmployeeContracts(profile.Id, nil)
	if err != nil {
		return nil, err
	}
	if len(contracts) > 0 {

		contractResponseItem, err := buildContractResponseItem(*contracts[0])
		if err != nil {
			return nil, err
		}
		userProfileResItem.Contract = contractResponseItem

		// need check user is judge or president
		if contractResponseItem.OrganizationUnit.Id > 0 {

			input := dto.GetJudgeResolutionListInputMS{
				Active: &active,
			}

			resolution, _ := getJudgeResolutionList(&input)

			if len(resolution.Data) > 0 {
				resolutionId := resolution.Data[0].Id
				judgeResolutionOrganizationUnit, _, err := getJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
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

func getEmployeeContracts(employeeID int, input *dto.GetEmployeeContracts) ([]*structs.Contracts, error) {
	res := &dto.GetEmployeeContractListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(employeeID)+"/contracts", input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func deleteEmployeeContract(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.EMPLOYEE_CONTRACTS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func createUserProfile(user structs.UserProfiles) (*structs.UserProfiles, error) {
	res := &dto.GetUserProfileResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.USER_PROFILES_ENDPOINT, user, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getUserProfiles(input *dto.GetUserProfilesInput) ([]*structs.UserProfiles, error) {
	res := &dto.GetUserProfileListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getUserProfileByUserAccountID(accountID int) (*structs.UserProfiles, error) {
	input := &dto.GetUserProfilesInput{AccountID: &accountID}
	res := &dto.GetUserProfileListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}
	if res.Total != 1 {
		return nil, fmt.Errorf("user profile not created for user account with ID %d", accountID)
	}

	return res.Data[0], nil
}

func getUserProfileById(id int) (*structs.UserProfiles, error) {
	res := &dto.GetUserProfileResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func DeleteUserProfile(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func updateUserProfile(userID int, user structs.UserProfiles) (*structs.UserProfiles, error) {
	res := &dto.GetUserProfileResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(userID), user, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getEmployeesInOrganizationUnitsByProfileId(profileId int) (*structs.EmployeesInOrganizationUnits, error) {
	res := &dto.GetEmployeesInOrganizationUnitsResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(profileId)+"/employee-in-organization-unit", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getEmployeesInOrganizationUnitList(input *dto.GetEmployeesInOrganizationUnitInput) ([]*structs.EmployeesInOrganizationUnits, error) {
	res := &dto.GetEmployeesInOrganizationUnitsListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.EMPLOYEES_IN_ORGANIZATION_UNITS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func updateEmployeeContract(id int, contract *structs.Contracts) (*structs.Contracts, error) {
	res := &dto.GetUserProfileContractResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.EMPLOYEE_CONTRACTS+"/"+strconv.Itoa(id), contract, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createEmployeeContract(contract *structs.Contracts) (*structs.Contracts, error) {
	res := &dto.GetUserProfileContractResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.EMPLOYEE_CONTRACTS, contract, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createEmployeeEducation(education *structs.Education) (*structs.Education, error) {
	res := &dto.GetEmployeeEducationResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.EMPLOYEE_EDUCATIONS, education, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateEmployeeEducation(id int, education *structs.Education) (*structs.Education, error) {
	res := &dto.GetEmployeeEducationResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.EMPLOYEE_EDUCATIONS+"/"+strconv.Itoa(id), education, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteEmployeeEducation(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.EMPLOYEE_EDUCATIONS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getEmployeeEducations(input dto.EducationInput) ([]structs.Education, error) {
	res := &dto.GetEmployeeEducationListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.EMPLOYEE_EDUCATIONS, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func updateExperience(id int, contract *structs.Experience) (*structs.Experience, error) {
	res := &dto.ExperienceItemResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.EMPLOYEE_EXPERIENCES+"/"+strconv.Itoa(id), contract, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createExperience(contract *structs.Experience) (*structs.Experience, error) {
	res := &dto.ExperienceItemResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.EMPLOYEE_EXPERIENCES, contract, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteExperience(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.EMPLOYEE_EXPERIENCES+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getEmployeeExperiences(employeeID int) ([]*structs.Experience, error) {
	res := &dto.GetEmployeeExperienceListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(employeeID)+"/experiences", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func createEmployeeFamilyMember(familyMember *structs.Family) (*structs.Family, error) {
	res := &dto.GetEmployeeFamilyMemberResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.EMPLOYEE_FAMILY_MEMBERS, familyMember, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateEmployeeFamilyMember(id int, education *structs.Family) (*structs.Family, error) {
	res := &dto.GetEmployeeFamilyMemberResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.EMPLOYEE_FAMILY_MEMBERS+"/"+strconv.Itoa(id), education, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteEmployeeFamilyMember(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.EMPLOYEE_FAMILY_MEMBERS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getEmployeeFamilyMembers(employeeID int) ([]*structs.Family, error) {
	res := &dto.GetEmployeeFamilyMemberListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(employeeID)+"/family-members", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getLoggedInUserProfile(token string) (*structs.UserProfiles, error) {
	userAccount, err := getLoggedInUser(token)
	if err != nil {
		return nil, err
	}

	userProfile, err := getUserProfileByUserAccountID(userAccount.Id)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}
