package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
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

	itemID := data.ID
	if itemID != 0 {
		assessmentResponse, err = r.Repo.UpdateAssessments(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	} else {
		assessmentResponse, err = r.Repo.CreateAssessments(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	items, err := BuildAssessmentResponse(r.Repo, assessmentResponse)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You inserted/updated this item!",
		Item:    items,
	}, nil
}

func (r *Resolver) BasicEXCLInventoryAssessmentsInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var dataArr []structs.BasicInventoryAssessmentsTypesItem
	var assessmentResponse *structs.BasicInventoryAssessmentsTypesItem
	var items []dto.BasicInventoryResponseAssessment
	var err error
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &dataArr)
	if len(dataArr) > 0 {
		for _, data := range dataArr {
			assessmentResponse, err = r.Repo.CreateAssessments(params.Context, &data)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			item, err := BuildAssessmentResponse(r.Repo, assessmentResponse)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			items = append(items, *item)
		}
	}

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You inserted/updated this item!",
		Items:   items,
	}, nil
}

func (r *Resolver) BasicInventoryAssessmentDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteAssessment(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func BuildAssessmentResponse(
	r repository.MicroserviceRepositoryInterface,
	item *structs.BasicInventoryAssessmentsTypesItem,
) (*dto.BasicInventoryResponseAssessment, error) {
	settings, err := r.GetDropdownSettingByID(item.DepreciationTypeID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get dropdown setting by id")
	}

	settingDropdownDepreciationTypeID := dto.DropdownSimple{}
	if settings != nil {
		settingDropdownDepreciationTypeID.ID = settings.ID
		settingDropdownDepreciationTypeID.Title = settings.Title
	}

	userDropdown := dto.DropdownSimple{}
	if item.UserProfileID != 0 {
		user, err := r.GetUserProfileByID(item.UserProfileID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}
		userDropdown.ID = user.ID
		userDropdown.Title = user.FirstName + " " + user.LastName
	}

	if item.EstimatedDuration == 0 {
		item.EstimatedDuration = 10000
	}

	depreciationRateInt := 100 / item.EstimatedDuration
	depreciationRateString := strconv.Itoa(depreciationRateInt) + "%"

	grossPriceNew := calculateMonthlyConsumption(*item.DateOfAssessment, item.GrossPriceDifference, item.EstimatedDuration)

	res := dto.BasicInventoryResponseAssessment{
		ID:                   item.ID,
		Type:                 item.Type,
		InventoryID:          item.InventoryID,
		DepreciationType:     settingDropdownDepreciationTypeID,
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
		FileID:               item.FileID,
	}

	return &res, nil
}

func calculateMonthlyConsumption(startDateStr string, initialPrice float32, estimatedDuration int) float32 {
	startDate, _ := time.Parse(config.ISO8601Format, startDateStr)
	//today := time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC)
	today := time.Now()
	endDate := startDate.AddDate(estimatedDuration, 0, 0)

	if endDate.Before(today) {
		return initialPrice
	}

	years := today.Year() - startDate.Year()
	months := int(today.Month()) - int(startDate.Month())
	if months < 0 {
		years--
		months += 12
	}

	months = years*12 + months

	totalConsumption := float32(0)
	var percentage float32
	if estimatedDuration == 0 {
		percentage = 0.0000000001
	} else {
		percentage = float32(100) / float32(estimatedDuration)
	}

	monthlyConsumption := initialPrice * percentage / 100 / 12
	for i := 0; i < months; i++ {
		totalConsumption += monthlyConsumption
	}

	return totalConsumption
}
