package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
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

	userEvaluationResponseList, err := buildEvaluationResponseItemList(r.Repo, userEvaluationList)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   userEvaluationResponseList,
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
		resItem, err := buildEvaluationResponseItem(r.Repo, item)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = resItem
	} else {
		item, err := r.Repo.CreateEmployeeEvaluation(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resItem, err := buildEvaluationResponseItem(r.Repo, item)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		response.Item = resItem
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

func buildEvaluationResponseItemList(repo repository.MicroserviceRepositoryInterface, items []*structs.Evaluation) (resItemList []*dto.EvaluationResponseItem, err error) {
	for _, item := range items {
		resItem, err := buildEvaluationResponseItem(repo, item)
		if err != nil {
			return nil, err
		}
		resItemList = append(resItemList, resItem)
	}
	return
}

func buildEvaluationResponseItem(repo repository.MicroserviceRepositoryInterface, item *structs.Evaluation) (*dto.EvaluationResponseItem, error) {
	var fileDropdown dto.FileDropdownSimple

	if item.FileID != 0 {
		file, err := repo.GetFileByID(item.FileID)

		if err != nil {
			return nil, err
		}

		fileDropdown.ID = file.ID
		fileDropdown.Name = file.Name

		if file.Type != nil {
			fileDropdown.Type = *file.Type
		}
	}

	evaluationType, err := repo.GetDropdownSettingByID(item.EvaluationTypeID)
	if err != nil {
		return nil, err
	}

	evaluationTypeDropdown := dto.DropdownSimple{
		ID:    evaluationType.ID,
		Title: evaluationType.Title,
	}

	res := dto.EvaluationResponseItem{
		ID:               item.ID,
		UserProfileID:    item.UserProfileID,
		EvaluationTypeID: item.EvaluationTypeID,
		EvaluationType:   evaluationTypeDropdown,
		Score:            item.Score,
		DateOfEvaluation: item.DateOfEvaluation,
		Evaluator:        item.Evaluator,
		IsRelevant:       item.IsRelevant,
		FileID:           item.FileID,
		CreatedAt:        item.CreatedAt,
		UpdatedAt:        item.UpdatedAt,
		File:             fileDropdown,
	}

	return &res, nil
}
