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

var JobPositionsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var (
		items []structs.JobPositions
		total int
	)

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	search, searchOk := params.Args["search"].(string)

	if id != nil && shared.IsInteger(id) && id != 0 {
		jobPositionResponse, err := getJobPositionById(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
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

		jobPositionsResponse, err := getJobPositions(&input)
		if err != nil {
			return shared.HandleAPIError(err)
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

var JobPositionsOrganizationUnitResolver = func(params graphql.ResolveParams) (interface{}, error) {
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
	jobPositionsInOrganizationUnitsResponse, err := getJobPositionsInOrganizationUnits(&input)
	if err != nil {
		return &jobPositionsInOrganizationUnitsResponse, err
	}

	for _, jobPositionsInOrganizationUnits := range jobPositionsInOrganizationUnitsResponse.Data {
		getJobPositionResponse, err := getJobPositionById(jobPositionsInOrganizationUnits.JobPositionId)
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

var JobPositionInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobPositions
	var jobPositionResponse *dto.GetJobPositionResponseMS
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		jobPositionResponse, err = updateJobPositions(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	} else {
		jobPositionResponse, err = createJobPositions(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    jobPositionResponse.Data,
	}, nil
}

var JobPositionDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	if !shared.IsInteger(itemId) && !(itemId.(int) <= 0) {
		return shared.ErrorResponse("You must pass the item id"), nil
	}

	err := deleteJobPositions(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

var JobPositionInOrganizationUnitInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobPositionsInOrganizationUnits
	var jobPositionInOrganizationUnit *dto.GetJobPositionInOrganizationUnitsResponseMS
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)
	if data.Id > 0 {
		jobPositionInOrganizationUnit, err = updateJobPositionsInOrganizationUnits(&data)
	} else {
		jobPositionInOrganizationUnit, err = createJobPositionsInOrganizationUnits(&data)
	}

	if err != nil {
		return shared.HandleAPIError(err)
	}

	err = deleteEmployeeInOrganizationUnit(jobPositionInOrganizationUnit.Data.Id)

	if len(data.Employees) > 0 {

		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, employeeId := range data.Employees {
			input := &structs.EmployeesInOrganizationUnits{
				PositionInOrganizationUnitId: jobPositionInOrganizationUnit.Data.Id,
				UserProfileId:                employeeId,
			}
			res, err := createEmployeesInOrganizationUnits(input)
			if err != nil {
				return shared.HandleAPIError(err)
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

var JobPositionInOrganizationUnitResolver = func(params graphql.ResolveParams) (interface{}, error) {
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

		systematizationsResponse, err := getSystematizations(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		if len(systematizationsResponse.Data) > 0 {
			for _, systematization := range systematizationsResponse.Data {
				input := dto.GetJobPositionInOrganizationUnitsInput{
					OrganizationUnitID: &officeUnitId,
					SystematizationID:  &systematization.Id,
				}
				jobPositionsInOrganizationUnits, err := getJobPositionsInOrganizationUnits(&input)
				if err != nil {
					return nil, err
				}
				for _, jobPositionOU := range jobPositionsInOrganizationUnits.Data {
					input := dto.GetEmployeesInOrganizationUnitInput{
						PositionInOrganizationUnit: &jobPositionOU.Id,
					}
					employeesInOrganizationUnit, _ := getEmployeesInOrganizationUnitList(&input)

					if len(employeesInOrganizationUnit) < jobPositionOU.AvailableSlots {
						jobPosition, err := getJobPositionById(jobPositionOU.JobPositionId)
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
var JobPositionInOrganizationUnitDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	if !shared.IsInteger(itemId) && !(itemId.(int) <= 0) {
		return shared.ErrorResponse("You must pass the item id"), nil
	}

	err := deleteJobPositionsInOrganizationUnits(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

func getJobPositionById(id int) (*structs.JobPositions, error) {
	res := &dto.GetJobPositionResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_POSITIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getJobPositions(input *dto.GetJobPositionsInput) (*dto.GetJobPositionsResponseMS, error) {
	res := &dto.GetJobPositionsResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_POSITIONS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func updateJobPositions(id int, data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error) {
	res := &dto.GetJobPositionResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.JOB_POSITIONS_ENDPOINT+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func createJobPositions(data *structs.JobPositions) (*dto.GetJobPositionResponseMS, error) {
	res := &dto.GetJobPositionResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.JOB_POSITIONS_ENDPOINT, data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func deleteJobPositions(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.JOB_POSITIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func createJobPositionsInOrganizationUnits(data *structs.JobPositionsInOrganizationUnits) (*dto.GetJobPositionInOrganizationUnitsResponseMS, error) {
	res := &dto.GetJobPositionInOrganizationUnitsResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.JOB_POSITIONS_IN_ORGANIZATION_UNITS_ENDPOINT, data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func updateJobPositionsInOrganizationUnits(data *structs.JobPositionsInOrganizationUnits) (*dto.GetJobPositionInOrganizationUnitsResponseMS, error) {
	res := &dto.GetJobPositionInOrganizationUnitsResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.JOB_POSITIONS_IN_ORGANIZATION_UNITS_ENDPOINT+"/"+strconv.Itoa(data.Id), data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getJobPositionsInOrganizationUnitsById(id int) (*structs.JobPositionsInOrganizationUnits, error) {
	res := &dto.GetJobPositionInOrganizationUnitsResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_POSITIONS_IN_ORGANIZATION_UNITS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteJobPositionsInOrganizationUnits(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.JOB_POSITIONS_IN_ORGANIZATION_UNITS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getJobPositionsInOrganizationUnits(input *dto.GetJobPositionInOrganizationUnitsInput) (*dto.GetJobPositionsInOrganizationUnitsResponseMS, error) {
	res := &dto.GetJobPositionsInOrganizationUnitsResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_POSITIONS_IN_ORGANIZATION_UNITS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func createEmployeesInOrganizationUnits(data *structs.EmployeesInOrganizationUnits) (*structs.EmployeesInOrganizationUnits, error) {
	res := &dto.GetEmployeesInOrganizationUnitsResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.EMPLOYEES_IN_ORGANIZATION_UNITS_ENDPOINT, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteEmployeeInOrganizationUnit(jobPositionInOrganizationUnitId int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.EMPLOYEES_IN_ORGANIZATION_UNITS_ENDPOINT+"/"+strconv.Itoa(jobPositionInOrganizationUnitId), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
