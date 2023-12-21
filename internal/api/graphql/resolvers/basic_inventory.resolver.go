package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) ReportValueClassInventoryResolver(params graphql.ResolveParams) (interface{}, error) {
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

		filter.ClassTypeID = &class.Id

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
			assessments, _ := r.Repo.GetMyInventoryAssessments(inventory.Id)

			if len(assessments) > 0 {
				for _, assessment := range assessments {
					if assessment.Id != 0 {
						assessmentResponse, _ := buildAssessmentResponse(r.Repo, &assessment)
						if assessmentResponse != nil {

							sumClassPurchaseGrossPrice += inventory.GrossPrice
							sumClassGrossPrice += assessmentResponse.GrossPriceDifference

							lifetimeOfAssessmentInMonths := 0

							if inventory.DepreciationTypeId != 0 {
								settings, _ := r.Repo.GetDropdownSettingById(inventory.DepreciationTypeId)

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
			Id:                 class.Id,
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

func (r *Resolver) BasicInventoryOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.BasicInventoryResponseListItem
	var filter dto.InventoryItemFilter
	var status string
	sourceTypeStr := ""
	var expireFilter bool

	if id, ok := params.Args["id"].(int); ok && id != 0 {
		filter.ID = &id
	}

	if classTypeID, ok := params.Args["class_type_id"].(int); ok && classTypeID != 0 {
		filter.ClassTypeID = &classTypeID
	}

	if officeID, ok := params.Args["office_id"].(int); ok && officeID != 0 {
		filter.OfficeID = &officeID
	}

	if typeFilter, ok := params.Args["type"].(string); ok && typeFilter != "" {
		filter.Type = &typeFilter
	}

	if search, ok := params.Args["search"].(string); ok && search != "" {
		filter.Search = &search
	}

	if sourceType, ok := params.Args["source_type"].(string); ok && sourceType != "" {
		sourceTypeStr = sourceType
	}

	if depreciationTypeID, ok := params.Args["depreciation_type_id"].(int); ok && depreciationTypeID != 0 {
		filter.DeprecationTypeID = &depreciationTypeID
	}
	if expire, ok := params.Args["expire"].(bool); ok && expire {
		expireFilter = expire
	}

	if st, ok := params.Args["status"].(string); ok && st != "" {
		status = st
	}

	if organizationUnitId, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitId != 0 {
		filter.OrganizationUnitID = &organizationUnitId
	}
	basicInventoryData, err := r.Repo.GetAllInventoryItem(filter)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	for _, item := range basicInventoryData.Data {
		resItem, err := buildInventoryResponse(r.Repo, item, *organizationUnitID)
		if len(sourceTypeStr) > 0 && sourceTypeStr != item.SourceType {
			continue
		}
		if status != "" && resItem.Status != status {
			continue
		}
		if expireFilter {
			check, _ := isCurrentOrExpiredDate(resItem.DateOfPurchase)
			if !check {
				continue
			}
		}

		if status == "Otpisan" || resItem.Active {
			items = append(items, resItem)
		}

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	}
	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   basicInventoryData.Total,
		Items:   items,
	}, nil
}

func isCurrentOrExpiredDate(dateStr string) (bool, error) {

	parsedDate, err := time.Parse("2006-01-02T00:00:00Z", dateStr)
	if err != nil {
		return false, err
	}

	currentDate := time.Now()

	return parsedDate.Year() == currentDate.Year() || parsedDate.Before(currentDate), nil
}

func (r *Resolver) BasicInventoryDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	Item, err := r.Repo.GetInventoryItem(id)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	items, err := buildInventoryItemResponse(r.Repo, Item, *organizationUnitID)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   items,
	}, nil
}

func (r *Resolver) BasicInventoryInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data []structs.BasicInventoryInsertItem
	var responseItemList []*dto.BasicInventoryResponseItem

	response := dto.Response{
		Status: "success",
	}

	loggedInProfile, _ := params.Context.Value(config.LoggedInProfileKey).(*structs.UserProfiles)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	responseItem, isSuccess, typeErr, err := r.Repo.CheckInsertInventoryData(data)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if !isSuccess && len(data) > 0 && (data[0].Type == "movable" || data[0].Type == "small") && data[0].Id == 0 {
		if typeErr == 1 {
			return apierrors.HandleAPIError(errors.New("Serijski broj artikla " + responseItem.Title + " već postoji!"))
		} else {
			return apierrors.HandleAPIError(errors.New("Inventarski broj artikla " + responseItem.Title + " već postoji!"))
		}
	}

	for _, item := range data {
		item.ArticleId = item.ContractArticleId
		item.Active = true
		item.OrganizationUnitId = *organizationUnitID
		if shared.IsInteger(item.Id) && item.Id != 0 {
			item.GrossPrice = float32(int(item.GrossPrice*100+0.5)) / 100
			itemRes, err := r.Repo.UpdateInventoryItem(item.Id, &item)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
			response.Message = "You updated this item/s!"

			items, err := buildInventoryItemResponse(r.Repo, itemRes, 0)

			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			responseItemList = append(responseItemList, items)
		} else {
			item.ArticleId = item.ContractArticleId
			item.OrganizationUnitId = *organizationUnitID
			itemRes, err := r.Repo.CreateInventoryItem(&item)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
			if item.ContractId > 0 && item.ContractArticleId > 0 {
				articles, _ := r.Repo.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
					ContractID: &item.ContractId,
					ArticleID:  &item.ContractArticleId,
				})

				if len(articles.Data) > 0 {
					article := articles.Data[0]

					article.UsedArticles++
					_, err := r.Repo.UpdateProcurementContractArticle(article.Id, article)
					if err != nil {
						return apierrors.HandleAPIError(err)
					}
				}
			}
			item.GrossPrice = float32(int(item.GrossPrice*100+0.5)) / 100
			assessment := structs.BasicInventoryAssessmentsTypesItem{
				DepreciationTypeId:   item.DepreciationTypeId,
				GrossPriceNew:        item.GrossPrice,
				GrossPriceDifference: item.GrossPrice,
				DateOfAssessment:     &itemRes.CreatedAt,
				InventoryId:          itemRes.Id,
				Active:               true,
				UserProfileId:        loggedInProfile.Id,
				Type:                 "financial",
			}

			_, err = r.Repo.CreateAssessments(&assessment)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			response.Message = "You created this item/s!"
			items, err := buildInventoryItemResponse(r.Repo, itemRes, 0)

			if err != nil {
				return apierrors.HandleAPIError(err)
			}
			responseItemList = append(responseItemList, items)
		}

	}
	response.Items = responseItemList
	return response, nil
}

