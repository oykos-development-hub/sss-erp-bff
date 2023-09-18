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
