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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
	}

	var roleID int
	if loginRes.Data.RoleID != nil {
		roleID = int(*loginRes.Data.RoleID)
	}
	httpResponseWriter := p.Context.Value((config.HTTPResponseWriterKey)).(http.ResponseWriter)
	for _, cookie := range cookies {
		http.SetCookie(httpResponseWriter, cookie)
	}

	permissions, err := r.Repo.GetPermissionList(roleID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
	}

	userProfile, err := r.Repo.GetUserProfileByUserAccountID(loginRes.Data.ID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
	}

	isActive := true
	contracts, _ := r.Repo.GetEmployeeContracts(userProfile.ID, &dto.GetEmployeeContracts{Active: &isActive})

	if len(contracts) > 0 {
		contractsResItem, _ = buildContractResponseItem(r.Repo, *contracts[0])
	}

	if userProfile.EngagementTypeID != nil {
		engagement, err = r.Repo.GetDropdownSettingByID(*userProfile.EngagementTypeID)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return apierrors.HandleAPPError(err)
		}
	}

	employeesInOrganizationUnit, _ := r.Repo.GetEmployeesInOrganizationUnitsByProfileID(userProfile.ID)
	if employeesInOrganizationUnit != nil {
		jobPositionInOrganizationUnit, err := r.Repo.GetJobPositionsInOrganizationUnitsByID(employeesInOrganizationUnit.PositionInOrganizationUnitID)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return apierrors.HandleAPPError(err)
		}

		jobPosition, err = r.Repo.GetJobPositionByID(jobPositionInOrganizationUnit.JobPositionID)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return apierrors.HandleAPPError(err)
		}

		systematization, err := r.Repo.GetSystematizationByID(jobPositionInOrganizationUnit.SystematizationID)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return apierrors.HandleAPPError(err)
		}

		organizationUnit, err = r.Repo.GetOrganizationUnitByID(systematization.OrganizationUnitID)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return apierrors.HandleAPPError(err)
		}
	}

	active := true
	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}

	resolution, err := r.Repo.GetJudgeResolutionList(&input)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
	}

	if len(resolution.Data) > 0 {

		filter := dto.JudgeResolutionsOrganizationUnitInput{
			ResolutionID:  &resolution.Data[0].ID,
			UserProfileID: &userProfile.ID,
		}

		judges, _, err := r.Repo.GetJudgeResolutionOrganizationUnit(&filter)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return apierrors.HandleAPPError(err)
		}

		if len(judges) > 0 {
			organizationUnit, err = r.Repo.GetOrganizationUnitByID(judges[0].OrganizationUnitID)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return apierrors.HandleAPPError(err)
			}
		}

	}

	var organizationUnitList []dto.OrganizationUnitsOverviewResponse

	loggedInAccount := p.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)

	hasPermission, err := r.HasPermission(*loggedInAccount, string(config.HR), config.OperationFullAccess)

	if err != nil {
		return nil, apierrors.Wrap(err, "repo has permission")
	}

	if hasPermission {
		isParent := true
		organizationUnits, err := r.Repo.GetOrganizationUnits(&dto.GetOrganizationUnitsInput{IsParent: &isParent})
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return apierrors.HandleAPPError(err)
		}

		for _, organizationUnit := range organizationUnits.Data {
			organizationUnitItem, err := buildOrganizationUnitOverviewResponse(r.Repo, &organizationUnit)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return apierrors.HandleAPPError(err)
			}
			organizationUnitList = append(organizationUnitList, *organizationUnitItem)
		}
	}

	supplierListRes, err := r.Repo.GetSupplierList(nil)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Password reset successful",
	}, nil
}

func (r *Resolver) RefreshTokenResolver(p graphql.ResolveParams) (interface{}, error) {
	request, ok := p.Context.Value(config.Requestkey).(*http.Request)
	if !ok {
		return apierrors.HandleAPPError(errors.New("could not get request from context"))
	}

	refreshCookie, err := request.Cookie("refresh_token")
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
	}

	refreshRes, cookies, err := r.Repo.RefreshToken(refreshCookie)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
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
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return apierrors.HandleAPPError(err)
	}

	r.NotificationsService.Wsmanager.RemoveClientByUserID(user.ID)

	return dto.ResponseSingle{
		Status:  "success",
		Message: "user logged out",
	}, nil

}
