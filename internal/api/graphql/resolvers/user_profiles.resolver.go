package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strings"
	"time"

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
	jobPositionID := params.Args["job_position_id"]
	isActive, isActiveOk := params.Args["is_active"].(bool)
	name, nameOk := params.Args["name"].(string)

	if id != nil && id != 0 {
		user, err := r.Repo.GetUserProfileByID(id.(int))
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

		loggedInAccount := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
		userOrganizationUnitID, _ := params.Context.Value(config.OrganizationUnitIDKey).(*int)

		for _, userProfile := range profiles {
			resItem, err := buildUserProfileOverviewResponse(r.Repo, userProfile)
			if err != nil {
				return errors.HandleAPIError(err)
			}

			if isActiveOk &&
				resItem.Active != isActive {
				continue
			}
			if nameOk && name != "" && !strings.Contains(strings.ToLower(resItem.FirstName), strings.ToLower(name)) && !strings.Contains(strings.ToLower(resItem.LastName), strings.ToLower(name)) {
				continue
			}

			if loggedInAccount.RoleID != structs.UserRoleAdmin && resItem.OrganizationUnit.ID != *userOrganizationUnitID {
				continue
			}

			if organizationUnitID != nil && organizationUnitID.(int) > 0 &&
				resItem.OrganizationUnit.ID != organizationUnitID {
				continue
			}
			if jobPositionID != nil && jobPositionID.(int) > 0 &&
				resItem.JobPosition.ID != jobPositionID {
				continue
			}

			items = append(items, *resItem)
		}
		total = len(items)
	}

	paginatedItems, err := shared.Paginate(items, page, size)
	if err != nil {
		return errors.HandleAPIError(err)
	}

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
		departmentDropdown       dto.DropdownSimple
		jobPositionDropdown      structs.SettingsDropdown
		isJudge, isPresident     bool
	)
	account, err := r.GetUserAccountByID(profile.UserAccountID)
	if err != nil {
		return nil, err
	}

	role, err := r.GetRole(account.RoleID)
	if err != nil {
		return nil, err
	}

	employeesInOrganizationUnit, _ := r.GetEmployeesInOrganizationUnitsByProfileID(profile.ID)

	if employeesInOrganizationUnit != nil {
		jobPositionInOrganizationUnit, err := r.GetJobPositionsInOrganizationUnitsByID(employeesInOrganizationUnit.PositionInOrganizationUnitID)
		if err != nil {
			return nil, err
		}

		jobPosition, err := r.GetJobPositionByID(jobPositionInOrganizationUnit.JobPositionID)
		if err != nil {
			return nil, err
		}
		jobPositionDropdown.ID = jobPosition.ID
		jobPositionDropdown.Title = jobPosition.Title

		systematization, _ := r.GetSystematizationByID(jobPositionInOrganizationUnit.SystematizationID)

		organizationUnit, err := r.GetOrganizationUnitByID(systematization.OrganizationUnitID)
		if err != nil {
			return nil, err
		}
		organizationUnitDropdown.ID = organizationUnit.ID
		organizationUnitDropdown.Title = organizationUnit.Title
	}

	contract, err := r.GetEmployeeContracts(profile.ID, nil)

	if err != nil {
		return nil, err
	}

	if len(contract) > 0 {
		orgUnit, err := r.GetOrganizationUnitByID(contract[0].OrganizationUnitID)
		if err != nil {
			return nil, err
		}
		organizationUnitDropdown.ID = orgUnit.ID
		organizationUnitDropdown.Title = orgUnit.Title

		if contract[0].OrganizationUnitDepartmentID != nil {
			department, err := r.GetOrganizationUnitByID(*contract[0].OrganizationUnitDepartmentID)
			if err != nil {
				return nil, err
			}
			departmentDropdown.ID = department.ID
			departmentDropdown.Title = department.Title
		}
	}

	active := true
	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}
	resolution, _ := r.GetJudgeResolutionList(&input)

	if len(resolution.Data) > 0 {

		judgeResolutionOrganizationUnit, _, err := r.GetJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
			UserProfileID: &profile.ID,
			ResolutionID:  &resolution.Data[0].ID,
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
			UserProfileID: &profile.ID,
		}
		judge, _, err := r.GetJudgeResolutionOrganizationUnit(&filter)

		if err != nil {
			return nil, err
		}

		orgUnit, err := r.GetOrganizationUnitByID(judge[0].OrganizationUnitID)

		if err != nil {
			return nil, err
		}

		organizationUnitDropdown.ID = orgUnit.ID
		organizationUnitDropdown.Title = orgUnit.Title
	}

	return &dto.UserProfileOverviewResponse{
		ID:          profile.ID,
		FirstName:   profile.FirstName,
		LastName:    profile.LastName,
		DateOfBirth: profile.DateOfBirth,
		Email:       account.Email,
		Phone:       account.Phone,
		Active:      account.Active,
		IsJudge:     isJudge,
		IsPresident: isPresident,
		Role: structs.SettingsDropdown{
			ID:    role.ID,
			Title: role.Title,
		},
		OrganizationUnit: organizationUnitDropdown,
		Department:       departmentDropdown,
		JobPosition:      jobPositionDropdown,
		CreatedAt:        profile.CreatedAt,
		UpdatedAt:        profile.UpdatedAt,
	}, nil
}

