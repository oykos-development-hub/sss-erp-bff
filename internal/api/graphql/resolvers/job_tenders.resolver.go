package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) JobTenderResolver(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.JobTenderResponseItem{}
	var total int

	id := params.Args["id"]
	page := params.Args["page"].(int)
	size := params.Args["size"].(int)
	organizationUnitID := params.Args["organization_unit_id"]
	active := params.Args["active"]
	typeID := params.Args["type_id"]

	if id != nil && shared.IsInteger(id) && id != 0 {
		jobTender, err := r.Repo.GetJobTender(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resItem, _ := buildJobTenderResponse(r.Repo, jobTender)

		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Items:   []dto.JobTenderResponseItem{*resItem},
			Total:   1,
		}, nil

	} else {
		jobTenders, err := r.Repo.GetJobTenderList()
		if err != nil {
			return errors.HandleAPIError(err)
		}
		total = len(jobTenders)

		for _, jobTender := range jobTenders {

			resItem, err := buildJobTenderResponse(r.Repo, jobTender)
			if err != nil {
				return errors.HandleAPIError(err)
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

func buildJobTenderResponse(r repository.MicroserviceRepositoryInterface, item *structs.JobTenders) (*dto.JobTenderResponseItem, error) {
	var (
		jobPosition      *structs.JobPositions
		organizationUnit *structs.OrganizationUnits
		err              error
	)

	tenderType, err := r.GetTenderType(item.TypeID)
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
		Active:              JobTenderIsActive(r, item),
		DateOfStart:         item.DateOfStart,
		DateOfEnd:           item.DateOfEnd,
		FileId:              item.FileId,
		NumberOfVacantSeats: item.NumberOfVacantSeats,
		CreatedAt:           item.CreatedAt,
		UpdatedAt:           item.UpdatedAt,
	}

	if item.OrganizationUnitID != 0 {
		organizationUnit, err = r.GetOrganizationUnitById(item.OrganizationUnitID)
		if err != nil {
			return nil, err
		}
		res.OrganizationUnit = *organizationUnit
	}

	return &res, nil
}

func JobTenderIsActive(r repository.MicroserviceRepositoryInterface, item *structs.JobTenders) bool {

	input := dto.GetJobTenderApplicationsInput{
		JobTenderID: &item.Id,
	}

	jobTenderApplications, err := r.GetTenderApplicationList(&input)

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

func buildJobTenderApplicationResponse(r repository.MicroserviceRepositoryInterface, item *structs.JobTenderApplications) (*dto.JobTenderApplicationResponseItem, error) {
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
		userProfile, err := r.GetUserProfileById(*item.UserProfileId)
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

		evaluation, err := r.GetEmployeeEvaluations(userProfile.Id)
		if err != nil {
			return nil, err
		}
		if len(evaluation) > 0 {
			res.Evaluation = evaluation[len(evaluation)-1].Score
		}
		res.UserProfile = userProfileDropdownItem
	}

	jobTender, err := r.GetJobTender(item.JobTenderId)
	if err != nil {
		return nil, err
	}

	jobTenderResponseItem, _ := buildJobTenderResponse(r, jobTender)

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

func (r *Resolver) JobTenderInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobTenders
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := r.Repo.UpdateJobTender(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		item, err := buildJobTenderResponse(r.Repo, res)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateJobTender(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		item, err := buildJobTenderResponse(r.Repo, res)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) JobTenderDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := r.Repo.DeleteJobTender(itemId)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) JobTenderApplicationsResolver(params graphql.ResolveParams) (interface{}, error) {
	items := []dto.JobTenderApplicationResponseItem{}

	id := params.Args["id"]
	page := params.Args["page"]
	size := params.Args["size"]
	user_profile_id := params.Args["user_profile_id"]

	if id != nil && shared.IsInteger(id) && id != 0 {
		tenderApplication, err := r.Repo.GetTenderApplication(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resItem, _ := buildJobTenderApplicationResponse(r.Repo, tenderApplication)
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

		tenderApplications, err := r.Repo.GetTenderApplicationList(&input)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		total := len(tenderApplications.Data)
		for _, jobTender := range tenderApplications.Data {
			resItem, err := buildJobTenderApplicationResponse(r.Repo, jobTender)
			if err != nil {
				return errors.HandleAPIError(err)
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

func (r *Resolver) JobTenderApplicationInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobTenderApplications
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	if data.UserProfileId != nil {
		userProfile, err := r.Repo.GetUserProfileById(*data.UserProfileId)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		data.FirstName = userProfile.FirstName
		data.LastName = userProfile.LastName
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := r.Repo.UpdateJobTenderApplication(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		item, err := buildJobTenderApplicationResponse(r.Repo, res)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateJobTenderApplication(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		item, err := buildJobTenderApplicationResponse(r.Repo, res)
		if err != nil {
			return errors.HandleAPIError(err)
		}

		response.Item = item
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) JobTenderApplicationDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := r.Repo.DeleteJobTenderApplication(itemId)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
