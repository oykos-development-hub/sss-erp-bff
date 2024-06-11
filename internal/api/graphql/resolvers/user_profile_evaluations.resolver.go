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

func (r *Resolver) JudgeEvaluationReportResolver(params graphql.ResolveParams) (interface{}, error) {
	isJudge := true

	var filter dto.GetEvaluationListInputMS
	filter.IsJudge = &isJudge

	if score, ok := params.Args["score"].(string); ok && score != "" {
		filter.Score = &score
	}
	if reasonForEvaluation, ok := params.Args["reason_for_evaluation"].(string); ok && reasonForEvaluation != "" {
		filter.ReasonForEvaluation = &reasonForEvaluation
	}

	var evaluationResItemList []*dto.JudgeEvaluationReportResponseItem
	evaluationList, err := r.Repo.GetEvaluationList(&filter)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	for _, item := range evaluationList {
		evaluationResItem, err := buildJudgeEvaluationReportResponseItem(r.Repo, item)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		if organizationUnitIDinput, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitIDinput != 0 {
			if evaluationResItem.UnitID != organizationUnitIDinput {
				continue
			}
		}
		evaluationResItemList = append(evaluationResItemList, evaluationResItem)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   evaluationResItemList,
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
		item, err := r.Repo.UpdateEmployeeEvaluation(params.Context, itemID, &data)
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
		item, err := r.Repo.CreateEmployeeEvaluation(params.Context, &data)
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

	err := r.Repo.DeleteEvaluation(params.Context, itemID)
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

func buildJudgeEvaluationReportResponseItem(repo repository.MicroserviceRepositoryInterface, item *structs.Evaluation) (*dto.JudgeEvaluationReportResponseItem, error) {
	userProfile, err := repo.GetUserProfileByID(item.UserProfileID)
	if err != nil {
		return nil, err
	}

	filter := dto.JudgeResolutionsOrganizationUnitInput{
		UserProfileID: &userProfile.ID,
	}
	judge, _, err := repo.GetJudgeResolutionOrganizationUnit(&filter)

	if err != nil {
		return nil, err
	}

	orgUnit, err := repo.GetOrganizationUnitByID(judge[0].OrganizationUnitID)

	if err != nil {
		return nil, err
	}

	res := dto.JudgeEvaluationReportResponseItem{
		ID:                  item.ID,
		FullName:            userProfile.GetFullName(),
		Judgment:            orgUnit.Title,
		UnitID:              orgUnit.ID,
		DateOfEvaluation:    *item.DateOfEvaluation,
		Score:               item.Score,
		ReasonForEvaluation: item.ReasonForEvaluation,
		DecisionNumber:      *item.DecisionNumber,
		EvaluationPeriod:    item.EvaluationPeriod,
	}

	return &res, nil
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
		ID:                  item.ID,
		UserProfileID:       item.UserProfileID,
		EvaluationTypeID:    item.EvaluationTypeID,
		EvaluationType:      evaluationTypeDropdown,
		Score:               item.Score,
		DateOfEvaluation:    item.DateOfEvaluation,
		Evaluator:           item.Evaluator,
		IsRelevant:          item.IsRelevant,
		FileID:              item.FileID,
		CreatedAt:           item.CreatedAt,
		UpdatedAt:           item.UpdatedAt,
		File:                fileDropdown,
		ReasonForEvaluation: item.ReasonForEvaluation,
		DecisionNumber:      item.DecisionNumber,
		EvaluationPeriod:    item.EvaluationPeriod,
	}

	return &res, nil
}