func (r *Resolver) UserProfileBasicResolver(params graphql.ResolveParams) (interface{}, error) {
	profileID := params.Args["user_profile_id"]

	profile, err := r.Repo.GetUserProfileByID(profileID.(int))
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

	userAccountData.ID = userProfileData.UserAccountID

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

	userProfileData.UserAccountID = userAccountRes.ID
	userProfileRes, err = r.Repo.CreateUserProfile(userProfileData)
	if err != nil {
		_ = r.Repo.DeleteUserAccount(userAccountRes.ID)
		return errors.HandleAPIError(err)
	}

	if activeContract.Contract != nil {
		activeContract.Contract.UserProfileID = userProfileRes.ID
		_, err := r.Repo.CreateEmployeeContract(activeContract.Contract)
		if err != nil {
			_ = r.Repo.DeleteUserAccount(userAccountRes.ID)
			_ = r.Repo.DeleteUserProfile(userProfileRes.ID)
			return errors.HandleAPIError(err)
		}

		if activeContract.Contract.JobPositionInOrganizationUnitID > 0 {
			input := &structs.EmployeesInOrganizationUnits{
				PositionInOrganizationUnitID: activeContract.Contract.JobPositionInOrganizationUnitID,
				UserProfileID:                userProfileRes.ID,
			}
			_, err = r.Repo.CreateEmployeesInOrganizationUnits(input)
			if err != nil {
				_ = r.Repo.DeleteUserAccount(userAccountRes.ID)
				_ = r.Repo.DeleteUserProfile(userProfileRes.ID)
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
				UserProfileID:      userProfileRes.ID,
				OrganizationUnitID: activeContract.Contract.OrganizationUnitID,
				ResolutionID:       resolution.Data[0].ID,
				IsPresident:        userProfileData.IsPresident,
			}
			_, err := r.Repo.CreateJudgeResolutionOrganizationUnit(&inputCreate)
			if err != nil {
				_ = r.Repo.DeleteUserAccount(userAccountRes.ID)
				_ = r.Repo.DeleteUserProfile(userProfileRes.ID)
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
		activeContract.Contract.UserProfileID = userProfileData.ID
		if activeContract.Contract.ID == 0 {
			_, err = r.Repo.CreateEmployeeContract(activeContract.Contract)
			if err != nil {
				return errors.HandleAPIError(err)
			}
		} else {
			_, err = r.Repo.UpdateEmployeeContract(activeContract.Contract.ID, activeContract.Contract)
			if err != nil {
				return errors.HandleAPIError(err)
			}
		}
		if activeContract.Contract.JobPositionInOrganizationUnitID > -1 {
			myBool := 2
			check := true
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
						SystematizationID: &systematization.ID,
					}
					jobPositionsInOrganizationUnits, err := r.Repo.GetJobPositionsInOrganizationUnits(&inputJpbPos)
					if err != nil {
						return errors.HandleAPIError(err)
					}
					if len(jobPositionsInOrganizationUnits.Data) > 0 {
						for _, job := range jobPositionsInOrganizationUnits.Data {

							input := dto.GetEmployeesInOrganizationUnitInput{
								PositionInOrganizationUnit: &job.ID,
								UserProfileID:              &userProfileData.ID,
							}
							employeesInOrganizationUnit, _ := r.Repo.GetEmployeesInOrganizationUnitList(&input)
							if len(employeesInOrganizationUnit) > 0 {
								for _, emp := range employeesInOrganizationUnit {
									if emp.UserProfileID == userProfileData.ID {
										if emp.PositionInOrganizationUnitID == activeContract.Contract.JobPositionInOrganizationUnitID {
											check = false
										} else {
											err := r.Repo.DeleteEmployeeInOrganizationUnit(emp.ID)
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
					PositionInOrganizationUnitID: activeContract.Contract.JobPositionInOrganizationUnitID,
					UserProfileID:                userProfileData.ID,
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
			OrganizationUnitID: &activeContract.Contract.OrganizationUnitID,
			UserProfileID:      &userProfileData.ID,
			ResolutionID:       &resolution.Data[0].ID,
		})

		if len(judgeResolutionOrganizationUnit) > 0 {
			if userProfileData.IsJudge {
				inputUpdate := dto.JudgeResolutionsOrganizationUnitItem{
					ID:                 judgeResolutionOrganizationUnit[0].ID,
					UserProfileID:      userProfileData.ID,
					OrganizationUnitID: activeContract.Contract.OrganizationUnitID,
					ResolutionID:       resolution.Data[0].ID,
					IsPresident:        userProfileData.IsPresident,
				}
				_, err := r.Repo.UpdateJudgeResolutionOrganizationUnit(&inputUpdate)
				if err != nil {
					return errors.HandleAPIError(err)
				}
			} else {
				err := r.Repo.DeleteJJudgeResolutionOrganizationUnit(judgeResolutionOrganizationUnit[0].ID)
				if err != nil {
					return errors.HandleAPIError(err)
				}
			}
		}
		if len(judgeResolutionOrganizationUnit) == 0 && userProfileData.IsJudge {
			inputCreate := dto.JudgeResolutionsOrganizationUnitItem{
				UserProfileID:      userProfileData.ID,
				OrganizationUnitID: activeContract.Contract.OrganizationUnitID,
				ResolutionID:       resolution.Data[0].ID,
				IsPresident:        userProfileData.IsPresident,
			}
			_, err := r.Repo.CreateJudgeResolutionOrganizationUnit(&inputCreate)
			if err != nil {
				return errors.HandleAPIError(err)
			}
		}

	}

	userProfileRes, err := r.Repo.UpdateUserProfile(userProfileData.ID, userProfileData)
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

	itemID := data.ID
	if itemID != 0 {
		item, err := r.Repo.UpdateEmployeeContract(itemID, &data)
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
	itemID := params.Args["id"]

	err := r.Repo.DeleteEmployeeContract(itemID.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildEducationResItem(r repository.MicroserviceRepositoryInterface, education structs.Education) (*dto.Education, error) {
	educationType, err := r.GetDropdownSettingByID(education.TypeID)
	if err != nil {
		return nil, err
	}

	educationResItem := &dto.Education{
		ID:                  education.ID,
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
		UserProfileID:       education.UserProfileID,
		ExpertiseLevel:      education.ExpertiseLevel,
	}

	educationResItem.Type = dto.DropdownSimple{ID: educationType.ID, Title: educationType.Title}
	if education.FileID != 0 {
		file, err := r.GetFileByID(education.FileID)
		if err != nil {
			return nil, err
		}
		educationResItem.File = dto.FileDropdownSimple{
			ID:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		}
	}

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
			TypeID:        &educationType.ID,
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

	itemID := data.ID
	if itemID != 0 {
		employeeEducation, err = r.Repo.UpdateEmployeeEducation(itemID, &data)
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
	itemID := params.Args["id"]

	err := r.Repo.DeleteEmployeeEducation(itemID.(int))
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
	experienceResponseItemList, err := buildExprienceResponseItemList(r.Repo, experiences)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   experienceResponseItemList,
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

	itemID := data.ID
	if itemID != 0 {
		item, err := r.Repo.UpdateExperience(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resItem, err := buildExprienceResponseItem(r.Repo, item)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		item, err := r.Repo.CreateExperience(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resItem, err := buildExprienceResponseItem(r.Repo, item)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		response.Item = resItem
	}

	return response, nil
}

func (r *Resolver) UserProfileExperiencesInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var data []structs.Experience

	response := dto.Response{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return errors.ErrorResponse("Error updating experience data"), nil
	}

	var responseItems []*dto.ExperienceResponseItem
	for _, item := range data {
		item, err := r.Repo.CreateExperience(&item)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resItem, err := buildExprienceResponseItem(r.Repo, item)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		responseItems = append(responseItems, resItem)
	}

	response.Items = responseItems

	return response, nil
}

func buildExprienceResponseItemList(repo repository.MicroserviceRepositoryInterface, items []*structs.Experience) (resItemList []*dto.ExperienceResponseItem, err error) {
	for _, item := range items {
		resItem, err := buildExprienceResponseItem(repo, item)
		if err != nil {
			return nil, err
		}
		resItemList = append(resItemList, resItem)
	}
	return
}

func buildExprienceResponseItem(repo repository.MicroserviceRepositoryInterface, item *structs.Experience) (*dto.ExperienceResponseItem, error) {
	var fileDropdown dto.FileDropdownSimple

	if item.ReferenceFileID != 0 {
		file, err := repo.GetFileByID(item.ReferenceFileID)

		if err != nil {
			return nil, err
		}

		fileDropdown.ID = file.ID
		fileDropdown.Name = file.Name

		if file.Type != nil {
			fileDropdown.Type = *file.Type
		}
	}

	dateOfEnd, _ := time.Parse("2006-01-02T00:00:00Z", item.DateOfEnd)
	dateOfStart, _ := time.Parse("2006-01-02T00:00:00Z", item.DateOfStart)

	insuredExperienceYears := item.YearsOfInsuredExperience
	insuredExperienceMonths := item.MonthsOfInsuredExperience
	insuredExperienceDays := item.DaysOfInsuredExperience
	var years, months, days int
	if insuredExperienceYears == 0 && insuredExperienceDays == 0 && insuredExperienceMonths == 0 {
		years = dateOfEnd.Year() - dateOfStart.Year()
		month := dateOfEnd.Month() - dateOfStart.Month()
		if month < 0 {
			month = 12 + dateOfEnd.Month() - dateOfStart.Month()
			years--
		}

		days = dateOfEnd.Day() - dateOfStart.Day()

		if days < 0 {
			days = 30 - dateOfEnd.Day() - dateOfStart.Day()
			month--
			if month < 0 {
				month = 12 + month
				years--
			}
		}
		months = int(month)
	}

	res := dto.ExperienceResponseItem{
		ID:                        item.ID,
		UserProfileID:             item.UserProfileID,
		OrganizationUnitID:        item.OrganizationUnitID,
		Relevant:                  item.Relevant,
		OrganizationUnit:          item.OrganizationUnit,
		YearsOfExperience:         years,
		YearsOfInsuredExperience:  insuredExperienceYears,
		MonthsOfExperience:        months,
		MonthsOfInsuredExperience: insuredExperienceMonths,
		DaysOfExperience:          days,
		DaysOfInsuredExperience:   insuredExperienceDays,
		DateOfStart:               item.DateOfStart,
		DateOfEnd:                 item.DateOfEnd,
		ReferenceFileID:           item.ReferenceFileID,
		CreatedAt:                 item.CreatedAt,
		UpdatedAt:                 item.UpdatedAt,
		File:                      fileDropdown,
	}

	return &res, nil
}

func (r *Resolver) UserProfileExperienceDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]

	err := r.Repo.DeleteExperience(itemID.(int))
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

	itemID := data.ID
	if itemID != 0 {
		res, err := r.Repo.UpdateEmployeeFamilyMember(itemID, &data)
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
	itemID := params.Args["id"]

	err := r.Repo.DeleteEmployeeFamilyMember(itemID.(int))
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
	var file dto.FileDropdownSimple

	if contract.FileID != nil && *contract.FileID != 0 {
		res, err := r.GetFileByID(*contract.FileID)

		if err != nil {
			return nil, err
		}

		file.ID = res.ID
		file.Name = res.Name
		file.Type = *res.Type
	}
	responseContract := &dto.Contract{
		ID:                 contract.ID,
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
		FileID:             contract.FileID,
		File:               file,
	}

	contractType, err := r.GetDropdownSettingByID(contract.ContractTypeID)
	if err != nil {
		return nil, err
	}
	responseContract.ContractType = dto.DropdownSimple{ID: contractType.ID, Title: contractType.Title}

	userProfile, err := r.GetUserProfileByID(contract.UserProfileID)
	if err != nil {
		return nil, err
	}
	responseContract.UserProfile = dto.DropdownSimple{ID: userProfile.ID, Title: userProfile.GetFullName()}

	organizationUnit, err := r.GetOrganizationUnitByID(contract.OrganizationUnitID)
	if err != nil {
		return nil, err
	}
	responseContract.OrganizationUnit = dto.DropdownSimple{ID: organizationUnit.ID, Title: organizationUnit.Title}

	if contract.Active {
		if contract.OrganizationUnitDepartmentID != nil {
			department, err := r.GetOrganizationUnitByID(*contract.OrganizationUnitDepartmentID)
			if err != nil {
				return nil, err
			}
			responseContract.Department = &dto.DropdownSimple{ID: department.ID, Title: department.Title}
		}

		myBool := 2
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
					SystematizationID: &systematization.ID,
				}
				jobPositionsInOrganizationUnits, err := r.GetJobPositionsInOrganizationUnits(&inputJpbPos)
				if err != nil {
					return nil, err
				}
				if len(jobPositionsInOrganizationUnits.Data) > 0 {
					for _, job := range jobPositionsInOrganizationUnits.Data {
						input := dto.GetEmployeesInOrganizationUnitInput{
							PositionInOrganizationUnit: &job.ID,
							UserProfileID:              &userProfile.ID,
						}
						employeesInOrganizationUnit, _ := r.GetEmployeesInOrganizationUnitList(&input)
						if len(employeesInOrganizationUnit) > 0 && employeesInOrganizationUnit[0].UserProfileID == userProfile.ID {
							jobPositionInOU = job
						}
					}
				}
			}
		}
		if jobPositionInOU.JobPositionID > 0 {
			jobPosition, err := r.GetJobPositionByID(jobPositionInOU.JobPositionID)
			if err != nil {
				return nil, err
			}
			responseContract.JobPositionInOrganizationUnit = dto.DropdownSimple{ID: jobPositionInOU.ID, Title: jobPosition.Title}
		}
	}

	return responseContract, nil
}

func buildUserProfileBasicResponse(
	r repository.MicroserviceRepositoryInterface,
	profile *structs.UserProfiles,
) (*dto.UserProfileBasicResponse, error) {
	account, err := r.GetUserAccountByID(profile.UserAccountID)

	if err != nil {
		return nil, err
	}

	var (
		jobPosition                     *structs.JobPositions
		organizationUnit                *structs.OrganizationUnits
		jobPositionInOrganizationUnitID int
	)

	employeesInOrganizationUnit, _ := r.GetEmployeesInOrganizationUnitsByProfileID(profile.ID)
	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok && apiErr.StatusCode != 404 {
			return nil, err
		}
	}

	if employeesInOrganizationUnit != nil {
		jobPositionInOrganizationUnit, err := r.GetJobPositionsInOrganizationUnitsByID(employeesInOrganizationUnit.PositionInOrganizationUnitID)
		if err != nil {
			return nil, err
		}
		jobPositionInOrganizationUnitID = jobPositionInOrganizationUnit.ID

		jobPosition, err = r.GetJobPositionByID(jobPositionInOrganizationUnit.JobPositionID)
		if err != nil {
			return nil, err
		}

		organizationUnit, err = r.GetOrganizationUnitByID(jobPositionInOrganizationUnit.ParentOrganizationUnitID)
		if err != nil {
			return nil, err
		}
	}

	userProfileResItem := &dto.UserProfileBasicResponse{
		ID:                             profile.ID,
		FirstName:                      profile.FirstName,
		LastName:                       profile.LastName,
		DateOfBirth:                    profile.DateOfBirth,
		BirthLastName:                  profile.BirthLastName,
		CountryOfBirth:                 profile.CountryOfBirth,
		CityOfBirth:                    profile.CityOfBirth,
		Nationality:                    profile.Nationality,
		Citizenship:                    profile.Citizenship,
		Address:                        profile.Address,
		FatherName:                     profile.FatherName,
		MotherName:                     profile.MotherName,
		MotherBirthLastName:            profile.MotherBirthLastName,
		BankAccount:                    profile.BankAccount,
		BankName:                       profile.BankName,
		PersonalID:                     profile.PersonalID,
		OfficialPersonalID:             profile.OfficialPersonalID,
		OfficialPersonalDocNumber:      profile.OfficialPersonalDocumentNumber,
		OfficialPersonalDocIssuer:      profile.OfficialPersonalDocumentIssuer,
		Gender:                         profile.Gender,
		SingleParent:                   profile.SingleParent,
		HousingDone:                    profile.HousingDone,
		HousingDescription:             profile.HousingDescription,
		IsPresident:                    false,
		IsJudge:                        false,
		MaritalStatus:                  profile.MaritalStatus,
		DateOfTakingOath:               profile.DateOfTakingOath,
		DateOfBecomingJudge:            profile.DateOfBecomingJudge,
		JudgeApplicationSubmissionDate: profile.JudgeApplicationSubmissionDate,
		Email:                          account.Email,
		Phone:                          account.Phone,
		OrganizationUnit:               organizationUnit,
		JobPosition:                    jobPosition,
		JobPositionInOrganizationUnit:  jobPositionInOrganizationUnitID,
		NationalMinority:               profile.NationalMinority,
	}
	active := true
	contracts, err := r.GetEmployeeContracts(profile.ID, nil)
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
		if contractResponseItem.OrganizationUnit.ID > 0 {

			input := dto.GetJudgeResolutionListInputMS{
				Active: &active,
			}

			resolution, _ := r.GetJudgeResolutionList(&input)

			if len(resolution.Data) > 0 {
				resolutionID := resolution.Data[0].ID
				judgeResolutionOrganizationUnit, _, err := r.GetJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
					ResolutionID:  &resolutionID,
					UserProfileID: &profile.ID,
				})
				if err != nil {
					return nil, err
				}

				if len(judgeResolutionOrganizationUnit) > 0 {
					for _, judge := range judgeResolutionOrganizationUnit {
						userProfileResItem.IsPresident = judge.IsPresident
						userProfileResItem.IsJudge = true
					}
					organizationUnitID, err := r.GetOrganizationUnitByID(judgeResolutionOrganizationUnit[0].OrganizationUnitID)
					if err != nil {
						return nil, err
					}

					userProfileResItem.OrganizationUnit = organizationUnitID
					userProfileResItem.Contract.OrganizationUnit = dto.DropdownSimple{
						ID:    organizationUnitID.ID,
						Title: organizationUnitID.Title,
					}
				}
			}
		}

	}

	evaluations, err := r.GetEmployeeEvaluations(userProfileResItem.ID)

	if err != nil {
		return nil, err
	}

	if len(evaluations) > 0 {
		evaluation, err := r.GetDropdownSettingByID(evaluations[0].EvaluationTypeID)
		if err != nil {
			return nil, err
		}
		userProfileResItem.Evaluation.ID = evaluation.ID
		userProfileResItem.Evaluation.Title = evaluation.Title
	}

	return userProfileResItem, nil
}
