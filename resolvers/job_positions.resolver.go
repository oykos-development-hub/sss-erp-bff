package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

var JobPositionsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	search := params.Args["search"]
	JobPositionType := &structs.JobPositions{}
	JobPositionData, JobPositionDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_positions.json", JobPositionType)

	if JobPositionDataErr != nil {
		fmt.Printf("Fetching Job Positions failed because of this error - %s.\n", JobPositionDataErr)
	}

	if id != nil && shared.IsInteger(id) && id != 0 {
		JobPositionData = shared.FindByProperty(JobPositionData, "Id", id)
	}

	if search != nil && shared.IsString(search) {
		JobPositionData = shared.FindByProperty(JobPositionData, "Title", search, true)
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the list you asked for!",
		"items":   JobPositionData,
	}, nil
}

var JobPositionInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.JobPositions
	dataBytes, _ := json.Marshal(params.Args["data"])
	JobPositionType := &structs.JobPositions{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	JobPositionData, JobPositionDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_positions.json", JobPositionType)

	if JobPositionDataErr != nil {
		fmt.Printf("Fetching Job Positions failed because of this error - %s.\n", JobPositionDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		JobPositionData = shared.FilterByProperty(JobPositionData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(JobPositionData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/job_positions.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var JobPositionDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	JobPositionType := &structs.JobPositions{}
	JobPositionData, JobPositionDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_positions.json", JobPositionType)

	if JobPositionDataErr != nil {
		fmt.Printf("Fetching Job Positions failed because of this error - %s.\n", JobPositionDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		JobPositionData = shared.FilterByProperty(JobPositionData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/job_positions.json"), JobPositionData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

var JobPositionInOrganizationUnitInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.JobPositionsInOrganizationUnits
	dataBytes, _ := json.Marshal(params.Args["data"])
	JobPositionType := &structs.JobPositionsInOrganizationUnits{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	JobPositionData, JobPositionDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_positions_in_organization_units.json", JobPositionType)

	if JobPositionDataErr != nil {
		fmt.Printf("Fetching Job Positions failed because of this error - %s.\n", JobPositionDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		JobPositionData = shared.FilterByProperty(JobPositionData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(JobPositionData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/job_positions_in_organization_units.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var JobPositionInOrganizationUnitDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	JobPositionType := &structs.JobPositionsInOrganizationUnits{}
	JobPositionData, JobPositionDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_positions_in_organization_units.json", JobPositionType)

	if JobPositionDataErr != nil {
		fmt.Printf("Fetching Job Positions failed because of this error - %s.\n", JobPositionDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		JobPositionData = shared.FilterByProperty(JobPositionData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/job_positions_in_organization_units.json"), JobPositionData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}

var EmployeeInOrganizationUnitInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.EmployeesInOrganizationUnits
	dataBytes, _ := json.Marshal(params.Args["data"])
	EmployeeType := &structs.EmployeesInOrganizationUnits{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	userProfileId := data.UserProfileId
	EmployeeData, EmployeeDataErr := shared.ReadJson(shared.GetDataRoot()+"/employees_in_organization_units.json", EmployeeType)

	if EmployeeDataErr != nil {
		fmt.Printf("Fetching Employees failed because of this error - %s.\n", EmployeeDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		EmployeeData = shared.FilterByProperty(EmployeeData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	// # Related User Profile
	var relatedUserProfile = shared.FetchByProperty(
		"user_profile",
		"Id",
		userProfileId,
	)

	if len(relatedUserProfile) > 0 {
		userProfile := shared.WriteStructToInterface(relatedUserProfile[0])
		data.UserAccountId = userProfile["user_account_id"].(int)
	}

	var updatedData = append(EmployeeData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/employees_in_organization_units.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var EmployeeInOrganizationUnitDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	EmployeeType := &structs.EmployeesInOrganizationUnits{}
	EmployeeData, EmployeeDataErr := shared.ReadJson(shared.GetDataRoot()+"/employees_in_organization_units.json", EmployeeType)

	if EmployeeDataErr != nil {
		fmt.Printf("Fetching Employees failed because of this error - %s.\n", EmployeeDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		EmployeeData = shared.FilterByProperty(EmployeeData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/employees_in_organization_units.json"), EmployeeData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
