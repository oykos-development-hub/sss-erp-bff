package resolvers

import (
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/structs"
	"errors"
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
						assessmentResponse, _ := BuildAssessmentResponse(r.Repo, &assessment)
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

func (r *Resolver) ReportInventoryListResolver(params graphql.ResolveParams) (interface{}, error) {
	var filter dto.InventoryItemFilter
	var date string
	var sourceType string
	var organizationUnitID int

	if dateParam, ok := params.Args["date"].(string); ok && dateParam != "" {
		date = dateParam
	}

	if sourceTypeParam, ok := params.Args["source_type"].(string); ok && sourceTypeParam != "" {
		sourceType = sourceTypeParam
	}

	if organizationUnitIDParam, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitIDParam != 0 {
		filter.OrganizationUnitID = &organizationUnitIDParam
		organizationUnitID = organizationUnitIDParam
	}

	items, err := r.Repo.GetAllInventoryItem(filter)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	var response []structs.InventoryReportStruct

	for _, item := range items.Data {

		if sourceType != "" {
			itemSourceType := getItemSourceType(*item, organizationUnitID)

			if itemSourceType != sourceType {
				continue
			}
		}

		reportItem, err := buildItemForInventoryListReport(r, *item, organizationUnitID, date)

		if err != nil {
			continue
		}

		if date != "" {
			dateEnd, err := time.Parse(time.RFC3339Nano, date)
			if err != nil {
				continue
			}
			dateStart, err := time.Parse(time.RFC3339Nano, reportItem.Date)
			if err != nil {
				continue
			}
			if dateStart.Before(dateEnd) {
				response = append(response, *reportItem)
			}
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    response,
	}, nil
}

func getItemSourceType(item structs.BasicInventoryInsertItem, organizationUnitID int) string {
	if item.Type == "immovable" {
		if item.OrganizationUnitID == item.TargetOrganizationUnitID || organizationUnitID == item.OrganizationUnitID {
			item.SourceType = "NS1"
		} else {
			item.SourceType = "NS2"
		}

		if item.IsExternalDonation {
			item.SourceType = "NS2"
		}
	}

	if item.Type == "movable" {
		if item.OrganizationUnitID == item.TargetOrganizationUnitID || organizationUnitID == item.OrganizationUnitID {
			item.SourceType = "PS1"
		} else {
			item.SourceType = "PS2"
		}

		if item.IsExternalDonation {
			item.SourceType = "PS2"
		}
	}

	return item.SourceType
}

func buildItemForInventoryListReport(r *Resolver, item structs.BasicInventoryInsertItem, organizationUnitID int, date string) (*structs.InventoryReportStruct, error) {
	response := structs.InventoryReportStruct{
		ID:               item.ID,
		Title:            item.Title,
		InventoryNumber:  item.InventoryNumber,
		ProcurementPrice: item.GrossPrice,
	}

	sourceType := getItemSourceType(item, organizationUnitID)

	if item.OfficeID != 0 {
		office, err := r.Repo.GetDropdownSettingByID(item.OfficeID)
		if err != nil {
			return nil, err
		}
		response.Office = office.Title
	} else {
		response.Office = "Lager"
	}

	if sourceType == "PS1" || sourceType == "NS1" {
		response.Date = item.DateOfPurchase
	} else if !item.IsExternalDonation || sourceType == "NS2" {
		response.Date = item.DateOfPurchase
	} else {
		movements, err := r.Repo.GetDispatchItemByInventoryID(item.ID)
		if err != nil {
			return nil, err
		}
		if len(movements) == 0 {
			return nil, errors.New("source type is ns2/ps2 but there is not movements")
		}

		for _, movement := range movements {
			dispatch, err := r.Repo.GetDispatchItemByID(movement.DispatchID)

			if err != nil {
				return nil, err
			}

			if dispatch.Type == "revers" {
				response.Date = dispatch.Date
				break
			}
		}
	}

	assessments, err := r.Repo.GetMyInventoryAssessments(item.ID)

	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse(time.RFC3339Nano, date)

	if err != nil {
		return nil, err
	}

	for _, assessment := range assessments {
		assessmentItem, err := BuildAssessmentResponse(r.Repo, &assessment)

		if err != nil {
			return nil, err
		}

		date1, err := time.Parse(time.RFC3339Nano, *assessmentItem.DateOfAssessment)

		if err != nil {
			return nil, err
		}

		if date1.Before(endDate) {
			response.Price = assessmentItem.GrossPriceDifference
			response.LostValue = response.ProcurementPrice - response.Price
			break
		}
	}

	return &response, nil
}
