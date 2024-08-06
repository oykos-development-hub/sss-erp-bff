package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) SystematizationsOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []dto.SystematizationOverviewResponse
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	active := params.Args["active"]
	organizationUnitID := params.Args["organization_unit_id"]
	year := params.Args["year"]
	search := params.Args["search"]

	if id != nil && id != 0 {
		systematization, err := r.Repo.GetSystematizationByID(id.(int))
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		systematizationResItem, err := buildSystematizationOverviewResponse(r.Repo, systematization)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		items = []dto.SystematizationOverviewResponse{systematizationResItem}
		total = 1
	} else {
		input := dto.GetSystematizationsInput{}
		if page != nil && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if size != nil && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if organizationUnitID != nil && organizationUnitID.(int) > 0 {
			organizationUnitID := organizationUnitID.(int)
			input.OrganizationUnitID = &organizationUnitID
		}
		if year != nil {
			yearInput := year.(string)
			input.Year = &yearInput
		}
		if search != nil {
			searchInput := search.(string)
			input.Search = &searchInput
		}
		if active != nil && active.(int) > 0 {
			activeValue := active.(int)
			input.Active = &activeValue
		}
		loggedInOrganizationUnitID, _ := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		loggedInUserAccount, _ := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

		systematizationsResponse, err := r.Repo.GetSystematizations(&input)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		for _, systematization := range systematizationsResponse.Data {
			systematizationResItem, err := buildSystematizationOverviewResponse(r.Repo, &systematization)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			hasPermission, err := r.HasPermission(*loggedInUserAccount, string(config.HR), config.OperationFullAccess)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			if !hasPermission && *loggedInOrganizationUnitID != systematizationResItem.OrganizationUnitID {
				continue
			}
			items = append(items, systematizationResItem)
		}
		total = systematizationsResponse.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
	}, nil
}

func (r *Resolver) SystematizationResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	systematization, err := r.Repo.GetSystematizationByID(id.(int))
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	systematizationResItem, err := buildSystematizationOverviewResponse(r.Repo, systematization)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    systematizationResItem,
	}, nil
}

func (r *Resolver) SystematizationInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	booActive := 2
	var data structs.Systematization
	var systematization *structs.Systematization
	var err error

	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		systematization, err = r.Repo.UpdateSystematization(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		systematization, err = r.Repo.CreateSystematization(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	input := dto.GetSystematizationsInput{}
	input.OrganizationUnitID = &systematization.OrganizationUnitID
	input.Active = &booActive

	systematizationsActiveResponse, _ := r.Repo.GetSystematizations(&input)
	// Getting Sectors
	inputOrganizationUnits := dto.GetOrganizationUnitsInput{
		ParentID: &systematization.OrganizationUnitID,
	}

	if systematization.Active == 2 {
		if len(systematizationsActiveResponse.Data) > 0 {
			for _, sys := range systematizationsActiveResponse.Data {
				if sys.ID != systematization.ID && sys.Active == 2 {
					sys.Active = 1
					_, _ = r.Repo.UpdateSystematization(params.Context, sys.ID, &sys)
				}
			}
		}
	}

	organizationUnitsResponse, err := r.Repo.GetOrganizationUnits(&inputOrganizationUnits)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	if (itemID == 0) && len(systematizationsActiveResponse.Data) > 0 {
		for _, sector := range organizationUnitsResponse.Data {
			input := dto.GetJobPositionInOrganizationUnitsInput{
				OrganizationUnitID: &sector.ID,
				SystematizationID:  &systematizationsActiveResponse.Data[0].ID,
			}
			jobPositionsInOrganizationUnits, err := r.Repo.GetJobPositionsInOrganizationUnits(&input)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			for _, jobPositionOU := range jobPositionsInOrganizationUnits.Data {
				var jobPositionsInOrganizationUnitRes *dto.GetJobPositionInOrganizationUnitsResponseMS
				jobPositionsInOrganizationUnitRes, _ = r.Repo.CreateJobPositionsInOrganizationUnits(&structs.JobPositionsInOrganizationUnits{
					SystematizationID:        systematization.ID,
					ParentOrganizationUnitID: jobPositionOU.ParentOrganizationUnitID,
					JobPositionID:            jobPositionOU.JobPositionID,
					AvailableSlots:           jobPositionOU.AvailableSlots,
					Requirements:             jobPositionOU.Requirements,
					Description:              jobPositionOU.Description,
				})

				input := dto.GetEmployeesInOrganizationUnitInput{
					PositionInOrganizationUnit: &jobPositionOU.ID,
				}
				employeesInOrganizationUnit, _ := r.Repo.GetEmployeesInOrganizationUnitList(&input)

				if len(employeesInOrganizationUnit) > 0 {
					for _, employee := range employeesInOrganizationUnit {
						input := &structs.EmployeesInOrganizationUnits{
							PositionInOrganizationUnitID: jobPositionsInOrganizationUnitRes.Data.ID,
							UserProfileID:                employee.UserProfileID,
						}
						_, err := r.Repo.CreateEmployeesInOrganizationUnits(input)
						if err != nil {
							_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
							return errors.HandleAPPError(err)
						}
					}

				}

			}
		}
	}

	systematizationResItem, err := buildSystematizationOverviewResponse(r.Repo, systematization)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    systematizationResItem,
	}, nil
}

