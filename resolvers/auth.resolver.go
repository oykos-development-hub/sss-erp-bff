package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"errors"
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

	userProfile, err := GetUserProfileByUserAccountID(loginRes.Data.Id)
	if err != nil {
		return shared.HandleAPIError(err)
	}

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
		systematization, _ := getSystematizationById(jobPositionInOrganizationUnit.SystematizationId)
		organizationUnit, _ = getOrganizationUnitById(systematization.OrganizationUnitId)
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

func RefreshTokenResolver(p graphql.ResolveParams) (interface{}, error) {
	request, ok := p.Context.Value(config.Requestkey).(*http.Request)
	if !ok {
		return shared.HandleAPIError(errors.New("could not get request from context"))
	}

	refreshCookie, err := request.Cookie("refresh_token")
	if err != nil {
		return shared.HandleAPIError(err)
	}

	refreshRes, cookies, err := refreshToken(refreshCookie)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	httpResponseWriter := p.Context.Value((config.HttpResponseWriterKey)).(http.ResponseWriter)
	for _, cookie := range cookies {
		http.SetCookie(httpResponseWriter, cookie)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Token refreshed successfully",
		Item:    refreshRes.Data, // Assuming that item holds the same data structure as data
	}, nil
}

var LogoutResolver = func(p graphql.ResolveParams) (interface{}, error) {
	var authToken = p.Context.Value(config.TokenKey).(string)

	err := logout(authToken)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "user logged out",
	}, nil

}

func logout(token string) error {
	_, err := shared.MakeAPIRequest("POST", config.LOGOUT_ENDPOINT, nil, nil, map[string]string{"Authorization": "Bearer " + token})
	if err != nil {
		return err
	}

	return nil
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

func refreshToken(cookie *http.Cookie) (*dto.RefreshTokenResponse, []*http.Cookie, error) {
	refreshResponse := &dto.RefreshTokenResponse{}
	cookies, err := shared.MakeAPIRequestWithCookie("GET", config.REFRESH_ENDPOINT, nil, refreshResponse, cookie)
	if err != nil {
		return nil, nil, err
	}

	return refreshResponse, cookies, nil
}
