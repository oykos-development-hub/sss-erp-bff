package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

var BasicInventoryOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.BasicInventoryResponseListItem
	var filter dto.InventoryItemFilter
	sourceTypeStr := ""

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return shared.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

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

	if page, ok := params.Args["page"].(int); ok && page != 0 {
		filter.Page = &page
	}

	if size, ok := params.Args["size"].(int); ok && size != 0 {
		filter.Size = &size
	}

	filter.OrganizationUnitID = organizationUnitID
	basicInventoryData, err := getAllInventoryItem(filter)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, item := range basicInventoryData.Data {
		resItem, err := buildInventoryResponse(item, *organizationUnitID)
		if len(sourceTypeStr) > 0 && sourceTypeStr != item.SourceType {
			continue
		}
		items = append(items, resItem)

		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   basicInventoryData.Total,
		Items:   items,
	}, nil
}

var BasicInventoryDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return shared.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	Item, err := getInventoryItem(id)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	items, err := buildInventoryItemResponse(Item, *organizationUnitID)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   items,
	}, nil
}

var BasicInventoryInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data []structs.BasicInventoryInsertItem
	var responseItemList []*dto.BasicInventoryResponseItem

	response := dto.Response{
		Status: "success",
	}

	loggedInProfile, _ := params.Context.Value(config.LoggedInProfileKey).(*structs.UserProfiles)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return shared.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, item := range data {
		if shared.IsInteger(item.Id) && item.Id != 0 {
			itemRes, err := updateInventoryItem(item.Id, &item)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			response.Message = "You updated this item/s!"

			items, err := buildInventoryItemResponse(itemRes, 0)

			if err != nil {
				return shared.HandleAPIError(err)
			}

			responseItemList = append(responseItemList, items)
		} else {
			item.OrganizationUnitId = *organizationUnitID
			itemRes, err := createInventoryItem(&item)
			if err != nil {
				return shared.HandleAPIError(err)
			}
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

			_, err = createAssessments(&assessment)
			if err != nil {
				return shared.HandleAPIError(err)
			}

			response.Message = "You created this item/s!"
			items, err := buildInventoryItemResponse(itemRes, 0)

			if err != nil {
				return shared.HandleAPIError(err)
			}
			responseItemList = append(responseItemList, items)
		}
	}
	response.Items = responseItemList
	return response, nil
}

var BasicInventoryDeactivateResolver = func(params graphql.ResolveParams) (interface{}, error) {
	response := dto.Response{
		Status: "success",
	}
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		item, err := getInventoryItem(id)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item.Active = false
		item.OfficeId = 0
		if deactivation_description, ok := params.Args["deactivation_description"].(string); ok {
			item.DeactivationDescription = deactivation_description
		}

		_, err = updateInventoryItem(id, item)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Message = "You Deactivate this item!"
	}

	return response, nil
}

func createInventoryItem(item *structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, error) {
	res := dto.GetBasicInventoryInsertItem{}

	_, err := shared.MakeAPIRequest("POST", config.INVENTORY_ITEM_ENDOPOINT, item, &res)
	if err != nil {
		return nil, err
	}

	if item.RealEstate != nil {
		item.RealEstate.ItemId = res.Data.Id
		_, err := shared.MakeAPIRequest("POST", config.REAL_ESTATES_ENDPOINT, item.RealEstate, nil)
		if err != nil {
			return nil, err
		}
	}

	return res.Data, nil
}