func (r *Resolver) BasicInventoryDeactivateResolver(params graphql.ResolveParams) (interface{}, error) {
	response := dto.Response{
		Status: "success",
	}
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		item, err := r.Repo.GetInventoryItem(id)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		item.Active = false
		item.OfficeId = 0
		if deactivation_description, ok := params.Args["deactivation_description"].(string); ok {
			item.DeactivationDescription = deactivation_description
		}
		if inactive, ok := params.Args["inactive"].(string); ok && inactive != "" {
			item.Inactive = &inactive
		}

		if fileId, ok := params.Args["file_id"].(int); ok && fileId != 0 {
			item.DeactivationFileID = fileId
		}

		_, err = r.Repo.UpdateInventoryItem(id, item)
		if err != nil {
			return apierrors.HandleAPIError(err)
		}
		response.Message = "You Deactivate this item!"
	}

	return response, nil
}

func buildInventoryResponse(r repository.MicroserviceRepositoryInterface, item *structs.BasicInventoryInsertItem, organizationUnitID int) (*dto.BasicInventoryResponseListItem, error) {
	settingDropdownClassType := dto.DropdownSimple{}
	if item.ClassTypeId != 0 {
		settings, err := r.GetDropdownSettingById(item.ClassTypeId)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownClassType = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
		}
	}

	if item.Type == "immovable" {
		if item.OrganizationUnitId == item.TargetOrganizationUnitId || organizationUnitID == item.OrganizationUnitId {
			item.SourceType = "NS1"
		} else {
			item.SourceType = "NS2"
		}

	}

	if item.Type == "movable" {
		if item.OrganizationUnitId == item.TargetOrganizationUnitId || organizationUnitID == item.OrganizationUnitId {
			item.SourceType = "PS1"
		} else {
			item.SourceType = "PS2"
		}
	}

	settingDropdownOfficeId := dto.DropdownSimple{}
	if item.OfficeId != 0 {
		settings, err := r.GetDropdownSettingById(item.OfficeId)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownOfficeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
		}
	}

	settingDropdownDepreciationTypeId := dto.DropdownSimple{}
	assessments, _ := r.GetMyInventoryAssessments(item.Id)
	var grossPrice float32
	var dateOfAssessment string
	hasAssessments := false
	indexAssessments := 0
	if len(assessments) > 0 {
		hasAssessments = true
		for i, assessment := range assessments {
			if assessment.Id != 0 {
				assessmentResponse, _ := buildAssessmentResponse(r, &assessment)
				if assessmentResponse != nil && i == indexAssessments && assessmentResponse.Type == "financial" {

					grossPrice = assessmentResponse.GrossPriceDifference
					if len(assessments) > 1 {
						dateOfAssessment = *assessmentResponse.DateOfAssessment
					}

					settings, _ := r.GetDropdownSettingById(assessments[0].DepreciationTypeId)

					if settings != nil {
						settingDropdownDepreciationTypeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
						if settings.Value != "" {
							//Racunanje Datuma Obračun amortizacije (get DateOfPurchase and add value from Depreciation and return Calculation of depreciation)
							layout := time.RFC3339Nano
							t, _ := time.Parse(layout, item.DateOfPurchase)
							valueDepreciation, err := strconv.Atoi(settings.Value)
							if err == nil {
								year := t.Year() + valueDepreciation
								newDate := time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
								item.DateOfPurchase = newDate.Format(layout)
							}

						}

					}
					break
				} else {
					indexAssessments++
				}
			}

		}

	}

	status := "Nezadužen"

	if item.Type == "movable" && item.Active {
		itemInventoryList, _ := r.GetDispatchItemByInventoryID(item.Id)
		for _, move := range itemInventoryList {
			dispatchRes, err := r.GetDispatchItemByID(move.DispatchId)
			if err != nil {
				return nil, err
			}
			if dispatchRes.TargetOrganizationUnitId == organizationUnitID || dispatchRes.SourceOrganizationUnitId == organizationUnitID {
				switch dispatchRes.Type {
				case "revers":
					status = "Revers"
				case "allocation":
					status = "Zadužen"
				case "return":
					status = "Nezadužen"
				}
				break
			}
		}

	}
	if !item.Active {
		status = "Otpisan"
	}

	realEstateStruct := &structs.BasicInventoryRealEstatesItemResponseForInventoryItem{}

	if item.Type == "immovable" {
		realEstate, err := r.GetMyInventoryRealEstate(item.Id)
		if err != nil {
			return nil, err
		}

		realEstateStruct = &structs.BasicInventoryRealEstatesItemResponseForInventoryItem{
			Id:                       realEstate.Id,
			TypeId:                   realEstate.TypeId,
			SquareArea:               realEstate.SquareArea,
			LandSerialNumber:         realEstate.LandSerialNumber,
			EstateSerialNumber:       realEstate.EstateSerialNumber,
			OwnershipType:            realEstate.OwnershipType,
			OwnershipScope:           realEstate.OwnershipScope,
			OwnershipInvestmentScope: realEstate.OwnershipInvestmentScope,
			LimitationsDescription:   realEstate.LimitationsDescription,
			LimitationsId:            realEstate.LimitationId,
			PropertyDocument:         realEstate.PropertyDocument,
			Document:                 realEstate.Document,
			FileId:                   realEstate.FileId,
		}
	}

	organizationUnitDropdown := dto.DropdownSimple{}
	if item.OrganizationUnitId != 0 {
		organizationUnit, err := r.GetOrganizationUnitById(item.OrganizationUnitId)
		if err != nil {
			return nil, err
		}
		if organizationUnit != nil {
			organizationUnitDropdown = dto.DropdownSimple{Id: organizationUnit.Id, Title: organizationUnit.Title}
		}
	}

	targetOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.TargetOrganizationUnitId != 0 {
		targetOrganizationUnit, err := r.GetOrganizationUnitById(item.TargetOrganizationUnitId)
		if err != nil {
			return nil, err
		}
		if targetOrganizationUnit != nil {
			targetOrganizationUnitDropdown = dto.DropdownSimple{Id: targetOrganizationUnit.Id, Title: targetOrganizationUnit.Title}
		}
	}

	res := dto.BasicInventoryResponseListItem{
		Id:                     item.Id,
		Active:                 item.Active,
		Type:                   item.Type,
		Title:                  item.Title,
		Location:               item.Location,
		InventoryNumber:        item.InventoryNumber,
		GrossPrice:             grossPrice,
		PurchaseGrossPrice:     item.GrossPrice,
		DateOfPurchase:         item.DateOfPurchase,
		DateOfAssessments:      dateOfAssessment,
		Status:                 status,
		SourceType:             item.SourceType,
		RealEstate:             realEstateStruct,
		DepreciationType:       settingDropdownDepreciationTypeId,
		OrganizationUnit:       organizationUnitDropdown,
		TargetOrganizationUnit: targetOrganizationUnitDropdown,
		ClassType:              settingDropdownClassType,
		Office:                 settingDropdownOfficeId,
		HasAssessments:         hasAssessments,
	}

	return &res, nil
}

