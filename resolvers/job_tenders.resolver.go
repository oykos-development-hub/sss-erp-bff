package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
)

var JobTenderResolver = func(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.JobTenderResponseItem{}
	var total int

	id := params.Args["id"]
	page := params.Args["page"].(int)
	size := params.Args["size"].(int)
	organizationUnitID := params.Args["organization_unit_id"]
	active := params.Args["active"]
	typeID := params.Args["type_id"]

	if id != nil && shared.IsInteger(id) && id != 0 {
		jobTender, err := getJobTender(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resItem, _ := buildJobTenderResponse(jobTender)

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []dto.JobTenderResponseItem{*resItem},
			Total:   1,
		}, nil

	} else {
		jobTenders, err := getJobTenderList()
		if err != nil {
			return shared.HandleAPIError(err)
		}
		total = len(jobTenders)

		for _, jobTender := range jobTenders {

			resItem, err := buildJobTenderResponse(jobTender)
			if err != nil {
				return shared.HandleAPIError(err)
			}

			if active != nil && active.(bool) != resItem.Active {
				total--
				continue
			}

			if organizationUnitID != nil &&
				organizationUnitID.(int) > 0 &&
				resItem.OrganizationUnit.Id != organizationUnitID {
				total--
				continue
			}

			if typeID != nil &&
				typeID.(int) > 0 &&
				resItem.Type.Id != typeID {
				total--
				continue
			}

			items = append(items, *resItem)
		}

		paginatedItems, _ := shared.Paginate(items, page, size)

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   paginatedItems,
			Total:   total,
		}, nil
	}
}

func buildJobTenderResponse(item *structs.JobTenders) (*dto.JobTenderResponseItem, error) {
	var (
		jobPosition      *structs.JobPositions
		organizationUnit *structs.OrganizationUnits
		err              error
	)

	tenderType, err := getTenderType(item.TypeID)
	if err != nil {
		return nil, err
	}

	res := dto.JobTenderResponseItem{
		Id:                  item.Id,
		JobPosition:         jobPosition,
		Type:                *tenderType,
		Description:         item.Description,
		SerialNumber:        item.SerialNumber,
		Title:               item.SerialNumber,
		Active:              JobTenderIsActive(item),
		DateOfStart:         item.DateOfStart,
		DateOfEnd:           item.DateOfEnd,
		FileId:              item.FileId,
		NumberOfVacantSeats: item.NumberOfVacantSeats,
		CreatedAt:           item.CreatedAt,
		UpdatedAt:           item.UpdatedAt,
	}

	if item.OrganizationUnitID != 0 {
		organizationUnit, err = getOrganizationUnitById(item.OrganizationUnitID)
		if err != nil {
			return nil, err
		}
		res.OrganizationUnit = *organizationUnit
	}

	return &res, nil
}

func JobTenderIsActive(item *structs.JobTenders) bool {

	input := dto.GetJobTenderApplicationsInput{
		JobTenderID: &item.Id,
	}

	jobTenderApplications, err := getTenderApplicationList(&input)

	if err != nil {
		return false
	}

	count := 0
	for _, tenderApp := range jobTenderApplications.Data {
		if tenderApp.Status == "Izabran" {
			count++
		}
	}

	if count == item.NumberOfVacantSeats {
		return false
	}

	start, _ := time.Parse(time.RFC3339, item.DateOfStart)

	var end *time.Time
	if item.DateOfEnd != nil {
		endDate, _ := time.Parse(time.RFC3339, *item.DateOfEnd)
		end = &endDate
	}

	currentDate := time.Now().UTC()

	return currentDate.After(start) && (end == nil || currentDate.Before(*end))
}

