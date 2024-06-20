package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserProfileForeignerResolver(params graphql.ResolveParams) (interface{}, error) {
	profileID := params.Args["user_profile_id"].(int)

	UserProfilesData, err := r.Repo.GetEmployeeForeigners(profileID)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the item you asked for!",
		Items:   UserProfilesData,
	}, nil
}

func (r *Resolver) UserProfileForeignerInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error

	var data structs.Foreigners
	response := dto.ResponseSingle{
		Status: "success",
	}

	dataBytes, _ := json.Marshal(params.Args["data"])

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	itemID := data.ID
	if itemID != 0 {
		item, err := r.Repo.UpdateEmployeeForeigner(itemID, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}
		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := r.Repo.CreateEmployeeForeigner(&data)
		if err != nil {
			return errors.HandleAPPError(err)
		}
		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func (r *Resolver) UserProfileForeignerDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]

	err := r.Repo.DeleteForeigner(itemID.(int))
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
