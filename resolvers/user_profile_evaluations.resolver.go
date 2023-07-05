package resolvers

import (
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

var UserProfileEvaluationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var evaluationItems []interface{}

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

	var evaluationTypes = shared.FetchByProperty(
		"evaluation_type",
		"",
		"",
	)

	var relatedEvaluation = shared.FetchByProperty(
		"evaluation",
		"UserProfileId",
		profileId,
	)
	// # Related Employee Evaluation
	if len(relatedEvaluation) > 0 {
		for _, relatedEvaluationItem := range relatedEvaluation {
			var relatedEvaluationItemData = shared.WriteStructToInterface(relatedEvaluationItem)
			var relatedEvaluationType = shared.FindByProperty(evaluationTypes, "Id", relatedEvaluationItemData["evaluation_type_id"])

			if len(relatedEvaluationType) > 0 {
				var relatedEvaluationData = shared.WriteStructToInterface(relatedEvaluationType[0])

				relatedEvaluationItemData["evaluation_type"] = map[string]interface{}{
					"title": relatedEvaluationData["title"],
					"id":    relatedEvaluationData["id"],
				}
			}

			evaluationItems = append(evaluationItems, relatedEvaluationItemData)
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Here's the item you asked for!",
		"items":   evaluationItems,
	}, nil
}

var UserProfileEvaluationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	var data structs.Evaluation
	dataBytes, _ := json.Marshal(params.Args["data"])
	EvaluationType := &structs.Evaluation{}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	evaluationData, evaluationDataErr := shared.ReadJson(shared.GetDataRoot()+"/user_profile_evaluations.json", EvaluationType)

	if evaluationDataErr != nil {
		fmt.Printf("Fetching User Profile's evaluation failed because of this error - %s.\n", evaluationDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		evaluationData = shared.FilterByProperty(evaluationData, "Id", itemId)
	} else {
		data.Id = shared.GetRandomNumber()
	}

	var updatedData = append(evaluationData, data)

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_evaluations.json"), updatedData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You updated this item!",
		"item":    data,
	}, nil
}

var UserProfileEvaluationDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var projectRoot, _ = shared.GetProjectRoot()
	itemId := params.Args["id"]
	EvaluationType := &structs.Evaluation{}
	evaluationData, evaluationDataErr := shared.ReadJson(shared.GetDataRoot()+"/user_profile_evaluations.json", EvaluationType)

	if evaluationDataErr != nil {
		fmt.Printf("Fetching User Profile's Evaluation failed because of this error - %s.\n", evaluationDataErr)
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		evaluationData = shared.FilterByProperty(evaluationData, "Id", itemId)
	}

	_ = shared.WriteJson(shared.FormatPath(projectRoot+"/mocked-data/user_profile_evaluations.json"), evaluationData)

	return map[string]interface{}{
		"status":  "success",
		"message": "You deleted this item!",
	}, nil
}