func buildJobTenderApplicationResponse(item *structs.JobTenderApplications) (*dto.JobTenderApplicationResponseItem, error) {
	res := dto.JobTenderApplicationResponseItem{
		Id:                 item.Id,
		Type:               item.Type,
		FirstName:          item.FirstName,
		LastName:           item.LastName,
		OfficialPersonalID: item.OfficialPersonalID,
		DateOfBirth:        item.DateOfBirth,
		Nationality:        item.Nationality,
		Evaluation:         item.Evaluation,
		DateOfAplication:   item.DateOfApplication,
		Active:             item.Active,
		FileId:             item.FileId,
		Status:             item.Status,
		CreatedAt:          item.CreatedAt,
		UpdatedAt:          item.UpdatedAt,
	}

	if item.UserProfileId != nil {
		userProfile, err := getUserProfileById(*item.UserProfileId)
		if err != nil {
			return nil, err
		}
		userProfileDropdownItem := &dto.DropdownSimple{
			Id:    userProfile.Id,
			Title: userProfile.GetFullName(),
		}
		res.FirstName = userProfile.FirstName
		res.LastName = userProfile.LastName
		res.OfficialPersonalID = userProfile.OfficialPersonalId
		res.DateOfBirth = userProfile.DateOfBirth
		res.Nationality = userProfile.Citizenship

		evaluation, err := getEmployeeEvaluations(userProfile.Id)
		if err != nil {
			return nil, err
		}
		if len(evaluation) > 0 {
			res.Evaluation = evaluation[len(evaluation)-1].Score
		}
		res.UserProfile = userProfileDropdownItem
	}

	jobTender, err := getJobTender(item.JobTenderId)
	if err != nil {
		return nil, err
	}

	jobTenderResponseItem, _ := buildJobTenderResponse(jobTender)

	res.JobTender = jobTenderResponseItem
	res.OrganizationUnit = &dto.DropdownSimple{
		Id:    jobTenderResponseItem.OrganizationUnit.Id,
		Title: jobTenderResponseItem.OrganizationUnit.Title,
	}
	res.TenderType = &dto.DropdownSimple{
		Id:    jobTenderResponseItem.Type.Id,
		Title: jobTenderResponseItem.Type.Title,
	}

	return &res, nil
}

var JobTenderInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobTenders
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateJobTender(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildJobTenderResponse(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := createJobTender(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildJobTenderResponse(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You created this item!"
	}

	return response, nil
}

var JobTenderDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteJobTender(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

var JobTenderApplicationsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.JobTenderApplicationResponseItem{}

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	user_profile_id := params.Args["user_profile_id"]

	if id != nil && shared.IsInteger(id) && id != 0 {
		tenderApplication, err := getTenderApplication(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resItem, _ := buildJobTenderApplicationResponse(tenderApplication)
		items = append(items, *resItem)

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   items,
			Total:   1,
		}, nil

	} else {
		input := dto.GetJobTenderApplicationsInput{}
		if shared.IsInteger(user_profile_id) && user_profile_id.(int) > 0 {
			userProfileId := user_profile_id.(int)
			input.UserProfileId = &userProfileId
		}
		if jobTenderID, ok := params.Args["job_tender_id"].(int); ok && jobTenderID != 0 {
			input.JobTenderID = &jobTenderID
		}
		if search, ok := params.Args["search"].(string); ok && search != "" {
			input.Search = &search
		}

		tenderApplications, err := getTenderApplicationList(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		total := len(tenderApplications.Data)
		for _, jobTender := range tenderApplications.Data {
			resItem, err := buildJobTenderApplicationResponse(jobTender)
			if err != nil {
				return shared.HandleAPIError(err)
			}

			if filerTenderTypeID, ok := params.Args["type_id"].(int); ok && filerTenderTypeID != 0 && resItem.JobTender.Type.Id != filerTenderTypeID {
				total--
				continue
			}
			if organizationUnitID, ok := params.Args["organization_unit_id"].(int); ok && organizationUnitID != 0 && resItem.JobTender.OrganizationUnit.Id != organizationUnitID {
				total--
				continue
			}
			items = append(items, *resItem)
		}

		paginatedItems, _ := shared.Paginate(items, page.(int), size.(int))

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   paginatedItems,
			Total:   total,
		}, nil
	}
}

var JobTenderApplicationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobTenderApplications
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	if data.UserProfileId != nil {
		userProfile, err := getUserProfileById(*data.UserProfileId)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		data.FirstName = userProfile.FirstName
		data.LastName = userProfile.LastName
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateJobTenderApplication(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		item, err := buildJobTenderApplicationResponse(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := createJobTenderApplication(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		item, err := buildJobTenderApplicationResponse(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Item = item
		response.Message = "You created this item!"
	}

	if data.UserProfileId != nil && data.Status == "Izabran" {
		active := true
		input := dto.GetJudgeResolutionListInputMS{
			Active: &active,
		}

		resolution, err := getJudgeResolutionList(&input)

		if err != nil {
			return shared.HandleAPIError(err)
		}

		filter := dto.JudgeResolutionsOrganizationUnitInput{
			UserProfileId: data.UserProfileId,
			ResolutionId:  &resolution.Data[0].Id,
		}

		judge, _, err := getJudgeResolutionOrganizationUnit(&filter)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		err = deleteJJudgeResolutionOrganizationUnit(judge[0].Id)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		contract, err := getEmployeeContracts(*data.UserProfileId, nil)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		now := time.Now()
		format := "2006-01-02T00:00:00Z"
		dateOfEnd := now.Format(format)
		dateOfStart, _ := time.Parse(format, *contract[0].DateOfStart)
		yearsDiff := now.Year() - dateOfStart.Year()
		monthsDiff := int(now.Month()) - int(dateOfStart.Month())

		if monthsDiff < 0 {
			monthsDiff += 12
			yearsDiff--
		}

		totalMonths := (yearsDiff * 12) + monthsDiff

		experience := structs.Experience{
			UserProfileId:             *data.UserProfileId,
			OrganizationUnitId:        judge[0].OrganizationUnitId,
			Relevant:                  true,
			DateOfStart:               *contract[0].DateOfStart,
			DateOfEnd:                 dateOfEnd,
			AmountOfExperience:        totalMonths,
			AmountOfInsuredExperience: totalMonths,
		}
		_, err = createExperience(&experience)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		jobTender, err := getJobTender(data.JobTenderId)

		if err != nil {
			return shared.HandleAPIError(err)
		}

		contract[0].DateOfStart = nil
		contract[0].DateOfEnd = nil
		contract[0].OrganizationUnitID = jobTender.OrganizationUnitID

		_, err = updateEmployeeContract(contract[0].Id, contract[0])

		if err != nil {
			return shared.HandleAPIError(err)
		}

		judge[0].OrganizationUnitId = jobTender.OrganizationUnitID
		_, err = createJudgeResolutionOrganizationUnit(&judge[0])

		if err != nil {
			return shared.HandleAPIError(err)
		}
	}

	return response, nil
}

var JobTenderApplicationDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteJobTenderApplication(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func createJobTender(jobTender *structs.JobTenders) (*structs.JobTenders, error) {
	res := &dto.GetJobTenderResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.JOB_TENDERS_ENDPOINT, jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateJobTender(id int, jobTender *structs.JobTenders) (*structs.JobTenders, error) {
	res := &dto.GetJobTenderResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.JOB_TENDERS_ENDPOINT+"/"+strconv.Itoa(id), jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getJobTender(id int) (*structs.JobTenders, error) {
	res := &dto.GetJobTenderResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_TENDERS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getJobTenderList() ([]*structs.JobTenders, error) {
	res := &dto.GetJobTenderListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_TENDERS_ENDPOINT, nil, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func deleteJobTender(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.JOB_TENDERS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func createJobTenderApplication(jobTender *structs.JobTenderApplications) (*structs.JobTenderApplications, error) {
	res := &dto.GetJobTenderApplicationResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.JOB_TENDER_APPLICATIONS_ENDPOINT, jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateJobTenderApplication(id int, jobTender *structs.JobTenderApplications) (*structs.JobTenderApplications, error) {
	currentTenderApplication, _ := getTenderApplication(id)
	if currentTenderApplication.Status != "Izabran" && jobTender.Status == "Izabran" {
		applications, _ := getTenderApplicationList(&dto.GetJobTenderApplicationsInput{JobTenderID: &currentTenderApplication.JobTenderId})
		for _, application := range applications.Data {
			if currentTenderApplication.Id != application.Id {
				application.Status = "Nije izabran"
				_, err := shared.MakeAPIRequest("PUT", config.JOB_TENDER_APPLICATIONS_ENDPOINT+"/"+strconv.Itoa(application.Id), application, nil)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	res := &dto.GetJobTenderApplicationResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.JOB_TENDER_APPLICATIONS_ENDPOINT+"/"+strconv.Itoa(id), jobTender, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteJobTenderApplication(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.JOB_TENDER_APPLICATIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getTenderApplication(id int) (*structs.JobTenderApplications, error) {
	res := &dto.GetJobTenderApplicationResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_TENDER_APPLICATIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getTenderApplicationList(input *dto.GetJobTenderApplicationsInput) (*dto.GetJobTenderApplicationListResponseMS, error) {
	res := &dto.GetJobTenderApplicationListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_TENDER_APPLICATIONS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
