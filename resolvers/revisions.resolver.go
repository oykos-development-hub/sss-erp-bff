package resolvers

import (
	"bff/config"
	"bff/dto"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"strconv"

	"github.com/graphql-go/graphql"
)

func buildRevisionDetailsItemResponse(revision *structs.Revision) (*dto.RevisionDetailsItem, error) {
	revisorUserProfileDropdown := structs.SettingsDropdown{Id: 0}
	revisionOrganizationUnit := structs.SettingsDropdown{Id: 0}
	responsibleUserProfile := structs.SettingsDropdown{Id: 0}
	implementationUserProfile := structs.SettingsDropdown{Id: 0}

	if revision.RevisorUserProfileID != nil {
		userProfile, err := getUserProfileById(*revision.RevisorUserProfileID)
		if err != nil {
			return nil, err
		}
		revisorUserProfileDropdown.Title = userProfile.FirstName + " " + userProfile.LastName
		revisorUserProfileDropdown.Id = userProfile.Id
	} else {
		if revision.RevisorUserProfile != nil {
			revisorUserProfileDropdown.Title = *revision.RevisorUserProfile
		}
	}

	if revision.ResponsibleUserProfileID != nil {
		userProfile, err := getUserProfileById(*revision.ResponsibleUserProfileID)
		if err != nil {
			return nil, err
		}
		responsibleUserProfile.Title = userProfile.FirstName + " " + userProfile.LastName
		responsibleUserProfile.Id = userProfile.Id
	} else if revision.ResponsibleUserProfile != nil {
		responsibleUserProfile.Title = *revision.ResponsibleUserProfile
	}

	if revision.ImplementationUserProfileID != nil {
		userProfile, err := getUserProfileById(*revision.ImplementationUserProfileID)
		if err != nil {
			return nil, err
		}
		implementationUserProfile.Title = userProfile.FirstName + " " + userProfile.LastName
		implementationUserProfile.Id = userProfile.Id
	} else if revision.ImplementationUserProfile != nil {
		implementationUserProfile.Title = *revision.ImplementationUserProfile
	}

	var err error

	revisionType := &structs.SettingsDropdown{}
	if revision.RevisionTypeID != nil {
		revisionType, err = getDropdownSettingById(*revision.RevisionTypeID)
		if err != nil {
			return nil, err
		}
	}

	if revision.InternalOrganizationUnitID != nil {
		organizationUnit, err := getOrganizationUnitById(*revision.InternalOrganizationUnitID)
		if err != nil {
			return nil, err
		}
		revisionOrganizationUnit.Value = "internal"
		revisionOrganizationUnit.Id = organizationUnit.Id
		revisionOrganizationUnit.Title = organizationUnit.Title
	} else {
		if revision.ExternalOrganizationUnitID != nil {
			organizationUnit, err := getDropdownSettingById(*revision.ExternalOrganizationUnitID)
			if err != nil {
				return nil, err
			}
			revisionOrganizationUnit.Value = "external"
			revisionOrganizationUnit.Id = organizationUnit.Id
			revisionOrganizationUnit.Title = organizationUnit.Title
		}
	}

	revisionItem := &dto.RevisionDetailsItem{
		ID:                              revision.ID,
		Name:                            revision.Name,
		RevisionType:                    *revisionType,
		RevisorUserProfile:              revisorUserProfileDropdown,
		RevisionOrganizationUnit:        revisionOrganizationUnit,
		ResponsibleUserProfile:          responsibleUserProfile,
		ImplementationUserProfile:       implementationUserProfile,
		Title:                           revision.Title,
		PlannedYear:                     revision.PlannedYear,
		PlannedQuarter:                  revision.PlannedQuarter,
		SerialNumber:                    revision.SerialNumber,
		Priority:                        revision.Priority,
		DateOfRevision:                  revision.DateOfRevision,
		DateOfAcceptance:                revision.DateOfAcceptance,
		DateOfRejection:                 revision.DateOfRejection,
		ImplementationSuggestion:        revision.ImplementationSuggestion,
		ImplementationMonthSpan:         revision.ImplementationMonthSpan,
		DateOfImplementation:            revision.DateOfImplementation,
		StateOfImplementation:           revision.StateOfImplementation,
		ImplementationFailedDescription: revision.ImplementationFailedDescription,
		SecondImplementationMonthSpan:   revision.SecondImplementationMonthSpan,
		SecondDateOfRevision:            revision.SecondDateOfRevision,
		FileID:                          revision.FileID,
		RefDocument:                     revision.RefDocument,
		CreatedAt:                       revision.CreatedAt,
		UpdatedAt:                       revision.UpdatedAt,
	}

	return revisionItem, nil
}

