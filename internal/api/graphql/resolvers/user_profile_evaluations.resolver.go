package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserProfileEvaluationResolver(params graphql.ResolveParams) (interface{}, error) {
	profileID := params.Args["user_profile_id"].(int)

	userEvaluationList, err := r.Repo.GetEmployeeEvaluations(profileID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	for _, evaluation := range userEvaluationList {
		evaluationType, err := r.Repo.GetDropdownSettingByID(evaluation.EvaluationTypeID)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		evaluation.EvaluationType = structs.SettingsDropdown{ID: evaluationType.ID, Title: evaluationType.Title}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   userEvaluationList,
	}, nil
}

func (r *Resolver) UserProfileEvaluationInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Evaluation
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return errors.ErrorResponse("Error updating settings data"), nil
	}

	itemID := data.ID
	if itemID != 0 {
		item, err := r.Repo.UpdateEmployeeEvaluation(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := r.Repo.CreateEmployeeEvaluation(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func (r *Resolver) UserProfileEvaluationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteEvaluation(itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
