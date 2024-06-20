package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) AbsentTypeResolver(_ graphql.ResolveParams) (interface{}, error) {
	absentTypesAll, err := r.Repo.GetAbsentTypes()
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   absentTypesAll.Data,
		Total:   absentTypesAll.Total,
	}, nil
}

func (r *Resolver) AbsentTypeInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var err error
	var data structs.AbsentType

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
		item, err := r.Repo.UpdateAbsentType(itemID, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		response.Message = "You updated this item!"
		response.Item = item
	} else {
		item, err := r.Repo.CreateAbsentType(&data)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		response.Message = "You created this item!"
		response.Item = item
	}

	return response, nil
}

func (r *Resolver) AbsentTypeDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"]

	err := r.Repo.DeleteAbsentType(itemID.(int))
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
