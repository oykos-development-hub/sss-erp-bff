package resolvers

import (
	"bff/config"
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) JudgesOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["user_profile_id"]
	page := params.Args["page"]
	size := params.Args["size"]

	orgUnitID := params.Args["organization_unit_id"]

	response := dto.Response{
		Status: "success",
	}

	filter := dto.JudgeResolutionsOrganizationUnitInput{}
	loggedInUser := params.Context.Value(config.LoggedInAccountKey).(*structs.UserAccounts)
	profileOrganizationUnit := params.Context.Value(config.OrganizationUnitIDKey).(*int)

	if page != nil {
		Page := page.(int)
		filter.Page = &Page
	}
	if size != nil {
		Size := size.(int)
		filter.PageSize = &Size
	}
	if orgUnitID != nil {
		OrgUnit := orgUnitID.(int)
		filter.OrganizationUnitID = &OrgUnit
	}
	if id != nil {
		user := id.(int)
		filter.UserProfileID = &user
	}
	if !loggedInUser.HasPermission(structs.PermissionManageOrganizationUnits) {
		filter.OrganizationUnitID = profileOrganizationUnit
	}

	active := true
	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}

	resolution, err := r.Repo.GetJudgeResolutionList(&input)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	if resolution.Data == nil {
		response.Message = "You must create active judge number resolution!"
		return response, nil
	}

	filter.ResolutionID = &resolution.Data[0].ID
	judges, total, err := r.Repo.GetJudgeResolutionOrganizationUnit(&filter)

	if err != nil {
		return errors.HandleAPIError(err)
	}

	var responseItems []*dto.Judges

	var normYear *int
	if params.Args["norm_year"] != nil {
		normYearRaw := params.Args["norm_year"].(int)
		normYear = &normYearRaw
	}

	for _, judge := range judges {
		judgeUser, err := buildJudgeResponseItem(r.Repo, judge.UserProfileID, judge.OrganizationUnitID, judge.IsPresident, normYear)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		responseItems = append(responseItems, judgeUser)
	}

	response.Items = responseItems
	response.Total = total
	return response, nil
}

func buildJudgeResponseItem(r repository.MicroserviceRepositoryInterface, userProfileID, organizationUnitID int, isPresident bool, normYear *int) (*dto.Judges, error) {
	userProfile, err := r.GetUserProfileByID(userProfileID)
	if err != nil {
		return nil, err
	}
	userAccount, err := r.GetUserAccountByID(userProfile.UserAccountID)
	if err != nil {
		return nil, err
	}

	organizationUnit, err := r.GetOrganizationUnitByID(organizationUnitID)
	if err != nil {
		return nil, err
	}
	organizationUnitDropdown := structs.SettingsDropdown{
		ID:    organizationUnit.ID,
		Title: organizationUnit.Title,
	}

	norms, err := r.GetJudgeNormListByEmployee(userProfile.ID, dto.GetUserNormFulfilmentListInput{NormYear: normYear})
	if err != nil {
		return nil, err
	}

	normResItemList, err := buildNormResItemList(r, norms)
	if err != nil {
		return nil, err
	}

	return &dto.Judges{
		ID:               userProfile.ID,
		FirstName:        userProfile.FirstName,
		LastName:         userProfile.LastName,
		IsJudgePresident: isPresident,
		OrganizationUnit: organizationUnitDropdown,
		Norms:            normResItemList,
		Gender:           userProfile.Gender,
		Age:              userProfile.GetAge(),
		FolderID:         userAccount.FolderID,
		CreatedAt:        userProfile.CreatedAt,
		UpdatedAt:        userProfile.UpdatedAt,
	}, nil
}

func buildNormResItemList(r repository.MicroserviceRepositoryInterface, norms []structs.JudgeNorms) ([]*dto.NormResItem, error) {
	var normResItems []*dto.NormResItem

	for _, norm := range norms {
		normResItem, err := buildNormResItem(r, norm)
		if err != nil {
			return nil, err
		}

		normResItems = append(normResItems, normResItem)
	}

	return normResItems, nil
}

