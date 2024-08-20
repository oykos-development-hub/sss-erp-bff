package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserProfileResolutionResolver(params graphql.ResolveParams) (interface{}, error) {
	userProfileID := params.Args["user_profile_id"].(int)

	resolutions, err := r.Repo.GetEmployeeResolutions(userProfileID, nil)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}
	resolutonResItemList, err := buildResolutionResponseItemList(r.Repo, resolutions)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   resolutonResItemList,
	}, nil
}

func (r *Resolver) UserProfileResolutionInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.Resolution
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	currentTime := time.Now()

	currentYear := currentTime.Year()

	if data.DateOfStart == "" {
		data.DateOfStart = fmt.Sprintf("%d-01-01T00:00:00Z", currentYear)
	}

	if data.DateOfEnd == nil || *data.DateOfEnd == "" {
		dateOfEnd := fmt.Sprintf("%d-12-31T00:00:00Z", currentYear)
		data.DateOfEnd = &dateOfEnd
	}

	itemID := data.ID
	if itemID != 0 {
		resolution, err := r.Repo.UpdateResolution(params.Context, itemID, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		resolutionResItem, err := buildResolutionResItem(r.Repo, resolution)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Item = resolutionResItem
		response.Message = "You updated this item!"
	} else {
		resolution, err := r.Repo.CreateResolution(params.Context, &data)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		resolutionResItem, err := buildResolutionResItem(r.Repo, resolution)
		if err != nil {
			_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
			return errors.HandleAPPError(err)
		}
		response.Item = resolutionResItem
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) UserProfileResolutionDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteResolution(params.Context, itemID)
	if err != nil {
		_ = r.Repo.CreateErrorLog(structs.ErrorLogs{Error: err.Error()})
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildResolutionResponseItemList(r repository.MicroserviceRepositoryInterface, items []*structs.Resolution) (resItemList []*dto.Resolution, err error) {
	for _, item := range items {

		resItem, err := buildResolutionResItem(r, item)
		if err != nil {
			return nil, errors.Wrap(err, "build resolution res item")
		}
		resItemList = append(resItemList, resItem)
	}
	return
}

func buildResolutionResItem(r repository.MicroserviceRepositoryInterface, item *structs.Resolution) (*dto.Resolution, error) {

	userProfile, err := r.GetUserProfileByID(item.UserProfileID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get user profile by id")
	}
	resolutionType, err := r.GetDropdownSettingByID(item.ResolutionTypeID)
	if err != nil {
		return nil, errors.Wrap(err, "repo get dropdown setting by id")
	}

	var file dto.FileDropdownSimple

	if item.FileID > 0 {
		res, _ := r.GetFileByID(item.FileID)
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

	}
	return &dto.Resolution{
		ID:                item.ID,
		ResolutionPurpose: item.ResolutionPurpose,
		UserProfile: dto.DropdownSimple{
			ID:    userProfile.ID,
			Title: userProfile.GetFullName(),
		},
		ResolutionType: dto.DropdownSimple{
			ID:    resolutionType.ID,
			Title: resolutionType.Title,
		},
		IsAffect:    item.IsAffect,
		DateOfStart: item.DateOfStart,
		DateOfEnd:   item.DateOfEnd,
		Value:       item.Value,
		File:        file,
		Year:        item.Year,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}, nil
}