func buildRevisionOverviewItemResponse(revision *structs.Revision) (*dto.RevisionOverviewItem, error) {
	userProfileDropdown := structs.SettingsDropdown{
		Id: 0,
	}

	if revision.RevisorUserProfileID != nil {
		userProfile, err := getUserProfileById(*revision.RevisorUserProfileID)
		if err != nil {
			return nil, err
		}
		userProfileDropdown.Title = userProfile.FirstName + " " + userProfile.LastName
		userProfileDropdown.Id = userProfile.Id
	} else {
		if revision.RevisorUserProfile != nil {
			userProfileDropdown.Title = *revision.RevisorUserProfile
		}
	}

	revisionType := &structs.SettingsDropdown{}
	var err error
	if revision.RevisionTypeID != nil {
		revisionType, err = getDropdownSettingById(*revision.RevisionTypeID)
		if err != nil {
			return nil, err
		}
	}

	organizationUnitDropdown := structs.SettingsDropdown{}
	if revision.InternalOrganizationUnitID != nil {
		organizationUnit, err := getOrganizationUnitById(*revision.InternalOrganizationUnitID)
		if err != nil {
			return nil, err
		}
		organizationUnitDropdown.Id = organizationUnit.Id
		organizationUnitDropdown.Title = organizationUnit.Title
	} else {
		if revision.ExternalOrganizationUnitID != nil {
			organizationUnit, err := getDropdownSettingById(*revision.ExternalOrganizationUnitID)
			if err != nil {
				return nil, err
			}
			organizationUnitDropdown.Id = organizationUnit.Id
			organizationUnitDropdown.Title = organizationUnit.Title
		}
	}

	revisionItem := &dto.RevisionOverviewItem{
		Id:                       revision.ID,
		Name:                     revision.Name,
		Title:                    revision.Title,
		RevisorUserProfile:       &userProfileDropdown,
		RevisionType:             revisionType,
		RevisionOrganizationUnit: &organizationUnitDropdown,
		PlannedQuarter:           revision.PlannedQuarter,
		PlannedYear:              revision.PlannedYear,
		CreatedAt:                &revision.CreatedAt,
		UpdatedAt:                &revision.UpdatedAt,
	}

	return revisionItem, nil
}

var RevisionsOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	response := dto.RevisionOverviewResponse{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}
	page := params.Args["page"]
	size := params.Args["size"]
	id := params.Args["id"]
	organizationUnitID := params.Args["organization_unit_id"]
	internal := params.Args["internal"]
	revisorUserProfileID := params.Args["revisor_user_profile_id"]
	revisionType := params.Args["revision_type"]

	if id != nil && id.(int) > 0 {
		revision, err := getRevisionById(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildRevisionOverviewItemResponse(revision)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Items = append(response.Items, *item)
		response.Total = 1
	} else {
		input := dto.GetRevisionsInput{}
		if shared.IsInteger(page) && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if shared.IsInteger(size) && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if organizationUnitID != nil && internal != nil {
			organizationUnitID := organizationUnitID.(int)
			if internal.(bool) {
				input.InternalOrganizationUnitID = &organizationUnitID
			} else {
				input.ExternalOrganizationUnitID = &organizationUnitID
			}
		}
		if revisionType != nil && revisionType.(int) > 0 {
			revisionType := revisionType.(int)
			input.RevisionType = &revisionType
		}
		if revisorUserProfileID != nil && revisorUserProfileID.(int) > 0 {
			revisorUserProfileID := revisorUserProfileID.(int)
			input.RevisorUserProfileID = &revisorUserProfileID
		}

		revisions, err := getRevisionList(&input)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		for _, revision := range revisions.Data {
			item, err := buildRevisionOverviewItemResponse(revision)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			response.Items = append(response.Items, *item)
		}
		response.Total = revisions.Total
	}

	revisorDropdownList, err := getRevisorListDropdown()
	if err != nil {
		return shared.HandleAPIError(err)
	}

	response.Revisors = revisorDropdownList

	return response, nil
}