func buildNormResItem(r repository.MicroserviceRepositoryInterface, norm structs.JudgeNorms) (*dto.NormResItem, error) {
	normResItem := &dto.NormResItem{
		ID:                       norm.ID,
		UserProfileID:            norm.UserProfileID,
		Topic:                    norm.Topic,
		Title:                    norm.Title,
		PercentageOfNormDecrease: norm.PercentageOfNormDecrease,
		NumberOfNormDecrease:     norm.NumberOfNormDecrease,
		NumberOfItems:            norm.NumberOfItems,
		NumberOfItemsSolved:      norm.NumberOfItemsSolved,
		DateOfEvaluationValidity: norm.DateOfEvaluationValidity,
		NormStartDate:            norm.NormStartDate,
		NormEndDate:              norm.NormEndDate,
		FileID:                   norm.FileID,
		CreatedAt:                norm.CreatedAt,
		UpdatedAt:                norm.UpdatedAt,
	}
	if norm.EvaluationID != nil {
		evaluation, err := r.GetEvaluation(*norm.EvaluationID)
		if err != nil {
			return nil, err
		}

		evaluationType, err := r.GetDropdownSettingByID(evaluation.EvaluationTypeID)
		if err != nil {
			return nil, err
		}
		normResItem.Evaluation = evaluation
		evaluation.EvaluationType = *evaluationType
	}
	if norm.RelocationID != nil {
		relocation, err := r.GetAbsentByID(*norm.RelocationID)
		if err != nil {
			return nil, err
		}

		relocationResItem, err := buildAbsentResponseItem(r, *relocation)
		if err != nil {
			return nil, err
		}

		normResItem.Relocation = relocationResItem
	}

	return normResItem, nil
}

func (r *Resolver) JudgeNormInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JudgeNorms
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		res, err := r.Repo.UpdateJudgeNorm(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = res
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateJudgeNorm(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = res
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) JudgeNormDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteJudgeNorm(itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) JudgeResolutionsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	page := params.Args["page"].(int)
	size := params.Args["size"].(int)

	resolutionList := []*structs.JudgeResolutions{}
	response := dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}

	if id != nil && id.(int) > 0 {
		resolution, err := r.Repo.GetJudgeResolution(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}

		resolutionList = append(resolutionList, resolution)
	} else {
		input := dto.GetJudgeResolutionListInputMS{}
		input.Page = &page
		input.Size = &size
		resolutions, err := r.Repo.GetJudgeResolutionList(&input)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resolutionList = append(resolutionList, resolutions.Data...)
	}

	resolutionResponseList, err := processResolutions(r.Repo, resolutionList)
	if err != nil {
		return dto.ErrorResponse(err), nil
	}

	response.Total = len(resolutionResponseList)
	paginatedItems, _ := shared.Paginate(resolutionResponseList, page, size)

	response.Items = paginatedItems

	return response, nil
}

func (r *Resolver) JudgeResolutionsActiveResolver(params graphql.ResolveParams) (interface{}, error) {

	response := dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}
	var item dto.JudgeResolutionsResponseItem
	page := 0
	size := 1000
	input := dto.GetJudgeResolutionListInputMS{}
	input.Page = &page
	input.Size = &size
	resolutions, err := r.Repo.GetJudgeResolutionList(&input)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	for _, res := range resolutions.Data {
		if res.Active {
			resolutionResponseItem, err := processJudgeResolution(r.Repo, res)
			if err != nil {
				fmt.Printf("Error processing JudgeResolution: %v\n", err)
				return nil, err
			}
			item = *resolutionResponseItem
			break
		}
	}
	response.Item = item
	return response, nil
}

