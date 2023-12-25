package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) SettingsDropdownResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	entity := params.Args["entity"].(string)
	page, pageOk := params.Args["page"].(int)
	size, sizeOk := params.Args["size"].(int)
	search, searchOk := params.Args["search"].(string)

	var (
		items []structs.SettingsDropdown
		total int
	)

	if id != nil && id != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}
		items = []structs.SettingsDropdown{*setting}
		total = 1
	} else {
		input := dto.GetSettingsInput{}
		if pageOk && page > 0 {
			input.Page = &page
		}
		if sizeOk && size > 0 {
			input.Size = &size
		}
		if searchOk && search != "" {
			input.Search = &search
		}
		input.Entity = entity

		if value, ok := params.Args["value"].(string); ok && value != "" {
			input.Value = &value
		}

		res, err := r.Repo.GetDropdownSettings(&input)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		items = res.Data
		total = res.Total
	}

	return dto.Response{
		Status:  "success",
		Message: "Here's the list you asked for!",
		Items:   items,
		Total:   total,
	}, nil
}

func (r *Resolver) SettingsDropdownInsertResolver(params graphql.ResolveParams) (interface{}, error) {
	var data structs.SettingsDropdown
	dataBytes, _ := json.Marshal(params.Args["data"])

	_ = json.Unmarshal(dataBytes, &data)

	itemID := data.ID

	response := dto.ResponseSingle{
		Status: "success",
	}

	if itemID != 0 {
		itemRes, err := r.Repo.UpdateDropdownSettings(itemID, &data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You updated this item!"
		response.Item = itemRes

	} else {
		itemRes, err := r.Repo.CreateDropdownSettings(&data)
		if err != nil {
			return errors.HandleAPIError(err)
		}
		response.Message = "You created this item!"
		response.Item = itemRes
	}

	return response, nil

}

func (r *Resolver) SettingsDropdownDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	itemID := params.Args["id"].(int)

	err := r.Repo.DeleteDropdownSettings(itemID)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
