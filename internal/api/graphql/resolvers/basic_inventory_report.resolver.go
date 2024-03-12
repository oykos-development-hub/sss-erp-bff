package resolvers

import (
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/structs"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) ReportValueClassInventoryResolver(params graphql.ResolveParams) (interface{}, error) {
	var classTypeID int
	if classTypeIDParam, ok := params.Args["class_type_id"].(int); ok && classTypeIDParam != 0 {
		classTypeID = classTypeIDParam
	}
	var classTypes []structs.SettingsDropdown

	if classTypeID == 0 {
		input := dto.GetSettingsInput{
			Entity: "inventory_class_type",
		}
		classTypesData, err := r.Repo.GetDropdownSettings(&input)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		classTypes = classTypesData.Data
	} else {
		classType, err := r.Repo.GetDropdownSettingByID(classTypeID)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		classTypes = append(classTypes, *classType)
	}

	var report []dto.ReportValueClassInventoryItem

	var (
		sumClassGrossPriceAllItem         float32
		sumClassPurchaseGrossPriceAllItem float32
		sumClassPriceOfAssessmentAllItem  float32
	)

	var filter dto.InventoryItemFilter

	if organizationUnitIDParam, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitIDParam != 0 {
		filter.SourceOrganizationUnitID = &organizationUnitIDParam
	}

	for _, class := range classTypes {

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
				assessment, err := BuildAssessmentResponse(r.Repo, &assessments[0])

				if err != nil {
					continue
				}
				sumClassPurchaseGrossPrice += inventory.GrossPrice
				//amortization := calculateAmortizationPrice(r.Repo, &assessment.DepreciationType.ID, assessment.DateOfAssessment, &assessment.GrossPriceDifference) // nabavna vrijednost
				amortization := inventory.AssessmentPrice
				sumClassGrossPrice += (assessment.GrossPriceDifference - amortization) //ispravak vrijednosti
				//

			}

		}
		sumClassPriceOfAssessment = sumClassPurchaseGrossPrice - sumClassGrossPrice //trenutna vrijednost
		sumClassGrossPriceAllItem += sumClassGrossPrice
		sumClassPurchaseGrossPriceAllItem += sumClassPurchaseGrossPrice
		sumClassPriceOfAssessmentAllItem += sumClassPriceOfAssessment
		report = append(report, dto.ReportValueClassInventoryItem{
			ID:                 class.ID,
			Title:              class.Title,
			Class:              class.Abbreviation,
			PurchaseGrossPrice: float32(int(sumClassPurchaseGrossPrice*100+0.5)) / 100,
			LostValue:          float32(int(sumClassPriceOfAssessment*100+0.5)) / 100,
			Price:              float32(int(sumClassGrossPrice*100+0.5)) / 100,
		})
	}
	response := dto.ReportValueClassInventory{
		Values:             report,
		PurchaseGrossPrice: float32(int(sumClassPurchaseGrossPriceAllItem*100+0.5)) / 100,
		LostValue:          float32(int(sumClassPriceOfAssessmentAllItem*100+0.5)) / 100,
		Price:              float32(int(sumClassGrossPriceAllItem*100+0.5)) / 100,
	}
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    response,
	}, nil
}

func (r *Resolver) ReportInventoryListResolver(params graphql.ResolveParams) (interface{}, error) {
	var filter dto.ItemReportFilterDTO

	if dateParam, ok := params.Args["date"].(string); ok && dateParam != "" {
		filter.Date = &dateParam
	}

	if sourceTypeParam, ok := params.Args["source_type"].(string); ok && sourceTypeParam != "" {
		filter.SourceType = &sourceTypeParam
	}

	if organizationUnitIDParam, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitIDParam != 0 {
		filter.OrganizationUnitID = &organizationUnitIDParam
	}

	if typeParam, ok := params.Args["type"].(string); ok && typeParam != "" {
		filter.Type = &typeParam
	}

	if officeParam, ok := params.Args["office_id"].(int); ok && officeParam != 0 {
		filter.OfficeID = &officeParam
	}

	items, err := r.Repo.GetAllInventoryItemForReport(filter)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	for i := 0; i < len(items); i++ {
		items[i].Type = items[i].SourceType
		if items[i].SourceType == "NS1" || items[i].SourceType == "NS2" {
			realEstate, err := r.Repo.GetMyInventoryRealEstate(items[i].ID)

			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			items[i].Title = realEstate.TypeID
		}

		if items[i].OfficeID != 0 {
			office, err := r.Repo.GetDropdownSettingByID(items[i].OfficeID)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
			items[i].Office = office.Title
		} else {
			items[i].Office = "Lager"
		}
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    items,
	}, nil
}