func (r *Resolver) CheckJudgeAndPresidentIsAvailable(params graphql.ResolveParams) (interface{}, error) {
	active := true
	input := dto.GetJudgeResolutionListInputMS{
		Active: &active,
	}

	response := dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}

	check := dto.CheckJudgeAndPresidentIsAvailableMS{
		Judge:     false,
		President: false,
	}

	organizationUnitID := params.Args["organization_unit_id"]
	if organizationUnitID != nil && organizationUnitID.(int) > 0 {

		resolution, _ := r.Repo.GetJudgeResolutionList(&input)

		if len(resolution.Data) > 0 {
			orgUnitID := organizationUnitID.(int)
			resolutionID := resolution.Data[0].ID

			judgeResolutionOrganizationUnit, _, err := r.Repo.GetJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
				OrganizationUnitID: &orgUnitID,
				ResolutionID:       &resolutionID,
			})

			if err != nil {
				return nil, err
			}

			numberOfJudges := len(judgeResolutionOrganizationUnit)

			itemsInput := dto.GetJudgeResolutionItemListInputMS{
				ResolutionID: &resolution.Data[0].ID,
			}
			resolutionItems, err := r.Repo.GetJudgeResolutionItemsList(&itemsInput)
			if err != nil {
				return nil, err
			}

			if len(resolutionItems) > 0 {
				for _, item := range resolutionItems {
					if item.OrganizationUnitID == organizationUnitID {
						if numberOfJudges < item.NumberOfJudges {
							check.Judge = true
						}
					}
				}
			} else {
				check.Judge = true
			}
			check.President = true
			for _, judge := range judgeResolutionOrganizationUnit {
				if judge.IsPresident {
					check.President = false
				}
			}
		}
	}
	response.Item = check
	return response, nil
}

func processResolutions(r repository.MicroserviceRepositoryInterface, resolutionList []*structs.JudgeResolutions) ([]*dto.JudgeResolutionsResponseItem, error) {
	var resolutionResponseList []*dto.JudgeResolutionsResponseItem

	// Process JudgeResolutions concurrently
	//var wg sync.WaitGroup
	//wg.Add(len(resolutionList))

	for _, resolution := range resolutionList {
		//		go func(resolution *structs.JudgeResolutions) {
		//		defer wg.Done()
		resolutionResponseItem, err := processJudgeResolution(r, resolution)
		if err != nil {
			fmt.Printf("Error processing JudgeResolution: %v\n", err)
			return nil, err
		}
		resolutionResponseList = append(resolutionResponseList, resolutionResponseItem)
		//		}(resolution)
	}

	//	wg.Wait()

	return resolutionResponseList, nil
}

func processJudgeResolution(r repository.MicroserviceRepositoryInterface, resolution *structs.JudgeResolutions) (*dto.JudgeResolutionsResponseItem, error) {
	itemsInput := dto.GetJudgeResolutionItemListInputMS{
		ResolutionID: &resolution.ID,
	}
	resolutionItems, err := r.GetJudgeResolutionItemsList(&itemsInput)
	if err != nil {
		return nil, err
	}

	resolutionItemResponseItemList, totalNumberOfAvailableSlotsJudges, totalNumberOfJudges, err := processResolutionItems(r, resolutionItems)
	if err != nil {
		return nil, err
	}

	resolutionResponseItem := &dto.JudgeResolutionsResponseItem{
		ID:                   resolution.ID,
		SerialNumber:         resolution.SerialNumber,
		CreatedAt:            resolution.CreatedAt,
		UpdatedAt:            resolution.UpdatedAt,
		Active:               resolution.Active,
		Items:                resolutionItemResponseItemList,
		AvailableSlotsJudges: totalNumberOfAvailableSlotsJudges,
		NumberOfJudges:       totalNumberOfJudges,
	}

	return resolutionResponseItem, nil
}

