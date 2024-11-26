package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func buildRevisionDetailsItemResponse(r repository.MicroserviceRepositoryInterface, revision *structs.Revision) (*dto.RevisionDetailsItem, error) {
	revisorUserProfileDropdown := structs.SettingsDropdown{ID: 0}
	revisionOrganizationUnit := structs.SettingsDropdown{ID: 0}
	responsibleUserProfile := structs.SettingsDropdown{ID: 0}
	implementationUserProfile := structs.SettingsDropdown{ID: 0}

	if revision.RevisorUserProfileID != nil {
		userProfile, err := r.GetUserProfileByID(*revision.RevisorUserProfileID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}
		revisorUserProfileDropdown.Title = userProfile.FirstName + " " + userProfile.LastName
		revisorUserProfileDropdown.ID = userProfile.ID
	} else {
		if revision.RevisorUserProfile != nil {
			revisorUserProfileDropdown.Title = *revision.RevisorUserProfile
		}
	}

	if revision.ResponsibleUserProfileID != nil {
		userProfile, err := r.GetUserProfileByID(*revision.ResponsibleUserProfileID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}
		responsibleUserProfile.Title = userProfile.FirstName + " " + userProfile.LastName
		responsibleUserProfile.ID = userProfile.ID
	} else if revision.ResponsibleUserProfile != nil {
		responsibleUserProfile.Title = *revision.ResponsibleUserProfile
	}

	if revision.ImplementationUserProfileID != nil {
		userProfile, err := r.GetUserProfileByID(*revision.ImplementationUserProfileID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}
		implementationUserProfile.Title = userProfile.FirstName + " " + userProfile.LastName
		implementationUserProfile.ID = userProfile.ID
	} else if revision.ImplementationUserProfile != nil {
		implementationUserProfile.Title = *revision.ImplementationUserProfile
	}

	var err error

	revisionType := &structs.SettingsDropdown{}
	if revision.RevisionTypeID != nil {
		revisionType, err = r.GetDropdownSettingByID(*revision.RevisionTypeID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}
	}

	if revision.InternalOrganizationUnitID != nil {
		organizationUnit, err := r.GetOrganizationUnitByID(*revision.InternalOrganizationUnitID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}
		revisionOrganizationUnit.Value = "internal"
		revisionOrganizationUnit.ID = organizationUnit.ID
		revisionOrganizationUnit.Title = organizationUnit.Title
	} else {
		if revision.ExternalOrganizationUnitID != nil {
			organizationUnit, err := r.GetDropdownSettingByID(*revision.ExternalOrganizationUnitID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get dropdown setting by id")
			}
			revisionOrganizationUnit.Value = "external"
			revisionOrganizationUnit.ID = organizationUnit.ID
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

	if revision.TipsFileID != nil && *revision.TipsFileID != 0 {
		file, err := r.GetFileByID(*revision.TipsFileID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}
		revisionItem.TipsFileID.ID = file.ID
		revisionItem.TipsFileID.Name = file.Name
		revisionItem.TipsFileID.Type = *file.Type
	}

	return revisionItem, nil
}

func buildRevisionOverviewItemResponse(r repository.MicroserviceRepositoryInterface, revision *structs.Revision) (*dto.RevisionOverviewItem, error) {
	userProfileDropdown := structs.SettingsDropdown{
		ID: 0,
	}

	if revision.RevisorUserProfileID != nil {
		userProfile, err := r.GetUserProfileByID(*revision.RevisorUserProfileID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}
		userProfileDropdown.Title = userProfile.FirstName + " " + userProfile.LastName
		userProfileDropdown.ID = userProfile.ID
	} else {
		if revision.RevisorUserProfile != nil {
			userProfileDropdown.Title = *revision.RevisorUserProfile
		}
	}

	revisionType := &structs.SettingsDropdown{}
	var err error
	if revision.RevisionTypeID != nil {
		revisionType, err = r.GetDropdownSettingByID(*revision.RevisionTypeID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}
	}

	organizationUnitDropdown := structs.SettingsDropdown{}
	if revision.InternalOrganizationUnitID != nil {
		organizationUnit, err := r.GetOrganizationUnitByID(*revision.InternalOrganizationUnitID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get organization unit by id")
		}
		organizationUnitDropdown.ID = organizationUnit.ID
		organizationUnitDropdown.Title = organizationUnit.Title
	} else {
		if revision.ExternalOrganizationUnitID != nil {
			organizationUnit, err := r.GetDropdownSettingByID(*revision.ExternalOrganizationUnitID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get dropdown setting by id")
			}
			organizationUnitDropdown.ID = organizationUnit.ID
			organizationUnitDropdown.Title = organizationUnit.Title
		}
	}

	revisionItem := &dto.RevisionOverviewItem{
		ID:                       revision.ID,
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

func (r *Resolver) RevisionsOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
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
		revision, err := r.Repo.GetRevisionByID(id.(int))
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildRevisionOverviewItemResponse(r.Repo, revision)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Items = append(response.Items, *item)
		response.Total = 1
	} else {
		input := dto.GetRevisionsInput{}
		if page != nil && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if size != nil && size.(int) > 0 {
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

		revisions, err := r.Repo.GetRevisionList(&input)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, revision := range revisions.Data {
			item, err := buildRevisionOverviewItemResponse(r.Repo, revision)
			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
			response.Items = append(response.Items, *item)
		}
		response.Total = revisions.Total
	}

	revisorDropdownList, err := getRevisorListDropdown(r.Repo)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Revisors = revisorDropdownList

	return response, nil
}

func (r *Resolver) RevisionResolver(params graphql.ResolveParams) (interface{}, error) {
	page := params.Args["page"]
	size := params.Args["size"]
	id := params.Args["id"]

	input := dto.GetRevisionsInput{}
	if page != nil && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}
	if size != nil && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}

	if id != nil && id.(int) > 0 {
		revision, err := r.Repo.GetRevisionByID(id.(int))
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildRevisionDetailsItemResponse(r.Repo, revision)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		return dto.Response{
			Status:  "success",
			Message: "Here's the list you asked for!",
			Total:   1,
			Items:   []dto.RevisionDetailsItem{*item},
		}, nil
	}

	revisions, err := r.Repo.GetRevisionList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	items := make([]dto.RevisionDetailsItem, 0, len(revisions.Data))
	for _, revision := range revisions.Data {
		item, err := buildRevisionDetailsItemResponse(r.Repo, revision)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
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

func (r *Resolver) RevisionDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	revision, err := r.Repo.GetRevisionByID(id)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	item, err := buildRevisionDetailsItemResponse(r.Repo, revision)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    *item,
	}, nil
}

func (r *Resolver) RevisionInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Revision
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		res, err := r.Repo.UpdateRevision(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildRevisionDetailsItemResponse(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateRevision(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildRevisionDetailsItemResponse(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Item = item
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) RevisionDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteRevision(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func getRevisorListDropdown(r repository.MicroserviceRepositoryInterface) ([]*structs.SettingsDropdown, error) {
	revisors, err := r.GetRevisors()
	if err != nil {
		return nil, errors.Wrap(err, "repo get revisors")
	}

	var revisorDropdownOptions []*structs.SettingsDropdown
	for _, revisor := range revisors {
		revisorDropdownOptions = append(revisorDropdownOptions, &structs.SettingsDropdown{
			ID:    revisor.ID,
			Title: revisor.FirstName + " " + revisor.LastName,
		})
	}
	return revisorDropdownOptions, nil
}

//---------------------------------------------------------------------------------------------- nova polja

func (r *Resolver) RevisionPlansOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	response := dto.RevisionPlanOverviewResponse{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}
	page := params.Args["page"]
	size := params.Args["size"]
	year := params.Args["year"]

	input := dto.GetPlansInput{}
	if page != nil && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}
	if size != nil && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}
	if year != nil {
		year := year.(string)
		input.Year = &year
	}

	revisions, err := r.Repo.GetRevisionPlanList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Items = revisions.Data
	response.Total = revisions.Total

	return response, nil
}

func (r *Resolver) RevisionPlansDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	revision, err := r.Repo.GetRevisionPlanByID(id)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    *revision,
	}, nil
}

