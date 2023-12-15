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

var BasicInventoryAssessmentsInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.BasicInventoryAssessmentsTypesItem
	var assessmentResponse *structs.BasicInventoryAssessmentsTypesItem
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		assessmentResponse, err = updateAssessments(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	} else {
		assessmentResponse, err = createAssessments(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	items, err := buildAssessmentResponse(assessmentResponse)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You inserted/updated this item!",
		Item:    items,
	}, nil
}

var BasicInventoryAssessmentDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteAssessment(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func createAssessments(data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.ASSESSMENTS_ENDPOINT, data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateAssessments(id int, data *structs.BasicInventoryAssessmentsTypesItem) (*structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.ASSESSMENTS_ENDPOINT+"/"+strconv.Itoa(id), data, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteAssessment(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.ASSESSMENTS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func buildAssessmentResponse(item *structs.BasicInventoryAssessmentsTypesItem) (*dto.BasicInventoryResponseAssessment, error) {
	settings, err := getDropdownSettingById(item.DepreciationTypeId)
	if err != nil {
		return nil, err
	}

	settingDropdownDepreciationTypeId := dto.DropdownSimple{}
	if settings != nil {
		settingDropdownDepreciationTypeId.Id = settings.Id
		settingDropdownDepreciationTypeId.Title = settings.Title
	}

	userDropdown := dto.DropdownSimple{}
	if item.UserProfileId != 0 {
		user, err := getUserProfileById(item.UserProfileId)
		if err != nil {
			return nil, err
		}
		userDropdown.Id = user.Id
		userDropdown.Title = user.FirstName + " " + user.LastName
	}

	res := dto.BasicInventoryResponseAssessment{
		Id:                   item.Id,
		Type:                 item.Type,
		InventoryId:          item.InventoryId,
		DepreciationType:     settingDropdownDepreciationTypeId,
		UserProfile:          userDropdown,
		ResidualPrice:        item.ResidualPrice,
		GrossPriceNew:        item.GrossPriceNew,
		GrossPriceDifference: item.GrossPriceDifference,
		Active:               item.Active,
		EstimatedDuration:    item.EstimatedDuration,
		DateOfAssessment:     item.DateOfAssessment,
		CreatedAt:            item.CreatedAt,
		UpdatedAt:            item.UpdatedAt,
		FileId:               item.FileId,
	}

	return &res, nil
}
