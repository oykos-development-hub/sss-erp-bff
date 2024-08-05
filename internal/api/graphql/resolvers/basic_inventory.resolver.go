package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
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
		filter.SourceType = &sourceType
	}

	if depreciationTypeID, ok := params.Args["depreciation_type_id"].(int); ok && depreciationTypeID != 0 {
		filter.DeprecationTypeID = &depreciationTypeID
	}
	if expire, ok := params.Args["expire"].(bool); ok && expire {
		filter.Expire = &expire
	}
	if isExternalDonation, ok := params.Args["is_external_donation"].(bool); ok && isExternalDonation {
		filter.IsExternalDonation = &isExternalDonation
	}

	if status, ok := params.Args["status"].(string); ok && status != "" {
		filter.Status = &status
	}

	if typeImmovement, ok := params.Args["type_of_immovable_property"].(string); ok && typeImmovement != "" {
		filter.TypeOfImmovableProperty = &typeImmovement
	}

	if organizationUnitID, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitID != 0 && status != "Arhiva" {
		filter.OrganizationUnitID = &organizationUnitID
	}

	if page, ok := params.Args["page"].(int); ok && page != 0 {
		filter.Page = &page
	}

	if size, ok := params.Args["size"].(int); ok && size != 0 {
		filter.Size = &size
	}

	var organizationUnitID *int
	if filter.OrganizationUnitID != nil {
		organizationUnitID = filter.OrganizationUnitID
	} else {
		organizationUnitIDFromParams, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitIDFromParams == nil {
			return errors.HandleAPPError(fmt.Errorf("user does not have organization unit assigned"))
		}
		organizationUnitID = organizationUnitIDFromParams
	}

	filter.CurrentOrganizationUnit = *organizationUnitID
	basicInventoryData, err := r.Repo.GetAllInventoryItem(filter)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, item := range basicInventoryData.Data {
		resItem, err := buildInventoryResponse(r.Repo, item, *organizationUnitID)

		items = append(items, resItem)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   basicInventoryData.Total,
		Items:   items,
	}, nil
}

func (r *Resolver) BasicInventoryDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	organizationUnitID, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
	if !ok || organizationUnitID == nil {
		return errors.HandleAPPError(fmt.Errorf("user does not have organization unit assigned"))
	}

	Item, err := r.Repo.GetInventoryItem(id)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	items, err := buildInventoryItemResponse(r.Repo, Item, *organizationUnitID)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
		return errors.HandleAPPError(fmt.Errorf("user does not have organization unit assigned"))
	}

	dataBytes, _ := json.Marshal(params.Args["data"])
	err := json.Unmarshal(dataBytes, &data)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	responseItem, err := r.Repo.CheckInsertInventoryData(data)

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
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
			itemRes, err := r.Repo.UpdateInventoryItem(params.Context, item.ID, &item)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			response.Message = "You updated this item/s!"

			items, err := buildInventoryItemResponse(r.Repo, itemRes, 0)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			responseItemList = append(responseItemList, items)
		} else {
			item.ArticleID = item.ContractArticleID
			item.OrganizationUnitID = *organizationUnitID
			itemRes, err := r.Repo.CreateInventoryItem(params.Context, &item)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			if item.DepreciationTypeID != 0 {

				depreciationType, err := r.Repo.GetDropdownSettingByID(item.DepreciationTypeID)

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
				value, err := strconv.Atoi(depreciationType.Value)

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
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

				_, err = r.Repo.CreateAssessments(params.Context, &assessment)
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}

			}

			typeOfDispatch := "created"

			if item.Type == "small" {
				typeOfDispatch = "allocation"
			}

			dispatch := structs.BasicInventoryDispatchItem{
				Type:                     typeOfDispatch,
				SourceUserProfileID:      item.TargetUserProfileID,
				SourceOrganizationUnitID: *organizationUnitID,
				Date:                     itemRes.CreatedAt,
				InventoryID:              []int{itemRes.ID},
				OfficeID:                 item.OfficeID,
			}

			_, err = r.Repo.CreateDispatchItem(params.Context, &dispatch)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			response.Message = "You created this item/s!"
			items, err := buildInventoryItemResponse(r.Repo, itemRes, 0)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			responseItemList = append(responseItemList, items)
		}

	}
	response.Items = responseItemList
	return response, nil
}

