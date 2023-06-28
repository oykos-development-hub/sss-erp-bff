package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

var UserProfileResolutionResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var resolutionItems []interface{}

	profileId := params.Args["user_profile_id"]
	accountId := params.Args["user_account_id"]

	if !shared.IsInteger(profileId) && !shared.IsInteger(accountId) {
		return map[string]interface{}{
			"status":  "error",
			"message": "Argument 'user_profile_id' must not be empty!",
			"item":    nil,
		}, nil
	}

	UserProfilesType := &structs.UserProfiles{}
	UserProfilesData, UserProfilesDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profiles.json", UserProfilesType)

	if UserProfilesDataErr != nil {
		fmt.Printf("Fetching User Profiles failed because of this error - %s.\n", UserProfilesDataErr)
	}

	var UserProfile = shared.FindByProperty(UserProfilesData, "Id", profileId)

	if UserProfile == nil || UserProfile[0] == nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "User Profile not found for provided 'user_profile_id'!",
			"item":    nil,
		}, nil
	}

	var resolutionTypes = shared.FetchByProperty(
		"resolution_type",
		"",
		"",
	)

	var relatedResolution = shared.FetchByProperty(
		"resolution",
		"UserProfileId",
		profileId,
	)
	// # Related Employee Resolution
	if len(relatedResolution) > 0 {
		for _, relatedResolutionItem := range relatedResolution {
			var relatedResolutionItemData = shared.WriteStructToInterface(relatedResolutionItem)
			var relatedResolutionType = shared.FindByProperty(resolutionTypes, "Id", relatedResolutionItemData["resolution_type_id"])

			if len(relatedResolutionType) > 0 {
				var relatedResolutionData = shared.WriteStructToInterface(relatedResolutionType[0])

				relatedResolutionItemData["resolution_type"] = map[string]interface{}{
					"title": relatedResolutionData["title"],
					"id":    relatedResolutionData["id"],
				}
			}

			resolutionItems = append(resolutionItems, relatedResolutionItemData)
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here are the items you asked for!",
		"items":   resolutionItems,
	}, nil
}

var UserProfileResolutionInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.Resolution
	dataBytes, _ := json.Marshal(params.Args["data"])
	ResolutionType := &structs.Resolution{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	resolutionData, resolutionDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profile_resolutions.json", ResolutionType)

	if resolutionDataErr != nil {
		fmt.Printf("Fetching User Profile's resolution failed because of this error - %s.\n", resolutionDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		resolutionData = shared.FilterByProperty(resolutionData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(resolutionData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_resolutions.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var UserProfileResolutionDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	ResolutionType := &structs.Resolution{}
	resolutionData, resolutionDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profile_resolutions.json", ResolutionType)

	if resolutionDataErr != nil {
		fmt.Printf("Fetching User Profile's Resolution failed because of this error - %s.\n", resolutionDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		resolutionData = shared.FilterByProperty(resolutionData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_resolutions.json"), resolutionData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
