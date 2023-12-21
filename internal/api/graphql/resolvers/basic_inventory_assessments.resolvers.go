package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) BasicInventoryAssessmentsInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.BasicInventoryAssessmentsTypesItem
	var assessmentResponse *structs.BasicInventoryAssessmentsTypesItem
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		assessmentResponse, err = r.Repo.UpdateAssessments(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	} else {
		assessmentResponse, err = r.Repo.CreateAssessments(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
	}

	items, err := buildAssessmentResponse(r.Repo, assessmentResponse)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You inserted/updated this item!",
		Item:    items,
	}, nil
}

func (r *Resolver) BasicInventoryAssessmentDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := r.Repo.DeleteAssessment(itemId)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildAssessmentResponse(
	r repository.MicroserviceRepositoryInterface,
	item *structs.BasicInventoryAssessmentsTypesItem,
) (*dto.BasicInventoryResponseAssessment, error) {
	settings, err := r.GetDropdownSettingById(item.DepreciationTypeId)
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
		user, err := r.GetUserProfileById(item.UserProfileId)
		if err != nil {
			return nil, err
		}
		userDropdown.Id = user.Id
		userDropdown.Title = user.FirstName + " " + user.LastName
	}

	depreciationRateInt := 100 / item.EstimatedDuration
	depreciationRateString := strconv.Itoa(depreciationRateInt) + "%"

	grossPriceNew := calculateMonthlyConsumption(*item.DateOfAssessment, depreciationRateInt, item.GrossPriceDifference, item.EstimatedDuration)

	res := dto.BasicInventoryResponseAssessment{
		Id:                   item.Id,
		Type:                 item.Type,
		InventoryId:          item.InventoryId,
		DepreciationType:     settingDropdownDepreciationTypeId,
		DepreciationRate:     depreciationRateString,
		UserProfile:          userDropdown,
		ResidualPrice:        item.ResidualPrice,
		GrossPriceNew:        grossPriceNew,
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

func calculateMonthlyConsumption(startDateStr string, annualPercentage int, initialPrice float32, estimatedDuration int) float32 {
	startDate, _ := time.Parse("2006-01-02T00:00:00Z", startDateStr)
	today := time.Now()

	endDate := startDate.AddDate(estimatedDuration, 0, 0)

	if today.Before(endDate) {
		return initialPrice
	}

	monthFloat := today.Sub(startDate).Abs().Hours() / 24 / 30.44
	months := int(monthFloat)

	totalConsumption := float32(0)
	for i := 0; i < months; i++ {
		monthlyConsumption := initialPrice * float32(annualPercentage) / 100 / 12
		totalConsumption += monthlyConsumption
	}

	return totalConsumption
}
