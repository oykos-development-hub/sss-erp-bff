package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/structs"
	"errors"
	"net/http"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) LoginResolver(p graphql.ResolveParams) (interface{}, error) {
	email := p.Args["email"].(string)
	password := p.Args["password"].(string)

	var (
		engagement       *structs.SettingsDropdown
		contractsResItem *dto.Contract
		jobPosition      *structs.JobPositions
		organizationUnit *structs.OrganizationUnits
	)

	loginRes, cookies, err := r.Repo.LoginUser(email, password)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	roleID := int(loginRes.Data.RoleID)

	httpResponseWriter := p.Context.Value((config.HTTPResponseWriterKey)).(http.ResponseWriter)
	for _, cookie := range cookies {
		http.SetCookie(httpResponseWriter, cookie)
	}

	permissions, err := r.Repo.GetPermissionList(roleID)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	userProfile, err := r.Repo.GetUserProfileByUserAccountID(loginRes.Data.ID)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	isActive := true
	contracts, _ := r.Repo.GetEmployeeContracts(userProfile.ID, &dto.GetEmployeeContracts{Active: &isActive})

	if len(contracts) > 0 {
		contractsResItem, _ = buildContractResponseItem(r.Repo, *contracts[0])
	}

	if userProfile.EngagementTypeID != nil {
		engagement, err = r.Repo.GetDropdownSettingByID(*userProfile.EngagementTypeID)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	}

	employeesInOrganizationUnit, _ := r.Repo.GetEmployeesInOrganizationUnitsByProfileID(userProfile.ID)
	if employeesInOrganizationUnit != nil {
		jobPositionInOrganizationUnit, _ := r.Repo.GetJobPositionsInOrganizationUnitsByID(employeesInOrganizationUnit.PositionInOrganizationUnitID)

		jobPosition, _ = r.Repo.GetJobPositionByID(jobPositionInOrganizationUnit.JobPositionID)
		systematization, _ := r.Repo.GetSystematizationByID(jobPositionInOrganizationUnit.SystematizationID)
		organizationUnit, _ = r.Repo.GetOrganizationUnitByID(systematization.OrganizationUnitID)
	}

	var organizationUnitList []dto.OrganizationUnitsOverviewResponse

	if loginRes.Data.HasPermission(structs.PermissionManageOrganizationUnits) {
		isParent := true
		organizationUnits, err := r.Repo.GetOrganizationUnits(&dto.GetOrganizationUnitsInput{IsParent: &isParent})
		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		for _, organizationUnit := range organizationUnits.Data {
			organizationUnitItem, err := buildOrganizationUnitOverviewResponse(r.Repo, &organizationUnit)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
			organizationUnitList = append(organizationUnitList, *organizationUnitItem)
		}
	}

	supplierListRes, err := r.Repo.GetSupplierList(nil)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return dto.LoginResponse{
		Status:               "success",
		Message:              "Welcome!",
		ID:                   userProfile.ID,
		RoleID:               roleID,
		FolderID:             0,
		Email:                loginRes.Data.Email,
		Phone:                loginRes.Data.Phone,
		Token:                loginRes.Data.Token.Token,
		CreatedAt:            userProfile.CreatedAt,
		FirstName:            userProfile.FirstName,
		LastName:             userProfile.LastName,
		BirthLastName:        userProfile.BirthLastName,
		Gender:               userProfile.Gender,
		DateOfBecomingJudge:  userProfile.DateOfBecomingJudge,
		Permissions:          permissions,
		Contract:             contractsResItem,
		Engagement:           engagement,
		JobPosition:          jobPosition,
		OrganizationUnit:     organizationUnit,
		OrganizationUnitList: organizationUnitList,
		SupplierList:         supplierListRes.Data,
	}, nil
}

func (r *Resolver) ForgotPasswordResolver(p graphql.ResolveParams) (interface{}, error) {
	email := p.Args["email"].(string)

	err := r.Repo.ForgotPassword(email)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Mail send successful",
	}, nil
}

func (r *Resolver) UserValidateMailResolver(p graphql.ResolveParams) (interface{}, error) {
	email := p.Args["email"].(string)
	token := p.Args["token"].(string)

	input := dto.ResetPasswordVerify{
		Email: email,
		Token: token,
	}

	res, err := r.Repo.ValidateMail(&input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:         "success",
		Message:        "Email is valid",
		EncryptedEmail: res.Data.EncryptedEmail,
	}, nil
}

func (r *Resolver) UserResetPasswordResolver(p graphql.ResolveParams) (interface{}, error) {
	encryptedEmail := p.Args["encrypted_email"].(string)
	password := p.Args["password"].(string)

	input := dto.ResetPassword{
		EncryptedEmail: encryptedEmail,
		Password:       password,
	}

	err := r.Repo.ResetPassword(&input)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Password reset successful",
	}, nil
}

func (r *Resolver) RefreshTokenResolver(p graphql.ResolveParams) (interface{}, error) {
	request, ok := p.Context.Value(config.Requestkey).(*http.Request)
	if !ok {
		return apierrors.HandleAPIError(errors.New("could not get request from context"))
	}

	refreshCookie, err := request.Cookie("refresh_token")
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	refreshRes, cookies, err := r.Repo.RefreshToken(refreshCookie)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	httpResponseWriter := p.Context.Value((config.HTTPResponseWriterKey)).(http.ResponseWriter)
	for _, cookie := range cookies {
		http.SetCookie(httpResponseWriter, cookie)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Token refreshed successfully",
		Item:    refreshRes.Data, // Assuming that item holds the same data structure as data
	}, nil
}

func (r *Resolver) LogoutResolver(p graphql.ResolveParams) (interface{}, error) {
	var authToken = p.Context.Value(config.TokenKey).(string)
	user := p.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

	err := r.Repo.Logout(authToken)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	r.NotificationsService.Wsmanager.RemoveClientByUserID(user.ID)

	return dto.ResponseSingle{
		Status:  "success",
		Message: "user logged out",
	}, nil

}
