package resolvers

import (
	"bff/shared"
	"bff/structs"
	"fmt"
	"github.com/graphql-go/graphql"
)

var LoginResolver = func(p graphql.ResolveParams) (interface{}, error) {
	email := p.Args["email"].(string)

	if shared.StringContains(email, "invalid") {
		return map[string]interface{}{
			"status":                 "error",
			"message":                "User account not found!",
			"role_id":                0,
			"folder_id":              0,
			"email":                  "",
			"phone":                  "",
			"token":                  "",
			"refresh_token":          "",
			"created_at":             "",
			"first_name":             "",
			"last_name":              "",
			"birth_last_name":        "",
			"gender":                 "",
			"date_of_becoming_judge": "",
			"permissions":            nil,
			"contract":               nil,
			"engagement":             nil,
			"job_position":           nil,
			"organization_unit":      nil,
		}, nil
	}

	PermissionsType := &structs.Permissions{}
	permissionsData, permissionsDataErr := shared.ReadJson(shared.GetDataRoot()+"/permissions_super_admin.json", PermissionsType)

	if permissionsDataErr != nil {
		fmt.Printf("Fetching permissions failed because of this error - %s.\n", permissionsDataErr)
	}

	ContractsType := &structs.Contracts{}
	contractsData, contractsDataErr := shared.ReadJson(shared.GetDataRoot()+"/contract_unlimited_type.json", ContractsType)

	if contractsDataErr != nil {
		fmt.Printf("Fetching contracts failed because of this error - %s.\n", contractsDataErr)
	}

	EngagementsType := &structs.EngagementType{}
	engagementsData, engagementsDataErr := shared.ReadJson(shared.GetDataRoot()+"/engagement_officer_type.json", EngagementsType)

	if engagementsDataErr != nil {
		fmt.Printf("Fetching engagements failed because of this error - %s.\n", engagementsDataErr)
	}

	JobPositionType := &structs.JobPositions{}
	jobPositionData, jobPositionDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_position_it_admin.json", JobPositionType)

	if jobPositionDataErr != nil {
		fmt.Printf("Fetching job positions failed because of this error - %s.\n", jobPositionDataErr)
	}

	OrganizationUnitType := &structs.OrganizationUnits{}
	organizationUnitData, organizationUnitDataErr := shared.ReadJson(shared.GetDataRoot()+"/organization_unit_sss.json", OrganizationUnitType)

	if organizationUnitDataErr != nil {
		fmt.Printf("Fetching organization units failed because of this error - %s.\n", organizationUnitDataErr)
	}

	return map[string]interface{}{
		"status":                 "success",
		"message":                "Welcome!",
		"role_id":                123,
		"folder_id":              456,
		"email":                  email,
		"phone":                  "555-555-1234",
		"token":                  "abc123",
		"refresh_token":          "def456",
		"created_at":             "2023-04-28T10:45:00Z",
		"first_name":             "John",
		"last_name":              "Doe",
		"birth_last_name":        "Smith",
		"gender":                 "Male",
		"date_of_becoming_judge": "2022-01-01",
		"permissions":            permissionsData,
		"contract":               contractsData[0],
		"engagement":             engagementsData[0],
		"job_position":           jobPositionData[0],
		"organization_unit":      organizationUnitData[0],
	}, nil
}
