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
		contractsResItem *dto.Contract
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

	if len(contracts) > 0 {
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
		Id:                  userProfile.Id,
		RoleId:              loginRes.Data.RoleId,
		FolderId:            0,
		Email:               loginRes.Data.Email,
		Phone:               loginRes.Data.Phone,
		Token:               loginRes.Data.Token.Token,
		CreatedAt:           userProfile.CreatedAt,
		FirstName:           userProfile.FirstName,
		LastName:            userProfile.LastName,
		BirthLastName:       userProfile.BirthLastName,
		Gender:              userProfile.Gender,
		DateOfBecomingJudge: userProfile.DateOfBecomingJudge,
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