func buildInventoryItemResponse(r repository.MicroserviceRepositoryInterface, item *structs.BasicInventoryInsertItem, organizationUnitID int) (*dto.BasicInventoryResponseItem, error) {
	settingDropdownClassType := dto.DropdownSimple{}
	if item.ClassTypeId != 0 {
		settings, err := r.GetDropdownSettingById(item.ClassTypeId)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownClassType = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
		}
	}

	suppliersDropdown := dto.DropdownSimple{}
	if item.SupplierId != 0 {
		suppliers, err := r.GetSupplier(item.SupplierId)
		if err != nil {
			return nil, err
		}

		if suppliers != nil {
			suppliersDropdown = dto.DropdownSimple{Id: suppliers.Id, Title: suppliers.Title}
		}
	}
	settingDropdownOfficeId := dto.DropdownSimple{}
	if item.OfficeId != 0 {
		settings, err := r.GetDropdownSettingById(item.OfficeId)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownOfficeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
		}
	}

	targetUserDropdown := dto.DropdownSimple{}
	if item.TargetUserProfileId != 0 {
		user, err := r.GetUserProfileById(item.TargetUserProfileId)
		if err != nil {
			return nil, err
		}
		if user != nil {
			targetUserDropdown = dto.DropdownSimple{Id: user.Id, Title: user.FirstName + " " + user.LastName}
		}
	}
	var currentOrganizationUnit *structs.OrganizationUnits
	organizationUnitDropdown := dto.DropdownSimple{}
	if item.OrganizationUnitId != 0 {
		organizationUnit, err := r.GetOrganizationUnitById(item.OrganizationUnitId)
		currentOrganizationUnit = organizationUnit
		if err != nil {
			return nil, err
		}
		if organizationUnit != nil {
			organizationUnitDropdown = dto.DropdownSimple{Id: organizationUnit.Id, Title: organizationUnit.Title}
		}
	}

	targetOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.TargetOrganizationUnitId != 0 {
		targetOrganizationUnit, err := r.GetOrganizationUnitById(item.TargetOrganizationUnitId)
		currentOrganizationUnit = targetOrganizationUnit
		if err != nil {
			return nil, err
		}
		if targetOrganizationUnit != nil {
			targetOrganizationUnitDropdown = dto.DropdownSimple{Id: targetOrganizationUnit.Id, Title: targetOrganizationUnit.Title}
		}
	}

	realEstate, err := r.GetMyInventoryRealEstate(item.Id)
	if err != nil {
		return nil, err
	}

	realEstateStruct := &structs.BasicInventoryRealEstatesItemResponseForInventoryItem{}

	if realEstate != nil {
		realEstateStruct = &structs.BasicInventoryRealEstatesItemResponseForInventoryItem{
			Id:                       realEstate.Id,
			TypeId:                   realEstate.TypeId,
			SquareArea:               realEstate.SquareArea,
			LandSerialNumber:         realEstate.LandSerialNumber,
			EstateSerialNumber:       realEstate.EstateSerialNumber,
			OwnershipType:            realEstate.OwnershipType,
			OwnershipScope:           realEstate.OwnershipScope,
			OwnershipInvestmentScope: realEstate.OwnershipInvestmentScope,
			LimitationsDescription:   realEstate.LimitationsDescription,
			LimitationsId:            realEstate.LimitationId,
			PropertyDocument:         realEstate.PropertyDocument,
			Document:                 realEstate.Document,
			FileId:                   realEstate.FileId,
		}
	}
	assessments, _ := r.GetMyInventoryAssessments(item.Id)
	depreciationTypeId := 0
	var grossPrice float32
	var residualPrice *float32
	indexAssessments := 0
	lifetimeOfAssessmentInMonths := 0
	var assessmentsResponse []*dto.BasicInventoryResponseAssessment
	for i, assessment := range assessments {
		if assessment.Id != 0 {
			assessmentResponse, _ := buildAssessmentResponse(r, &assessment)
			if assessmentResponse != nil && i == indexAssessments && assessmentResponse.Type == "financial" {
				depreciationTypeId = assessmentResponse.DepreciationType.Id
				grossPrice = assessmentResponse.GrossPriceDifference
				residualPrice = assessmentResponse.ResidualPrice
				lifetimeOfAssessmentInMonths = assessmentResponse.EstimatedDuration
			} else {
				indexAssessments++
			}
			assessmentsResponse = append(assessmentsResponse, assessmentResponse)
		}
	}

	settingDropdownDepreciationTypeId := dto.DropdownSimple{}
	var amortizationValue float32
	depreciationRate := 100
	if depreciationTypeId != 0 {
		settings, _ := r.GetDropdownSettingById(depreciationTypeId)

		if settings != nil {
			settingDropdownDepreciationTypeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
			num, _ := strconv.Atoi(settings.Value)
			if num > -1 && lifetimeOfAssessmentInMonths == 0 {
				lifetimeOfAssessmentInMonths = num
			}
			if lifetimeOfAssessmentInMonths > 0 {
				depreciationRate = 100 / lifetimeOfAssessmentInMonths
				layout := time.RFC3339Nano

				t, _ := time.Parse(layout, item.CreatedAt)

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
					amortizationValue = grossPrice / float32(lifetimeOfAssessmentInMonths) / 12 * float32(totalMonths)
				}
			}
		}
	}

	itemInventoryList, _ := r.GetDispatchItemByInventoryID(item.Id)

	status := "Nezadužen"
	var movements []*dto.InventoryDispatchResponse
	indexMovements := 0
	if len(itemInventoryList) > 0 {
		for i, move := range itemInventoryList {
			dispatchRes, err := r.GetDispatchItemByID(move.DispatchId)
			if err != nil {
				return nil, err
			}
			if (dispatchRes.TargetOrganizationUnitId == organizationUnitID || dispatchRes.SourceOrganizationUnitId == organizationUnitID) && i == indexMovements {
				switch dispatchRes.Type {
				case "revers":
					status = "Revers"
				case "allocation":
					status = "Zadužen"
				case "return":
					status = "Nezadužen"
				}
			} else {
				indexAssessments++
			}
			dispatch, _ := buildInventoryDispatchResponse(r, dispatchRes)
			movements = append(movements, dispatch)
		}
	}

	if item.Type == "immovable" {
		if item.OrganizationUnitId == item.TargetOrganizationUnitId || organizationUnitID == item.OrganizationUnitId {
			item.SourceType = "NS1"
		} else {
			item.SourceType = "NS2"
		}

	}

	if item.Type == "movable" {
		if item.OrganizationUnitId == item.TargetOrganizationUnitId || organizationUnitID == item.OrganizationUnitId {
			item.SourceType = "PS1"
		} else {
			item.SourceType = "PS2"
		}
	}

	if !item.Active && item.DeactivationFileID != 0 {

		file, err := r.GetFileByID(item.DeactivationFileID)

		if err != nil {
			return nil, err
		}

		fileDropdown := dto.FileDropdownSimple{
			Id:   file.ID,
			Type: *file.Type,
			Name: file.Name,
		}

		movements = append(movements, &dto.InventoryDispatchResponse{
			DeactivationDescription: item.DeactivationDescription,
			DateOfDeactivation:      *item.Inactive,
			DeactivationFile:        fileDropdown,
		})
	}

	res := dto.BasicInventoryResponseItem{
		Id:                           item.Id,
		ArticleId:                    item.ArticleId,
		Type:                         item.Type,
		SourceType:                   item.SourceType,
		ClassType:                    settingDropdownClassType,
		DepreciationType:             settingDropdownDepreciationTypeId,
		Supplier:                     suppliersDropdown,
		RealEstate:                   realEstateStruct,
		Assessments:                  assessmentsResponse,
		Movements:                    movements,
		SerialNumber:                 item.SerialNumber,
		InventoryNumber:              item.InventoryNumber,
		Title:                        item.Title,
		Abbreviation:                 item.Abbreviation,
		InternalOwnership:            item.InternalOwnership,
		Office:                       settingDropdownOfficeId,
		Location:                     item.Location,
		TargetUserProfile:            targetUserDropdown,
		Unit:                         item.Unit,
		Amount:                       item.Amount,
		NetPrice:                     item.NetPrice,
		GrossPrice:                   grossPrice,
		ResidualPrice:                residualPrice,
		PurchaseGrossPrice:           item.GrossPrice,
		Description:                  item.Description,
		DateOfPurchase:               item.DateOfPurchase,
		Source:                       item.Source,
		DonorTitle:                   item.DonorTitle,
		InvoiceNumber:                item.InvoiceNumber,
		Active:                       item.Active,
		Inactive:                     item.Inactive,
		DeactivationDescription:      item.DeactivationDescription,
		DateOfAssessment:             item.DateOfAssessment,
		PriceOfAssessment:            item.PriceOfAssessment,
		LifetimeOfAssessmentInMonths: lifetimeOfAssessmentInMonths,
		DepreciationRate:             fmt.Sprintf("%d%%", depreciationRate),
		AmortizationValue:            amortizationValue,
		OrganizationUnit:             organizationUnitDropdown,
		TargetOrganizationUnit:       targetOrganizationUnitDropdown,
		City:                         currentOrganizationUnit.City,
		Address:                      currentOrganizationUnit.Address,
		Status:                       status,
		CreatedAt:                    item.CreatedAt,
		UpdatedAt:                    item.UpdatedAt,
	}

	return &res, nil
}