func (r *Resolver) InvoicesForInventoryOverview(params graphql.ResolveParams) (interface{}, error) {
	response := dto.Response{
		Status: "success",
	}

	supplierID := params.Args["supplier_id"].(int)
	organizationUnitID, ok := params.Args["organization_unit_id"].(int)

	if !ok {
		organizationUnitIDPtr, ok := params.Context.Value(config.OrganizationUnitIDKey).(*int)
		if !ok || organizationUnitIDPtr == nil {
			return errors.HandleAPPError(fmt.Errorf("user does not have organization unit assigned"))
		}

		organizationUnitID = *organizationUnitIDPtr
	}

	invoices, _, err := r.Repo.GetInvoiceList(&dto.GetInvoiceListInputMS{
		SupplierID:         &supplierID,
		OrganizationUnitID: &organizationUnitID,
	})

	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var invoicesResponse []structs.Invoice

	for _, invoice := range invoices {

		invoiceResponse, _ := buildInvoiceResponseItem(r, invoice)
		var invoiceArticles []structs.InvoiceArticles
		for _, article := range invoiceResponse.Articles {
			data, err := r.Repo.GetAllInventoryItem(dto.InventoryItemFilter{
				InvoiceArticleID: &article.ID,
			})

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			if article.Amount-data.Total > 0 {
				invoiceArticle := structs.InvoiceArticles{
					ID:       article.ID,
					Title:    article.Title,
					NetPrice: article.NetPrice,
					Amount:   article.Amount - data.Total,
				}

				invoiceArticles = append(invoiceArticles, invoiceArticle)
			}
		}

		if len(invoiceArticles) > 0 {
			invoicesResponse = append(invoicesResponse, structs.Invoice{
				ID:            invoice.ID,
				InvoiceNumber: invoice.InvoiceNumber,
				DateOfInvoice: invoice.DateOfInvoice,
				Articles:      invoiceArticles,
			})
		}
	}

	response.Items = invoicesResponse

	return response, nil
}

