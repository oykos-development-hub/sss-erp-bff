package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
)

var UserProfileSalaryParamsResolver = func(params graphql.ResolveParams) (interface{}, error) {
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

	var salaryParamsItems = shared.FetchByProperty(
		"salary_params",
		"UserProfileId",
		profileId,
	)

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the item you asked for!",
		"items":   salaryParamsItems,
	}, nil
}

var UserProfileSalaryParamsInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.SalaryParams
	dataBytes, _ := json.Marshal(params.Args["data"])
	SalaryParamsType := &structs.SalaryParams{}

	json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	SalaryParamsData, SalaryParamsDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profile_salary_params.json", SalaryParamsType)

	if SalaryParamsDataErr != nil {
		fmt.Printf("Fetching User Profile's SalaryParams failed because of this error - %s.\n", SalaryParamsDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		SalaryParamsData = shared.FilterByProperty(SalaryParamsData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(SalaryParamsData, data)

	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_salary_params.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var UserProfileSalaryParamsDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	SalaryParamsType := &structs.SalaryParams{}
	SalaryParamsData, SalaryParamsDataErr := shared.ReadJson("http://localhost:8080/mocked-data/user_profile_salary_params.json", SalaryParamsType)

	if SalaryParamsDataErr != nil {
		fmt.Printf("Fetching User Profile's SalaryParams failed because of this error - %s.\n", SalaryParamsDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		SalaryParamsData = shared.FilterByProperty(SalaryParamsData, "Id", itemId)
	}

	shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_salary_params.json"), SalaryParamsData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
