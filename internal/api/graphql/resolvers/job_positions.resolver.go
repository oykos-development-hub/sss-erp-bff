package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

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

	if id != nil && id != 0 {
		jobPositionResponse, err := r.Repo.GetJobPositionByID(id.(int))
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		items = append(items, *jobPositionResponse)
		total = 1
	} else {
		input := dto.GetJobPositionsInput{}
		if page != nil && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if size != nil && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if searchOk && search != "" {
			input.Search = &search
		}

		jobPositionsResponse, err := r.Repo.GetJobPositions(&input)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
		organizationUnitID int
	)

	if params.Args["organization_unit_id"] == nil {
		organizationUnitID = 0
	} else {
		organizationUnitID = params.Args["organization_unit_id"].(int)
	}

	input := dto.GetJobPositionInOrganizationUnitsInput{
		OrganizationUnitID: &organizationUnitID,
	}
	jobPositionsInOrganizationUnitsResponse, err := r.Repo.GetJobPositionsInOrganizationUnits(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, jobPositionsInOrganizationUnits := range jobPositionsInOrganizationUnitsResponse.Data {
		getJobPositionResponse, err := r.Repo.GetJobPositionByID(jobPositionsInOrganizationUnits.JobPositionID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item := structs.JobPositionsInOrganizationUnitsSettings{
			ID:    jobPositionsInOrganizationUnits.ID,
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

	itemID := data.ID
	if itemID != 0 {
		jobPositionResponse, err = r.Repo.UpdateJobPositions(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		jobPositionResponse, err = r.Repo.CreateJobPositions(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    jobPositionResponse.Data,
	}, nil
}

func (r *Resolver) JobPositionDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]

	err := r.Repo.DeleteJobPositions(params.Context, itemID.(int))
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
	if data.ID > 0 {
		jobPositionInOrganizationUnit, err = r.Repo.UpdateJobPositionsInOrganizationUnits(&data)
	} else {
		jobPositionInOrganizationUnit, err = r.Repo.CreateJobPositionsInOrganizationUnits(&data)
	}

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	err = r.Repo.DeleteEmployeeInOrganizationUnit(jobPositionInOrganizationUnit.Data.ID)

	if len(data.Employees) > 0 {

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, employeeID := range data.Employees {
			input := &structs.EmployeesInOrganizationUnits{
				PositionInOrganizationUnitID: jobPositionInOrganizationUnit.Data.ID,
				UserProfileID:                employeeID,
				Active:                       true,
			}
			res, err := r.Repo.CreateEmployeesInOrganizationUnits(input)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			jobPositionInOrganizationUnit.Data.Employees = append(jobPositionInOrganizationUnit.Data.Employees, res.ID)
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
		organizationUnitID int
		officeUnitID       int
	)

	if params.Args["organization_unit_id"] != nil && params.Args["office_unit_id"] != nil {
		organizationUnitID = params.Args["organization_unit_id"].(int)
		officeUnitID = params.Args["office_unit_id"].(int)
		myBool := 2
		input := dto.GetSystematizationsInput{}
		input.OrganizationUnitID = &organizationUnitID
		input.Active = &myBool

		systematizationsResponse, err := r.Repo.GetSystematizations(&input)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		if len(systematizationsResponse.Data) > 0 {
			for _, systematization := range systematizationsResponse.Data {
				input := dto.GetJobPositionInOrganizationUnitsInput{
					OrganizationUnitID: &officeUnitID,
					SystematizationID:  &systematization.ID,
				}
				jobPositionsInOrganizationUnits, err := r.Repo.GetJobPositionsInOrganizationUnits(&input)
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
				for _, jobPositionOU := range jobPositionsInOrganizationUnits.Data {
					input := dto.GetEmployeesInOrganizationUnitInput{
						PositionInOrganizationUnit: &jobPositionOU.ID,
					}
					employeesInOrganizationUnit, _ := r.Repo.GetEmployeesInOrganizationUnitList(&input)

					if len(employeesInOrganizationUnit) < jobPositionOU.AvailableSlots {
						jobPosition, err := r.Repo.GetJobPositionByID(jobPositionOU.JobPositionID)
						if err != nil {
							_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
							return errors.HandleAPPError(err)
						}
						items = append(items, structs.JobPositionsInOrganizationUnitsSettings{
							ID:    jobPositionOU.ID,
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
	itemID := params.Args["id"]

	err := r.Repo.DeleteJobPositionsInOrganizationUnits(itemID.(int))
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