func updateInventoryItem(id int, item *structs.BasicInventoryInsertItem) (*structs.BasicInventoryInsertItem, error) {
	res := dto.GetBasicInventoryInsertItem{}
	res1 := dto.GetBasicInventoryInsertItem{}

	_, err := shared.MakeAPIRequest("PUT", config.INVENTORY_ITEM_ENDOPOINT+"/"+strconv.Itoa(id), item, &res)
	if err != nil {
		return nil, err
	}

	if item.RealEstate != nil {
		item.RealEstate.ItemId = res.Data.Id
		if item.RealEstate.Id != 0 {
			_, err := shared.MakeAPIRequest("PUT", config.REAL_ESTATES_ENDPOINT+"/"+strconv.Itoa(item.RealEstate.Id), item.RealEstate, &res1)
			if err != nil {
				return nil, err
			}
		} else {
			_, err := shared.MakeAPIRequest("POST", config.REAL_ESTATES_ENDPOINT, item.RealEstate, &res1)
			if err != nil {
				return nil, err
			}
		}
	}

	return res.Data, nil
}

func getInventoryItem(id int) (*structs.BasicInventoryInsertItem, error) {
	res := dto.GetBasicInventoryInsertItem{}
	_, err := shared.MakeAPIRequest("GET", config.INVENTORY_ITEM_ENDOPOINT+"/"+strconv.Itoa(id), nil, &res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func getAllInventoryItem(filter dto.InventoryItemFilter) (*dto.GetAllBasicInventoryItem, error) {
	res := &dto.GetAllBasicInventoryItem{}
	_, err := shared.MakeAPIRequest("GET", config.INVENTORY_ITEM_ENDOPOINT, filter, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func buildInventoryResponse(item *structs.BasicInventoryInsertItem, organizationUnitID int) (*dto.BasicInventoryResponseListItem, error) {

	settingDropdownClassType := dto.DropdownSimple{}
	if item.ClassTypeId != 0 {
		settings, err := getDropdownSettingById(item.ClassTypeId)
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
		settings, err := getDropdownSettingById(item.OfficeId)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownOfficeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
		}
	}

	settingDropdownDepreciationTypeId := dto.DropdownSimple{}
	assessments, _ := getMyInventoryAssessments(item.Id)

	if len(assessments) > 0 {
		item.GrossPrice = assessments[0].GrossPriceDifference
		settings, _ := getDropdownSettingById(assessments[0].DepreciationTypeId)

		if settings != nil {
			settingDropdownDepreciationTypeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
		}
	}

	status := "Lager"

	if item.Type == "movable" {
		itemInventoryList, _ := getDispatchItemByInventoryID(item.Id)
		if len(itemInventoryList) > 0 {
			dispatchRes, err := getDispatchItemByID(itemInventoryList[0].DispatchId)
			if err != nil {
				return nil, err
			}
			if status == "" && dispatchRes.TargetOrganizationUnitId == organizationUnitID || dispatchRes.SourceOrganizationUnitId == organizationUnitID {
				switch dispatchRes.Type {
				case "revers":
					status = "Revers"
				case "allocation":
					status = "Zadužen"
				case "return":
					status = "Lager"
				}
			}
		}
	}

	realEstateStruct := &structs.BasicInventoryRealEstatesItemResponseForInventoryItem{}

	if item.Type == "immovable" {
		realEstate, err := getMyInventoryRealEstate(item.Id)
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
		organizationUnit, err := getOrganizationUnitById(item.OrganizationUnitId)
		if err != nil {
			return nil, err
		}
		if organizationUnit != nil {
			organizationUnitDropdown = dto.DropdownSimple{Id: organizationUnit.Id, Title: organizationUnit.Title}
		}
	}

	targetOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.TargetOrganizationUnitId != 0 {
		targetOrganizationUnit, err := getOrganizationUnitById(item.TargetOrganizationUnitId)
		if err != nil {
			return nil, err
		}
		if targetOrganizationUnit != nil {
			targetOrganizationUnitDropdown = dto.DropdownSimple{Id: targetOrganizationUnit.Id, Title: targetOrganizationUnit.Title}
		}
	}

	res := dto.BasicInventoryResponseListItem{
		Id:                     item.Id,
		Type:                   item.Type,
		Title:                  item.Title,
		Location:               item.Location,
		InventoryNumber:        item.InventoryNumber,
		GrossPrice:             item.GrossPrice,
		DateOfPurchase:         item.DateOfPurchase,
		Status:                 status,
		SourceType:             item.SourceType,
		RealEstate:             realEstateStruct,
		DepreciationType:       settingDropdownDepreciationTypeId,
		OrganizationUnit:       organizationUnitDropdown,
		TargetOrganizationUnit: targetOrganizationUnitDropdown,
		ClassType:              settingDropdownClassType,
		Office:                 settingDropdownOfficeId,
	}

	return &res, nil
}

func buildInventoryItemResponse(item *structs.BasicInventoryInsertItem, organizationUnitID int) (*dto.BasicInventoryResponseItem, error) {
	settingDropdownClassType := dto.DropdownSimple{}
	if item.ClassTypeId != 0 {
		settings, err := getDropdownSettingById(item.ClassTypeId)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownClassType = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
		}
	}

	suppliersDropdown := dto.DropdownSimple{}
	if item.SupplierId != 0 {
		suppliers, err := getSupplier(item.SupplierId)
		if err != nil {
			return nil, err
		}

		if suppliers != nil {
			suppliersDropdown = dto.DropdownSimple{Id: suppliers.Id, Title: suppliers.Title}
		}
	}
	settingDropdownOfficeId := dto.DropdownSimple{}
	if item.OfficeId != 0 {
		settings, err := getDropdownSettingById(item.OfficeId)
		if err != nil {
			return nil, err
		}

		if settings != nil {
			settingDropdownOfficeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
		}
	}

	targetUserDropdown := dto.DropdownSimple{}
	if item.TargetUserProfileId != 0 {
		user, err := getUserProfileById(item.TargetUserProfileId)
		if err != nil {
			return nil, err
		}
		if user != nil {
			targetUserDropdown = dto.DropdownSimple{Id: user.Id, Title: user.FirstName + " " + user.LastName}
		}
	}

	organizationUnitDropdown := dto.DropdownSimple{}
	if item.OrganizationUnitId != 0 {
		organizationUnit, err := getOrganizationUnitById(item.OrganizationUnitId)
		if err != nil {
			return nil, err
		}
		if organizationUnit != nil {
			organizationUnitDropdown = dto.DropdownSimple{Id: organizationUnit.Id, Title: organizationUnit.Title}
		}
	}

	targetOrganizationUnitDropdown := dto.DropdownSimple{}
	if item.TargetOrganizationUnitId != 0 {
		targetOrganizationUnit, err := getOrganizationUnitById(item.TargetOrganizationUnitId)
		if err != nil {
			return nil, err
		}
		if targetOrganizationUnit != nil {
			targetOrganizationUnitDropdown = dto.DropdownSimple{Id: targetOrganizationUnit.Id, Title: targetOrganizationUnit.Title}
		}
	}

	realEstate, err := getMyInventoryRealEstate(item.Id)
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
	assessments, _ := getMyInventoryAssessments(item.Id)
	depreciationTypeId := 0
	var assessmentsResponse []*dto.BasicInventoryResponseAssessment
	for i, assessment := range assessments {
		if assessment.Id != 0 {
			assessmentResponse, err := buildAssessmentResponse(&assessment)
			if i == 0 {
				depreciationTypeId = assessmentResponse.Id
				item.GrossPrice = assessmentResponse.GrossPriceDifference
			}
			if err != nil {
				return nil, err
			}
			assessmentsResponse = append(assessmentsResponse, assessmentResponse)
		}

	}

	settingDropdownDepreciationTypeId := dto.DropdownSimple{}
	lifetimeOfAssessmentInMonths := 0
	amortizationValue := 0
	depreciationRate := 100
	if item.DepreciationTypeId != 0 {
		settings, _ := getDropdownSettingById(depreciationTypeId)

		if settings != nil {
			settingDropdownDepreciationTypeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
			num, _ := strconv.Atoi(settings.Value)
			if num > -1 {
				lifetimeOfAssessmentInMonths = num
			}
			if lifetimeOfAssessmentInMonths > 0 {
				depreciationRate = 100 / lifetimeOfAssessmentInMonths
				layout := time.RFC3339Nano

				t, _ := time.Parse(layout, item.CreatedAt)

				currentTime := time.Now()
				years := currentTime.Year() - t.Year()

				if currentTime.YearDay() < t.YearDay() {
					years--
				}
				if years > 0 {
					amortizationValue = item.GrossPrice / lifetimeOfAssessmentInMonths * years
				}
			}
		}
	}

	itemInventoryList, _ := getDispatchItemByInventoryID(item.Id)

	status := "Lager"
	var movements []*dto.InventoryDispatchResponse
	if item.Type == "movable" {
		if len(itemInventoryList) > 0 {
			for i, move := range itemInventoryList {
				dispatchRes, err := getDispatchItemByID(move.DispatchId)
				if err != nil {
					return nil, err
				}
				if i == 0 && (dispatchRes.TargetOrganizationUnitId == organizationUnitID || dispatchRes.SourceOrganizationUnitId == organizationUnitID) {
					switch dispatchRes.Type {
					case "revers":
						status = "Revers"
					case "allocation":
						status = "Zadužen"
					case "return":
						status = "Lager"
					}
				}
				dispatch, _ := buildInventoryDispatchResponse(dispatchRes)
				movements = append(movements, dispatch)
			}
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
		GrossPrice:                   item.GrossPrice,
		Description:                  item.Description,
		DateOfPurchase:               item.DateOfPurchase,
		Source:                       item.Source,
		DonorTitle:                   item.DonorTitle,
		InvoiceNumber:                item.InvoiceNumber,
		Active:                       item.Active,
		DeactivationDescription:      item.DeactivationDescription,
		DateOfAssessment:             item.DateOfAssessment,
		PriceOfAssessment:            item.PriceOfAssessment,
		LifetimeOfAssessmentInMonths: lifetimeOfAssessmentInMonths,
		DepreciationRate:             fmt.Sprintf("%d%%", depreciationRate),
		AmortizationValue:            amortizationValue,
		OrganizationUnit:             organizationUnitDropdown,
		TargetOrganizationUnit:       targetOrganizationUnitDropdown,
		Status:                       status,
		CreatedAt:                    item.CreatedAt,
		UpdatedAt:                    item.UpdatedAt,
	}

	return &res, nil
}

func getMyInventoryRealEstate(id int) (*structs.BasicInventoryRealEstatesItem, error) {
	res := &dto.GetMyInventoryRealEstateResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.BASIC_INVENTORY_MS_BASE_URL+"/item/"+strconv.Itoa(id)+"/real-estates", nil, res)

	if err != nil {
		fmt.Printf("Fetching Real Estate failed because of this error - %s.\n", err)
		return nil, err
	}

	return &res.Data, nil
}

func getMyInventoryAssessments(id int) ([]structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseArrayMS{}
	_, err := shared.MakeAPIRequest("GET", config.ASSESSMENTS_ENDPOINT+"/"+strconv.Itoa(id)+"/item", nil, res)

	if err != nil {
		if apiErr, ok := err.(*shared.APIError); ok && apiErr.StatusCode != 404 {
			fmt.Printf("Fetching Assessments failed because of this error - %s.\n", err)
			return nil, err
		}
	}

	return res.Data, nil
}

func getDispatchItemByInventoryID(id int) ([]*structs.BasicInventoryDispatchItemsItem, error) {
	res1 := dto.GetAllBasicInventoryDispatchItems{}
	_, err := shared.MakeAPIRequest("GET", config.BASIC_INVENTORY_MS_BASE_URL+"/item/"+strconv.Itoa(id)+"/dispatch-items", nil, &res1)

	if err != nil {
		return nil, err
	}

	return res1.Data, nil
}
