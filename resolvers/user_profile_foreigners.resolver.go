package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

var UserProfileForeignerResolver = func(params graphql.ResolveParams) (interface{}, error) {
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
	UserProfilesData, UserProfilesDataErr := shared.ReadJson(shared.GetDataRoot()+"/user_profiles.json", UserProfilesType)

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

	var foreignerItems = shared.FetchByProperty(
		"foreigner",
		"UserProfileId",
		profileId,
	)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the item you asked for!",
		"items":   foreignerItems,
	}, nil
}

var UserProfileForeignerInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.Foreigners
	dataBytes, _ := json.Marshal(params.Args["data"])
	ForeignerType := &structs.Foreigners{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	ForeignerData, ForeignerDataErr := shared.ReadJson(shared.GetDataRoot()+"/user_profile_foreigners.json", ForeignerType)

	if ForeignerDataErr != nil {
		fmt.Printf("Fetching User Profile's Foreigner failed because of this error - %s.\n", ForeignerDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		ForeignerData = shared.FilterByProperty(ForeignerData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(ForeignerData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_foreigners.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var UserProfileForeignerDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	ForeignerType := &structs.Foreigners{}
	ForeignerData, ForeignerDataErr := shared.ReadJson(shared.GetDataRoot()+"/user_profile_foreigners.json", ForeignerType)

	if ForeignerDataErr != nil {
		fmt.Printf("Fetching User Profile's Foreigner failed because of this error - %s.\n", ForeignerDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		ForeignerData = shared.FilterByProperty(ForeignerData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_foreigners.json"), ForeignerData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
