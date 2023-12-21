package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/shared"
	"bff/structs"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) UserProfileForeignerResolver(params graphql.ResolveParams) (interface{}, error) {
	profileId := params.Args["user_profile_id"].(int)

	UserProfilesData, err := r.Repo.GetEmployeeForeigners(profileId)
	if err != nil {
		return errors.HandleAPIError(err)
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
		fmt.Printf("Error JSON parsing because of this error - %s.\n", err)
		return errors.ErrorResponse("Error updating settings data"), nil
	}

	itemId := data.Id
	if shared.IsInteger(itemId) && itemId != 0 {
		item, err := r.Repo.UpdateEmployeeForeigner(itemId, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := r.Repo.CreateEmployeeForeigner(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func (r *Resolver) UserProfileForeignerDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemId := params.Args["id"]

	err := r.Repo.DeleteForeigner(itemId.(int))
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
