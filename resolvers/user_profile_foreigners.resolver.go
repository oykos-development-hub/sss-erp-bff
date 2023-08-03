package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

var UserProfileForeignerResolver = func(params graphql.ResolveParams) (interface{}, error) {
	profileId := params.Args["user_profile_id"].(int)

	UserProfilesData, err := getEmployeeForeigners(profileId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   UserProfilesData,
	}, nil
}

var UserProfileForeignerInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var err error

	var data structs.Foreigners
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating settings data"), nil
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		item, err := updateEmployeeForeigner(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := createEmployeeForeigner(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

var UserProfileForeignerDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := deleteForeigner(itemId.(int))
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func getEmployeeForeigners(userProfileID int) ([]*structs.Foreigners, error) {
	res := dto.GetEmployeeForeignersListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(userProfileID)+"/foreigners", nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func updateEmployeeForeigner(id int, foreigner *structs.Foreigners) (*structs.Foreigners, error) {
	res := dto.GetEmployeeForeignersResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.FOREIGNERS+"/"+strconv.Itoa(id), foreigner, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func createEmployeeForeigner(foreigner *structs.Foreigners) (*structs.Foreigners, error) {
	res := dto.GetEmployeeForeignersResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.FOREIGNERS, foreigner, &res)
	//foreigners
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func deleteForeigner(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.FOREIGNERS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
