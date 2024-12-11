package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"crypto/rand"
	"encoding/json"
	"math/big"
	"strconv"
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		resItem, err := buildUserProfileOverviewResponse(r.Repo, user)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		items = []dto.UserProfileOverviewResponse{*resItem}
		total = 1
	} else {
		input := dto.GetUserProfilesInput{}
		profiles, err := r.Repo.GetUserProfiles(&input)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		loggedInAccount := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
		userOrganizationUnitID, _ := params.Context.Value(config.OrganizationUnitIDKey).(*int)

		hasPermission, err := r.HasPermission(*loggedInAccount, string(config.HR), config.OperationFullAccess)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, userProfile := range profiles {
			resItem, err := buildUserProfileOverviewResponse(r.Repo, userProfile)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			if isActiveOk &&
				resItem.Active != isActive {
				continue
			}
			if nameOk && name != "" && !strings.Contains(strings.ToLower(resItem.FirstName), strings.ToLower(name)) && !strings.Contains(strings.ToLower(resItem.LastName), strings.ToLower(name)) {
				continue
			}

			if !hasPermission && resItem.OrganizationUnit.ID != *userOrganizationUnitID {
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   paginatedItems,
		Total:   total,
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
	account, _ := r.GetUserAccountByID(profile.UserAccountID)
	/*	if err != nil {
			return nil, errors.Wrap(err, "repo get user account by id")
		}
	*/

	var role *structs.Roles

	if account != nil {

		if account.RoleID != nil && *account.RoleID != 0 {
			role, _ = r.GetRole(*account.RoleID)
			/*	if err != nil {
				return nil, errors.Wrap(err, "repo get role")
			}*/
		}
	}

	employeesInOrganizationUnit, _ := r.GetEmployeesInOrganizationUnitsByProfileID(profile.ID)

	if employeesInOrganizationUnit != nil && employeesInOrganizationUnit.PositionInOrganizationUnitID != 0 {
		jobPositionInOrganizationUnit, err := r.GetJobPositionsInOrganizationUnitsByID(employeesInOrganizationUnit.PositionInOrganizationUnitID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get job positions in organization units by id")
		}

		jobPosition, err := r.GetJobPositionByID(jobPositionInOrganizationUnit.JobPositionID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get job position by id")
		}

		systematization, err := r.GetSystematizationByID(jobPositionInOrganizationUnit.SystematizationID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get systematization by id")
		}

		if systematization.Active == 2 {

			jobPositionDropdown.ID = jobPosition.ID
			jobPositionDropdown.Title = jobPosition.Title

			organizationUnit, err := r.GetOrganizationUnitByID(systematization.OrganizationUnitID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get organization unit by id")
			}
			organizationUnitDropdown.ID = organizationUnit.ID
			organizationUnitDropdown.Title = organizationUnit.Title

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
			return nil, errors.Wrap(err, "repo get judge resolution organization unit")
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
			return nil, errors.Wrap(err, "repo get judge resolution organization unit")
		}

		orgUnit, err := r.GetOrganizationUnitByID(judge[0].OrganizationUnitID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}

		organizationUnitDropdown.ID = orgUnit.ID
		organizationUnitDropdown.Title = orgUnit.Title
	}

	if role == nil {
		role = &structs.Roles{
			ID:    0,
			Title: "",
		}
	}

	if organizationUnitDropdown.ID == 0 {
		contracts, err := r.GetEmployeeContracts(profile.ID, nil)

		if err != nil {
			return nil, errors.Wrap(err, "repo get employee contracts")
		}

		if len(contracts) > 0 {
			organizationUnit, err := r.GetOrganizationUnitByID(contracts[0].OrganizationUnitID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get organization unit by id")
			}
			organizationUnitDropdown.ID = organizationUnit.ID
			organizationUnitDropdown.Title = organizationUnit.Title
		}
	}

	return &dto.UserProfileOverviewResponse{
		ID:               profile.ID,
		FirstName:        profile.FirstName,
		LastName:         profile.LastName,
		DateOfBirth:      profile.DateOfBirth,
		Email:            account.Email,
		Phone:            account.Phone,
		Active:           account.Active,
		IsJudge:          isJudge,
		IsPresident:      isPresident,
		IsJudgePresident: isPresident,
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	res, err := buildUserProfileBasicResponse(r.Repo, profile)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	err = json.Unmarshal(dataBytes, &userProfileData)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	err = json.Unmarshal(dataBytes, &activeContract)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	active := true
	inactive := false

	if activeContract.Contract != nil {
		userProfileData.ActiveContract = &active
	} else {
		userProfileData.ActiveContract = &inactive
	}

	password, err := generateRandomString(8)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	userAccountData.Password = password

	pin, err := generateRandomString(4)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	userAccountData.Pin = pin

	userAccountRes, err = r.Repo.CreateUserAccount(params.Context, userAccountData)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	userProfileData.UserAccountID = userAccountRes.ID
	userProfileRes, err = r.Repo.CreateUserProfile(params.Context, userProfileData)
	if err != nil {
		_ = r.Repo.DeleteUserAccount(params.Context, userAccountRes.ID)
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if activeContract.Contract != nil {
		activeContract.Contract.UserProfileID = userProfileRes.ID
		_, err := r.Repo.CreateEmployeeContract(params.Context, activeContract.Contract)
		if err != nil {
			_ = r.Repo.DeleteUserAccount(params.Context, userAccountRes.ID)
			_ = r.Repo.DeleteUserProfile(params.Context, userProfileRes.ID)
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if activeContract.Contract.JobPositionInOrganizationUnitID > 0 {
			input := &structs.EmployeesInOrganizationUnits{
				PositionInOrganizationUnitID: activeContract.Contract.JobPositionInOrganizationUnitID,
				UserProfileID:                userProfileRes.ID,
			}
			_, err = r.Repo.CreateEmployeesInOrganizationUnits(input)
			if err != nil {
				_ = r.Repo.DeleteUserAccount(params.Context, userAccountRes.ID)
				_ = r.Repo.DeleteUserProfile(params.Context, userProfileRes.ID)
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
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
				_ = r.Repo.DeleteUserAccount(params.Context, userAccountRes.ID)
				_ = r.Repo.DeleteUserProfile(params.Context, userProfileRes.ID)
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}
	}

	res, err := buildUserProfileBasicResponse(r.Repo, userProfileRes)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	err = json.Unmarshal(dataBytes, &activeContract)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	oldData, err := r.Repo.GetUserProfileByID(userProfileData.ID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	active := true
	inactive := false
	hasContract := false
	if activeContract.Contract != nil {
		hasContract = true
		userProfileData.ActiveContract = &active
		activeContract.Contract.UserProfileID = userProfileData.ID
		if activeContract.Contract.ID == 0 {
			_, err = r.Repo.CreateEmployeeContract(params.Context, activeContract.Contract)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		} else {
			_, err = r.Repo.UpdateEmployeeContract(params.Context, activeContract.Contract.ID, activeContract.Contract)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
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
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			if len(systematizationsResponse.Data) > 0 {
				for _, systematization := range systematizationsResponse.Data {
					inputJpbPos := dto.GetJobPositionInOrganizationUnitsInput{
						SystematizationID: &systematization.ID,
					}
					jobPositionsInOrganizationUnits, err := r.Repo.GetJobPositionsInOrganizationUnits(&inputJpbPos)
					if err != nil {
						_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
						return errors.HandleAPPError(err)
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
												_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
												return errors.HandleAPPError(err)
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
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
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
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			} else {
				err := r.Repo.DeleteJJudgeResolutionOrganizationUnit(judgeResolutionOrganizationUnit[0].ID)
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
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
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}

	}

	userProfileData.ActiveContract = &hasContract
	userProfileRes, err := r.Repo.UpdateUserProfile(params.Context, userProfileData.ID, userProfileData)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	if oldData.IsJudge && !userProfileData.IsJudge {

		if len(resolution.Data) > 0 {
			judgeResolutionOrganizationUnit, _, err := r.Repo.GetJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
				OrganizationUnitID: &activeContract.Contract.OrganizationUnitID,
				UserProfileID:      &userProfileData.ID,
				ResolutionID:       &resolution.Data[0].ID,
			})

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			for _, item := range judgeResolutionOrganizationUnit {
				err := r.Repo.DeleteJJudgeResolutionOrganizationUnit(item.ID)
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}
	}

	if !oldData.IsJudge && userProfileData.IsJudge {
		item, err := r.Repo.GetEmployeesInOrganizationUnitsByProfileID(userProfileData.ID)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if item != nil && item.ID != 0 {
			err := r.Repo.DeleteEmployeeInOrganizationUnitByID(item.ID)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}
	}

	if hasContract {
		account, err := r.Repo.GetUserAccountByID(userProfileRes.UserAccountID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if !account.Active {
			account.Active = true

			_, err = r.Repo.UpdateUserAccount(params.Context, account.ID, *account)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}
	}

	res, err := buildUserProfileBasicResponse(r.Repo, userProfileRes)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    res,
	}, nil
}

func buildEducationResItem(r repository.MicroserviceRepositoryInterface, education structs.Education) (*dto.Education, error) {
	educationType, _ := r.GetDropdownSettingByID(education.TypeID)
	/*if err != nil {
		return nil, errors.Wrap(err, "repo get dropdown setting by id")
	}*/

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

	if educationType != nil {

		educationResItem.Type = dto.DropdownSimple{ID: educationType.ID, Title: educationType.Title}
	}

	var fileList []dto.FileDropdownSimple
	for i := range education.FileIDs {
		file, _ := r.GetFileByID(education.FileIDs[i])
		/*if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}*/

		if file != nil {
			fileDropdown := dto.FileDropdownSimple{
				ID:   file.ID,
				Name: file.Name,
			}

			if file.Type != nil {
				fileDropdown.Type = *file.Type
			}

			fileList = append(fileList, fileDropdown)
		}
	}

	educationResItem.Files = fileList

	return educationResItem, nil
}

func buildEducationResItemList(r repository.MicroserviceRepositoryInterface, educations []structs.Education) (educationResItemList []*dto.Education, err error) {
	for _, education := range educations {
		educationResItem, err := buildEducationResItem(r, education)
		if err != nil {
			return nil, errors.Wrap(err, "build education res item")
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, educationType := range educationTypes.Data {
		educations, err := r.Repo.GetEmployeeEducations(dto.EducationInput{
			UserProfileID: userProfileID,
			TypeID:        &educationType.ID,
		})
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		educationResItemList, err := buildEducationResItemList(r.Repo, educations)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Message = "You updated this item!"
	} else {
		employeeEducation, err = r.Repo.CreateEmployeeEducation(&data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Message = "You created this item!"
	}

	responseItem, err := buildEducationResItem(r.Repo, *employeeEducation)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	response.Item = responseItem

	return response, nil
}

func (r *Resolver) UserProfileEducationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]

	err := r.Repo.DeleteEmployeeEducation(itemID.(int))
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	experienceResponseItemList, err := buildExprienceResponseItemList(r.Repo, experiences)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	itemID := data.ID
	if itemID != 0 {
		item, err := r.Repo.UpdateExperience(itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		resItem, err := buildExprienceResponseItem(r.Repo, item)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		item, err := r.Repo.CreateExperience(&data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		resItem, err := buildExprienceResponseItem(r.Repo, item)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var responseItems []*dto.ExperienceResponseItem
	for _, item := range data {
		item, err := r.Repo.CreateExperience(&item)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		resItem, err := buildExprienceResponseItem(r.Repo, item)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
			return nil, errors.Wrap(err, "build exprience response item")
		}
		resItemList = append(resItemList, resItem)
	}
	return
}

func calculateDifference(dateOfStart, dateOfEnd time.Time) (int, int, int) {
	years := dateOfEnd.Year() - dateOfStart.Year()
	months := int(dateOfEnd.Month()) - int(dateOfStart.Month())
	days := 1 + dateOfEnd.Day() - dateOfStart.Day()

	if months < 0 {
		years--
		months += 12
	}

	isLeapStartYear := isLeapYear(dateOfStart.Year())
	isLeapEndYear := isLeapYear(dateOfEnd.Year())

	if days < 0 {
		if int(dateOfStart.Month()) == 12 || int(dateOfStart.Month()) == 10 || int(dateOfStart.Month()) == 8 ||
			int(dateOfStart.Month()) == 7 || int(dateOfStart.Month()) == 5 || int(dateOfStart.Month()) == 3 || int(dateOfStart.Month()) == 1 {
			days += 31
		} else if int(dateOfStart.Month()) == 11 || int(dateOfStart.Month()) == 9 || int(dateOfStart.Month()) == 6 || int(dateOfStart.Month()) == 4 {
			days += 30
		} else if int(dateOfStart.Month()) == 2 && isLeapStartYear {
			days += 29
		} else if int(dateOfStart.Month()) == 2 && !isLeapStartYear {
			days += 28
		}
		months--

		if months < 0 {
			years--
			months += 12
		}
	}

	if (days == 31 && (int(dateOfEnd.Month()) == 12 || int(dateOfEnd.Month()) == 10 || int(dateOfEnd.Month()) == 8 ||
		int(dateOfEnd.Month()) == 7 || int(dateOfEnd.Month()) == 5 || int(dateOfEnd.Month()) == 3 || int(dateOfEnd.Month()) == 1)) ||
		(days == 30 && (int(dateOfEnd.Month()) == 11 || int(dateOfEnd.Month()) == 9 || int(dateOfEnd.Month()) == 6 || int(dateOfEnd.Month()) == 4)) ||
		(days == 29 && int(dateOfEnd.Month()) == 2 && isLeapEndYear) ||
		(days == 28 && int(dateOfEnd.Month()) == 2 && !isLeapEndYear) {
		months++
		days = 0
		if months == 12 {
			months = 0
			years++
		}
	}

	return years, months, days
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

func buildExprienceResponseItem(repo repository.MicroserviceRepositoryInterface, item *structs.Experience) (*dto.ExperienceResponseItem, error) {
	var fileDropdownList []dto.FileDropdownSimple

	for i := range item.FileIDs {
		var fileDropdown dto.FileDropdownSimple

		file, _ := repo.GetFileByID(item.FileIDs[i])

		/*	if err != nil {
				return nil, errors.Wrap(err, "repo get file by id")
			}
		*/

		if file != nil {
			fileDropdown.ID = file.ID
			fileDropdown.Name = file.Name

			if file.Type != nil {
				fileDropdown.Type = *file.Type
			}
		}

		fileDropdownList = append(fileDropdownList, fileDropdown)
	}

	dateOfEnd, _ := time.Parse(config.ISO8601Format, item.DateOfEnd)
	dateOfStart, _ := time.Parse(config.ISO8601Format, item.DateOfStart)

	years, months, days := calculateDifference(dateOfStart, dateOfEnd)

	insuredExperienceYears := item.YearsOfInsuredExperience
	insuredExperienceMonths := item.MonthsOfInsuredExperience
	insuredExperienceDays := item.DaysOfInsuredExperience
	if insuredExperienceYears == 0 && insuredExperienceDays == 0 && insuredExperienceMonths == 0 {
		insuredExperienceYears = years
		insuredExperienceMonths = months
		insuredExperienceDays = days
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
		CreatedAt:                 item.CreatedAt,
		UpdatedAt:                 item.UpdatedAt,
		Files:                     fileDropdownList,
	}

	if item.OrganizationUnitID != 0 {
		organizationUnit, err := repo.GetOrganizationUnitByID(item.OrganizationUnitID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get organziation unit by id")
		}

		res.OrganizationUnitTitle = organizationUnit.Title
	}

	return &res, nil
}

func (r *Resolver) UserProfileExperienceDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]

	err := r.Repo.DeleteExperience(itemID.(int))
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildUserProfileBasicResponse(
	r repository.MicroserviceRepositoryInterface,
	profile *structs.UserProfiles,
) (*dto.UserProfileBasicResponse, error) {
	account, err := r.GetUserAccountByID(profile.UserAccountID)

	if err != nil {
		return nil, errors.Wrap(err, "repo get user account by id")
	}

	var (
		jobPosition                     *structs.JobPositions
		organizationUnit                *structs.OrganizationUnits
		jobPositionInOrganizationUnitID int
	)

	employeesInOrganizationUnit, err := r.GetEmployeesInOrganizationUnitsByProfileID(profile.ID)
	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok && apiErr.StatusCode != 404 {
			return nil, errors.Wrap(err, "repo get employees in organizatin units by profile id")
		}
	}

	if employeesInOrganizationUnit != nil && employeesInOrganizationUnit.PositionInOrganizationUnitID != 0 {
		jobPositionInOrganizationUnit, err := r.GetJobPositionsInOrganizationUnitsByID(employeesInOrganizationUnit.PositionInOrganizationUnitID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get job positions in organization units by id")
		}

		systematization, err := r.GetSystematizationByID(jobPositionInOrganizationUnit.SystematizationID)

		if err != nil {
			return nil, errors.Wrap(err, "repo systematization by id")
		}

		if systematization.Active == 2 {

			jobPositionInOrganizationUnitID = jobPositionInOrganizationUnit.ID

			jobPosition, err = r.GetJobPositionByID(jobPositionInOrganizationUnit.JobPositionID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get job positions by id")
			}

			organizationUnit, err = r.GetOrganizationUnitByID(jobPositionInOrganizationUnit.ParentOrganizationUnitID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get organization unit by id")
			}
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
		return nil, errors.Wrap(err, "repo get employee contracts")
	}

	if len(contracts) > 0 {

		contractResponseItem, err := buildContractResponseItem(r, *contracts[0])
		if err != nil {
			return nil, errors.Wrap(err, "build contract response item")
		}

		if organizationUnit != nil || profile.IsJudge {
			userProfileResItem.Contract = contractResponseItem
		}

		if organizationUnit == nil {
			organizationUnit = &structs.OrganizationUnits{
				ID:    contractResponseItem.OrganizationUnit.ID,
				Title: contractResponseItem.OrganizationUnit.Title,
			}
			userProfileResItem.OrganizationUnit = organizationUnit
		}
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
					return nil, errors.Wrap(err, "repo get judge resoluton organization unit")
				}

				if len(judgeResolutionOrganizationUnit) > 0 {
					for _, judge := range judgeResolutionOrganizationUnit {
						userProfileResItem.IsPresident = judge.IsPresident
						userProfileResItem.IsJudge = true
					}
					organizationUnitID, err := r.GetOrganizationUnitByID(judgeResolutionOrganizationUnit[0].OrganizationUnitID)
					if err != nil {
						return nil, errors.Wrap(err, "repo get organization unit by id")
					}

					userProfileResItem.OrganizationUnit = organizationUnitID
					if userProfileResItem.Contract != nil {
						userProfileResItem.Contract.OrganizationUnit = dto.DropdownSimple{
							ID:    organizationUnitID.ID,
							Title: organizationUnitID.Title,
						}
					} else {
						contract := dto.Contract{
							OrganizationUnit: dto.DropdownSimple{
								ID:    organizationUnitID.ID,
								Title: organizationUnitID.Title,
							},
						}
						userProfileResItem.Contract = &contract
					}
				}
			}
		}

	}

	evaluations, err := r.GetEmployeeEvaluations(userProfileResItem.ID)

	if err != nil {
		return nil, errors.Wrap(err, "repo get employee evaluations")
	}

	if len(evaluations) > 0 {
		evaluation, err := r.GetDropdownSettingByID(evaluations[0].EvaluationTypeID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}
		userProfileResItem.Evaluation.ID = evaluation.ID
		userProfileResItem.Evaluation.Title = evaluation.Title
	}

	return userProfileResItem, nil
}

const (
	smallLetterBytes = "abcdefghijklmnopqrstuvwxyz"
	bigLetterBytes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes      = "0123456789"
	symbolBytes      = "!@#$%&"
	allBytes         = smallLetterBytes + bigLetterBytes + numberBytes + symbolBytes
)

func generateRandomString(n int) (string, error) {

	result := make([]byte, n)

	// Ensure one uppercase letter, one number, and one symbol
	categories := []string{
		smallLetterBytes,
		bigLetterBytes,
		numberBytes,
		symbolBytes,
	}

	// Place mandatory characters in the result
	for i, category := range categories {
		char, err := randomCharFromCategory(category)
		if err != nil {
			return "", err
		}
		result[i] = char
	}

	// Fill the remaining positions with random characters
	for i := len(categories); i < n; i++ {
		char, err := randomCharFromCategory(allBytes)
		if err != nil {
			return "", err
		}
		result[i] = char
	}

	// Shuffle the result to mix mandatory characters
	shuffle(result)

	return string(result), nil
}

func randomCharFromCategory(category string) (byte, error) {
	max := big.NewInt(int64(len(category)))
	randNum, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return category[randNum.Int64()], nil
}

func shuffle(data []byte) {
	for i := range data {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		data[i], data[j.Int64()] = data[j.Int64()], data[i]
	}
}

func (r *Resolver) DataForTemplateResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	dataForTemplate, err := r.buildDataForTemplate(itemID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You get this item!",
		Item:    dataForTemplate,
	}, nil
}

func (r *Resolver) buildDataForTemplate(id int) (*dto.UserDataForTemplate, error) {

	var organizationUnitName string
	var departmentName string
	var jobPositionName string
	var jobPositionRequirments string
	var systematizationNumber string
	var systematizationDate string
	var contractStartDate string
	var contractEndDate string
	var workStartDate string
	var contractDuration string
	var acquiredVacationDays string
	var remainingVacationDays string
	var usedVacationDays string
	var weeklyWorkingHours string
	var vacationStartDate string
	var vacationEndDate string
	var rating string
	var education string
	var yearsOfExperience string
	var monthsOfExperience string
	var daysOfExperience string

	employee, err := r.Repo.GetUserProfileByID(id)

	if err != nil {
		return nil, errors.Wrap(err, "repo get user profile by id")
	}

	currentDate := time.Now()

	currentYear := strconv.Itoa(currentDate.Year())
	currentMonth := strconv.Itoa(int(currentDate.Month()))
	currentDateString := currentDate.Format("02.01.2006")

	experiences, err := r.Repo.GetEmployeeExperiences(id)
	if err != nil {
		return nil, errors.Wrap(err, "repo get employee experiences")
	}
	experienceResponseItemList, err := buildExprienceResponseItemList(r.Repo, experiences)
	if err != nil {
		return nil, errors.Wrap(err, "build exprience response item list")
	}

	var yearsOfExperienceInt int
	var monthsOfExperienceInt int
	var daysOfExperienceInt int

	for _, item := range experienceResponseItemList {
		daysOfExperienceInt += item.DaysOfExperience
		monthsOfExperienceInt += item.MonthsOfExperience
		yearsOfExperienceInt += item.YearsOfExperience

		if daysOfExperienceInt > 29 {
			monthsOfExperienceInt++
			daysOfExperienceInt -= 30
		}

		if monthsOfExperienceInt > 11 {
			yearsOfExperienceInt++
			monthsOfExperienceInt -= 12
		}
	}

	isActive := true
	contracts, err := r.Repo.GetEmployeeContracts(employee.ID, &dto.GetEmployeeContracts{Active: &isActive})

	if err != nil {
		return nil, errors.Wrap(err, "repo get employee contracts")
	}

	if len(contracts) > 0 {
		if contracts[0].DateOfStart != nil {
			contractStartDateTime, err := parseDate(*contracts[0].DateOfStart)

			if err != nil {
				return nil, errors.Wrap(err, "repo parse date")
			}
			contractStartDate = contractStartDateTime.Format("02.01.2006")

			startDay, startMonth, startYear := contractStartDateTime.Day(), contractStartDateTime.Month(), contractStartDateTime.Year()
			endDay, endMonth, endYear := time.Now().Day(), time.Now().Month(), time.Now().Year()

			dayDiff := endDay - startDay
			if dayDiff < 0 {
				dayDiff += 30
				endMonth--
			}

			monthDiff := int(endMonth) - int(startMonth)
			if monthDiff < 0 {
				monthDiff += 12
				endYear--
			}

			yearDiff := endYear - startYear

			yearsOfExperienceInt += yearDiff
			monthsOfExperienceInt += monthDiff
			daysOfExperienceInt += dayDiff

			if daysOfExperienceInt > 29 {
				monthsOfExperienceInt++
				daysOfExperienceInt -= 30
			}

			if monthsOfExperienceInt > 11 {
				yearsOfExperienceInt++
				monthsOfExperienceInt -= 12
			}
		}

		if contracts[0].DateOfEnd != nil {
			contractEndDateTime, err := parseDate(*contracts[0].DateOfEnd)

			if err != nil {
				return nil, errors.Wrap(err, "repo parse date")
			}
			contractEndDate = contractEndDateTime.Format("02.01.2006")
		}
		if contracts[0].DateOfSignature != nil {
			contractDateOfSignatureTime, err := parseDate(*contracts[0].DateOfSignature)

			if err != nil {
				return nil, errors.Wrap(err, "repo parse date")
			}

			workStartDate = contractDateOfSignatureTime.Format("02.01.2006")
		}

		if contracts[0].DateOfStart != nil && contracts[0].DateOfEnd != nil {
			contractStartDateTime, err := parseDate(*contracts[0].DateOfStart)

			if err != nil {
				return nil, errors.Wrap(err, "repo parse date")
			}

			contractEndDateTime, err := parseDate(*contracts[0].DateOfEnd)

			if err != nil {
				return nil, errors.Wrap(err, "repo parse date")
			}

			duration := contractEndDateTime.Sub(contractStartDateTime)

			days := int(duration.Hours() / 24)
			contractDuration = strconv.Itoa(days)
		}
	}

	employeesInOrganizationUnit, _ := r.Repo.GetEmployeesInOrganizationUnitsByProfileID(employee.ID)
	if employeesInOrganizationUnit != nil && employeesInOrganizationUnit.PositionInOrganizationUnitID != 0 {
		jobPositionInOrganizationUnit, err := r.Repo.GetJobPositionsInOrganizationUnitsByID(employeesInOrganizationUnit.PositionInOrganizationUnitID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get job position in organization units by id")
		}

		jobPosition, err := r.Repo.GetJobPositionByID(jobPositionInOrganizationUnit.JobPositionID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get job position by id")
		}

		systematization, err := r.Repo.GetSystematizationByID(jobPositionInOrganizationUnit.SystematizationID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get systematization by id")
		}

		if systematization.Active == 2 {
			fullSystematization, err := buildSystematizationOverviewResponse(r.Repo, systematization)

			if err != nil {
				return nil, errors.Wrap(err, "repo build systematization overview response")
			}

			jobPositionName = jobPosition.Title
			systematizationNumber = systematization.SerialNumber

			if systematization.DateOfActivation != nil {
				systematizationDateTime, err := parseDate(*systematization.DateOfActivation)

				if err != nil {
					return nil, errors.Wrap(err, "repo parse date")
				}
				systematizationDate = systematizationDateTime.Format("02.01.2006")
			}

			for _, item := range fullSystematization.ActiveEmployees {
				if item.ID == id {
					departmentName = item.Sector
					break
				}
			}

			for _, item := range *fullSystematization.Sectors {
				for _, employeeItem := range item.JobPositionsOrganizationUnits {
					for _, employee := range employeeItem.Employees {
						if employee.ID == id && employeeItem.Requirements != nil {
							jobPositionRequirments = *employeeItem.Requirements
						}
					}
				}
			}

			organizationUnit, err := r.Repo.GetOrganizationUnitByID(systematization.OrganizationUnitID)

			if err != nil {
				return nil, errors.Wrap(err, "repo get organization unit by id")
			}

			organizationUnitName = organizationUnit.Title
		}
	}

	sumDaysOfCurrentYear, availableDaysOfCurrentYear, availableDaysOfPreviousYear, err := GetNumberOfCurrentAndPreviousYearAvailableDays(r.Repo, id)
	if err != nil {
		return nil, errors.Wrap(err, "repo get number of current and previous year avaliable days")
	}

	remainingVacationDays = strconv.Itoa(availableDaysOfCurrentYear + availableDaysOfPreviousYear)
	usedVacationDays = strconv.Itoa(sumDaysOfCurrentYear - availableDaysOfCurrentYear)
	acquiredVacationDays = strconv.Itoa(sumDaysOfCurrentYear + availableDaysOfPreviousYear)

	salaries, err := r.Repo.GetEmployeeSalaryParams(id)
	if err != nil {
		return nil, errors.Wrap(err, "repo get employee salary params")
	}

	if len(salaries) > 0 {
		weeklyWorkingHours = salaries[0].WeeklyWorkHours
	}

	allAbsents, err := r.Repo.GetEmployeeAbsents(id, nil)

	if err != nil {
		return nil, errors.Wrap(err, "repo get employee absents")
	}

	if len(allAbsents) > 0 {
		vacationStartDateTime, err := parseDate(allAbsents[0].DateOfStart)

		if err != nil {
			return nil, errors.Wrap(err, "repo parse date")
		}
		vacationStartDate = vacationStartDateTime.Format("02.01.2006")

		vacationEndDateTime, err := parseDate(allAbsents[0].DateOfEnd)

		if err != nil {
			return nil, errors.Wrap(err, "repo parse date")
		}
		vacationEndDate = vacationEndDateTime.Format("02.01.2006")
	}

	employeeRatings, err := r.Repo.GetEmployeeEvaluations(id)

	if err != nil {
		return nil, errors.Wrap(err, "repo get employee evaluations")
	}

	if len(employeeRatings) > 0 {
		rating = employeeRatings[0].Score
	}

	educationTypes, err := r.Repo.GetDropdownSettings(&dto.GetSettingsInput{
		Entity: "education_academic_types",
	})

	if err != nil {
		return nil, errors.Wrap(err, "repo get droodown settings")
	}

	if len(educationTypes.Data) > 0 {
		educations, err := r.Repo.GetEmployeeEducations(dto.EducationInput{
			UserProfileID: id,
			TypeID:        &educationTypes.Data[0].ID,
		})

		if err != nil {
			return nil, errors.Wrap(err, "repo get employee educations")
		}

		if len(educations) > 0 {
			education = educations[0].AcademicTitle
		}
	}

	yearsOfExperience = strconv.Itoa(yearsOfExperienceInt)
	monthsOfExperience = strconv.Itoa(monthsOfExperienceInt)
	daysOfExperience = strconv.Itoa(daysOfExperienceInt)

	active := true
	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}

	resolution, err := r.Repo.GetJudgeResolutionList(&input)

	if err != nil {
		return nil, errors.Wrap(err, "repo get resolution list")
	}

	if len(resolution.Data) > 0 {

		filter := dto.JudgeResolutionsOrganizationUnitInput{
			ResolutionID:  &resolution.Data[0].ID,
			UserProfileID: &id,
		}

		judges, _, err := r.Repo.GetJudgeResolutionOrganizationUnit(&filter)

		if err != nil {
			return nil, errors.Wrap(err, "repo get judge resolution organization unit")
		}

		if len(judges) > 0 && judges[0].IsPresident {
			jobPositionName = "Predsjednik suda"
		} else if len(judges) > 0 {
			jobPositionName = "Sudija"
		}

		if len(judges) > 0 {
			organizationUnit, err := r.Repo.GetOrganizationUnitByID(judges[0].OrganizationUnitID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get organization unit by id")
			}
			organizationUnitName = organizationUnit.Title
		}

	}

	dataForTemplate := dto.UserDataForTemplate{
		CurrentYear:            currentYear,
		CurrentMonth:           currentMonth,
		CurrentDate:            currentDateString,
		FullName:               employee.FirstName + " " + employee.LastName,
		JMBG:                   employee.OfficialPersonalDocumentNumber,
		Street:                 employee.Address,
		OrganizationalUnit:     organizationUnitName,
		Department:             departmentName,
		Position:               jobPositionName,
		JobPositionRequirments: jobPositionRequirments,
		SystematizationNumber:  systematizationNumber,
		SystematizationDate:    systematizationDate,
		ContractStartDate:      contractStartDate,
		ContractEndDate:        contractEndDate,
		WorkStartDate:          workStartDate,
		ContractDurationInDays: contractDuration,
		AcquiredVacationDays:   acquiredVacationDays,
		RemainingVacationDays:  remainingVacationDays,
		UsedVacationDays:       usedVacationDays,
		WeeklyWorkingHours:     weeklyWorkingHours,
		VacationStartDate:      vacationStartDate,
		VacationEndDate:        vacationEndDate,
		Rating:                 rating,
		Education:              education,
		YearsOfExperience:      yearsOfExperience,
		MonthsOfExperience:     monthsOfExperience,
		DaysOfExperience:       daysOfExperience,
	}

	return &dataForTemplate, nil

}
