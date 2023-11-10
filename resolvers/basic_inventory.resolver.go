package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/pdfutils"
	"bff/shared"
	"bff/structs"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

var BasicInventoryOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.BasicInventoryResponseListItem
	var filter dto.InventoryItemFilter
	var status string
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

	if st, ok := params.Args["status"].(string); ok && st != "" {
		status = st
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
		if status != "" && resItem.Status != status {
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

	for i, item := range data {
		if i == 0 && item.OrderListId != 0 {
			orderList, _ := getOrderListById(item.OrderListId)
			orderList.IsUsed = true
			_, err := updateOrderListItem(item.OrderListId, orderList)
			if err != nil {
				return shared.HandleAPIError(err)
			}
		}
		item.Active = true
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
		if inactive, ok := params.Args["inactive"].(string); ok && inactive != "" {
			item.Inactive = &inactive
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
	var grossPrice float32
	indexAssessments := 0
	if len(assessments) > 0 {
		for i, assessment := range assessments {
			if assessment.Id != 0 {
				assessmentResponse, _ := buildAssessmentResponse(&assessment)
				if assessmentResponse != nil && i == indexAssessments && assessmentResponse.Type == "financial" {

					grossPrice = assessmentResponse.GrossPriceDifference

					settings, _ := getDropdownSettingById(assessments[0].DepreciationTypeId)

					if settings != nil {
						settingDropdownDepreciationTypeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
					}
					break
				} else {
					indexAssessments++
				}
			}

		}

	}

	status := "Lager"

	if item.Type == "movable" && item.Active {
		itemInventoryList, _ := getDispatchItemByInventoryID(item.Id)
		if len(itemInventoryList) > 0 {
			for _, move := range itemInventoryList {

				dispatchRes, err := getDispatchItemByID(move.DispatchId)
				if err != nil {
					return nil, err
				}
				if dispatchRes.TargetOrganizationUnitId == organizationUnitID || dispatchRes.SourceOrganizationUnitId == organizationUnitID {
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
					break
				}
			}
		}
	}
	if !item.Active {
		status = "Deaktiviran"
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
		Active:                 item.Active,
		Type:                   item.Type,
		Title:                  item.Title,
		Location:               item.Location,
		InventoryNumber:        item.InventoryNumber,
		GrossPrice:             grossPrice,
		PurchaseGrossPrice:     item.GrossPrice,
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
	var grossPrice float32
	indexAssessments := 0
	var assessmentsResponse []*dto.BasicInventoryResponseAssessment
	for i, assessment := range assessments {
		if assessment.Id != 0 {
			assessmentResponse, _ := buildAssessmentResponse(&assessment)
			if assessmentResponse != nil && i == indexAssessments && assessmentResponse.Type == "financial" {
				depreciationTypeId = assessmentResponse.DepreciationType.Id
				grossPrice = assessmentResponse.GrossPriceDifference
			} else {
				indexAssessments++
			}
			assessmentsResponse = append(assessmentsResponse, assessmentResponse)
		}

	}

	settingDropdownDepreciationTypeId := dto.DropdownSimple{}
	lifetimeOfAssessmentInMonths := 0
	var amortizationValue float32
	depreciationRate := 100
	if depreciationTypeId != 0 {
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
					amortizationValue = grossPrice / float32(lifetimeOfAssessmentInMonths) * float32(years)
				}
			}
		}
	}

	itemInventoryList, _ := getDispatchItemByInventoryID(item.Id)

	status := "Lager"
	var movements []*dto.InventoryDispatchResponse
	indexMovements := 0
	if len(itemInventoryList) > 0 {
		for i, move := range itemInventoryList {
			dispatchRes, err := getDispatchItemByID(move.DispatchId)
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
					status = "Lager"
				}
			} else {
				indexAssessments++
			}
			dispatch, _ := buildInventoryDispatchResponse(dispatchRes)
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

var FormNS1PDFResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return shared.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	organizationUnit, err := getOrganizationUnitById(*organizationUnitID)
	if err != nil {
		return nil, err
	}

	Item, err := getInventoryItem(id)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	items, err := buildInventoryItemResponse(Item, *organizationUnitID)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	c := creator.New()
	c.SetPageMargins(50, 50, 50, 30)
	c.NewPage()

	// Define fonts
	fontRegular, err := model.NewCompositePdfFontFromTTFFile(config.BASE_APP_DIR + "fonts/RobotoSlab-VariableFont_wght.ttf")
	if err != nil {
		return shared.HandleAPIError(err)
	}

	fontBold, err := model.NewCompositePdfFontFromTTFFile(config.BASE_APP_DIR + "fonts/RobotoSlab-Bold.ttf")
	if err != nil {
		return shared.HandleAPIError(err)
	}
	textTitle := "Obrazac " + items.SourceType

	// Add Title
	title, err := pdfutils.CreateTitle(c, textTitle, fontBold)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	err = c.Draw(title)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	// Create table with a single column.
	table := c.NewTable(1)
	table.SetMargins(0, 0, 10, 0)

	//title
	pTitleParagraph := c.NewStyledParagraph()
	pTitleParagraph.SetMargins(5, 5, 10, 10)
	pTitleParagraph.SetTextAlignment(creator.TextAlignmentCenter)
	var pTitleParagraphBoldText string
	var pTitleParagraphText string

	if items.SourceType == "NS1" {
		pTitleParagraphBoldText = "Korisnik nepokretnih stvari u državnoj svojini "
		pTitleParagraphText = "(državni organi,organi lokalne samouprave i javne službe čiji je osnivač Crna Gora, odnosno lokalna samouprava)"
	} else {
		pTitleParagraphBoldText = "Organi u čijoj su nadležnosti nepokretne stvari "
		pTitleParagraphText = "(organi u čijoj su nadležnosti nepokretne stvari za koje vrše popis, odnosno identifikaciju)"
	}

	pTitleParagraph.Append(pTitleParagraphBoldText).Style.Font = fontBold
	pTitleParagraph.Append(pTitleParagraphText).Style.Font = fontRegular
	cell := table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	_ = cell.SetContent(pTitleParagraph)

	//Name of Organization unit
	pNameOUParagraph := c.NewStyledParagraph()
	pNameOUParagraph.SetMargins(5, 5, 7, 7)
	pNameOUParagraphText := "1. Naziv: "
	pNameOUParagraphBoldText := items.OrganizationUnit.Title

	pNameOUParagraph.Append(pNameOUParagraphText).Style.Font = fontRegular
	pNameOUParagraph.Append(pNameOUParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pNameOUParagraph)

	//City
	pCityOUParagraph := c.NewStyledParagraph()
	pCityOUParagraph.SetMargins(5, 5, 7, 7)
	pCityOUParagraphText := "2. Sjedište (mjesto,opština): "
	pCityOUParagraphBoldText := organizationUnit.City

	pCityOUParagraph.Append(pCityOUParagraphText).Style.Font = fontRegular
	pCityOUParagraph.Append(pCityOUParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pCityOUParagraph)

	//Address
	pAddressOUParagraph := c.NewStyledParagraph()
	pAddressOUParagraph.SetMargins(5, 5, 7, 7)
	pAddressOUParagraphText := "3. Adresa (ulica,broj,sprat,kancelarija): "
	pAddressOUParagraphBoldText := organizationUnit.Address

	pAddressOUParagraph.Append(pAddressOUParagraphText).Style.Font = fontRegular
	pAddressOUParagraph.Append(pAddressOUParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)

	_ = cell.SetContent(pAddressOUParagraph)

	//Djelatnost
	pDjelatnostParagraph := c.NewStyledParagraph()
	pDjelatnostParagraph.SetMargins(5, 5, 7, 7)
	pDjelatnostParagraphText := "4. Djelatnost (šifra): "
	pDjelatnostParagraphBoldText := "84.23 Sudske i pravosudne djelatnosti"

	pDjelatnostParagraph.Append(pDjelatnostParagraphText).Style.Font = fontRegular
	pDjelatnostParagraph.Append(pDjelatnostParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pDjelatnostParagraph)

	//Movable inventory title
	pMovableTitleParagraph := c.NewStyledParagraph()
	pMovableTitleParagraph.SetMargins(5, 5, 7, 7)
	pMovableTitleParagraph.SetTextAlignment(creator.TextAlignmentCenter)
	pMovableTitleParagraphBoldText := "Nepokretne stvari"

	pMovableTitleParagraph.Append(pMovableTitleParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	_ = cell.SetContent(pMovableTitleParagraph)

	//Name Inventory
	pNameInventoryParagraph := c.NewStyledParagraph()
	pNameInventoryParagraph.SetMargins(5, 5, 7, 7)
	pNameInventoryParagraphText := "1. Vrsta (oprema, prevozna sredstva i druge pokretne stvari koje se koriste za obavljanje funkcije): "
	pNameInventoryParagraphBoldText := items.RealEstate.TypeId

	pNameInventoryParagraph.Append(pNameInventoryParagraphText).Style.Font = fontRegular
	pNameInventoryParagraph.Append(pNameInventoryParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pNameInventoryParagraph)

	//Amount
	pAmountParagraph := c.NewStyledParagraph()
	pAmountParagraph.SetMargins(5, 5, 7, 7)
	pAmountParagraphText := "2. . Mjesto gdje se nepokretnost nalazi: "
	pAmountParagraphBoldText := items.Location

	pAmountParagraph.Append(pAmountParagraphText).Style.Font = fontRegular
	pAmountParagraph.Append(pAmountParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pAmountParagraph)

	//InventoryNumber
	pInventoryNumberParagraph := c.NewStyledParagraph()
	pInventoryNumberParagraph.SetMargins(5, 5, 7, 7)
	pInventoryNumberParagraphText := "3. Površina (u m2): "
	pInventoryNumberParagraphBoldText := fmt.Sprintf("%v", items.RealEstate.SquareArea)

	pInventoryNumberParagraph.Append(pInventoryNumberParagraphText).Style.Font = fontRegular
	pInventoryNumberParagraph.Append(pInventoryNumberParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pInventoryNumberParagraph)

	//Source
	pSourceParagraph := c.NewStyledParagraph()
	pSourceParagraph.SetMargins(5, 5, 7, 7)
	pSourceParagraphText := "4. Broj katastarske parcele, list nepokretnosti i katastarska opština: "
	pSourceParagraphBoldText := items.RealEstate.LandSerialNumber

	pSourceParagraph.Append(pSourceParagraphText).Style.Font = fontRegular
	pSourceParagraph.Append(pSourceParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pSourceParagraph)

	//PropertyDocument
	pPropertyDocumentParagraph := c.NewStyledParagraph()
	pPropertyDocumentParagraph.SetMargins(5, 5, 7, 7)
	pPropertyDocumentParagraphText := "5. Isprave o svojini (osnov sticanja i prestanka prava svojine): "
	pPropertyDocumentParagraphBoldText := items.RealEstate.PropertyDocument

	pPropertyDocumentParagraph.Append(pPropertyDocumentParagraphText).Style.Font = fontRegular
	pPropertyDocumentParagraph.Append(pPropertyDocumentParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pPropertyDocumentParagraph)

	//PurchaseGrossPrice
	pOwnershipScopeParagraph := c.NewStyledParagraph()
	pOwnershipScopeParagraph.SetMargins(5, 5, 7, 7)
	pOwnershipScopeParagraphText := "6. Obim prava: "
	pOwnershipScopeParagraphBoldText := items.RealEstate.OwnershipScope

	pOwnershipScopeParagraph.Append(pOwnershipScopeParagraphText).Style.Font = fontRegular
	pOwnershipScopeParagraph.Append(pOwnershipScopeParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pOwnershipScopeParagraph)

	//OwnershipInvestmentScope
	indexTable := 7
	if items.SourceType == "NS1" {
		pOwnershipInvestmentScopeParagraph := c.NewStyledParagraph()
		pOwnershipInvestmentScopeParagraph.SetMargins(5, 5, 7, 7)
		pOwnershipInvestmentScopeParagraphText := fmt.Sprintf("%d. Obim prava za imovinu stečenu zajedničkim ulaganjem: ", indexTable)
		pOwnershipInvestmentScopeParagraphBoldText := items.RealEstate.OwnershipInvestmentScope

		pOwnershipInvestmentScopeParagraph.Append(pOwnershipInvestmentScopeParagraphText).Style.Font = fontRegular
		pOwnershipInvestmentScopeParagraph.Append(pOwnershipInvestmentScopeParagraphBoldText).Style.Font = fontBold
		cell = table.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
		_ = cell.SetContent(pOwnershipInvestmentScopeParagraph)
		indexTable++
	}

	//LifetimeOfAssessmentInMonths

	pLifetimeOfAssessmentInMonthsParagraph := c.NewStyledParagraph()
	pLifetimeOfAssessmentInMonthsParagraph.SetMargins(5, 5, 7, 7)
	pLifetimeOfAssessmentInMonthsParagraphText := fmt.Sprintf("%d. Vijek trajanja: ", indexTable)
	pLifetimeOfAssessmentInMonthsParagraphBoldText := fmt.Sprintf("%d", items.LifetimeOfAssessmentInMonths)

	pLifetimeOfAssessmentInMonthsParagraph.Append(pLifetimeOfAssessmentInMonthsParagraphText).Style.Font = fontRegular
	pLifetimeOfAssessmentInMonthsParagraph.Append(pLifetimeOfAssessmentInMonthsParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pLifetimeOfAssessmentInMonthsParagraph)

	//PurchaseGrossPrice
	indexTable++
	pPurchaseGrossPriceParagraph := c.NewStyledParagraph()
	pPurchaseGrossPriceParagraph.SetMargins(5, 5, 7, 7)
	pPurchaseGrossPriceParagraphText := fmt.Sprintf("%d. Nabavna vrijednost: ", indexTable)
	pPurchaseGrossPriceParagraphBoldText := fmt.Sprintf("€%v", items.PurchaseGrossPrice)

	pPurchaseGrossPriceParagraph.Append(pPurchaseGrossPriceParagraphText).Style.Font = fontRegular
	pPurchaseGrossPriceParagraph.Append(pPurchaseGrossPriceParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pPurchaseGrossPriceParagraph)

	//AmortizationValue
	indexTable++
	pAmortizationValueParagraph := c.NewStyledParagraph()
	pAmortizationValueParagraph.SetMargins(5, 5, 7, 7)
	pAmortizationValueParagraphText := fmt.Sprintf("%d. Ispravka vrijednosti (ispravka/otpis vrijednosti predhodnih godina + amortizacija tekuće godine): ", indexTable)
	pAmortizationValueParagraphBoldText := fmt.Sprintf("€%v", items.AmortizationValue)

	pAmortizationValueParagraph.Append(pAmortizationValueParagraphText).Style.Font = fontRegular
	pAmortizationValueParagraph.Append(pAmortizationValueParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pAmortizationValueParagraph)

	//GrossPrice
	indexTable++
	pGrossPriceParagraph := c.NewStyledParagraph()
	pGrossPriceParagraph.SetMargins(5, 5, 7, 7)
	pGrossPriceParagraphText := fmt.Sprintf("%d. Knjigovodstvena vrijednost / fer vrijednost (procijenjena vrijednost): ", indexTable)
	pGrossPriceParagraphBoldText := fmt.Sprintf("€%v", items.GrossPrice)

	pGrossPriceParagraph.Append(pGrossPriceParagraphText).Style.Font = fontRegular
	pGrossPriceParagraph.Append(pGrossPriceParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pGrossPriceParagraph)

	//LimitationsDescription
	if items.SourceType == "NS1" {
		indexTable++
		pLimitationsDescriptionParagraph := c.NewStyledParagraph()
		pLimitationsDescriptionParagraph.SetMargins(5, 5, 7, 7)
		pLimitationsDescriptionParagraphText := fmt.Sprintf("%d. Tereti i ograničenja (založna prava, službenosti,restitucija i dr): ", indexTable)
		pLimitationsDescriptionParagraphBoldText := items.RealEstate.LimitationsDescription

		pLimitationsDescriptionParagraph.Append(pLimitationsDescriptionParagraphText).Style.Font = fontRegular
		pLimitationsDescriptionParagraph.Append(pLimitationsDescriptionParagraphBoldText).Style.Font = fontBold
		cell = table.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
		_ = cell.SetContent(pLimitationsDescriptionParagraph)
	}

	//Note
	pNoteParagraph := c.NewStyledParagraph()
	pNoteParagraph.SetMargins(5, 5, 7, 7)
	pNoteParagraphBoldText := "Napomena: "

	pNoteParagraph.Append(pNoteParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pNoteParagraph)

	//Footer
	pFooterParagraph := c.NewStyledParagraph()
	pFooterParagraph.SetMargins(5, 5, 7, 7)
	pFooterLeftParagraphText := "Datum _____________                                     M.P                                     Starješina organa______________"

	cell = table.NewCell()
	pFooterParagraph.Append(pFooterLeftParagraphText).Style.Font = fontRegular
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	_ = cell.SetContent(pFooterParagraph)

	err = c.Draw(table)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	var buf bytes.Buffer
	err = c.Write(&buf)
	if err != nil {
		return nil, err
	}

	encodedStr := base64.StdEncoding.EncodeToString(buf.Bytes())
	// Return the path or a URL to the file
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's is your PDF file in base64 encode format!",
		Item:    encodedStr,
	}, nil
}

var FormPS1PDFResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return shared.HandleAPIError(fmt.Errorf("user does not have organization unit assigned"))
	}

	organizationUnit, err := getOrganizationUnitById(*organizationUnitID)
	if err != nil {
		return nil, err
	}

	Item, err := getInventoryItem(id)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	items, err := buildInventoryItemResponse(Item, *organizationUnitID)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	c := creator.New()
	c.SetPageMargins(50, 50, 50, 50)
	c.NewPage()

	// Define fonts
	fontRegular, err := model.NewCompositePdfFontFromTTFFile(config.BASE_APP_DIR + "fonts/RobotoSlab-VariableFont_wght.ttf")
	if err != nil {
		return shared.HandleAPIError(err)
	}

	fontBold, err := model.NewCompositePdfFontFromTTFFile(config.BASE_APP_DIR + "fonts/RobotoSlab-Bold.ttf")
	if err != nil {
		return shared.HandleAPIError(err)
	}
	textTitle := "Obrazac " + items.SourceType

	// Add Title
	title, err := pdfutils.CreateTitle(c, textTitle, fontBold)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	err = c.Draw(title)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	// Create table with a single column.
	table := c.NewTable(1)
	table.SetMargins(0, 0, 10, 0)

	//title
	pTitleParagraph := c.NewStyledParagraph()
	pTitleParagraph.SetMargins(5, 5, 10, 10)
	pTitleParagraph.SetTextAlignment(creator.TextAlignmentCenter)
	var pTitleParagraphBoldText string
	var pTitleParagraphText string
	if items.SourceType == "PS1" {
		pTitleParagraphBoldText = "Korisnik pokretnih stvari u državnoj svojini "
		pTitleParagraphText = "(državni organi,organi lokalne samouprave i javne službe čiji je osnivač Crna Gora, odnosno lokalna samouprava)"
	} else {
		pTitleParagraphBoldText = "Organi u čijoj su nadležnosti pokretne stvari "
		pTitleParagraphText = "(organi u čijoj su nadležnosti pokretne stvari za koje vrše popis,odnosno indentifikacija)"
	}
	pTitleParagraph.Append(pTitleParagraphBoldText).Style.Font = fontBold
	pTitleParagraph.Append(pTitleParagraphText).Style.Font = fontRegular
	cell := table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	_ = cell.SetContent(pTitleParagraph)

	//Name of Organization unit
	pNameOUParagraph := c.NewStyledParagraph()
	pNameOUParagraph.SetMargins(5, 5, 10, 10)
	pNameOUParagraphText := "1. Naziv: "
	pNameOUParagraphBoldText := items.OrganizationUnit.Title

	pNameOUParagraph.Append(pNameOUParagraphText).Style.Font = fontRegular
	pNameOUParagraph.Append(pNameOUParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pNameOUParagraph)

	//City
	pCityOUParagraph := c.NewStyledParagraph()
	pCityOUParagraph.SetMargins(5, 5, 10, 10)
	pCityOUParagraphText := "2. Sjedište (mjesto,opština): "
	pCityOUParagraphBoldText := organizationUnit.City

	pCityOUParagraph.Append(pCityOUParagraphText).Style.Font = fontRegular
	pCityOUParagraph.Append(pCityOUParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pCityOUParagraph)

	//Address
	pAddressOUParagraph := c.NewStyledParagraph()
	pAddressOUParagraph.SetMargins(5, 5, 10, 10)
	pAddressOUParagraphText := "3. Adresa (ulica,broj,sprat,kancelarija): "
	pAddressOUParagraphBoldText := organizationUnit.Address

	pAddressOUParagraph.Append(pAddressOUParagraphText).Style.Font = fontRegular
	pAddressOUParagraph.Append(pAddressOUParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)

	_ = cell.SetContent(pAddressOUParagraph)

	//Djelatnost
	pDjelatnostParagraph := c.NewStyledParagraph()
	pDjelatnostParagraph.SetMargins(5, 5, 10, 10)
	pDjelatnostParagraphText := "4. Djelatnost (šifra): "
	pDjelatnostParagraphBoldText := "84.23 Sudske i pravosudne djelatnosti"

	pDjelatnostParagraph.Append(pDjelatnostParagraphText).Style.Font = fontRegular
	pDjelatnostParagraph.Append(pDjelatnostParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pDjelatnostParagraph)

	//Movable inventory title
	pMovableTitleParagraph := c.NewStyledParagraph()
	pMovableTitleParagraph.SetMargins(5, 5, 10, 10)
	pMovableTitleParagraph.SetTextAlignment(creator.TextAlignmentCenter)
	pMovableTitleParagraphBoldText := "Pokretne stvari"

	pMovableTitleParagraph.Append(pMovableTitleParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	_ = cell.SetContent(pMovableTitleParagraph)

	//Name Inventory
	pNameInventoryParagraph := c.NewStyledParagraph()
	pNameInventoryParagraph.SetMargins(5, 5, 10, 10)
	pNameInventoryParagraphText := "1. Vrsta (oprema, prevozna sredstva i druge pokretne stvari koje se koriste za obavljanje funkcije): "
	pNameInventoryParagraphBoldText := items.Title

	pNameInventoryParagraph.Append(pNameInventoryParagraphText).Style.Font = fontRegular
	pNameInventoryParagraph.Append(pNameInventoryParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pNameInventoryParagraph)

	//Amount
	pAmountParagraph := c.NewStyledParagraph()
	pAmountParagraph.SetMargins(5, 5, 10, 10)
	pAmountParagraphText := "2. Količina,komad i broj: "
	pAmountParagraphBoldText := "1"

	pAmountParagraph.Append(pAmountParagraphText).Style.Font = fontRegular
	pAmountParagraph.Append(pAmountParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pAmountParagraph)

	//InventoryNumber
	pInventoryNumberParagraph := c.NewStyledParagraph()
	pInventoryNumberParagraph.SetMargins(5, 5, 10, 10)
	pInventoryNumberParagraphText := "3. Inventarski broj: "
	pInventoryNumberParagraphBoldText := items.InventoryNumber

	pInventoryNumberParagraph.Append(pInventoryNumberParagraphText).Style.Font = fontRegular
	pInventoryNumberParagraph.Append(pInventoryNumberParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pInventoryNumberParagraph)

	//Source
	pSourceParagraph := c.NewStyledParagraph()
	pSourceParagraph.SetMargins(5, 5, 10, 10)
	pSourceParagraphText := "4. Način sticanja: "
	pSourceParagraphBoldText := items.Source

	pSourceParagraph.Append(pSourceParagraphText).Style.Font = fontRegular
	pSourceParagraph.Append(pSourceParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pSourceParagraph)

	//LifetimeOfAssessmentInMonths
	pLifetimeOfAssessmentInMonthsParagraph := c.NewStyledParagraph()
	pLifetimeOfAssessmentInMonthsParagraph.SetMargins(5, 5, 10, 10)
	pLifetimeOfAssessmentInMonthsParagraphText := "5. Vijek trajanja: "
	pLifetimeOfAssessmentInMonthsParagraphBoldText := fmt.Sprintf("%d", items.LifetimeOfAssessmentInMonths)

	pLifetimeOfAssessmentInMonthsParagraph.Append(pLifetimeOfAssessmentInMonthsParagraphText).Style.Font = fontRegular
	pLifetimeOfAssessmentInMonthsParagraph.Append(pLifetimeOfAssessmentInMonthsParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pLifetimeOfAssessmentInMonthsParagraph)

	//PurchaseGrossPrice
	pPurchaseGrossPriceParagraph := c.NewStyledParagraph()
	pPurchaseGrossPriceParagraph.SetMargins(5, 5, 10, 10)
	pPurchaseGrossPriceParagraphText := "6. Nabavna vrijednost: "
	pPurchaseGrossPriceParagraphBoldText := fmt.Sprintf("€%v", items.PurchaseGrossPrice)

	pPurchaseGrossPriceParagraph.Append(pPurchaseGrossPriceParagraphText).Style.Font = fontRegular
	pPurchaseGrossPriceParagraph.Append(pPurchaseGrossPriceParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pPurchaseGrossPriceParagraph)

	//AmortizationValue
	pAmortizationValueParagraph := c.NewStyledParagraph()
	pAmortizationValueParagraph.SetMargins(5, 5, 10, 10)
	pAmortizationValueParagraphText := "7. Ispravka  vrijednosti(ispravka/otpis vrijednosti predhodnih godina + amortizacija) : "
	pAmortizationValueParagraphBoldText := fmt.Sprintf("€%v", items.AmortizationValue)

	pAmortizationValueParagraph.Append(pAmortizationValueParagraphText).Style.Font = fontRegular
	pAmortizationValueParagraph.Append(pAmortizationValueParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pAmortizationValueParagraph)

	//GrossPrice
	pGrossPriceParagraph := c.NewStyledParagraph()
	pGrossPriceParagraph.SetMargins(5, 5, 10, 10)
	pGrossPriceParagraphText := "8. Knjigovodstvena vrijednost / fer vrijednost (procijenjena vrijednost): "
	pGrossPriceParagraphBoldText := fmt.Sprintf("€%v", items.GrossPrice)

	pGrossPriceParagraph.Append(pGrossPriceParagraphText).Style.Font = fontRegular
	pGrossPriceParagraph.Append(pGrossPriceParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pGrossPriceParagraph)

	//Inactive
	pInactiveParagraph := c.NewStyledParagraph()
	pInactiveParagraph.SetMargins(5, 5, 10, 10)
	pInactiveParagraphText := "9. Broj i datum odluke, o utvrdjenom manjku,višku ili rashodovanju stvari: "
	pInactiveParagraphBoldText := fmt.Sprintf("€%v", items.Inactive)

	pInactiveParagraph.Append(pInactiveParagraphText).Style.Font = fontRegular
	pInactiveParagraph.Append(pInactiveParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pInactiveParagraph)

	//Note
	pNoteParagraph := c.NewStyledParagraph()
	pNoteParagraph.SetMargins(5, 5, 10, 10)
	pNoteParagraphBoldText := "Napomena: "

	pNoteParagraph.Append(pNoteParagraphBoldText).Style.Font = fontBold
	cell = table.NewCell()
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	_ = cell.SetContent(pNoteParagraph)

	//Footer
	pFooterParagraph := c.NewStyledParagraph()
	pFooterParagraph.SetMargins(5, 5, 10, 10)
	pFooterLeftParagraphText := "Datum _____________                                     M.P                                     Starješina organa______________"

	cell = table.NewCell()
	pFooterParagraph.Append(pFooterLeftParagraphText).Style.Font = fontRegular
	cell.SetHorizontalAlignment(creator.CellHorizontalAlignmentLeft)
	cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
	_ = cell.SetContent(pFooterParagraph)

	err = c.Draw(table)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	var buf bytes.Buffer
	err = c.Write(&buf)
	if err != nil {
		return nil, err
	}

	encodedStr := base64.StdEncoding.EncodeToString(buf.Bytes())
	// Return the path or a URL to the file
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's is your PDF file in base64 encode format!",
		Item:    encodedStr,
	}, nil
}
