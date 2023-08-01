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
			return dto.Response{
				Status:  "error",
				Message: err.Error(),
			}, nil
		}
		resItem, err := buildUserProfileOverviewResponse(user)
		if err != nil {
			return dto.Response{
				Status:  "error",
				Message: err.Error(),
			}, nil
		}
		items = []dto.UserProfileOverviewResponse{*resItem}
		total = 1
	} else {
		input := dto.GetUserProfilesInput{}
		profiles, err := getUserProfiles(&input)
		if err != nil {
			return dto.Response{
				Status:  "error",
				Message: err.Error(),
			}, nil
		}

		total = len(profiles)
		for _, userProfile := range profiles {
			resItem, err := buildUserProfileOverviewResponse(userProfile)
			if err != nil {
				return dto.Response{
					Status:  "error",
					Message: err.Error(),
				}, nil
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

	paginatedItems, err := shared.Paginate(items, page, size)
	if err != nil {
		fmt.Printf("Error paginating items: %v", err)
	}
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
		return dto.ErrorResponse(err), nil
	}
	contractResponseItems, err := buildContractResponseItemList(contracts)
	if err != nil {
		return dto.ErrorResponse(err), nil
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
		isJudge, isJudgePresdent bool
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
		isJudge = jobPosition.IsJudge
		isJudgePresdent = jobPosition.IsJudgePresident

		organizationUnit, err := getOrganizationUnitById(jobPositionInOrganizationUnit.ParentOrganizationUnitId)
		if err != nil {
			return nil, err
		}
		organizationUnitDropdown.Id = organizationUnit.Id
		organizationUnitDropdown.Title = organizationUnit.Title
	}

	return &dto.UserProfileOverviewResponse{
		ID:               profile.Id,
		FirstName:        profile.FirstName,
		LastName:         profile.LastName,
		DateOfBirth:      profile.DateOfBirth,
		Email:            account.Email,
		Phone:            account.Phone,
		Active:           account.Active,
		IsJudge:          isJudge,
		IsJudgePresident: isJudgePresdent,
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

	if !shared.IsInteger(profileId) {
		return map[string]interface{}{
			"status":  "error",
			"message": "Argument 'user_profile_id' must not be empty!",
			"item":    nil,
		}, nil
	}

	profile, err := getUserProfileById(profileId.(int))
	if err != nil {
		fmt.Printf("Error getting user profile because of this error - %s.\n", err)
		return shared.ErrorResponse("Error getting user profile data"), nil
	}

	res, err := buildUserProfileBasicResponse(profile)
	if err != nil {
		fmt.Printf("Building user response failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error building user data"), nil
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
	var employeesInOrganizationUnits structs.EmployeesInOrganizationUnits
	var employeeContracts dto.CreateUserProfileContractList

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
	err = json.Unmarshal(dataBytes, &employeesInOrganizationUnits)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating settings data"), nil
	}
	err = json.Unmarshal(dataBytes, &employeeContracts)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating settings data"), nil
	}

	userAccountRes, err = CreateUserAccount(userAccountData)
	if err != nil {
		fmt.Printf("Creating the user account failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error creating the user account data"), nil
	}

	userProfileData.UserAccountId = userAccountRes.Id
	userProfileRes, err = createUserProfile(userProfileData)
	if err != nil {
		fmt.Printf("Creating the user profile failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error creating the user profile data"), nil
	}

	for _, contractInput := range employeeContracts.Contracts {
		contractInput.UserProfileId = userProfileRes.Id
		_, err := createEmployeeContract(contractInput)
		if err != nil {
			fmt.Printf("Creating the user profile contracts failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating the user profile contracts data"), nil
		}
	}

	if employeesInOrganizationUnits.PositionInOrganizationUnitId != 0 {
		employeesInOrganizationUnits.UserAccountId = userAccountRes.Id
		employeesInOrganizationUnits.UserProfileId = userProfileRes.Id
		employeesInOrganizationUnits.Active = true
		_, err = createEmployeesInOrganizationUnits(&employeesInOrganizationUnits)
		if err != nil {
			fmt.Printf("Inserting employees in organization unit failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating the employees in organization unit data"), nil
		}
	}

	res, err := buildUserProfileBasicResponse(userProfileRes)
	if err != nil {
		fmt.Printf("Building user response failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error getting user data"), nil
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
	var positionInOrganizationUnitData structs.EmployeesInOrganizationUnits

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &userProfileData)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating settings data"), nil
	}

	userProfileRes, err := updateUserProfile(userProfileData.Id, userProfileData)
	if err != nil {
		fmt.Printf("Creating the user profile failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error creating the user profile data"), nil
	}

	err = json.Unmarshal(dataBytes, &positionInOrganizationUnitData)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating settings data"), nil
	}

	if positionInOrganizationUnitData.PositionInOrganizationUnitId != 0 {
		positionInOrganizationUnitData.UserProfileId = userProfileRes.Id
		_, err := updateEmployeePositionInOrganizationUnitByProfile(userProfileData.Id, &positionInOrganizationUnitData)
		if err != nil {
			fmt.Printf("Inserting employees in organization unit failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating the employees in organization unit data"), nil
		}
	}

	res, err := buildUserProfileBasicResponse(userProfileRes)
	if err != nil {
		fmt.Printf("Building user response failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error getting user data"), nil
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
			fmt.Printf("Updating organization type failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error updating user profile contract data"), nil
		}
		contractResponseItem, err := buildContractResponseItem(*item)
		if err != nil {
			return dto.ErrorResponse(err), nil
		}
		response.Message = "You updated this item!"
		response.Item = contractResponseItem
	} else {
		item, err := createEmployeeContract(&data)
		if err != nil {
			fmt.Printf("Creating organization type failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating user profile contract data"), nil
		}
		contractResponseItem, err := buildContractResponseItem(*item)
		if err != nil {
			return dto.ErrorResponse(err), nil
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
		fmt.Printf("Deleting employee's contract failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error deleting the contract"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

var UserProfileEducationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		response []dto.EducationTypeWithEducationsResponse
	)

	userProfileID := params.Args["user_profile_id"]

	settingsInput := dto.GetSettingsInput{
		Entity: "education_types",
	}
	educationTypes, err := getDropdownSettings(&settingsInput)
	if err != nil {
		return dto.Response{
			Status:  "error",
			Message: err.Error(),
		}, nil
	}

	educationTypeMap := make(map[int]dto.EducationTypeWithEducationsResponse)

	for _, educationType := range educationTypes.Data {
		educationTypeResponse := dto.EducationTypeWithEducationsResponse{
			ID:           educationType.Id,
			Abbreviation: educationType.Abbreviation,
			Title:        educationType.Title,
			Value:        educationType.Value,
		}
		educationTypeMap[educationType.Id] = educationTypeResponse
	}

	userProfileEducations, err := getEmployeeEducations(userProfileID.(int))
	if err != nil {
		return dto.Response{
			Status:  "error",
			Message: err.Error(),
		}, nil
	}

	for _, education := range userProfileEducations {
		if educationTypeResponse, ok := educationTypeMap[education.EducationTypeId]; ok {
			educationTypeResponse.Educations = append(educationTypeResponse.Educations, education)
			educationTypeMap[education.EducationTypeId] = educationTypeResponse
		}
	}

	for _, educationTypeResponse := range educationTypeMap {
		response = append(response, educationTypeResponse)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

var UserProfileEducationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Education
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		employeeEducationResponse, err := updateEmployeeEducation(itemId, &data)
		if err != nil {
			fmt.Printf("Updating employee's education failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error updating employee's education data"), nil
		}
		response.Item = employeeEducationResponse
		response.Message = "You updated this item!"
	} else {
		employeeEducationResponse, err := createEmployeeEducation(&data)
		if err != nil {
			fmt.Printf("Creating employee's education failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating employee's education data"), nil
		}
		response.Item = employeeEducationResponse
		response.Message = "You created this item!"
	}

	return response, nil
}

var UserProfileEducationDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := deleteEmployeeEducation(itemId.(int))
	if err != nil {
		fmt.Printf("Deleting employee's education failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error deleting the id"), nil
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
		return dto.Response{
			Status:  "error",
			Message: err.Error(),
		}, nil
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
			fmt.Printf("Updating experience failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error updating experience data"), nil
		}
		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := createExperience(&data)
		if err != nil {
			fmt.Printf("Creating experience failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating experience data"), nil
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
		fmt.Printf("Deleting experience failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error deleting the experience"), nil
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
		return dto.Response{
			Status:  "error",
			Message: err.Error(),
		}, nil
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
			fmt.Printf("Updating employee's family member failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error updating employee's family member data"), nil
		}
		response.Item = res
		response.Message = "You updated this item!"
	} else {
		res, err := createEmployeeFamilyMember(&data)
		if err != nil {
			fmt.Printf("Creating employee's family member failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating employee's family member data"), nil
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
		fmt.Printf("Deleting Family member failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error deleting the id"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildContractResponseItemList(contracts []*structs.Contracts) (contractResponseItemList []*structs.Contracts, err error) {
	for _, contract := range contracts {
		contractResItem, err := buildContractResponseItem(*contract)
		if err != nil {
			return nil, err
		}
		contractResponseItemList = append(contractResponseItemList, contractResItem)
	}
	return
}

func buildContractResponseItem(contract structs.Contracts) (*structs.Contracts, error) {
	contractType, err := getDropdownSettingById(contract.ContractTypeId)
	if err != nil {
		return nil, err
	}
	contract.ContractType = *contractType

	return &contract, nil
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

	contracts, err := getEmployeeContracts(profile.Id, nil)
	if err != nil {
		return nil, err
	}
	contractResponseItems, err := buildContractResponseItemList(contracts)
	if err != nil {
		return nil, err
	}

	return &dto.UserProfileBasicResponse{
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
		OfficialPersonalID:            profile.OfficialPersonalId,
		OfficialPersonalDocNumber:     profile.OfficialPersonalDocumentNumber,
		OfficialPersonalDocIssuer:     profile.OfficialPersonalDocumentIssuer,
		Gender:                        profile.Gender,
		SingleParent:                  profile.SingleParent,
		HousingDone:                   profile.HousingDone,
		HousingDescription:            profile.HousingDescription,
		RevisorRole:                   profile.RevisorRole,
		MaritalStatus:                 profile.MaritalStatus,
		DateOfTakingOath:              profile.DateOfTakingOath,
		DateOfBecomingJudge:           profile.DateOfBecomingJudge,
		Email:                         account.Email,
		Phone:                         account.Phone,
		OrganizationUnit:              organizationUnit,
		JobPosition:                   jobPosition,
		Contracts:                     contractResponseItems,
		JobPositionInOrganizationUnit: jobPositionInOrganizationUnitID,
	}, nil
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

func updateEmployeePositionInOrganizationUnitByProfile(profileId int, data *structs.EmployeesInOrganizationUnits) (*structs.EmployeesInOrganizationUnits, error) {
	res := &dto.GetEmployeesInOrganizationUnitsResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.EMPLOYEES_IN_ORGANIZATION_UNITS_ENDPOINT+"/"+strconv.Itoa(profileId), data, res)
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

	return &res.Data, nil
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

func getEmployeeEducations(userProfileID int) ([]structs.Education, error) {
	res := &dto.GetEmployeeEducationListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(userProfileID)+"/educations", nil, res)
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
