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
	page := params.Args["page"]
	size := params.Args["size"]
	search, searchOk := params.Args["search"].(string)
	parentID, parentIDOK := params.Args["parent_id"].(int)

	var (
		items []structs.SettingsDropdown
		total int
	)

	if id != nil && id != 0 {
		setting, err := r.Repo.GetDropdownSettingByID(id.(int))
		if err != nil {
			return errors.HandleAPPError(err)
		}
		items = []structs.SettingsDropdown{*setting}
		total = 1
	} else {
		input := dto.GetSettingsInput{}
		if page != nil && page.(int) > 0 {
			pageNum := page.(int)
			input.Page = &pageNum
		}
		if size != nil && size.(int) > 0 {
			sizeNum := size.(int)
			input.Size = &sizeNum
		}
		if searchOk && search != "" {
			input.Search = &search
		}

		if parentIDOK && parentID != 0 {
			input.ParentID = &parentID
		}
		input.Entity = entity

		if value, ok := params.Args["value"].(string); ok && value != "" {
			input.Value = &value
		}

		res, err := r.Repo.GetDropdownSettings(&input)
		if err != nil {
			return errors.HandleAPPError(err)
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
			return errors.HandleAPPError(err)
		}
		response.Message = "You updated this item!"
		response.Item = itemRes

	} else {
		itemRes, err := r.Repo.CreateDropdownSettings(&data)
		if err != nil {
			return errors.HandleAPPError(err)
		}
		response.Message = "You created this item!"
		response.Item = itemRes
	}

	return response, nil

}

func (r *Resolver) SettingsDropdownDeleteResolver(params graphql.ResolveParams) (interface{}, error) {
	/*itemID := params.Args["id"].(int)

	err := r.Repo.DeleteDropdownSettings(itemID)
	if err != nil {
		return errors.HandleAPPError(err)
	}

	return dto.ResponseSingle{
		Status:  "success",
		Message: "You deleted this item!",
	}, nil*/
	return dto.ResponseSingle{
		Status:  "failed",
		Message: "You can not delete this item!",
	}, nil
}