func (r *Resolver) RevisionPlanDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteRevisionPlan(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func (r *Resolver) RevisionPlanInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data dto.RevisionPlanItem
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		res, err := r.Repo.UpdateRevisionPlan(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Item = res
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateRevisionPlan(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		response.Item = res
		response.Message = "You created this item!"
	}

	return response, nil
}

//--------------------------------------------------------------------------

func buildRevisionItemResponse(r repository.MicroserviceRepositoryInterface, revision *structs.Revisions) (*dto.RevisionsOverviewItem, error) {

	revisiontype, _ := r.GetDropdownSettingByID(revision.RevisionType)
	/*if err != nil {
		return nil, errors.Wrap(err, "repo get dropdown setting by id")
	}*/

	var revisionType dto.DropdownSimple
	if revisiontype != nil {

		revisionType = dto.DropdownSimple{
			ID:    revisiontype.ID,
			Title: revisiontype.Title,
		}
	}

	var revisorDropdown []dto.DropdownSimple

	input := dto.RevisionRevisorFilter{
		RevisionID: &revision.ID,
	}

	revisors, err := r.GetRevisionRevisorList(&input)

	if err != nil {
		return nil, errors.Wrap(err, "repo get revision revisor list")
	}

	for _, revisorID := range revisors {
		revisor, err := r.GetUserProfileByID(revisorID.RevisorID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}

		revisorSingleDropdown := dto.DropdownSimple{
			ID:    revisor.ID,
			Title: revisor.FirstName + " " + revisor.LastName,
		}

		revisorDropdown = append(revisorDropdown, revisorSingleDropdown)
	}

	internalUnitDropdown := []dto.DropdownSimple{}
	externalUnitDropdown := &dto.DropdownSimple{}

	if revision.InternalRevisionsubject != nil {
		filt := dto.RevisionOrgUnitFilter{
			RevisionID: &revision.ID,
		}

		revisions, err := r.GetRevisionOrgUnitList(&filt)

		if err != nil {
			return nil, errors.Wrap(err, "repo get revision org unit list")
		}

		for _, revision := range revisions {
			orgUnit, err := r.GetOrganizationUnitByID(revision.OrganizationUnitID)
			if err != nil {
				return nil, errors.Wrap(err, "repo get organization unit by id")
			}

			orgUnitDropdown := dto.DropdownSimple{
				ID:    orgUnit.ID,
				Title: orgUnit.Title,
			}

			internalUnitDropdown = append(internalUnitDropdown, orgUnitDropdown)
		}
	}

	if revision.ExternalRevisionsubject != nil {
		organizationUnit, err := r.GetDropdownSettingByID(*revision.ExternalRevisionsubject)
		if err != nil {
			return nil, errors.Wrap(err, "repo get dropdown setting by id")
		}
		externalUnitDropdown.ID = organizationUnit.ID
		externalUnitDropdown.Title = organizationUnit.Title
	}

	var fileList []dto.FileDropdownSimple

	for i := range revision.FileIDs {
		var file dto.FileDropdownSimple

		res, _ := r.GetFileByID(revision.FileIDs[i])
		/*
			if err != nil {
				return nil, errors.Wrap(err, "repo get file by id")
			}
		*/

		if res != nil {
			file.ID = res.ID
			file.Name = res.Name
			file.Type = *res.Type
		}

		fileList = append(fileList, file)
	}

	revisionItem := &dto.RevisionsOverviewItem{
		ID:                      revision.ID,
		Title:                   revision.Title,
		PlanID:                  revision.PlanID,
		SerialNumber:            revision.SerialNumber,
		DateOfRevision:          revision.DateOfRevision,
		RevisionQuartal:         revision.RevisionQuartal,
		InternalRevisionsubject: &internalUnitDropdown,
		ExternalRevisionsubject: externalUnitDropdown,
		Revisor:                 revisorDropdown,
		RevisionType:            revisionType,
		Files:                   fileList,
		CreatedAt:               revision.CreatedAt,
		UpdatedAt:               revision.UpdatedAt,
	}

	var fileTipsList []dto.FileDropdownSimple

	for i := range revision.TipsFileIDs {
		var file dto.FileDropdownSimple
		res, _ := r.GetFileByID(revision.TipsFileIDs[i])

		/*if err != nil {
			return nil, errors.Wrap(err, "repo get file by id")
		}*/

		if res != nil {

			file.ID = res.ID
			file.Name = res.Name
			file.Type = *res.Type
		}

		fileTipsList = append(fileTipsList, file)
	}

	revisionItem.TipsFiles = fileTipsList

	return revisionItem, nil
}

func (r *Resolver) RevisionOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
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
	if page != nil && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}

	if size != nil && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}

	if revisionType != nil && revisionType.(int) > 0 {
		temp := revisionType.(int)
		input.RevisionType = &temp
	}

	if plan != nil && plan.(int) > 0 {
		temp := plan.(int)
		input.PlanID = &temp
	}

	revisions, err := r.Repo.GetRevisionsList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	var revisionsOrgUnit []*dto.RevisionOrgUnit
	revisionOrgUnit := false

	if internal != nil && internal.(int) > 0 {
		temp := internal.(int)
		filter := dto.RevisionOrgUnitFilter{
			OrganizationUnitID: &temp,
		}

		revisionsOrgUnit, err = r.Repo.GetRevisionOrgUnitList(&filter)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		revisionOrgUnit = true
	}

	var revisionsRevisor []*dto.RevisionRevisor
	revisionRevisor := false

	if revisor != nil && revisor.(int) > 0 {
		temp := revisor.(int)
		filter := dto.RevisionRevisorFilter{
			RevisorID: &temp,
		}

		revisionsRevisor, err = r.Repo.GetRevisionRevisorList(&filter)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		revisionRevisor = true
	}

	for _, revision := range revisions.Data {

		found := false
		if revisionRevisor && !revisionOrgUnit {
			for _, revisor := range revisionsRevisor {
				if revisor.RevisionID == revision.ID {
					found = true
				}
			}
		} else if revisionOrgUnit && !revisionRevisor {
			for _, orgUnit := range revisionsOrgUnit {
				if orgUnit.RevisionID == revision.ID {
					found = true
				}
			}
		} else if revisionRevisor && revisionOrgUnit {
			for _, revisor := range revisionsRevisor {
				if revisor.RevisionID == revision.ID {
					found = true
				}
			}
			if !found {
				continue
			}

			found = false
			for _, orgUnit := range revisionsOrgUnit {
				if orgUnit.RevisionID == revision.ID {
					found = true
				}
			}
		} else if !revisionRevisor && !revisionOrgUnit {
			found = true
		}

		if !found {
			continue
		}

		item, err := buildRevisionItemResponse(r.Repo, revision)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Items = append(response.Items, *item)
	}
	response.Total = revisions.Total

	revisorDropdownList, err := getRevisorListDropdown(r.Repo)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Revisors = revisorDropdownList

	return response, nil
}

