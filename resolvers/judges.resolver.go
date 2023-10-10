package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/graphql-go/graphql"
)

var JudgesOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["user_profile_id"]
	page := params.Args["page"]
	size := params.Args["size"]

	orgUnitID := params.Args["organization_unit_id"]

	response := dto.Response{
		Status: "success",
	}

	filter := dto.JudgeResolutionsOrganizationUnitInput{}

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
		filter.OrganizationUnitId = &OrgUnit
	}
	if id != nil {
		user := id.(int)
		filter.UserProfileId = &user
	}

	active := true
	input := dto.GetJudgeResolutionItemListInputMS{
		Active: &active,
	}

	resolution, err := getJudgeResolutionItemsList(&input)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	filter.ResolutionId = &resolution[0].ResolutionId

	judges, total, err := getJudgeResolutionOrganizationUnit(&filter)

	if err != nil {
		return shared.HandleAPIError(err)
	}

	var responseItems []*dto.Judges

	for _, judge := range judges {
		judgeUser, err := buildJudgeResponseItem(judge.UserProfileId, judge.OrganizationUnitId, judge.IsPresident)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		responseItems = append(responseItems, judgeUser)
	}

	response.Items = responseItems
	response.Total = total
	return response, nil
}

func buildJudgeResponseItem(userProfileID, organizationUnitID int, isPresident bool) (*dto.Judges, error) {
	userProfile, err := getUserProfileById(userProfileID)
	if err != nil {
		return nil, err
	}
	userAccount, err := GetUserAccountById(userProfile.UserAccountId)
	if err != nil {
		return nil, err
	}

	organizationUnit, err := getOrganizationUnitById(organizationUnitID)
	if err != nil {
		return nil, err
	}
	organizationUnitDropdown := structs.SettingsDropdown{
		Id:    organizationUnit.Id,
		Title: organizationUnit.Title,
	}

	norms, err := getJudgeNormListByEmployee(userProfile.Id)
	if err != nil {
		return nil, err
	}

	normResItemList, err := buildNormResItemList(norms)
	if err != nil {
		return nil, err
	}

	return &dto.Judges{
		ID:               userProfile.Id,
		FirstName:        userProfile.FirstName,
		LastName:         userProfile.LastName,
		IsJudgePresident: isPresident,
		OrganizationUnit: organizationUnitDropdown,
		Norms:            normResItemList,
		FolderID:         userAccount.FolderId,
		CreatedAt:        userProfile.CreatedAt,
		UpdatedAt:        userProfile.UpdatedAt,
	}, nil
}

func buildNormResItemList(norms []structs.JudgeNorms) ([]*dto.NormResItem, error) {
	var normResItems []*dto.NormResItem

	for _, norm := range norms {
		normResItem, err := buildNormResItem(norm)
		if err != nil {
			return nil, err
		}

		normResItems = append(normResItems, normResItem)
	}

	return normResItems, nil
}

func buildNormResItem(norm structs.JudgeNorms) (*dto.NormResItem, error) {
	normResItem := &dto.NormResItem{
		Id:                       norm.Id,
		UserProfileId:            norm.UserProfileId,
		Topic:                    norm.Topic,
		Title:                    norm.Title,
		PercentageOfNormDecrease: norm.PercentageOfNormDecrease,
		NumberOfNormDecrease:     norm.NumberOfNormDecrease,
		NumberOfItems:            norm.NumberOfItems,
		NumberOfItemsSolved:      norm.NumberOfItemsSolved,
		DateOfEvaluationValidity: norm.DateOfEvaluationValidity,
		FileID:                   norm.FileID,
		CreatedAt:                norm.CreatedAt,
		UpdatedAt:                norm.UpdatedAt,
	}
	if norm.EvaluationID != nil {
		evaluation, err := getEvaluation(*norm.EvaluationID)
		if err != nil {
			return nil, err
		}

		evaluationType, err := getDropdownSettingById(evaluation.EvaluationTypeId)
		if err != nil {
			return nil, err
		}
		normResItem.Evaluation = evaluation
		evaluation.EvaluationType = *evaluationType
	}
	if norm.RelocationID != nil {
		relocation, err := getAbsentById(*norm.RelocationID)
		if err != nil {
			return nil, err
		}

		relocationResItem, err := buildAbsentResponseItem(*relocation)
		if err != nil {
			return nil, err
		}

		normResItem.Relocation = relocationResItem
	}

	return normResItem, nil
}

var JudgeNormInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JudgeNorms
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateJudgeNorm(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = res
		response.Message = "You updated this item!"
	} else {
		res, err := createJudgeNorm(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = res
		response.Message = "You created this item!"
	}

	return response, nil
}

var JudgeNormDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteJudgeNorm(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

var JudgeResolutionsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	page := params.Args["page"].(int)
	size := params.Args["size"].(int)

	resolutionList := []*structs.JudgeResolutions{}
	response := dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}

	if id != nil && id.(int) > 0 {
		resolution, err := getJudgeResolution(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}

		resolutionList = append(resolutionList, resolution)
	} else {
		input := dto.GetJudgeResolutionListInputMS{}
		input.Page = &page
		input.Size = &size
		resolutions, err := getJudgeResolutionList(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resolutionList = append(resolutionList, resolutions.Data...)
	}

	resolutionResponseList, err := processResolutions(resolutionList, page, size)
	if err != nil {
		return dto.ErrorResponse(err), nil
	}

	response.Total = len(resolutionResponseList)
	paginatedItems, _ := shared.Paginate(resolutionResponseList, page, size)

	response.Items = paginatedItems

	return response, nil
}

func CheckJudgeAndPresidentIsAvailable(params graphql.ResolveParams) (interface{}, error) {
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

	organizationUnitId := params.Args["organization_unit_id"]
	if shared.IsInteger(organizationUnitId) && organizationUnitId.(int) > 0 {

		resolution, _ := getJudgeResolutionList(&input)

		if len(resolution.Data) > 0 {
			orgUnitID := organizationUnitId.(int)
			resolutionID := resolution.Data[0].Id

			judgeResolutionOrganizationUnit, _, err := getJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
				OrganizationUnitId: &orgUnitID,
				ResolutionId:       &resolutionID,
			})

			if err != nil {
				return nil, err
			}

			numberOfJudges := len(judgeResolutionOrganizationUnit)

			itemsInput := dto.GetJudgeResolutionItemListInputMS{
				ResolutionID: &resolution.Data[0].Id,
			}
			resolutionItems, err := getJudgeResolutionItemsList(&itemsInput)
			if err != nil {
				return nil, err
			}

			if len(resolutionItems) > 0 {
				for _, item := range resolutionItems {
					if item.OrganizationUnitId == organizationUnitId {
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

func processResolutions(resolutionList []*structs.JudgeResolutions, page, size int) ([]*dto.JudgeResolutionsResponseItem, error) {
	var resolutionResponseList []*dto.JudgeResolutionsResponseItem

	// Process JudgeResolutions concurrently
	//var wg sync.WaitGroup
	//wg.Add(len(resolutionList))

	for _, resolution := range resolutionList {
		//		go func(resolution *structs.JudgeResolutions) {
		//		defer wg.Done()
		resolutionResponseItem, err := processJudgeResolution(resolution)
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

func processJudgeResolution(resolution *structs.JudgeResolutions) (*dto.JudgeResolutionsResponseItem, error) {
	itemsInput := dto.GetJudgeResolutionItemListInputMS{
		ResolutionID: &resolution.Id,
	}
	resolutionItems, err := getJudgeResolutionItemsList(&itemsInput)
	if err != nil {
		return nil, err
	}

	resolutionItemResponseItemList, totalNumberOfAvailableSlotsJudges, totalNumberOfJudges, err := processResolutionItems(resolutionItems)
	if err != nil {
		return nil, err
	}

	resolutionResponseItem := &dto.JudgeResolutionsResponseItem{
		Id:                   resolution.Id,
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

func processResolutionItems(resolutionItems []*structs.JudgeResolutionItems) ([]*dto.JudgeResolutionItemResponseItem, int, int, error) {
	var resolutionItemResponseItemList []*dto.JudgeResolutionItemResponseItem
	var totalNumberOfAvailableSlotsJudges, totalNumberOfJudges int

	// Process JudgeResolutionItems concurrently
	var wg sync.WaitGroup
	wg.Add(len(resolutionItems))

	for _, resolutionItem := range resolutionItems {
		go func(resolutionItem *structs.JudgeResolutionItems) {
			defer wg.Done()
			resolutionItemResponseItem, err := buildResolutionItemResponseItem(resolutionItem)
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

func buildResolutionItemResponseItem(item *structs.JudgeResolutionItems) (*dto.JudgeResolutionItemResponseItem, error) {
	organizationUnit, err := getOrganizationUnitById(item.OrganizationUnitId)
	if err != nil {
		return nil, err
	}
	organizationUnitDropdown := structs.SettingsDropdown{Id: organizationUnit.Id, Title: organizationUnit.Title}

	numberOfJudgesInOU, numberOfPresidents, numberOfEmployees, numberOfRelocations, err := calculateEmployeeStats(organizationUnit.Id, item.Id)
	if err != nil {
		fmt.Printf("Calculating number of presindents failed beacuse of error: %v\n", err)
	}

	return &dto.JudgeResolutionItemResponseItem{
		Id:                       item.Id,
		ResolutionId:             item.ResolutionId,
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

func OrganizationUintCalculateEmployeeStats(params graphql.ResolveParams) (interface{}, error) {
	var response []dto.JudgeResolutionItemResponseItem
	var page int = 1
	var size int = 1000
	resID := params.Args["resolution_id"]
	input := dto.GetOrganizationUnitsInput{
		Page: &page,
		Size: &size,
	}
	organizationUnits, err := getOrganizationUnits(&input)
	if err != nil {
		return nil, err
	}

	for _, organizationUnit := range organizationUnits.Data {

		if organizationUnit.ParentId != nil {
			continue
		}
		organizationUnitDropdown := structs.SettingsDropdown{Id: organizationUnit.Id, Title: organizationUnit.Title}

		numberOfJudgesInOU, numberOfPresidents, numberOfEmployees, numberOfRelocations, err := calculateEmployeeStats(organizationUnit.Id, resID.(int))
		if err != nil {
			fmt.Printf("Calculating number of presindents failed beacuse of error: %v\n", err)
		}

		response = append(response, dto.JudgeResolutionItemResponseItem{
			OrganizationUnit:        organizationUnitDropdown,
			NumberOfJudges:          numberOfJudgesInOU,
			NumberOfPresidents:      numberOfPresidents,
			NumberOfEmployees:       numberOfEmployees,
			NumberOfSuspendedJudges: 0,
			NumberOfRelocatedJudges: numberOfRelocations,
		})
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   response,
	}, nil
}

func calculateEmployeeStats(id int, resID int) (int, int, int, int, error) {
	var numberOfEmployees, numberOfJudges, totalRelocations, numberOfJudgePresidents int

	input := &dto.JudgeResolutionsOrganizationUnitInput{
		OrganizationUnitId: &id,
		ResolutionId:       &resID,
	}
	judgeResolutionOrganizationUnit, _, err := getJudgeResolutionOrganizationUnit(input)

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
// 		absents, err := getEmployeeAbsents(employee.UserProfileId, &dto.EmployeeAbsentsInput{Date: &today})
// 		if err != nil {
// 			return 0, err
// 		}
// 		for _, absent := range absents {
// 			absentType, err := getAbsentTypeById(absent.AbsentTypeId)
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

var JudgeResolutionInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
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

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		judgeResolution := structs.JudgeResolutions{
			SerialNumber: data.SerialNumber,
			Active:       data.Active,
		}
		resolution, err = updateJudgeResolutions(itemId, &judgeResolution)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		updatedItems, err := insertOrUpdateResolutionItemList(data.Items, resolution.Id)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resolution.Items = updatedItems
		response.Message = "You updated this item!"

	} else {
		judgeResolution := structs.JudgeResolutions{
			SerialNumber: data.SerialNumber,
			Active:       true,
		}
		resolution, err = createJudgeResolutions(&judgeResolution)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		/*input := dto.GetJudgeResolutionListInputMS{}

		resolutions, err := getJudgeResolutionList(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}

			for _, res := range resolutions.Data {
			//if res.Active && resolution.Id != res.Id {
				judgeResolution := structs.JudgeResolutions{
					SerialNumber: res.SerialNumber,
					Active:       false,
				}
				_, err = updateJudgeResolutions(res.Id, &judgeResolution)
				if err != nil {
					return shared.HandleAPIError(err)
				}
		*/
		oldResID := resolution.Id - 1
		judgesResolution, _, err := getJudgeResolutionOrganizationUnit(&dto.JudgeResolutionsOrganizationUnitInput{
			ResolutionId: &oldResID,
		})
		if err != nil {
			return shared.HandleAPIError(err)
		}

		if len(judgesResolution) > 0 {
			for _, judge := range judgesResolution {
				inputCreate := dto.JudgeResolutionsOrganizationUnitItem{
					UserProfileId:      judge.UserProfileId,
					OrganizationUnitId: judge.OrganizationUnitId,
					ResolutionId:       resolution.Id,
					IsPresident:        judge.IsPresident,
				}
				_, err := createJudgeResolutionOrganizationUnit(&inputCreate)
				if err != nil {
					return shared.HandleAPIError(err)
				}
				//		}
				//			}

			}
		}

		updatedItems, err := insertOrUpdateResolutionItemList(data.Items, resolution.Id)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		resolution.Items = updatedItems
		response.Message = "You created this item!"
	}

	judgeResolution, err := processJudgeResolution(resolution)
	if err != nil {
		return dto.ErrorResponse(err), nil
	}

	response.Item = judgeResolution

	return response, nil
}

func insertOrUpdateResolutionItemList(items []*structs.JudgeResolutionItems, resolutionID int) ([]*structs.JudgeResolutionItems, error) {
	var updateItemsList []*structs.JudgeResolutionItems
	for _, item := range items {
		judgeResolutionItem := structs.JudgeResolutionItems{
			ResolutionId:       resolutionID,
			OrganizationUnitId: item.OrganizationUnitId,
			NumberOfJudges:     item.NumberOfJudges,
			NumberOfPresidents: item.NumberOfPresidents,
		}
		if item.Id > 0 {
			item, err := updateJudgeResolutionItems(item.Id, &judgeResolutionItem)
			if err != nil {
				fmt.Printf("Updating Judge Resolution Items failed because of this error - %s.\n", err)
				return nil, err
			}
			updateItemsList = append(updateItemsList, item)
		} else {
			item, err := createJudgeResolutionItems(&judgeResolutionItem)
			if err != nil {
				fmt.Printf("Creating Judge Resolution Items failed because of this error - %s.\n", err)
				return nil, err
			}
			updateItemsList = append(updateItemsList, item)
		}
	}
	return updateItemsList, nil
}

var JudgeResolutionDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteJudgeResolution(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func updateJudgeNorm(id int, norm *structs.JudgeNorms) (*structs.JudgeNorms, error) {
	res := &dto.GetJudgeNormResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.JUDGE_NORM_ENDPOINT+"/"+strconv.Itoa(id), norm, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createJudgeNorm(norm *structs.JudgeNorms) (*structs.JudgeNorms, error) {
	res := &dto.GetJudgeNormResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.JUDGE_NORM_ENDPOINT, norm, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteJudgeNorm(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.JUDGE_NORM_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func updateJudgeResolutionItems(id int, item *structs.JudgeResolutionItems) (*structs.JudgeResolutionItems, error) {
	res := &dto.GetJudgeResolutionItemResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.JUDGE_RESOLUTION_ITEMS_ENDPOINT+"/"+strconv.Itoa(id), item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createJudgeResolutionItems(item *structs.JudgeResolutionItems) (*structs.JudgeResolutionItems, error) {
	res := &dto.GetJudgeResolutionItemResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.JUDGE_RESOLUTION_ITEMS_ENDPOINT, item, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getJudgeResolutionItemsList(input *dto.GetJudgeResolutionItemListInputMS) ([]*structs.JudgeResolutionItems, error) {
	res := &dto.GetJudgeResolutionItemListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JUDGE_RESOLUTION_ITEMS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func updateJudgeResolutions(id int, resolution *structs.JudgeResolutions) (*structs.JudgeResolutions, error) {
	res := &dto.GetJudgeResolutionResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.JUDGE_RESOLUTIONS_ENDPOINT+"/"+strconv.Itoa(id), resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func createJudgeResolutions(resolution *structs.JudgeResolutions) (*structs.JudgeResolutions, error) {
	res := &dto.GetJudgeResolutionResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.JUDGE_RESOLUTIONS_ENDPOINT, resolution, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteJudgeResolution(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.JUDGE_RESOLUTIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getJudgeResolutionList(input *dto.GetJudgeResolutionListInputMS) (*dto.GetJudgeResolutionListResponseMS, error) {
	res := &dto.GetJudgeResolutionListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JUDGE_RESOLUTIONS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getJudgeResolution(id int) (*structs.JudgeResolutions, error) {
	res := &dto.GetJudgeResolutionResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JUDGE_RESOLUTIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getJudgeNormListByEmployee(userProfileID int) ([]structs.JudgeNorms, error) {
	res := &dto.GetEmployeeNormListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT+"/"+strconv.Itoa(userProfileID)+"/norms", nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func createJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitItem) (*dto.JudgeResolutionsOrganizationUnitItem, error) {
	res := &dto.GetJudgeResolutionsOrganizationUnitResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.JUDGES, input, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitItem) (*dto.JudgeResolutionsOrganizationUnitItem, error) {
	res := &dto.GetJudgeResolutionsOrganizationUnitResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.JUDGES+"/"+strconv.Itoa(input.Id), input, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getJudgeResolutionOrganizationUnit(input *dto.JudgeResolutionsOrganizationUnitInput) ([]dto.JudgeResolutionsOrganizationUnitItem, int, error) {
	res := &dto.GetJudgeResolutionsOrganizationUnitListMS{}
	_, err := shared.MakeAPIRequest("GET", config.JUDGES, input, res)
	if err != nil {
		return nil, 0, err
	}

	return res.Data, res.Total, nil
}

func deleteJJudgeResolutionOrganizationUnit(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.JUDGES+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}
