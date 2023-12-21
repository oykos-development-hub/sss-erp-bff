package resolvers

import (
	"bff/internal/api/dto"
	"bff/internal/api/errors"
	"bff/shared"
	"bff/structs"
	"encoding/json"

	"github.com/graphql-go/graphql"
)

func (r *Resolver) SettingsDropdownResolver(params graphql.ResolveParams) (interface{}, error) {
	id := params.Args["id"]
	entity := params.Args["entity"].(string)
	page := params.Args["page"]
	size := params.Args["size"]
	search, searchOk := params.Args["search"].(string)

	var (
		items []structs.SettingsDropdown
		total int
	)

	if id != nil && shared.IsInteger(id) && id != 0 {
		setting, err := r.Repo.GetDropdownSettingById(id.(int))
		if err != nil {
			return errors.HandleAPIError(err)
		}
		items = []structs.SettingsDropdown{*setting}
		total = 1
	} else {
		input := dto.GetSettingsInput{}
		if shared.IsInteger(page) && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if shared.IsInteger(size) && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
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

	itemId := data.Id

	response := dto.ResponseSingle{
		Status: "success",
	}

	if shared.IsInteger(itemId) && itemId != 0 {
		itemRes, err := r.Repo.UpdateDropdownSettings(itemId, &data)
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
	itemId := params.Args["id"].(int)

	err := r.Repo.DeleteDropdownSettings(itemId)
	if err != nil {
		return errors.HandleAPIError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil
}