var RevisionResolver = func(params graphql.ResolveParams) (interface{}, error) {
	page := params.Args["page"]
	size := params.Args["size"]
	id := params.Args["id"]

	input := dto.GetRevisionsInput{}
	if shared.IsInteger(page) && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}
	if shared.IsInteger(size) && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}

	if shared.IsInteger(id) && id.(int) > 0 {
		revision, err := getRevisionById(id.(int))
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildRevisionDetailsItemResponse(revision)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Total:   1,
			Items:   []dto.RevisionDetailsItem{*item},
		}, nil
	}

	revisions, err := getRevisionList(&input)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	items := make([]dto.RevisionDetailsItem, 0, len(revisions.Data))
	for _, revision := range revisions.Data {
		item, err := buildRevisionDetailsItemResponse(revision)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		items = append(items, *item)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Total:   revisions.Total,
		Items:   items,
	}, nil
}

var RevisionDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	revision, err := getRevisionById(id)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	item, err := buildRevisionDetailsItemResponse(revision)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    *item,
	}, nil
}

var RevisionInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Revision
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.ID
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateRevision(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildRevisionDetailsItemResponse(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := createRevision(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildRevisionDetailsItemResponse(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You created this item!"
	}

	return response, nil
}

var RevisionDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteRevision(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func getRevisorListDropdown() ([]*structs.SettingsDropdown, error) {
	revisors, err := getRevisors()
	if err != nil {
		return nil, err
	}

	var revisorDropdownOptions []*structs.SettingsDropdown
	for _, revisor := range revisors {
		revisorDropdownOptions = append(revisorDropdownOptions, &structs.SettingsDropdown{
			Id:    revisor.Id,
			Title: revisor.FirstName + " " + revisor.LastName,
		})
	}
	return revisorDropdownOptions, nil
}

func createRevision(revision *structs.Revision) (*structs.Revision, error) {
	res := &dto.GetRevisionResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.REVISIONS_ENDPOINT, revision, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getRevisors() ([]*structs.UserProfiles, error) {
	res := &dto.GetUserProfileListResponseMS{}
	isRevisor := true
	input := &dto.GetUserProfilesInput{
		IsRevisor: &isRevisor,
	}
	_, err := shared.MakeAPIRequest("GET", config.USER_PROFILES_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func updateRevision(id int, revision *structs.Revision) (*structs.Revision, error) {
	res := &dto.GetRevisionResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.REVISIONS_ENDPOINT+"/"+strconv.Itoa(id), revision, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteRevision(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.REVISIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func getRevisionById(id int) (*structs.Revision, error) {
	res := &dto.GetRevisionResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.REVISIONS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func getRevisionList(input *dto.GetRevisionsInput) (*dto.GetRevisionListResponseMS, error) {
	res := &dto.GetRevisionListResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.REVISIONS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//---------------------------------------------------------------------------------------------- nova polja

var RevisionPlansOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	response := dto.RevisionPlanOverviewResponse{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}
	page := params.Args["page"]
	size := params.Args["size"]
	year := params.Args["year"]

	input := dto.GetPlansInput{}
	if shared.IsInteger(page) && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}
	if shared.IsInteger(size) && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}
	if year != nil {
		year := year.(string)
		input.Year = &year
	}

	revisions, err := getRevisionPlanList(&input)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	response.Items = revisions.Data
	response.Total = revisions.Total

	return response, nil
}

var RevisionPlansDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	revision, err := getRevisionPlanByID(id)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    *revision,
	}, nil
}

var RevisionPlanDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteRevisionPlan(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

var RevisionPlanInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data dto.RevisionPlanItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateRevisionPlan(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Item = res
		response.Message = "You updated this item!"
	} else {
		res, err := createRevisionPlan(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}

		response.Item = res
		response.Message = "You created this item!"
	}

	return response, nil
}

func getRevisionPlanList(input *dto.GetPlansInput) (*dto.GetRevisionPlanResponseMS, error) {
	res := &dto.GetRevisionPlanResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.REVISION_PLAN_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getRevisionPlanByID(id int) (*dto.RevisionPlanItem, error) {
	res := &dto.GetPlanResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.REVISION_PLAN_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteRevisionPlan(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.REVISION_PLAN_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func createRevisionPlan(plan *dto.RevisionPlanItem) (*dto.RevisionPlanItem, error) {
	res := &dto.GetPlanResponseMS{}
	_, err := shared.MakeAPIRequest("POST", config.REVISION_PLAN_ENDPOINT, plan, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateRevisionPlan(id int, plan *dto.RevisionPlanItem) (*dto.RevisionPlanItem, error) {
	res := &dto.GetPlanResponseMS{}
	_, err := shared.MakeAPIRequest("PUT", config.REVISION_PLAN_ENDPOINT+"/"+strconv.Itoa(id), plan, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

//--------------------------------------------------------------------------

func buildRevisionItemResponse(revision *structs.Revisions) (*dto.RevisionsOverviewItem, error) {

	revisiontype, err := getDropdownSettingById(revision.RevisionType)
	if err != nil {
		return nil, err
	}

	revisionType := dto.DropdownSimple{
		Id:    revisiontype.Id,
		Title: revisiontype.Title,
	}

	revisor, err := getUserProfileById(revision.Revisor)
	if err != nil {
		return nil, err
	}

	revisorDropdown := dto.DropdownSimple{
		Id:    revisor.Id,
		Title: revisor.FirstName + " " + revisor.LastName,
	}

	internalUnitDropdown := &dto.DropdownSimple{}
	externalUnitDropdown := &dto.DropdownSimple{}

	if revision.InternalRevisionSubject != nil {
		organizationUnit, err := getOrganizationUnitById(*revision.InternalRevisionSubject)
		if err != nil {
			return nil, err
		}
		internalUnitDropdown.Id = organizationUnit.Id
		internalUnitDropdown.Title = organizationUnit.Title
	}
	if revision.ExternalRevisionSubject != nil {
		organizationUnit, err := getDropdownSettingById(*revision.ExternalRevisionSubject)
		if err != nil {
			return nil, err
		}
		externalUnitDropdown.Id = organizationUnit.Id
		externalUnitDropdown.Title = organizationUnit.Title
	}

	revisionItem := &dto.RevisionsOverviewItem{
		ID:                      revision.ID,
		Title:                   revision.Title,
		PlanID:                  revision.PlanID,
		SerialNumber:            revision.SerialNumber,
		DateOfRevision:          revision.DateOfRevision,
		RevisionQuartal:         revision.RevisionQuartal,
		InternalRevisionSubject: internalUnitDropdown,
		ExternalRevisionSubject: externalUnitDropdown,
		Revisor:                 revisorDropdown,
		RevisionType:            revisionType,
		FileID:                  revision.FileID,
		CreatedAt:               revision.CreatedAt,
		UpdatedAt:               revision.UpdatedAt,
	}

	return revisionItem, nil
}

var RevisionOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	response := dto.RevisionsOverviewResponse{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}
	page := params.Args["page"]
	size := params.Args["size"]
	revisor := params.Args["revisor_id"]
	revisionType := params.Args["revision_type_id"]
	internal := params.Args["internal_revision_subject_id"]
	plan := params.Args["plan_id"]

	input := dto.GetRevisionFilter{}
	if shared.IsInteger(page) && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}

	if shared.IsInteger(size) && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}

	if shared.IsInteger(revisor) && revisor.(int) > 0 {
		temp := revisor.(int)
		input.Revisor = &temp
	}

	if shared.IsInteger(revisionType) && revisionType.(int) > 0 {
		temp := revisionType.(int)
		input.RevisionType = &temp
	}

	if shared.IsInteger(internal) && internal.(int) > 0 {
		temp := internal.(int)
		input.InternalRevisionSubject = &temp
	}

	if shared.IsInteger(plan) && plan.(int) > 0 {
		temp := plan.(int)
		input.PlanID = &temp
	}

	revisions, err := getRevisionsList(&input)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, revision := range revisions.Data {
		item, err := buildRevisionItemResponse(revision)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Items = append(response.Items, *item)
	}
	response.Total = revisions.Total

	revisorDropdownList, err := getRevisorListDropdown()
	if err != nil {
		return shared.HandleAPIError(err)
	}

	response.Revisors = revisorDropdownList

	return response, nil
}

var RevisionDetailResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	revision, err := getRevisionsByID(id)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	item, err := buildRevisionItemResponse(revision)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	revisorDropdownList, err := getRevisorListDropdown()
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.RevisionsDetailsResponse{
		Status:   "success",
		Message:  "Here's the list you asked for!",
		Item:     *item,
		Revisors: revisorDropdownList,
	}, nil
}

var RevisionsInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.RevisionsInsert
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	data1 := structs.Revisions{
		ID:                      data.ID,
		Title:                   data.Title,
		PlanID:                  data.PlanID,
		SerialNumber:            data.SerialNumber,
		DateOfRevision:          data.DateOfRevision,
		RevisionQuartal:         data.RevisionQuartal,
		ExternalRevisionSubject: data.ExternalRevisionSubject,
		Revisor:                 data.Revisor,
		RevisionType:            data.RevisionType,
		FileID:                  data.FileID,
	}

	itemId := data.ID
	if shared.IsInteger(itemId) && itemId != 0 {
		data1.InternalRevisionSubject = data.InternalRevisionSubject[0]
		res, err := updateRevisions(itemId, &data1)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildRevisionItemResponse(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		var responseItems []*dto.RevisionsOverviewItem
		if data.InternalRevisionSubject != nil {
			for _, orgUnit := range data.InternalRevisionSubject {
				data1.InternalRevisionSubject = orgUnit
				res, err := createRevisions(&data1)
				if err != nil {
					return shared.HandleAPIError(err)
				}
				item, err := buildRevisionItemResponse(res)
				if err != nil {
					return shared.HandleAPIError(err)
				}
				responseItems = append(responseItems, item)
				response.Item = responseItems
			}
		} else {
			res, err := createRevisions(&data1)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			item, err := buildRevisionItemResponse(res)
			if err != nil {
				return shared.HandleAPIError(err)
			}
			response.Item = item
		}
		response.Message = "You created this item!"
	}

	return response, nil
}

var RevisionsDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteRevisions(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func getRevisionsList(input *dto.GetRevisionFilter) (*dto.GetRevisionsResponseMS, error) {
	res := &dto.GetRevisionsResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.REVISION_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getRevisionsByID(id int) (*structs.Revisions, error) {
	res := &dto.GetRevisionMS{}
	_, err := shared.MakeAPIRequest("GET", config.REVISION_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteRevisions(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.REVISION_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func createRevisions(plan *structs.Revisions) (*structs.Revisions, error) {
	res := &dto.GetRevisionMS{}
	_, err := shared.MakeAPIRequest("POST", config.REVISION_ENDPOINT, plan, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateRevisions(id int, plan *structs.Revisions) (*structs.Revisions, error) {
	res := &dto.GetRevisionMS{}
	_, err := shared.MakeAPIRequest("PUT", config.REVISION_ENDPOINT+"/"+strconv.Itoa(id), plan, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

//-------------------------------------------------------------------------

func buildRevisionTipItemResponse(revision *structs.RevisionTips) (*dto.RevisionTipsOverviewItem, error) {

	revisorDropdown := structs.SettingsDropdown{}

	if revision.UserProfileID != nil {
		revisor, err := getUserProfileById(*revision.UserProfileID)
		if err != nil {
			return nil, err
		}

		revisorDropdown = structs.SettingsDropdown{
			Id:    revisor.Id,
			Title: revisor.FirstName + " " + revisor.LastName,
		}
	}

	revisionTipItem := &dto.RevisionTipsOverviewItem{
		ID:                     revision.ID,
		RevisionID:             revision.RevisionID,
		UserProfile:            revisorDropdown,
		DateOfAccept:           revision.DateOfAccept,
		DueDate:                revision.DueDate,
		NewDueDate:             revision.NewDueDate,
		EndDate:                revision.EndDate,
		DateOfReject:           revision.DateOfReject,
		DateOfExecution:        revision.DateOfExecution,
		NewDateOfExecution:     revision.NewDateOfExecution,
		Recommendation:         revision.Recommendation,
		RevisionPriority:       revision.RevisionPriority,
		Status:                 revision.Status,
		ResponsiblePerson:      revision.ResponsiblePerson,
		Documents:              revision.Documents,
		ReasonsForNonExecuting: revision.ReasonsForNonExecuting,
		FileID:                 revision.FileID,
		CreatedAt:              revision.CreatedAt,
		UpdatedAt:              revision.UpdatedAt,
	}

	return revisionTipItem, nil
}

var RevisionTipsOverviewResolver = func(params graphql.ResolveParams) (interface{}, error) {
	response := dto.RevisionTipsOverviewResponse{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}
	page := params.Args["page"]
	size := params.Args["size"]
	revision := params.Args["revision_id"]

	input := dto.GetRevisionTipFilter{}
	if shared.IsInteger(page) && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}

	if shared.IsInteger(size) && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}

	if shared.IsInteger(revision) && revision.(int) > 0 {
		temp := revision.(int)
		input.RevisionID = &temp
	}

	revisions, err := getRevisionTipsList(&input)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	for _, revision := range revisions.Data {
		item, err := buildRevisionTipItemResponse(revision)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Items = append(response.Items, *item)
	}
	response.Total = revisions.Total

	revisorDropdownList, err := getRevisorListDropdown()
	if err != nil {
		return shared.HandleAPIError(err)
	}

	response.Revisors = revisorDropdownList

	return response, nil
}

var RevisionTipsDetailsResolver = func(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	revision, err := getRevisionTipByID(id)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	item, err := buildRevisionTipItemResponse(revision)
	if err != nil {
		return shared.HandleAPIError(err)
	}
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    *item,
	}, nil
}

var RevisionTipsInsertResolver = func(params graphql.ResolveParams) (interface{}, error) {
	var data structs.RevisionTips
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemId := data.ID
	if shared.IsInteger(itemId) && itemId != 0 {
		res, err := updateRevisionTips(itemId, &data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildRevisionTipItemResponse(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := createRevisionTips(&data)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		item, err := buildRevisionTipItemResponse(res)
		if err != nil {
			return shared.HandleAPIError(err)
		}
		response.Item = item
		response.Message = "You created this item!"
	}

	return response, nil
}

var RevisionTipsDeleteResolver = func(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"].(int)

	err := deleteRevisionTips(itemId)
	if err != nil {
		return shared.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func getRevisionTipsList(input *dto.GetRevisionTipFilter) (*dto.GetRevisionTipsResponseMS, error) {
	res := &dto.GetRevisionTipsResponseMS{}
	_, err := shared.MakeAPIRequest("GET", config.REVISION_TIPS_ENDPOINT, input, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getRevisionTipByID(id int) (*structs.RevisionTips, error) {
	res := &dto.GetRevisionTipMS{}
	_, err := shared.MakeAPIRequest("GET", config.REVISION_TIPS_ENDPOINT+"/"+strconv.Itoa(id), nil, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func deleteRevisionTips(id int) error {
	_, err := shared.MakeAPIRequest("DELETE", config.REVISION_TIPS_ENDPOINT+"/"+strconv.Itoa(id), nil, nil)
	if err != nil {
		return err
	}

	return nil
}

func createRevisionTips(plan *structs.RevisionTips) (*structs.RevisionTips, error) {
	res := &dto.GetRevisionTipMS{}
	_, err := shared.MakeAPIRequest("POST", config.REVISION_TIPS_ENDPOINT, plan, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}

func updateRevisionTips(id int, plan *structs.RevisionTips) (*structs.RevisionTips, error) {
	res := &dto.GetRevisionTipMS{}
	_, err := shared.MakeAPIRequest("PUT", config.REVISION_TIPS_ENDPOINT+"/"+strconv.Itoa(id), plan, res)
	if err != nil {
		return nil, err
	}

	return &res.Data, nil
}
