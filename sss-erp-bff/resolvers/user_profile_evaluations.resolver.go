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

var UserProfileEvaluationResolver = func(params graphql.ResolveParams) (interface{}, error) {
	profileId := params.Args["user_profile_id"].(int)

	userProfilesData, err := getEmployeeEvaluations(profileId)
	if err != nil {
		fmt.Printf("Fetching User Profiles failed because of this error - %s.\n", err)
	}

	items := shared.ConvertToInterfaceSlice(userProfilesData)
	_ = hydrateSettings("ContractType", "ContractTypeId", items)

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   userProfilesData,
	}, nil
}

var UserProfileEvaluationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Evaluation
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return shared.ErrorResponse("Error updating settings data"), nil
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		item, err := updateEmployeeEvaluation(itemId, &data)
		if err != nil {
			fmt.Printf("Updating evaluation failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error updating organization type data"), nil
		}
		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := createEmployeeEvaluation(&data)
		if err != nil {
			fmt.Printf("Creating evaluation failed because of this error - %s.\n", err)
			return shared.ErrorResponse("Error creating organization type data"), nil
		}
		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

var UserProfileEvaluationDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteEvaluation(itemId)

	if err != nil {
		fmt.Printf("Fetching User Profile's Evaluation failed because of this error - %s.\n", err)
		return dto.Response{
			Status:  "failed",
			Message: err.Error(),
		}, nil
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func getEmployeeEvaluations(userProfileID int) ([]*structs.Evaluation, error) {
	res := &dto.GetEmployeeEvaluationListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(userProfileID)+"/evaluations", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func updateEmployeeEvaluation(id int, evaluation *structs.Evaluation) (*structs.Evaluation, error) {
	res := dto.GetEvaluationResponse{}
	_, err := shared.MakeAPIRequest("PUT", config.EVALUATIONS+"/"+strconv.Itoa(id), evaluation, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func createEmployeeEvaluation(evaluation *structs.Evaluation) (*structs.Evaluation, error) {
	res := dto.GetEvaluationResponse{}
	_, err := shared.MakeAPIRequest("POST", config.EVALUATIONS, evaluation, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func deleteEvaluation(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.EVALUATIONS+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
