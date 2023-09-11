package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"
)

var UserProfileResolutionResolver = func(params graphql.ResolveParams) (interface{}, error) {
	userProfileId := params.Args["user_profile_id"].(int)

	resolutions, err := getEmployeeResolutions(userProfileId, nil)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	resolutonResItemList, err := buildResolutionResponseItemList(resolutions)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   resolutonResItemList,
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
		resolution, err := updateResolution(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resolutionResItem, err := buildResolutionResItem(resolution)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = resolutionResItem
		response.Message = "You updated this item!"
	} else {
		resolution, err := createResolution(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resolutionResItem, err := buildResolutionResItem(resolution)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = resolutionResItem
		response.Message = "You created this item!"
	}

	return response, nil
}

var UserProfileResolutionDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteResolution(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildResolutionResponseItemList(items []*structs.Resolution) (resItemList []*dto.Resolution, err error) {
	for _, item := range items {
		resItem, err := buildResolutionResItem(item)
		if err != nil {
			return nil, err
		}
		resItemList = append(resItemList, resItem)
	}
	return
}

func buildResolutionResItem(item *structs.Resolution) (*dto.Resolution, error) {
	userProfile, err := getUserProfileById(item.UserProfileId)
	if err != nil {
		return nil, err
	}
	resolutionType, err := getDropdownSettingById(item.ResolutionTypeId)
	if err != nil {
		return nil, err
	}

	return &dto.Resolution{
		Id:                item.Id,
		ResolutionPurpose: item.ResolutionPurpose,
		UserProfile: dto.DropdownSimple{
			Id:    userProfile.Id,
			Title: userProfile.GetFullName(),
		},
		ResolutionType: dto.DropdownSimple{
			Id:    resolutionType.Id,
			Title: resolutionType.Title,
		},
		DateOfStart: item.DateOfStart,
		DateOfEnd:   item.DateOfEnd,
		Value:       item.Value,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}, nil
}

func getEmployeeResolutions(employeeID int, input *dto.EmployeeResolutionListInput) ([]*structs.Resolution, error) {
	res := &dto.GetResolutionListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(employeeID)+"/resolutions", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getEmployeeResolution(id int) (*structs.Resolution, error) {
	res := &dto.GetResolutionResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.RESOLUTIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
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
