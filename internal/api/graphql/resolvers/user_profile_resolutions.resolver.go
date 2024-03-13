package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserProfileResolutionResolver(params graphql.ResolveParams) (interface{}, error) {
	userProfileID := params.Args["user_profile_id"].(int)

	resolutions, err := r.Repo.GetEmployeeResolutions(userProfileID, nil)
	if err != nil {
		return errors.HandleAPIError(err)
	}
	resolutonResItemList, err := buildResolutionResponseItemList(r.Repo, resolutions)
	if err != nil {
		return errors.HandleAPIError(err)
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

	itemID := data.ID
	if itemID != 0 {
		resolution, err := r.Repo.UpdateResolution(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resolutionResItem, err := buildResolutionResItem(r.Repo, resolution)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = resolutionResItem
		response.Message = "You updated this item!"
	} else {
		resolution, err := r.Repo.CreateResolution(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		resolutionResItem, err := buildResolutionResItem(r.Repo, resolution)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Item = resolutionResItem
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) UserProfileResolutionDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteResolution(itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}

func buildResolutionResponseItemList(r repository.MicroserviceRepositoryInterface, items []*structs.Resolution) (resItemList []*dto.Resolution, err error) {
	for _, item := range items {
		if item.Value != "" {
			continue
		}
		resItem, err := buildResolutionResItem(r, item)
		if err != nil {
			return nil, err
		}
		resItemList = append(resItemList, resItem)
	}
	return
}

func buildResolutionResItem(r repository.MicroserviceRepositoryInterface, item *structs.Resolution) (*dto.Resolution, error) {
	/*if item.Value != config.VacationTypeValueResolutionType {
		return nil, nil
	}*/

	userProfile, err := r.GetUserProfileByID(item.UserProfileID)
	if err != nil {
		return nil, err
	}
	resolutionType, err := r.GetDropdownSettingByID(item.ResolutionTypeID)
	if err != nil {
		return nil, err
	}

	var file dto.FileDropdownSimple

	if item.FileID > 0 {
		res, err := r.GetFileByID(item.FileID)

		if err != nil {
			return nil, err
		}

		file.ID = res.ID
		file.Name = res.Name
		file.Type = *res.Type
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
