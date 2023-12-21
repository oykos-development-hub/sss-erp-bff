package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
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
	organizationUnitId := params.Args["organization_unit_id"]
	year := params.Args["year"]
	search := params.Args["search"]

	if id != nil && shared.IsInteger(id) && id != 0 {
		systematization, err := r.Repo.GetSystematizationById(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}
		systematizationResItem, err := buildSystematizationOverviewResponse(r.Repo, systematization)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		items = []dto.SystematizationOverviewResponse{systematizationResItem}
		total = 1
	} else {
		input := dto.GetSystematizationsInput{}
		if shared.IsInteger(page) && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if shared.IsInteger(size) && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if shared.IsInteger(organizationUnitId) && organizationUnitId.(int) > 0 {
			organizationUnitId := organizationUnitId.(int)
			input.OrganizationUnitID = &organizationUnitId
		}
		if year != nil {
			yearInput := year.(string)
			input.Year = &yearInput
		}
		if search != nil {
			searchInput := search.(string)
			input.Search = &searchInput
		}
		if shared.IsInteger(active) && active.(int) > 0 {
			activeValue := active.(int)
			input.Active = &activeValue
		}

		systematizationsResponse, err := r.Repo.GetSystematizations(&input)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		for _, systematization := range systematizationsResponse.Data {
			systematizationResItem, err := buildSystematizationOverviewResponse(r.Repo, &systematization)
			if err != nil {
				return errors.HandleAPIError(err)
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
	systematization, err := r.Repo.GetSystematizationById(id.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}
	systematizationResItem, err := buildSystematizationOverviewResponse(r.Repo, systematization)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    systematizationResItem,
	}, nil
}

func (r *Resolver) SystematizationInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var booActive int = 2
	var data structs.Systematization
	var systematization *structs.Systematization
	var err error

	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		systematization, err = r.Repo.UpdateSystematization(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		systematization, err = r.Repo.CreateSystematization(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	input := dto.GetSystematizationsInput{}
	input.OrganizationUnitID = &systematization.OrganizationUnitId
	input.Active = &booActive

	systematizationsActiveResponse, _ := r.Repo.GetSystematizations(&input)
	// Getting Sectors
	inputOrganizationUnits := dto.GetOrganizationUnitsInput{
		ParentID: &systematization.OrganizationUnitId,
	}
	organizationUnitsResponse, err := r.Repo.GetOrganizationUnits(&inputOrganizationUnits)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	if (!shared.IsInteger(itemId) || itemId == 0) && len(systematizationsActiveResponse.Data) > 0 {
		for _, sector := range organizationUnitsResponse.Data {
			input := dto.GetJobPositionInOrganizationUnitsInput{
				OrganizationUnitID: &sector.Id,
				SystematizationID:  &systematizationsActiveResponse.Data[0].Id,
			}
			jobPositionsInOrganizationUnits, err := r.Repo.GetJobPositionsInOrganizationUnits(&input)
			if err != nil {
				return errors.HandleAPIError(err)
			}
			for _, jobPositionOU := range jobPositionsInOrganizationUnits.Data {
				var jobPositionsInOrganizationUnitRes *dto.GetJobPositionInOrganizationUnitsResponseMS
				jobPositionsInOrganizationUnitRes, _ = r.Repo.CreateJobPositionsInOrganizationUnits(&structs.JobPositionsInOrganizationUnits{
					SystematizationId:        systematization.Id,
					ParentOrganizationUnitId: jobPositionOU.ParentOrganizationUnitId,
					JobPositionId:            jobPositionOU.JobPositionId,
					AvailableSlots:           jobPositionOU.AvailableSlots,
					Requirements:             jobPositionOU.Requirements,
					Description:              jobPositionOU.Description,
				})

				input := dto.GetEmployeesInOrganizationUnitInput{
					PositionInOrganizationUnit: &jobPositionOU.Id,
				}
				employeesInOrganizationUnit, _ := r.Repo.GetEmployeesInOrganizationUnitList(&input)

				if len(employeesInOrganizationUnit) > 0 {
					for _, employee := range employeesInOrganizationUnit {
						input := &structs.EmployeesInOrganizationUnits{
							PositionInOrganizationUnitId: jobPositionsInOrganizationUnitRes.Data.Id,
							UserProfileId:                employee.UserProfileId,
						}
						_, err := r.Repo.CreateEmployeesInOrganizationUnits(input)
						if err != nil {
							return errors.HandleAPIError(err)
						}
					}

				}

			}
		}
	}
	if err != nil {
		return errors.HandleAPIError(err)
	}

	if systematization.Active == 2 {
		if len(systematizationsActiveResponse.Data) > 0 {
			for _, sys := range systematizationsActiveResponse.Data {
				if sys.Id != systematization.Id {
					sys.Active = 1
					_, _ = r.Repo.UpdateSystematization(sys.Id, &sys)
				}
			}
		}
	}

	systematizationResItem, err := buildSystematizationOverviewResponse(r.Repo, systematization)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    systematizationResItem,
	}, nil
}

func (r *Resolver) SystematizationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	if !shared.IsInteger(itemId) && !(itemId.(int) <= 0) {
		return errors.ErrorResponse("You must pass the item id"), nil
	}

	err := r.Repo.DeleteSystematization(itemId.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildSystematizationOverviewResponse(r repository.MicroserviceRepositoryInterface, systematization *structs.Systematization) (dto.SystematizationOverviewResponse, error) {

	result := dto.SystematizationOverviewResponse{
		Id:                 systematization.Id,
		UserProfileId:      systematization.UserProfileId,
		OrganizationUnitId: systematization.OrganizationUnitId,
		Description:        systematization.Description,
		SerialNumber:       systematization.SerialNumber,
		FileId:             systematization.FileId,
		Active:             systematization.Active,
		DateOfActivation:   systematization.DateOfActivation,
		Sectors:            &[]dto.OrganizationUnitsSectorResponse{},
		CreatedAt:          systematization.CreatedAt,
		UpdatedAt:          systematization.UpdatedAt,
		ActiveEmployees:    []structs.ActiveEmployees{},
	}

	// Getting Organization Unit
	var relatedOrganizationUnit, err = r.GetOrganizationUnitById(systematization.OrganizationUnitId)
	if err != nil {
		return result, err
	}
	result.OrganizationUnit = relatedOrganizationUnit

	// Getting Sectors
	inputOrganizationUnits := dto.GetOrganizationUnitsInput{
		ParentID: &systematization.OrganizationUnitId,
	}
	organizationUnitsResponse, err := r.GetOrganizationUnits(&inputOrganizationUnits)
	if err != nil {
		return result, err
	}
	for _, organizationUnit := range organizationUnitsResponse.Data {
		*result.Sectors = append(*result.Sectors, *dto.ToOrganizationUnitsSectorResponse(organizationUnit))
	}

	// Getting Job positions
	if result.Sectors != nil {
		for i, sector := range *result.Sectors {
			// jobPositionInOrganizationBySectorId getJobPositionsInOrganizationUnits
			input := dto.GetJobPositionInOrganizationUnitsInput{
				OrganizationUnitID: &sector.Id,
				SystematizationID:  &systematization.Id,
			}
			jobPositionsInOrganizationUnits, err := r.GetJobPositionsInOrganizationUnits(&input)
			if err != nil {
				return result, err
			}
			var jobPositionsOrganizationUnits []dto.JobPositionsOrganizationUnits
			for _, jobPositionOU := range jobPositionsInOrganizationUnits.Data {
				jobPosition, err := r.GetJobPositionById(jobPositionOU.JobPositionId)
				if err != nil {
					return result, err
				}

				input := dto.GetEmployeesInOrganizationUnitInput{
					PositionInOrganizationUnit: &jobPositionOU.Id,
				}
				employeesInOrganizationUnit, _ := r.GetEmployeesInOrganizationUnitList(&input)
				var employees []dto.DropdownSimple
				for _, employeeID := range employeesInOrganizationUnit {
					employee, err := r.GetUserProfileById(employeeID.UserProfileId)
					if err != nil {
						return result, err
					}
					employees = append(employees, dto.DropdownSimple{
						Id:    employeeID.UserProfileId,
						Title: employee.FirstName + " " + employee.LastName,
					})
					if !jobPosition.IsJudgePresident {
						result.ActiveEmployees = append(result.ActiveEmployees, structs.ActiveEmployees{
							Id:       employeeID.UserProfileId,
							FullName: employee.FirstName + " " + employee.LastName,
							JobPositions: structs.SettingsDropdown{
								Id:    jobPosition.Id,
								Title: jobPosition.Title,
							},
							Sector: sector.Title,
						})
					}
				}
				// jobEmployeesByPositionInOrganizationId
				jobPositionsOrganizationUnits = append(jobPositionsOrganizationUnits, dto.JobPositionsOrganizationUnits{
					Id: jobPositionOU.Id,
					JobPositions: dto.DropdownSimple{
						Id:    jobPosition.Id,
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
