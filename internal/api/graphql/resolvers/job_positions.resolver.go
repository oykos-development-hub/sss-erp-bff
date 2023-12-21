package resolvers

import (
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"errors"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) JobPositionsResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []structs.JobPositions
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	search, searchOk := params.Args["search"].(string)

	if id != nil && shared.IsInteger(id) && id != 0 {
		jobPositionResponse, err := r.Repo.GetJobPositionById(id.(int))
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		items = append(items, *jobPositionResponse)
		total = 1
	} else {
		input := dto.GetJobPositionsInput{}
		if shared.IsInteger(page) && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if shared.IsInteger(size) && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if searchOk && search != "" {
			input.Search = &search
		}

		jobPositionsResponse, err := r.Repo.GetJobPositions(&input)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		items = jobPositionsResponse.Data
		total = jobPositionsResponse.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
	}, nil
}

func (r *Resolver) JobPositionsOrganizationUnitResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		items              []structs.JobPositionsInOrganizationUnitsSettings
		total              int
		organizationUnitId int
	)

	if params.Args["organization_unit_id"] == nil {
		organizationUnitId = 0
	} else {
		organizationUnitId = params.Args["organization_unit_id"].(int)
	}

	input := dto.GetJobPositionInOrganizationUnitsInput{
		OrganizationUnitID: &organizationUnitId,
	}
	jobPositionsInOrganizationUnitsResponse, err := r.Repo.GetJobPositionsInOrganizationUnits(&input)
	if err != nil {
		return &jobPositionsInOrganizationUnitsResponse, err
	}

	for _, jobPositionsInOrganizationUnits := range jobPositionsInOrganizationUnitsResponse.Data {
		getJobPositionResponse, err := r.Repo.GetJobPositionById(jobPositionsInOrganizationUnits.JobPositionId)
		if err != nil {
			return &items, err
		}
		item := structs.JobPositionsInOrganizationUnitsSettings{
			Id:    jobPositionsInOrganizationUnits.Id,
			Title: getJobPositionResponse.Title,
		}
		items = append(items, item)
	}

	total = len(items)

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
	}, nil
}

func (r *Resolver) JobPositionInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobPositions
	var jobPositionResponse *dto.GetJobPositionResponseMS
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		jobPositionResponse, err = r.Repo.UpdateJobPositions(itemId, &data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	} else {
		jobPositionResponse, err = r.Repo.CreateJobPositions(&data)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    jobPositionResponse.Data,
	}, nil
}

func (r *Resolver) JobPositionDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	if !shared.IsInteger(itemId) && !(itemId.(int) <= 0) {
		return apierrors.HandleAPIError(errors.New("You must pass the item id"))
	}

	err := r.Repo.DeleteJobPositions(itemId.(int))
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

func (r *Resolver) JobPositionInOrganizationUnitInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobPositionsInOrganizationUnits
	var jobPositionInOrganizationUnit *dto.GetJobPositionInOrganizationUnitsResponseMS
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)
	if data.Id > 0 {
		jobPositionInOrganizationUnit, err = r.Repo.UpdateJobPositionsInOrganizationUnits(&data)
	} else {
		jobPositionInOrganizationUnit, err = r.Repo.CreateJobPositionsInOrganizationUnits(&data)
	}

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	err = r.Repo.DeleteEmployeeInOrganizationUnit(jobPositionInOrganizationUnit.Data.Id)

	if len(data.Employees) > 0 {

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		for _, employeeId := range data.Employees {
			input := &structs.EmployeesInOrganizationUnits{
				PositionInOrganizationUnitId: jobPositionInOrganizationUnit.Data.Id,
				UserProfileId:                employeeId,
			}
			res, err := r.Repo.CreateEmployeesInOrganizationUnits(input)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			jobPositionInOrganizationUnit.Data.Employees = append(jobPositionInOrganizationUnit.Data.Employees, res.Id)
		}

	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    jobPositionInOrganizationUnit.Data,
	}, nil
}

func (r *Resolver) JobPositionInOrganizationUnitResolver(params graphql.ResolveParams) (interface{}, error) {
	var (
		items              []structs.JobPositionsInOrganizationUnitsSettings
		total              int
		organizationUnitId int
		officeUnitId       int
	)

	if params.Args["organization_unit_id"] != nil && params.Args["office_unit_id"] != nil {
		organizationUnitId = params.Args["organization_unit_id"].(int)
		officeUnitId = params.Args["office_unit_id"].(int)
		var myBool int = 2
		input := dto.GetSystematizationsInput{}
		input.OrganizationUnitID = &organizationUnitId
		input.Active = &myBool

		systematizationsResponse, err := r.Repo.GetSystematizations(&input)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		if len(systematizationsResponse.Data) > 0 {
			for _, systematization := range systematizationsResponse.Data {
				input := dto.GetJobPositionInOrganizationUnitsInput{
					OrganizationUnitID: &officeUnitId,
					SystematizationID:  &systematization.Id,
				}
				jobPositionsInOrganizationUnits, err := r.Repo.GetJobPositionsInOrganizationUnits(&input)
				if err != nil {
					return nil, err
				}
				for _, jobPositionOU := range jobPositionsInOrganizationUnits.Data {
					input := dto.GetEmployeesInOrganizationUnitInput{
						PositionInOrganizationUnit: &jobPositionOU.Id,
					}
					employeesInOrganizationUnit, _ := r.Repo.GetEmployeesInOrganizationUnitList(&input)

					if len(employeesInOrganizationUnit) < jobPositionOU.AvailableSlots {
						jobPosition, err := r.Repo.GetJobPositionById(jobPositionOU.JobPositionId)
						if err != nil {
							return nil, err
						}
						items = append(items, structs.JobPositionsInOrganizationUnitsSettings{
							Id:    jobPositionOU.Id,
							Title: jobPosition.Title,
						})
					}

				}
			}
		}

	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
		Items:   items,
	}, nil
}
func (r *Resolver) JobPositionInOrganizationUnitDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	if !shared.IsInteger(itemId) && !(itemId.(int) <= 0) {
		return apierrors.HandleAPIError(errors.New("You must pass the item id"))
	}

	err := r.Repo.DeleteJobPositionsInOrganizationUnits(itemId.(int))
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
