package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) JobTenderTypesResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	search := params.Args["search"]
	var items []*structs.JobTenderTypes

	if id != nil && id != 0 {
		tenderType, err := r.Repo.GetTenderType(id.(int))
		if err != nil {
			return errors.HandleAPPError(err)
		}
		items = append(items, tenderType)
	} else {
		input := dto.GetJobTenderTypeInputMS{}
		if search != nil {
			search := search.(string)
			input.Search = &search
		}
		tenderTypes, err := r.Repo.GetTenderTypeList(&input)
		if err != nil {
			return errors.HandleAPPError(err)
		}
		items = tenderTypes
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
	}, nil
}

func (r *Resolver) JobTenderTypeInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.JobTenderTypes
	dataBytes, _ := json.Marshal(params.Args["data"])
	response := dto.ResponseSingle{
		Status: "success",
	}

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID
	if itemID != 0 {
		res, err := r.Repo.UpdateJobTenderType(itemID, &data)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		response.Item = res
		response.Message = "You updated this item!"
	} else {
		res, err := r.Repo.CreateJobTenderType(&data)
		if err != nil {
			return errors.HandleAPPError(err)
		}

		response.Item = res
		response.Message = "You created this item!"
	}

	return response, nil
}

func (r *Resolver) JobTenderTypeDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteJobTenderType(itemID)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
