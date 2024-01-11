package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	apierrors "bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) BasicInventoryOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.BasicInventoryResponseListItem
	var filter dto.InventoryItemFilter
	var status string
	var typeOfImmovable string
	sourceTypeStr := ""
	var expireFilter bool
	var isDonation bool

	if id, ok := params.Args["id"].(int); ok && id != 0 {
		filter.ID = &id
	}

	if classTypeID, ok := params.Args["class_type_id"].(int); ok && classTypeID != 0 {
		filter.ClassTypeID = &classTypeID
	}

	if OfficeID, ok := params.Args["office_id"].(int); ok && OfficeID != 0 {
		filter.OfficeID = &OfficeID
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
	if isExternalDonation, ok := params.Args["is_external_donation"].(bool); ok && isExternalDonation {
		isDonation = isExternalDonation
	}

	if st, ok := params.Args["status"].(string); ok && st != "" {
		status = st
	}

	if typeImmovement, ok := params.Args["type_of_immovable_property"].(string); ok && typeImmovement != "" {
		typeOfImmovable = typeImmovement
	}

	if organizationUnitID, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitID != 0 && status != "Arhiva" {
		filter.OrganizationUnitID = &organizationUnitID
	}
	basicInventoryData, err := r.Repo.GetAllInventoryItem(filter)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}
	var organizationUnitID *int
	if filter.OrganizationUnitID != nil {
		organizationUnitID = filter.OrganizationUnitID
	} else {
		organizationUnitIDFromParams, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitIDFromParams == nil {
			return apierrors.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
		}
		organizationUnitID = organizationUnitIDFromParams
	}
	for _, item := range basicInventoryData.Data {
		resItem, err := buildInventoryResponse(r.Repo, item, *organizationUnitID)

		if err != nil {
			continue
		}

		if len(sourceTypeStr) > 0 && sourceTypeStr != item.SourceType {
			continue
		}
		if status != "" && status != "Arhiva" && resItem.Status != status {
			continue
		}

		date, err := time.Parse("2006-01-02T00:00:00Z", resItem.DateOfAssessments)
		if err != nil {
			continue
		}
		dateOfExpiry := date.AddDate(resItem.EstimatedDuration, 0, 0)

		newDateStr := dateOfExpiry.Format("2006-01-02T00:00:00Z")
		resItem.DateOfEndOfAssessment = newDateStr

		if expireFilter {
			check, _ := isCurrentOrExpiredDate(newDateStr)
			if !check || (item.SourceType != "PS1" && item.SourceType != "NS1") {
				continue
			}
		}

		if status == "Otpisano" || resItem.Active {
			if len(typeOfImmovable) == 0 || (typeOfImmovable != "" && resItem.RealEstate.TypeID == typeOfImmovable) {
				if !isDonation || (isDonation && item.IsExternalDonation) {
					items = append(items, resItem)
				}
			}
		}

		if err != nil {
			return apierrors.HandleAPIError(err)
		}
	}

	if page, ok := params.Args["page"].(int); ok && page != 0 {
		filter.Page = &page
	}

	if size, ok := params.Args["size"].(int); ok && size != 0 {
		filter.Size = &size
	}

	if status == "Arhiva" {
		var response []*dto.BasicInventoryResponseListItem
		var responseItem *dto.BasicInventoryResponseListItem
		basicItems, err := r.Repo.GetAllInventoryItemInOrgUnits(*organizationUnitID)

		if err != nil {
			return apierrors.HandleAPIError(err)
		}

		for _, item := range basicItems {
			var exists bool
			for _, inventoryItem := range items {
				if inventoryItem.ID == item.ItemID {
					exists = true
					responseItem = inventoryItem
				}
			}
			if exists {
				var existsInResponse bool
				for _, responseItem := range response {
					if responseItem.ID == item.ItemID {
						existsInResponse = true
					}
				}
				if !existsInResponse {
					responseItem.SourceType = "Arhiva"
					response = append(response, responseItem)
				}
			}
		}
		items = response
	}

	total := len(items)
	if filter.Page != nil && filter.Size != nil {
		start := (*filter.Page - 1) * *filter.Size
		end := start + *filter.Size

		if start >= len(items) {
			items = []*dto.BasicInventoryResponseListItem{}
		} else {

			if end > len(items) {
				end = len(items)
			}

			items = items[start:end]
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   total,
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

	responseItem, err := r.Repo.CheckInsertInventoryData(data)

	if err != nil {
		return apierrors.HandleAPIError(err)
	}

	if len(responseItem) > 0 {
		return dto.Response{
			Status:    "failed",
			Validator: responseItem,
		}, nil
	}

	for _, item := range data {
		item.ArticleID = item.ContractArticleID
		item.Active = true
		item.OrganizationUnitID = *organizationUnitID
		if item.ID != 0 {
			item.GrossPrice = float32(int(item.GrossPrice*100+0.5)) / 100
			itemRes, err := r.Repo.UpdateInventoryItem(item.ID, &item)
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
			item.ArticleID = item.ContractArticleID
			item.OrganizationUnitID = *organizationUnitID
			itemRes, err := r.Repo.CreateInventoryItem(&item)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}
			if item.ContractID > 0 && item.ContractArticleID > 0 {
				articles, _ := r.Repo.GetProcurementContractArticlesList(&dto.GetProcurementContractArticlesInput{
					ContractID: &item.ContractID,
					ArticleID:  &item.ContractArticleID,
				})

				if len(articles.Data) > 0 {
					article := articles.Data[0]

					article.UsedArticles++
					_, err := r.Repo.UpdateProcurementContractArticle(article.ID, article)
					if err != nil {
						return apierrors.HandleAPIError(err)
					}
				}
			}

			depreciationType, err := r.Repo.GetDropdownSettingByID(item.DepreciationTypeID)

			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			value, err := strconv.Atoi(depreciationType.Value)

			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			var estimatedDuration int

			if value != 0 {
				estimatedDuration = 100 / value
			} else {
				estimatedDuration = 10000
			}

			item.GrossPrice = float32(int(item.GrossPrice*100+0.5)) / 100
			assessment := structs.BasicInventoryAssessmentsTypesItem{
				EstimatedDuration:    estimatedDuration,
				DepreciationTypeID:   item.DepreciationTypeID,
				GrossPriceNew:        item.GrossPrice,
				GrossPriceDifference: item.GrossPrice,
				DateOfAssessment:     &itemRes.CreatedAt,
				InventoryID:          itemRes.ID,
				Active:               true,
				UserProfileID:        loggedInProfile.ID,
				Type:                 "financial",
			}

			_, err = r.Repo.CreateAssessments(&assessment)
			if err != nil {
				return apierrors.HandleAPIError(err)
			}

			dispatch := structs.BasicInventoryDispatchItem{
				Type:                     "created",
				SourceUserProfileID:      item.TargetUserProfileID,
				SourceOrganizationUnitID: *organizationUnitID,
				Date:                     itemRes.CreatedAt,
				InventoryID:              []int{itemRes.ID},
			}

			_, err = r.Repo.CreateDispatchItem(&dispatch)
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
		item.OfficeID = 0
		if deactivationDescription, ok := params.Args["deactivation_description"].(string); ok {
			item.DeactivationDescription = deactivationDescription
		}
		if inactive, ok := params.Args["inactive"].(string); ok && inactive != "" {
			item.Inactive = &inactive
		}

		if fileID, ok := params.Args["file_id"].(int); ok && fileID != 0 {
			item.DeactivationFileID = fileID
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
	var estimatedDuration int
	if item.ClassTypeID != 0 {
		settings, err := r.GetDropdownSettingByID(item.ClassTypeID)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownClassType = dto.DropdownSimple{ID: settings.ID, Title: settings.Title}
		}
	}

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

	settingDropdownOfficeID := dto.DropdownSimple{}
	if item.OfficeID != 0 {
		settings, err := r.GetDropdownSettingByID(item.OfficeID)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownOfficeID = dto.DropdownSimple{ID: settings.ID, Title: settings.Title}
		}
	}

	settingDropdownDepreciationTypeID := dto.DropdownSimple{}

	if item.DepreciationTypeID != 0 {
		depreciationTypeDropDown, err := r.GetDropdownSettingByID(item.DepreciationTypeID)

		if err != nil {
			return nil, err
		}
		settingDropdownDepreciationTypeID.ID = depreciationTypeDropDown.ID
		settingDropdownDepreciationTypeID.Title = depreciationTypeDropDown.Title
		value, err := strconv.Atoi(depreciationTypeDropDown.Value)

		if err != nil {
			return nil, err
		}

		if value != 0 {
			estimatedDuration = 100 / value
		} else {
			estimatedDuration = 10000
		}
	}

	assessments, _ := r.GetMyInventoryAssessments(item.ID)
	var grossPrice float32
	var dateOfAssessment string
	hasAssessments := false
	var amortizationValue float32
	indexAssessments := 0
	if len(assessments) > 0 {
		hasAssessments = true
		for i, assessment := range assessments {
			if assessment.ID != 0 {
				assessmentResponse, _ := BuildAssessmentResponse(r, &assessment)
				if assessmentResponse != nil && i == indexAssessments && assessmentResponse.Type == "financial" {
					grossPrice = assessmentResponse.GrossPriceDifference
					dateOfAssessment = *assessmentResponse.DateOfAssessment
					estimatedDuration = assessmentResponse.EstimatedDuration
					amortizationValue = calculateMonthlyConsumption(dateOfAssessment, 100/estimatedDuration, grossPrice, estimatedDuration)
					break
				}
			}
		}
	} else {
		dateOfAssessment = item.DateOfPurchase
	}

	status := "Nezaduženo"

	if item.Type == "movable" && item.Active {
		itemInventoryList, _ := r.GetDispatchItemByInventoryID(item.ID)
		if len(itemInventoryList) > 0 {
			dispatchRes, err := r.GetDispatchItemByID(itemInventoryList[0].DispatchID)
			if err != nil {
				return nil, err
			}
			if dispatchRes.Type == "revers" && !dispatchRes.IsAccepted {
				status = "Poslato"
			} else if (item.TargetOrganizationUnitID != 0 && item.TargetOrganizationUnitID != organizationUnitID) || (dispatchRes.Type == "revers" && dispatchRes.IsAccepted && item.OrganizationUnitID == organizationUnitID) {
				status = "Prihvaćeno"
			} else if dispatchRes.Type == "allocation" {
				status = "Zaduženo"
			} else if dispatchRes.Type == "return-revers" && dispatchRes.SourceOrganizationUnitID == organizationUnitID {
				status = "Povraćaj"
			} else {
				status = "Nezaduženo"
			}

		}
	}
	if !item.Active {
		status = "Otpisano"
	}

	realEstateStruct := &structs.BasicInventoryRealEstatesItemResponseForInventoryItem{}

	if item.Type == "immovable" {
		realEstate, err := r.GetMyInventoryRealEstate(item.ID)
		if err != nil {
			return nil, err
		}

		realEstateStruct = &structs.BasicInventoryRealEstatesItemResponseForInventoryItem{
			ID:                       realEstate.ID,
			TypeID:                   realEstate.TypeID,
			SquareArea:               realEstate.SquareArea,
			LandSerialNumber:         realEstate.LandSerialNumber,
			EstateSerialNumber:       realEstate.EstateSerialNumber,
			OwnershipType:            realEstate.OwnershipType,
			OwnershipScope:           realEstate.OwnershipScope,
			OwnershipInvestmentScope: realEstate.OwnershipInvestmentScope,
			LimitationsDescription:   realEstate.LimitationsDescription,
			LimitationsID:            realEstate.LimitationID,
			PropertyDocument:         realEstate.PropertyDocument,
			Document:                 realEstate.Document,
			FileID:                   realEstate.FileID,
		}
	}

	organizationUnitDropdown := dto.DropdownOUSimple{}
	if item.OrganizationUnitID != 0 {
		organizationUnit, err := r.GetOrganizationUnitByID(item.OrganizationUnitID)
		if err != nil {
			return nil, err
		}
		if organizationUnit != nil {
			organizationUnitDropdown = dto.DropdownOUSimple{
				ID:      organizationUnit.ID,
				Title:   organizationUnit.Title,
				City:    organizationUnit.City,
				Address: organizationUnit.Address,
			}
		}
	}

	targetOrganizationUnitDropdown := dto.DropdownOUSimple{}
	if item.TargetOrganizationUnitID != 0 {
		targetOrganizationUnit, err := r.GetOrganizationUnitByID(item.TargetOrganizationUnitID)
		if err != nil {
			return nil, err
		}
		if targetOrganizationUnit != nil {
			targetOrganizationUnitDropdown = dto.DropdownOUSimple{
				ID:      targetOrganizationUnit.ID,
				Title:   targetOrganizationUnit.Title,
				City:    targetOrganizationUnit.City,
				Address: targetOrganizationUnit.Address}
		}
	}

	//get invoice

	res := dto.BasicInventoryResponseListItem{
		ID:                           item.ID,
		Active:                       item.Active,
		Type:                         item.Type,
		Title:                        item.Title,
		Location:                     item.Location,
		InventoryNumber:              item.InventoryNumber,
		EstimatedDuration:            estimatedDuration,
		GrossPrice:                   grossPrice,
		Inactive:                     item.Inactive,
		PurchaseGrossPrice:           item.GrossPrice,
		DateOfPurchase:               item.DateOfPurchase,
		DateOfAssessments:            dateOfAssessment,
		LifetimeOfAssessmentInMonths: estimatedDuration,
		AmortizationValue:            amortizationValue,
		City:                         &organizationUnitDropdown.City,
		Address:                      &organizationUnitDropdown.Address,
		Status:                       status,
		SourceType:                   item.SourceType,
		RealEstate:                   realEstateStruct,
		DepreciationType:             settingDropdownDepreciationTypeID,
		OrganizationUnit:             organizationUnitDropdown,
		TargetOrganizationUnit:       targetOrganizationUnitDropdown,
		ClassType:                    settingDropdownClassType,
		Office:                       settingDropdownOfficeID,
		Invoice:                      dto.DropdownSimple{}, // add invoice dropdown
		HasAssessments:               hasAssessments,
		IsExternalDonation:           item.IsExternalDonation,
		Source:                       item.Source,
	}

	return &res, nil
}

func buildInventoryItemResponse(r repository.MicroserviceRepositoryInterface, item *structs.BasicInventoryInsertItem, organizationUnitID int) (*dto.BasicInventoryResponseItem, error) {
	settingDropdownClassType := dto.DropdownSimple{}
	if item.ClassTypeID != 0 {
		settings, err := r.GetDropdownSettingByID(item.ClassTypeID)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownClassType = dto.DropdownSimple{ID: settings.ID, Title: settings.Title}
		}
	}

	suppliersDropdown := dto.DropdownSimple{}
	if item.SupplierID != 0 {
		suppliers, err := r.GetSupplier(item.SupplierID)
		if err != nil {
			return nil, err
		}

		if suppliers != nil {
			suppliersDropdown = dto.DropdownSimple{ID: suppliers.ID, Title: suppliers.Title}
		}
	}

	donorDropdown := dto.DropdownSimple{}
	if item.DonorID != 0 {
		donor, err := r.GetSupplier(item.DonorID)
		if err != nil {
			return nil, err
		}

		if donor != nil {
			donorDropdown = dto.DropdownSimple{ID: donor.ID, Title: donor.Title}
		}
	}

	settingDropdownOfficeID := dto.DropdownSimple{}
	if item.OfficeID != 0 {
		settings, err := r.GetDropdownSettingByID(item.OfficeID)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownOfficeID = dto.DropdownSimple{ID: settings.ID, Title: settings.Title}
		}
	}

	targetUserDropdown := dto.DropdownSimple{}
	if item.TargetUserProfileID != 0 {
		user, err := r.GetUserProfileByID(item.TargetUserProfileID)
		if err != nil {
			return nil, err
		}
		if user != nil {
			targetUserDropdown = dto.DropdownSimple{ID: user.ID, Title: user.FirstName + " " + user.LastName}
		}
	}
	var currentOrganizationUnit *structs.OrganizationUnits
	organizationUnitDropdown := dto.DropdownSimple{}
	if item.OrganizationUnitID != 0 {
		organizationUnit, err := r.GetOrganizationUnitByID(item.OrganizationUnitID)
		currentOrganizationUnit = organizationUnit
		if err != nil {
			return nil, err
		}
		if organizationUnit != nil {
			organizationUnitDropdown = dto.DropdownSimple{ID: organizationUnit.ID, Title: organizationUnit.Title}
		}
	}

	targetOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.TargetOrganizationUnitID != 0 {
		targetOrganizationUnit, err := r.GetOrganizationUnitByID(item.TargetOrganizationUnitID)
		currentOrganizationUnit = targetOrganizationUnit
		if err != nil {
			return nil, err
		}
		if targetOrganizationUnit != nil {
			targetOrganizationUnitDropdown = dto.DropdownSimple{ID: targetOrganizationUnit.ID, Title: targetOrganizationUnit.Title}
		}
	}

	realEstate, err := r.GetMyInventoryRealEstate(item.ID)
	if err != nil {
		return nil, err
	}

	realEstateStruct := &structs.BasicInventoryRealEstatesItemResponseForInventoryItem{}

	if realEstate != nil {
		realEstateStruct = &structs.BasicInventoryRealEstatesItemResponseForInventoryItem{
			ID:                       realEstate.ID,
			TypeID:                   realEstate.TypeID,
			SquareArea:               realEstate.SquareArea,
			LandSerialNumber:         realEstate.LandSerialNumber,
			EstateSerialNumber:       realEstate.EstateSerialNumber,
			OwnershipType:            realEstate.OwnershipType,
			OwnershipScope:           realEstate.OwnershipScope,
			OwnershipInvestmentScope: realEstate.OwnershipInvestmentScope,
			LimitationsDescription:   realEstate.LimitationsDescription,
			LimitationsID:            realEstate.LimitationID,
			PropertyDocument:         realEstate.PropertyDocument,
			Document:                 realEstate.Document,
			FileID:                   realEstate.FileID,
		}
	}
	assessments, _ := r.GetMyInventoryAssessments(item.ID)
	depreciationTypeID := 0
	var grossPrice float32
	var residualPrice *float32
	var dateOfAssessment string
	var amortizationValue float32
	indexAssessments := 0
	lifetimeOfAssessmentInMonths := 0
	var assessmentsResponse []*dto.BasicInventoryResponseAssessment
	for i, assessment := range assessments {
		if assessment.ID != 0 {
			assessmentResponse, _ := BuildAssessmentResponse(r, &assessment)
			if assessmentResponse != nil && i == 0 && assessmentResponse.Type == "financial" {
				depreciationTypeID = assessmentResponse.DepreciationType.ID
				grossPrice = assessmentResponse.GrossPriceDifference
				residualPrice = assessmentResponse.ResidualPrice
				lifetimeOfAssessmentInMonths = assessmentResponse.EstimatedDuration
				dateOfAssessment = *assessmentResponse.DateOfAssessment
				amortizationValue = calculateMonthlyConsumption(dateOfAssessment, 100/lifetimeOfAssessmentInMonths, grossPrice, lifetimeOfAssessmentInMonths)

			}
			assessmentsResponse = append(assessmentsResponse, assessmentResponse)
		}
	}

	settingDropdownDepreciationTypeID := dto.DropdownSimple{}
	depreciationRate := 100
	if depreciationTypeID != 0 {
		settings, _ := r.GetDropdownSettingByID(depreciationTypeID)

		if settings != nil {
			settingDropdownDepreciationTypeID = dto.DropdownSimple{ID: settings.ID, Title: settings.Title}
			num, _ := strconv.Atoi(settings.Value)
			if num > -1 && lifetimeOfAssessmentInMonths == 0 {
				lifetimeOfAssessmentInMonths = num
			}
		}
	}

	itemInventoryList, _ := r.GetDispatchItemByInventoryID(item.ID)

	status := "Nezaduženo"
	var movements []*dto.InventoryDispatchResponse
	if len(itemInventoryList) > 0 {
		for i, move := range itemInventoryList {
			dispatchRes, err := r.GetDispatchItemByID(move.DispatchID)
			if err != nil {
				return nil, err
			}
			if i == 0 {
				if dispatchRes.Type == "revers" && !dispatchRes.IsAccepted {
					status = "Poslato"
				} else if (item.TargetOrganizationUnitID != 0 && item.TargetOrganizationUnitID != organizationUnitID) || (dispatchRes.Type == "revers" && dispatchRes.IsAccepted && item.OrganizationUnitID == organizationUnitID) {
					status = "Prihvaćeno"
				} else if dispatchRes.Type == "allocation" {
					status = "Zaduženo"
				} else {
					status = "Nezaduženo"
				}
			} else {
				indexAssessments++
			}
			dispatch, _ := buildInventoryDispatchResponse(r, dispatchRes, organizationUnitID)
			movements = append(movements, dispatch)
		}
	}

	if !item.Active {
		status = "Otpisano"
	}

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

	if item.SourceType == "PS2" && !item.IsExternalDonation {
		var movementResponse []*dto.InventoryDispatchResponse
		var addMovement bool
		for i := len(movements) - 1; i >= 0; i-- {
			movement := movements[i]

			if movement.Type == "revers" && movement.TargetOrganizationUnit.ID == organizationUnitID {
				addMovement = true
			}

			if addMovement && movement.Type != "created" {
				movementResponse = append([]*dto.InventoryDispatchResponse{movement}, movementResponse...)
			}

			if movement.Type == "return-revers" {
				addMovement = false
			}
		}

		movements = movementResponse
		if len(movements) > 0 && movements[0].Type == "return-revers" && item.OrganizationUnitID != organizationUnitID {
			status = "Arhiva"
		}
	}

	if !item.Active && item.DeactivationFileID != 0 {

		file, err := r.GetFileByID(item.DeactivationFileID)

		if err != nil {
			return nil, err
		}

		fileDropdown := dto.FileDropdownSimple{
			ID:   file.ID,
			Type: *file.Type,
			Name: file.Name,
		}

		movement := &dto.InventoryDispatchResponse{
			DeactivationDescription: item.DeactivationDescription,
			DateOfDeactivation:      *item.Inactive,
			File:                    fileDropdown,
		}

		movements = append([]*dto.InventoryDispatchResponse{movement}, movements...)
	}

	var donationFiles []dto.FileDropdownSimple

	for _, fileID := range item.DonationFiles {
		file, err := r.GetFileByID(fileID)

		if err != nil {
			return nil, err
		}

		donationFiles = append(donationFiles, dto.FileDropdownSimple{
			ID:   file.ID,
			Name: file.Name,
			Type: *file.Type,
		})
	}

	/*
		get invoice
	*/

	res := dto.BasicInventoryResponseItem{
		ID:                           item.ID,
		ArticleID:                    item.ArticleID,
		Type:                         item.Type,
		SourceType:                   item.SourceType,
		ClassType:                    settingDropdownClassType,
		DepreciationType:             settingDropdownDepreciationTypeID,
		Invoice:                      dto.DropdownSimple{}, //add invoice dropdown
		Supplier:                     suppliersDropdown,
		Donor:                        donorDropdown,
		RealEstate:                   realEstateStruct,
		Assessments:                  assessmentsResponse,
		Movements:                    movements,
		SerialNumber:                 item.SerialNumber,
		InventoryNumber:              item.InventoryNumber,
		Title:                        item.Title,
		Abbreviation:                 item.Abbreviation,
		InternalOwnership:            item.InternalOwnership,
		Office:                       settingDropdownOfficeID,
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
		DateOfAssessment:             &dateOfAssessment,
		PriceOfAssessment:            item.PriceOfAssessment,
		LifetimeOfAssessmentInMonths: lifetimeOfAssessmentInMonths,
		DepreciationRate:             fmt.Sprintf("%d%%", depreciationRate),
		AmortizationValue:            amortizationValue,
		OrganizationUnit:             organizationUnitDropdown,
		TargetOrganizationUnit:       targetOrganizationUnitDropdown,
		City:                         currentOrganizationUnit.City,
		Address:                      currentOrganizationUnit.Address,
		Status:                       status,
		DonationDescription:          item.DonationDescription,
		DonationFiles:                donationFiles,
		CreatedAt:                    item.CreatedAt,
		UpdatedAt:                    item.UpdatedAt,
		IsExternalDonation:           item.IsExternalDonation,
		Owner:                        item.Owner,
	}

	return &res, nil
}

func calculateAmortizationPrice(r repository.MicroserviceRepositoryInterface, depreciationTypeID *int, CreatedAt *string, grossPrice *float32) float32 {

	if depreciationTypeID != nil && *depreciationTypeID != 0 {
		settings, _ := r.GetDropdownSettingByID(*depreciationTypeID)
		lifetimeOfAssessmentInMonths := 0
		if settings != nil {
			num, _ := strconv.Atoi(settings.Value)
			if num > -1 {
				lifetimeOfAssessmentInMonths = num
			}
			if lifetimeOfAssessmentInMonths > 0 {
				layout := time.RFC3339Nano

				t, _ := time.Parse(layout, *CreatedAt)

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
					return *grossPrice / float32(lifetimeOfAssessmentInMonths) / 12 * float32(totalMonths)
				}
			}
		}
	}
	return 0
}
