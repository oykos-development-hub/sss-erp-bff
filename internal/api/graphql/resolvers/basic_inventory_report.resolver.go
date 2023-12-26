package resolvers

import (
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) ReportValueClassInventoryResolver(_ graphql.ResolveParams) (interface{}, error) {
	input := dto.GetSettingsInput{
		Entity: "inventory_class_type",
	}
	var report []dto.ReportValueClassInventoryItem
	classTypes, err := r.Repo.GetDropdownSettings(&input)
	var (
		sumClassGrossPriceAllItem         float32
		sumClassPurchaseGrossPriceAllItem float32
		sumClassPriceOfAssessmentAllItem  float32
	)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	for _, class := range classTypes.Data {
		var filter dto.InventoryItemFilter

		filter.ClassTypeID = &class.ID

		basicInventoryData, err := r.Repo.GetAllInventoryItem(filter)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		var (
			sumClassGrossPrice         float32
			sumClassPurchaseGrossPrice float32
			sumClassPriceOfAssessment  float32
		)
		for _, inventory := range basicInventoryData.Data {
			assessments, _ := r.Repo.GetMyInventoryAssessments(inventory.ID)

			if len(assessments) > 0 {
				for _, assessment := range assessments {
					if assessment.ID != 0 {
						assessmentResponse, _ := buildAssessmentResponse(r.Repo, &assessment)
						if assessmentResponse != nil {

							sumClassPurchaseGrossPrice += inventory.GrossPrice
							sumClassGrossPrice += assessmentResponse.GrossPriceDifference

							lifetimeOfAssessmentInMonths := 0

							if inventory.DepreciationTypeID != 0 {
								settings, _ := r.Repo.GetDropdownSettingByID(inventory.DepreciationTypeID)

								if settings != nil {
									num, _ := strconv.Atoi(settings.Value)
									if num > -1 {
										lifetimeOfAssessmentInMonths = num
									}
									if lifetimeOfAssessmentInMonths > 0 {
										layout := time.RFC3339Nano

										t, _ := time.Parse(layout, inventory.CreatedAt)

										currentTime := time.Now()
										years := currentTime.Year() - t.Year()
										months := int(currentTime.Month() - t.Month())

										if currentTime.Day() < t.Day() {
											months--
										}

										if currentTime.YearDay() < t.YearDay() {
											years--
										}

										if months < 0 {
											years--
											months += 12
										}

										totalMonths := years*12 + months

										if totalMonths > 0 {
											sumClassPriceOfAssessment += assessmentResponse.GrossPriceDifference / float32(lifetimeOfAssessmentInMonths) / 12 * float32(totalMonths)
										}
									}
								}
							}
							break
						}
					}

				}

			}

		}
		sumClassGrossPriceAllItem += sumClassGrossPrice
		sumClassPurchaseGrossPriceAllItem += sumClassPurchaseGrossPrice
		sumClassPriceOfAssessmentAllItem += sumClassPriceOfAssessment
		report = append(report, dto.ReportValueClassInventoryItem{
			ID:                 class.ID,
			Title:              class.Title,
			Class:              class.Abbreviation,
			PurchaseGrossPrice: float32(int(sumClassPurchaseGrossPrice*100+0.5)) / 100,
			PriceOfAssessment:  float32(int(sumClassPriceOfAssessment*100+0.5)) / 100,
			GrossPrice:         float32(int(sumClassGrossPrice*100+0.5)) / 100,
		})
	}
	response := dto.ReportValueClassInventory{
		Values:             report,
		PurchaseGrossPrice: float32(int(sumClassPurchaseGrossPriceAllItem*100+0.5)) / 100,
		PriceOfAssessment:  float32(int(sumClassPriceOfAssessmentAllItem*100+0.5)) / 100,
		GrossPrice:         float32(int(sumClassGrossPriceAllItem*100+0.5)) / 100,
	}
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    response,
	}, nil
}

func (r *Resolver) ReportInventoryListResolver(_ graphql.ResolveParams) (interface{}, error) {
	input := dto.GetSettingsInput{
		Entity: "inventory_class_type",
	}
	var report []dto.ReportValueClassInventoryItem
	classTypes, err := r.Repo.GetDropdownSettings(&input)
	var (
		sumClassGrossPriceAllItem         float32
		sumClassPurchaseGrossPriceAllItem float32
		sumClassPriceOfAssessmentAllItem  float32
	)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	for _, class := range classTypes.Data {
		var filter dto.InventoryItemFilter

		filter.ClassTypeID = &class.ID

		basicInventoryData, err := r.Repo.GetAllInventoryItem(filter)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		var (
			sumClassGrossPrice         float32
			sumClassPurchaseGrossPrice float32
			sumClassPriceOfAssessment  float32
		)
		for _, inventory := range basicInventoryData.Data {
			assessments, _ := r.Repo.GetMyInventoryAssessments(inventory.ID)

			if len(assessments) > 0 {
				for _, assessment := range assessments {
					if assessment.ID != 0 {
						assessmentResponse, _ := buildAssessmentResponse(r.Repo, &assessment)
						if assessmentResponse != nil {

							sumClassPurchaseGrossPrice += inventory.GrossPrice
							sumClassGrossPrice += assessmentResponse.GrossPriceDifference

							lifetimeOfAssessmentInMonths := 0

							if inventory.DepreciationTypeID != 0 {
								settings, _ := r.Repo.GetDropdownSettingByID(inventory.DepreciationTypeID)

								if settings != nil {
									num, _ := strconv.Atoi(settings.Value)
									if num > -1 {
										lifetimeOfAssessmentInMonths = num
									}
									if lifetimeOfAssessmentInMonths > 0 {
										layout := time.RFC3339Nano

										t, _ := time.Parse(layout, inventory.CreatedAt)

										currentTime := time.Now()
										years := currentTime.Year() - t.Year()
										months := int(currentTime.Month() - t.Month())

										if currentTime.Day() < t.Day() {
											months--
										}

										if currentTime.YearDay() < t.YearDay() {
											years--
										}

										if months < 0 {
											years--
											months += 12
										}

										totalMonths := years*12 + months

										if totalMonths > 0 {
											sumClassPriceOfAssessment += assessmentResponse.GrossPriceDifference / float32(lifetimeOfAssessmentInMonths) / 12 * float32(totalMonths)
										}
									}
								}
							}
							break
						}
					}

				}

			}

		}
		sumClassGrossPriceAllItem += sumClassGrossPrice
		sumClassPurchaseGrossPriceAllItem += sumClassPurchaseGrossPrice
		sumClassPriceOfAssessmentAllItem += sumClassPriceOfAssessment
		report = append(report, dto.ReportValueClassInventoryItem{
			ID:                 class.ID,
			Title:              class.Title,
			Class:              class.Abbreviation,
			PurchaseGrossPrice: float32(int(sumClassPurchaseGrossPrice*100+0.5)) / 100,
			PriceOfAssessment:  float32(int(sumClassPriceOfAssessment*100+0.5)) / 100,
			GrossPrice:         float32(int(sumClassGrossPrice*100+0.5)) / 100,
		})
	}
	response := dto.ReportValueClassInventory{
		Values:             report,
		PurchaseGrossPrice: float32(int(sumClassPurchaseGrossPriceAllItem*100+0.5)) / 100,
		PriceOfAssessment:  float32(int(sumClassPriceOfAssessmentAllItem*100+0.5)) / 100,
		GrossPrice:         float32(int(sumClassGrossPriceAllItem*100+0.5)) / 100,
	}
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    response,
	}, nil
}
