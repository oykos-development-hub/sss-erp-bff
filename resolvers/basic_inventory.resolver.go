package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/graphql-go/graphql"
)

var BasicInventoryOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var items []*dto.BasicInventoryResponseItem
	var filter dto.InventoryItemFilter

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
		filter.SourceType = &sourceType
	}

	if depreciationTypeID, ok := params.Args["depreciation_type_id"].(int); ok && depreciationTypeID != 0 {
		filter.DeprecationTypeID = &depreciationTypeID
	}

	if organizationUnitID, ok := params.Args["organization_unit"].(int); ok && organizationUnitID != 0 {
		filter.OrganizationUnitID = &organizationUnitID
	}

	if page, ok := params.Args["page"].(int); ok && page != 0 {
		filter.Page = &page
	}

	if size, ok := params.Args["size"].(int); ok && size != 0 {
		filter.Size = &size
	}

	basicInventoryData, err := getAllInventoryItem(filter)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, item := range basicInventoryData.Data {
		resItem, err := buildInventoryItemResponse(item)
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

	Item, err := getInventoryItem(id)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	items, err := buildInventoryItemResponse(Item)

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

			items, err := buildInventoryItemResponse(itemRes)

			if err != nil {
				return shared.HandleAPIError(err)
			}

			responseItemList = append(responseItemList, items)
		} else {
			itemRes, err := createInventoryItem(&item)
			if err != nil {
				return shared.HandleAPIError(err)
			}

			response.Message = "You created this item/s!"
			items, err := buildInventoryItemResponse(itemRes)

			if err != nil {
				return shared.HandleAPIError(err)
			}
			responseItemList = append(responseItemList, items)
		}
	}
	response.Items = responseItemList
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

func buildInventoryItemResponse(item *structs.BasicInventoryInsertItem) (*dto.BasicInventoryResponseItem, error) {
	settings, err := getDropdownSettingById(item.ClassTypeId)
	if err != nil {
		return nil, err
	}

	settingDropdownClassType := dto.DropdownSimple{}
	if settings != nil {
		settingDropdownClassType = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
	}

	settings, err = getDropdownSettingById(item.DepreciationTypeId)
	if err != nil {
		return nil, err
	}

	settingDropdownDepreciationTypeId := dto.DropdownSimple{}
	if settings != nil {
		settingDropdownDepreciationTypeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
	}

	suppliers, err := getSupplier(item.SupplierId)
	if err != nil {
		return nil, err
	}

	suppliersDropdown := dto.DropdownSimple{}
	if suppliers != nil {
		suppliersDropdown = dto.DropdownSimple{Id: suppliers.Id, Title: suppliers.Title}
	}

	settings, err = getDropdownSettingById(item.OfficeId)
	if err != nil {
		return nil, err
	}

	settingDropdownOfficeId := dto.DropdownSimple{}
	if settings != nil {
		settingDropdownOfficeId = dto.DropdownSimple{Id: settings.Id, Title: settings.Title}
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
	assessment, err := getMyInventoryAssessments(item.Id)

	if err != nil {
		return nil, err
	}
	assessmentResponse := (*dto.BasicInventoryResponseAssessment)(nil) // Inicijalizacija sa nil

	if assessment.Id != 0 {
		assessmentResponse, err = buildAssessmentResponse(assessment)

		if err != nil {
			return nil, err
		}
	} else {
		assessmentResponse = nil
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
		Assessments:                  assessmentResponse,
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
		LifetimeOfAssessmentInMonths: item.LifetimeOfAssessmentInMonths,
		OrganizationUnit:             organizationUnitDropdown,
		TargetOrganizationUnit:       targetOrganizationUnitDropdown,
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

func getMyInventoryAssessments(id int) (*structs.BasicInventoryAssessmentsTypesItem, error) {
	res := &dto.AssessmentResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.ASSESSMENTS_ENDPOINT+"/"+strconv.Itoa(id)+"/item", nil, res)

	if err != nil {
		fmt.Printf("Fetching Assessments failed because of this error - %s.\n", err)
		return nil, err
	}

	return &res.Data, nil
}
