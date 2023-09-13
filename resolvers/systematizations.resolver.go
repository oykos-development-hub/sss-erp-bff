package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"
)

var SystematizationsOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []dto.SystematizationOverviewResponse
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	organizationUnitId := params.Args["organization_unit_id"]
	year := params.Args["year"]
	search := params.Args["search"]

	if id != nil && shared.IsInteger(id) && id != 0 {
		systematization, err := getSystematizationById(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		systematizationResItem, err := buildSystematizationOverviewResponse(systematization)
		if err != nil {
			return shared.HandleAPIError(err)
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

		systematizationsResponse, err := getSystematizations(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		for _, systematization := range systematizationsResponse.Data {
			systematizationResItem, err := buildSystematizationOverviewResponse(&systematization)
			if err != nil {
				return shared.HandleAPIError(err)
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

var SystematizationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	systematization, err := getSystematizationById(id.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}
	systematizationResItem, err := buildSystematizationOverviewResponse(systematization)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    systematizationResItem,
	}, nil
}

var SystematizationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var booActive bool = true
	var data structs.Systematization
	var systematization *structs.Systematization
	var err error

	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		systematization, err = updateSystematization(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	} else {
		systematization, err = createSystematization(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	input := dto.GetSystematizationsInput{}
	input.OrganizationUnitID = &systematization.OrganizationUnitId
	input.Active = &booActive

	systematizationsActiveResponse, _ := getSystematizations(&input)
	// Getting Sectors
	inputOrganizationUnits := dto.GetOrganizationUnitsInput{
		ParentID: &systematization.OrganizationUnitId,
	}
	organizationUnitsResponse, err := getOrganizationUnits(&inputOrganizationUnits)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	if !shared.IsInteger(itemId) || itemId == 0 && len(systematizationsActiveResponse.Data) > 0 {
		for _, sector := range organizationUnitsResponse.Data {
			input := dto.GetJobPositionInOrganizationUnitsInput{
				OrganizationUnitID: &sector.Id,
				SystematizationID:  &systematizationsActiveResponse.Data[0].Id,
			}
			jobPositionsInOrganizationUnits, err := getJobPositionsInOrganizationUnits(&input)
			if err != nil {
				return nil, err
			}
			for _, jobPositionOU := range jobPositionsInOrganizationUnits.Data {
				var jobPositionsInOrganizationUnitRes *dto.GetJobPositionInOrganizationUnitsResponseMS
				jobPositionsInOrganizationUnitRes, _ = createJobPositionsInOrganizationUnits(&structs.JobPositionsInOrganizationUnits{
					SystematizationId:        systematization.Id,
					ParentOrganizationUnitId: jobPositionOU.ParentOrganizationUnitId,
					JobPositionId:            jobPositionOU.JobPositionId,
					AvailableSlots:           jobPositionOU.AvailableSlots,
				})

				input := dto.GetEmployeesInOrganizationUnitInput{
					PositionInOrganizationUnit: &jobPositionOU.Id,
				}
				employeesInOrganizationUnit, _ := getEmployeesInOrganizationUnitList(&input)

				if len(employeesInOrganizationUnit) > 0 {
					for _, employee := range employeesInOrganizationUnit {
						input := &structs.EmployeesInOrganizationUnits{
							PositionInOrganizationUnitId: jobPositionsInOrganizationUnitRes.Data.Id,
							UserProfileId:                employee.UserProfileId,
						}
						_, err := createEmployeesInOrganizationUnits(input)
						if err != nil {
							return nil, err
						}
					}

				}

			}
		}
	}
	if err != nil {
		return shared.HandleAPIError(err)
	}

	if systematization.Active {

		if len(systematizationsActiveResponse.Data) > 0 {
			for _, sys := range systematizationsActiveResponse.Data {
				if sys.Id != systematization.Id {
					sys.Active = false
					updateSystematization(sys.Id, &sys)
				}
			}
		}
	}

	systematizationResItem, err := buildSystematizationOverviewResponse(systematization)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    systematizationResItem,
	}, nil
}

var SystematizationDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	if !shared.IsInteger(itemId) && !(itemId.(int) <= 0) {
		return shared.ErrorResponse("You must pass the item id"), nil
	}

	err := deleteSystematization(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

func getSystematizationById(id int) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.SYSTEMATIZATIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getSystematizations(input *dto.GetSystematizationsInput) (*dto.GetSystematizationsResponseMS, error) {
	res := &dto.GetSystematizationsResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.SYSTEMATIZATIONS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func updateSystematization(id int, data *structs.Systematization) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.SYSTEMATIZATIONS_ENDPOINT+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createSystematization(data *structs.Systematization) (*structs.Systematization, error) {
	res := &dto.GetSystematizationResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.SYSTEMATIZATIONS_ENDPOINT, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteSystematization(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.SYSTEMATIZATIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func buildSystematizationOverviewResponse(systematization *structs.Systematization) (dto.SystematizationOverviewResponse, error) {
	result := dto.SystematizationOverviewResponse{
		Id:                 systematization.Id,
		UserProfileId:      systematization.UserProfileId,
		OrganizationUnitId: systematization.OrganizationUnitId,
		Description:        systematization.Description,
		SerialNumber:       systematization.SerialNumber,
		Active:             systematization.Active,
		FileId:             systematization.FileId,
		DateOfActivation:   systematization.DateOfActivation,
		Sectors:            &[]dto.OrganizationUnitsSectorResponse{},
		CreatedAt:          systematization.CreatedAt,
		UpdatedAt:          systematization.UpdatedAt,
		ActiveEmployees:    []structs.ActiveEmployees{},
	}

	// Getting Organization Unit
	var relatedOrganizationUnit, err = getOrganizationUnitById(systematization.OrganizationUnitId)
	if err != nil {
		return result, err
	}
	result.OrganizationUnit = relatedOrganizationUnit

	// Getting Sectors
	inputOrganizationUnits := dto.GetOrganizationUnitsInput{
		ParentID: &systematization.OrganizationUnitId,
	}
	organizationUnitsResponse, err := getOrganizationUnits(&inputOrganizationUnits)
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
			jobPositionsInOrganizationUnits, err := getJobPositionsInOrganizationUnits(&input)
			if err != nil {
				return result, err
			}
			var jobPositionsOrganizationUnits []dto.JobPositionsOrganizationUnits
			for _, jobPositionOU := range jobPositionsInOrganizationUnits.Data {
				jobPosition, err := getJobPositionById(jobPositionOU.JobPositionId)
				if err != nil {
					return result, err
				}

				input := dto.GetEmployeesInOrganizationUnitInput{
					PositionInOrganizationUnit: &jobPositionOU.Id,
				}
				employeesInOrganizationUnit, _ := getEmployeesInOrganizationUnitList(&input)
				var employees []dto.DropdownSimple
				for _, employeeID := range employeesInOrganizationUnit {
					employee, err := getUserProfileById(employeeID.UserProfileId)
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
					Description:    jobPosition.Description,
					SerialNumber:   jobPosition.SerialNumber,
					Requirements:   jobPosition.Requirements,
					AvailableSlots: jobPositionOU.AvailableSlots,
					Employees:      employees,
				})

			}
			(*result.Sectors)[i].JobPositionsOrganizationUnits = jobPositionsOrganizationUnits

		}
	}

	return result, nil
}