func (r *Resolver) BasicInventoryDeactivateResolver(params graphql.ResolveParams) (interface{}, error) {
	response := dto.Response{
		Status: "success",
	}
	if id, ok := params.Args["id"].(int); ok && id != 0 {
		item, err := r.Repo.GetInventoryItem(id)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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

		_, err = r.Repo.UpdateInventoryItem(params.Context, id, item)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
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
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}

		if settings != nil {
			settingDropdownOfficeID = dto.DropdownSimple{ID: settings.ID, Title: settings.Title}
		}
	}

	settingDropdownDepreciationTypeID := dto.DropdownSimple{}

	if item.DepreciationTypeID != 0 {
		depreciationTypeDropDown, err := r.GetDropdownSettingByID(item.DepreciationTypeID)

		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}
		settingDropdownDepreciationTypeID.ID = depreciationTypeDropDown.ID
		settingDropdownDepreciationTypeID.Title = depreciationTypeDropDown.Title
		value, err := strconv.Atoi(depreciationTypeDropDown.Value)

		if err != nil {
			return nil, errors.Wrap(err, "strconv atoi")
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
	var dateOfEndOfAssessment string
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
					date, err := time.Parse(time.RFC3339, dateOfAssessment)
					if err != nil {
						return nil, errors.Wrap(err, "time parse")
					}
					newDate := date.AddDate(estimatedDuration, 0, 0)
					dateOfEndOfAssessment = newDate.Format(config.ISO8601Format)
					amortizationValue = calculateMonthlyConsumption(dateOfAssessment, grossPrice, estimatedDuration)
					//amortizationValue = item.AssessmentPrice
					break
				}
			}
		}
	} else {
		if item.DateOfPurchase != nil {
			dateOfAssessment = *item.DateOfPurchase
		}
	}

	status := "Nezaduženo"

	if item.Type == "movable" && item.Active {
		itemInventoryList, _ := r.GetDispatchItemByInventoryID(item.ID)
		if len(itemInventoryList) > 0 {
			dispatchRes, err := r.GetDispatchItemByID(itemInventoryList[0].DispatchID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get dispatch item by id")
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
			return nil, errors.Wrap(err, "repo get my inventory real estate")
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
			return nil, errors.Wrap(err, "repo get organization unit by id")
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
			return nil, errors.Wrap(err, "repo get organization unit by id")
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
		GrossPrice:                   grossPrice - amortizationValue,
		Inactive:                     item.Inactive,
		PurchaseGrossPrice:           item.GrossPrice,
		DateOfPurchase:               item.DateOfPurchase,
		DateOfAssessments:            dateOfAssessment,
		DateOfEndOfAssessment:        dateOfEndOfAssessment,
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
		Description:                  item.Description,
	}

	return &res, nil
}

func buildInventoryItemResponse(r repository.MicroserviceRepositoryInterface, item *structs.BasicInventoryInsertItem, organizationUnitID int) (*dto.BasicInventoryResponseItem, error) {
	settingDropdownClassType := dto.DropdownSimple{}
	if item.ClassTypeID != 0 {
		settings, err := r.GetDropdownSettingByID(item.ClassTypeID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}

		if settings != nil {
			settingDropdownClassType = dto.DropdownSimple{ID: settings.ID, Title: settings.Title}
		}
	}

	suppliersDropdown := dto.DropdownSimple{}
	if item.SupplierID != 0 {
		suppliers, err := r.GetSupplier(item.SupplierID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get supplier")
		}

		if suppliers != nil {
			suppliersDropdown = dto.DropdownSimple{ID: suppliers.ID, Title: suppliers.Title}
		}
	}

	donorDropdown := dto.DropdownSimple{}
	if item.DonorID != 0 {
		donor, err := r.GetSupplier(item.DonorID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get supplier")
		}

		if donor != nil {
			donorDropdown = dto.DropdownSimple{ID: donor.ID, Title: donor.Title}
		}
	}

	settingDropdownOfficeID := dto.DropdownSimple{}
	if item.OfficeID != 0 {
		settings, err := r.GetDropdownSettingByID(item.OfficeID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}

		if settings != nil {
			settingDropdownOfficeID = dto.DropdownSimple{ID: settings.ID, Title: settings.Title}
		}
	}

	targetUserDropdown := dto.DropdownSimple{}
	if item.TargetUserProfileID != 0 {
		user, err := r.GetUserProfileByID(item.TargetUserProfileID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
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
			return nil, errors.Wrap(err, "repo get organization unit by id")
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
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}
		if targetOrganizationUnit != nil {
			targetOrganizationUnitDropdown = dto.DropdownSimple{ID: targetOrganizationUnit.ID, Title: targetOrganizationUnit.Title}
		}
	}

	realEstate, err := r.GetMyInventoryRealEstate(item.ID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get my inventory real estate")
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
				residualPrice = assessmentResponse.ResidualPrice
				lifetimeOfAssessmentInMonths = assessmentResponse.EstimatedDuration
				dateOfAssessment = *assessmentResponse.DateOfAssessment
				grossPrice = assessmentResponse.GrossPriceDifference
				amortizationValue = calculateMonthlyConsumption(dateOfAssessment, grossPrice, lifetimeOfAssessmentInMonths)
				grossPrice = assessmentResponse.GrossPriceDifference - amortizationValue

				//grossPrice = assessmentResponse.GrossPriceDifference - item.AssessmentPrice
				//amortizationValue = item.AssessmentPrice
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
				return nil, errors.Wrap(err, "repo get dispatch item by id")
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
			if dispatchRes.Type == "created" {
				continue
			}
			dispatch, _ := buildInventoryDispatchResponse(r, dispatchRes)
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

			if addMovement {
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
			return nil, errors.Wrap(err, "repo get file by id")
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
			return nil, errors.Wrap(err, "repo get file by id")
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

	if lifetimeOfAssessmentInMonths == 0 {
		lifetimeOfAssessmentInMonths = 999999999
	}

	res := dto.BasicInventoryResponseItem{
		ID:                           item.ID,
		InvoiceArticleID:             item.InvoiceArticleID,
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
		DepreciationRate:             fmt.Sprintf("%d", depreciationRate/lifetimeOfAssessmentInMonths),
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

/*
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
*/
