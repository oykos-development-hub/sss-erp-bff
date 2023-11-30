package websocketmanager

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"net/http"
)

func authenticateUser(r *http.Request) (*structs.UserAccounts, error) {
	token := r.URL.Query().Get("token")

	loggedInAccount, err := GetLoggedInUser(token)
	if err != nil {
		return nil, err
	}

	return loggedInAccount, nil
}

func GetLoggedInUser(token string) (*structs.UserAccounts, error) {
	res := &dto.GetUserAccountResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.LOGGED_IN_USER_ENDPOINT, nil, res, map[string]string{"Authorization": "Bearer " + token})
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
