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

	jobPositionInOrganizationUnit, err = createJobPositionsInOrganizationUnits(&data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You updated this item!",
		Item:    jobPositionInOrganizationUnit.Data,
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

var EmployeeInOrganizationUnitInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.EmployeesInOrganizationUnits
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}
	_ = json.Unmarshal(dataBytes, &data)

	userProfile, err := getUserProfileById(data.UserProfileId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	data.UserAccountId = userProfile.UserAccountId

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateEmployeesInOrganizationUnits(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = res
		response.Message = "You updated this item!"
	} else {
		res, err := createEmployeesInOrganizationUnits(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = res
		response.Message = "You created this item!"
	}

	return response, nil
}

var EmployeeInOrganizationUnitDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := deleteEmployeeInOrganizationUnit(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
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

func updateEmployeesInOrganizationUnits(id int, data *structs.EmployeesInOrganizationUnits) (*structs.EmployeesInOrganizationUnits, error) {
	res := &dto.GetEmployeesInOrganizationUnitsResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.EMPLOYEES_IN_ORGANIZATION_UNITS_ENDPOINT+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteEmployeeInOrganizationUnit(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.EMPLOYEES_IN_ORGANIZATION_UNITS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