func processResolutionItems(r repository.MicroserviceRepositoryInterface, resolutionItems []*structs.JudgeResolutionItems) ([]*dto.JudgeResolutionItemResponseItem, int, int, error) {
	var resolutionItemResponseItemList []*dto.JudgeResolutionItemResponseItem
	var totalNumberOfAvailableSlotsJudges, totalNumberOfJudges int

	// Process JudgeResolutionItems concurrently
	var wg sync.WaitGroup
	wg.Add(len(resolutionItems))

	for _, resolutionItem := range resolutionItems {
		go func(resolutionItem *structs.JudgeResolutionItems) {
			defer wg.Done()
			resolutionItemResponseItem, err := buildResolutionItemResponseItem(r, resolutionItem)
			if err != nil {
				fmt.Printf("Error building ResolutionItemResponseItem: %v\n", err)
				return
			}

			resolutionItemResponseItemList = append(resolutionItemResponseItemList, resolutionItemResponseItem)

			totalNumberOfAvailableSlotsJudges += resolutionItemResponseItem.AvailableSlotsJudges
			totalNumberOfJudges += resolutionItemResponseItem.AvailableSlotsPredisents + resolutionItemResponseItem.AvailableSlotsJudges
		}(resolutionItem)
	}

	wg.Wait()

	return resolutionItemResponseItemList, totalNumberOfAvailableSlotsJudges, totalNumberOfJudges, nil
}

func buildResolutionItemResponseItem(r repository.MicroserviceRepositoryInterface, item *structs.JudgeResolutionItems) (*dto.JudgeResolutionItemResponseItem, error) {
	organizationUnit, err := r.GetOrganizationUnitByID(item.OrganizationUnitID)
	if err != nil {
		return nil, err
	}
	organizationUnitDropdown := structs.SettingsDropdown{ID: organizationUnit.ID, Title: organizationUnit.Title}

	numberOfJudgesInOU, numberOfPresidents, numberOfEmployees, numberOfRelocations, err := calculateEmployeeStats(r, organizationUnit.ID, item.ID)
	if err != nil {
		fmt.Printf("Calculating number of presindents failed beacuse of error: %v\n", err)
	}

	return &dto.JudgeResolutionItemResponseItem{
		ID:                       item.ID,
		ResolutionID:             item.ResolutionID,
		OrganizationUnit:         organizationUnitDropdown,
		AvailableSlotsJudges:     item.NumberOfJudges,
		AvailableSlotsPredisents: item.NumberOfPresidents,
		NumberOfJudges:           numberOfJudgesInOU,
		NumberOfPresidents:       numberOfPresidents,
		NumberOfEmployees:        numberOfEmployees,
		NumberOfSuspendedJudges:  0,
		NumberOfRelocatedJudges:  numberOfRelocations,
	}, nil
}

