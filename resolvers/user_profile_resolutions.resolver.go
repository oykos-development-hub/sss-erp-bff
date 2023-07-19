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

var UserProfileResolutionResolver = func(params graphql.ResolveParams) (interface{}, error) {
	userProfileId := params.Args["user_profile_id"].(int)

	resolutions, err := getEmployeeResolutions(userProfileId)
	if err != nil {
		return dto.Response{
			Status:  "error",
			Message: err.Error(),
		}, nil
	}
	items := shared.ConvertToInterfaceSlice(resolutions)

	_ = hydrateSettings("ResolutionType", "ResolutionTypeId", items...)

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   resolutions,
	}, nil
}

var UserProfileResolutionInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Resolution
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		resolutionResponse, err := updateResolution(itemId, &data)
		if err != nil {
			fmt.Printf("Updating employee's resolution failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error updating employee's resolution data"), nil
		}
		response.Item = resolutionResponse
		response.Message = "You updated this item!"
	} else {
		resolutionResponse, err := createResolution(&data)
		if err != nil {
			fmt.Printf("Creating employee's resolution failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating employee's resolution data"), nil
		}
		response.Item = resolutionResponse
		response.Message = "You created this item!"
	}

	return response, nil
}

var UserProfileResolutionDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteResolution(itemId)
	if err != nil {
		fmt.Printf("Deleting employee's resolution failed because of this error - %s.\n", err)
		return shared.ErrorResponse("Error deleting the resolution"), nil
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func getEmployeeResolutions(employeeID int) ([]*structs.Resolution, error) {
	res := &dto.GetResolutionListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(employeeID)+"/resolutions", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func updateResolution(id int, resolution *structs.Resolution) (*structs.Resolution, error) {
	res := &dto.GetResolutionResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.RESOLUTIONS_ENDPOINT+"/"+strconv.Itoa(id), resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createResolution(resolution *structs.Resolution) (*structs.Resolution, error) {
	res := &dto.GetResolutionResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.RESOLUTIONS_ENDPOINT, resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteResolution(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.RESOLUTIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
