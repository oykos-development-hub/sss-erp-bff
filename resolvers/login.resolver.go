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

	var (
		engagement       *structs.SettingsDropdown
		contractsResItem *structs.Contracts
		jobPosition      *structs.JobPositions
		organizationUnit *structs.OrganizationUnits
	)

	loginRes, cookies, err := loginUser(email, password)
	if err != nil {
		return shared.HandleAPIError(err)
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

	userProfile, _ := getUserProfileByUserAccountID(loginRes.Data.Id)

	isActive := true
	contracts, _ := getEmployeeContracts(userProfile.Id, &dto.GetEmployeeContracts{Active: &isActive})

	if len(contracts) != 1 {
		fmt.Printf("employee must have exactly one active contract assigned")
	} else {
		contractsResItem, _ = buildContractResponseItem(*contracts[0])
	}

	if userProfile.EngagementTypeId != nil {
		engagement, err = getDropdownSettingById(*userProfile.EngagementTypeId)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	employeesInOrganizationUnit, _ := getEmployeesInOrganizationUnitsByProfileId(userProfile.Id)
	if employeesInOrganizationUnit != nil {
		jobPositionInOrganizationUnit, _ := getJobPositionsInOrganizationUnitsById(employeesInOrganizationUnit.PositionInOrganizationUnitId)

		jobPosition, _ = getJobPositionById(jobPositionInOrganizationUnit.JobPositionId)
		organizationUnit, _ = getOrganizationUnitById(jobPositionInOrganizationUnit.ParentOrganizationUnitId)
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
		Contract:            contractsResItem,
		Engagement:          engagement,
		JobPosition:         jobPosition,
		OrganizationUnit:    organizationUnit,
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