func (r *Resolver) OrganizationUintCalculateEmployeeStats(params graphql.ResolveParams) (interface{}, error) {
	var response []dto.JudgeResolutionItemResponseItem
	page := 1
	size := 1000
	resID := params.Args["resolution_id"]
	active := params.Args["active"]
	input := dto.GetOrganizationUnitsInput{
		Page: &page,
		Size: &size,
	}
	organizationUnits, err := r.Repo.GetOrganizationUnits(&input)
	if err != nil {
		return nil, err
	}

	var resolutionID int

	if active != nil && active.(bool) {
		isActive := true
		filter := dto.GetJudgeResolutionListInputMS{
			Active: &isActive,
		}
		resolution, err := r.Repo.GetJudgeResolutionList(&filter)

		if err != nil {
			return nil, err
		}

		if resolution.Data == nil {
			return dto.Response{
				Status:  "success",
				Message: "Here's the list you asked for!",
				Items:   dto.JudgeResolutionItemResponseItem{},
			}, nil
		}

		resolutionID = resolution.Data[0].ID
	} else if resID.(int) != 0 {
		resolutionID = resID.(int)
	} else {
		return dto.Response{
			Status:  "failed",
			Message: "You must provide one argument!",
			Items:   dto.JudgeResolutionItemResponseItem{},
		}, nil
	}

	for _, organizationUnit := range organizationUnits.Data {

		if organizationUnit.ParentID != nil {
			continue
		}
		organizationUnitDropdown := structs.SettingsDropdown{ID: organizationUnit.ID, Title: organizationUnit.Title}

		numberOfJudgesInOU, numberOfPresidents, numberOfEmployees, numberOfRelocations, err := calculateEmployeeStats(r.Repo, organizationUnit.ID, resolutionID)
		if err != nil {
			fmt.Printf("Calculating number of presindents failed beacuse of error: %v\n", err)
		}

		filter := dto.GetJudgeResolutionItemListInputMS{
			ResolutionID:       &resolutionID,
			OrganizationUnitID: &organizationUnit.ID,
		}

		judgeResolutionItem, err := r.Repo.GetJudgeResolutionItemsList(&filter)
		if err != nil {
			fmt.Printf("Calculating number of slots failed beacuse of error: %v\n", err)
		}

		var numberOfJudgesSlots int
		var numberOfPresidentsSlots int
		if len(judgeResolutionItem) > 0 {
			numberOfJudgesSlots = judgeResolutionItem[0].NumberOfJudges
			numberOfPresidentsSlots = judgeResolutionItem[0].NumberOfPresidents
		}
		response = append(response, dto.JudgeResolutionItemResponseItem{
			OrganizationUnit:         organizationUnitDropdown,
			NumberOfJudges:           numberOfJudgesInOU,
			NumberOfPresidents:       numberOfPresidents,
			TotalNumber:              numberOfJudgesInOU + numberOfPresidents,
			NumberOfEmployees:        numberOfEmployees,
			NumberOfSuspendedJudges:  0,
			NumberOfRelocatedJudges:  numberOfRelocations,
			AvailableSlotsPredisents: numberOfPresidentsSlots,
			AvailableSlotsJudges:     numberOfJudgesSlots,
			AvailableSlotsTotal:      numberOfPresidentsSlots + numberOfJudgesSlots,
			VacantSlotsJudges:        numberOfJudgesSlots - numberOfJudgesInOU,
			VacantSlotsPresidents:    numberOfPresidentsSlots - numberOfPresidents,
			VacantSlots:              numberOfJudgesSlots - numberOfJudgesInOU + numberOfPresidentsSlots - numberOfPresidents,
		})
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

func calculateEmployeeStats(r repository.MicroserviceRepositoryInterface, id int, resID int) (int, int, int, int, error) {
	var numberOfEmployees, numberOfJudges, totalRelocations, numberOfJudgePresidents int

	input := &dto.JudgeResolutionsOrganizationUnitInput{
		OrganizationUnitID: &id,
		ResolutionID:       &resID,
	}
	judgeResolutionOrganizationUnit, _, err := r.GetJudgeResolutionOrganizationUnit(input)

	if err != nil {
		return numberOfEmployees, numberOfJudges, totalRelocations, numberOfJudgePresidents, err
	}

	numberOfJudges = len(judgeResolutionOrganizationUnit)
	totalRelocations = numberOfJudges

	for _, judge := range judgeResolutionOrganizationUnit {
		if judge.IsPresident {
			numberOfJudgePresidents = 1
			break
		}
	}

	return numberOfJudges, numberOfJudgePresidents, numberOfEmployees, totalRelocations, nil
}

// func getNumberOfRelocatedJudges(employees []*structs.EmployeesInOrganizationUnits) (int, error) {
// 	var numberOfRelocations int
// 	for _, employee := range employees {
// 		today := time.Now()
// 		absents, err := r.Repo.GetEmployeeAbsents(employee.UserProfileID, &dto.EmployeeAbsentsInput{Date: &today})
// 		if err != nil {
// 			return 0, err
// 		}
// 		for _, absent := range absents {
// 			absentType, err := r.Repo.GetAbsentTypeByID(absent.AbsentTypeID)
// 			if err != nil {
// 				return 0, err
// 			}
// 			if absentType.Relocation {
// 				numberOfRelocations++
// 			}
// 		}
// 	}
// 	return numberOfRelocations, nil
// }

func (r *Resolver) JudgeResolutionInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JudgeResolutions
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	var (
		resolution *structs.JudgeResolutions
		err        error
	)

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		judgeResolution := structs.JudgeResolutions{
			SerialNumber: data.SerialNumber,
			Active:       data.Active,
		}
		resolution, err = r.Repo.UpdateJudgeResolutions(itemID, &judgeResolution)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		updatedItems, err := insertOrUpdateResolutionItemList(r.Repo, data.Items, resolution.ID)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resolution.Items = updatedItems
		response.Message = "You updated this item!"

	} else {
		judgeResolution := structs.JudgeResolutions{
			SerialNumber: data.SerialNumber,
			Active:       true,
		}
		resolution, err = r.Repo.CreateJudgeResolutions(&judgeResolution)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		/*input := dto.GetJudgeResolutionListInputMS{}

		resolutions, err := r.Repo.GetJudgeResolutionList(&input)
		if err != nil {
			return errors.HandleAPIError(err)
		}

			for _, res := range resolutions.Data {
			//if res.Active && resolution.ID != res.ID {
				judgeResolution := structs.JudgeResolutions{
					SerialNumber: res.SerialNumber,
					Active:       false,
				}
				_, err = r.Repo.UpdateJudgeResolutions(res.ID, &judgeResolution)
				if err != nil {
					return errors.HandleAPIError(err)
				}
		*/
		oldResID := resolution.ID - 1
		judgesResolution, _, err := r.Repo.GetJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
			ResolutionID: &oldResID,
		})
		if err != nil {
			return errors.HandleAPIError(err)
		}

		if len(judgesResolution) > 0 {
			for _, judge := range judgesResolution {
				inputCreate := dto.JudgeResolutionsOrganizationUnitItem{
					UserProfileID:      judge.UserProfileID,
					OrganizationUnitID: judge.OrganizationUnitID,
					ResolutionID:       resolution.ID,
					IsPresident:        judge.IsPresident,
				}
				_, err := r.Repo.CreateJudgeResolutionOrganizationUnit(&inputCreate)
				if err != nil {
					return errors.HandleAPIError(err)
				}
				//		}
				//			}

			}
		}

		updatedItems, err := insertOrUpdateResolutionItemList(r.Repo, data.Items, resolution.ID)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		resolution.Items = updatedItems
		response.Message = "You created this item!"
	}

	judgeResolution, err := processJudgeResolution(r.Repo, resolution)
	if err != nil {
		return dto.ErrorResponse(err), nil
	}

	response.Item = judgeResolution

	return response, nil
}

