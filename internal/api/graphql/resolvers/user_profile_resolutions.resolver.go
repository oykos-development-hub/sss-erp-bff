package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/internal/api/repository"
	"bff/shared"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserProfileResolutionResolver(params graphql.ResolveParams) (interface{}, error) {
	userProfileId := params.Args["user_profile_id"].(int)

	resolutions, err := r.Repo.GetEmployeeResolutions(userProfileId, nil)
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

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		resolution, err := r.Repo.UpdateResolution(itemId, &data)
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
	itemId := params.Args["id"].(int)

	err := r.Repo.DeleteResolution(itemId)
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
		resItem, err := buildResolutionResItem(r, item)
		if err != nil {
			return nil, err
		}
		resItemList = append(resItemList, resItem)
	}
	return
}

func buildResolutionResItem(r repository.MicroserviceRepositoryInterface, item *structs.Resolution) (*dto.Resolution, error) {
	userProfile, err := r.GetUserProfileById(item.UserProfileId)
	if err != nil {
		return nil, err
	}
	resolutionType, err := r.GetDropdownSettingById(item.ResolutionTypeId)
	if err != nil {
		return nil, err
	}

	return &dto.Resolution{
		Id:                item.Id,
		ResolutionPurpose: item.ResolutionPurpose,
		UserProfile: dto.DropdownSimple{
			Id:    userProfile.Id,
			Title: userProfile.GetFullName(),
		},
		ResolutionType: dto.DropdownSimple{
			Id:    resolutionType.Id,
			Title: resolutionType.Title,
		},
		IsAffect:    item.IsAffect,
		DateOfStart: item.DateOfStart,
		DateOfEnd:   item.DateOfEnd,
		Value:       item.Value,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}, nil
}