func (r *Resolver) RevisionDetailResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	revision, err := r.Repo.GetRevisionsByID(id)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	item, err := buildRevisionItemResponse(r.Repo, revision)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	revisorDropdownList, err := getRevisorListDropdown(r.Repo)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.RevisionsDetailsResponse{
		Status:   "success",
		Message:  "Here's the list you asked for!",
		Item:     *item,
		Revisors: revisorDropdownList,
	}, nil
}

func (r *Resolver) RevisionsInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Revisions
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {

		input := dto.RevisionRevisorFilter{
			RevisionID: &itemID,
		}

		revisors, err := r.Repo.GetRevisionRevisorList(&input)

		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, revisor := range revisors {
			err := r.Repo.DeleteRevisionRevisor(revisor.ID)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}

		for _, revisor := range data.Revisor {
			inp := dto.RevisionRevisor{
				RevisionID: itemID,
				RevisorID:  revisor,
			}

			err := r.Repo.CreateRevisionRevisor(&inp)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}

		if data.InternalRevisionsubject != nil {
			filt := dto.RevisionOrgUnitFilter{
				RevisionID: &itemID,
			}

			revisions, err := r.Repo.GetRevisionOrgUnitList(&filt)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}

			for _, revision := range revisions {
				err = r.Repo.DeleteRevisionOrgUnit(revision.ID)

				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}

			for _, orgUnit := range data.InternalRevisionsubject {
				dataRevision := dto.RevisionOrgUnit{
					OrganizationUnitID: orgUnit,
					RevisionID:         itemID,
				}

				err := r.Repo.CreateRevisionOrgUnit(&dataRevision)
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}

		res, err := r.Repo.UpdateRevisions(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		item, err := buildRevisionItemResponse(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateRevisions(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}

		for _, revisor := range data.Revisor {
			inp := dto.RevisionRevisor{
				RevisionID: res.ID,
				RevisorID:  revisor,
			}

			err := r.Repo.CreateRevisionRevisor(&inp)

			if err != nil {
				_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
				return errors.HandleAPPError(err)
			}
		}

		if data.InternalRevisionsubject != nil {
			for _, orgUnit := range data.InternalRevisionsubject {
				dataRevision := dto.RevisionOrgUnit{
					OrganizationUnitID: orgUnit,
					RevisionID:         res.ID,
				}

				err := r.Repo.CreateRevisionOrgUnit(&dataRevision)
				if err != nil {
					_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
					return errors.HandleAPPError(err)
				}
			}
		}

		item, err := buildRevisionItemResponse(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Item = item
	}
	response.Message = "You created this item!"

	return response, nil
}

func (r *Resolver) RevisionsDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteRevisions(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

//-------------------------------------------------------------------------

func buildRevisionTipItemResponse(r repository.MicroserviceRepositoryInterface, revision *structs.RevisionTips) (*dto.RevisionTipsOverviewItem, error) {

	revisorDropdown := structs.SettingsDropdown{}

	if revision.UserProfileID != nil {
		revisor, err := r.GetUserProfileByID(*revision.UserProfileID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}

		revisorDropdown = structs.SettingsDropdown{
			ID:    revisor.ID,
			Title: revisor.FirstName + " " + revisor.LastName,
		}
	}

	var files []dto.FileDropdownSimple

	for i := range revision.FileIDs {
		res, _ := r.GetFileByID(revision.FileIDs[i])
		/*
			if err != nil {
				return nil, errors.Wrap(err, "repo get file by id")
			}
		*/

		if res != nil {
			files = append(files, dto.FileDropdownSimple{
				ID:   res.ID,
				Name: res.Name,
				Type: *res.Type,
			})
		}
	}
	revisionTipItem := &dto.RevisionTipsOverviewItem{
		ID:                revision.ID,
		RevisionID:        revision.RevisionID,
		UserProfile:       revisorDropdown,
		DateOfAccept:      revision.DateOfAccept,
		DueDate:           revision.DueDate,
		EndDate:           revision.EndDate,
		DateOfReject:      revision.DateOfReject,
		DateOfExecution:   revision.DateOfExecution,
		Recommendation:    revision.Recommendation,
		RevisionPriority:  revision.RevisionPriority,
		Status:            revision.Status,
		ResponsiblePerson: revision.ResponsiblePerson,
		Files:             files,
		CreatedAt:         revision.CreatedAt,
		UpdatedAt:         revision.UpdatedAt,
	}

	return revisionTipItem, nil
}

func (r *Resolver) RevisionTipsOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	response := dto.RevisionTipsOverviewResponse{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}
	page := params.Args["page"]
	size := params.Args["size"]
	revision := params.Args["revision_id"]

	input := dto.GetRevisionTipFilter{}
	if page != nil && page.(int) > 0 {
		pageNum := page.(int)
		input.Page = &pageNum
	}

	if size != nil && size.(int) > 0 {
		sizeNum := size.(int)
		input.Size = &sizeNum
	}

	if revision != nil && revision.(int) > 0 {
		temp := revision.(int)
		input.RevisionID = &temp
	}

	revisions, err := r.Repo.GetRevisionTipsList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, revision := range revisions.Data {
		item, err := buildRevisionTipItemResponse(r.Repo, revision)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Items = append(response.Items, *item)
	}
	response.Total = revisions.Total

	revisorDropdownList, err := getRevisorListDropdown(r.Repo)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Revisors = revisorDropdownList

	return response, nil
}

func (r *Resolver) RevisionTipsDetailsResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"].(int)

	revision, err := r.Repo.GetRevisionTipByID(id)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	item, err := buildRevisionTipItemResponse(r.Repo, revision)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	return dto.ResponseSingle{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Item:    *item,
	}, nil
}

func (r *Resolver) RevisionTipsInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.RevisionTips
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		res, err := r.Repo.UpdateRevisionTips(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildRevisionTipItemResponse(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateRevisionTips(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildRevisionTipItemResponse(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Item = item
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) RevisionTipsDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteRevisionTips(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

//-------------------------------------------------------------------------

func buildRevisionTipImplementationItemResponse(r repository.MicroserviceRepositoryInterface, tipImpl *structs.RevisionTipImplementations) (*dto.RevisionTipImplementationsOverviewItem, error) {
	revisorDropdown := dto.DropdownSimple{}

	if tipImpl.RevisorID != nil {
		revisor, err := r.GetUserProfileByID(*tipImpl.RevisorID)
		if err != nil {
			return nil, errors.Wrap(err, "repo get user profile by id")
		}

		revisorDropdown = dto.DropdownSimple{
			ID:    revisor.ID,
			Title: revisor.FirstName + " " + revisor.LastName,
		}
	}

	var files []dto.FileDropdownSimple

	for i := range tipImpl.FileIDs {
		res, _ := r.GetFileByID(tipImpl.FileIDs[i])

		if res != nil {
			files = append(files, dto.FileDropdownSimple{
				ID:   res.ID,
				Name: res.Name,
				Type: *res.Type,
			})
		}
	}

	revisionTipImplementationItem := &dto.RevisionTipImplementationsOverviewItem{
		ID:                     tipImpl.ID,
		TipID:                  tipImpl.TipID,
		Status:                 tipImpl.Status,
		NewDueDate:             tipImpl.NewDueDate,
		NewDateOfExecution:     tipImpl.NewDateOfExecution,
		ReasonsForNonExecuting: tipImpl.ReasonsForNonExecuting,
		Revisor:                revisorDropdown,
		Documents:              tipImpl.Documents,
		Files:                  files,
		CreatedAt:              tipImpl.CreatedAt,
		UpdatedAt:              tipImpl.UpdatedAt,
	}

	return revisionTipImplementationItem, nil
}

func (r *Resolver) RevisionTipImplementationOverviewResolver(params graphql.ResolveParams) (interface{}, error) {
	response := dto.RevisionTipImplementationsOverviewResponse{
		Status:  "success",
		Message: "Here's the list you asked for!",
	}

	tipID := params.Args["tip_id"]

	input := dto.GetRevisionTipImplementationFilter{}
	if tipID != nil && tipID.(int) > 0 {
		temp := tipID.(int)
		input.TipID = &temp
	}

	tipImplementations, err := r.Repo.GetRevisionTipImplementationList(&input)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	for _, tipImpl := range tipImplementations.Data {
		item, err := buildRevisionTipImplementationItemResponse(r.Repo, tipImpl)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Items = append(response.Items, *item)
	}

	revisorDropdownList, err := getRevisorListDropdown(r.Repo)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	response.Revisors = revisorDropdownList

	return response, nil
}

func (r *Resolver) RevisionTipImplementationInsertResolver(params graphql.ResolveParams) (any, error) {
	var data structs.RevisionTipImplementations
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		res, err := r.Repo.UpdateRevisionTipImplementation(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildRevisionTipImplementationItemResponse(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Item = item
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateRevisionTipImplementation(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		item, err := buildRevisionTipImplementationItemResponse(r.Repo, res)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Item = item
		response.Message = "You created this item!"
	}

	return response, nil
}