func insertOrUpdateResolutionItemList(r repository.MicroserviceRepositoryInterface, items []*structs.JudgeResolutionItems, resolutionID int) ([]*structs.JudgeResolutionItems, error) {
	var updateItemsList []*structs.JudgeResolutionItems
	for _, item := range items {
		judgeResolutionItem := structs.JudgeResolutionItems{
			ResolutionID:       resolutionID,
			OrganizationUnitID: item.OrganizationUnitID,
			NumberOfJudges:     item.NumberOfJudges,
			NumberOfPresidents: item.NumberOfPresidents,
		}
		if item.ID > 0 {
			item, err := r.UpdateJudgeResolutionItems(item.ID, &judgeResolutionItem)
			if err != nil {
				fmt.Printf("Updating Judge Resolution Items failed because of this error - %s.\n", err)
				return nil, err
			}
			updateItemsList = append(updateItemsList, item)
		} else {
			item, err := r.CreateJudgeResolutionItems(&judgeResolutionItem)
			if err != nil {
				fmt.Printf("Creating Judge Resolution Items failed because of this error - %s.\n", err)
				return nil, err
			}
			updateItemsList = append(updateItemsList, item)
		}
	}
	return updateItemsList, nil
}

func (r *Resolver) JudgeResolutionDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteJudgeResolution(itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