func (r *Resolver) SystematizationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]

	err := r.Repo.DeleteSystematization(params.Context, itemID.(int))
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildSystematizationOverviewResponse(r repository.MicroserviceRepositoryInterface, systematization *structs.Systematization) (dto.SystematizationOverviewResponse, error) {

	result := dto.SystematizationOverviewResponse{
		ID:                 systematization.ID,
		UserProfileID:      systematization.UserProfileID,
		OrganizationUnitID: systematization.OrganizationUnitID,
		Description:        systematization.Description,
		SerialNumber:       systematization.SerialNumber,
		FileID:             systematization.FileID,
		Active:             systematization.Active,
		DateOfActivation:   systematization.DateOfActivation,
		Sectors:            &[]dto.OrganizationUnitsSectorResponse{},
		CreatedAt:          systematization.CreatedAt,
		UpdatedAt:          systematization.UpdatedAt,
		ActiveEmployees:    []structs.ActiveEmployees{},
	}

	// Getting Organization Unit
	var relatedOrganizationUnit, err = r.GetOrganizationUnitByID(systematization.OrganizationUnitID)
	if err != nil {
		return result, errors.Wrap(err, "repo get organization unit by id")
	}
	result.OrganizationUnit = relatedOrganizationUnit

	// Getting Sectors
	inputOrganizationUnits := dto.GetOrganizationUnitsInput{
		ParentID: &systematization.OrganizationUnitID,
	}
	organizationUnitsResponse, err := r.GetOrganizationUnits(&inputOrganizationUnits)
	if err != nil {
		return result, errors.Wrap(err, "repo get organization units")
	}
	for _, organizationUnit := range organizationUnitsResponse.Data {
		*result.Sectors = append(*result.Sectors, *dto.ToOrganizationUnitsSectorResponse(organizationUnit))
	}

	if systematization.FileID != 0 {
		file, err := r.GetFileByID(systematization.FileID)
		if err != nil {
			return result, errors.Wrap(err, "repo get file by id")
		}
		result.File = dto.FileDropdownSimple{
			ID:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		}
	}

	// Getting Job positions
	if result.Sectors != nil {
		for i, sector := range *result.Sectors {
			// jobPositionInOrganizationBySectorID getJobPositionsInOrganizationUnits
			input := dto.GetJobPositionInOrganizationUnitsInput{
				OrganizationUnitID: &sector.ID,
				SystematizationID:  &systematization.ID,
			}
			jobPositionsInOrganizationUnits, err := r.GetJobPositionsInOrganizationUnits(&input)
			if err != nil {
				return result, errors.Wrap(err, "repo get job positions in organization units")
			}
			var jobPositionsOrganizationUnits []dto.JobPositionsOrganizationUnits
			for _, jobPositionOU := range jobPositionsInOrganizationUnits.Data {
				jobPosition, err := r.GetJobPositionByID(jobPositionOU.JobPositionID)
				if err != nil {
					return result, errors.Wrap(err, "repo get job position by id")
				}

				input := dto.GetEmployeesInOrganizationUnitInput{
					PositionInOrganizationUnit: &jobPositionOU.ID,
				}
				employeesInOrganizationUnit, _ := r.GetEmployeesInOrganizationUnitList(&input)
				var employees []dto.DropdownSimple
				for _, employeeID := range employeesInOrganizationUnit {
					employee, err := r.GetUserProfileByID(employeeID.UserProfileID)
					if err != nil {
						return result, errors.Wrap(err, "repo get user profile by id")
					}
					employees = append(employees, dto.DropdownSimple{
						ID:    employeeID.UserProfileID,
						Title: employee.FirstName + " " + employee.LastName,
					})
					if !jobPosition.IsJudgePresident {
						result.ActiveEmployees = append(result.ActiveEmployees, structs.ActiveEmployees{
							ID:       employeeID.UserProfileID,
							FullName: employee.FirstName + " " + employee.LastName,
							JobPositions: structs.SettingsDropdown{
								ID:    jobPosition.ID,
								Title: jobPosition.Title,
							},
							Sector: sector.Title,
						})
					}
				}
				// jobEmployeesByPositionInOrganizationID
				jobPositionsOrganizationUnits = append(jobPositionsOrganizationUnits, dto.JobPositionsOrganizationUnits{
					ID: jobPositionOU.ID,
					JobPositions: dto.DropdownSimple{
						ID:    jobPosition.ID,
						Title: jobPosition.Title,
					},
					Description:    jobPositionOU.Description,
					SerialNumber:   jobPosition.SerialNumber,
					Requirements:   jobPositionOU.Requirements,
					AvailableSlots: jobPositionOU.AvailableSlots,
					Employees:      employees,
				})

			}
			(*result.Sectors)[i].JobPositionsOrganizationUnits = jobPositionsOrganizationUnits

		}
	}

	return result, nil
}
