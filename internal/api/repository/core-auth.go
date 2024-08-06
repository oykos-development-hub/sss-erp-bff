package repository

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"context"
	"net/http"
	"strconv"
)

func (repo *MicroserviceRepository) Logout(token string) error {
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.Logout, nil, nil, map[string]string{"Authorization": "Bearer " + token})
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) ForgotPassword(email string) error {
	reqBody := dto.ResetRequestMS{
		Email: email,
	}
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.ForgotPassword, reqBody, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}
	return nil
}

func (repo *MicroserviceRepository) ValidateMail(input *dto.ResetPasswordVerify) (*dto.ResetPasswordVerifyResponseMS, error) {
	res := &dto.ResetPasswordVerifyResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.ValidateMail, input, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}
	return res, nil
}

func (repo *MicroserviceRepository) ResetPassword(input *dto.ResetPassword) error {
	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.ResetPassword, input, nil)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}
	return nil
}

func (repo *MicroserviceRepository) LoginUser(email, password string) (*dto.LoginResponseMS, []*http.Cookie, error) {
	reqBody := dto.LoginRequestMS{
		Email:    email,
		Password: password,
	}

	loginResponse := &dto.LoginResponseMS{}
	cookies, err := makeAPIRequest("POST", repo.Config.Microservices.Core.Login, reqBody, loginResponse)
	if err != nil {
		return nil, nil, errors.Wrap(err, "make api request")
	}

	return loginResponse, cookies, nil
}

func (repo *MicroserviceRepository) RefreshToken(cookie *http.Cookie) (*dto.RefreshTokenResponse, []*http.Cookie, error) {
	refreshResponse := &dto.RefreshTokenResponse{}
	cookies, err := makeAPIRequestWithCookie("GET", repo.Config.Microservices.Core.Refresh, nil, refreshResponse, cookie)
	if err != nil {
		return nil, nil, errors.Wrap(err, "make api request")
	}

	return refreshResponse, cookies, nil
}

func (repo *MicroserviceRepository) GetRole(id int) (*structs.Roles, error) {
	res := &dto.GetRoleResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Roles+"/"+strconv.Itoa(int(id)), nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) GetRoleList() ([]structs.Roles, error) {
	res := &dto.GeRoleListResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.Roles, nil, res)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return res.Data, nil
}

func (repo *MicroserviceRepository) UpdateRole(ctx context.Context, id int, data structs.Roles) (*structs.Roles, error) {
	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	res := &dto.GetRoleResponseMS{}
	_, err := makeAPIRequest("PUT", repo.Config.Microservices.Core.Roles+"/"+strconv.Itoa(id), data, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) CreateRole(ctx context.Context, data structs.Roles) (*structs.Roles, error) {
	res := &dto.GetRoleResponseMS{}

	header := make(map[string]string)
	account := ctx.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	header["UserID"] = strconv.Itoa(account.ID)

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.Roles, data, res, header)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}

func (repo *MicroserviceRepository) ValidatePin(pin string, headers map[string]string) error {
	reqBody := dto.PinRequestMS{
		Pin: pin,
	}

	_, err := makeAPIRequest("POST", repo.Config.Microservices.Core.Pin, reqBody, nil, headers)
	if err != nil {
		return errors.Wrap(err, "make api request")
	}

	return nil
}

func (repo *MicroserviceRepository) AuthenticateUser(r *http.Request) (*structs.UserAccounts, error) {
	token := r.URL.Query().Get("token")

	loggedInAccount, err := repo.GetLoggedInUser(token)
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return loggedInAccount, nil
}

func (repo *MicroserviceRepository) GetLoggedInUser(token string) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	_, err := makeAPIRequest("GET", repo.Config.Microservices.Core.LoggedInUser, nil, res, map[string]string{"Authorization": "Bearer " + token})
	if err != nil {
		return nil, errors.Wrap(err, "make api request")
	}

	return &res.Data, nil
}
