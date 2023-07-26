package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
)

var LoginResolver = func(p graphql.ResolveParams) (interface{}, error) {
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)

	loginRes, cookies, err := loginUser(email, password)
	if err != nil {
		return dto.LoginResponse{
			Status:  "error",
			Message: err.Error(),
		}, nil
	}

	httpResponseWriter := p.Context.Value((config.HttpResponseWriterKey)).(http.ResponseWriter)
	for _, cookie := range cookies {
		http.SetCookie(httpResponseWriter, cookie)
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
		contractsData = []interface{}{}
	}

	EngagementsType := &structs.EngagementType{}
	engagementsData, engagementsDataErr := shared.ReadJson(shared.GetDataRoot()+"/engagement_officer_type.json", EngagementsType)

	if engagementsDataErr != nil {
		fmt.Printf("Fetching engagements failed because of this error - %s.\n", engagementsDataErr)
		engagementsData = []interface{}{}
	}

	JobPositionType := &structs.JobPositions{}
	jobPositionData, jobPositionDataErr := shared.ReadJson(shared.GetDataRoot()+"/job_position_it_admin.json", JobPositionType)

	if jobPositionDataErr != nil {
		fmt.Printf("Fetching job positions failed because of this error - %s.\n", jobPositionDataErr)
		jobPositionData = []interface{}{}
	}

	OrganizationUnitType := &structs.OrganizationUnits{}
	organizationUnitData, organizationUnitDataErr := shared.ReadJson(shared.GetDataRoot()+"/organization_unit_sss.json", OrganizationUnitType)

	if organizationUnitDataErr != nil {
		fmt.Printf("Fetching organization units failed because of this error - %s.\n", organizationUnitDataErr)
		organizationUnitData = []interface{}{}
	}

	return dto.LoginResponse{
		Status:              "success",
		Message:             "Welcome!",
		RoleId:              123,
		FolderId:            456,
		Email:               email,
		Phone:               "555-555-1234",
		Token:               loginRes.Data.Token.Token,
		CreatedAt:           "2023-04-28T10:45:00Z",
		FirstName:           "John",
		LastName:            "Doe",
		BirthLastName:       "Smith",
		Gender:              "Male",
		DateOfBecomingJudge: "2022-01-01",
		Permissions:         permissionsData,
		Contract:            contractsData[0],
		Engagement:          engagementsData[0],
		JobPosition:         jobPositionData[0],
		OrganizationUnit:    organizationUnitData[0],
	}, nil
}

func loginUser(email, password string) (*dto.LoginResponseMS, []*http.Cookie, error) {
	reqBody := dto.LoginRequestMS{
		Email:    email,
		Password: password,
	}

	loginResponse := &dto.LoginResponseMS{}
	cookies, err := shared.MakeAPIRequest("POST", config.LOGIN_ENDPOINT, reqBody, loginResponse)
	if err != nil {
		return nil, nil, err
	}

	return loginResponse, cookies, nil
}
