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

var JobTenderResolver = func(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.JobTenderResponseItem{}
	var total int

	id := params.Args["id"]
	page := params.Args["page"].(int)
	size := params.Args["size"].(int)
	organizationUnitID := params.Args["organization_unit_id"]
	active := params.Args["active"]
	typeID := params.Args["type_id"]
	if active != nil {
		active = active.(bool)
	}
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
		input := dto.GetJobTendersInput{}

		jobTenders, err := getJobTenderList(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		total = len(jobTenders)

		for _, jobTender := range jobTenders {

			resItem, err := buildJobTenderResponse(jobTender)
			if err != nil {
				return shared.HandleAPIError(err)
			}

			if active != nil && active != resItem.Active {
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

		paginatedItems, err := shared.Paginate(items, page, size)
		if err != nil {
			fmt.Printf("Error paginating items: %v", err)
		}
		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   paginatedItems,
			Total:   total,
		}, nil
	}
}

func buildJobTenderResponse(item *structs.JobTenders) (*dto.JobTenderResponseItem, error) {
	jobPositionInOrganizationUnit, err := getJobPositionsInOrganizationUnitsById(item.PositionInOrganizationUnitId)
	if err != nil {
		return nil, err
	}
	jobPosition, err := getJobPositionById(jobPositionInOrganizationUnit.JobPositionId)
	if err != nil {
		return nil, err
	}
	organizationUnit, err := getOrganizationUnitById(jobPositionInOrganizationUnit.ParentOrganizationUnitId)
	if err != nil {
		return nil, err
	}

	tenderType, err := getTenderType(item.TypeID)
	if err != nil {
		return nil, err
	}

	res := dto.JobTenderResponseItem{
		Id:               item.Id,
		OrganizationUnit: *organizationUnit,
		JobPosition:      *jobPosition,
		Type:             *tenderType,
		Description:      item.Description,
		SerialNumber:     item.SerialNumber,
		AvailableSlots:   item.AvailableSlots,
		Active:           JobTenderIsActive(item),
		DateOfStart:      item.DateOfStart,
		DateOfEnd:        item.DateOfEnd,
		FileId:           item.FileId,
		CreatedAt:        item.CreatedAt,
		UpdatedAt:        item.UpdatedAt,
	}

	return &res, nil
}

func JobTenderIsActive(item *structs.JobTenders) bool {
	const dateFormat = "2006-01-02T15:04:05Z"

	startDate, err := time.Parse(dateFormat, string(item.DateOfStart))
	if err != nil {
		return false
	}

	endDate, err := time.Parse(dateFormat, string(item.DateOfEnd))
	if err != nil {
		return false
	}

	currentDate := time.Now().UTC()

	return currentDate.After(startDate) && currentDate.Before(endDate)
}

func buildJobTenderApplicationResponse(item *structs.JobTenderApplications) (*dto.JobTenderApplicationResponseItem, error) {
	userProfile, err := getUserProfileById(item.UserProfileId)
	if err != nil {
		return nil, err
	}
	jobTender, err := getJobTender(item.JobTenderId)
	if err != nil {
		return nil, err
	}

	userProfileDropdownItem := structs.SettingsDropdown{
		Id:    userProfile.Id,
		Title: userProfile.FirstName + " " + userProfile.LastName,
	}

	jobTenderDropdownItem := structs.SettingsDropdown{
		Id:    jobTender.Id,
		Title: jobTender.SerialNumber,
	}

	res := dto.JobTenderApplicationResponseItem{
		Id:          item.Id,
		UserProfile: userProfileDropdownItem,
		JobTender:   jobTenderDropdownItem,
		Active:      item.Active,
		FileId:      item.FileId,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
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
	total := 0

	if id != nil && shared.IsInteger(id) && id != 0 {
		tenderApplication, err := getTenderApplication(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		resItem, _ := buildJobTenderApplicationResponse(tenderApplication)
		items = append(items, *resItem)
		total = 1
	} else {
		input := dto.GetJobTenderApplicationsInput{}
		if shared.IsInteger(page) && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if shared.IsInteger(size) && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if jobTenderID, ok := params.Args["job_tender_id"].(int); ok && jobTenderID != 0 {
			input.JobTenderID = &jobTenderID
		}

		tenderApplications, err := getTenderApplicationList(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, jobTender := range tenderApplications.Data {
			resItem, err := buildJobTenderApplicationResponse(jobTender)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			items = append(items, *resItem)
		}
		total = tenderApplications.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   total,
	}, nil
}

var JobTenderApplicationInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobTenderApplications
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

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

func getJobTenderList(input *dto.GetJobTendersInput) ([]*structs.JobTenders, error) {
	res := &dto.GetJobTenderListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.JOB_TENDERS_ENDPOINT, input, res)
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
