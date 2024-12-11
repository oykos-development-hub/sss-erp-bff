package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) ActiveContractMutateResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var contract dto.ContractInsert

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &contract)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	contract.Active = true

	active := true

	userData, err := r.Repo.GetUserProfileByID(contract.UserProfileID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	entityData := contract.ToEntity()

	if contract.ID == 0 {
		_, err = r.Repo.CreateEmployeeContract(params.Context, entityData)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		_, err = r.Repo.UpdateEmployeeContract(params.Context, contract.ID, entityData)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	if contract.JobPositionInOrganizationUnitID > -1 {
		myBool := 2
		check := true
		inputSys := dto.GetSystematizationsInput{}
		inputSys.OrganizationUnitID = &contract.OrganizationUnitID
		inputSys.Active = &myBool

		systematizationsResponse, err := r.Repo.GetSystematizations(&inputSys)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, systematization := range systematizationsResponse.Data {
			jobPositionsInOrganizationUnits, err := r.Repo.GetJobPositionsInOrganizationUnits(&dto.GetJobPositionInOrganizationUnitsInput{
				SystematizationID: &systematization.ID,
			})
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			for _, job := range jobPositionsInOrganizationUnits.Data {
				employeesInOrganizationUnit, _ := r.Repo.GetEmployeesInOrganizationUnitList(&dto.GetEmployeesInOrganizationUnitInput{
					PositionInOrganizationUnit: &job.ID,
					UserProfileID:              &userData.ID,
				})
				for _, emp := range employeesInOrganizationUnit {
					if emp.UserProfileID == userData.ID {
						if emp.PositionInOrganizationUnitID == contract.JobPositionInOrganizationUnitID {
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

		if contract.JobPositionInOrganizationUnitID > 0 && check {
			input := &structs.EmployeesInOrganizationUnits{
				PositionInOrganizationUnitID: contract.JobPositionInOrganizationUnitID,
				UserProfileID:                userData.ID,
			}
			_, err = r.Repo.CreateEmployeesInOrganizationUnits(input)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}
	}

	resolution, _ := r.Repo.GetJudgeResolutionList(&dto.GetJudgeResolutionListInputMS{
		Active: &active,
	})

	if len(resolution.Data) > 0 {
		judgeResolutionOrganizationUnit, _, _ := r.Repo.GetJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
			OrganizationUnitID: &contract.OrganizationUnitID,
			UserProfileID:      &userData.ID,
			ResolutionID:       &resolution.Data[0].ID,
		})

		if len(judgeResolutionOrganizationUnit) > 0 {
			if userData.IsJudge {
				inputUpdate := dto.JudgeResolutionsOrganizationUnitItem{
					ID:                 judgeResolutionOrganizationUnit[0].ID,
					UserProfileID:      userData.ID,
					OrganizationUnitID: contract.OrganizationUnitID,
					ResolutionID:       resolution.Data[0].ID,
					IsPresident:        userData.IsPresident,
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
		if len(judgeResolutionOrganizationUnit) == 0 && userData.IsJudge {
			inputCreate := dto.JudgeResolutionsOrganizationUnitItem{
				UserProfileID:      userData.ID,
				OrganizationUnitID: contract.OrganizationUnitID,
				ResolutionID:       resolution.Data[0].ID,
				IsPresident:        userData.IsPresident,
			}
			_, err := r.Repo.CreateJudgeResolutionOrganizationUnit(&inputCreate)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}

	}

	if userData.IsJudge && !contract.IsJudge {
		if len(resolution.Data) > 0 {
			judgeResolutionOrganizationUnit, _, err := r.Repo.GetJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
				OrganizationUnitID: &contract.OrganizationUnitID,
				UserProfileID:      &userData.ID,
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

	if !userData.IsJudge && contract.IsJudge {
		item, err := r.Repo.GetEmployeesInOrganizationUnitsByProfileID(userData.ID)

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

	userData.ActiveContract = &active
	userData.IsJudge = contract.IsJudge
	userData.IsPresident = contract.IsPresident
	userData.JudgeApplicationSubmissionDate = contract.JudgeApplicationSubmissionDate

	_, err = r.Repo.UpdateUserProfile(params.Context, userData.ID, *userData)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	account, err := r.Repo.GetUserAccountByID(userData.UserAccountID)
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

	res, err := buildContractResponseItem(r.Repo, *entityData)
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

func (r *Resolver) UserProfileContractsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]

	active := true
	contracts, err := r.Repo.GetEmployeeContracts(id.(int), &dto.GetEmployeeContracts{Active: &active})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var res *dto.Contract

	if len(contracts) > 0 {
		res, err = buildContractResponseItem(r.Repo, *contracts[0])
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    res,
	}, nil
}

func (r *Resolver) TerminateEmployment(params graphql.ResolveParams) (interface{}, error) {
	userID := params.Args["user_profile_id"].(int)

	user, err := r.Repo.GetUserProfileByID(userID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	userResponse, err := buildUserProfileOverviewResponse(r.Repo, user)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	active := true

	if userResponse.IsJudge {
		input := dto.GetJudgeResolutionListInputMS{
			Active: &active,
		}

		resolution, err := r.Repo.GetJudgeResolutionList(&input)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		filter := dto.JudgeResolutionsOrganizationUnitInput{
			UserProfileID: &userID,
			ResolutionID:  &resolution.Data[0].ID,
		}

		judge, _, err := r.Repo.GetJudgeResolutionOrganizationUnit(&filter)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		err = r.Repo.DeleteJJudgeResolutionOrganizationUnit(judge[0].ID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		contract, err := r.Repo.GetEmployeeContracts(userID, nil)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		now := time.Now()
		format := config.ISO8601Format
		dateOfEnd := now.Format(format)
		dateOfStart, _ := time.Parse(format, *contract[0].DateOfStart)

		yearsDiff := time.Now().Year() - dateOfStart.Year()
		monthsDiff := int(time.Now().Month()) - int(dateOfStart.Month())

		if monthsDiff < 0 {
			monthsDiff += 12
			yearsDiff--
		}

		daysDiff := int(time.Now().Day()) - int(dateOfStart.Day())

		if daysDiff < 0 {
			monthsDiff--
			daysDiff += 30
			if monthsDiff < 0 {
				yearsDiff--
				monthsDiff += 12
			}
		}

		experience := structs.Experience{
			UserProfileID:             userID,
			OrganizationUnitID:        judge[0].OrganizationUnitID,
			Relevant:                  true,
			DateOfStart:               *contract[0].DateOfStart,
			DateOfEnd:                 dateOfEnd,
			YearsOfExperience:         yearsDiff,
			YearsOfInsuredExperience:  yearsDiff,
			MonthsOfExperience:        monthsDiff,
			MonthsOfInsuredExperience: monthsDiff,
			DaysOfExperience:          daysDiff,
			DaysOfInsuredExperience:   daysDiff,
		}
		_, err = r.Repo.CreateExperience(&experience)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

	} else {
		contract, err := r.Repo.GetEmployeeContracts(userID, nil)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		employeeInOrgUnit, err := r.Repo.GetEmployeesInOrganizationUnitsByProfileID(userID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if employeeInOrgUnit != nil {
			err = r.Repo.DeleteEmployeeInOrganizationUnitByID(employeeInOrgUnit.ID)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			now := time.Now()
			format := config.ISO8601Format
			dateOfEnd := now.Format(format)
			dateOfStart, _ := time.Parse(format, *contract[0].DateOfStart)
			yearsDiff := time.Now().Year() - dateOfStart.Year()
			monthsDiff := int(time.Now().Month()) - int(dateOfStart.Month())

			if monthsDiff < 0 {
				monthsDiff += 12
				yearsDiff--
			}

			daysDiff := int(time.Now().Day()) - int(dateOfStart.Day())

			if daysDiff < 0 {
				monthsDiff--
				daysDiff += 30
				if monthsDiff < 0 {
					yearsDiff--
					monthsDiff += 12
				}
			}

			if monthsDiff < 0 {
				monthsDiff += 12
				yearsDiff--
			}

			experience := structs.Experience{
				UserProfileID:             userID,
				OrganizationUnitID:        contract[0].OrganizationUnitID,
				Relevant:                  true,
				DateOfStart:               *contract[0].DateOfStart,
				DateOfEnd:                 dateOfEnd,
				YearsOfExperience:         yearsDiff,
				YearsOfInsuredExperience:  yearsDiff,
				MonthsOfExperience:        monthsDiff,
				MonthsOfInsuredExperience: monthsDiff,
				DaysOfExperience:          daysDiff,
				DaysOfInsuredExperience:   daysDiff,
			}
			_, err = r.Repo.CreateExperience(&experience)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			contract[0].JobPositionInOrganizationUnitID = 0
			contract[0].OrganizationUnitDepartmentID = nil
			contract[0].Active = false
			_, err = r.Repo.UpdateEmployeeContract(params.Context, contract[0].ID, contract[0])
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}
	}

	terminateResolutionTypeValue := config.EmploymentTerminationResolutionType

	terminateResolutionType, err := r.Repo.GetDropdownSettings(&dto.GetSettingsInput{Value: &terminateResolutionTypeValue, Entity: config.ResolutionTypes})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	now := time.Now()
	rawFileIDs, ok := params.Args["file_ids"].([]interface{})
	if !ok {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: "file ids shlold be slice of integers"})
		return errors.HandleAPPError(err)
	}

	fileIDs := make([]int, len(rawFileIDs))
	for i, rawID := range rawFileIDs {
		id, ok := rawID.(int)
		if !ok {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: "file ids shlold be slice of integers"})
			return errors.HandleAPPError(err)
		}
		fileIDs[i] = id
	}

	_, err = r.Repo.CreateResolution(params.Context, &structs.Resolution{
		UserProfileID:    userID,
		ResolutionTypeID: terminateResolutionType.Data[0].ID,
		IsAffect:         true,
		DateOfStart:      now.Format(config.ISO8601Format),
		FileIDs:          fileIDs,
	})
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	inactive := false
	user.ActiveContract = &inactive
	_, err = r.Repo.UpdateUserProfile(params.Context, user.ID, *user)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	_, err = r.Repo.DeactivateUserAccount(params.Context, user.UserAccountID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deactivated this user!",
	}, nil
}

func buildContractResponseItem(r repository.MicroserviceRepositoryInterface, contract structs.Contracts) (*dto.Contract, error) {
	var files []dto.FileDropdownSimple

	if len(contract.FileIDs) > 0 {
		for i := range contract.FileIDs {
			res, _ := r.GetFileByID(contract.FileIDs[i])
			/*
				if err != nil {
					return nil, errors.Wrap(err, "repo get file by id")
				}
			*/

			if res != nil {
				files = append(files, dto.FileDropdownSimple{
					ID:   res.ID,
					Name: res.Name,
					Type: *res.Type,
				})
			}
		}

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
		Files:              files,
	}

	contractType, _ := r.GetDropdownSettingByID(contract.ContractTypeID)
	/*if err != nil {
		return nil, errors.Wrap(err, "repo get dropdown setting by id")
	}
	*/

	if contractType != nil {
		responseContract.ContractType = dto.DropdownSimple{ID: contractType.ID, Title: contractType.Title}
	}

	userProfile, err := r.GetUserProfileByID(contract.UserProfileID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get user profile by id")
	}
	responseContract.UserProfile = dto.DropdownSimple{ID: userProfile.ID, Title: userProfile.GetFullName()}
	responseContract.IsJudge = userProfile.IsJudge
	responseContract.IsPresident = userProfile.IsPresident
	responseContract.JudgeApplicationSubmissionDate = userProfile.JudgeApplicationSubmissionDate

	organizationUnit, err := r.GetOrganizationUnitByID(contract.OrganizationUnitID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get organization unit by id")
	}
	responseContract.OrganizationUnit = dto.DropdownSimple{ID: organizationUnit.ID, Title: organizationUnit.Title}

	if contract.Active {
		if contract.OrganizationUnitDepartmentID != nil {
			department, err := r.GetOrganizationUnitByID(*contract.OrganizationUnitDepartmentID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get organization unit by id")
			}
			responseContract.Department = &dto.DropdownSimple{ID: department.ID, Title: department.Title}
		}

		myBool := 2
		inputSys := dto.GetSystematizationsInput{}
		inputSys.OrganizationUnitID = &contract.OrganizationUnitID
		inputSys.Active = &myBool

		systematizationsResponse, err := r.GetSystematizations(&inputSys)
		if err != nil {
			return nil, errors.Wrap(err, "repo get systematizations")
		}
		var jobPositionInOU structs.JobPositionsInOrganizationUnits
		if len(systematizationsResponse.Data) > 0 {
			for _, systematization := range systematizationsResponse.Data {

				inputJpbPos := dto.GetJobPositionInOrganizationUnitsInput{
					SystematizationID: &systematization.ID,
				}
				jobPositionsInOrganizationUnits, err := r.GetJobPositionsInOrganizationUnits(&inputJpbPos)
				if err != nil {
					return nil, errors.Wrap(err, "repo get job positions in organization units")
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
				return nil, errors.Wrap(err, "repo get job position by id")
			}
			responseContract.JobPositionInOrganizationUnit = dto.DropdownSimple{ID: jobPositionInOU.ID, Title: jobPosition.Title}
		}
	}

	return responseContract, nil
}
